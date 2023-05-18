package auth

import (
	"strings"
	"net/http"
	"errors"
)

func GetAPIKey(headers http.Header) (string, error) {
	val:=headers.Get("Authorization")
	if val == "" {
		return "", errors.New("no Authorization header found")
	}
	vals:=strings.Split(val, " ")
	if len(vals)!= 2 {
        return "", errors.New("Invalid-Malformed Authorization header")
    }
	if vals[0]!= "ApiKey"{
		return "", errors.New("Invalid-Malformed first-part Authorization header")
	}
	return vals[1], nil
}