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
// AdditionalDataWallets struct for AdditionalDataWallets
type AdditionalDataWallets struct {
	// The Android Pay token retrieved from the SDK.
	AndroidpayToken string `json:"androidpay.token,omitempty"`
	// The Mastercard Masterpass Transaction ID retrieved from the SDK.
	MasterpassTransactionId string `json:"masterpass.transactionId,omitempty"`
	// The Apple Pay token retrieved from the SDK.
	PaymentToken string `json:"payment.token,omitempty"`
	// The Google Pay token retrieved from the SDK.
	PaywithgoogleToken string `json:"paywithgoogle.token,omitempty"`
	// The Samsung Pay token retrieved from the SDK.
	SamsungpayToken string `json:"samsungpay.token,omitempty"`
	// The Visa Checkout Call ID retrieved from the SDK.
	VisacheckoutCallId string `json:"visacheckout.callId,omitempty"`
}
