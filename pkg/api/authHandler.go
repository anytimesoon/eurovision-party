package api

import (
	"encoding/json"
	"github.com/anytimesoon/eurovision-party/pkg/api/dto"
	"github.com/anytimesoon/eurovision-party/pkg/errs"
	"github.com/anytimesoon/eurovision-party/pkg/service"
	"io"
	"log"
	"net/http"
)

type AuthHandler struct {
	Service service.AuthService
}

func (ah AuthHandler) Login(resp http.ResponseWriter, req *http.Request) {
	body, err := io.ReadAll(req.Body)
	if err != nil {
		log.Println("FAILED to read body of login request!", err)
		return
	}

	var authDTO dto.Auth
	err = json.Unmarshal(body, &authDTO)
	if err != nil {
		log.Println("FAILED to unmarshal json!", err)
		appErr := errs.NewUnexpectedError(errs.Common.BadlyFormedObject)
		WriteResponse(resp, appErr.Code, dto.SessionAuth{}, appErr.Message)
		return
	}

	auth, user, appErr := ah.Service.Login(authDTO)

	if appErr != nil {
		WriteResponse(resp, appErr.Code, dto.SessionAuth{}, appErr.Message)
	} else {
		WriteResponse(resp, http.StatusOK, auth.ToSession(*user), "")
	}
}
