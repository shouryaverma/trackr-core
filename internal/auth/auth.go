package auth

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	uuid "github.com/satori/go.uuid"
)

// CreateToken ...
func CreateToken(userID uuid.UUID) (string, error) {
	return "", nil
}

// TokenValid ...
func TokenValid(request *http.Request) error {
	return nil
}

// ExtractToken ...
func ExtractToken(request *http.Request) error {
	return nil
}

// ExtractUserID ...
func ExtractUserID(request *http.Request) (string, error) {
	return "", nil
}

// Pretty display the claims nicely in the terminal
func Pretty(data interface{}) {
	b, err := json.MarshalIndent(data, "", " ")
	if err != nil {
		log.Println(err)
		return
	}

	fmt.Println(string(b))
}
