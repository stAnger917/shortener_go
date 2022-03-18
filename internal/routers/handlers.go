package routers

import (
	"context"
	"encoding/json"
	"net/http"
	"shortener/pkg/logging"
	"strings"
)

type HandlePostShortenReqBody struct {
	Url string `json:"url"`
}

func (h *Handler) HandlePostShorten(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if req.Method == http.MethodPost {
		if req.URL.Path == "/" {
			var requestBody HandlePostShortenReqBody
			decoder := json.NewDecoder(req.Body)
			err := decoder.Decode(&requestBody)
			if err != nil {
				logging.EasyLogError("routers", "failed to parse request body for method",
					" POST /", err)
				http.Error(w, "Error! Failed to parse request body!", http.StatusBadRequest)
			}

			res, err := h.services.UrlService.TransformUrl(context.TODO(), requestBody.Url)
			if err != nil {
				logging.EasyLogError("routers", "failed to get short url for method",
					" POST /", err)
				http.Error(w, "Error! Failed to generate short url!", http.StatusBadRequest)
			}
			result, err := json.Marshal(res)
			if err != nil {
				logging.EasyLogError("routers", "failed to parse response body for method",
					" POST /", err)
				http.Error(w, "Error! Failed to return short url!", http.StatusBadRequest)
			}

			w.WriteHeader(201)
			w.Write(result)
			return
		}
	} else if req.Method == http.MethodGet {
		path := strings.Trim(req.URL.Path, "/")
		res, err := h.services.UrlService.ReTransformUrl(context.TODO(), path)
		if err != nil {
			logging.EasyLogError("routers", "failed to get response body for method",
				" GET /:id", err)
			http.Error(w, "Error! Failed to return full url!", http.StatusBadRequest)
		}
		w.Header().Add("Location", res)
		w.WriteHeader(307)
		return
	} else {
		logging.EasyLogInfo("routers", "request denied - unhallowed method for: ",
			" POST /")
		http.Error(w, "Error! Method not allowed", http.StatusBadRequest)
	}
}
