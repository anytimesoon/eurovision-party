package enum

type Categories string

const (
	Song        = "song"
	Costume     = "costume"
	Performance = "performance"
	Props       = "props"
)

func (c Categories) String() string {
	return string(c)
}
