/*
Transfers API

API version: 3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package transfers

import (
	"encoding/json"

	"github.com/adyen/adyen-go-api-library/v7/src/common"
)

// checks if the NameLocation type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &NameLocation{}

// NameLocation struct for NameLocation
type NameLocation struct {
	// The city where the merchant is located.
	City *string `json:"city,omitempty"`
	// The country where the merchant is located in [three-letter country code](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-3) format.
	Country *string `json:"country,omitempty"`
	// The home country in [three-digit country code](https://en.wikipedia.org/wiki/ISO_3166-1_numeric) format, used for government-controlled merchants such as embassies.
	CountryOfOrigin *string `json:"countryOfOrigin,omitempty"`
	// The name of the merchant's shop or service.
	Name *string `json:"name,omitempty"`
	// The raw data.
	RawData *string `json:"rawData,omitempty"`
	// The state where the merchant is located.
	State *string `json:"state,omitempty"`
}

// NewNameLocation instantiates a new NameLocation object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNameLocation() *NameLocation {
	this := NameLocation{}
	return &this
}

// NewNameLocationWithDefaults instantiates a new NameLocation object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNameLocationWithDefaults() *NameLocation {
	this := NameLocation{}
	return &this
}

// GetCity returns the City field value if set, zero value otherwise.
func (o *NameLocation) GetCity() string {
	if o == nil || common.IsNil(o.City) {
		var ret string
		return ret
	}
	return *o.City
}

// GetCityOk returns a tuple with the City field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NameLocation) GetCityOk() (*string, bool) {
	if o == nil || common.IsNil(o.City) {
		return nil, false
	}
	return o.City, true
}

// HasCity returns a boolean if a field has been set.
func (o *NameLocation) HasCity() bool {
	if o != nil && !common.IsNil(o.City) {
		return true
	}

	return false
}

// SetCity gets a reference to the given string and assigns it to the City field.
func (o *NameLocation) SetCity(v string) {
	o.City = &v
}

// GetCountry returns the Country field value if set, zero value otherwise.
func (o *NameLocation) GetCountry() string {
	if o == nil || common.IsNil(o.Country) {
		var ret string
		return ret
	}
	return *o.Country
}

// GetCountryOk returns a tuple with the Country field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NameLocation) GetCountryOk() (*string, bool) {
	if o == nil || common.IsNil(o.Country) {
		return nil, false
	}
	return o.Country, true
}

// HasCountry returns a boolean if a field has been set.
func (o *NameLocation) HasCountry() bool {
	if o != nil && !common.IsNil(o.Country) {
		return true
	}

	return false
}

// SetCountry gets a reference to the given string and assigns it to the Country field.
func (o *NameLocation) SetCountry(v string) {
	o.Country = &v
}

// GetCountryOfOrigin returns the CountryOfOrigin field value if set, zero value otherwise.
func (o *NameLocation) GetCountryOfOrigin() string {
	if o == nil || common.IsNil(o.CountryOfOrigin) {
		var ret string
		return ret
	}
	return *o.CountryOfOrigin
}

// GetCountryOfOriginOk returns a tuple with the CountryOfOrigin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NameLocation) GetCountryOfOriginOk() (*string, bool) {
	if o == nil || common.IsNil(o.CountryOfOrigin) {
		return nil, false
	}
	return o.CountryOfOrigin, true
}

// HasCountryOfOrigin returns a boolean if a field has been set.
func (o *NameLocation) HasCountryOfOrigin() bool {
	if o != nil && !common.IsNil(o.CountryOfOrigin) {
		return true
	}

	return false
}

// SetCountryOfOrigin gets a reference to the given string and assigns it to the CountryOfOrigin field.
func (o *NameLocation) SetCountryOfOrigin(v string) {
	o.CountryOfOrigin = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *NameLocation) GetName() string {
	if o == nil || common.IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NameLocation) GetNameOk() (*string, bool) {
	if o == nil || common.IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *NameLocation) HasName() bool {
	if o != nil && !common.IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *NameLocation) SetName(v string) {
	o.Name = &v
}

// GetRawData returns the RawData field value if set, zero value otherwise.
func (o *NameLocation) GetRawData() string {
	if o == nil || common.IsNil(o.RawData) {
		var ret string
		return ret
	}
	return *o.RawData
}

// GetRawDataOk returns a tuple with the RawData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NameLocation) GetRawDataOk() (*string, bool) {
	if o == nil || common.IsNil(o.RawData) {
		return nil, false
	}
	return o.RawData, true
}

// HasRawData returns a boolean if a field has been set.
func (o *NameLocation) HasRawData() bool {
	if o != nil && !common.IsNil(o.RawData) {
		return true
	}

	return false
}

// SetRawData gets a reference to the given string and assigns it to the RawData field.
func (o *NameLocation) SetRawData(v string) {
	o.RawData = &v
}

// GetState returns the State field value if set, zero value otherwise.
func (o *NameLocation) GetState() string {
	if o == nil || common.IsNil(o.State) {
		var ret string
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NameLocation) GetStateOk() (*string, bool) {
	if o == nil || common.IsNil(o.State) {
		return nil, false
	}
	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *NameLocation) HasState() bool {
	if o != nil && !common.IsNil(o.State) {
		return true
	}

	return false
}

// SetState gets a reference to the given string and assigns it to the State field.
func (o *NameLocation) SetState(v string) {
	o.State = &v
}

func (o NameLocation) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NameLocation) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.City) {
		toSerialize["city"] = o.City
	}
	if !common.IsNil(o.Country) {
		toSerialize["country"] = o.Country
	}
	if !common.IsNil(o.CountryOfOrigin) {
		toSerialize["countryOfOrigin"] = o.CountryOfOrigin
	}
	if !common.IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !common.IsNil(o.RawData) {
		toSerialize["rawData"] = o.RawData
	}
	if !common.IsNil(o.State) {
		toSerialize["state"] = o.State
	}
	return toSerialize, nil
}

type NullableNameLocation struct {
	value *NameLocation
	isSet bool
}

func (v NullableNameLocation) Get() *NameLocation {
	return v.value
}

func (v *NullableNameLocation) Set(val *NameLocation) {
	v.value = val
	v.isSet = true
}

func (v NullableNameLocation) IsSet() bool {
	return v.isSet
}

func (v *NullableNameLocation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNameLocation(val *NameLocation) *NullableNameLocation {
	return &NullableNameLocation{value: val, isSet: true}
}

func (v NullableNameLocation) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNameLocation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
