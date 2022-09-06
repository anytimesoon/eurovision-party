package handler

import (
	"bytes"
	"encoding/json"
	"eurovision/mocks/service"
	"eurovision/pkg/dto"
	"eurovision/pkg/errs"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

var commentRouter *mux.Router
var comhan CommentHandler
var mockCommentService *service.MockCommentService
var mockComments []dto.Comment
var mockComment dto.Comment
var commentJSON []byte
var commentBody *bytes.Buffer
var invalidComment dto.Comment
var invalidCommentJSON []byte
var invalidCommentBody *bytes.Buffer

func setupCommentTest(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockCommentService = service.NewMockCommentService(ctrl)
	comhan = CommentHandler{mockCommentService}
	mockComments = []dto.Comment{
		{UUID: uuid.New(), UserId: uuid.New(), Text: "tEsTcOmMeNt", CreatedAt: time.Now()},
		{UUID: uuid.New(), UserId: uuid.New(), Text: "tEsTcOmMeNt AGAIN!", CreatedAt: time.Now()},
	}

	mockComment = mockComments[0]
	commentJSON, _ = json.Marshal(mockComment)
	commentBody = bytes.NewBuffer(commentJSON)

	invalidComment = dto.Comment{UUID: uuid.New(), UserId: uuid.New(), Text: "", CreatedAt: time.Now()}
	invalidCommentJSON, _ = json.Marshal(mockComment)
	invalidCommentBody = bytes.NewBuffer(commentJSON)

	commentRouter = mux.NewRouter()
	commentRouter.HandleFunc("/comment", comhan.FindAllComments).Methods(http.MethodGet)
	commentRouter.HandleFunc("/comment", comhan.CreateComment).Methods(http.MethodPost)
	commentRouter.HandleFunc("/comment/{uuid}", comhan.RemoveComment).Methods(http.MethodDelete)
}

func Test_all_comments_route_should_return_comments_with_200_code(t *testing.T) {
	setupCommentTest(t)

	mockCommentService.EXPECT().FindAllComments().Return(mockComments, nil)

	req, _ := http.NewRequest(http.MethodGet, "/comment", nil)

	recorder := httptest.NewRecorder()
	commentRouter.ServeHTTP(recorder, req)

	if recorder.Code != http.StatusOK {
		t.Error("Expected status code 200, but got", recorder.Code)
	}

	comments := make([]dto.Comment, 0)
	json.Unmarshal(recorder.Body.Bytes(), &comments)

	if len(comments) != 2 {
		t.Error("Expected 2 users, but found", len(comments))
	}
}

func Test_all_comments_route_should_return_500_code(t *testing.T) {
	setupCommentTest(t)

	mockCommentService.EXPECT().FindAllComments().Return(nil, errs.NewUnexpectedError("Couldn't find comments"))

	req, _ := http.NewRequest(http.MethodGet, "/comment", nil)

	recorder := httptest.NewRecorder()
	commentRouter.ServeHTTP(recorder, req)

	if recorder.Code != http.StatusInternalServerError {
		t.Error("Expected status code 500, but got", recorder.Code)
	}
}

func Test_new_comment_route_returns_500_code(t *testing.T) {
	setupCommentTest(t)

	mockCommentService.EXPECT().CreateComment(commentJSON).Return(nil, errs.NewUnexpectedError("Couldn't create new comment"))

	req, _ := http.NewRequest(http.MethodPost, "/comment", commentBody)

	recorder := httptest.NewRecorder()
	commentRouter.ServeHTTP(recorder, req)

	if recorder.Code != http.StatusInternalServerError {
		t.Error("Expected status code 500, but got", recorder.Code)
	}
}

func Test_new_comment_route_returns_400_error(t *testing.T) {
	setupCommentTest(t)

	mockCommentService.EXPECT().CreateComment(invalidCommentJSON).Return(nil, errs.NewInvalidError("Comment name must not be blank"))

	req, _ := http.NewRequest(http.MethodPost, "/comment", invalidCommentBody)

	recorder := httptest.NewRecorder()
	commentRouter.ServeHTTP(recorder, req)

	if recorder.Code != http.StatusBadRequest {
		t.Errorf("Expected 400 error, but got %d", recorder.Code)
	}
}

func Test_new_comment_route_returns_200_code(t *testing.T) {
	setupCommentTest(t)

	mockCommentService.EXPECT().CreateComment(commentJSON).Return(&mockComment, nil)

	req, _ := http.NewRequest(http.MethodPost, "/comment", commentBody)

	recorder := httptest.NewRecorder()
	commentRouter.ServeHTTP(recorder, req)

	if recorder.Code != http.StatusOK {
		t.Error("Expected status code 200, but got", recorder.Code)
	}
}

func Test_delete_comment_route_returns_500_code(t *testing.T) {
	setupCommentTest(t)
	commentId := mockComment.UUID.String()

	mockCommentService.EXPECT().DeleteComment(commentId).Return(errs.NewUnexpectedError("Couldn't delete user"))
	path := "/comment/" + commentId

	req, _ := http.NewRequest(http.MethodDelete, path, nil)

	recorder := httptest.NewRecorder()
	commentRouter.ServeHTTP(recorder, req)

	if recorder.Code != http.StatusInternalServerError {
		t.Error("Expected status code 500, but got", recorder.Code)
	}
}

func Test_delete_comment_route_returns_200_code(t *testing.T) {
	setupCommentTest(t)

	mockCommentService.EXPECT().DeleteComment(mockComment.UUID.String()).Return(nil)
	path := "/comment/" + mockComment.UUID.String()

	req, _ := http.NewRequest(http.MethodDelete, path, nil)

	recorder := httptest.NewRecorder()
	commentRouter.ServeHTTP(recorder, req)

	if recorder.Code != http.StatusOK {
		t.Error("Expected status code 200, but got", recorder.Code)
	}
}
