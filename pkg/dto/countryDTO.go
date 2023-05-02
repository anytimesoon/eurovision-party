package dto

type Country struct {
	Name          string `json:"name"`
	Slug          string `json:"slug"`
	BandName      string `json:"bandName"`
	SongName      string `json:"songName"`
	Flag          string `json:"flag"`
	Participating bool   `json:"participating"`
}
