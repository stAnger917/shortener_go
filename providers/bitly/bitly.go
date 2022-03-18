package bitly

type UrlService interface {
	TransformUrl(longUrl string) (string, error)
	ReTransformUrl(modUrl string) (string, error)
}

type UrlServiceConfig struct {
	Token     string
	GroupGuid string
	Domain    string
}

func NewUrlServiceConfiguration(token, guid, domain string) *UrlServiceConfig {
	return &UrlServiceConfig{
		Token:     token,
		GroupGuid: guid,
		Domain:    domain,
	}
}

type ModUrlService struct {
	ServiceManager UrlService
}

func NewUrlService(token, guid, domain string) *ModUrlService {
	return &ModUrlService{
		ServiceManager: NewUrlServiceConfiguration(token, guid, domain),
	}
}
