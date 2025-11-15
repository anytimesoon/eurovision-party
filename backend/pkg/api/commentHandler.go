package api

import (
	"io"
	"log"
	"net/http"

	"github.com/anytimesoon/eurovision-party/pkg/errs"
	"github.com/anytimesoon/eurovision-party/pkg/service"
	dto2 "github.com/anytimesoon/eurovision-party/pkg/service/dto"
)

type CommentHandler struct {
	Service service.CommentService
}

func (ch *CommentHandler) React(resp http.ResponseWriter, req *http.Request) {
	var commentReaction *dto2.CommentReaction
	var appErr *errs.AppError
	body, err := io.ReadAll(req.Body)
	if err != nil {
		log.Println("FAILED to read body of USER UPDATE.", err)
		WriteResponse(resp, http.StatusBadRequest, dto2.CommentReaction{}, "Failed to read body of comment reaction.")
		return
	}

	reaction, err := dto2.Deserialize[dto2.CommentReaction](body)
	if err != nil {
		log.Println("FAILED to deserialize comment reaction.", err)
		WriteResponse(resp, http.StatusBadRequest, dto2.CommentReaction{}, "Failed to deserialize comment reaction.")
		return
	}
	if req.Context().Value("auth").(dto2.Auth).UserId == reaction.UserId {
		commentReaction, appErr = ch.Service.UpdateReaction(*reaction)
	} else {
		commentReaction = &dto2.CommentReaction{}
		appErr = errs.NewUnauthorizedError(errs.Common.Unauthorized)
	}

	if appErr != nil {
		WriteResponse(resp, appErr.Code, dto2.CommentReaction{}, appErr.Message)
	} else {
		WriteResponse(resp, http.StatusOK, *commentReaction, "")
	}
}
