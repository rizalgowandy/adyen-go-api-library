/*
 * Adyen Payout API
 *
 * A set of API endpoints that allow you to store payout details, confirm, or decline a payout.  For more information, refer to [Online payouts](https://docs.adyen.com/checkout/online-payouts).
 *
 * API version: 52
 * Contact: support@adyen.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package payout

import (
	"time"
)

// MerchantRiskIndicator struct for MerchantRiskIndicator
type MerchantRiskIndicator struct {
	// Whether the chosen delivery address is identical to the billing address.
	AddressMatch bool `json:"addressMatch,omitempty"`
	// Indicator regarding the delivery address. Allowed values: * `shipToBillingAddress` * `shipToVerifiedAddress` * `shipToNewAddress` * `shipToStore` * `digitalGoods` * `goodsNotShipped` * `other`
	DeliveryAddressIndicator string `json:"deliveryAddressIndicator,omitempty"`
	// The delivery email address (for digital goods).
	DeliveryEmail string `json:"deliveryEmail,omitempty"`
	// The estimated delivery time for the shopper to receive the goods. Allowed values: * `electronicDelivery` * `sameDayShipping` * `overnightShipping` * `twoOrMoreDaysShipping`
	DeliveryTimeframe string  `json:"deliveryTimeframe,omitempty"`
	GiftCardAmount    *Amount `json:"giftCardAmount,omitempty"`
	// Number of individual prepaid or gift cards used for this purchase.
	GiftCardCount int32 `json:"giftCardCount,omitempty"`
	// For pre-order purchases, the expected date this product will be available to the shopper.
	PreOrderDate *time.Time `json:"preOrderDate,omitempty"`
	// Indicator for whether this transaction is for pre-ordering a product.
	PreOrderPurchase bool `json:"preOrderPurchase,omitempty"`
	// Indicator for whether the shopper has already purchased the same items in the past.
	ReorderItems bool `json:"reorderItems,omitempty"`
}
