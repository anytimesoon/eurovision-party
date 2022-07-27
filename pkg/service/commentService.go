package service

import (
	"encoding/json"
	"eurovision/pkg/domain"
	"eurovision/pkg/dto"
	"log"
)

type CommentService interface {
	FindAllComments() ([]dto.Comment, error)
	CreateComment([]byte) (dto.Comment, error)
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

func (service DefaultCommentService) CreateComment(body []byte) (dto.Comment, error) {
	var commentDTO dto.Comment
	err := json.Unmarshal(body, &commentDTO)
	if err != nil {
		log.Println("FAILED to unmarshal json!", err)
		return commentDTO, err
	}

	comment, err := service.repo.CreateComment(commentDTO)
	if err != nil {
		log.Println("FAILED to create user", err)
		return commentDTO, err
	}

	return comment.ToDto(), nil
}
