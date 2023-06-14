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

// checks if the CardOrderItemDeliveryStatus type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &CardOrderItemDeliveryStatus{}

// CardOrderItemDeliveryStatus struct for CardOrderItemDeliveryStatus
type CardOrderItemDeliveryStatus struct {
	// Status of the delivery.
	Status *string `json:"status,omitempty"`
	// Error status, if any.
	StatusError *string `json:"statusError,omitempty"`
	// Error message, if any.
	StatusErrorMessage *string `json:"statusErrorMessage,omitempty"`
	// Tracking number of the delivery.
	TrackingNumber *string `json:"trackingNumber,omitempty"`
}

// NewCardOrderItemDeliveryStatus instantiates a new CardOrderItemDeliveryStatus object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCardOrderItemDeliveryStatus() *CardOrderItemDeliveryStatus {
	this := CardOrderItemDeliveryStatus{}
	return &this
}

// NewCardOrderItemDeliveryStatusWithDefaults instantiates a new CardOrderItemDeliveryStatus object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCardOrderItemDeliveryStatusWithDefaults() *CardOrderItemDeliveryStatus {
	this := CardOrderItemDeliveryStatus{}
	return &this
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *CardOrderItemDeliveryStatus) GetStatus() string {
	if o == nil || common.IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CardOrderItemDeliveryStatus) GetStatusOk() (*string, bool) {
	if o == nil || common.IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *CardOrderItemDeliveryStatus) HasStatus() bool {
	if o != nil && !common.IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *CardOrderItemDeliveryStatus) SetStatus(v string) {
	o.Status = &v
}

// GetStatusError returns the StatusError field value if set, zero value otherwise.
func (o *CardOrderItemDeliveryStatus) GetStatusError() string {
	if o == nil || common.IsNil(o.StatusError) {
		var ret string
		return ret
	}
	return *o.StatusError
}

// GetStatusErrorOk returns a tuple with the StatusError field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CardOrderItemDeliveryStatus) GetStatusErrorOk() (*string, bool) {
	if o == nil || common.IsNil(o.StatusError) {
		return nil, false
	}
	return o.StatusError, true
}

// HasStatusError returns a boolean if a field has been set.
func (o *CardOrderItemDeliveryStatus) HasStatusError() bool {
	if o != nil && !common.IsNil(o.StatusError) {
		return true
	}

	return false
}

// SetStatusError gets a reference to the given string and assigns it to the StatusError field.
func (o *CardOrderItemDeliveryStatus) SetStatusError(v string) {
	o.StatusError = &v
}

// GetStatusErrorMessage returns the StatusErrorMessage field value if set, zero value otherwise.
func (o *CardOrderItemDeliveryStatus) GetStatusErrorMessage() string {
	if o == nil || common.IsNil(o.StatusErrorMessage) {
		var ret string
		return ret
	}
	return *o.StatusErrorMessage
}

// GetStatusErrorMessageOk returns a tuple with the StatusErrorMessage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CardOrderItemDeliveryStatus) GetStatusErrorMessageOk() (*string, bool) {
	if o == nil || common.IsNil(o.StatusErrorMessage) {
		return nil, false
	}
	return o.StatusErrorMessage, true
}

// HasStatusErrorMessage returns a boolean if a field has been set.
func (o *CardOrderItemDeliveryStatus) HasStatusErrorMessage() bool {
	if o != nil && !common.IsNil(o.StatusErrorMessage) {
		return true
	}

	return false
}

// SetStatusErrorMessage gets a reference to the given string and assigns it to the StatusErrorMessage field.
func (o *CardOrderItemDeliveryStatus) SetStatusErrorMessage(v string) {
	o.StatusErrorMessage = &v
}

// GetTrackingNumber returns the TrackingNumber field value if set, zero value otherwise.
func (o *CardOrderItemDeliveryStatus) GetTrackingNumber() string {
	if o == nil || common.IsNil(o.TrackingNumber) {
		var ret string
		return ret
	}
	return *o.TrackingNumber
}

// GetTrackingNumberOk returns a tuple with the TrackingNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CardOrderItemDeliveryStatus) GetTrackingNumberOk() (*string, bool) {
	if o == nil || common.IsNil(o.TrackingNumber) {
		return nil, false
	}
	return o.TrackingNumber, true
}

// HasTrackingNumber returns a boolean if a field has been set.
func (o *CardOrderItemDeliveryStatus) HasTrackingNumber() bool {
	if o != nil && !common.IsNil(o.TrackingNumber) {
		return true
	}

	return false
}

// SetTrackingNumber gets a reference to the given string and assigns it to the TrackingNumber field.
func (o *CardOrderItemDeliveryStatus) SetTrackingNumber(v string) {
	o.TrackingNumber = &v
}

func (o CardOrderItemDeliveryStatus) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CardOrderItemDeliveryStatus) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !common.IsNil(o.StatusError) {
		toSerialize["statusError"] = o.StatusError
	}
	if !common.IsNil(o.StatusErrorMessage) {
		toSerialize["statusErrorMessage"] = o.StatusErrorMessage
	}
	if !common.IsNil(o.TrackingNumber) {
		toSerialize["trackingNumber"] = o.TrackingNumber
	}
	return toSerialize, nil
}

type NullableCardOrderItemDeliveryStatus struct {
	value *CardOrderItemDeliveryStatus
	isSet bool
}

func (v NullableCardOrderItemDeliveryStatus) Get() *CardOrderItemDeliveryStatus {
	return v.value
}

func (v *NullableCardOrderItemDeliveryStatus) Set(val *CardOrderItemDeliveryStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableCardOrderItemDeliveryStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableCardOrderItemDeliveryStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCardOrderItemDeliveryStatus(val *CardOrderItemDeliveryStatus) *NullableCardOrderItemDeliveryStatus {
	return &NullableCardOrderItemDeliveryStatus{value: val, isSet: true}
}

func (v NullableCardOrderItemDeliveryStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCardOrderItemDeliveryStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

func (o *CardOrderItemDeliveryStatus) isValidStatus() bool {
	var allowedEnumValues = []string{"created", "delivered", "processing", "produced", "rejected", "shipped", "unknown"}
	for _, allowed := range allowedEnumValues {
		if o.GetStatus() == allowed {
			return true
		}
	}
	return false
}
