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

// checks if the PaymentMethodIssuer type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &PaymentMethodIssuer{}

// PaymentMethodIssuer struct for PaymentMethodIssuer
type PaymentMethodIssuer struct {
	// A boolean value indicating whether this issuer is unavailable. Can be `true` whenever the issuer is offline.
	Disabled *bool `json:"disabled,omitempty"`
	// The unique identifier of this issuer, to submit in requests to /payments.
	Id string `json:"id"`
	// A localized name of the issuer.
	Name string `json:"name"`
}

// NewPaymentMethodIssuer instantiates a new PaymentMethodIssuer object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPaymentMethodIssuer(id string, name string) *PaymentMethodIssuer {
	this := PaymentMethodIssuer{}
	var disabled bool = false
	this.Disabled = &disabled
	this.Id = id
	this.Name = name
	return &this
}

// NewPaymentMethodIssuerWithDefaults instantiates a new PaymentMethodIssuer object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPaymentMethodIssuerWithDefaults() *PaymentMethodIssuer {
	this := PaymentMethodIssuer{}
	var disabled bool = false
	this.Disabled = &disabled
	return &this
}

// GetDisabled returns the Disabled field value if set, zero value otherwise.
func (o *PaymentMethodIssuer) GetDisabled() bool {
	if o == nil || common.IsNil(o.Disabled) {
		var ret bool
		return ret
	}
	return *o.Disabled
}

// GetDisabledOk returns a tuple with the Disabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentMethodIssuer) GetDisabledOk() (*bool, bool) {
	if o == nil || common.IsNil(o.Disabled) {
		return nil, false
	}
	return o.Disabled, true
}

// HasDisabled returns a boolean if a field has been set.
func (o *PaymentMethodIssuer) HasDisabled() bool {
	if o != nil && !common.IsNil(o.Disabled) {
		return true
	}

	return false
}

// SetDisabled gets a reference to the given bool and assigns it to the Disabled field.
func (o *PaymentMethodIssuer) SetDisabled(v bool) {
	o.Disabled = &v
}

// GetId returns the Id field value
func (o *PaymentMethodIssuer) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *PaymentMethodIssuer) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *PaymentMethodIssuer) SetId(v string) {
	o.Id = v
}

// GetName returns the Name field value
func (o *PaymentMethodIssuer) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *PaymentMethodIssuer) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *PaymentMethodIssuer) SetName(v string) {
	o.Name = v
}

func (o PaymentMethodIssuer) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PaymentMethodIssuer) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.Disabled) {
		toSerialize["disabled"] = o.Disabled
	}
	toSerialize["id"] = o.Id
	toSerialize["name"] = o.Name
	return toSerialize, nil
}

type NullablePaymentMethodIssuer struct {
	value *PaymentMethodIssuer
	isSet bool
}

func (v NullablePaymentMethodIssuer) Get() *PaymentMethodIssuer {
	return v.value
}

func (v *NullablePaymentMethodIssuer) Set(val *PaymentMethodIssuer) {
	v.value = val
	v.isSet = true
}

func (v NullablePaymentMethodIssuer) IsSet() bool {
	return v.isSet
}

func (v *NullablePaymentMethodIssuer) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaymentMethodIssuer(val *PaymentMethodIssuer) *NullablePaymentMethodIssuer {
	return &NullablePaymentMethodIssuer{value: val, isSet: true}
}

func (v NullablePaymentMethodIssuer) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaymentMethodIssuer) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
