package page

import (
	"fmt"
	"log/slog"
	"net/http"
)

var specialErrors map[int]string

func init() {
	specialErrors = map[int]string{
		// 400
		http.StatusBadRequest: "Whoops! The server didn't understand your request",
		// 403
		http.StatusForbidden: "You don't have permission to access this link",
		// 404
		http.StatusNotFound: "Whoops! We couldn't find what you're looking for",
		// 418
		http.StatusTeapot: "I'm a little \xf0\x9f\x8d\xb5",
		// 500
		http.StatusInternalServerError: "Whoops! Something went wrong",
	}
}

// 400
func (p *Page) SetBadRequest(w http.ResponseWriter, r *http.Request, log *slog.Logger, err error) {
	statusCode := http.StatusBadRequest
	log.Info("Served 400 Bad Request", "statusCode", statusCode, "path", r.URL.Path, "error", err)
	p.setErrorVars(w, r, statusCode, log)
}

// 404
func (p *Page) SetNotFound(w http.ResponseWriter, r *http.Request, log *slog.Logger) {
	statusCode := http.StatusNotFound
	log.Info("Served 404 Not Found", "statusCode", statusCode, "path", r.URL.Path)
	p.setErrorVars(w, r, statusCode, log)
}

// 500
func (p *Page) SetInternalServerError(w http.ResponseWriter, r *http.Request, log *slog.Logger) {
	statusCode := http.StatusInternalServerError
	log.Info("Served 500 Internal Server Error", "status", statusCode, "path", r.URL.Path)
	p.setErrorVars(w, r, statusCode, log)
}

func (p *Page) SetError(w http.ResponseWriter, r *http.Request, statusCode int, log *slog.Logger) {
	log.Info("Served some error page with generic handler", "status", statusCode, "path", r.URL.Path)
	p.setErrorVars(w, r, statusCode, log)
}

func (p *Page) setErrorVars(w http.ResponseWriter, r *http.Request, statusCode int, log *slog.Logger) {
	message := getErrorMessage(statusCode, r.URL.Path, log)

	p.AddVar("StatusCode", statusCode)
	p.AddVar("Message", message)

	w.WriteHeader(statusCode)
}

var codes = [...]int{
	400, 401, 402, 403, 404, 405, 406, 407, 408, 409, 410, 411,
	412, 413, 414, 415, 416, 417, 418, 421, 422, 423, 424, 426,
	428, 429, 431, 451, 500, 501, 502, 503, 504, 505, 506, 507,
	508, 509, 510, 511,
}

func getErrorMessage(statusCode int, path string, log *slog.Logger) string {
	message, ok := specialErrors[statusCode]
	if ok {
		return message
	}

	for _, e := range codes {
		if e == statusCode {
			return fmt.Sprintf("Whoops, that's an error (%d)", statusCode)
		}
	}

	log.Warn("Served unknown error page", "status", statusCode, "request", path)
	return fmt.Sprintf("Huh? That's not an error I know of (%d)", statusCode)
}
