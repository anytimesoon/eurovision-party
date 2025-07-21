package api

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/anytimesoon/eurovision-party/conf"
	"github.com/anytimesoon/eurovision-party/pkg/api/enum"
	"github.com/anytimesoon/eurovision-party/pkg/api/enum/chatMsgType"
	"github.com/anytimesoon/eurovision-party/pkg/data/dao"
	"github.com/anytimesoon/eurovision-party/pkg/service"
	"github.com/anytimesoon/eurovision-party/pkg/service/dto"
	"github.com/gorilla/mux"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"
)

func TestVoteHandler_GetResults(t *testing.T) {
	type fields struct {
		Service service.VoteService
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
			name: "successful get results for all users",
			fields: fields{
				Service: newTestVoteService(),
			},
			args: args{
				resp: httptest.NewRecorder(),
				req:  httptest.NewRequest(http.MethodGet, "/api/vote/results", nil),
			},
			expected: expected{
				statusCode: http.StatusOK,
				countries:  countryNames,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			vh := VoteHandler{
				Service: tt.fields.Service,
			}
			vh.GetResults(tt.args.resp, tt.args.req)
			result := tt.args.resp.(*httptest.ResponseRecorder).Result()
			defer func(Body io.ReadCloser) {
				err := Body.Close()
				if err != nil {
					panic(err)
				}
			}(result.Body)

			if result.StatusCode != tt.expected.statusCode {
				t.Errorf("Login return status code = %v, want %v", result.StatusCode, http.StatusOK)
			}

			resultBody, err := io.ReadAll(result.Body)
			if err != nil {
				panic(err)
			}

			var resultDto dto.ApiPayload[[]dto.Result]
			err = json.Unmarshal(resultBody, &resultDto)
			if err != nil {
				panic(err)
			}

			returnedResult := resultDto.Body
			returnedResultCount := len(returnedResult)

			if returnedResultCount != len(tt.expected.countries) {
				t.Errorf("Returned result count = %v, want %v", returnedResultCount, len(tt.expected.countries))
			}

			for i := 0; i < returnedResultCount; i++ {
				forwardIdx := i
				reverseIdx := len(tt.expected.countries) - 1 - i

				if returnedResult[forwardIdx].CountrySlug != tt.expected.countries[reverseIdx] {
					t.Errorf("Returned result country slug = %v, want %v", returnedResult[forwardIdx].CountrySlug, tt.expected.countries[reverseIdx])
				}

				expectedScore := reverseIdx * 2

				if returnedResult[forwardIdx].Total != expectedScore*4 {
					t.Errorf("Returned result total = %v, want %v", returnedResult[forwardIdx].Total, expectedScore*4)
				}

				if returnedResult[forwardIdx].Costume != expectedScore {
					t.Errorf("Returned result costume = %v, want %v", returnedResult[forwardIdx].Costume, expectedScore)
				}

				if returnedResult[forwardIdx].Props != expectedScore {
					t.Errorf("Returned result props = %v, want %v", returnedResult[forwardIdx].Props, expectedScore)
				}

				if returnedResult[forwardIdx].Song != expectedScore {
					t.Errorf("Returned result song = %v, want %v", returnedResult[forwardIdx].Song, expectedScore)
				}

				if returnedResult[forwardIdx].Props != expectedScore {
					t.Errorf("Returned result props = %v, want %v", returnedResult[forwardIdx].Props, expectedScore)
				}
			}
		})
	}
}

