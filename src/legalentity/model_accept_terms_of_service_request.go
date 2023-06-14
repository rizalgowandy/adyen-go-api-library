/*
Legal Entity Management API

API version: 3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package legalentity

import (
	"encoding/json"

	"github.com/adyen/adyen-go-api-library/v7/src/common"
)

// checks if the AcceptTermsOfServiceRequest type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &AcceptTermsOfServiceRequest{}

// AcceptTermsOfServiceRequest struct for AcceptTermsOfServiceRequest
type AcceptTermsOfServiceRequest struct {
	// The unique identifier of the user accepting the Terms of Service.
	AcceptedBy *string `json:"acceptedBy,omitempty"`
	// The IP address of the user accepting the Terms of Service.
	IpAddress *string `json:"ipAddress,omitempty"`
}

// NewAcceptTermsOfServiceRequest instantiates a new AcceptTermsOfServiceRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAcceptTermsOfServiceRequest() *AcceptTermsOfServiceRequest {
	this := AcceptTermsOfServiceRequest{}
	return &this
}

// NewAcceptTermsOfServiceRequestWithDefaults instantiates a new AcceptTermsOfServiceRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAcceptTermsOfServiceRequestWithDefaults() *AcceptTermsOfServiceRequest {
	this := AcceptTermsOfServiceRequest{}
	return &this
}

// GetAcceptedBy returns the AcceptedBy field value if set, zero value otherwise.
func (o *AcceptTermsOfServiceRequest) GetAcceptedBy() string {
	if o == nil || common.IsNil(o.AcceptedBy) {
		var ret string
		return ret
	}
	return *o.AcceptedBy
}

// GetAcceptedByOk returns a tuple with the AcceptedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AcceptTermsOfServiceRequest) GetAcceptedByOk() (*string, bool) {
	if o == nil || common.IsNil(o.AcceptedBy) {
		return nil, false
	}
	return o.AcceptedBy, true
}

// HasAcceptedBy returns a boolean if a field has been set.
func (o *AcceptTermsOfServiceRequest) HasAcceptedBy() bool {
	if o != nil && !common.IsNil(o.AcceptedBy) {
		return true
	}

	return false
}

// SetAcceptedBy gets a reference to the given string and assigns it to the AcceptedBy field.
func (o *AcceptTermsOfServiceRequest) SetAcceptedBy(v string) {
	o.AcceptedBy = &v
}

// GetIpAddress returns the IpAddress field value if set, zero value otherwise.
func (o *AcceptTermsOfServiceRequest) GetIpAddress() string {
	if o == nil || common.IsNil(o.IpAddress) {
		var ret string
		return ret
	}
	return *o.IpAddress
}

// GetIpAddressOk returns a tuple with the IpAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AcceptTermsOfServiceRequest) GetIpAddressOk() (*string, bool) {
	if o == nil || common.IsNil(o.IpAddress) {
		return nil, false
	}
	return o.IpAddress, true
}

// HasIpAddress returns a boolean if a field has been set.
func (o *AcceptTermsOfServiceRequest) HasIpAddress() bool {
	if o != nil && !common.IsNil(o.IpAddress) {
		return true
	}

	return false
}

// SetIpAddress gets a reference to the given string and assigns it to the IpAddress field.
func (o *AcceptTermsOfServiceRequest) SetIpAddress(v string) {
	o.IpAddress = &v
}

func (o AcceptTermsOfServiceRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AcceptTermsOfServiceRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.AcceptedBy) {
		toSerialize["acceptedBy"] = o.AcceptedBy
	}
	if !common.IsNil(o.IpAddress) {
		toSerialize["ipAddress"] = o.IpAddress
	}
	return toSerialize, nil
}

type NullableAcceptTermsOfServiceRequest struct {
	value *AcceptTermsOfServiceRequest
	isSet bool
}

func (v NullableAcceptTermsOfServiceRequest) Get() *AcceptTermsOfServiceRequest {
	return v.value
}

func (v *NullableAcceptTermsOfServiceRequest) Set(val *AcceptTermsOfServiceRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableAcceptTermsOfServiceRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableAcceptTermsOfServiceRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAcceptTermsOfServiceRequest(val *AcceptTermsOfServiceRequest) *NullableAcceptTermsOfServiceRequest {
	return &NullableAcceptTermsOfServiceRequest{value: val, isSet: true}
}

func (v NullableAcceptTermsOfServiceRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAcceptTermsOfServiceRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
