package data

import (
	"log"

	"github.com/anytimesoon/eurovision-party/pkg/data/dao"
	"github.com/anytimesoon/eurovision-party/pkg/enum/reactAction"
	"github.com/google/uuid"
	"github.com/timshannon/bolthold"
)

type CommentRepository interface {
	GetAllComments() ([]dao.Comment, error)
	CreateComment(dao.Comment) (*dao.Comment, error)
	DeleteComment(string) error
	GetCommentsAfter(string) ([]dao.Comment, error)
	UpdateReaction(uuid.UUID, uuid.UUID, reactAction.ReactAction, string) (*dao.Comment, error)
}

type CommentRepositoryDb struct {
	store *bolthold.Store
}

func NewCommentRepositoryDb(store *bolthold.Store) CommentRepositoryDb {
	return CommentRepositoryDb{store}
}

func (db CommentRepositoryDb) GetAllComments() ([]dao.Comment, error) {
	comments := make([]dao.Comment, 0)

	q := &bolthold.Query{}
	err := db.store.Find(&comments, q.SortBy("CreatedAt"))
	if err != nil {
		log.Println("Error while querying comment table", err)
		return nil, err
	}

	return comments, nil
}

func (db CommentRepositoryDb) GetCommentsAfter(commentId string) ([]dao.Comment, error) {
	comments := make([]dao.Comment, 0)
	var latestComment dao.Comment

	err := db.store.Get(commentId, &latestComment)
	if err != nil {
		log.Println("Error while querying comment table", err)
		return nil, err
	}

	err = db.store.Find(&comments, bolthold.Where("CreatedAt").Gt(latestComment.CreatedAt).SortBy("CreatedAt"))
	if err != nil {
		log.Println("Error while querying comment table", err)
		return nil, err
	}

	return comments, nil
}

func (db CommentRepositoryDb) CreateComment(comment dao.Comment) (*dao.Comment, error) {
	err := db.store.Insert(comment.UUID.String(), &comment)
	if err != nil {
		log.Printf("Error when creating comment from user %s, %s", comment.UserId, err)
		return nil, err
	}

	return &comment, nil
}

func (db CommentRepositoryDb) DeleteComment(uuid string) error {
	err := db.store.Delete(uuid, dao.Comment{})
	if err != nil {
		log.Println("Error when deleting comment", err)
		return err
	}

	return nil
}

func (db CommentRepositoryDb) UpdateReaction(commentId uuid.UUID, userId uuid.UUID, action reactAction.ReactAction, reaction string) (*dao.Comment, error) {
	var comment dao.Comment
	err := db.store.Get(commentId.String(), &comment)
	if err != nil {
		log.Println("Error when fetching comment", err)
		return nil, err
	}

	switch action {
	case reactAction.ADD:
		comment.Reactions[reaction] = append(comment.Reactions[reaction], userId)
	case reactAction.DELETE:
		users := comment.Reactions[reaction]
		for i, u := range users {
			if u == userId {
				users = append(users[:i], users[i+1:]...)
				break
			}
		}
		if len(users) == 0 {
			delete(comment.Reactions, reaction)
		} else {
			comment.Reactions[reaction] = users
		}
	default:
		log.Println("Invalid action. This should never happen.")
		return nil, nil
	}

	err = db.store.Update(commentId.String(), comment)
	if err != nil {
		log.Println("Error when updating comment", err)
		return nil, err
	}

	return &comment, nil
}
