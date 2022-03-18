package bitly

type PostShortUrlRequestBody struct {
	Domain  string `json:"domain"`
	LongUrl string `json:"long_url"`
}

type PostShortUrlResponseBody struct {
	CreatedAt      string        `json:"created_at"`
	ID             string        `json:"id"`
	Link           string        `json:"link"`
	CustomBitlinks []interface{} `json:"custom_bitlinks"`
	LongURL        string        `json:"long_url"`
	Archived       bool          `json:"archived"`
	Tags           []interface{} `json:"tags"`
	Deeplinks      []interface{} `json:"deeplinks"`
	References     struct {
		Group string `json:"group"`
	} `json:"references"`
}

type PostExpandRequestBody struct {
	BitlinkId string `json:"bitlink_id"`
}

type PostExpandResponseBody struct {
	CreatedAt string `json:"created_at"`
	Link      string `json:"link"`
	ID        string `json:"id"`
	LongURL   string `json:"long_url"`
}
