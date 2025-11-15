package dto

import (
	"github.com/anytimesoon/eurovision-party/pkg/enum/reactAction"
	"github.com/google/uuid"
)

type CommentReaction struct {
	Action    reactAction.ReactAction `json:"action"`
	UserId    uuid.UUID               `json:"userId"`
	CommentId uuid.UUID               `json:"commentId"`
	Reaction  string                  `json:"reaction"`
}
