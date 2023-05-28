package service

import (
	"firstWebService/src/main/model"
	"firstWebService/src/main/util"
	"fmt"
	"log"
)

func AddAnime(anime model.Anime) model.HttpStatus {
	db := util.InitDb()

	_, err := db.Exec("INSERT INTO anime(title, author) VALUES($1, $2) returning id;", anime.Title, anime.Author)

	if err != nil {
		log.Print("\nQuery exec failed.")
		fmt.Println(err)
		return model.HttpStatus{Code: 500, Message: "Failed to add anime: Query exec failed."}
	} else {
		log.Print("Inserted anime successfully.")
	}

	return model.HttpStatus{Code: 200, Message: "Added anime."}
}

func DeleteAnime(animeTitle string) model.HttpStatus {
	db := util.InitDb()
	_, err := db.Exec("DELETE FROM anime WHERE title=$1;", animeTitle)

	if err != nil {
		log.Print("\nQuery exec failed.")
		log.Print(err)
		return model.HttpStatus{Code: 500, Message: "Failed to add anime: Query exec failed."}
	} else {
		log.Print(animeTitle + " deleted successfully.")
		return model.HttpStatus{Code: 200, Message: "Deletion success."}
	}
}
