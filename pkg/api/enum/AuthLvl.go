package enum

type AuthLvl uint8

const (
	NONE AuthLvl = iota
	ADMIN
	BOT
)
