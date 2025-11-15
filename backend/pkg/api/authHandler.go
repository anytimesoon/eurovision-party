package api

import (
	"encoding/json"
	"io"
	"log"
	"net/http"

	"github.com/anytimesoon/eurovision-party/pkg/errs"
	"github.com/anytimesoon/eurovision-party/pkg/service"
	dto2 "github.com/anytimesoon/eurovision-party/pkg/service/dto"
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

	var authDTO dto2.Auth
	err = json.Unmarshal(body, &authDTO)
	if err != nil {
		log.Println("FAILED to unmarshal json!", err)
		appErr := errs.NewUnexpectedError(errs.Common.BadlyFormedObject)
		WriteResponse(resp, appErr.Code, &dto2.Session{}, appErr.Message)
		return
	}

	session, appErr := ah.Service.Login(authDTO)

	if appErr != nil {
		WriteResponse(resp, appErr.Code, &dto2.Session{}, appErr.Message)
	} else {
		WriteResponse(resp, http.StatusOK, session, "")
	}
}
