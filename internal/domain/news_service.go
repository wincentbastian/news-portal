package domain

type NewsService interface {
	GetBitcoinNews() ([]NewsArticle, error)
}
