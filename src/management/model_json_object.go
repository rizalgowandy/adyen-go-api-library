/*
Management API

API version: 1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package management

import (
	"encoding/json"

	"github.com/adyen/adyen-go-api-library/v7/src/common"
)

// checks if the JSONObject type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &JSONObject{}

// JSONObject struct for JSONObject
type JSONObject struct {
	Paths    []JSONPath `json:"paths,omitempty"`
	RootPath *JSONPath  `json:"rootPath,omitempty"`
}

// NewJSONObject instantiates a new JSONObject object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewJSONObject() *JSONObject {
	this := JSONObject{}
	return &this
}

// NewJSONObjectWithDefaults instantiates a new JSONObject object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewJSONObjectWithDefaults() *JSONObject {
	this := JSONObject{}
	return &this
}

// GetPaths returns the Paths field value if set, zero value otherwise.
func (o *JSONObject) GetPaths() []JSONPath {
	if o == nil || common.IsNil(o.Paths) {
		var ret []JSONPath
		return ret
	}
	return o.Paths
}

// GetPathsOk returns a tuple with the Paths field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JSONObject) GetPathsOk() ([]JSONPath, bool) {
	if o == nil || common.IsNil(o.Paths) {
		return nil, false
	}
	return o.Paths, true
}

// HasPaths returns a boolean if a field has been set.
func (o *JSONObject) HasPaths() bool {
	if o != nil && !common.IsNil(o.Paths) {
		return true
	}

	return false
}

// SetPaths gets a reference to the given []JSONPath and assigns it to the Paths field.
func (o *JSONObject) SetPaths(v []JSONPath) {
	o.Paths = v
}

// GetRootPath returns the RootPath field value if set, zero value otherwise.
func (o *JSONObject) GetRootPath() JSONPath {
	if o == nil || common.IsNil(o.RootPath) {
		var ret JSONPath
		return ret
	}
	return *o.RootPath
}

// GetRootPathOk returns a tuple with the RootPath field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JSONObject) GetRootPathOk() (*JSONPath, bool) {
	if o == nil || common.IsNil(o.RootPath) {
		return nil, false
	}
	return o.RootPath, true
}

// HasRootPath returns a boolean if a field has been set.
func (o *JSONObject) HasRootPath() bool {
	if o != nil && !common.IsNil(o.RootPath) {
		return true
	}

	return false
}

// SetRootPath gets a reference to the given JSONPath and assigns it to the RootPath field.
func (o *JSONObject) SetRootPath(v JSONPath) {
	o.RootPath = &v
}

func (o JSONObject) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o JSONObject) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.Paths) {
		toSerialize["paths"] = o.Paths
	}
	if !common.IsNil(o.RootPath) {
		toSerialize["rootPath"] = o.RootPath
	}
	return toSerialize, nil
}

type NullableJSONObject struct {
	value *JSONObject
	isSet bool
}

func (v NullableJSONObject) Get() *JSONObject {
	return v.value
}

func (v *NullableJSONObject) Set(val *JSONObject) {
	v.value = val
	v.isSet = true
}

func (v NullableJSONObject) IsSet() bool {
	return v.isSet
}

func (v *NullableJSONObject) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableJSONObject(val *JSONObject) *NullableJSONObject {
	return &NullableJSONObject{value: val, isSet: true}
}

func (v NullableJSONObject) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableJSONObject) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
