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

// checks if the ResponsePaymentMethod type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &ResponsePaymentMethod{}

// ResponsePaymentMethod struct for ResponsePaymentMethod
type ResponsePaymentMethod struct {
	// The card brand that the shopper used to pay. Only returned if `paymentMethod.type` is **scheme**.
	Brand *string `json:"brand,omitempty"`
	// The `paymentMethod.type` value used in the request.
	Type *string `json:"type,omitempty"`
}

// NewResponsePaymentMethod instantiates a new ResponsePaymentMethod object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewResponsePaymentMethod() *ResponsePaymentMethod {
	this := ResponsePaymentMethod{}
	return &this
}

// NewResponsePaymentMethodWithDefaults instantiates a new ResponsePaymentMethod object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewResponsePaymentMethodWithDefaults() *ResponsePaymentMethod {
	this := ResponsePaymentMethod{}
	return &this
}

// GetBrand returns the Brand field value if set, zero value otherwise.
func (o *ResponsePaymentMethod) GetBrand() string {
	if o == nil || common.IsNil(o.Brand) {
		var ret string
		return ret
	}
	return *o.Brand
}

// GetBrandOk returns a tuple with the Brand field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponsePaymentMethod) GetBrandOk() (*string, bool) {
	if o == nil || common.IsNil(o.Brand) {
		return nil, false
	}
	return o.Brand, true
}

// HasBrand returns a boolean if a field has been set.
func (o *ResponsePaymentMethod) HasBrand() bool {
	if o != nil && !common.IsNil(o.Brand) {
		return true
	}

	return false
}

// SetBrand gets a reference to the given string and assigns it to the Brand field.
func (o *ResponsePaymentMethod) SetBrand(v string) {
	o.Brand = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *ResponsePaymentMethod) GetType() string {
	if o == nil || common.IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponsePaymentMethod) GetTypeOk() (*string, bool) {
	if o == nil || common.IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *ResponsePaymentMethod) HasType() bool {
	if o != nil && !common.IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *ResponsePaymentMethod) SetType(v string) {
	o.Type = &v
}

func (o ResponsePaymentMethod) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ResponsePaymentMethod) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.Brand) {
		toSerialize["brand"] = o.Brand
	}
	if !common.IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	return toSerialize, nil
}

type NullableResponsePaymentMethod struct {
	value *ResponsePaymentMethod
	isSet bool
}

func (v NullableResponsePaymentMethod) Get() *ResponsePaymentMethod {
	return v.value
}

func (v *NullableResponsePaymentMethod) Set(val *ResponsePaymentMethod) {
	v.value = val
	v.isSet = true
}

func (v NullableResponsePaymentMethod) IsSet() bool {
	return v.isSet
}

func (v *NullableResponsePaymentMethod) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResponsePaymentMethod(val *ResponsePaymentMethod) *NullableResponsePaymentMethod {
	return &NullableResponsePaymentMethod{value: val, isSet: true}
}

func (v NullableResponsePaymentMethod) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableResponsePaymentMethod) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
