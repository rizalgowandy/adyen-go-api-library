/*
Adyen Recurring API

API version: 68
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package recurring

import (
	"encoding/json"

	"github.com/adyen/adyen-go-api-library/v6/src/common"
)

// checks if the ScheduleAccountUpdaterRequest type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &ScheduleAccountUpdaterRequest{}

// ScheduleAccountUpdaterRequest struct for ScheduleAccountUpdaterRequest
type ScheduleAccountUpdaterRequest struct {
	// This field contains additional data, which may be required for a particular request.
	AdditionalData *map[string]string `json:"additionalData,omitempty"`
	Card           *Card              `json:"card,omitempty"`
	// Account of the merchant.
	MerchantAccount string `json:"merchantAccount"`
	// A reference that merchants can apply for the call.
	Reference string `json:"reference"`
	// The selected detail recurring reference.  Optional if `card` is provided.
	SelectedRecurringDetailReference *string `json:"selectedRecurringDetailReference,omitempty"`
	// The reference of the shopper that owns the recurring contract.  Optional if `card` is provided.
	ShopperReference *string `json:"shopperReference,omitempty"`
}

// NewScheduleAccountUpdaterRequest instantiates a new ScheduleAccountUpdaterRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewScheduleAccountUpdaterRequest(merchantAccount string, reference string) *ScheduleAccountUpdaterRequest {
	this := ScheduleAccountUpdaterRequest{}
	this.MerchantAccount = merchantAccount
	this.Reference = reference
	return &this
}

// NewScheduleAccountUpdaterRequestWithDefaults instantiates a new ScheduleAccountUpdaterRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewScheduleAccountUpdaterRequestWithDefaults() *ScheduleAccountUpdaterRequest {
	this := ScheduleAccountUpdaterRequest{}
	return &this
}

// GetAdditionalData returns the AdditionalData field value if set, zero value otherwise.
func (o *ScheduleAccountUpdaterRequest) GetAdditionalData() map[string]string {
	if o == nil || common.IsNil(o.AdditionalData) {
		var ret map[string]string
		return ret
	}
	return *o.AdditionalData
}

// GetAdditionalDataOk returns a tuple with the AdditionalData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScheduleAccountUpdaterRequest) GetAdditionalDataOk() (*map[string]string, bool) {
	if o == nil || common.IsNil(o.AdditionalData) {
		return nil, false
	}
	return o.AdditionalData, true
}

// HasAdditionalData returns a boolean if a field has been set.
func (o *ScheduleAccountUpdaterRequest) HasAdditionalData() bool {
	if o != nil && !common.IsNil(o.AdditionalData) {
		return true
	}

	return false
}

// SetAdditionalData gets a reference to the given map[string]string and assigns it to the AdditionalData field.
func (o *ScheduleAccountUpdaterRequest) SetAdditionalData(v map[string]string) {
	o.AdditionalData = &v
}

// GetCard returns the Card field value if set, zero value otherwise.
func (o *ScheduleAccountUpdaterRequest) GetCard() Card {
	if o == nil || common.IsNil(o.Card) {
		var ret Card
		return ret
	}
	return *o.Card
}

// GetCardOk returns a tuple with the Card field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScheduleAccountUpdaterRequest) GetCardOk() (*Card, bool) {
	if o == nil || common.IsNil(o.Card) {
		return nil, false
	}
	return o.Card, true
}

// HasCard returns a boolean if a field has been set.
func (o *ScheduleAccountUpdaterRequest) HasCard() bool {
	if o != nil && !common.IsNil(o.Card) {
		return true
	}

	return false
}

// SetCard gets a reference to the given Card and assigns it to the Card field.
func (o *ScheduleAccountUpdaterRequest) SetCard(v Card) {
	o.Card = &v
}

// GetMerchantAccount returns the MerchantAccount field value
func (o *ScheduleAccountUpdaterRequest) GetMerchantAccount() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MerchantAccount
}

// GetMerchantAccountOk returns a tuple with the MerchantAccount field value
// and a boolean to check if the value has been set.
func (o *ScheduleAccountUpdaterRequest) GetMerchantAccountOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MerchantAccount, true
}

// SetMerchantAccount sets field value
func (o *ScheduleAccountUpdaterRequest) SetMerchantAccount(v string) {
	o.MerchantAccount = v
}

// GetReference returns the Reference field value
func (o *ScheduleAccountUpdaterRequest) GetReference() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Reference
}

// GetReferenceOk returns a tuple with the Reference field value
// and a boolean to check if the value has been set.
func (o *ScheduleAccountUpdaterRequest) GetReferenceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Reference, true
}

// SetReference sets field value
func (o *ScheduleAccountUpdaterRequest) SetReference(v string) {
	o.Reference = v
}

// GetSelectedRecurringDetailReference returns the SelectedRecurringDetailReference field value if set, zero value otherwise.
func (o *ScheduleAccountUpdaterRequest) GetSelectedRecurringDetailReference() string {
	if o == nil || common.IsNil(o.SelectedRecurringDetailReference) {
		var ret string
		return ret
	}
	return *o.SelectedRecurringDetailReference
}

// GetSelectedRecurringDetailReferenceOk returns a tuple with the SelectedRecurringDetailReference field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScheduleAccountUpdaterRequest) GetSelectedRecurringDetailReferenceOk() (*string, bool) {
	if o == nil || common.IsNil(o.SelectedRecurringDetailReference) {
		return nil, false
	}
	return o.SelectedRecurringDetailReference, true
}

// HasSelectedRecurringDetailReference returns a boolean if a field has been set.
func (o *ScheduleAccountUpdaterRequest) HasSelectedRecurringDetailReference() bool {
	if o != nil && !common.IsNil(o.SelectedRecurringDetailReference) {
		return true
	}

	return false
}

// SetSelectedRecurringDetailReference gets a reference to the given string and assigns it to the SelectedRecurringDetailReference field.
func (o *ScheduleAccountUpdaterRequest) SetSelectedRecurringDetailReference(v string) {
	o.SelectedRecurringDetailReference = &v
}

// GetShopperReference returns the ShopperReference field value if set, zero value otherwise.
func (o *ScheduleAccountUpdaterRequest) GetShopperReference() string {
	if o == nil || common.IsNil(o.ShopperReference) {
		var ret string
		return ret
	}
	return *o.ShopperReference
}

// GetShopperReferenceOk returns a tuple with the ShopperReference field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScheduleAccountUpdaterRequest) GetShopperReferenceOk() (*string, bool) {
	if o == nil || common.IsNil(o.ShopperReference) {
		return nil, false
	}
	return o.ShopperReference, true
}

// HasShopperReference returns a boolean if a field has been set.
func (o *ScheduleAccountUpdaterRequest) HasShopperReference() bool {
	if o != nil && !common.IsNil(o.ShopperReference) {
		return true
	}

	return false
}

// SetShopperReference gets a reference to the given string and assigns it to the ShopperReference field.
func (o *ScheduleAccountUpdaterRequest) SetShopperReference(v string) {
	o.ShopperReference = &v
}

func (o ScheduleAccountUpdaterRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ScheduleAccountUpdaterRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.AdditionalData) {
		toSerialize["additionalData"] = o.AdditionalData
	}
	if !common.IsNil(o.Card) {
		toSerialize["card"] = o.Card
	}
	toSerialize["merchantAccount"] = o.MerchantAccount
	toSerialize["reference"] = o.Reference
	if !common.IsNil(o.SelectedRecurringDetailReference) {
		toSerialize["selectedRecurringDetailReference"] = o.SelectedRecurringDetailReference
	}
	if !common.IsNil(o.ShopperReference) {
		toSerialize["shopperReference"] = o.ShopperReference
	}
	return toSerialize, nil
}

type NullableScheduleAccountUpdaterRequest struct {
	value *ScheduleAccountUpdaterRequest
	isSet bool
}

func (v NullableScheduleAccountUpdaterRequest) Get() *ScheduleAccountUpdaterRequest {
	return v.value
}

func (v *NullableScheduleAccountUpdaterRequest) Set(val *ScheduleAccountUpdaterRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableScheduleAccountUpdaterRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableScheduleAccountUpdaterRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableScheduleAccountUpdaterRequest(val *ScheduleAccountUpdaterRequest) *NullableScheduleAccountUpdaterRequest {
	return &NullableScheduleAccountUpdaterRequest{value: val, isSet: true}
}

func (v NullableScheduleAccountUpdaterRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableScheduleAccountUpdaterRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
