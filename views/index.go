package views

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type UserStruct struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

func IndexView(w http.ResponseWriter, r *http.Request) {
	user := UserStruct{
		Name:  "John Doe",
		Email: "test@email.ru",
	}

	marshalledData, err := json.Marshal(user)
	if err != nil {
		log.Fatalf("impossible to marshal user: %s", err)
	}

	r.Header.Set("Content-Type", "application/json")
	fmt.Fprintf(w, string(marshalledData))
}
