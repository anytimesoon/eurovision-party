package authLvl

type AuthLvl uint8

const (
	USER AuthLvl = iota
	ADMIN
	BOT
	FRIEND_OF_FRIEND
)
