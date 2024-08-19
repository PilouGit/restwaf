package config

type OpenApi struct {
	Enabled bool   `json:"enabled"  validate:"required"`
	Url     string `json:"url"`
}
