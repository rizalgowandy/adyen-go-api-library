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

// AdditionalDataOpenInvoice struct for AdditionalDataOpenInvoice
type AdditionalDataOpenInvoice struct {
	// Holds different merchant data points like product, purchase, customer, and so on. It takes data in a Base64 encoded string.  The `merchantData` parameter needs to be added to the `openinvoicedata` signature at the end.  Since the field is optional, if it's not included it does not impact computing the merchant signature.  Applies only to Klarna.  You can contact Klarna for the format and structure of the string.
	OpeninvoicedataMerchantData string `json:"openinvoicedata.merchantData,omitempty"`
	// The number of invoice lines included in `openinvoicedata`.  There needs to be at least one line, so `numberOfLines` needs to be at least 1.
	OpeninvoicedataNumberOfLines string `json:"openinvoicedata.numberOfLines,omitempty"`
	// First name of the recipient. If the delivery address and the billing address are different, specify the `recipientFirstName` and `recipientLastName` to share the delivery address with Klarna. Otherwise, only the billing address is shared with Klarna.
	OpeninvoicedataRecipientFirstName string `json:"openinvoicedata.recipientFirstName,omitempty"`
	// Last name of the recipient. If the delivery address and the billing address are different, specify the `recipientFirstName` and `recipientLastName` to share the delivery address with Klarna. Otherwise, only the billing address is shared with Klarna.
	OpeninvoicedataRecipientLastName string `json:"openinvoicedata.recipientLastName,omitempty"`
	// The three-character ISO currency code.
	OpeninvoicedataLineItemNrCurrencyCode string `json:"openinvoicedataLine[itemNr].currencyCode,omitempty"`
	// A text description of the product the invoice line refers to.
	OpeninvoicedataLineItemNrDescription string `json:"openinvoicedataLine[itemNr].description,omitempty"`
	// The price for one item in the invoice line, represented in minor units.  The due amount for the item, VAT excluded.
	OpeninvoicedataLineItemNrItemAmount string `json:"openinvoicedataLine[itemNr].itemAmount,omitempty"`
	// A unique id for this item. Required for RatePay if the description of each item is not unique.
	OpeninvoicedataLineItemNrItemId string `json:"openinvoicedataLine[itemNr].itemId,omitempty"`
	// The VAT due for one item in the invoice line, represented in minor units.
	OpeninvoicedataLineItemNrItemVatAmount string `json:"openinvoicedataLine[itemNr].itemVatAmount,omitempty"`
	// The VAT percentage for one item in the invoice line, represented in minor units.  For example, 19% VAT is specified as 1900.
	OpeninvoicedataLineItemNrItemVatPercentage string `json:"openinvoicedataLine[itemNr].itemVatPercentage,omitempty"`
	// The number of units purchased of a specific product.
	OpeninvoicedataLineItemNrNumberOfItems string `json:"openinvoicedataLine[itemNr].numberOfItems,omitempty"`
	// Name of the shipping company handling the the return shipment.
	OpeninvoicedataLineItemNrReturnShippingCompany string `json:"openinvoicedataLine[itemNr].returnShippingCompany,omitempty"`
	// The tracking number for the return of the shipment.
	OpeninvoicedataLineItemNrReturnTrackingNumber string `json:"openinvoicedataLine[itemNr].returnTrackingNumber,omitempty"`
	// URI where the customer can track the return of their shipment.
	OpeninvoicedataLineItemNrReturnTrackingUri string `json:"openinvoicedataLine[itemNr].returnTrackingUri,omitempty"`
	// Name of the shipping company handling the delivery.
	OpeninvoicedataLineItemNrShippingCompany string `json:"openinvoicedataLine[itemNr].shippingCompany,omitempty"`
	// Shipping method.
	OpeninvoicedataLineItemNrShippingMethod string `json:"openinvoicedataLine[itemNr].shippingMethod,omitempty"`
	// The tracking number for the shipment.
	OpeninvoicedataLineItemNrTrackingNumber string `json:"openinvoicedataLine[itemNr].trackingNumber,omitempty"`
	// URI where the customer can track their shipment.
	OpeninvoicedataLineItemNrTrackingUri string `json:"openinvoicedataLine[itemNr].trackingUri,omitempty"`
}
