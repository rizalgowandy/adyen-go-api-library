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
// UpiCollectDetails struct for UpiCollectDetails
type UpiCollectDetails struct {
	// The sequence number for the debit. For example, send **2** if this is the second debit for the subscription. The sequence number is included in the notification sent to the shopper.
	BillingSequenceNumber string `json:"billingSequenceNumber"`
	// This is the `recurringDetailReference` returned in the response when you created the token.
	RecurringDetailReference string `json:"recurringDetailReference,omitempty"`
	// The `shopperNotificationReference` returned in the response when you requested to notify the shopper. Used for recurring payment only.
	ShopperNotificationReference string `json:"shopperNotificationReference,omitempty"`
	// This is the `recurringDetailReference` returned in the response when you created the token.
	StoredPaymentMethodId string `json:"storedPaymentMethodId,omitempty"`
	// **upi_collect**
	Type string `json:"type"`
	// The virtual payment address for UPI.
	VirtualPaymentAddress string `json:"virtualPaymentAddress,omitempty"`
}
