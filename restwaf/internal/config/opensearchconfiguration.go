package config

type OpenSearchConfiguration struct {
	Insecureskipverify bool     `json:"insecureskipverify"  validate:"required"`
	Urls               []string `json:"url" validate:"required"`
	Username           string   `json:"username" validate:"required"`
	Password           string   `json:"password" validate:"required"`
	Index              string   `json:"index" validate:"required"`
}
