/*
Configuration webhooks

API version: 1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package configurationwebhook

import (
	"encoding/json"

	"github.com/adyen/adyen-go-api-library/v8/src/common"
)

// checks if the PaymentInstrumentNotificationData type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &PaymentInstrumentNotificationData{}

// PaymentInstrumentNotificationData struct for PaymentInstrumentNotificationData
type PaymentInstrumentNotificationData struct {
	// The unique identifier of the balance platform.
	BalancePlatform   *string            `json:"balancePlatform,omitempty"`
	PaymentInstrument *PaymentInstrument `json:"paymentInstrument,omitempty"`
}

// NewPaymentInstrumentNotificationData instantiates a new PaymentInstrumentNotificationData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPaymentInstrumentNotificationData() *PaymentInstrumentNotificationData {
	this := PaymentInstrumentNotificationData{}
	return &this
}

// NewPaymentInstrumentNotificationDataWithDefaults instantiates a new PaymentInstrumentNotificationData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPaymentInstrumentNotificationDataWithDefaults() *PaymentInstrumentNotificationData {
	this := PaymentInstrumentNotificationData{}
	return &this
}

// GetBalancePlatform returns the BalancePlatform field value if set, zero value otherwise.
func (o *PaymentInstrumentNotificationData) GetBalancePlatform() string {
	if o == nil || common.IsNil(o.BalancePlatform) {
		var ret string
		return ret
	}
	return *o.BalancePlatform
}

// GetBalancePlatformOk returns a tuple with the BalancePlatform field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentInstrumentNotificationData) GetBalancePlatformOk() (*string, bool) {
	if o == nil || common.IsNil(o.BalancePlatform) {
		return nil, false
	}
	return o.BalancePlatform, true
}

// HasBalancePlatform returns a boolean if a field has been set.
func (o *PaymentInstrumentNotificationData) HasBalancePlatform() bool {
	if o != nil && !common.IsNil(o.BalancePlatform) {
		return true
	}

	return false
}

// SetBalancePlatform gets a reference to the given string and assigns it to the BalancePlatform field.
func (o *PaymentInstrumentNotificationData) SetBalancePlatform(v string) {
	o.BalancePlatform = &v
}

// GetPaymentInstrument returns the PaymentInstrument field value if set, zero value otherwise.
func (o *PaymentInstrumentNotificationData) GetPaymentInstrument() PaymentInstrument {
	if o == nil || common.IsNil(o.PaymentInstrument) {
		var ret PaymentInstrument
		return ret
	}
	return *o.PaymentInstrument
}

// GetPaymentInstrumentOk returns a tuple with the PaymentInstrument field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentInstrumentNotificationData) GetPaymentInstrumentOk() (*PaymentInstrument, bool) {
	if o == nil || common.IsNil(o.PaymentInstrument) {
		return nil, false
	}
	return o.PaymentInstrument, true
}

// HasPaymentInstrument returns a boolean if a field has been set.
func (o *PaymentInstrumentNotificationData) HasPaymentInstrument() bool {
	if o != nil && !common.IsNil(o.PaymentInstrument) {
		return true
	}

	return false
}

// SetPaymentInstrument gets a reference to the given PaymentInstrument and assigns it to the PaymentInstrument field.
func (o *PaymentInstrumentNotificationData) SetPaymentInstrument(v PaymentInstrument) {
	o.PaymentInstrument = &v
}

func (o PaymentInstrumentNotificationData) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PaymentInstrumentNotificationData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.BalancePlatform) {
		toSerialize["balancePlatform"] = o.BalancePlatform
	}
	if !common.IsNil(o.PaymentInstrument) {
		toSerialize["paymentInstrument"] = o.PaymentInstrument
	}
	return toSerialize, nil
}

type NullablePaymentInstrumentNotificationData struct {
	value *PaymentInstrumentNotificationData
	isSet bool
}

func (v NullablePaymentInstrumentNotificationData) Get() *PaymentInstrumentNotificationData {
	return v.value
}

func (v *NullablePaymentInstrumentNotificationData) Set(val *PaymentInstrumentNotificationData) {
	v.value = val
	v.isSet = true
}

func (v NullablePaymentInstrumentNotificationData) IsSet() bool {
	return v.isSet
}

func (v *NullablePaymentInstrumentNotificationData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaymentInstrumentNotificationData(val *PaymentInstrumentNotificationData) *NullablePaymentInstrumentNotificationData {
	return &NullablePaymentInstrumentNotificationData{value: val, isSet: true}
}

func (v NullablePaymentInstrumentNotificationData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaymentInstrumentNotificationData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
