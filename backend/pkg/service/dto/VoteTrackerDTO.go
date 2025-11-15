package dto

type VoteTracker struct {
	Count   int     `json:"count"`
	Country Country `json:"country"`
	Comment Comment `json:"comment"`
}
