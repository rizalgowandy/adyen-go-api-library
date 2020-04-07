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

// SplitAmount struct for SplitAmount
type SplitAmount struct {
	// The three-character [ISO currency code](https://docs.adyen.com/development-resources/currency-codes).  If this value is not provided, the currency in which the payment is made will be used.
	Currency string `json:"currency,omitempty"`
	// The payable amount that can be charged for the transaction.  The transaction amount needs to be represented in minor units according to the [following table](https://docs.adyen.com/development-resources/currency-codes).
	Value int64 `json:"value"`
}
