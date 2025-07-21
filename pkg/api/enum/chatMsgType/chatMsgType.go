package chatMsgType

type ChatMsgType string

const (
	COMMENT           = "comment"
	COMMENT_ARRAY     = "comments"
	UPDATE_USER       = "updateUser"
	VOTE_NOTIFICATION = "voteNotification"
	LATEST_COMMENT    = "latestComment"
	NEW_USER          = "newUser"
	ERROR             = "error"
)
