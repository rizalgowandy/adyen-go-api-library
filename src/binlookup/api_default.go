/*
 * Adyen BinLookup API
 *
 * The BIN Lookup API provides endpoints for retrieving information, such as cost estimates, and 3D Secure supported version based on a given BIN.
 *
 * API version: 50
 * Contact: support@adyen.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package binlookup

import (
	_context "context"
	_nethttp "net/http"

	"github.com/adyen/adyen-go-api-library/v3/src/common"
)

// BinLookup BinLookup service
type BinLookup common.Service

/*
Get3dsAvailability Checks 3D Secure availability.
Verifies whether 3D Secure is available for the specified BIN or card brand. For 3D Secure 2, this endpoint also returns device fingerprinting keys.  For more information, refer to [3D Secure 2](https://docs.adyen.com/checkout/3d-secure/native-3ds2).
 * @param request ThreeDSAvailabilityRequest - reference of ThreeDSAvailabilityRequest).
 * @param ctxs ..._context.Context - optional, for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@return ThreeDSAvailabilityResponse
*/
func (a BinLookup) Get3dsAvailability(req *ThreeDSAvailabilityRequest, ctxs ..._context.Context) (ThreeDSAvailabilityResponse, *_nethttp.Response, error) {
	res := &ThreeDSAvailabilityResponse{}
	httpRes, err := a.Client.MakeHTTPPostRequest(req, res, a.BasePath()+"/get3dsAvailability", ctxs...)
	return *res, httpRes, err
}

/*
GetCostEstimate Gets a cost estimate.
Use the Adyen Cost Estimation API to pre-calculate interchange and scheme fee costs. Knowing these costs prior actual payment authorisation gives you an opportunity to charge those costs to the cardholder, if necessary.  To retrieve this information, make the call to the &#x60;/getCostEstimate&#x60; endpoint. The response to this call contains the amount of the interchange and scheme fees charged by the network for this transaction, and also which surcharging policy is possible (based on current regulations).  &gt; Since not all information is known in advance (for example, if the cardholder will successfully authenticate via 3D Secure or if you also plan to provide additional Level 2/3 data), the returned amounts are based on a set of assumption criteria you define in the &#x60;assumptions&#x60; parameter.
 * @param request CostEstimateRequest - reference of CostEstimateRequest).
 * @param ctxs ..._context.Context - optional, for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@return CostEstimateResponse
*/
func (a BinLookup) GetCostEstimate(req *CostEstimateRequest, ctxs ..._context.Context) (CostEstimateResponse, *_nethttp.Response, error) {
	res := &CostEstimateResponse{}
	httpRes, err := a.Client.MakeHTTPPostRequest(req, res, a.BasePath()+"/getCostEstimate", ctxs...)
	return *res, httpRes, err
}
