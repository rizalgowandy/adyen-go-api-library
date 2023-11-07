/*
Configuration API

API version: 2
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package balanceplatform

import (
	"encoding/json"

	"github.com/adyen/adyen-go-api-library/v8/src/common"
)

// checks if the MccsRestriction type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &MccsRestriction{}

// MccsRestriction struct for MccsRestriction
type MccsRestriction struct {
	// Defines how the condition must be evaluated.
	Operation string `json:"operation"`
	// List of merchant category codes (MCCs).
	Value []string `json:"value,omitempty"`
}

// NewMccsRestriction instantiates a new MccsRestriction object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMccsRestriction(operation string) *MccsRestriction {
	this := MccsRestriction{}
	this.Operation = operation
	return &this
}

// NewMccsRestrictionWithDefaults instantiates a new MccsRestriction object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMccsRestrictionWithDefaults() *MccsRestriction {
	this := MccsRestriction{}
	return &this
}

// GetOperation returns the Operation field value
func (o *MccsRestriction) GetOperation() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Operation
}

// GetOperationOk returns a tuple with the Operation field value
// and a boolean to check if the value has been set.
func (o *MccsRestriction) GetOperationOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Operation, true
}

// SetOperation sets field value
func (o *MccsRestriction) SetOperation(v string) {
	o.Operation = v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *MccsRestriction) GetValue() []string {
	if o == nil || common.IsNil(o.Value) {
		var ret []string
		return ret
	}
	return o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MccsRestriction) GetValueOk() ([]string, bool) {
	if o == nil || common.IsNil(o.Value) {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *MccsRestriction) HasValue() bool {
	if o != nil && !common.IsNil(o.Value) {
		return true
	}

	return false
}

// SetValue gets a reference to the given []string and assigns it to the Value field.
func (o *MccsRestriction) SetValue(v []string) {
	o.Value = v
}

func (o MccsRestriction) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MccsRestriction) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["operation"] = o.Operation
	if !common.IsNil(o.Value) {
		toSerialize["value"] = o.Value
	}
	return toSerialize, nil
}

type NullableMccsRestriction struct {
	value *MccsRestriction
	isSet bool
}

func (v NullableMccsRestriction) Get() *MccsRestriction {
	return v.value
}

func (v *NullableMccsRestriction) Set(val *MccsRestriction) {
	v.value = val
	v.isSet = true
}

func (v NullableMccsRestriction) IsSet() bool {
	return v.isSet
}

func (v *NullableMccsRestriction) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMccsRestriction(val *MccsRestriction) *NullableMccsRestriction {
	return &NullableMccsRestriction{value: val, isSet: true}
}

func (v NullableMccsRestriction) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMccsRestriction) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
