package controllers

import (
	"clickship/models"
	"clickship/utils"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"net/http"
)

type LoginBody struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type RegisterBody struct {
	LoginBody
	Username string `json:"username"`
}

func Login(w http.ResponseWriter, r *http.Request) {
	user := models.User{}
	loginBody := LoginBody{}
	database := models.GetDatabase(r)
	err := json.NewDecoder(r.Body).Decode(&loginBody)
	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Invalid request")
		return
	}

	hash := sha256.New()
	hash.Write([]byte(loginBody.Password))
	passwordHash := hex.EncodeToString(hash.Sum(nil))

	database.Where("email = ? AND password = ?", loginBody.Email, passwordHash).First(&user)
	if user.ID == 0 {
		utils.RespondWithError(w, http.StatusNotFound, "User not found")
		return
	}

	utils.RespondWithJSON(w, user)
}

func Register(w http.ResponseWriter, r *http.Request) {
	user := models.User{}
	loginBody := RegisterBody{}
	database := models.GetDatabase(r)
	err := json.NewDecoder(r.Body).Decode(&loginBody)
	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Invalid request")
		return
	}

	hash := sha256.New()
	hash.Write([]byte(loginBody.Password))
	passwordHash := hex.EncodeToString(hash.Sum(nil))

	user.Email = loginBody.Email
	user.Password = passwordHash
	user.Username = loginBody.Username

	database.Create(&user)
	utils.RespondWithJSON(w, user)
}
