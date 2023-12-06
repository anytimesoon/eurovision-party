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
	CreateComment([]byte) ([]byte, *errs.AppError)
	DeleteComment(string) *errs.AppError
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

func (service DefaultCommentService) CreateComment(body []byte) ([]byte, *errs.AppError) {
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

	commentDTO.UUID = uuid.New()
	commentDTO.CreatedAt = time.Now()

	comment, appErr := service.repo.CreateComment(commentDTO)
	if appErr != nil {
		return nil, appErr
	}

	commentDTO = comment.ToDto()

	commentJSON, err := json.Marshal(commentDTO)
	if err != nil {
		log.Println("FAILED to marshal commentDTO!", err)
		return nil, errs.NewUnexpectedError(errs.Common.BadlyFormedObject)
	}

	return commentJSON, nil
}

func (service DefaultCommentService) DeleteComment(uuid string) *errs.AppError {
	err := service.repo.DeleteComment(uuid)
	if err != nil {
		return err
	}

	return nil
}
