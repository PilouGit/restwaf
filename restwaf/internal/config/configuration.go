package config

type Configuration struct {
	GlobalConfiguration  *GlobalConfiguration `json:"configuration" validate:"required"`
	OpenApiConfiguration *OpenApi             `json:"openapi" validate:"required"`
	WafConfiguration     *WafConfiguration    `json:"waf" validate:"required"`
}
