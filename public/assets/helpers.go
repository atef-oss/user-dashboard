package userdashboard

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/gorilla/mux"
)

func extractParam(r *http.Request, key string) (string, error) {
	vars := mux.Vars(r)
	if vars[key] == "" {
		return "", fmt.Errorf("missing %s parameter", key)
	}
	return vars[key], nil
}

func extractIntParam(r *http.Request, key string) (int, error) {
	str, err := extractParam(r, key)
	if err != nil {
		return 0, err
	}
	if str == "" {
		return 0, fmt.Errorf("missing %s parameter", key)
	}
	i, err := strconv.Atoi(str)
	if err != nil {
		return 0, err
	}
	return i, nil
}

func isNumeric(s string) bool {
	_, err := strconv.ParseFloat(s, 64)
	return err == nil
}

func isEmail(s string) bool {
	return strings.Contains(s, "@")
}

func getJson(r *http.Request, target interface{}) error {
	err := json.NewDecoder(r.Body).Decode(target)
	if err != nil {
		return err
	}
	return nil
}

func getJsonWithAuth(r *http.Request, target interface{}) error {
	return getJson(r, target)
}

func getJsonWithValidation(r *http.Request, target interface{}) error {
	return getJson(r, target)
}

func validateEmail(email string) bool {
	return strings.Contains(email, "@")
}

func validatePassword(password string) bool {
	if len(password) < 6 || len(password) > 32 {
		return false
	}
	for _, c := range password {
		if !isAlphaNum(c) {
			return false
		}
	}
	return true
}

func validateUsername(username string) bool {
	if len(username) < 3 || len(username) > 32 {
		return false
	}
	for _, c := range username {
		if !isAlphaNum(c) {
			return false
		}
	}
	return true
}

func toUnixNano(t time.Time) int64 {
	return t.UnixNano()
}

func isAlphaNum(c rune) bool {
	return (c >= 'a' && c <= 'z') || (c >= 'A' && c <= 'Z') || (c >= '0' && c <= '9')
}