package controllers

import (
	"net/http"

	"github.com/leepuppychow/fishies/helpers"
	"github.com/leepuppychow/fishies/models"
)

func FishIndex(w http.ResponseWriter, r *http.Request) {
	data, err := models.AllFish()
	helpers.WriteResponse(data, err, 400, w)
}

func FishCreate(w http.ResponseWriter, r *http.Request) {
	data, err := models.CreateFish(r.Body)
	helpers.WriteResponse(data, err, 400, w)
}
