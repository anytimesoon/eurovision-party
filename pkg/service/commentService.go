package service

import (
	"encoding/json"
	"eurovision/pkg/domain"
	"eurovision/pkg/dto"
	"eurovision/pkg/errs"
	"log"
)

type CommentService interface {
	FindAllComments() ([]dto.Comment, *errs.AppError)
	CreateComment([]byte) (*dto.Comment, *errs.AppError)
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

func (service DefaultCommentService) CreateComment(body []byte) (*dto.Comment, *errs.AppError) {
	var commentDTO dto.Comment
	err := json.Unmarshal(body, &commentDTO)
	if err != nil {
		log.Println("FAILED to unmarshal json!", err)
		return nil, errs.NewUnexpectedError(errs.Common.BadlyFormedObject)
	}

	comment, appErr := service.repo.CreateComment(commentDTO)
	if err != nil {
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
