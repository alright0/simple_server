package middleware

import (
	"net/http"
	"os"

	"github.com/golang-jwt/jwt/v5"
)

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tokenString := r.Header.Get("Authorization")
		if tokenString == "" {
			http.Error(w, "You are not authorized", http.StatusUnauthorized)
			return
		}

		tokenString = tokenString[len("Bearer "):]
		var jwtSecretKey = []byte(os.Getenv("SECRET_KEY"))
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if token.Method.Alg() != jwt.SigningMethodHS256.Alg() {
				http.Error(w, "Token not valid", http.StatusUnauthorized)
			}
			return jwtSecretKey, nil
		})

		if err != nil || !token.Valid {
			http.Error(w, "Error decoding token", http.StatusUnauthorized)
			return
		}

		//claims, ok := token.Claims.(jwt.MapClaims)
		//if !ok {
		//	http.Error(w, err.Error(), http.StatusUnauthorized)
		//}

		next.ServeHTTP(w, r)
	})
}
