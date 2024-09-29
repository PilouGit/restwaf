package engine

import (
	"fmt"
	"io"
	"net/http"
	"restwaf/internal/cache"
	"restwaf/internal/config"
	"restwaf/internal/model"
	"restwaf/internal/siem"
	"restwaf/internal/validator"

	"github.com/corazawaf/coraza/v3"
	"github.com/corazawaf/coraza/v3/debuglog"
	"github.com/corazawaf/coraza/v3/types"
	"go.uber.org/zap"
)

type Engine struct {
	configuration    *config.Configuration
	requestCache     *cache.RequestGCache
	openApiValidator *validator.OpenApiValidator
	Waf              *validator.WafValidator
	logger           *zap.Logger
	siem             *siem.Siem
}

func (engine *Engine) ProcessRequest(request *model.Request) *validator.ValidatorResponse {
	return engine.Waf.ProcessRequest(request)

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
	logger, _ := zap.NewProduction()
	defer logger.Sync()
	engine.logger = logger
	var error = engine.initRequestCache()
	if error != nil {
		return error
	}
	error = engine.initRestWaf()
	if error != nil {
		engine.logger.Error("unable to create RestWaf instance", zap.Error(error))
		return error
	}
	error = engine.initWaf()

	if error != nil {
		engine.logger.Error("unable to create Waf instance", zap.Error(error))
		return error
	}
	error = engine.initSiem()
	if error != nil {
		engine.logger.Error("unable to create Siem instance", zap.Error(error))
		return error
	}
	return error
}
func logError(error types.MatchedRule) {
	msg := error.ErrorLog()
	fmt.Printf("[logError][%s] %s\n", error.Rule().Severity(), msg)
}
func (engine *Engine) initWaf() error {
	var wafConfiguration = engine.configuration.WafConfiguration
	if wafConfiguration.Enabled {
		var conf = coraza.NewWAFConfig()
		for _, directiveFromFile := range wafConfiguration.DirectivesFromFile {
			conf = conf.WithDirectivesFromFile(directiveFromFile)
		}
		conf = conf.WithErrorCallback(logError)
		conf = conf.WithDebugLogger(debuglog.Default())

		waf, err := coraza.NewWAF(conf)
		if err != nil {
			engine.logger.Error("unable to create waf instance", zap.Error(err))
			return err
		}

		engine.Waf = new(validator.WafValidator)
		engine.Waf.CreateWafValidator(&waf)

	}
	return nil
}
func (engine *Engine) initRequestCache() error {

	var cacheConfiguration = engine.configuration.GlobalConfiguration.CacheConfiguration
	if cacheConfiguration.IsGCache() {
		var gcache = cache.CreateRequestGCache(cacheConfiguration.Nbtransaction)
		engine.requestCache = gcache
	}
	return nil
}

func (engine *Engine) initSiem() error {
	engine.siem = siem.CreateSiem()
	return engine.siem.Initialization(engine.configuration.GlobalConfiguration.OpenSearchConfiguration)

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
