/*
Legal Entity Management API

The Legal Entity Management API enables you to manage legal entities that contain information required for verification.  ## Authentication To connect to the Legal Entity Management API, you must use the basic authentication credentials of your web service user. If you don't have one, contact the [Adyen Support Team](https://www.adyen.help/hc/en-us/requests/new). Use the web service user credentials to authenticate your request, for example:  ``` curl -U \"ws12345@Scope.BalancePlatform_YourBalancePlatform\":\"YourWsPassword\" \\ -H \"Content-Type: application/json\" \\ ... ``` Note that when going live, you need to generate new web service user credentials to access the [live endpoints](https://docs.adyen.com/development-resources/live-endpoints).  ## Versioning The Legal Entity Management API supports versioning of its endpoints through a version suffix in the endpoint URL. This suffix has the following format: \"vXX\", where XX is the version number.  For example: ``` https://kyc-test.adyen.com/lem/v2/legalEntities ``` ## Going live When going live, your Adyen contact will provide your API credential for the live environment. You can then use the username and password to send requests to `https://kyc-live.adyen.com/lem/v2`.  

API version: 2
Contact: developer-experience@adyen.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package legalentity

import (
	"encoding/json"
)

// checks if the PhoneNumber type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PhoneNumber{}

// PhoneNumber struct for PhoneNumber
type PhoneNumber struct {
	// The full phone number, including the country code. For example, **+3112345678**.
	Number string `json:"number"`
	// The type of phone number.  Possible values: **mobile**, **landline**, **sip**, **fax.** 
	Type *string `json:"type,omitempty"`
}

// NewPhoneNumber instantiates a new PhoneNumber object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPhoneNumber(number string) *PhoneNumber {
	this := PhoneNumber{}
	this.Number = number
	return &this
}

// NewPhoneNumberWithDefaults instantiates a new PhoneNumber object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPhoneNumberWithDefaults() *PhoneNumber {
	this := PhoneNumber{}
	return &this
}

// GetNumber returns the Number field value
func (o *PhoneNumber) GetNumber() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Number
}

// GetNumberOk returns a tuple with the Number field value
// and a boolean to check if the value has been set.
func (o *PhoneNumber) GetNumberOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Number, true
}

// SetNumber sets field value
func (o *PhoneNumber) SetNumber(v string) {
	o.Number = v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *PhoneNumber) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PhoneNumber) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *PhoneNumber) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *PhoneNumber) SetType(v string) {
	o.Type = &v
}

func (o PhoneNumber) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PhoneNumber) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["number"] = o.Number
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	return toSerialize, nil
}

type NullablePhoneNumber struct {
	value *PhoneNumber
	isSet bool
}

func (v NullablePhoneNumber) Get() *PhoneNumber {
	return v.value
}

func (v *NullablePhoneNumber) Set(val *PhoneNumber) {
	v.value = val
	v.isSet = true
}

func (v NullablePhoneNumber) IsSet() bool {
	return v.isSet
}

func (v *NullablePhoneNumber) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePhoneNumber(val *PhoneNumber) *NullablePhoneNumber {
	return &NullablePhoneNumber{value: val, isSet: true}
}

func (v NullablePhoneNumber) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePhoneNumber) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


