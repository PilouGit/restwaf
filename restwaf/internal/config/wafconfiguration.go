package config

type WafConfiguration struct {
	Enabled               bool     `json:"enabled"  validate:"required"`
	DirectivesFromFile    []string `json:"directivesFromFile"`
	WithOpenApiDirectives bool     `json:"withopenapidirectives"`
}

/*func (wafConfiguration *WafConfiguration) Validate() error {
for

}*/
