package bitly

func (us *UrlServiceConfig) TransformUrl(longUrl string) (string, error) {
	return longUrl, nil
}

func (us *UrlServiceConfig) ReTransformUrl(longUrl string) (string, error) {
	return longUrl, nil
}
