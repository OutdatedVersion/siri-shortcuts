package main

import (
	"os"
	"strings"

	"github.com/valyala/fasthttp"
)

const (
	// Test
	headerName = "Authorization"

	// Test
	authType = "Bearer "
)

var tokens = strings.Split(os.Getenv("VALID_TOKENS"), ",")

func getTokenFromHeader(context *fasthttp.RequestCtx) string {
	header := context.Request.Header.Peek(headerName)

	if header == nil {
		context.Error("Missing/empty 'Authorization' header", fasthttp.StatusBadRequest)
		return ""
	}

	text := string(header)

	if !strings.HasPrefix(text, authType) {
		context.Error("Invalid authorization type", fasthttp.StatusBadRequest)
		return ""
	}

	token := text[len(authType):]

	if token == "" || len(token) < 1 {
		context.Error("Empty token", fasthttp.StatusBadRequest)
		return ""
	}

	return token
}

// IsAuthorized verifies whether or not the client may run the provided action
//
// At the moment, any token has the ability to execute every action.
func IsAuthorized(context *fasthttp.RequestCtx, permission string) bool {
	providedToken := getTokenFromHeader(context)

	if providedToken == "" {
		return false
	}

	for _, token := range tokens {
		if providedToken == token {
			return true
		}
	}

	context.Error("You may not perform "+permission, fasthttp.StatusUnauthorized)

	return false
}
