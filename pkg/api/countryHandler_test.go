package api

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/anytimesoon/eurovision-party/pkg/api/enum"
	"github.com/anytimesoon/eurovision-party/pkg/data/dao"
	"github.com/anytimesoon/eurovision-party/pkg/service"
	"github.com/anytimesoon/eurovision-party/pkg/service/dto"
	"github.com/gorilla/mux"
	"github.com/timshannon/bolthold"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestCountryHandler_FindAllCountries(t *testing.T) {
	type fields struct {
		Service service.CountryService
	}
	type args struct {
		resp http.ResponseWriter
		req  *http.Request
	}
	type expected struct {
		statusCode int
		countries  []string
	}
	tests := []struct {
		name     string
		fields   fields
		args     args
		expected expected
	}{
		{
			name: "successful get all countries",
			fields: fields{
				Service: newTestCountryService(),
			},
			args: args{
				resp: httptest.NewRecorder(),
				req:  httptest.NewRequest(http.MethodGet, "/api/countries", nil),
			},
			expected: expected{
				statusCode: http.StatusOK,
				countries:  countryNames,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ch := &CountryHandler{
				Service: tt.fields.Service,
			}
			ch.GetAllCountries(tt.args.resp, tt.args.req)
			result := tt.args.resp.(*httptest.ResponseRecorder).Result()
			defer func(Body io.ReadCloser) {
				err := Body.Close()
				if err != nil {
					panic(err)
				}
			}(result.Body)

			if result.StatusCode != tt.expected.statusCode {
				t.Errorf("GetAllCountries return status code = %v, want %v", result.StatusCode, http.StatusOK)
			}

			resultBody, err := io.ReadAll(result.Body)
			if err != nil {
				panic(err)
			}

			var resultDto dto.ApiPayload[[]dto.Country]
			err = json.Unmarshal(resultBody, &resultDto)
			if err != nil {
				panic(err)
			}

			for _, expectedCountry := range tt.expected.countries {
				found := false
				for _, resultCountry := range resultDto.Body {
					if resultCountry.Slug == expectedCountry {
						found = true
						break
					}
				}
				if !found {
					t.Errorf("Country %s not found in response", expectedCountry)
				}
			}

		})
	}
}

func TestCountryHandler_FindOneCountry(t *testing.T) {
	type fields struct {
		Service service.CountryService
	}
	type args struct {
		resp http.ResponseWriter
		req  *http.Request
	}
	type expected struct {
		statusCode int
		country    string
	}
	tests := []struct {
		name     string
		fields   fields
		args     args
		expected expected
	}{
		{
			name: "successful get one country",
			fields: fields{
				Service: newTestCountryService(),
			},
			args: args{
				resp: httptest.NewRecorder(),
				req:  httptest.NewRequest(http.MethodGet, "/api/countries/Austria", nil),
			},
			expected: expected{
				statusCode: http.StatusOK,
				country:    countryNames[0],
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ch := &CountryHandler{
				Service: tt.fields.Service,
			}
			tt.args.req = mux.SetURLVars(tt.args.req, map[string]string{"slug": countryNames[0]})
			ch.GetOneCountry(tt.args.resp, tt.args.req)
			result := tt.args.resp.(*httptest.ResponseRecorder).Result()
			defer func(Body io.ReadCloser) {
				err := Body.Close()
				if err != nil {
					panic(err)
				}
			}(result.Body)

			if result.StatusCode != tt.expected.statusCode {
				t.Errorf("GetAllCountries return status code = %v, want %v", result.StatusCode, http.StatusOK)
			}

			resultBody, err := io.ReadAll(result.Body)
			if err != nil {
				panic(err)
			}

			var resultDto dto.ApiPayload[dto.Country]
			err = json.Unmarshal(resultBody, &resultDto)
			if err != nil {
				panic(err)
			}

			if resultDto.Body.Slug != tt.expected.country {
				t.Errorf("Country %s not found in response", tt.expected.country)
			}
		})
	}
}

