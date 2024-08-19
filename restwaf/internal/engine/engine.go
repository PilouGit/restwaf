package engine

import (
	"io"
	"net/http"
	"restwaf/internal/cache"
	"restwaf/internal/config"
	"restwaf/internal/validator"
)

type Engine struct {
	configuration    *config.Configuration
	requestCache     *cache.RequestGCache
	openApiValidator *validator.OpenApiValidator
}

func (engine *Engine) CreateFromConfigurationFile(filename string) error {

	configuration := new(config.Configuration)
	var error = configuration.ReadConfiguration(filename)
	if error != nil {
		return error
	}
	error = configuration.Validate()
	if error != nil {
		return error
	}
	engine.configuration = configuration
	return nil
}
func (engine *Engine) Init() error {
	var error = engine.initRequestCache()
	if error != nil {
		return error
	}
	error = engine.initRestWaf()

	return error
}
func (engine *Engine) initRequestCache() error {

	var cacheConfiguration = engine.configuration.GlobalConfiguration.CacheConfiguration
	if cacheConfiguration.IsGCache() {
		var gcache = cache.CreateRequestGCache(cacheConfiguration.Nbtransaction)
		engine.requestCache = gcache
	}
	return nil
}

func (engine *Engine) initRestWaf() error {

	var openApiConfiguration = engine.configuration.OpenApiConfiguration
	if openApiConfiguration != nil && openApiConfiguration.Enabled {
		var openApiValidator = new(validator.OpenApiValidator)
		response, error := http.Get(openApiConfiguration.Url)
		if error != nil {
			return error
		}
		// read response body
		body, error := io.ReadAll(response.Body)
		if error != nil {
			return error
		}
		// close response body
		response.Body.Close()
		error = openApiValidator.CreateOpenApiValidator(body)
		if error != nil {
			return error
		}
		engine.openApiValidator = openApiValidator

	}
	return nil
}
