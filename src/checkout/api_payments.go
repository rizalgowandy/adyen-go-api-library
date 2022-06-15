/*
 * Adyen Checkout API
 *
 * Adyen Checkout API provides a simple and flexible way to initiate and authorise online payments. You can use the same integration for payments made with cards (including 3D Secure), mobile wallets, and local payment methods (for example, iDEAL and Sofort).  This API reference provides information on available endpoints and how to interact with them. To learn more about the API, visit [Checkout documentation](https://docs.adyen.com/online-payments).  ## Authentication Each request to the Checkout API must be signed with an API key. For this, obtain an API Key from your Customer Area, as described in [How to get the API key](https://docs.adyen.com/development-resources/api-credentials#generate-api-key). Then set this key to the `X-API-Key` header value, for example:  ``` curl -H \"Content-Type: application/json\" \\ -H \"X-API-Key: Your_Checkout_API_key\" \\ ... ``` Note that when going live, you need to generate a new API Key to access the [live endpoints](https://docs.adyen.com/development-resources/live-endpoints).  ## Versioning Checkout API supports [versioning](https://docs.adyen.com/development-resources/versioning) using a version suffix in the endpoint URL. This suffix has the following format: \"vXX\", where XX is the version number.  For example: ``` https://checkout-test.adyen.com/v68/payments ```
 *
 * API version: 68
 * Contact: developer-experience@adyen.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package checkout

import (
	_context "context"
	_nethttp "net/http"

	"github.com/adyen/adyen-go-api-library/v5/src/common"
)

/*
Donations Start a transaction for donations
Takes in the donation token generated by the &#x60;/payments&#x60; request and uses it to make the donation for the donation account specified in the request.  For more information, see [Donations](https://docs.adyen.com/online-payments/donations).
 * @param request PaymentDonationRequest - reference of PaymentDonationRequest).
 * @param ctxs ..._context.Context - optional, for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@return DonationResponse
*/
func (a Checkout) Donations(req *PaymentDonationRequest, ctxs ..._context.Context) (DonationResponse, *_nethttp.Response, error) {
    res := &DonationResponse{}
    httpRes, err := a.Client.MakeHTTPPostRequest(req, res, a.BasePath() + "/donations", ctxs...)
    return *res, httpRes, err
}

/*
PaymentMethods Get a list of available payment methods
Queries the available payment methods for a transaction based on the transaction context (like amount, country, and currency). Besides giving back a list of the available payment methods, the response also returns which input details you need to collect from the shopper (to be submitted to &#x60;/payments&#x60;).  Although we highly recommend using this endpoint to ensure you are always offering the most up-to-date list of payment methods, its usage is optional. You can, for example, also cache the &#x60;/paymentMethods&#x60; response and update it once a week.
 * @param request PaymentMethodsRequest - reference of PaymentMethodsRequest). 
 * @param ctxs ..._context.Context - optional, for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@return PaymentMethodsResponse
*/
func (a Checkout) PaymentMethods(req *PaymentMethodsRequest, ctxs ..._context.Context) (PaymentMethodsResponse, *_nethttp.Response, error) {
    res := &PaymentMethodsResponse{}
    httpRes, err := a.Client.MakeHTTPPostRequest(req, res, a.BasePath() + "/paymentMethods", ctxs...)
    return *res, httpRes, err
}

/*
Payments Start a transaction
Sends payment parameters (like amount, country, and currency) together with other required input details collected from the shopper. To know more about required parameters for specific payment methods, refer to our [payment method guides](https://docs.adyen.com/payment-methods).  The response depends on the [payment flow](https://docs.adyen.com/payment-methods#payment-flow): * For a direct flow, the response includes a &#x60;pspReference&#x60; and a &#x60;resultCode&#x60; with the payment result, for example **Authorised** or **Refused**.  * For a redirect or additional action, the response contains an &#x60;action&#x60; object. 
 * @param request PaymentRequest - reference of PaymentRequest). 
 * @param ctxs ..._context.Context - optional, for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@return PaymentResponse
*/
func (a Checkout) Payments(req *PaymentRequest, ctxs ..._context.Context) (PaymentResponse, *_nethttp.Response, error) {
    res := &PaymentResponse{}

	adyenLib := &CommonField{
		Name:    common.LibName,
		Version: common.LibVersion,
	}
	if req.ApplicationInfo == nil {
		req.ApplicationInfo = &ApplicationInfo{}
	}
	req.ApplicationInfo.AdyenLibrary = adyenLib

    httpRes, err := a.Client.MakeHTTPPostRequest(req, res, a.BasePath() + "/payments", ctxs...)
    return *res, httpRes, err
}

/*
PaymentsDetails Submit details for a payment
Submits details for a payment created using &#x60;/payments&#x60;. This step is only needed when no final state has been reached on the &#x60;/payments&#x60; request, for example when the shopper was redirected to another page to complete the payment.  
 * @param request DetailsRequest - reference of DetailsRequest). 
 * @param ctxs ..._context.Context - optional, for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@return PaymentDetailsResponse
*/
func (a Checkout) PaymentsDetails(req *DetailsRequest, ctxs ..._context.Context) (PaymentDetailsResponse, *_nethttp.Response, error) {
    res := &PaymentDetailsResponse{}
    httpRes, err := a.Client.MakeHTTPPostRequest(req, res, a.BasePath() + "/payments/details", ctxs...)
    return *res, httpRes, err
}

/*
Sessions Create a payment session
Creates a payment session for [Web Drop-in](https://docs.adyen.com/online-payments/web-drop-in) and [Web Components](https://docs.adyen.com/online-payments/web-components) integrations.  The response contains encrypted payment session data. The front end then uses the session data to make any required server-side calls for the payment flow.  You get the payment outcome asynchronously, in an [AUTHORISATION](https://docs.adyen.com/api-explorer/#/Webhooks/latest/post/AUTHORISATION) webhook.
 * @param request CreateCheckoutSessionRequest - reference of CreateCheckoutSessionRequest). 
 * @param ctxs ..._context.Context - optional, for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@return CreateCheckoutSessionResponse
*/
func (a Checkout) Sessions(req *CreateCheckoutSessionRequest, ctxs ..._context.Context) (CreateCheckoutSessionResponse, *_nethttp.Response, error) {
    res := &CreateCheckoutSessionResponse{}

    // add ApplicationInfo data 
    adyenLib := &CommonField{
		Name:    common.LibName,
		Version: common.LibVersion,
	}
	if req.ApplicationInfo == nil {
		req.ApplicationInfo = &ApplicationInfo{}
	}
	req.ApplicationInfo.AdyenLibrary = adyenLib

    httpRes, err := a.Client.MakeHTTPPostRequest(req, res, a.BasePath() + "/sessions", ctxs...)
    return *res, httpRes, err
}