func TestVoteHandler_GetResultsByUser(t *testing.T) {
	type fields struct {
		Service service.VoteService
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
			name: "successful get results for one user",
			fields: fields{
				Service: newTestVoteService(),
			},
			args: args{
				resp: httptest.NewRecorder(),
				req:  httptest.NewRequest(http.MethodGet, "/api/vote/results/"+adminUserId.String(), nil),
			},
			expected: expected{
				statusCode: http.StatusOK,
				countries:  countryNames,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			vh := VoteHandler{
				Service: tt.fields.Service,
			}
			tt.args.req = mux.SetURLVars(tt.args.req, map[string]string{"userId": adminUserId.String()})
			vh.GetResultsByUser(tt.args.resp, tt.args.req)
			result := tt.args.resp.(*httptest.ResponseRecorder).Result()
			defer func(Body io.ReadCloser) {
				err := Body.Close()
				if err != nil {
					panic(err)
				}
			}(result.Body)

			if result.StatusCode != tt.expected.statusCode {
				t.Errorf("Login return status code = %v, want %v", result.StatusCode, http.StatusOK)
			}

			resultBody, err := io.ReadAll(result.Body)
			if err != nil {
				panic(err)
			}

			var resultDto dto.ApiPayload[[]dto.Result]
			err = json.Unmarshal(resultBody, &resultDto)
			if err != nil {
				panic(err)
			}

			returnedResult := resultDto.Body
			returnedResultCount := len(returnedResult)

			if returnedResultCount != len(tt.expected.countries) {
				t.Errorf("Returned result count = %v, want %v", returnedResultCount, len(tt.expected.countries))
			}

			for i := 0; i < returnedResultCount; i++ {
				forwardIdx := i
				reverseIdx := len(tt.expected.countries) - 1 - i

				if returnedResult[forwardIdx].CountrySlug != tt.expected.countries[reverseIdx] {
					t.Errorf("Returned result country slug = %v, want %v", returnedResult[forwardIdx].CountrySlug, tt.expected.countries[reverseIdx])
				}

				expectedScore := reverseIdx

				if returnedResult[forwardIdx].Total != expectedScore*4 {
					t.Errorf("Returned result total = %v, want %v", returnedResult[forwardIdx].Total, expectedScore)
				}

				if returnedResult[forwardIdx].Costume != expectedScore {
					t.Errorf("Returned result costume = %v, want %v", returnedResult[forwardIdx].Costume, expectedScore)
				}

				if returnedResult[forwardIdx].Props != expectedScore {
					t.Errorf("Returned result props = %v, want %v", returnedResult[forwardIdx].Props, expectedScore)
				}

				if returnedResult[forwardIdx].Song != expectedScore {
					t.Errorf("Returned result song = %v, want %v", returnedResult[forwardIdx].Song, expectedScore)
				}

				if returnedResult[forwardIdx].Props != expectedScore {
					t.Errorf("Returned result props = %v, want %v", returnedResult[forwardIdx].Props, expectedScore)
				}
			}
		})
	}
}