func TestCountryHandler_Participating(t *testing.T) {
	participatingCountries := []string{countryNames[0], countryNames[1], countryNames[2]}
	type fields struct {
		Service service.CountryService
	}
	type args struct {
		resp http.ResponseWriter
		req  *http.Request
	}
	type expected struct {
		statusCode int
		countries  []string
	}
	tests := []struct {
		name     string
		fields   fields
		args     args
		expected expected
	}{
		{
			name: "successful get participating countries",
			fields: fields{
				Service: newTestCountryService(),
			},
			args: args{
				resp: httptest.NewRecorder(),
				req:  httptest.NewRequest(http.MethodGet, "/api/countries", nil),
			},
			expected: expected{
				statusCode: http.StatusOK,
				countries:  participatingCountries,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ch := &CountryHandler{
				Service: tt.fields.Service,
			}
			fmt.Printf("%T: %v\n", participatingCountries, participatingCountries)
			err := testDB.UpdateMatching(&dao.Country{},
				bolthold.Where("Name").In(participatingCountries[0], participatingCountries[1], participatingCountries[2]),
				func(record interface{}) error {
					country := record.(*dao.Country)
					country.Participating = true
					return nil
				})
			if err != nil {
				panic(err)
			}
			ch.GetParticipatingCountries(tt.args.resp, tt.args.req)
			result := tt.args.resp.(*httptest.ResponseRecorder).Result()
			defer func(Body io.ReadCloser) {
				err := Body.Close()
				if err != nil {
					panic(err)
				}
			}(result.Body)

			if result.StatusCode != tt.expected.statusCode {
				t.Errorf("GetParticipatingCountries return status code = %v, want %v", result.StatusCode, http.StatusOK)
			}

			resultBody, err := io.ReadAll(result.Body)
			if err != nil {
				panic(err)
			}

			var resultDto dto.ApiPayload[[]dto.Country]
			err = json.Unmarshal(resultBody, &resultDto)
			if err != nil {
				panic(err)
			}

			for _, expectedCountry := range tt.expected.countries {
				found := false
				for _, resultCountry := range resultDto.Body {
					if resultCountry.Slug == expectedCountry {
						found = true
						break
					}
				}
				if !found {
					t.Errorf("Country %s not found in response", expectedCountry)
				}
			}
		})
	}
}

func TestCountryHandler_UpdateCountry(t *testing.T) {
	type fields struct {
		Service service.CountryService
	}
	type args struct {
		resp    http.ResponseWriter
		req     *http.Request
		authLvl enum.AuthLvl
	}
	type expected struct {
		statusCode int
		country    dao.Country
	}
	tests := []struct {
		name     string
		fields   fields
		args     args
		expected expected
	}{
		{
			name: "successful update one country",
			fields: fields{
				Service: newTestCountryService(),
			},
			args: args{
				resp: httptest.NewRecorder(),
				req: httptest.NewRequest(http.MethodPut, "/api/countries", strings.NewReader(fmt.Sprintf(
					`{
					"name": "%s",
					"slug": "%s",
					"bandName": "band",
					"songName": "song",
					"flag": "",
					"participating": true
                }`, countryNames[0], countryNames[0]))),
				authLvl: enum.ADMIN,
			},
			expected: expected{
				statusCode: http.StatusOK,
				country: dao.Country{
					Name:          countryNames[0],
					Slug:          countryNames[0],
					BandName:      "band",
					SongName:      "song",
					Flag:          "",
					Participating: true,
				},
			},
		},
		{
			name: "fail update one country if user is not admin",
			fields: fields{
				Service: newTestCountryService(),
			},
			args: args{
				resp: httptest.NewRecorder(),
				req: httptest.NewRequest(http.MethodPut, "/api/countries", strings.NewReader(fmt.Sprintf(
					`{
					"name": "%s",
					"slug": "%s",
					"bandName": "band",
					"songName": "song",
					"flag": "",
					"participating": true
                }`, countryNames[0], countryNames[0]))),
				authLvl: enum.NONE,
			},
			expected: expected{
				statusCode: http.StatusUnauthorized,
				country: dao.Country{
					Name:          countryNames[0],
					Slug:          countryNames[0],
					BandName:      "band",
					SongName:      "song",
					Flag:          "",
					Participating: false,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ch := &CountryHandler{
				Service: tt.fields.Service,
			}

			mockAuth := dto.Auth{
				AuthLvl: tt.args.authLvl,
			}
			tt.args.req = tt.args.req.WithContext(context.WithValue(context.Background(), "auth", mockAuth))

			ch.UpdateCountry(tt.args.resp, tt.args.req)
			result := tt.args.resp.(*httptest.ResponseRecorder).Result()
			defer func(Body io.ReadCloser) {
				err := Body.Close()
				if err != nil {
					panic(err)
				}
			}(result.Body)

			if result.StatusCode != tt.expected.statusCode {
				t.Errorf("UpdateCountry return status code = %v, want %v", result.StatusCode, http.StatusOK)
			}

			var updatedCountry dao.Country
			err := testDB.Get(countryNames[0], &updatedCountry)
			if err != nil {
				panic(err)
			}

			err = testDB.Upsert(countryNames[0], dao.Country{
				Name:          countryNames[0],
				Slug:          countryNames[0],
				BandName:      "",
				SongName:      "",
				Flag:          "",
				Participating: false,
			})
			if err != nil {
				panic(err)
			}
		})
	}
}
