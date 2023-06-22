package usecase

import (
	"news-portal/internal/domain"
	"news-portal/internal/infrastructure"
)

type NewsUsecase interface {
	GetBitcoinNews() ([]domain.NewsArticle, error)
	GetSearchNews(keyword string) ([]domain.NewsArticle, error)
}

type newsUsecase struct {
	newsRepo infrastructure.NewsRepository
}

func NewNewsUsecase(newsRepo infrastructure.NewsRepository) NewsUsecase {
	return &newsUsecase{
		newsRepo: newsRepo,
	}
}

func (uc *newsUsecase) GetBitcoinNews() ([]domain.NewsArticle, error) {
	return uc.newsRepo.GetBitcoinNews()
}

func (uc *newsUsecase) GetSearchNews(keyword string) ([]domain.NewsArticle, error) {
	return uc.newsRepo.GetSearchNews(keyword)
}
