package service

import "context"

func (us *UrlServiceManager) TransformUrl(ctx context.Context, longUrl string) (string, error) {
	return longUrl, nil
}

func (us *UrlServiceManager) ReTransformUrl(ctx context.Context, modUrl string) (string, error) {
	return modUrl, nil
}
