/*
Adyen Recurring API

API version: 68
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package recurring

import (
	"encoding/json"

	"github.com/adyen/adyen-go-api-library/v8/src/common"
)

// checks if the CreatePermitResult type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &CreatePermitResult{}

// CreatePermitResult struct for CreatePermitResult
type CreatePermitResult struct {
	// List of new permits.
	PermitResultList []PermitResult `json:"permitResultList,omitempty"`
	// A unique reference associated with the request. This value is globally unique; quote it when communicating with us about this request.
	PspReference *string `json:"pspReference,omitempty"`
}

// NewCreatePermitResult instantiates a new CreatePermitResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreatePermitResult() *CreatePermitResult {
	this := CreatePermitResult{}
	return &this
}

// NewCreatePermitResultWithDefaults instantiates a new CreatePermitResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreatePermitResultWithDefaults() *CreatePermitResult {
	this := CreatePermitResult{}
	return &this
}

// GetPermitResultList returns the PermitResultList field value if set, zero value otherwise.
func (o *CreatePermitResult) GetPermitResultList() []PermitResult {
	if o == nil || common.IsNil(o.PermitResultList) {
		var ret []PermitResult
		return ret
	}
	return o.PermitResultList
}

// GetPermitResultListOk returns a tuple with the PermitResultList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreatePermitResult) GetPermitResultListOk() ([]PermitResult, bool) {
	if o == nil || common.IsNil(o.PermitResultList) {
		return nil, false
	}
	return o.PermitResultList, true
}

// HasPermitResultList returns a boolean if a field has been set.
func (o *CreatePermitResult) HasPermitResultList() bool {
	if o != nil && !common.IsNil(o.PermitResultList) {
		return true
	}

	return false
}

// SetPermitResultList gets a reference to the given []PermitResult and assigns it to the PermitResultList field.
func (o *CreatePermitResult) SetPermitResultList(v []PermitResult) {
	o.PermitResultList = v
}

// GetPspReference returns the PspReference field value if set, zero value otherwise.
func (o *CreatePermitResult) GetPspReference() string {
	if o == nil || common.IsNil(o.PspReference) {
		var ret string
		return ret
	}
	return *o.PspReference
}

// GetPspReferenceOk returns a tuple with the PspReference field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreatePermitResult) GetPspReferenceOk() (*string, bool) {
	if o == nil || common.IsNil(o.PspReference) {
		return nil, false
	}
	return o.PspReference, true
}

// HasPspReference returns a boolean if a field has been set.
func (o *CreatePermitResult) HasPspReference() bool {
	if o != nil && !common.IsNil(o.PspReference) {
		return true
	}

	return false
}

// SetPspReference gets a reference to the given string and assigns it to the PspReference field.
func (o *CreatePermitResult) SetPspReference(v string) {
	o.PspReference = &v
}

func (o CreatePermitResult) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreatePermitResult) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.PermitResultList) {
		toSerialize["permitResultList"] = o.PermitResultList
	}
	if !common.IsNil(o.PspReference) {
		toSerialize["pspReference"] = o.PspReference
	}
	return toSerialize, nil
}

type NullableCreatePermitResult struct {
	value *CreatePermitResult
	isSet bool
}

func (v NullableCreatePermitResult) Get() *CreatePermitResult {
	return v.value
}

func (v *NullableCreatePermitResult) Set(val *CreatePermitResult) {
	v.value = val
	v.isSet = true
}

func (v NullableCreatePermitResult) IsSet() bool {
	return v.isSet
}

func (v *NullableCreatePermitResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreatePermitResult(val *CreatePermitResult) *NullableCreatePermitResult {
	return &NullableCreatePermitResult{value: val, isSet: true}
}

func (v NullableCreatePermitResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreatePermitResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
