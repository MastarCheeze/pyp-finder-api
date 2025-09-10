package internal

import (
	"net/http"
)

func PageExists(url string) bool {
	resp, err := http.Head(url)
	if err != nil || resp.StatusCode != http.StatusOK {
		return false
	} else {
		return true
	}
}
