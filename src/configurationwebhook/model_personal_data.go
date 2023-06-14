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

// checks if the PersonalData type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &PersonalData{}

// PersonalData struct for PersonalData
type PersonalData struct {
	// The date of birth of the person. The date should be in ISO-8601 format yyyy-mm-dd (e.g. 2000-01-31).
	DateOfBirth *string `json:"dateOfBirth,omitempty"`
	// An ID number of the person.
	IdNumber *string `json:"idNumber,omitempty"`
	// The nationality of the person represented by a two-character country code. >The permitted country codes are defined in ISO-3166-1 alpha-2 (e.g. 'NL').
	Nationality *string `json:"nationality,omitempty"`
}

// NewPersonalData instantiates a new PersonalData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPersonalData() *PersonalData {
	this := PersonalData{}
	return &this
}

// NewPersonalDataWithDefaults instantiates a new PersonalData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPersonalDataWithDefaults() *PersonalData {
	this := PersonalData{}
	return &this
}

// GetDateOfBirth returns the DateOfBirth field value if set, zero value otherwise.
func (o *PersonalData) GetDateOfBirth() string {
	if o == nil || common.IsNil(o.DateOfBirth) {
		var ret string
		return ret
	}
	return *o.DateOfBirth
}

// GetDateOfBirthOk returns a tuple with the DateOfBirth field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PersonalData) GetDateOfBirthOk() (*string, bool) {
	if o == nil || common.IsNil(o.DateOfBirth) {
		return nil, false
	}
	return o.DateOfBirth, true
}

// HasDateOfBirth returns a boolean if a field has been set.
func (o *PersonalData) HasDateOfBirth() bool {
	if o != nil && !common.IsNil(o.DateOfBirth) {
		return true
	}

	return false
}

// SetDateOfBirth gets a reference to the given string and assigns it to the DateOfBirth field.
func (o *PersonalData) SetDateOfBirth(v string) {
	o.DateOfBirth = &v
}

// GetIdNumber returns the IdNumber field value if set, zero value otherwise.
func (o *PersonalData) GetIdNumber() string {
	if o == nil || common.IsNil(o.IdNumber) {
		var ret string
		return ret
	}
	return *o.IdNumber
}

// GetIdNumberOk returns a tuple with the IdNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PersonalData) GetIdNumberOk() (*string, bool) {
	if o == nil || common.IsNil(o.IdNumber) {
		return nil, false
	}
	return o.IdNumber, true
}

// HasIdNumber returns a boolean if a field has been set.
func (o *PersonalData) HasIdNumber() bool {
	if o != nil && !common.IsNil(o.IdNumber) {
		return true
	}

	return false
}

// SetIdNumber gets a reference to the given string and assigns it to the IdNumber field.
func (o *PersonalData) SetIdNumber(v string) {
	o.IdNumber = &v
}

// GetNationality returns the Nationality field value if set, zero value otherwise.
func (o *PersonalData) GetNationality() string {
	if o == nil || common.IsNil(o.Nationality) {
		var ret string
		return ret
	}
	return *o.Nationality
}

// GetNationalityOk returns a tuple with the Nationality field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PersonalData) GetNationalityOk() (*string, bool) {
	if o == nil || common.IsNil(o.Nationality) {
		return nil, false
	}
	return o.Nationality, true
}

// HasNationality returns a boolean if a field has been set.
func (o *PersonalData) HasNationality() bool {
	if o != nil && !common.IsNil(o.Nationality) {
		return true
	}

	return false
}

// SetNationality gets a reference to the given string and assigns it to the Nationality field.
func (o *PersonalData) SetNationality(v string) {
	o.Nationality = &v
}

func (o PersonalData) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PersonalData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.DateOfBirth) {
		toSerialize["dateOfBirth"] = o.DateOfBirth
	}
	if !common.IsNil(o.IdNumber) {
		toSerialize["idNumber"] = o.IdNumber
	}
	if !common.IsNil(o.Nationality) {
		toSerialize["nationality"] = o.Nationality
	}
	return toSerialize, nil
}

type NullablePersonalData struct {
	value *PersonalData
	isSet bool
}

func (v NullablePersonalData) Get() *PersonalData {
	return v.value
}

func (v *NullablePersonalData) Set(val *PersonalData) {
	v.value = val
	v.isSet = true
}

func (v NullablePersonalData) IsSet() bool {
	return v.isSet
}

func (v *NullablePersonalData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePersonalData(val *PersonalData) *NullablePersonalData {
	return &NullablePersonalData{value: val, isSet: true}
}

func (v NullablePersonalData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePersonalData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
