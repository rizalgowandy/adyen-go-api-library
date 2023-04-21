/*
Adyen Payout API

API version: 68
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package payout

import (
	"encoding/json"

	"github.com/adyen/adyen-go-api-library/v6/src/common"
)

// checks if the ModifyResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &ModifyResponse{}

// ModifyResponse struct for ModifyResponse
type ModifyResponse struct {
	// This field contains additional data, which may be returned in a particular response.
	AdditionalData *map[string]string `json:"additionalData,omitempty"`
	// Adyen's 16-character string reference associated with the transaction. This value is globally unique; quote it when communicating with us about this response.
	PspReference string `json:"pspReference"`
	// The response: * In case of success, it is either `payout-confirm-received` or `payout-decline-received`. * In case of an error, an informational message is returned.
	Response string `json:"response"`
}

// NewModifyResponse instantiates a new ModifyResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewModifyResponse(pspReference string, response string) *ModifyResponse {
	this := ModifyResponse{}
	this.PspReference = pspReference
	this.Response = response
	return &this
}

// NewModifyResponseWithDefaults instantiates a new ModifyResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewModifyResponseWithDefaults() *ModifyResponse {
	this := ModifyResponse{}
	return &this
}

// GetAdditionalData returns the AdditionalData field value if set, zero value otherwise.
func (o *ModifyResponse) GetAdditionalData() map[string]string {
	if o == nil || common.IsNil(o.AdditionalData) {
		var ret map[string]string
		return ret
	}
	return *o.AdditionalData
}

// GetAdditionalDataOk returns a tuple with the AdditionalData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModifyResponse) GetAdditionalDataOk() (*map[string]string, bool) {
	if o == nil || common.IsNil(o.AdditionalData) {
		return nil, false
	}
	return o.AdditionalData, true
}

// HasAdditionalData returns a boolean if a field has been set.
func (o *ModifyResponse) HasAdditionalData() bool {
	if o != nil && !common.IsNil(o.AdditionalData) {
		return true
	}

	return false
}

// SetAdditionalData gets a reference to the given map[string]string and assigns it to the AdditionalData field.
func (o *ModifyResponse) SetAdditionalData(v map[string]string) {
	o.AdditionalData = &v
}

// GetPspReference returns the PspReference field value
func (o *ModifyResponse) GetPspReference() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PspReference
}

// GetPspReferenceOk returns a tuple with the PspReference field value
// and a boolean to check if the value has been set.
func (o *ModifyResponse) GetPspReferenceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PspReference, true
}

// SetPspReference sets field value
func (o *ModifyResponse) SetPspReference(v string) {
	o.PspReference = v
}

// GetResponse returns the Response field value
func (o *ModifyResponse) GetResponse() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Response
}

// GetResponseOk returns a tuple with the Response field value
// and a boolean to check if the value has been set.
func (o *ModifyResponse) GetResponseOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Response, true
}

// SetResponse sets field value
func (o *ModifyResponse) SetResponse(v string) {
	o.Response = v
}

func (o ModifyResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ModifyResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.AdditionalData) {
		toSerialize["additionalData"] = o.AdditionalData
	}
	toSerialize["pspReference"] = o.PspReference
	toSerialize["response"] = o.Response
	return toSerialize, nil
}

type NullableModifyResponse struct {
	value *ModifyResponse
	isSet bool
}

func (v NullableModifyResponse) Get() *ModifyResponse {
	return v.value
}

func (v *NullableModifyResponse) Set(val *ModifyResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableModifyResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableModifyResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableModifyResponse(val *ModifyResponse) *NullableModifyResponse {
	return &NullableModifyResponse{value: val, isSet: true}
}

func (v NullableModifyResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableModifyResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
