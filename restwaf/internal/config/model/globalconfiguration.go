package modelconfig

type GlobalConfiguration struct {
	Port               int                 `json:"port"  validate:"required"`
	CacheConfiguration *CacheConfiguration `json:"cache"  validate:"required"`
}
