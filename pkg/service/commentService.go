package service

import (
	"eurovision/pkg/domain"
	"eurovision/pkg/dto"
)

type CommentService interface {
	FindAllComments() ([]dto.Comment, error)
}

type DefaultCommentService struct {
	repo domain.CommentRepository
}

func NewCommentService(repo domain.CommentRepository) DefaultCommentService {
	return DefaultCommentService{repo}
}

func (service DefaultCommentService) FindAllComments() ([]dto.Comment, error) {
	commentsDTO := make([]dto.Comment, 0)

	commentData, err := service.repo.FindAllComments()
	if err != nil {
		return commentsDTO, err
	}

	for _, comment := range commentData {
		commentsDTO = append(commentsDTO, comment.ToDto())
	}

	return commentsDTO, nil
}
