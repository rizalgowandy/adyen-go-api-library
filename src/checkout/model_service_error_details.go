/*
Adyen Checkout API

API version: 71
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package checkout

import (
	"encoding/json"

	"github.com/adyen/adyen-go-api-library/v8/src/common"
)

// checks if the ServiceErrorDetails type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &ServiceErrorDetails{}

// ServiceErrorDetails struct for ServiceErrorDetails
type ServiceErrorDetails struct {
	ErrorCode    *string `json:"errorCode,omitempty"`
	ErrorType    *string `json:"errorType,omitempty"`
	Message      *string `json:"message,omitempty"`
	PspReference *string `json:"pspReference,omitempty"`
}

// NewServiceErrorDetails instantiates a new ServiceErrorDetails object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewServiceErrorDetails() *ServiceErrorDetails {
	this := ServiceErrorDetails{}
	return &this
}

// NewServiceErrorDetailsWithDefaults instantiates a new ServiceErrorDetails object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewServiceErrorDetailsWithDefaults() *ServiceErrorDetails {
	this := ServiceErrorDetails{}
	return &this
}

// GetErrorCode returns the ErrorCode field value if set, zero value otherwise.
func (o *ServiceErrorDetails) GetErrorCode() string {
	if o == nil || common.IsNil(o.ErrorCode) {
		var ret string
		return ret
	}
	return *o.ErrorCode
}

// GetErrorCodeOk returns a tuple with the ErrorCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceErrorDetails) GetErrorCodeOk() (*string, bool) {
	if o == nil || common.IsNil(o.ErrorCode) {
		return nil, false
	}
	return o.ErrorCode, true
}

// HasErrorCode returns a boolean if a field has been set.
func (o *ServiceErrorDetails) HasErrorCode() bool {
	if o != nil && !common.IsNil(o.ErrorCode) {
		return true
	}

	return false
}

// SetErrorCode gets a reference to the given string and assigns it to the ErrorCode field.
func (o *ServiceErrorDetails) SetErrorCode(v string) {
	o.ErrorCode = &v
}

// GetErrorType returns the ErrorType field value if set, zero value otherwise.
func (o *ServiceErrorDetails) GetErrorType() string {
	if o == nil || common.IsNil(o.ErrorType) {
		var ret string
		return ret
	}
	return *o.ErrorType
}

// GetErrorTypeOk returns a tuple with the ErrorType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceErrorDetails) GetErrorTypeOk() (*string, bool) {
	if o == nil || common.IsNil(o.ErrorType) {
		return nil, false
	}
	return o.ErrorType, true
}

// HasErrorType returns a boolean if a field has been set.
func (o *ServiceErrorDetails) HasErrorType() bool {
	if o != nil && !common.IsNil(o.ErrorType) {
		return true
	}

	return false
}

// SetErrorType gets a reference to the given string and assigns it to the ErrorType field.
func (o *ServiceErrorDetails) SetErrorType(v string) {
	o.ErrorType = &v
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *ServiceErrorDetails) GetMessage() string {
	if o == nil || common.IsNil(o.Message) {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceErrorDetails) GetMessageOk() (*string, bool) {
	if o == nil || common.IsNil(o.Message) {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *ServiceErrorDetails) HasMessage() bool {
	if o != nil && !common.IsNil(o.Message) {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *ServiceErrorDetails) SetMessage(v string) {
	o.Message = &v
}

// GetPspReference returns the PspReference field value if set, zero value otherwise.
func (o *ServiceErrorDetails) GetPspReference() string {
	if o == nil || common.IsNil(o.PspReference) {
		var ret string
		return ret
	}
	return *o.PspReference
}

// GetPspReferenceOk returns a tuple with the PspReference field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceErrorDetails) GetPspReferenceOk() (*string, bool) {
	if o == nil || common.IsNil(o.PspReference) {
		return nil, false
	}
	return o.PspReference, true
}

// HasPspReference returns a boolean if a field has been set.
func (o *ServiceErrorDetails) HasPspReference() bool {
	if o != nil && !common.IsNil(o.PspReference) {
		return true
	}

	return false
}

// SetPspReference gets a reference to the given string and assigns it to the PspReference field.
func (o *ServiceErrorDetails) SetPspReference(v string) {
	o.PspReference = &v
}

func (o ServiceErrorDetails) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ServiceErrorDetails) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.ErrorCode) {
		toSerialize["errorCode"] = o.ErrorCode
	}
	if !common.IsNil(o.ErrorType) {
		toSerialize["errorType"] = o.ErrorType
	}
	if !common.IsNil(o.Message) {
		toSerialize["message"] = o.Message
	}
	if !common.IsNil(o.PspReference) {
		toSerialize["pspReference"] = o.PspReference
	}
	return toSerialize, nil
}

type NullableServiceErrorDetails struct {
	value *ServiceErrorDetails
	isSet bool
}

func (v NullableServiceErrorDetails) Get() *ServiceErrorDetails {
	return v.value
}

func (v *NullableServiceErrorDetails) Set(val *ServiceErrorDetails) {
	v.value = val
	v.isSet = true
}

func (v NullableServiceErrorDetails) IsSet() bool {
	return v.isSet
}

func (v *NullableServiceErrorDetails) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableServiceErrorDetails(val *ServiceErrorDetails) *NullableServiceErrorDetails {
	return &NullableServiceErrorDetails{value: val, isSet: true}
}

func (v NullableServiceErrorDetails) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableServiceErrorDetails) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
