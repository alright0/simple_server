package views

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type Credentials struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

type Token struct {
	Token string `json:"token"`
}

type ErrorResponse struct {
	Error string `json:"error"`
}

func LoginView(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	bodyBytes, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if len(bodyBytes) == 0 {
		http.Error(w, "body is empty", http.StatusBadRequest)
		return
	}

	var credentials Credentials

	e := json.Unmarshal(bodyBytes, &credentials)
	if e != nil {
		http.Error(w, e.Error(), http.StatusUnauthorized)
		return
	}

	token, err := generateJWT(credentials.Email)
	if err != nil {
		http.Error(w, "Not authorized", http.StatusUnauthorized)
		return
	}

	err = json.NewEncoder(w).Encode(Token{token})
	if err != nil {
		log.Printf("Error encoding response: %v", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

}

func generateJWT(email string) (string, error) {
	var jwtSecretKey = []byte(os.Getenv("SECRET_KEY"))

	claims := jwt.MapClaims{
		"email": email,
		"exp":   time.Now().Add(time.Hour * 24).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtSecretKey)
}
