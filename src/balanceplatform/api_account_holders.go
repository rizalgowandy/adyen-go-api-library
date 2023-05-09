/*
Configuration API

API version: 2
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package balanceplatform

import (
	"context"
	_nethttp "net/http"
	"net/url"
	"strings"

	"github.com/adyen/adyen-go-api-library/v6/src/common"
)

// AccountHoldersApi AccountHoldersApi service
type AccountHoldersApi common.Service

type CreateAccountHolderConfig struct {
	ctx               context.Context
	accountHolderInfo *AccountHolderInfo
}

func (r CreateAccountHolderConfig) AccountHolderInfo(accountHolderInfo AccountHolderInfo) CreateAccountHolderConfig {
	r.accountHolderInfo = &accountHolderInfo
	return r
}

/*
CreateAccountHolder Create an account holder

Creates an account holder linked to a [legal entity](https://docs.adyen.com/api-explorer/#/legalentity/latest/post/legalEntities).



 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return CreateAccountHolderConfig
*/
func (a *AccountHoldersApi) CreateAccountHolderConfig(ctx context.Context) CreateAccountHolderConfig {
	return CreateAccountHolderConfig{
		ctx: ctx,
	}
}

/*
Create an account holder
Creates an account holder linked to a [legal entity](https://docs.adyen.com/api-explorer/#/legalentity/latest/post/legalEntities).
 * @param req AccountHolderInfo - reference of AccountHolderInfo).
 * @param ctxs ..._context.Context - optional, for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@return AccountHolder
*/

func (a *AccountHoldersApi) CreateAccountHolder(r CreateAccountHolderConfig) (AccountHolder, *_nethttp.Response, error) {
	res := &AccountHolder{}
	path := "/accountHolders"
	httpRes, err := a.Client.MakeHTTPPostRequest(r.accountHolderInfo, res, a.BasePath()+path, r.ctx)
	return *res, httpRes, err
}

type GetAccountHolderConfig struct {
	ctx context.Context
	id  string
}

/*
GetAccountHolder Get an account holder

Returns an account holder.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id The unique identifier of the account holder.
 @return GetAccountHolderConfig
*/
func (a *AccountHoldersApi) GetAccountHolderConfig(ctx context.Context, id string) GetAccountHolderConfig {
	return GetAccountHolderConfig{
		ctx: ctx,
		id:  id,
	}
}

/*
Get an account holder
Returns an account holder.
 * @param id The unique identifier of the account holder.
 * @param ctxs ..._context.Context - optional, for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@return AccountHolder
*/

func (a *AccountHoldersApi) GetAccountHolder(r GetAccountHolderConfig) (AccountHolder, *_nethttp.Response, error) {
	res := &AccountHolder{}
	path := "/accountHolders/{id}"
	path = strings.Replace(path, "{"+"id"+"}", url.PathEscape(common.ParameterValueToString(r.id, "id")), -1)
	httpRes, err := a.Client.MakeHTTPGetRequest(res, a.BasePath()+path, r.ctx)
	return *res, httpRes, err
}

type GetAllBalanceAccountsOfAccountHolderConfig struct {
	ctx    context.Context
	id     string
	offset *int32
	limit  *int32
}

// The number of items that you want to skip.
func (r GetAllBalanceAccountsOfAccountHolderConfig) Offset(offset int32) GetAllBalanceAccountsOfAccountHolderConfig {
	r.offset = &offset
	return r
}

// The number of items returned per page, maximum 100 items. By default, the response returns 10 items per page.
func (r GetAllBalanceAccountsOfAccountHolderConfig) Limit(limit int32) GetAllBalanceAccountsOfAccountHolderConfig {
	r.limit = &limit
	return r
}

/*
GetAllBalanceAccountsOfAccountHolder Get all balance accounts of an account holder

Returns a paginated list of the balance accounts associated with an account holder. To fetch multiple pages, use the query parameters.

For example, to limit the page to 5 balance accounts and skip the first 10, use `/accountHolders/{id}/balanceAccounts?limit=5&offset=10`.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id The unique identifier of the account holder.
 @return GetAllBalanceAccountsOfAccountHolderConfig
*/
func (a *AccountHoldersApi) GetAllBalanceAccountsOfAccountHolderConfig(ctx context.Context, id string) GetAllBalanceAccountsOfAccountHolderConfig {
	return GetAllBalanceAccountsOfAccountHolderConfig{
		ctx: ctx,
		id:  id,
	}
}

/*
Get all balance accounts of an account holder
Returns a paginated list of the balance accounts associated with an account holder. To fetch multiple pages, use the query parameters.   For example, to limit the page to 5 balance accounts and skip the first 10, use &#x60;/accountHolders/{id}/balanceAccounts?limit&#x3D;5&amp;offset&#x3D;10&#x60;.
 * @param id The unique identifier of the account holder.
 * @param ctxs ..._context.Context - optional, for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@return PaginatedBalanceAccountsResponse
*/

func (a *AccountHoldersApi) GetAllBalanceAccountsOfAccountHolder(r GetAllBalanceAccountsOfAccountHolderConfig) (PaginatedBalanceAccountsResponse, *_nethttp.Response, error) {
	res := &PaginatedBalanceAccountsResponse{}
	path := "/accountHolders/{id}/balanceAccounts"
	path = strings.Replace(path, "{"+"id"+"}", url.PathEscape(common.ParameterValueToString(r.id, "id")), -1)
	queryString := url.Values{}
	if r.offset != nil {
		common.ParameterAddToQuery(queryString, "offset", r.offset, "")
	}
	if r.limit != nil {
		common.ParameterAddToQuery(queryString, "limit", r.limit, "")
	}
	httpRes, err := a.Client.MakeHTTPGetRequest(res, a.BasePath()+path+"?"+queryString.Encode(), r.ctx)
	return *res, httpRes, err
}

type UpdateAccountHolderConfig struct {
	ctx           context.Context
	id            string
	accountHolder *AccountHolder
}

func (r UpdateAccountHolderConfig) AccountHolder(accountHolder AccountHolder) UpdateAccountHolderConfig {
	r.accountHolder = &accountHolder
	return r
}

/*
UpdateAccountHolder Update an account holder

Updates an account holder. When updating an account holder resource, if a parameter is not provided in the request, it is left unchanged.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id The unique identifier of the account holder.
 @return UpdateAccountHolderConfig
*/
func (a *AccountHoldersApi) UpdateAccountHolderConfig(ctx context.Context, id string) UpdateAccountHolderConfig {
	return UpdateAccountHolderConfig{
		ctx: ctx,
		id:  id,
	}
}

/*
Update an account holder
Updates an account holder. When updating an account holder resource, if a parameter is not provided in the request, it is left unchanged.
 * @param id The unique identifier of the account holder.
 * @param req AccountHolder - reference of AccountHolder).
 * @param ctxs ..._context.Context - optional, for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@return AccountHolder
*/

func (a *AccountHoldersApi) UpdateAccountHolder(r UpdateAccountHolderConfig) (AccountHolder, *_nethttp.Response, error) {
	res := &AccountHolder{}
	path := "/accountHolders/{id}"
	path = strings.Replace(path, "{"+"id"+"}", url.PathEscape(common.ParameterValueToString(r.id, "id")), -1)
	httpRes, err := a.Client.MakeHTTPPatchRequest(r.accountHolder, res, a.BasePath()+path, r.ctx)
	return *res, httpRes, err
}
