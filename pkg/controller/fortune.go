package controller

import (
	"encoding/json"
	"fortune-back-go/pkg/repo"
	"io/ioutil"
	"net/http"
)

type FortuneController struct {
	FortuneRepo *repo.FortuneRepo
}

func NewFortuneController() *FortuneController {
	return &FortuneController{
		FortuneRepo: repo.NewFortuneRepo(),
	}
}

func (c *FortuneController) GetFortune(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Failed to read request body", http.StatusInternalServerError)
		return
	}

	params := struct {
		BirthYear  string `json:"birthYear"`
		BirthMonth string `json:"birthMonth"`
		BirthDay   string `json:"birthDay"`
		IsLunar    bool   `json:"isLunar"`
	}{}

	err = json.Unmarshal(body, &params)
	if err != nil {
		http.Error(w, "Failed to decode JSON body", http.StatusBadRequest)
		return
	}

	var ganji string
	if params.IsLunar {
		ganji, err = repo.GetGanjiWithLunar(params.BirthYear, params.BirthMonth, params.BirthDay)
	} else {
		ganji, err = repo.GetGanjiWithSolar(params.BirthYear, params.BirthMonth, params.BirthDay)
	}

	if err != nil {
		http.Error(w, "Failed to get ganji", http.StatusInternalServerError)
		return
	}

	resp, err := c.FortuneRepo.GetFortune(ganji)
	if err != nil {
		http.Error(w, "Failed to get fortune", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(resp)
	if err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
		return
	}
}
