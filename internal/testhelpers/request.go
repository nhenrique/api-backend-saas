package testhelpers

import "net/http"

func AuthRequest(req *http.Request, token string) {
	req.Header.Set("Authorization", "Bearer "+token)
}
