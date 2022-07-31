package main

import (
	"encoding/json"
	"io"
	"net/http"
	"os"
	"regexp"
	"strconv"
)

// Estrutura do json de resposta de erro
type error struct {
	Error string `json:"error"`
}

// Retorna um JSON com uma mensagem de erro
func (e *error) SpitError(w http.ResponseWriter, message string, code int) {
	errorResponse := error{
		Error: "resource not found",
	}

	jsonBytes, _ := json.Marshal(errorResponse)

	w.WriteHeader(code)
	w.Write(jsonBytes)
}

// Estrutura do json de resposta
type code struct {
	Code     int    `json:"code"`
	ImageUrl string `json:"image"`
}

// Handler de requests
type codeHandler struct{}

// Regex para roteamento
var (
	getRacoonResponse      = regexp.MustCompile(`^\/(\d+)$`)
	getRacoonImageResponse = regexp.MustCompile(`^\/images\/(\d+)$`)
)

// Roteador
func (h *codeHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")

	url := r.URL.Path

	isGet := r.Method == http.MethodGet

	switch {
	case isGet && getRacoonImageResponse.MatchString(url):
		h.getImage(w, r)
		return
	case isGet && getRacoonResponse.MatchString(url):
		h.get(w, r)
		return
	default:
		h.notFound(w, r)
	}
}

// Endpoint JSON
func (h *codeHandler) get(w http.ResponseWriter, r *http.Request) {
	url := r.URL.Path

	matches := getRacoonResponse.FindStringSubmatch(url)

	if len(matches) < 2 {
		h.notFound(w, r)
		return
	}

	httpCode, err := strconv.Atoi(matches[1])

	if err != nil {
		h.badRequest(w, r)
		return
	}

	if httpCode >= 600 {
		h.notFound(w, r)
		return
	}

	codeR := code{
		Code:     httpCode,
		ImageUrl: "/images/" + matches[1],
	}

	jsonBytes, err := json.Marshal(codeR)

	if err != nil {
		h.internalServerError(w, r)
	}

	w.Write(jsonBytes)
}

// Endpoint que envia a imagem correspondente
func (h *codeHandler) getImage(w http.ResponseWriter, r *http.Request) {
	url := r.URL.Path

	matches := getRacoonImageResponse.FindStringSubmatch(url)

	if len(matches) < 2 {
		h.notFound(w, r)
		return
	}

	httpCode, err := strconv.Atoi(matches[1])

	if err != nil {
		h.badRequest(w, r)
		return
	}

	if httpCode >= 600 {
		h.notFound(w, r)
		return
	}

	dir, err := os.Getwd()

	imageUrl := dir + "/static/" + matches[1] + ".png"

	img, err := os.Open(imageUrl)

	if err != nil {
		h.internalServerError(w, r)
	}

	defer img.Close()

	w.Header().Set("Content-Type", "image/png")
	io.Copy(w, img)
}

// Respostas de erro
func (h *codeHandler) notFound(w http.ResponseWriter, r *http.Request) {
	errorResponse := new(error)
	errorResponse.SpitError(w, "not found", http.StatusNotFound)
}

func (h *codeHandler) internalServerError(w http.ResponseWriter, r *http.Request) {
	errorResponse := new(error)
	errorResponse.SpitError(w, "internal server error", http.StatusNotFound)
}

func (h *codeHandler) badRequest(w http.ResponseWriter, r *http.Request) {
	errorResponse := new(error)
	errorResponse.SpitError(w, "bad request", http.StatusNotFound)
}

func main() {
	mux := http.NewServeMux()

	codeH := &codeHandler{}

	mux.Handle("/", codeH)

	http.ListenAndServe("localhost:3000", mux)
}
