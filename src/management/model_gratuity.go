/*
Management API

API version: 3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package management

import (
	"encoding/json"

	"github.com/adyen/adyen-go-api-library/v8/src/common"
)

// checks if the Gratuity type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &Gratuity{}

// Gratuity struct for Gratuity
type Gratuity struct {
	// Indicates whether one of the predefined tipping options is to let the shopper enter a custom tip. If **true**, only three of the other options defined in `predefinedTipEntries` are shown.
	AllowCustomAmount *bool `json:"allowCustomAmount,omitempty"`
	// The currency that the tipping settings apply to.
	Currency *string `json:"currency,omitempty"`
	// Tipping options the shopper can choose from if `usePredefinedTipEntries` is **true**. The maximum number of predefined options is four, or three plus the option to enter a custom tip. The options can be a mix of:  - A percentage of the transaction amount. Example: **5%** - A tip amount in [minor units](https://docs.adyen.com/development-resources/currency-codes). Example: **500** for a EUR 5 tip.
	PredefinedTipEntries []string `json:"predefinedTipEntries,omitempty"`
	// Indicates whether the terminal shows a prompt to enter a tip (**false**), or predefined tipping options to choose from (**true**).
	UsePredefinedTipEntries *bool `json:"usePredefinedTipEntries,omitempty"`
}

// NewGratuity instantiates a new Gratuity object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGratuity() *Gratuity {
	this := Gratuity{}
	return &this
}

// NewGratuityWithDefaults instantiates a new Gratuity object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGratuityWithDefaults() *Gratuity {
	this := Gratuity{}
	return &this
}

// GetAllowCustomAmount returns the AllowCustomAmount field value if set, zero value otherwise.
func (o *Gratuity) GetAllowCustomAmount() bool {
	if o == nil || common.IsNil(o.AllowCustomAmount) {
		var ret bool
		return ret
	}
	return *o.AllowCustomAmount
}

// GetAllowCustomAmountOk returns a tuple with the AllowCustomAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Gratuity) GetAllowCustomAmountOk() (*bool, bool) {
	if o == nil || common.IsNil(o.AllowCustomAmount) {
		return nil, false
	}
	return o.AllowCustomAmount, true
}

// HasAllowCustomAmount returns a boolean if a field has been set.
func (o *Gratuity) HasAllowCustomAmount() bool {
	if o != nil && !common.IsNil(o.AllowCustomAmount) {
		return true
	}

	return false
}

// SetAllowCustomAmount gets a reference to the given bool and assigns it to the AllowCustomAmount field.
func (o *Gratuity) SetAllowCustomAmount(v bool) {
	o.AllowCustomAmount = &v
}

// GetCurrency returns the Currency field value if set, zero value otherwise.
func (o *Gratuity) GetCurrency() string {
	if o == nil || common.IsNil(o.Currency) {
		var ret string
		return ret
	}
	return *o.Currency
}

// GetCurrencyOk returns a tuple with the Currency field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Gratuity) GetCurrencyOk() (*string, bool) {
	if o == nil || common.IsNil(o.Currency) {
		return nil, false
	}
	return o.Currency, true
}

// HasCurrency returns a boolean if a field has been set.
func (o *Gratuity) HasCurrency() bool {
	if o != nil && !common.IsNil(o.Currency) {
		return true
	}

	return false
}

// SetCurrency gets a reference to the given string and assigns it to the Currency field.
func (o *Gratuity) SetCurrency(v string) {
	o.Currency = &v
}

// GetPredefinedTipEntries returns the PredefinedTipEntries field value if set, zero value otherwise.
func (o *Gratuity) GetPredefinedTipEntries() []string {
	if o == nil || common.IsNil(o.PredefinedTipEntries) {
		var ret []string
		return ret
	}
	return o.PredefinedTipEntries
}

// GetPredefinedTipEntriesOk returns a tuple with the PredefinedTipEntries field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Gratuity) GetPredefinedTipEntriesOk() ([]string, bool) {
	if o == nil || common.IsNil(o.PredefinedTipEntries) {
		return nil, false
	}
	return o.PredefinedTipEntries, true
}

// HasPredefinedTipEntries returns a boolean if a field has been set.
func (o *Gratuity) HasPredefinedTipEntries() bool {
	if o != nil && !common.IsNil(o.PredefinedTipEntries) {
		return true
	}

	return false
}

// SetPredefinedTipEntries gets a reference to the given []string and assigns it to the PredefinedTipEntries field.
func (o *Gratuity) SetPredefinedTipEntries(v []string) {
	o.PredefinedTipEntries = v
}

// GetUsePredefinedTipEntries returns the UsePredefinedTipEntries field value if set, zero value otherwise.
func (o *Gratuity) GetUsePredefinedTipEntries() bool {
	if o == nil || common.IsNil(o.UsePredefinedTipEntries) {
		var ret bool
		return ret
	}
	return *o.UsePredefinedTipEntries
}

// GetUsePredefinedTipEntriesOk returns a tuple with the UsePredefinedTipEntries field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Gratuity) GetUsePredefinedTipEntriesOk() (*bool, bool) {
	if o == nil || common.IsNil(o.UsePredefinedTipEntries) {
		return nil, false
	}
	return o.UsePredefinedTipEntries, true
}

// HasUsePredefinedTipEntries returns a boolean if a field has been set.
func (o *Gratuity) HasUsePredefinedTipEntries() bool {
	if o != nil && !common.IsNil(o.UsePredefinedTipEntries) {
		return true
	}

	return false
}

// SetUsePredefinedTipEntries gets a reference to the given bool and assigns it to the UsePredefinedTipEntries field.
func (o *Gratuity) SetUsePredefinedTipEntries(v bool) {
	o.UsePredefinedTipEntries = &v
}

func (o Gratuity) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Gratuity) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.AllowCustomAmount) {
		toSerialize["allowCustomAmount"] = o.AllowCustomAmount
	}
	if !common.IsNil(o.Currency) {
		toSerialize["currency"] = o.Currency
	}
	if !common.IsNil(o.PredefinedTipEntries) {
		toSerialize["predefinedTipEntries"] = o.PredefinedTipEntries
	}
	if !common.IsNil(o.UsePredefinedTipEntries) {
		toSerialize["usePredefinedTipEntries"] = o.UsePredefinedTipEntries
	}
	return toSerialize, nil
}

type NullableGratuity struct {
	value *Gratuity
	isSet bool
}

func (v NullableGratuity) Get() *Gratuity {
	return v.value
}

func (v *NullableGratuity) Set(val *Gratuity) {
	v.value = val
	v.isSet = true
}

func (v NullableGratuity) IsSet() bool {
	return v.isSet
}

func (v *NullableGratuity) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGratuity(val *Gratuity) *NullableGratuity {
	return &NullableGratuity{value: val, isSet: true}
}

func (v NullableGratuity) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGratuity) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
