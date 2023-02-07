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

// checks if the OnboardingThemes type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OnboardingThemes{}

// OnboardingThemes struct for OnboardingThemes
type OnboardingThemes struct {
	// The next page. Only present if there is a next page.
	Next *string `json:"next,omitempty"`
	// The previous page. Only present if there is a previous page.
	Previous *string `json:"previous,omitempty"`
	// List of onboarding themes.
	Themes []OnboardingTheme `json:"themes"`
}

// NewOnboardingThemes instantiates a new OnboardingThemes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOnboardingThemes(themes []OnboardingTheme) *OnboardingThemes {
	this := OnboardingThemes{}
	this.Themes = themes
	return &this
}

// NewOnboardingThemesWithDefaults instantiates a new OnboardingThemes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOnboardingThemesWithDefaults() *OnboardingThemes {
	this := OnboardingThemes{}
	return &this
}

// GetNext returns the Next field value if set, zero value otherwise.
func (o *OnboardingThemes) GetNext() string {
	if o == nil || IsNil(o.Next) {
		var ret string
		return ret
	}
	return *o.Next
}

// GetNextOk returns a tuple with the Next field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OnboardingThemes) GetNextOk() (*string, bool) {
	if o == nil || IsNil(o.Next) {
		return nil, false
	}
	return o.Next, true
}

// HasNext returns a boolean if a field has been set.
func (o *OnboardingThemes) HasNext() bool {
	if o != nil && !IsNil(o.Next) {
		return true
	}

	return false
}

// SetNext gets a reference to the given string and assigns it to the Next field.
func (o *OnboardingThemes) SetNext(v string) {
	o.Next = &v
}

// GetPrevious returns the Previous field value if set, zero value otherwise.
func (o *OnboardingThemes) GetPrevious() string {
	if o == nil || IsNil(o.Previous) {
		var ret string
		return ret
	}
	return *o.Previous
}

// GetPreviousOk returns a tuple with the Previous field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OnboardingThemes) GetPreviousOk() (*string, bool) {
	if o == nil || IsNil(o.Previous) {
		return nil, false
	}
	return o.Previous, true
}

// HasPrevious returns a boolean if a field has been set.
func (o *OnboardingThemes) HasPrevious() bool {
	if o != nil && !IsNil(o.Previous) {
		return true
	}

	return false
}

// SetPrevious gets a reference to the given string and assigns it to the Previous field.
func (o *OnboardingThemes) SetPrevious(v string) {
	o.Previous = &v
}

// GetThemes returns the Themes field value
func (o *OnboardingThemes) GetThemes() []OnboardingTheme {
	if o == nil {
		var ret []OnboardingTheme
		return ret
	}

	return o.Themes
}

// GetThemesOk returns a tuple with the Themes field value
// and a boolean to check if the value has been set.
func (o *OnboardingThemes) GetThemesOk() ([]OnboardingTheme, bool) {
	if o == nil {
		return nil, false
	}
	return o.Themes, true
}

// SetThemes sets field value
func (o *OnboardingThemes) SetThemes(v []OnboardingTheme) {
	o.Themes = v
}

func (o OnboardingThemes) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OnboardingThemes) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Next) {
		toSerialize["next"] = o.Next
	}
	if !IsNil(o.Previous) {
		toSerialize["previous"] = o.Previous
	}
	toSerialize["themes"] = o.Themes
	return toSerialize, nil
}

type NullableOnboardingThemes struct {
	value *OnboardingThemes
	isSet bool
}

func (v NullableOnboardingThemes) Get() *OnboardingThemes {
	return v.value
}

func (v *NullableOnboardingThemes) Set(val *OnboardingThemes) {
	v.value = val
	v.isSet = true
}

func (v NullableOnboardingThemes) IsSet() bool {
	return v.isSet
}

func (v *NullableOnboardingThemes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOnboardingThemes(val *OnboardingThemes) *NullableOnboardingThemes {
	return &NullableOnboardingThemes{value: val, isSet: true}
}

func (v NullableOnboardingThemes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOnboardingThemes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


