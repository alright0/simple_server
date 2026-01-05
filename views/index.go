package views

import (
	"encoding/json"
	"fmt"
	"log"
	"main/dto"
	"net/http"
)

func IndexView(w http.ResponseWriter, r *http.Request) {
	var credentials dto.LoginRequest

	marshalledData, err := json.Marshal(credentials)
	if err != nil {
		log.Fatalf("impossible to marshal user: %s", err)
	}

	r.Header.Set("Content-Type", "application/json")
	fmt.Fprintf(w, string(marshalledData))

}
