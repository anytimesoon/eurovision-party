package service

import (
	"encoding/json"
	"eurovision/pkg/domain"
	"eurovision/pkg/dto"
	"eurovision/pkg/errs"
	"log"
	"time"

	"github.com/google/uuid"
)

//go:generate mockgen -source=commentService.go -destination=../../mocks/service/mockCommentService.go -package=service eurovision/pkg/service
type CommentService interface {
	FindAllComments() ([]dto.Comment, *errs.AppError)
	CreateComment([]byte, uuid.UUID) ([]byte, *errs.AppError)
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

func (service DefaultCommentService) CreateComment(message []byte, userId uuid.UUID) ([]byte, *errs.AppError) {
	commentDTO := dto.Comment{
		UUID:      uuid.New(),
		UserId:    userId,
		Text:      string(message),
		CreatedAt: time.Now(),
	}

	appErr := commentDTO.Validate()
	if appErr != nil {
		return nil, appErr
	}

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
