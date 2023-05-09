/*
Configuration API

API version: 2
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package balanceplatform

import (
	"context"
	_nethttp "net/http"

	"github.com/adyen/adyen-go-api-library/v6/src/common"
)

// BankAccountValidationApi BankAccountValidationApi service
type BankAccountValidationApi common.Service

type ValidateBankAccountIdentificationConfig struct {
	ctx                                        context.Context
	bankAccountIdentificationValidationRequest *BankAccountIdentificationValidationRequest
}

func (r ValidateBankAccountIdentificationConfig) BankAccountIdentificationValidationRequest(bankAccountIdentificationValidationRequest BankAccountIdentificationValidationRequest) ValidateBankAccountIdentificationConfig {
	r.bankAccountIdentificationValidationRequest = &bankAccountIdentificationValidationRequest
	return r
}

/*
ValidateBankAccountIdentification Validate a bank account

Validates bank account identification details. You can use this endpoint to validate bank account details before you [make a transfer](https://docs.adyen.com/api-explorer/transfers/latest/post/transfers) or [create a transfer instrument](https://docs.adyen.com/api-explorer/legalentity/latest/post/transferInstruments).

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ValidateBankAccountIdentificationConfig
*/
func (a *BankAccountValidationApi) ValidateBankAccountIdentificationConfig(ctx context.Context) ValidateBankAccountIdentificationConfig {
	return ValidateBankAccountIdentificationConfig{
		ctx: ctx,
	}
}

/*
Validate a bank account
Validates bank account identification details. You can use this endpoint to validate bank account details before you [make a transfer](https://docs.adyen.com/api-explorer/transfers/latest/post/transfers) or [create a transfer instrument](https://docs.adyen.com/api-explorer/legalentity/latest/post/transferInstruments).
 * @param req BankAccountIdentificationValidationRequest - reference of BankAccountIdentificationValidationRequest).
 * @param ctxs ..._context.Context - optional, for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@return map[string]interface{}
*/

func (a *BankAccountValidationApi) ValidateBankAccountIdentification(r ValidateBankAccountIdentificationConfig) (map[string]interface{}, *_nethttp.Response, error) {
	res := &map[string]interface{}{}
	path := "/validateBankAccountIdentification"
	httpRes, err := a.Client.MakeHTTPPostRequest(r.bankAccountIdentificationValidationRequest, res, a.BasePath()+path, r.ctx)
	return *res, httpRes, err
}
