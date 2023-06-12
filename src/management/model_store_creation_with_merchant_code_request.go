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

// checks if the StoreCreationWithMerchantCodeRequest type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &StoreCreationWithMerchantCodeRequest{}

// StoreCreationWithMerchantCodeRequest struct for StoreCreationWithMerchantCodeRequest
type StoreCreationWithMerchantCodeRequest struct {
	Address StoreLocation `json:"address"`
	// The unique identifiers of the [business lines](https://docs.adyen.com/api-explorer/#/legalentity/latest/post/businesslines__resParam_id) that the store is associated with. If not specified, the business line of the merchant account is used. Required when there are multiple business lines under the merchant account.
	BusinessLineIds []string `json:"businessLineIds,omitempty"`
	// Your description of the store.
	Description string `json:"description"`
	// When using the Zip payment method: The location ID that Zip has assigned to your store.
	ExternalReferenceId *string `json:"externalReferenceId,omitempty"`
	// The unique identifier of the merchant account that the store belongs to.
	MerchantId string `json:"merchantId"`
	// The phone number of the store, including '+' and country code.
	PhoneNumber string `json:"phoneNumber"`
	// Your reference to recognize the store by. Also known as the store code.  Allowed characters: Lowercase and uppercase letters without diacritics, numbers 0 through 9, hyphen (-), and underscore (_).
	Reference *string `json:"reference,omitempty"`
	// The store name to be shown on the shopper's bank or credit card statement and on the shopper receipt. Maximum length: 22 characters; can't be all numbers.
	ShopperStatement   string                   `json:"shopperStatement"`
	SplitConfiguration *StoreSplitConfiguration `json:"splitConfiguration,omitempty"`
}

// NewStoreCreationWithMerchantCodeRequest instantiates a new StoreCreationWithMerchantCodeRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStoreCreationWithMerchantCodeRequest(address StoreLocation, description string, merchantId string, phoneNumber string, shopperStatement string) *StoreCreationWithMerchantCodeRequest {
	this := StoreCreationWithMerchantCodeRequest{}
	this.Address = address
	this.Description = description
	this.MerchantId = merchantId
	this.PhoneNumber = phoneNumber
	this.ShopperStatement = shopperStatement
	return &this
}

// NewStoreCreationWithMerchantCodeRequestWithDefaults instantiates a new StoreCreationWithMerchantCodeRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStoreCreationWithMerchantCodeRequestWithDefaults() *StoreCreationWithMerchantCodeRequest {
	this := StoreCreationWithMerchantCodeRequest{}
	return &this
}

// GetAddress returns the Address field value
func (o *StoreCreationWithMerchantCodeRequest) GetAddress() StoreLocation {
	if o == nil {
		var ret StoreLocation
		return ret
	}

	return o.Address
}

// GetAddressOk returns a tuple with the Address field value
// and a boolean to check if the value has been set.
func (o *StoreCreationWithMerchantCodeRequest) GetAddressOk() (*StoreLocation, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Address, true
}

// SetAddress sets field value
func (o *StoreCreationWithMerchantCodeRequest) SetAddress(v StoreLocation) {
	o.Address = v
}

// GetBusinessLineIds returns the BusinessLineIds field value if set, zero value otherwise.
func (o *StoreCreationWithMerchantCodeRequest) GetBusinessLineIds() []string {
	if o == nil || common.IsNil(o.BusinessLineIds) {
		var ret []string
		return ret
	}
	return o.BusinessLineIds
}

// GetBusinessLineIdsOk returns a tuple with the BusinessLineIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StoreCreationWithMerchantCodeRequest) GetBusinessLineIdsOk() ([]string, bool) {
	if o == nil || common.IsNil(o.BusinessLineIds) {
		return nil, false
	}
	return o.BusinessLineIds, true
}

// HasBusinessLineIds returns a boolean if a field has been set.
func (o *StoreCreationWithMerchantCodeRequest) HasBusinessLineIds() bool {
	if o != nil && !common.IsNil(o.BusinessLineIds) {
		return true
	}

	return false
}

// SetBusinessLineIds gets a reference to the given []string and assigns it to the BusinessLineIds field.
func (o *StoreCreationWithMerchantCodeRequest) SetBusinessLineIds(v []string) {
	o.BusinessLineIds = v
}

// GetDescription returns the Description field value
func (o *StoreCreationWithMerchantCodeRequest) GetDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *StoreCreationWithMerchantCodeRequest) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Description, true
}

// SetDescription sets field value
func (o *StoreCreationWithMerchantCodeRequest) SetDescription(v string) {
	o.Description = v
}

// GetExternalReferenceId returns the ExternalReferenceId field value if set, zero value otherwise.
func (o *StoreCreationWithMerchantCodeRequest) GetExternalReferenceId() string {
	if o == nil || common.IsNil(o.ExternalReferenceId) {
		var ret string
		return ret
	}
	return *o.ExternalReferenceId
}

// GetExternalReferenceIdOk returns a tuple with the ExternalReferenceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StoreCreationWithMerchantCodeRequest) GetExternalReferenceIdOk() (*string, bool) {
	if o == nil || common.IsNil(o.ExternalReferenceId) {
		return nil, false
	}
	return o.ExternalReferenceId, true
}

// HasExternalReferenceId returns a boolean if a field has been set.
func (o *StoreCreationWithMerchantCodeRequest) HasExternalReferenceId() bool {
	if o != nil && !common.IsNil(o.ExternalReferenceId) {
		return true
	}

	return false
}

// SetExternalReferenceId gets a reference to the given string and assigns it to the ExternalReferenceId field.
func (o *StoreCreationWithMerchantCodeRequest) SetExternalReferenceId(v string) {
	o.ExternalReferenceId = &v
}

