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

// checks if the SupportingEntityCapability type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &SupportingEntityCapability{}

// SupportingEntityCapability struct for SupportingEntityCapability
type SupportingEntityCapability struct {
	// Indicates whether the supporting entity capability is allowed.  If a supporting entity is allowed but its parent legal entity is not, it means there are other supporting entities that failed validation.  **The allowed supporting entity can still be used**
	Allowed *bool `json:"allowed,omitempty"`
	// Supporting entity reference
	Id *string `json:"id,omitempty"`
	// Indicates whether the supporting entity capability is requested.
	Requested *bool `json:"requested,omitempty"`
	// The status of the verification checks for the supporting entity capability.  Possible values:  * **pending**: Adyen is running the verification.  * **invalid**: The verification failed. Check if the `errors` array contains more information.  * **valid**: The verification has been successfully completed.  * **rejected**: Adyen has verified the information, but found reasons to not allow the capability.
	VerificationStatus *string `json:"verificationStatus,omitempty"`
}

// NewSupportingEntityCapability instantiates a new SupportingEntityCapability object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSupportingEntityCapability() *SupportingEntityCapability {
	this := SupportingEntityCapability{}
	return &this
}

// NewSupportingEntityCapabilityWithDefaults instantiates a new SupportingEntityCapability object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSupportingEntityCapabilityWithDefaults() *SupportingEntityCapability {
	this := SupportingEntityCapability{}
	return &this
}

// GetAllowed returns the Allowed field value if set, zero value otherwise.
func (o *SupportingEntityCapability) GetAllowed() bool {
	if o == nil || common.IsNil(o.Allowed) {
		var ret bool
		return ret
	}
	return *o.Allowed
}

// GetAllowedOk returns a tuple with the Allowed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SupportingEntityCapability) GetAllowedOk() (*bool, bool) {
	if o == nil || common.IsNil(o.Allowed) {
		return nil, false
	}
	return o.Allowed, true
}

// HasAllowed returns a boolean if a field has been set.
func (o *SupportingEntityCapability) HasAllowed() bool {
	if o != nil && !common.IsNil(o.Allowed) {
		return true
	}

	return false
}

// SetAllowed gets a reference to the given bool and assigns it to the Allowed field.
func (o *SupportingEntityCapability) SetAllowed(v bool) {
	o.Allowed = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *SupportingEntityCapability) GetId() string {
	if o == nil || common.IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SupportingEntityCapability) GetIdOk() (*string, bool) {
	if o == nil || common.IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *SupportingEntityCapability) HasId() bool {
	if o != nil && !common.IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *SupportingEntityCapability) SetId(v string) {
	o.Id = &v
}

// GetRequested returns the Requested field value if set, zero value otherwise.
func (o *SupportingEntityCapability) GetRequested() bool {
	if o == nil || common.IsNil(o.Requested) {
		var ret bool
		return ret
	}
	return *o.Requested
}

// GetRequestedOk returns a tuple with the Requested field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SupportingEntityCapability) GetRequestedOk() (*bool, bool) {
	if o == nil || common.IsNil(o.Requested) {
		return nil, false
	}
	return o.Requested, true
}

// HasRequested returns a boolean if a field has been set.
func (o *SupportingEntityCapability) HasRequested() bool {
	if o != nil && !common.IsNil(o.Requested) {
		return true
	}

	return false
}

// SetRequested gets a reference to the given bool and assigns it to the Requested field.
func (o *SupportingEntityCapability) SetRequested(v bool) {
	o.Requested = &v
}

// GetVerificationStatus returns the VerificationStatus field value if set, zero value otherwise.
func (o *SupportingEntityCapability) GetVerificationStatus() string {
	if o == nil || common.IsNil(o.VerificationStatus) {
		var ret string
		return ret
	}
	return *o.VerificationStatus
}

// GetVerificationStatusOk returns a tuple with the VerificationStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SupportingEntityCapability) GetVerificationStatusOk() (*string, bool) {
	if o == nil || common.IsNil(o.VerificationStatus) {
		return nil, false
	}
	return o.VerificationStatus, true
}

// HasVerificationStatus returns a boolean if a field has been set.
func (o *SupportingEntityCapability) HasVerificationStatus() bool {
	if o != nil && !common.IsNil(o.VerificationStatus) {
		return true
	}

	return false
}

// SetVerificationStatus gets a reference to the given string and assigns it to the VerificationStatus field.
func (o *SupportingEntityCapability) SetVerificationStatus(v string) {
	o.VerificationStatus = &v
}

func (o SupportingEntityCapability) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SupportingEntityCapability) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.Allowed) {
		toSerialize["allowed"] = o.Allowed
	}
	if !common.IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !common.IsNil(o.Requested) {
		toSerialize["requested"] = o.Requested
	}
	if !common.IsNil(o.VerificationStatus) {
		toSerialize["verificationStatus"] = o.VerificationStatus
	}
	return toSerialize, nil
}

type NullableSupportingEntityCapability struct {
	value *SupportingEntityCapability
	isSet bool
}

func (v NullableSupportingEntityCapability) Get() *SupportingEntityCapability {
	return v.value
}

func (v *NullableSupportingEntityCapability) Set(val *SupportingEntityCapability) {
	v.value = val
	v.isSet = true
}

func (v NullableSupportingEntityCapability) IsSet() bool {
	return v.isSet
}

func (v *NullableSupportingEntityCapability) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSupportingEntityCapability(val *SupportingEntityCapability) *NullableSupportingEntityCapability {
	return &NullableSupportingEntityCapability{value: val, isSet: true}
}

func (v NullableSupportingEntityCapability) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSupportingEntityCapability) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
