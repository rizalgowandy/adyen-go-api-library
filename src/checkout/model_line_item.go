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

// LineItem struct for LineItem
type LineItem struct {
	// Item amount excluding the tax, in minor units.
	AmountExcludingTax int64 `json:"amountExcludingTax,omitempty"`
	// Item amount including the tax, in minor units.
	AmountIncludingTax int64 `json:"amountIncludingTax,omitempty"`
	// Brand of the item.
	Brand string `json:"brand,omitempty"`
	// Color of the item.
	Color string `json:"color,omitempty"`
	// Description of the line item.
	Description string `json:"description,omitempty"`
	// ID of the line item.
	Id string `json:"id,omitempty"`
	// Link to the picture of the purchased item.
	ImageUrl string `json:"imageUrl,omitempty"`
	// Item category, used by the RatePay payment method.
	ItemCategory string `json:"itemCategory,omitempty"`
	// Manufacturer of the item.
	Manufacturer string `json:"manufacturer,omitempty"`
	// Link to the purchased item.
	ProductUrl string `json:"productUrl,omitempty"`
	// Number of items.
	Quantity int64 `json:"quantity,omitempty"`
	// Email associated with the given product in the basket (usually in electronic gift cards).
	ReceiverEmail string `json:"receiverEmail,omitempty"`
	// Size of the item.
	Size string `json:"size,omitempty"`
	// Stock keeping unit.
	Sku string `json:"sku,omitempty"`
	// Tax amount, in minor units.
	TaxAmount int64 `json:"taxAmount,omitempty"`
	// Tax percentage, in minor units.
	TaxPercentage int64 `json:"taxPercentage,omitempty"`
	// Universal Product Code.
	Upc string `json:"upc,omitempty"`
}
