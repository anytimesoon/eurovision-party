package dto

type UpdateMessage struct {
	UpdatedUser User    `json:"updatedUser"`
	Comment     Comment `json:"comment"`
}
