package bitly

import "context"

type UrlService interface {
	TransformUrl(ctx context.Context, longUrl string) (string, error)
	ReTransformUrl(ctx context.Context, bitlinkId string) (string, error)
}

type UrlServiceConfig struct {
	Token  string
	URL    string
	Domain string
}

func NewUrlServiceConfiguration(token, guid, domain string) *UrlServiceConfig {
	return &UrlServiceConfig{
		Token:  token,
		URL:    guid,
		Domain: domain,
	}
}

type ModUrlService struct {
	ServiceManager UrlService
}

func NewUrlService(token, url, domain string) *ModUrlService {
	return &ModUrlService{
		ServiceManager: NewUrlServiceConfiguration(token, url, domain),
	}
}
