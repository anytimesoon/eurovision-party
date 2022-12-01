package router

import (
	"encoding/json"
	"eurovision/pkg/errs"
	"net/http"
)

func writeResponse(resp http.ResponseWriter, code int, data interface{}) {
	resp.WriteHeader(code)
	if err := json.NewEncoder(resp).Encode(data); err != nil {
		panic(err)
	}
}

func (currentSessions sessionStore) authorize(req *http.Request) (bool, *errs.AppError) {
	// cookie, err := req.Cookie("token")
	// if err != nil {
	// 	log.Println("Error when trying to read cookie", err)
	// 	return false, errs.NewUnauthorizedError(errs.Common.Unauthorized)
	// }

	// params := mux.Vars(req)

	// session, found := currentSessions.sessions[cookie.Value]
	// if found && session.authLvl == domain.Admin {
	// 	return true, nil
	// } else if found && session.slug == params["slug"] {
	// 	return true, nil
	// } else {
	// 	return false, errs.NewUnauthorizedError(errs.Common.Unauthorized)
	// }
	return true, nil
}
