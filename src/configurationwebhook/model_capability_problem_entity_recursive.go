/*
Configuration webhooks

API version: 1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package configurationwebhook

import (
	"encoding/json"

	"github.com/adyen/adyen-go-api-library/v7/src/common"
)

// checks if the CapabilityProblemEntityRecursive type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &CapabilityProblemEntityRecursive{}

// CapabilityProblemEntityRecursive struct for CapabilityProblemEntityRecursive
type CapabilityProblemEntityRecursive struct {
	// List of document IDs to which the verification errors related to the capabilities correspond to.
	Documents []string `json:"documents,omitempty"`
	// The ID of the entity.
	Id *string `json:"id,omitempty"`
	// Type of entity.   Possible values: **LegalEntity**, **BankAccount**, **Document**.
	Type *string `json:"type,omitempty"`
}

// NewCapabilityProblemEntityRecursive instantiates a new CapabilityProblemEntityRecursive object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCapabilityProblemEntityRecursive() *CapabilityProblemEntityRecursive {
	this := CapabilityProblemEntityRecursive{}
	return &this
}

// NewCapabilityProblemEntityRecursiveWithDefaults instantiates a new CapabilityProblemEntityRecursive object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCapabilityProblemEntityRecursiveWithDefaults() *CapabilityProblemEntityRecursive {
	this := CapabilityProblemEntityRecursive{}
	return &this
}

// GetDocuments returns the Documents field value if set, zero value otherwise.
func (o *CapabilityProblemEntityRecursive) GetDocuments() []string {
	if o == nil || common.IsNil(o.Documents) {
		var ret []string
		return ret
	}
	return o.Documents
}

// GetDocumentsOk returns a tuple with the Documents field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CapabilityProblemEntityRecursive) GetDocumentsOk() ([]string, bool) {
	if o == nil || common.IsNil(o.Documents) {
		return nil, false
	}
	return o.Documents, true
}

// HasDocuments returns a boolean if a field has been set.
func (o *CapabilityProblemEntityRecursive) HasDocuments() bool {
	if o != nil && !common.IsNil(o.Documents) {
		return true
	}

	return false
}

// SetDocuments gets a reference to the given []string and assigns it to the Documents field.
func (o *CapabilityProblemEntityRecursive) SetDocuments(v []string) {
	o.Documents = v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *CapabilityProblemEntityRecursive) GetId() string {
	if o == nil || common.IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CapabilityProblemEntityRecursive) GetIdOk() (*string, bool) {
	if o == nil || common.IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *CapabilityProblemEntityRecursive) HasId() bool {
	if o != nil && !common.IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *CapabilityProblemEntityRecursive) SetId(v string) {
	o.Id = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *CapabilityProblemEntityRecursive) GetType() string {
	if o == nil || common.IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CapabilityProblemEntityRecursive) GetTypeOk() (*string, bool) {
	if o == nil || common.IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *CapabilityProblemEntityRecursive) HasType() bool {
	if o != nil && !common.IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *CapabilityProblemEntityRecursive) SetType(v string) {
	o.Type = &v
}

func (o CapabilityProblemEntityRecursive) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CapabilityProblemEntityRecursive) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.Documents) {
		toSerialize["documents"] = o.Documents
	}
	if !common.IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !common.IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	return toSerialize, nil
}

type NullableCapabilityProblemEntityRecursive struct {
	value *CapabilityProblemEntityRecursive
	isSet bool
}

func (v NullableCapabilityProblemEntityRecursive) Get() *CapabilityProblemEntityRecursive {
	return v.value
}

func (v *NullableCapabilityProblemEntityRecursive) Set(val *CapabilityProblemEntityRecursive) {
	v.value = val
	v.isSet = true
}

func (v NullableCapabilityProblemEntityRecursive) IsSet() bool {
	return v.isSet
}

func (v *NullableCapabilityProblemEntityRecursive) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCapabilityProblemEntityRecursive(val *CapabilityProblemEntityRecursive) *NullableCapabilityProblemEntityRecursive {
	return &NullableCapabilityProblemEntityRecursive{value: val, isSet: true}
}

func (v NullableCapabilityProblemEntityRecursive) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCapabilityProblemEntityRecursive) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

func (o *CapabilityProblemEntityRecursive) isValidType() bool {
	var allowedEnumValues = []string{"BankAccount", "Document", "LegalEntity"}
	for _, allowed := range allowedEnumValues {
		if o.GetType() == allowed {
			return true
		}
	}
	return false
}
