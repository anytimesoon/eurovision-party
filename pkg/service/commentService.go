package service

import (
	"encoding/json"
	"github.com/anytimesoon/eurovision-party/pkg/domain"
	"github.com/anytimesoon/eurovision-party/pkg/dto"
	"github.com/anytimesoon/eurovision-party/pkg/errs"
	"github.com/google/uuid"
	"log"
	"time"
)

//go:generate mockgen -source=commentService.go -destination=../../mocks/service/mockCommentService.go -package=service eurovision/pkg/service
type CommentService interface {
	FindAllComments() ([]dto.Comment, *errs.AppError)
	CreateComment([]byte) (*dto.Comment, *errs.AppError)
	DeleteComment(string) *errs.AppError
	FindCommentsAfter(json.RawMessage) ([]byte, *errs.AppError)
}

type DefaultCommentService struct {
	repo domain.CommentRepository
}

func NewCommentService(repo domain.CommentRepository) DefaultCommentService {
	return DefaultCommentService{repo}
}

func (service DefaultCommentService) FindAllComments() ([]dto.Comment, *errs.AppError) {
	commentsDTO := make([]dto.Comment, 0)

	commentData, err := service.repo.FindAllComments()
	if err != nil {
		return nil, err
	}

	for _, comment := range commentData {
		commentsDTO = append(commentsDTO, comment.ToDto())
	}

	return commentsDTO, nil
}

func (service DefaultCommentService) FindCommentsAfter(commentIdJSON json.RawMessage) ([]byte, *errs.AppError) {
	var commentId string
	var comments []domain.Comment
	var appErr *errs.AppError
	commentsDTO := make([]dto.Comment, 0)

	err := json.Unmarshal(commentIdJSON, &commentId)
	if err != nil {
		log.Println("Failed to unmarshal comment id.", err)
		return nil, errs.NewUnexpectedError(errs.Common.BadlyFormedObject)
	}

	if commentId == "" {
		comments, appErr = service.repo.FindAllComments()
	} else {
		comments, appErr = service.repo.FindCommentsAfter(commentId)
	}

	if appErr != nil {
		return nil, appErr
	}

	for _, comment := range comments {
		commentsDTO = append(commentsDTO, comment.ToDto())
	}

	commentsJSON, err := json.Marshal(commentsDTO)
	if err != nil {
		log.Println("Failed to marshal latest comments.", err)
		return nil, errs.NewUnexpectedError(errs.Common.BadlyFormedObject)
	}

	return commentsJSON, nil
}

func (service DefaultCommentService) CreateComment(body []byte) (*dto.Comment, *errs.AppError) {
	commentDTO := dto.Comment{}
	err := json.Unmarshal(body, &commentDTO)
	if err != nil {
		log.Println("Failed to unmarshal comment.", err)
		return nil, errs.NewUnexpectedError(errs.Common.BadlyFormedObject)
	}

	appErr := commentDTO.Validate()
	if appErr != nil {
		return nil, appErr
	}

	_, err = uuid.Parse(commentDTO.UUID.String())
	if err != nil {
		commentDTO.UUID = uuid.New()
	}
	commentDTO.CreatedAt = time.Now()

	comment, appErr := service.repo.CreateComment(commentDTO)
	if appErr != nil {
		return nil, appErr
	}

	commentDTO = comment.ToDto()

	return &commentDTO, nil
}

func (service DefaultCommentService) DeleteComment(uuid string) *errs.AppError {
	err := service.repo.DeleteComment(uuid)
	if err != nil {
		return err
	}

	return nil
}
