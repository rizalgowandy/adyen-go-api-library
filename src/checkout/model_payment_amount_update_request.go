/*
Adyen Checkout API

API version: 71
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package checkout

import (
	"encoding/json"

	"github.com/adyen/adyen-go-api-library/v8/src/common"
)

// checks if the PaymentAmountUpdateRequest type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &PaymentAmountUpdateRequest{}

// PaymentAmountUpdateRequest struct for PaymentAmountUpdateRequest
type PaymentAmountUpdateRequest struct {
	Amount          Amount           `json:"amount"`
	ApplicationInfo *ApplicationInfo `json:"applicationInfo,omitempty"`
	// The reason for the amount update. Possible values:  * **delayedCharge**  * **noShow**  * **installment**
	IndustryUsage *string `json:"industryUsage,omitempty"`
	// Price and product information of the refunded items, required for [partial refunds](https://docs.adyen.com/online-payments/refund#refund-a-payment). > This field is required for partial refunds with 3x 4x Oney, Affirm, Afterpay, Atome, Clearpay, Klarna, Ratepay, Walley, and Zip.
	LineItems []LineItem `json:"lineItems,omitempty"`
	// The merchant account that is used to process the payment.
	MerchantAccount string `json:"merchantAccount"`
	// Your reference for the amount update request. Maximum length: 80 characters.
	Reference *string `json:"reference,omitempty"`
	// An array of objects specifying how the amount should be split between accounts when using Adyen for Platforms. For details, refer to [Providing split information](https://docs.adyen.com/marketplaces-and-platforms/processing-payments#providing-split-information).
	Splits []Split `json:"splits,omitempty"`
}

// NewPaymentAmountUpdateRequest instantiates a new PaymentAmountUpdateRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPaymentAmountUpdateRequest(amount Amount, merchantAccount string) *PaymentAmountUpdateRequest {
	this := PaymentAmountUpdateRequest{}
	this.Amount = amount
	this.MerchantAccount = merchantAccount
	return &this
}

// NewPaymentAmountUpdateRequestWithDefaults instantiates a new PaymentAmountUpdateRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPaymentAmountUpdateRequestWithDefaults() *PaymentAmountUpdateRequest {
	this := PaymentAmountUpdateRequest{}
	return &this
}

// GetAmount returns the Amount field value
func (o *PaymentAmountUpdateRequest) GetAmount() Amount {
	if o == nil {
		var ret Amount
		return ret
	}

	return o.Amount
}

// GetAmountOk returns a tuple with the Amount field value
// and a boolean to check if the value has been set.
func (o *PaymentAmountUpdateRequest) GetAmountOk() (*Amount, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Amount, true
}

// SetAmount sets field value
func (o *PaymentAmountUpdateRequest) SetAmount(v Amount) {
	o.Amount = v
}

// GetApplicationInfo returns the ApplicationInfo field value if set, zero value otherwise.
func (o *PaymentAmountUpdateRequest) GetApplicationInfo() ApplicationInfo {
	if o == nil || common.IsNil(o.ApplicationInfo) {
		var ret ApplicationInfo
		return ret
	}
	return *o.ApplicationInfo
}

// GetApplicationInfoOk returns a tuple with the ApplicationInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentAmountUpdateRequest) GetApplicationInfoOk() (*ApplicationInfo, bool) {
	if o == nil || common.IsNil(o.ApplicationInfo) {
		return nil, false
	}
	return o.ApplicationInfo, true
}

// HasApplicationInfo returns a boolean if a field has been set.
func (o *PaymentAmountUpdateRequest) HasApplicationInfo() bool {
	if o != nil && !common.IsNil(o.ApplicationInfo) {
		return true
	}

	return false
}

// SetApplicationInfo gets a reference to the given ApplicationInfo and assigns it to the ApplicationInfo field.
func (o *PaymentAmountUpdateRequest) SetApplicationInfo(v ApplicationInfo) {
	o.ApplicationInfo = &v
}

// GetIndustryUsage returns the IndustryUsage field value if set, zero value otherwise.
func (o *PaymentAmountUpdateRequest) GetIndustryUsage() string {
	if o == nil || common.IsNil(o.IndustryUsage) {
		var ret string
		return ret
	}
	return *o.IndustryUsage
}

// GetIndustryUsageOk returns a tuple with the IndustryUsage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentAmountUpdateRequest) GetIndustryUsageOk() (*string, bool) {
	if o == nil || common.IsNil(o.IndustryUsage) {
		return nil, false
	}
	return o.IndustryUsage, true
}

// HasIndustryUsage returns a boolean if a field has been set.
func (o *PaymentAmountUpdateRequest) HasIndustryUsage() bool {
	if o != nil && !common.IsNil(o.IndustryUsage) {
		return true
	}

	return false
}

// SetIndustryUsage gets a reference to the given string and assigns it to the IndustryUsage field.
func (o *PaymentAmountUpdateRequest) SetIndustryUsage(v string) {
	o.IndustryUsage = &v
}

// GetLineItems returns the LineItems field value if set, zero value otherwise.
func (o *PaymentAmountUpdateRequest) GetLineItems() []LineItem {
	if o == nil || common.IsNil(o.LineItems) {
		var ret []LineItem
		return ret
	}
	return o.LineItems
}

// GetLineItemsOk returns a tuple with the LineItems field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentAmountUpdateRequest) GetLineItemsOk() ([]LineItem, bool) {
	if o == nil || common.IsNil(o.LineItems) {
		return nil, false
	}
	return o.LineItems, true
}

// HasLineItems returns a boolean if a field has been set.
func (o *PaymentAmountUpdateRequest) HasLineItems() bool {
	if o != nil && !common.IsNil(o.LineItems) {
		return true
	}

	return false
}

// SetLineItems gets a reference to the given []LineItem and assigns it to the LineItems field.
func (o *PaymentAmountUpdateRequest) SetLineItems(v []LineItem) {
	o.LineItems = v
}

// GetMerchantAccount returns the MerchantAccount field value
func (o *PaymentAmountUpdateRequest) GetMerchantAccount() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MerchantAccount
}

// GetMerchantAccountOk returns a tuple with the MerchantAccount field value
// and a boolean to check if the value has been set.
func (o *PaymentAmountUpdateRequest) GetMerchantAccountOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MerchantAccount, true
}

// SetMerchantAccount sets field value
func (o *PaymentAmountUpdateRequest) SetMerchantAccount(v string) {
	o.MerchantAccount = v
}

// GetReference returns the Reference field value if set, zero value otherwise.
func (o *PaymentAmountUpdateRequest) GetReference() string {
	if o == nil || common.IsNil(o.Reference) {
		var ret string
		return ret
	}
	return *o.Reference
}

// GetReferenceOk returns a tuple with the Reference field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentAmountUpdateRequest) GetReferenceOk() (*string, bool) {
	if o == nil || common.IsNil(o.Reference) {
		return nil, false
	}
	return o.Reference, true
}

// HasReference returns a boolean if a field has been set.
func (o *PaymentAmountUpdateRequest) HasReference() bool {
	if o != nil && !common.IsNil(o.Reference) {
		return true
	}

	return false
}

// SetReference gets a reference to the given string and assigns it to the Reference field.
func (o *PaymentAmountUpdateRequest) SetReference(v string) {
	o.Reference = &v
}

// GetSplits returns the Splits field value if set, zero value otherwise.
func (o *PaymentAmountUpdateRequest) GetSplits() []Split {
	if o == nil || common.IsNil(o.Splits) {
		var ret []Split
		return ret
	}
	return o.Splits
}

// GetSplitsOk returns a tuple with the Splits field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentAmountUpdateRequest) GetSplitsOk() ([]Split, bool) {
	if o == nil || common.IsNil(o.Splits) {
		return nil, false
	}
	return o.Splits, true
}

// HasSplits returns a boolean if a field has been set.
func (o *PaymentAmountUpdateRequest) HasSplits() bool {
	if o != nil && !common.IsNil(o.Splits) {
		return true
	}

	return false
}

// SetSplits gets a reference to the given []Split and assigns it to the Splits field.
func (o *PaymentAmountUpdateRequest) SetSplits(v []Split) {
	o.Splits = v
}

func (o PaymentAmountUpdateRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PaymentAmountUpdateRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["amount"] = o.Amount
	if !common.IsNil(o.ApplicationInfo) {
		toSerialize["applicationInfo"] = o.ApplicationInfo
	}
	if !common.IsNil(o.IndustryUsage) {
		toSerialize["industryUsage"] = o.IndustryUsage
	}
	if !common.IsNil(o.LineItems) {
		toSerialize["lineItems"] = o.LineItems
	}
	toSerialize["merchantAccount"] = o.MerchantAccount
	if !common.IsNil(o.Reference) {
		toSerialize["reference"] = o.Reference
	}
	if !common.IsNil(o.Splits) {
		toSerialize["splits"] = o.Splits
	}
	return toSerialize, nil
}

type NullablePaymentAmountUpdateRequest struct {
	value *PaymentAmountUpdateRequest
	isSet bool
}

func (v NullablePaymentAmountUpdateRequest) Get() *PaymentAmountUpdateRequest {
	return v.value
}

func (v *NullablePaymentAmountUpdateRequest) Set(val *PaymentAmountUpdateRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePaymentAmountUpdateRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePaymentAmountUpdateRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaymentAmountUpdateRequest(val *PaymentAmountUpdateRequest) *NullablePaymentAmountUpdateRequest {
	return &NullablePaymentAmountUpdateRequest{value: val, isSet: true}
}

func (v NullablePaymentAmountUpdateRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaymentAmountUpdateRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

func (o *PaymentAmountUpdateRequest) isValidIndustryUsage() bool {
	var allowedEnumValues = []string{"delayedCharge", "installment", "noShow"}
	for _, allowed := range allowedEnumValues {
		if o.GetIndustryUsage() == allowed {
			return true
		}
	}
	return false
}
