/*
Adyen Payment API

API version: 68
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package payments

import (
	"encoding/json"

	"github.com/adyen/adyen-go-api-library/v7/src/common"
)

// checks if the AdditionalDataRetry type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &AdditionalDataRetry{}

// AdditionalDataRetry struct for AdditionalDataRetry
type AdditionalDataRetry struct {
	// The number of times the transaction (not order) has been retried between different payment service providers. For instance, the `chainAttemptNumber` set to 2 means that this transaction has been recently tried on another provider before being sent to Adyen.  > If you submit `retry.chainAttemptNumber`, `retry.orderAttemptNumber`, and `retry.skipRetry` values, we also recommend you provide the `merchantOrderReference` to facilitate linking payment attempts together.
	RetryChainAttemptNumber *string `json:"retry.chainAttemptNumber,omitempty"`
	// The index of the attempt to bill a particular order, which is identified by the `merchantOrderReference` field. For example, if a recurring transaction fails and is retried one day later, then the order number for these attempts would be 1 and 2, respectively.  > If you submit `retry.chainAttemptNumber`, `retry.orderAttemptNumber`, and `retry.skipRetry` values, we also recommend you provide the `merchantOrderReference` to facilitate linking payment attempts together.
	RetryOrderAttemptNumber *string `json:"retry.orderAttemptNumber,omitempty"`
	// The Boolean value indicating whether Adyen should skip or retry this transaction, if possible.  > If you submit `retry.chainAttemptNumber`, `retry.orderAttemptNumber`, and `retry.skipRetry` values, we also recommend you provide the `merchantOrderReference` to facilitate linking payment attempts together.
	RetrySkipRetry *string `json:"retry.skipRetry,omitempty"`
}

// NewAdditionalDataRetry instantiates a new AdditionalDataRetry object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAdditionalDataRetry() *AdditionalDataRetry {
	this := AdditionalDataRetry{}
	return &this
}

// NewAdditionalDataRetryWithDefaults instantiates a new AdditionalDataRetry object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAdditionalDataRetryWithDefaults() *AdditionalDataRetry {
	this := AdditionalDataRetry{}
	return &this
}

// GetRetryChainAttemptNumber returns the RetryChainAttemptNumber field value if set, zero value otherwise.
func (o *AdditionalDataRetry) GetRetryChainAttemptNumber() string {
	if o == nil || common.IsNil(o.RetryChainAttemptNumber) {
		var ret string
		return ret
	}
	return *o.RetryChainAttemptNumber
}

// GetRetryChainAttemptNumberOk returns a tuple with the RetryChainAttemptNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdditionalDataRetry) GetRetryChainAttemptNumberOk() (*string, bool) {
	if o == nil || common.IsNil(o.RetryChainAttemptNumber) {
		return nil, false
	}
	return o.RetryChainAttemptNumber, true
}

// HasRetryChainAttemptNumber returns a boolean if a field has been set.
func (o *AdditionalDataRetry) HasRetryChainAttemptNumber() bool {
	if o != nil && !common.IsNil(o.RetryChainAttemptNumber) {
		return true
	}

	return false
}

// SetRetryChainAttemptNumber gets a reference to the given string and assigns it to the RetryChainAttemptNumber field.
func (o *AdditionalDataRetry) SetRetryChainAttemptNumber(v string) {
	o.RetryChainAttemptNumber = &v
}

// GetRetryOrderAttemptNumber returns the RetryOrderAttemptNumber field value if set, zero value otherwise.
func (o *AdditionalDataRetry) GetRetryOrderAttemptNumber() string {
	if o == nil || common.IsNil(o.RetryOrderAttemptNumber) {
		var ret string
		return ret
	}
	return *o.RetryOrderAttemptNumber
}

// GetRetryOrderAttemptNumberOk returns a tuple with the RetryOrderAttemptNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdditionalDataRetry) GetRetryOrderAttemptNumberOk() (*string, bool) {
	if o == nil || common.IsNil(o.RetryOrderAttemptNumber) {
		return nil, false
	}
	return o.RetryOrderAttemptNumber, true
}

// HasRetryOrderAttemptNumber returns a boolean if a field has been set.
func (o *AdditionalDataRetry) HasRetryOrderAttemptNumber() bool {
	if o != nil && !common.IsNil(o.RetryOrderAttemptNumber) {
		return true
	}

	return false
}

// SetRetryOrderAttemptNumber gets a reference to the given string and assigns it to the RetryOrderAttemptNumber field.
func (o *AdditionalDataRetry) SetRetryOrderAttemptNumber(v string) {
	o.RetryOrderAttemptNumber = &v
}

// GetRetrySkipRetry returns the RetrySkipRetry field value if set, zero value otherwise.
func (o *AdditionalDataRetry) GetRetrySkipRetry() string {
	if o == nil || common.IsNil(o.RetrySkipRetry) {
		var ret string
		return ret
	}
	return *o.RetrySkipRetry
}

// GetRetrySkipRetryOk returns a tuple with the RetrySkipRetry field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdditionalDataRetry) GetRetrySkipRetryOk() (*string, bool) {
	if o == nil || common.IsNil(o.RetrySkipRetry) {
		return nil, false
	}
	return o.RetrySkipRetry, true
}

// HasRetrySkipRetry returns a boolean if a field has been set.
func (o *AdditionalDataRetry) HasRetrySkipRetry() bool {
	if o != nil && !common.IsNil(o.RetrySkipRetry) {
		return true
	}

	return false
}

// SetRetrySkipRetry gets a reference to the given string and assigns it to the RetrySkipRetry field.
func (o *AdditionalDataRetry) SetRetrySkipRetry(v string) {
	o.RetrySkipRetry = &v
}

func (o AdditionalDataRetry) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AdditionalDataRetry) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.RetryChainAttemptNumber) {
		toSerialize["retry.chainAttemptNumber"] = o.RetryChainAttemptNumber
	}
	if !common.IsNil(o.RetryOrderAttemptNumber) {
		toSerialize["retry.orderAttemptNumber"] = o.RetryOrderAttemptNumber
	}
	if !common.IsNil(o.RetrySkipRetry) {
		toSerialize["retry.skipRetry"] = o.RetrySkipRetry
	}
	return toSerialize, nil
}

type NullableAdditionalDataRetry struct {
	value *AdditionalDataRetry
	isSet bool
}

func (v NullableAdditionalDataRetry) Get() *AdditionalDataRetry {
	return v.value
}

func (v *NullableAdditionalDataRetry) Set(val *AdditionalDataRetry) {
	v.value = val
	v.isSet = true
}

func (v NullableAdditionalDataRetry) IsSet() bool {
	return v.isSet
}

func (v *NullableAdditionalDataRetry) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAdditionalDataRetry(val *AdditionalDataRetry) *NullableAdditionalDataRetry {
	return &NullableAdditionalDataRetry{value: val, isSet: true}
}

func (v NullableAdditionalDataRetry) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAdditionalDataRetry) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