// GetMerchantId returns the MerchantId field value
func (o *StoreCreationWithMerchantCodeRequest) GetMerchantId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MerchantId
}

// GetMerchantIdOk returns a tuple with the MerchantId field value
// and a boolean to check if the value has been set.
func (o *StoreCreationWithMerchantCodeRequest) GetMerchantIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MerchantId, true
}

// SetMerchantId sets field value
func (o *StoreCreationWithMerchantCodeRequest) SetMerchantId(v string) {
	o.MerchantId = v
}

// GetPhoneNumber returns the PhoneNumber field value
func (o *StoreCreationWithMerchantCodeRequest) GetPhoneNumber() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PhoneNumber
}

// GetPhoneNumberOk returns a tuple with the PhoneNumber field value
// and a boolean to check if the value has been set.
func (o *StoreCreationWithMerchantCodeRequest) GetPhoneNumberOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PhoneNumber, true
}

// SetPhoneNumber sets field value
func (o *StoreCreationWithMerchantCodeRequest) SetPhoneNumber(v string) {
	o.PhoneNumber = v
}

// GetReference returns the Reference field value if set, zero value otherwise.
func (o *StoreCreationWithMerchantCodeRequest) GetReference() string {
	if o == nil || common.IsNil(o.Reference) {
		var ret string
		return ret
	}
	return *o.Reference
}

// GetReferenceOk returns a tuple with the Reference field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StoreCreationWithMerchantCodeRequest) GetReferenceOk() (*string, bool) {
	if o == nil || common.IsNil(o.Reference) {
		return nil, false
	}
	return o.Reference, true
}

// HasReference returns a boolean if a field has been set.
func (o *StoreCreationWithMerchantCodeRequest) HasReference() bool {
	if o != nil && !common.IsNil(o.Reference) {
		return true
	}

	return false
}

// SetReference gets a reference to the given string and assigns it to the Reference field.
func (o *StoreCreationWithMerchantCodeRequest) SetReference(v string) {
	o.Reference = &v
}

// GetShopperStatement returns the ShopperStatement field value
func (o *StoreCreationWithMerchantCodeRequest) GetShopperStatement() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ShopperStatement
}

// GetShopperStatementOk returns a tuple with the ShopperStatement field value
// and a boolean to check if the value has been set.
func (o *StoreCreationWithMerchantCodeRequest) GetShopperStatementOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ShopperStatement, true
}

// SetShopperStatement sets field value
func (o *StoreCreationWithMerchantCodeRequest) SetShopperStatement(v string) {
	o.ShopperStatement = v
}

// GetSplitConfiguration returns the SplitConfiguration field value if set, zero value otherwise.
func (o *StoreCreationWithMerchantCodeRequest) GetSplitConfiguration() StoreSplitConfiguration {
	if o == nil || common.IsNil(o.SplitConfiguration) {
		var ret StoreSplitConfiguration
		return ret
	}
	return *o.SplitConfiguration
}

// GetSplitConfigurationOk returns a tuple with the SplitConfiguration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StoreCreationWithMerchantCodeRequest) GetSplitConfigurationOk() (*StoreSplitConfiguration, bool) {
	if o == nil || common.IsNil(o.SplitConfiguration) {
		return nil, false
	}
	return o.SplitConfiguration, true
}

// HasSplitConfiguration returns a boolean if a field has been set.
func (o *StoreCreationWithMerchantCodeRequest) HasSplitConfiguration() bool {
	if o != nil && !common.IsNil(o.SplitConfiguration) {
		return true
	}

	return false
}

// SetSplitConfiguration gets a reference to the given StoreSplitConfiguration and assigns it to the SplitConfiguration field.
func (o *StoreCreationWithMerchantCodeRequest) SetSplitConfiguration(v StoreSplitConfiguration) {
	o.SplitConfiguration = &v
}

func (o StoreCreationWithMerchantCodeRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o StoreCreationWithMerchantCodeRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["address"] = o.Address
	if !common.IsNil(o.BusinessLineIds) {
		toSerialize["businessLineIds"] = o.BusinessLineIds
	}
	toSerialize["description"] = o.Description
	if !common.IsNil(o.ExternalReferenceId) {
		toSerialize["externalReferenceId"] = o.ExternalReferenceId
	}
	toSerialize["merchantId"] = o.MerchantId
	toSerialize["phoneNumber"] = o.PhoneNumber
	if !common.IsNil(o.Reference) {
		toSerialize["reference"] = o.Reference
	}
	toSerialize["shopperStatement"] = o.ShopperStatement
	if !common.IsNil(o.SplitConfiguration) {
		toSerialize["splitConfiguration"] = o.SplitConfiguration
	}
	return toSerialize, nil
}

type NullableStoreCreationWithMerchantCodeRequest struct {
	value *StoreCreationWithMerchantCodeRequest
	isSet bool
}

func (v NullableStoreCreationWithMerchantCodeRequest) Get() *StoreCreationWithMerchantCodeRequest {
	return v.value
}

func (v *NullableStoreCreationWithMerchantCodeRequest) Set(val *StoreCreationWithMerchantCodeRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableStoreCreationWithMerchantCodeRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableStoreCreationWithMerchantCodeRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStoreCreationWithMerchantCodeRequest(val *StoreCreationWithMerchantCodeRequest) *NullableStoreCreationWithMerchantCodeRequest {
	return &NullableStoreCreationWithMerchantCodeRequest{value: val, isSet: true}
}

func (v NullableStoreCreationWithMerchantCodeRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStoreCreationWithMerchantCodeRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
