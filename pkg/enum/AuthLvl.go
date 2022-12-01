package enum

type AuthLvl uint8

const (
	None AuthLvl = iota
	Admin
	Bot
)
