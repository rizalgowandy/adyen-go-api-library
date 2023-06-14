/*
Management API

API version: 1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package management

import (
	"encoding/json"

	"github.com/adyen/adyen-go-api-library/v7/src/common"
)

// checks if the TerminalModelsResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &TerminalModelsResponse{}

// TerminalModelsResponse struct for TerminalModelsResponse
type TerminalModelsResponse struct {
	// The terminal models that the API credential has access to.
	Data []IdName `json:"data,omitempty"`
}

// NewTerminalModelsResponse instantiates a new TerminalModelsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTerminalModelsResponse() *TerminalModelsResponse {
	this := TerminalModelsResponse{}
	return &this
}

// NewTerminalModelsResponseWithDefaults instantiates a new TerminalModelsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTerminalModelsResponseWithDefaults() *TerminalModelsResponse {
	this := TerminalModelsResponse{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *TerminalModelsResponse) GetData() []IdName {
	if o == nil || common.IsNil(o.Data) {
		var ret []IdName
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TerminalModelsResponse) GetDataOk() ([]IdName, bool) {
	if o == nil || common.IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *TerminalModelsResponse) HasData() bool {
	if o != nil && !common.IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given []IdName and assigns it to the Data field.
func (o *TerminalModelsResponse) SetData(v []IdName) {
	o.Data = v
}

func (o TerminalModelsResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TerminalModelsResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	return toSerialize, nil
}

type NullableTerminalModelsResponse struct {
	value *TerminalModelsResponse
	isSet bool
}

func (v NullableTerminalModelsResponse) Get() *TerminalModelsResponse {
	return v.value
}

func (v *NullableTerminalModelsResponse) Set(val *TerminalModelsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableTerminalModelsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableTerminalModelsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTerminalModelsResponse(val *TerminalModelsResponse) *NullableTerminalModelsResponse {
	return &NullableTerminalModelsResponse{value: val, isSet: true}
}

func (v NullableTerminalModelsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTerminalModelsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
