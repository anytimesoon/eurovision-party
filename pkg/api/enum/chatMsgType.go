package enum

type ChatMsgType string

const (
	COMMENT        = "comment"
	COMMENT_ARRAY  = "comments"
	UPDATE_USER    = "updateUser"
	LATEST_COMMENT = "latestComment"
	NEW_USER       = "newUser"
	ERROR          = "error"
)
