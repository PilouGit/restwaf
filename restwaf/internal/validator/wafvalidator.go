package validator

import (
	"fmt"
	"log"
	"restwaf/internal/model"

	"github.com/corazawaf/coraza/v3"
	"github.com/corazawaf/coraza/v3/types"
)

type WafValidator struct {
	waf *coraza.WAF
}

func (wafvalidator *WafValidator) CreateWafValidator(corazza *coraza.WAF) {
	wafvalidator.waf = corazza

}
func logError(error types.MatchedRule) {
	msg := error.ErrorLog()
	fmt.Printf("[logError][%s] %s\n", error.Rule().Severity(), msg)
}
func (wafvalidator *WafValidator) ProcessRequest(request *model.Request) {
	var id = request.Id
	var coraza = *wafvalidator.waf
	var transaction = coraza.NewTransactionWithID(id)
	defer transaction.ProcessLogging()
	if transaction.IsRuleEngineOff() {
		log.Printf("coraza is off")
	}
	httprequest, _ := request.ToHttpRequest()
	for key, values := range httprequest.Header {
		for _, v := range values {
			transaction.AddRequestHeader(key, v)
		}
	}
	transaction.WriteRequestBody(request.Body)
	interruption := transaction.ProcessRequestHeaders()
	log.Printf("Interruption %v", interruption)
	interruption, err := transaction.ProcessRequestBody()
	log.Printf("Interruption %v", interruption)
	log.Printf("Error %v", err)

}
