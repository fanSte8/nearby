package httperrors

import (
	"log/slog"
	"nearby/common/jsonutils"
	"net/http"
)

type HttpErrors struct {
	logger *slog.Logger
}

func NewHttpErrors(logger *slog.Logger) HttpErrors {
	return HttpErrors{logger}
}

func (errors HttpErrors) logError(r *http.Request, err error) {
	errors.logger.Error("Request error", "error", err, "requestMethod", r.Method, "requestUrl", r.URL.String())
}

func (errors HttpErrors) errorResponse(w http.ResponseWriter, r *http.Request, status int, message any) {
	env := map[string]any{"error": message}

	err := jsonutils.WriteJSON(w, status, env, nil)
	if err != nil {
		errors.logError(r, err)
		w.WriteHeader(500)
	}
}

func (errors HttpErrors) ServerErrorResponse(w http.ResponseWriter, r *http.Request, err error) {
	errors.logError(r, err)

	errors.errorResponse(w, r, http.StatusInternalServerError, "the server encountered a problem and could not process your request")
}

func (errors HttpErrors) NotFoundResponse(w http.ResponseWriter, r *http.Request) {
	errors.errorResponse(w, r, http.StatusNotFound, "the requested resource could not be found")
}

func (errors HttpErrors) BadRequestResponse(w http.ResponseWriter, r *http.Request, err error) {
	errors.errorResponse(w, r, http.StatusBadRequest, err.Error())
}

func (errors HttpErrors) InvalidCredentialsResponse(w http.ResponseWriter, r *http.Request) {
	message := "invalid authentication credentials"
	errors.errorResponse(w, r, http.StatusUnauthorized, message)
}

func (errors HttpErrors) InvalidAuthenticationTokenResponse(w http.ResponseWriter, r *http.Request) {
	errors.errorResponse(w, r, http.StatusUnauthorized, "invalid or missing authentication token")
}

func (errors HttpErrors) AuthenticationRequiredResponse(w http.ResponseWriter, r *http.Request) {
	errors.errorResponse(w, r, http.StatusUnauthorized, "you must be authenticated to access this resource")
}

func (errors HttpErrors) InactiveAccountResponse(w http.ResponseWriter, r *http.Request) {
	message := "your user account must be activated to access this resource"
	errors.errorResponse(w, r, http.StatusForbidden, message)
}
