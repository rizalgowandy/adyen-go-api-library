/*
 * Adyen Checkout API
 *
 * Adyen Checkout API provides a simple and flexible way to initiate and authorise online payments. You can use the same integration for payments made with cards (including 3D Secure), mobile wallets, and local payment methods (for example, iDEAL and Sofort).  This API reference provides information on available endpoints and how to interact with them. To learn more about the API, visit [online payments documentation](https://docs.adyen.com/online-payments).  ## Authentication Each request to Checkout API must be signed with an API key. For this, [get your API key](https://docs.adyen.com/development-resources/api-credentials#generate-api-key) from your Customer Area, and set this key to the `X-API-Key` header value, for example:  ``` curl -H \"Content-Type: application/json\" \\ -H \"X-API-Key: YOUR_API_KEY\" \\ ... ``` ## Versioning Checkout API supports [versioning](https://docs.adyen.com/development-resources/versioning) using a version suffix in the endpoint URL. This suffix has the following format: \"vXX\", where XX is the version number.  For example: ``` https://checkout-test.adyen.com/v70/payments ```  ## Going live  To access the live endpoints, you need an API key from your live Customer Area.  The live endpoint URLs contain a prefix which is unique to your company account, for example: ``` https://{PREFIX}-checkout-live.adyenpayments.com/checkout/v70/payments ```  Get your `{PREFIX}` from your live Customer Area under **Developers** > **API URLs** > **Prefix**.  When preparing to do live transactions with Checkout API, follow the [go-live checklist](https://docs.adyen.com/online-payments/go-live-checklist) to make sure you've got all the required configuration in place.  ## Release notes Have a look at the [release notes](https://docs.adyen.com/online-payments/release-notes?integration_type=api&version=70) to find out what changed in this version!
 *
 * API version: 70
 * Contact: developer-experience@adyen.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package checkout

// CardDetails struct for CardDetails
type CardDetails struct {
	// Secondary brand of the card. For example: **plastix**, **hmclub**.
	Brand string `json:"brand,omitempty"`
	// The checkout attempt identifier.
	CheckoutAttemptId    string `json:"checkoutAttemptId,omitempty"`
	CupsecureplusSmscode string `json:"cupsecureplus.smscode,omitempty"`
	// The card verification code. Only collect raw card data if you are [fully PCI compliant](https://docs.adyen.com/development-resources/pci-dss-compliance-guide).
	Cvc string `json:"cvc,omitempty"`
	// The encrypted card number.
	EncryptedCardNumber string `json:"encryptedCardNumber"`
	// The encrypted card expiry month.
	EncryptedExpiryMonth string `json:"encryptedExpiryMonth"`
	// The encrypted card expiry year.
	EncryptedExpiryYear string `json:"encryptedExpiryYear"`
	// The encrypted card verification code.
	EncryptedSecurityCode string `json:"encryptedSecurityCode,omitempty"`
	// The card expiry month. Only collect raw card data if you are [fully PCI compliant](https://docs.adyen.com/development-resources/pci-dss-compliance-guide).
	ExpiryMonth string `json:"expiryMonth,omitempty"`
	// The card expiry year. Only collect raw card data if you are [fully PCI compliant](https://docs.adyen.com/development-resources/pci-dss-compliance-guide).
	ExpiryYear string `json:"expiryYear,omitempty"`
	// The funding source that should be used when multiple sources are available. For Brazilian combo cards, by default the funding source is credit. To use debit, set this value to **debit**.
	FundingSource string `json:"fundingSource,omitempty"`
	// The name of the card holder.
	HolderName string `json:"holderName,omitempty"`
	// The network token reference. This is the [`networkTxReference`](https://docs.adyen.com/api-explorer/#/CheckoutService/latest/post/payments__resParam_additionalData-ResponseAdditionalDataCommon-networkTxReference) from the response to the first payment.
	NetworkPaymentReference string `json:"networkPaymentReference,omitempty"`
	// The card number. Only collect raw card data if you are [fully PCI compliant](https://docs.adyen.com/development-resources/pci-dss-compliance-guide).
	Number string `json:"number,omitempty"`
	// This is the `recurringDetailReference` returned in the response when you created the token.
	RecurringDetailReference string `json:"recurringDetailReference,omitempty"`
	// The `shopperNotificationReference` returned in the response when you requested to notify the shopper. Used only for recurring payments in India.
	ShopperNotificationReference string `json:"shopperNotificationReference,omitempty"`
	// This is the `recurringDetailReference` returned in the response when you created the token.
	StoredPaymentMethodId string `json:"storedPaymentMethodId,omitempty"`
	// Version of the 3D Secure 2 mobile SDK.
	ThreeDS2SdkVersion string `json:"threeDS2SdkVersion,omitempty"`
	// Default payment method details. Common for scheme payment methods, and for simple payment method details.
	Type string `json:"type,omitempty"`
}