func TestVoteHandler_GetVoteByUserAndCountry(t *testing.T) {
	type fields struct {
		Service service.VoteService
	}
	type args struct {
		resp http.ResponseWriter
		req  *http.Request
	}
	type expected struct {
		statusCode int
		vote       dto.Vote
	}
	tests := []struct {
		name     string
		fields   fields
		args     args
		expected expected
	}{
		{
			name: "successful get votes for one user and one country",
			fields: fields{
				Service: newTestVoteService(),
			},
			args: args{
				resp: httptest.NewRecorder(),
				req:  httptest.NewRequest(http.MethodGet, "/api/vote/countryanduser/"+countryNames[0], nil),
			},
			expected: expected{
				statusCode: http.StatusOK,
				vote: dto.Vote{
					UserId:      adminUserId,
					CountrySlug: countryNames[0],
					Costume:     0,
					Song:        0,
					Performance: 0,
					Props:       0,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			vh := VoteHandler{
				Service: tt.fields.Service,
			}

			mockAuth := dto.Auth{
				UserId: adminUserId,
			}
			tt.args.req = tt.args.req.WithContext(context.WithValue(context.Background(), "auth", mockAuth))
			tt.args.req = mux.SetURLVars(tt.args.req, map[string]string{"slug": countryNames[0]})

			vh.GetVoteByUserAndCountry(tt.args.resp, tt.args.req)
			result := tt.args.resp.(*httptest.ResponseRecorder).Result()
			defer func(Body io.ReadCloser) {
				err := Body.Close()
				if err != nil {
					panic(err)
				}
			}(result.Body)

			if result.StatusCode != tt.expected.statusCode {
				t.Errorf("Login return status code = %v, want %v", result.StatusCode, http.StatusOK)
			}

			resultBody, err := io.ReadAll(result.Body)
			if err != nil {
				panic(err)
			}

			var resultDto dto.ApiPayload[dto.Vote]
			err = json.Unmarshal(resultBody, &resultDto)
			if err != nil {
				panic(err)
			}

			if resultDto.Body.UserId != tt.expected.vote.UserId {
				t.Errorf("Returned vote user id = %v, want %v", resultDto.Body.UserId, tt.expected.vote.UserId)
			}

			if resultDto.Body.CountrySlug != tt.expected.vote.CountrySlug {
				t.Errorf("Returned vote country slug = %v, want %v", resultDto.Body.CountrySlug, tt.expected.vote.CountrySlug)
			}

			if resultDto.Body.Costume != tt.expected.vote.Costume {
				t.Errorf("Returned vote costume = %v, want %v", resultDto.Body.Costume, tt.expected.vote.Costume)
			}

			if resultDto.Body.Song != tt.expected.vote.Song {
				t.Errorf("Returned vote song = %v, want %v", resultDto.Body.Song, tt.expected.vote.Song)
			}

			if resultDto.Body.Performance != tt.expected.vote.Performance {
				t.Errorf("Returned vote performance = %v, want %v", resultDto.Body.Performance, tt.expected.vote.Performance)
			}

			if resultDto.Body.Props != tt.expected.vote.Props {
				t.Errorf("Returned vote props = %v, want %v", resultDto.Body.Props, tt.expected.vote.Props)
			}
		})
	}
}

func TestVoteHandler_UpdateVote(t *testing.T) {
	type fields struct {
		Service service.VoteService
	}
	type args struct {
		resp http.ResponseWriter
		req  *http.Request
	}
	type expected struct {
		statusCode int
		vote       dao.Vote
	}
	tests := []struct {
		name     string
		fields   fields
		args     args
		expected expected
	}{
		{
			name: "successful update song for one user and one country",
			fields: fields{
				Service: newTestVoteService(),
			},
			args: args{
				resp: httptest.NewRecorder(),
				req: httptest.NewRequest(http.MethodPut, "/api/vote/", strings.NewReader(fmt.Sprintf(`{
					"userId": "%s",
					"countrySlug": "%s",
					"cat": "song",
					"score": 10
                }`, adminUserId.String(), countryNames[0]))),
			},
			expected: expected{
				statusCode: http.StatusOK,
				vote: dao.Vote{
					UserId:      adminUserId,
					CountrySlug: countryNames[0],
					Costume:     0,
					Song:        10,
					Performance: 0,
					Props:       0,
					Total:       10,
				},
			},
		},
		{
			name: "successful update costume for one user and one country",
			fields: fields{
				Service: newTestVoteService(),
			},
			args: args{
				resp: httptest.NewRecorder(),
				req: httptest.NewRequest(http.MethodPost, "/api/vote/", strings.NewReader(fmt.Sprintf(`{
					"userId": "%s",
					"countrySlug": "%s",
					"cat": "costume",
					"score": 10
                }`, adminUserId.String(), countryNames[0]))),
			},
			expected: expected{
				statusCode: http.StatusOK,
				vote: dao.Vote{
					UserId:      adminUserId,
					CountrySlug: countryNames[0],
					Costume:     10,
					Song:        0,
					Performance: 0,
					Props:       0,
					Total:       10,
				},
			},
		},
		{
			name: "successful update perf for one user and one country",
			fields: fields{
				Service: newTestVoteService(),
			},
			args: args{
				resp: httptest.NewRecorder(),
				req: httptest.NewRequest(http.MethodPost, "/api/vote/", strings.NewReader(fmt.Sprintf(`{
					"userId": "%s",
					"countrySlug": "%s",
					"cat": "performance",
					"score": 10
                }`, adminUserId.String(), countryNames[0]))),
			},
			expected: expected{
				statusCode: http.StatusOK,
				vote: dao.Vote{
					UserId:      adminUserId,
					CountrySlug: countryNames[0],
					Costume:     0,
					Song:        0,
					Performance: 10,
					Props:       0,
					Total:       10,
				},
			},
		},
		{
			name: "successful update prop for one user and one country",
			fields: fields{
				Service: newTestVoteService(),
			},
			args: args{
				resp: httptest.NewRecorder(),
				req: httptest.NewRequest(http.MethodPost, "/api/vote/", strings.NewReader(fmt.Sprintf(`{
					"userId": "%s",
					"countrySlug": "%s",
					"cat": "props",
					"score": 10
                }`, adminUserId.String(), countryNames[0]))),
			},
			expected: expected{
				statusCode: http.StatusOK,
				vote: dao.Vote{
					UserId:      adminUserId,
					CountrySlug: countryNames[0],
					Costume:     0,
					Song:        0,
					Performance: 0,
					Props:       10,
					Total:       10,
				},
			},
		},
		{
			name: "failed update if users do not match",
			fields: fields{
				Service: newTestVoteService(),
			},
			args: args{
				resp: httptest.NewRecorder(),
				req: httptest.NewRequest(http.MethodPost, "/api/vote/", strings.NewReader(fmt.Sprintf(`{
					"userId": "%s",
					"countrySlug": "%s",
					"cat": "props",
					"score": 10
                }`, regularUserId.String(), countryNames[0]))),
			},
			expected: expected{
				statusCode: http.StatusUnauthorized,
				vote:       dao.Vote{},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			vh := VoteHandler{
				Service: tt.fields.Service,
			}

			mockAuth := dto.Auth{
				UserId: adminUserId,
			}
			tt.args.req = tt.args.req.WithContext(context.WithValue(context.Background(), "auth", mockAuth))
			tt.args.req = mux.SetURLVars(tt.args.req, map[string]string{"slug": countryNames[0]})

			vh.UpdateVote(tt.args.resp, tt.args.req)

			response := tt.args.resp.(*httptest.ResponseRecorder).Result()
			defer func(Body io.ReadCloser) {
				err := Body.Close()
				if err != nil {
					panic(err)
				}
			}(response.Body)

			if response.StatusCode != tt.expected.statusCode {
				t.Errorf("Login return status code = %v, want %v", response.StatusCode, http.StatusOK)
			}
			if response.StatusCode != http.StatusOK {
				// the vote was not updated, no need to check the result
				return
			}

			var result dao.Vote
			err := testDB.Get(fmt.Sprintf("%s_%s", adminUserId.String(), countryNames[0]), &result)
			if err != nil {
				panic(err)
			}

			if result.UserId != tt.expected.vote.UserId {
				t.Errorf("Returned vote user id = %v, want %v", result.UserId, tt.expected.vote.UserId)
			}
			if result.CountrySlug != tt.expected.vote.CountrySlug {
				t.Errorf("Returned vote country slug = %v, want %v", result.CountrySlug, tt.expected.vote.CountrySlug)
			}
			if result.Costume != tt.expected.vote.Costume {
				t.Errorf("Returned vote costume = %v, want %v", result.Costume, tt.expected.vote.Costume)
			}
			if result.Song != tt.expected.vote.Song {
				t.Errorf("Returned vote song = %v, want %v", result.Song, tt.expected.vote.Song)
			}
			if result.Performance != tt.expected.vote.Performance {
				t.Errorf("Returned vote performance = %v, want %v", result.Performance, tt.expected.vote.Performance)
			}
			if result.Props != tt.expected.vote.Props {
				t.Errorf("Returned vote props = %v, want %v", result.Props, tt.expected.vote.Props)
			}
			if result.Total != tt.expected.vote.Total {
				t.Errorf("Returned vote total = %v, want %v", result.Total, tt.expected.vote.Total)
			}

			err = testDB.Upsert(fmt.Sprintf("%s_%s", adminUserId.String(), countryNames[0]),
				dao.Vote{
					UserId:      adminUserId,
					CountrySlug: countryNames[0],
					Costume:     0,
					Song:        0,
					Performance: 0,
					Props:       0,
					Total:       0,
				})
			if err != nil {
				panic(err)
			}
		})
	}
}

func TestDefaultVoteService_UpdateVote_Broadcasting(t *testing.T) {
	vs := newTestVoteService()
	categories := make([]enum.Categories, 4)
	categories[0] = enum.Costume
	categories[1] = enum.Song
	categories[2] = enum.Performance
	categories[3] = enum.Props
	conf.App.VoteCountTrigger = 5

	tests := []struct {
		name          string
		updateCount   int
		expectMessage int
		countrySlug   string
	}{
		{
			name: "two vote updates - one broadcast message",

			updateCount:   conf.App.VoteCountTrigger,
			expectMessage: 1,
			countrySlug:   countryNames[1],
		},
		{
			name:          "three vote updates - one broadcast message",
			updateCount:   conf.App.VoteCountTrigger + 1,
			expectMessage: 1,
			countrySlug:   countryNames[2],
		},
		{
			name:          "single vote update - no broadcast",
			updateCount:   conf.App.VoteCountTrigger - 2,
			expectMessage: 0,
			countrySlug:   countryNames[3],
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			for len(testVoteBroadcastChan) > 0 {
				<-testVoteBroadcastChan
			}

			if len(testVoteBroadcastChan) != 0 {
				t.Errorf("testVoteBroadcastChan should be empty, got %v", len(testVoteBroadcastChan))
			}

			for i := 1; i <= tt.updateCount; i++ {
				voteUpdate := dto.VoteSingle{
					UserId:      adminUserId,
					CountrySlug: tt.countrySlug,
					Cat:         categories[i%len(categories)],
					Score:       5,
				}

				_, err := vs.UpdateVote(voteUpdate)
				if err != nil {
					panic(err)
				}
			}

			messageReceivedCount := 0
			select {
			case msg := <-testVoteBroadcastChan:

				if msg.Category == chatMsgType.VOTE_NOTIFICATION {
					var notification dto.VoteTracker
					err := json.Unmarshal(msg.Body, &notification)
					if err != nil {
						panic(err)
					}

					if notification.Country.Slug != tt.countrySlug {
						fmt.Printf("Received broadcast message: %+v\n", notification)
						messageReceivedCount++
					}
				}
			case <-time.After(2000 * time.Millisecond):
			}

			if tt.expectMessage != messageReceivedCount {
				t.Errorf("Expected %d broadcast message but got %d", tt.expectMessage, messageReceivedCount)
			}
		})
	}
}
