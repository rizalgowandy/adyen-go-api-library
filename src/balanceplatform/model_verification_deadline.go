/*
Configuration API

API version: 2
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package balanceplatform

import (
	"encoding/json"
	"time"

	"github.com/adyen/adyen-go-api-library/v7/src/common"
)

// checks if the VerificationDeadline type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &VerificationDeadline{}

// VerificationDeadline struct for VerificationDeadline
type VerificationDeadline struct {
	// The names of the capabilities to be disallowed.
	Capabilities []string `json:"capabilities"`
	// The date that verification is due by before capabilities are disallowed.
	ExpiresAt time.Time `json:"expiresAt"`
}

// NewVerificationDeadline instantiates a new VerificationDeadline object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVerificationDeadline(capabilities []string, expiresAt time.Time) *VerificationDeadline {
	this := VerificationDeadline{}
	this.Capabilities = capabilities
	this.ExpiresAt = expiresAt
	return &this
}

// NewVerificationDeadlineWithDefaults instantiates a new VerificationDeadline object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVerificationDeadlineWithDefaults() *VerificationDeadline {
	this := VerificationDeadline{}
	return &this
}

// GetCapabilities returns the Capabilities field value
func (o *VerificationDeadline) GetCapabilities() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Capabilities
}

// GetCapabilitiesOk returns a tuple with the Capabilities field value
// and a boolean to check if the value has been set.
func (o *VerificationDeadline) GetCapabilitiesOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Capabilities, true
}

// SetCapabilities sets field value
func (o *VerificationDeadline) SetCapabilities(v []string) {
	o.Capabilities = v
}

// GetExpiresAt returns the ExpiresAt field value
func (o *VerificationDeadline) GetExpiresAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.ExpiresAt
}

// GetExpiresAtOk returns a tuple with the ExpiresAt field value
// and a boolean to check if the value has been set.
func (o *VerificationDeadline) GetExpiresAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ExpiresAt, true
}

// SetExpiresAt sets field value
func (o *VerificationDeadline) SetExpiresAt(v time.Time) {
	o.ExpiresAt = v
}

func (o VerificationDeadline) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o VerificationDeadline) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["capabilities"] = o.Capabilities
	toSerialize["expiresAt"] = o.ExpiresAt
	return toSerialize, nil
}

type NullableVerificationDeadline struct {
	value *VerificationDeadline
	isSet bool
}

func (v NullableVerificationDeadline) Get() *VerificationDeadline {
	return v.value
}

func (v *NullableVerificationDeadline) Set(val *VerificationDeadline) {
	v.value = val
	v.isSet = true
}

func (v NullableVerificationDeadline) IsSet() bool {
	return v.isSet
}

func (v *NullableVerificationDeadline) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVerificationDeadline(val *VerificationDeadline) *NullableVerificationDeadline {
	return &NullableVerificationDeadline{value: val, isSet: true}
}

func (v NullableVerificationDeadline) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVerificationDeadline) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
