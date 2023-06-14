/*
Configuration API

API version: 2
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package balanceplatform

import (
	"encoding/json"

	"github.com/adyen/adyen-go-api-library/v7/src/common"
)

// checks if the PaymentInstrumentUpdateRequest type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &PaymentInstrumentUpdateRequest{}

// PaymentInstrumentUpdateRequest struct for PaymentInstrumentUpdateRequest
type PaymentInstrumentUpdateRequest struct {
	// The unique identifier of the balance account associated with this payment instrument. >You can only change the balance account ID if the payment instrument has **inactive** status.
	BalanceAccountId *string   `json:"balanceAccountId,omitempty"`
	Card             *CardInfo `json:"card,omitempty"`
	// The status of the payment instrument. If a status is not specified when creating a payment instrument, it is set to **active** by default. However, there can be exceptions for cards based on the `card.formFactor` and the `issuingCountryCode`. For example, when issuing physical cards in the US, the default status is **inactive**.  Possible values:    * **active**:  The payment instrument is active and can be used to make payments.    * **inactive**: The payment instrument is inactive and cannot be used to make payments.    * **suspended**: The payment instrument is suspended, either because it was stolen or lost.    * **closed**: The payment instrument is permanently closed. This action cannot be undone.
	Status *string `json:"status,omitempty"`
	// Comment for the status of the payment instrument.  Required if `statusReason` is **other**.
	StatusComment *string `json:"statusComment,omitempty"`
	// The reason for updating the status of the payment instrument.  Possible values: **lost**, **stolen**, **damaged**, **suspectedFraud**, **expired**, **endOfLife**, **accountClosure**, **other**. If the reason is **other**, you must also send the `statusComment` parameter describing the status change.
	StatusReason *string `json:"statusReason,omitempty"`
}

// NewPaymentInstrumentUpdateRequest instantiates a new PaymentInstrumentUpdateRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPaymentInstrumentUpdateRequest() *PaymentInstrumentUpdateRequest {
	this := PaymentInstrumentUpdateRequest{}
	return &this
}

// NewPaymentInstrumentUpdateRequestWithDefaults instantiates a new PaymentInstrumentUpdateRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPaymentInstrumentUpdateRequestWithDefaults() *PaymentInstrumentUpdateRequest {
	this := PaymentInstrumentUpdateRequest{}
	return &this
}

// GetBalanceAccountId returns the BalanceAccountId field value if set, zero value otherwise.
func (o *PaymentInstrumentUpdateRequest) GetBalanceAccountId() string {
	if o == nil || common.IsNil(o.BalanceAccountId) {
		var ret string
		return ret
	}
	return *o.BalanceAccountId
}

// GetBalanceAccountIdOk returns a tuple with the BalanceAccountId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentInstrumentUpdateRequest) GetBalanceAccountIdOk() (*string, bool) {
	if o == nil || common.IsNil(o.BalanceAccountId) {
		return nil, false
	}
	return o.BalanceAccountId, true
}

// HasBalanceAccountId returns a boolean if a field has been set.
func (o *PaymentInstrumentUpdateRequest) HasBalanceAccountId() bool {
	if o != nil && !common.IsNil(o.BalanceAccountId) {
		return true
	}

	return false
}

// SetBalanceAccountId gets a reference to the given string and assigns it to the BalanceAccountId field.
func (o *PaymentInstrumentUpdateRequest) SetBalanceAccountId(v string) {
	o.BalanceAccountId = &v
}

// GetCard returns the Card field value if set, zero value otherwise.
func (o *PaymentInstrumentUpdateRequest) GetCard() CardInfo {
	if o == nil || common.IsNil(o.Card) {
		var ret CardInfo
		return ret
	}
	return *o.Card
}

// GetCardOk returns a tuple with the Card field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentInstrumentUpdateRequest) GetCardOk() (*CardInfo, bool) {
	if o == nil || common.IsNil(o.Card) {
		return nil, false
	}
	return o.Card, true
}

// HasCard returns a boolean if a field has been set.
func (o *PaymentInstrumentUpdateRequest) HasCard() bool {
	if o != nil && !common.IsNil(o.Card) {
		return true
	}

	return false
}

// SetCard gets a reference to the given CardInfo and assigns it to the Card field.
func (o *PaymentInstrumentUpdateRequest) SetCard(v CardInfo) {
	o.Card = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *PaymentInstrumentUpdateRequest) GetStatus() string {
	if o == nil || common.IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentInstrumentUpdateRequest) GetStatusOk() (*string, bool) {
	if o == nil || common.IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *PaymentInstrumentUpdateRequest) HasStatus() bool {
	if o != nil && !common.IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *PaymentInstrumentUpdateRequest) SetStatus(v string) {
	o.Status = &v
}

// GetStatusComment returns the StatusComment field value if set, zero value otherwise.
func (o *PaymentInstrumentUpdateRequest) GetStatusComment() string {
	if o == nil || common.IsNil(o.StatusComment) {
		var ret string
		return ret
	}
	return *o.StatusComment
}

// GetStatusCommentOk returns a tuple with the StatusComment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentInstrumentUpdateRequest) GetStatusCommentOk() (*string, bool) {
	if o == nil || common.IsNil(o.StatusComment) {
		return nil, false
	}
	return o.StatusComment, true
}

// HasStatusComment returns a boolean if a field has been set.
func (o *PaymentInstrumentUpdateRequest) HasStatusComment() bool {
	if o != nil && !common.IsNil(o.StatusComment) {
		return true
	}

	return false
}

// SetStatusComment gets a reference to the given string and assigns it to the StatusComment field.
func (o *PaymentInstrumentUpdateRequest) SetStatusComment(v string) {
	o.StatusComment = &v
}

// GetStatusReason returns the StatusReason field value if set, zero value otherwise.
func (o *PaymentInstrumentUpdateRequest) GetStatusReason() string {
	if o == nil || common.IsNil(o.StatusReason) {
		var ret string
		return ret
	}
	return *o.StatusReason
}

// GetStatusReasonOk returns a tuple with the StatusReason field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentInstrumentUpdateRequest) GetStatusReasonOk() (*string, bool) {
	if o == nil || common.IsNil(o.StatusReason) {
		return nil, false
	}
	return o.StatusReason, true
}

// HasStatusReason returns a boolean if a field has been set.
func (o *PaymentInstrumentUpdateRequest) HasStatusReason() bool {
	if o != nil && !common.IsNil(o.StatusReason) {
		return true
	}

	return false
}

// SetStatusReason gets a reference to the given string and assigns it to the StatusReason field.
func (o *PaymentInstrumentUpdateRequest) SetStatusReason(v string) {
	o.StatusReason = &v
}

func (o PaymentInstrumentUpdateRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PaymentInstrumentUpdateRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.BalanceAccountId) {
		toSerialize["balanceAccountId"] = o.BalanceAccountId
	}
	if !common.IsNil(o.Card) {
		toSerialize["card"] = o.Card
	}
	if !common.IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !common.IsNil(o.StatusComment) {
		toSerialize["statusComment"] = o.StatusComment
	}
	if !common.IsNil(o.StatusReason) {
		toSerialize["statusReason"] = o.StatusReason
	}
	return toSerialize, nil
}

type NullablePaymentInstrumentUpdateRequest struct {
	value *PaymentInstrumentUpdateRequest
	isSet bool
}

func (v NullablePaymentInstrumentUpdateRequest) Get() *PaymentInstrumentUpdateRequest {
	return v.value
}

func (v *NullablePaymentInstrumentUpdateRequest) Set(val *PaymentInstrumentUpdateRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePaymentInstrumentUpdateRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePaymentInstrumentUpdateRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaymentInstrumentUpdateRequest(val *PaymentInstrumentUpdateRequest) *NullablePaymentInstrumentUpdateRequest {
	return &NullablePaymentInstrumentUpdateRequest{value: val, isSet: true}
}

func (v NullablePaymentInstrumentUpdateRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaymentInstrumentUpdateRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

func (o *PaymentInstrumentUpdateRequest) isValidStatus() bool {
	var allowedEnumValues = []string{"active", "closed", "inactive", "suspended"}
	for _, allowed := range allowedEnumValues {
		if o.GetStatus() == allowed {
			return true
		}
	}
	return false
}
func (o *PaymentInstrumentUpdateRequest) isValidStatusReason() bool {
	var allowedEnumValues = []string{"accountClosure", "damaged", "endOfLife", "expired", "lost", "other", "stolen", "suspectedFraud", "transactionRule"}
	for _, allowed := range allowedEnumValues {
		if o.GetStatusReason() == allowed {
			return true
		}
	}
	return false
}
