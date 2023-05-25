package model

type Anime struct {
	Title  string `json:"title"`
	Author string `json:"author"`
}

func NewAnime(name string, author string) *Anime {
	a := Anime{Title: name, Author: author}
	return &a
}
