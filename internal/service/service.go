package service

import "shortener/providers/bitly"

type Services struct {
	UrlService UrlService
}

type UrlService interface {
	TransformUrl(longUrl string) (string, error)
	ReTransformUrl(modUrl string) (string, error)
}

type Deps struct {
	ProviderService *bitly.ModUrlService
}

type UrlServiceManager struct {
	urlService bitly.UrlService
}

func NewUrlServiceManager(provider bitly.UrlService) *UrlServiceManager {
	return &UrlServiceManager{
		urlService: provider,
	}
}

func AppServices(deps Deps) *Services {
	newService := NewUrlServiceManager(deps.ProviderService.ServiceManager)
	return &Services{
		UrlService: newService,
	}
}
