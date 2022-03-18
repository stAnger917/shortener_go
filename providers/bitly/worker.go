package bitly

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
	"shortener/pkg/logging"
)

func (us *UrlServiceConfig) TransformUrl(ctx context.Context, longUrl string) (string, error) {
	var requestBody = PostShortUrlRequestBody{LongUrl: longUrl, Domain: us.Domain}
	var bearerToken = "Bearer" + us.Token
	jsonRequestBody, err := json.Marshal(requestBody)
	if err != nil {
		logging.EasyLogError("providers", "bitly: error, failed to marshal request body",
			"", err)
		return "", err
	}
	url := us.URL + "shorten"
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonRequestBody))
	if err != nil {
		logging.EasyLogError("providers", "bitly: failed to create new request for method: ",
			"POST /shorten", err)
		return "", err
	}
	req.Header.Add("Authorization", bearerToken)
	req.Header.Add("Content-Type", "application/json")
	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		logging.EasyLogError("providers", "bitly: failed to get response for method: ",
			"POST /shorten", err)
		return "", err
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			logging.EasyLogError("providers", "bitly: failed to close response body for method: ",
				"POST /shorten", err)
		}
	}(resp.Body)
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		logging.EasyLogError("providers", "bitly: failed to parse response body for method: ",
			"POST /shorten", err)
		return "", err
	}
	var formattedRespBody PostShortUrlResponseBody
	err = json.Unmarshal(body, &formattedRespBody)
	if err != nil {
		logging.EasyLogError("providers", "bitly: failed to parse response body into structure for method: ",
			"POST /shorten", err)
		return "", err
	}
	return formattedRespBody.LongURL, nil
}

func (us *UrlServiceConfig) ReTransformUrl(ctx context.Context, longUrl string) (string, error) {
	return longUrl, nil
}
