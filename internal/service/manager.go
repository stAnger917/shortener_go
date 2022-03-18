package service

import "context"

func (us *UrlServiceManager) TransformUrl(ctx context.Context, longUrl string) (string, error) {
	res, err := us.urlService.TransformUrl(ctx, longUrl)
	if err != nil {
		return "", err
	}
	return res, nil
}

func (us *UrlServiceManager) ReTransformUrl(ctx context.Context, bitlinkId string) (string, error) {
	res, err := us.urlService.ReTransformUrl(ctx, bitlinkId)
	if err != nil {
		return "", err
	}
	return res, nil
}
