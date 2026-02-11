package handlers

import (
	"errors"
	"html/template"
	"net/http"

	"ascii-art-web/internal/ascii"
)

type Handler struct {
	tmpl *template.Template
}

type PageData struct {
	Text string
	Banner string
	Result string
	Error string
}

func NewHandler(t *template.Template) *Handler {
	return &Handler{tmpl: t}
}

func (h *Handler) Home(w http.ResponseWriter, r *http.Request) {
	// Handle non-existent paths
	if r.URL.Path != "/" {
		http.Error(w, "404 Page Not Found", http.StatusNotFound)
		return
	}
	
	// Accept only GET requests
	if r.Method != http.MethodGet {
		http.Error(w, "Method Not Allowed", http.StatusBadRequest)
		return
	}

	h.tmpl.Execute(w, nil)
}

func (h *Handler) Generate(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	// Handler expects data to come as an html form named text and bannner for the input text and banner file name respectively
	text := r.FormValue("text")
	banner := r.FormValue("banner")

	if text == "" || banner == "" {
		w.WriteHeader(http.StatusBadRequest)
		h.tmpl.Execute(w, PageData{
			Error: "Text and banner are required",
		})
		return
	}

	result, err := ascii.Generate(text, banner)
	if err != nil {
		switch {
		case errors.Is(err, ascii.ErrInvalidInput):
			w.WriteHeader(http.StatusBadRequest)
		case errors.Is(err, ascii.ErrBannerNotFound):
			w.WriteHeader(http.StatusNotFound)
		case errors.Is(err, ascii.ErrBannerUnreadable):
			w.WriteHeader(http.StatusInternalServerError)
		case errors.Is(err, ascii.ErrInvalidBannerFormat):
			w.WriteHeader(http.StatusInternalServerError)
		default:
			w.WriteHeader(http.StatusInternalServerError)
		}

		h.tmpl.Execute(w, PageData{
			Text: text,
			Banner: banner,
			Error: err.Error(),
		})
		return
	}

	h.tmpl.Execute(w, PageData{
		Text: text,
		Banner: banner,
		Result: result,
	})
}
//check that not all errors give 404,it should be able to do,other errors
//based on the changes made