package router

import (
	"eurovision/pkg/service"
	"io/ioutil"
	"net/http"
)

type AuthHandler struct {
	Service service.AuthService
}

func (ah AuthHandler) Register(resp http.ResponseWriter, req *http.Request) {
	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		panic(err)
	}

	auth, appErr := ah.Service.Register(body)
	if appErr != nil {
		writeResponse(resp, appErr.Code, appErr.AsMessage())
	} else {
		writeResponse(resp, http.StatusOK, auth)
	}
}

func (ah AuthHandler) Login(resp http.ResponseWriter, req *http.Request) {
	// params := mux.Vars(req)
	values := req.URL.Query()
	params := make(map[string]string, 0)
	for k, v := range values {
		params[k] = v[0]
	}

	auth, err := ah.Service.Login(params["t"], params["u"])
	if err != nil {
		writeResponse(resp, err.Code, err.AsMessage())
	} else {
		cookie := &http.Cookie{
			Name:   "token",
			Value:  auth.Token,
			MaxAge: 432000, // 5 days
		}
		http.SetCookie(resp, cookie)
		currentSessions.sessions[auth.Token] = auth.UserId.String()
		writeResponse(resp, http.StatusOK, auth)
	}
}

func (ah AuthHandler) Authenticate(resp http.ResponseWriter, req *http.Request) {

}
