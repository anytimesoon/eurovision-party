package service

import (
	"encoding/json"
	"github.com/anytimesoon/eurovision-party/pkg/data"
	"github.com/anytimesoon/eurovision-party/pkg/data/dao"
	"github.com/anytimesoon/eurovision-party/pkg/errs"
	"github.com/anytimesoon/eurovision-party/pkg/service/dto"
	"github.com/google/uuid"
	"log"
	"time"
)

type CommentService interface {
	FindAllComments() ([]dto.Comment, *errs.AppError)
	CreateComment([]byte) (*dto.Comment, *errs.AppError)
	DeleteComment(string) *errs.AppError
	FindCommentsAfter(json.RawMessage) ([]byte, *errs.AppError)
}

type DefaultCommentService struct {
	repo data.CommentRepository
}

func NewCommentService(repo data.CommentRepository) DefaultCommentService {
	return DefaultCommentService{repo}
}

func (cs DefaultCommentService) FindAllComments() ([]dto.Comment, *errs.AppError) {
	commentsDTO := make([]dto.Comment, 0)

	commentData, err := cs.repo.GetAllComments()
	if err != nil {
		return nil, errs.NewUnexpectedError(errs.Common.DBFail)
	}

	for _, comment := range commentData {
		commentsDTO = append(commentsDTO, comment.ToDto())
	}

	return commentsDTO, nil
}

func (cs DefaultCommentService) FindCommentsAfter(commentIdJSON json.RawMessage) ([]byte, *errs.AppError) {
	var commentId string
	var comments []dao.Comment
	commentsDTO := make([]dto.Comment, 0)

	err := json.Unmarshal(commentIdJSON, &commentId)
	if err != nil {
		log.Println("Failed to unmarshal comment id.", err)
		return nil, errs.NewUnexpectedError(errs.Common.BadlyFormedObject)
	}

	if commentId == "" {
		comments, err = cs.repo.GetAllComments()
	} else {
		comments, err = cs.repo.GetCommentsAfter(commentId)
	}

	if err != nil {
		return nil, errs.NewUnexpectedError(errs.Common.DBFail)
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

func (cs DefaultCommentService) CreateComment(body []byte) (*dto.Comment, *errs.AppError) {
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

	comment, err := cs.repo.CreateComment(dao.Comment{}.FromDTO(commentDTO))
	if err != nil {
		return nil, errs.NewUnexpectedError(errs.Common.NotCreated + "your comment")
	}

	commentDTO = comment.ToDto()

	return &commentDTO, nil
}

func (cs DefaultCommentService) DeleteComment(uuid string) *errs.AppError {
	err := cs.repo.DeleteComment(uuid)
	if err != nil {
		return errs.NewUnexpectedError(errs.Common.NotDeleted + "your comment")
	}

	return nil
}
