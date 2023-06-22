package infrastructure

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"news-portal/internal/domain"
)

type NewsRepository struct {
	apiKey string
}

func NewNewsRepository() *NewsRepository {
	// Ganti apiKey dengan API key yang valid dari News API
	apiKey := "2f2e085b9b0344c48be4ae9cda5f5368"

	return &NewsRepository{
		apiKey: apiKey,
	}
}

func (r *NewsRepository) GetBitcoinNews() ([]domain.NewsArticle, error) {
	// Membuat HTTP GET request ke API News API
	url := "https://newsapi.org/v2/everything?q=bitcoin&apiKey=" + r.apiKey
	response, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	// Membaca response body sebagai JSON
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	// Mengubah JSON menjadi struktur NewsResponse
	var newsResponse struct {
		Articles []domain.NewsArticle `json:"articles"`
	}
	if err := json.Unmarshal(body, &newsResponse); err != nil {
		return nil, err
	}

	// Mengembalikan daftar artikel berita
	return newsResponse.Articles, nil
}

func (r *NewsRepository) GetSearchNews(keyword string) ([]domain.NewsArticle, error) {
	url := fmt.Sprintf("https://newsapi.org/v2/everything?q=%s&apiKey=%s", keyword, r.apiKey)
	response, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	var searchResponse struct {
		Articles []domain.NewsArticle `json:"articles"`
	}
	if err := json.Unmarshal(body, &searchResponse); err != nil {
		return nil, err
	}

	return searchResponse.Articles, nil
}
