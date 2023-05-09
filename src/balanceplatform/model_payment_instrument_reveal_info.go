/*
Configuration API

API version: 2
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package balanceplatform

import (
	"encoding/json"

	"github.com/adyen/adyen-go-api-library/v6/src/common"
)

// checks if the PaymentInstrumentRevealInfo type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &PaymentInstrumentRevealInfo{}

// PaymentInstrumentRevealInfo struct for PaymentInstrumentRevealInfo
type PaymentInstrumentRevealInfo struct {
	// The CVC2 value of the card.
	Cvc        string `json:"cvc"`
	Expiration Expiry `json:"expiration"`
	// The primary account number (PAN) of the card.
	Pan string `json:"pan"`
}

// NewPaymentInstrumentRevealInfo instantiates a new PaymentInstrumentRevealInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPaymentInstrumentRevealInfo(cvc string, expiration Expiry, pan string) *PaymentInstrumentRevealInfo {
	this := PaymentInstrumentRevealInfo{}
	this.Cvc = cvc
	this.Expiration = expiration
	this.Pan = pan
	return &this
}

// NewPaymentInstrumentRevealInfoWithDefaults instantiates a new PaymentInstrumentRevealInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPaymentInstrumentRevealInfoWithDefaults() *PaymentInstrumentRevealInfo {
	this := PaymentInstrumentRevealInfo{}
	return &this
}

// GetCvc returns the Cvc field value
func (o *PaymentInstrumentRevealInfo) GetCvc() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Cvc
}

// GetCvcOk returns a tuple with the Cvc field value
// and a boolean to check if the value has been set.
func (o *PaymentInstrumentRevealInfo) GetCvcOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Cvc, true
}

// SetCvc sets field value
func (o *PaymentInstrumentRevealInfo) SetCvc(v string) {
	o.Cvc = v
}

// GetExpiration returns the Expiration field value
func (o *PaymentInstrumentRevealInfo) GetExpiration() Expiry {
	if o == nil {
		var ret Expiry
		return ret
	}

	return o.Expiration
}

// GetExpirationOk returns a tuple with the Expiration field value
// and a boolean to check if the value has been set.
func (o *PaymentInstrumentRevealInfo) GetExpirationOk() (*Expiry, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Expiration, true
}

// SetExpiration sets field value
func (o *PaymentInstrumentRevealInfo) SetExpiration(v Expiry) {
	o.Expiration = v
}

// GetPan returns the Pan field value
func (o *PaymentInstrumentRevealInfo) GetPan() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Pan
}

// GetPanOk returns a tuple with the Pan field value
// and a boolean to check if the value has been set.
func (o *PaymentInstrumentRevealInfo) GetPanOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Pan, true
}

// SetPan sets field value
func (o *PaymentInstrumentRevealInfo) SetPan(v string) {
	o.Pan = v
}

func (o PaymentInstrumentRevealInfo) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PaymentInstrumentRevealInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["cvc"] = o.Cvc
	toSerialize["expiration"] = o.Expiration
	toSerialize["pan"] = o.Pan
	return toSerialize, nil
}

type NullablePaymentInstrumentRevealInfo struct {
	value *PaymentInstrumentRevealInfo
	isSet bool
}

func (v NullablePaymentInstrumentRevealInfo) Get() *PaymentInstrumentRevealInfo {
	return v.value
}

func (v *NullablePaymentInstrumentRevealInfo) Set(val *PaymentInstrumentRevealInfo) {
	v.value = val
	v.isSet = true
}

func (v NullablePaymentInstrumentRevealInfo) IsSet() bool {
	return v.isSet
}

func (v *NullablePaymentInstrumentRevealInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaymentInstrumentRevealInfo(val *PaymentInstrumentRevealInfo) *NullablePaymentInstrumentRevealInfo {
	return &NullablePaymentInstrumentRevealInfo{value: val, isSet: true}
}

func (v NullablePaymentInstrumentRevealInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaymentInstrumentRevealInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
