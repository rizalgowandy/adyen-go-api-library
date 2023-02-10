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

import (
	"time"
)

// PaymentLinkResponse struct for PaymentLinkResponse
type PaymentLinkResponse struct {
	// List of payment methods to be presented to the shopper. To refer to payment methods, use their `paymentMethod.type` from [Payment methods overview](https://docs.adyen.com/payment-methods).  Example: `\"allowedPaymentMethods\":[\"ideal\",\"giropay\"]`
	AllowedPaymentMethods []string `json:"allowedPaymentMethods,omitempty"`
	Amount                Amount   `json:"amount"`
	BillingAddress        *Address `json:"billingAddress,omitempty"`
	// List of payment methods to be hidden from the shopper. To refer to payment methods, use their `paymentMethod.type` from [Payment methods overview](https://docs.adyen.com/payment-methods).  Example: `\"blockedPaymentMethods\":[\"ideal\",\"giropay\"]`
	BlockedPaymentMethods []string `json:"blockedPaymentMethods,omitempty"`
	// The delay between the authorisation and scheduled auto-capture, specified in hours.
	CaptureDelayHours int32 `json:"captureDelayHours,omitempty"`
	// The shopper's two-letter country code.
	CountryCode string `json:"countryCode,omitempty"`
	// The shopper's date of birth.  Format [ISO-8601](https://www.w3.org/TR/NOTE-datetime): YYYY-MM-DD
	DateOfBirth string `json:"dateOfBirth,omitempty"`
	// The date and time when the purchased goods should be delivered.  [ISO 8601](https://www.w3.org/TR/NOTE-datetime) format: YYYY-MM-DDThh:mm:ss+TZD, for example, **2020-12-18T10:15:30+01:00**.
	DeliverAt       *time.Time `json:"deliverAt,omitempty"`
	DeliveryAddress *Address   `json:"deliveryAddress,omitempty"`
	// A short description visible on the payment page. Maximum length: 280 characters.
	Description string `json:"description,omitempty"`
	// The date when the payment link expires.  [ISO 8601](https://www.w3.org/TR/NOTE-datetime) format: YYYY-MM-DDThh:mm:ss+TZD, for example, **2020-12-18T10:15:30+01:00**.  The maximum expiry date is 70 days after the payment link is created.  If not provided, the payment link expires 24 hours after it was created.
	ExpiresAt string `json:"expiresAt,omitempty"`
	// A unique identifier of the payment link.
	Id string `json:"id"`
	// A set of key-value pairs that specifies the installment options available per payment method. The key must be a payment method name in lowercase. For example, **card** to specify installment options for all cards, or **visa** or **mc**. The value must be an object containing the installment options.
	InstallmentOptions *map[string]InstallmentOption `json:"installmentOptions,omitempty"`
	// Price and product information about the purchased items, to be included on the invoice sent to the shopper. This parameter is required for open invoice (_buy now, pay later_) payment methods such Afterpay, Clearpay, Klarna, RatePay, and Zip.
	LineItems *[]LineItem `json:"lineItems,omitempty"`
	// Indicates if the payment must be [captured manually](https://docs.adyen.com/online-payments/capture).
	ManualCapture bool `json:"manualCapture,omitempty"`
	// The [merchant category code](https://en.wikipedia.org/wiki/Merchant_category_code) (MCC) is a four-digit number, which relates to a particular market segment. This code reflects the predominant activity that is conducted by the merchant.
	Mcc string `json:"mcc,omitempty"`
	// The merchant account identifier for which the payment link is created.
	MerchantAccount string `json:"merchantAccount"`
	// This reference allows linking multiple transactions to each other for reporting purposes (for example, order auth-rate). The reference should be unique per billing cycle.
	MerchantOrderReference string `json:"merchantOrderReference,omitempty"`
	// Metadata consists of entries, each of which includes a key and a value. Limitations: * Maximum 20 key-value pairs per request. Otherwise, error \"177\" occurs: \"Metadata size exceeds limit\" * Maximum 20 characters per key. Otherwise, error \"178\" occurs: \"Metadata key size exceeds limit\" * A key cannot have the name `checkout.linkId`. Any value that you provide with this key is going to be replaced by the real payment link ID.
	Metadata map[string]string `json:"metadata,omitempty"`
	// Defines a recurring payment type. Required when creating a token to store payment details. Possible values: * **Subscription** – A transaction for a fixed or variable amount, which follows a fixed schedule. * **CardOnFile** – With a card-on-file (CoF) transaction, card details are stored to enable one-click or omnichannel journeys, or simply to streamline the checkout process. Any subscription not following a fixed schedule is also considered a card-on-file transaction. * **UnscheduledCardOnFile** – An unscheduled card-on-file (UCoF) transaction is a transaction that occurs on a non-fixed schedule and/or has variable amounts. For example, automatic top-ups when a cardholder's balance drops below a certain amount.
	RecurringProcessingModel string `json:"recurringProcessingModel,omitempty"`
	// A reference that is used to uniquely identify the payment in future communications about the payment status.
	Reference string `json:"reference"`
	// List of fields that the shopper has to provide on the payment page before completing the payment. For more information, refer to [Provide shopper information](https://docs.adyen.com/unified-commerce/pay-by-link/payment-links/api#shopper-information).  Possible values: * **billingAddress** – The address where to send the invoice. * **deliveryAddress** – The address where the purchased goods should be delivered. * **shopperEmail** – The shopper's email address. * **shopperName** – The shopper's full name. * **telephoneNumber** – The shopper's phone number.
	RequiredShopperFields []string `json:"requiredShopperFields,omitempty"`
	// Website URL used for redirection after payment is completed. If provided, a **Continue** button will be shown on the payment page. If shoppers select the button, they are redirected to the specified URL.
	ReturnUrl string `json:"returnUrl,omitempty"`
	// Indicates whether the payment link can be reused for multiple payments. If not provided, this defaults to **false** which means the link can be used for one successful payment only.
	Reusable bool      `json:"reusable,omitempty"`
	RiskData *RiskData `json:"riskData,omitempty"`
	// The shopper's email address.
	ShopperEmail string `json:"shopperEmail,omitempty"`
	// The language to be used in the payment page, specified by a combination of a language and country code. For example, `en-US`.  For a list of shopper locales that Pay by Link supports, refer to [Language and localization](https://docs.adyen.com/online-payments/pay-by-link#language-and-localization).
	ShopperLocale string `json:"shopperLocale,omitempty"`
	ShopperName   *Name  `json:"shopperName,omitempty"`
	// Your reference to uniquely identify this shopper, for example user ID or account ID. Minimum length: 3 characters. > Your reference must not include personally identifiable information (PII), for example name or email address.
	ShopperReference string `json:"shopperReference,omitempty"`
	// The text to be shown on the shopper's bank statement.  We recommend sending a maximum of 22 characters, otherwise banks might truncate the string.  Allowed characters: **a-z**, **A-Z**, **0-9**, spaces, and special characters **. , ' _ - ? + * /_**.
	ShopperStatement string `json:"shopperStatement,omitempty"`
	// Set to **false** to hide the button that lets the shopper remove a stored payment method.
	ShowRemovePaymentMethodButton bool `json:"showRemovePaymentMethodButton,omitempty"`
	// The shopper's social security number.
	SocialSecurityNumber string `json:"socialSecurityNumber,omitempty"`
	// Boolean value indicating whether the card payment method should be split into separate debit and credit options.
	SplitCardFundingSources bool `json:"splitCardFundingSources,omitempty"`
	// An array of objects specifying how the payment should be split between accounts when using Adyen for Platforms. For details, refer to [Providing split information](https://docs.adyen.com/marketplaces-and-platforms/processing-payments#providing-split-information).
	Splits *[]Split `json:"splits,omitempty"`
	// Status of the payment link. Possible values: * **active**: The link can be used to make payments. * **expired**: The expiry date for the payment link has passed. Shoppers can no longer use the link to make payments. * **completed**: The shopper completed the payment. * **paymentPending**: The shopper is in the process of making the payment. Applies to payment methods with an asynchronous flow.
	Status string `json:"status"`
	// The physical store, for which this payment is processed.
	Store string `json:"store,omitempty"`
	// Indicates if the details of the payment method will be stored for the shopper. Possible values: * **disabled** – No details will be stored (default). * **askForConsent** – If the `shopperReference` is provided, the UI lets the shopper choose if they want their payment details to be stored. * **enabled** – If the `shopperReference` is provided, the details will be stored without asking the shopper for consent.
	StorePaymentMethodMode string `json:"storePaymentMethodMode,omitempty"`
	// The shopper's telephone number.
	TelephoneNumber string `json:"telephoneNumber,omitempty"`
	// A [theme](https://docs.adyen.com/unified-commerce/pay-by-link/payment-links/api#themes) to customize the appearance of the payment page. If not specified, the payment page is rendered according to the theme set as default in your Customer Area.
	ThemeId string `json:"themeId,omitempty"`
	// The date when the payment link status was updated.  [ISO 8601](https://www.w3.org/TR/NOTE-datetime) format: YYYY-MM-DDThh:mm:ss+TZD, for example, **2020-12-18T10:15:30+01:00**.
	UpdatedAt *time.Time `json:"updatedAt,omitempty"`
	// The URL at which the shopper can complete the payment.
	Url string `json:"url"`
}
