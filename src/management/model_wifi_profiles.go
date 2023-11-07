/*
Management API

API version: 3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package management

import (
	"encoding/json"

	"github.com/adyen/adyen-go-api-library/v8/src/common"
)

// checks if the WifiProfiles type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &WifiProfiles{}

// WifiProfiles struct for WifiProfiles
type WifiProfiles struct {
	// List of remote Wi-Fi profiles.
	Profiles []Profile `json:"profiles,omitempty"`
	Settings *Settings `json:"settings,omitempty"`
}

// NewWifiProfiles instantiates a new WifiProfiles object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWifiProfiles() *WifiProfiles {
	this := WifiProfiles{}
	return &this
}

// NewWifiProfilesWithDefaults instantiates a new WifiProfiles object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWifiProfilesWithDefaults() *WifiProfiles {
	this := WifiProfiles{}
	return &this
}

// GetProfiles returns the Profiles field value if set, zero value otherwise.
func (o *WifiProfiles) GetProfiles() []Profile {
	if o == nil || common.IsNil(o.Profiles) {
		var ret []Profile
		return ret
	}
	return o.Profiles
}

// GetProfilesOk returns a tuple with the Profiles field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WifiProfiles) GetProfilesOk() ([]Profile, bool) {
	if o == nil || common.IsNil(o.Profiles) {
		return nil, false
	}
	return o.Profiles, true
}

// HasProfiles returns a boolean if a field has been set.
func (o *WifiProfiles) HasProfiles() bool {
	if o != nil && !common.IsNil(o.Profiles) {
		return true
	}

	return false
}

// SetProfiles gets a reference to the given []Profile and assigns it to the Profiles field.
func (o *WifiProfiles) SetProfiles(v []Profile) {
	o.Profiles = v
}

// GetSettings returns the Settings field value if set, zero value otherwise.
func (o *WifiProfiles) GetSettings() Settings {
	if o == nil || common.IsNil(o.Settings) {
		var ret Settings
		return ret
	}
	return *o.Settings
}

// GetSettingsOk returns a tuple with the Settings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WifiProfiles) GetSettingsOk() (*Settings, bool) {
	if o == nil || common.IsNil(o.Settings) {
		return nil, false
	}
	return o.Settings, true
}

// HasSettings returns a boolean if a field has been set.
func (o *WifiProfiles) HasSettings() bool {
	if o != nil && !common.IsNil(o.Settings) {
		return true
	}

	return false
}

// SetSettings gets a reference to the given Settings and assigns it to the Settings field.
func (o *WifiProfiles) SetSettings(v Settings) {
	o.Settings = &v
}

func (o WifiProfiles) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o WifiProfiles) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.Profiles) {
		toSerialize["profiles"] = o.Profiles
	}
	if !common.IsNil(o.Settings) {
		toSerialize["settings"] = o.Settings
	}
	return toSerialize, nil
}

type NullableWifiProfiles struct {
	value *WifiProfiles
	isSet bool
}

func (v NullableWifiProfiles) Get() *WifiProfiles {
	return v.value
}

func (v *NullableWifiProfiles) Set(val *WifiProfiles) {
	v.value = val
	v.isSet = true
}

func (v NullableWifiProfiles) IsSet() bool {
	return v.isSet
}

func (v *NullableWifiProfiles) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWifiProfiles(val *WifiProfiles) *NullableWifiProfiles {
	return &NullableWifiProfiles{value: val, isSet: true}
}

func (v NullableWifiProfiles) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWifiProfiles) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
