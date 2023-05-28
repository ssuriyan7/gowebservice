package endpoint

import (
	"encoding/json"
	"firstWebService/src/main/model"
	"firstWebService/src/main/service"
	"net/http"
)

func AddAnime(w http.ResponseWriter, r *http.Request) {

	animeName := r.FormValue("name")
	animeAuthor := r.FormValue("author")

	var response = model.HttpStatus{}
	if animeName == "" || animeAuthor == "" {
		response = model.HttpStatus{Code: 400, Message: "Invalid input."}
	} else {
		newAnime := model.Anime{
			Title:  animeName,
			Author: animeAuthor,
		}

		response = service.AddAnime(newAnime)
	}

	json.NewEncoder(w).Encode(response)

}

func DeleteAnime(w http.ResponseWriter, r *http.Request) {

	animeName := r.FormValue("name")

	var response = model.HttpStatus{}
	if animeName == "" {
		response = model.HttpStatus{Code: 400, Message: "Invalid input."}
	} else {
		response = service.DeleteAnime(animeName)
	}

	json.NewEncoder(w).Encode(response)
}
