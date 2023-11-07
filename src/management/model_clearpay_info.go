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

// checks if the ClearpayInfo type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &ClearpayInfo{}

// ClearpayInfo struct for ClearpayInfo
type ClearpayInfo struct {
	// Support Url
	SupportUrl string `json:"supportUrl"`
}

// NewClearpayInfo instantiates a new ClearpayInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewClearpayInfo(supportUrl string) *ClearpayInfo {
	this := ClearpayInfo{}
	this.SupportUrl = supportUrl
	return &this
}

// NewClearpayInfoWithDefaults instantiates a new ClearpayInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewClearpayInfoWithDefaults() *ClearpayInfo {
	this := ClearpayInfo{}
	return &this
}

// GetSupportUrl returns the SupportUrl field value
func (o *ClearpayInfo) GetSupportUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SupportUrl
}

// GetSupportUrlOk returns a tuple with the SupportUrl field value
// and a boolean to check if the value has been set.
func (o *ClearpayInfo) GetSupportUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SupportUrl, true
}

// SetSupportUrl sets field value
func (o *ClearpayInfo) SetSupportUrl(v string) {
	o.SupportUrl = v
}

func (o ClearpayInfo) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ClearpayInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["supportUrl"] = o.SupportUrl
	return toSerialize, nil
}

type NullableClearpayInfo struct {
	value *ClearpayInfo
	isSet bool
}

func (v NullableClearpayInfo) Get() *ClearpayInfo {
	return v.value
}

func (v *NullableClearpayInfo) Set(val *ClearpayInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableClearpayInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableClearpayInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableClearpayInfo(val *ClearpayInfo) *NullableClearpayInfo {
	return &NullableClearpayInfo{value: val, isSet: true}
}

func (v NullableClearpayInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableClearpayInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
