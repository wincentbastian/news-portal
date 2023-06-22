package delivery

import (
	"fmt"
	"html/template"
	"net/http"
	"news-portal/internal/infrastructure"
	"news-portal/internal/usecase"
)

func newsHandler(w http.ResponseWriter, r *http.Request) {
	// Membuat instance NewsUsecase
	newsRepo := infrastructure.NewNewsRepository()
	newsUsecase := usecase.NewNewsUsecase(*newsRepo)

	// Mengambil daftar berita terkait Bitcoin
	articles, err := newsUsecase.GetBitcoinNews()
	if err != nil {
		http.Error(w, "Failed to get news", http.StatusInternalServerError)
		return
	}

	// Memuat template HTML
	tmpl, err := template.ParseFiles("../web/templates/index.html")
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to load template: %v", err), http.StatusInternalServerError)
		return
	}

	// Menampilkan halaman HTML dengan daftar berita
	err = tmpl.Execute(w, articles)
	if err != nil {
		http.Error(w, "Failed to render template", http.StatusInternalServerError)
		return
	}
}

func searchHandler(w http.ResponseWriter, r *http.Request) {
	// Membuat instance NewsUsecase
	newsRepo := infrastructure.NewNewsRepository()
	newsUsecase := usecase.NewNewsUsecase(*newsRepo)

	articles, err := newsUsecase.GetSearchNews(r.FormValue("searchValue"))
	if err != nil {
		http.Error(w, "Failed to get news", http.StatusInternalServerError)
		return
	}

	// Memuat template HTML
	tmpl, err := template.ParseFiles("../web/templates/index.html")
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to load template: %v", err), http.StatusInternalServerError)
		return
	}

	// Menampilkan halaman HTML dengan daftar berita
	err = tmpl.Execute(w, articles)
	if err != nil {
		http.Error(w, "Failed to render template", http.StatusInternalServerError)
		return
	}
}

// InitializeHandlers menginisialisasi semua handler
func InitializeHandlers() {
	http.HandleFunc("/", newsHandler)
	http.HandleFunc("/search", searchHandler)
}
