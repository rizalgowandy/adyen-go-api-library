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
// PaymentMethodsRequest struct for PaymentMethodsRequest
type PaymentMethodsRequest struct {
	// This field contains additional data, which may be required for a particular payment request.  The `additionalData` object consists of entries, each of which includes the key and value.
	AdditionalData map[string]string `json:"additionalData,omitempty"`
	// List of payment methods to be presented to the shopper. To refer to payment methods, use their `paymentMethod.type` from [Payment methods overview](https://docs.adyen.com/payment-methods).  Example: `\"allowedPaymentMethods\":[\"ideal\",\"giropay\"]`
	AllowedPaymentMethods []string `json:"allowedPaymentMethods,omitempty"`
	Amount *Amount `json:"amount,omitempty"`
	// List of payment methods to be hidden from the shopper. To refer to payment methods, use their `paymentMethod.type` from [Payment methods overview](https://docs.adyen.com/payment-methods).  Example: `\"blockedPaymentMethods\":[\"ideal\",\"giropay\"]`
	BlockedPaymentMethods []string `json:"blockedPaymentMethods,omitempty"`
	// The platform where a payment transaction takes place. This field can be used for filtering out payment methods that are only available on specific platforms. Possible values: * iOS * Android * Web
	Channel string `json:"channel,omitempty"`
	// The shopper's country code.
	CountryCode string `json:"countryCode,omitempty"`
	// The merchant account identifier, with which you want to process the transaction.
	MerchantAccount string `json:"merchantAccount"`
	Order *CheckoutOrder `json:"order,omitempty"`
	// The combination of a language code and a country code to specify the language to be used in the payment.
	ShopperLocale string `json:"shopperLocale,omitempty"`
	// Required for recurring payments.  Your reference to uniquely identify this shopper, for example user ID or account ID. Minimum length: 3 characters. > Your reference must not include personally identifiable information (PII), for example name or email address.
	ShopperReference string `json:"shopperReference,omitempty"`
	// Boolean value indicating whether the card payment method should be split into separate debit and credit options.
	SplitCardFundingSources bool `json:"splitCardFundingSources,omitempty"`
	// The ecommerce or point-of-sale store that is processing the payment. Used in [partner arrangement integrations](https://docs.adyen.com/platforms/platforms-for-partners#route-payments) for Adyen for Platforms.
	Store string `json:"store,omitempty"`
}
