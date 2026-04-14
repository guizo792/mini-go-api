package middleware

import (
	"errors"
	"net/http"
	"strings"

	"github.com/guizo792/mini-go-api/api"
	"github.com/guizo792/mini-go-api/internal/tools"
	log "github.com/sirupsen/logrus"
)

var UnauthorizedError = errors.New("Unauthorized Access. Invalid username or token")

func Authorization(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var username string = strings.TrimSpace(r.URL.Query().Get("username"))
		var token string = strings.TrimSpace(r.Header.Get("Authorization"))

		var err error

		if username == "" || token == "" {
			log.Error(UnauthorizedError)
			api.RequestErrorHandler(w, UnauthorizedError)
			return
		}

		var database tools.DatabaseInterface
		database, err = tools.NewDatabase(false)

		if err != nil {
			api.InternalErrorHandler(w)
			return
		}

		var loginDetails *tools.LoginDetails
		loginDetails, err = database.GetUserLoginDetails(username)

		if (err != nil || loginDetails == nil || (token != (*loginDetails).AuthToken)) {
			log.Error(UnauthorizedError)
			api.RequestErrorHandler(w, UnauthorizedError)
			return
		}

		next.ServeHTTP(w, r)
	})	
}
