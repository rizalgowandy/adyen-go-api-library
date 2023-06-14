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

// checks if the CapitalGrantAccount type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &CapitalGrantAccount{}

// CapitalGrantAccount struct for CapitalGrantAccount
type CapitalGrantAccount struct {
	// The balances of the grant account.
	Balances []CapitalBalance `json:"balances,omitempty"`
	// The unique identifier of the balance account used to fund the grant.
	FundingBalanceAccountId *string `json:"fundingBalanceAccountId,omitempty"`
	// The identifier of the grant account.
	Id *string `json:"id,omitempty"`
	// The limits of the grant account.
	Limits []GrantLimit `json:"limits,omitempty"`
}

// NewCapitalGrantAccount instantiates a new CapitalGrantAccount object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCapitalGrantAccount() *CapitalGrantAccount {
	this := CapitalGrantAccount{}
	return &this
}

// NewCapitalGrantAccountWithDefaults instantiates a new CapitalGrantAccount object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCapitalGrantAccountWithDefaults() *CapitalGrantAccount {
	this := CapitalGrantAccount{}
	return &this
}

// GetBalances returns the Balances field value if set, zero value otherwise.
func (o *CapitalGrantAccount) GetBalances() []CapitalBalance {
	if o == nil || common.IsNil(o.Balances) {
		var ret []CapitalBalance
		return ret
	}
	return o.Balances
}

// GetBalancesOk returns a tuple with the Balances field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CapitalGrantAccount) GetBalancesOk() ([]CapitalBalance, bool) {
	if o == nil || common.IsNil(o.Balances) {
		return nil, false
	}
	return o.Balances, true
}

// HasBalances returns a boolean if a field has been set.
func (o *CapitalGrantAccount) HasBalances() bool {
	if o != nil && !common.IsNil(o.Balances) {
		return true
	}

	return false
}

// SetBalances gets a reference to the given []CapitalBalance and assigns it to the Balances field.
func (o *CapitalGrantAccount) SetBalances(v []CapitalBalance) {
	o.Balances = v
}

// GetFundingBalanceAccountId returns the FundingBalanceAccountId field value if set, zero value otherwise.
func (o *CapitalGrantAccount) GetFundingBalanceAccountId() string {
	if o == nil || common.IsNil(o.FundingBalanceAccountId) {
		var ret string
		return ret
	}
	return *o.FundingBalanceAccountId
}

// GetFundingBalanceAccountIdOk returns a tuple with the FundingBalanceAccountId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CapitalGrantAccount) GetFundingBalanceAccountIdOk() (*string, bool) {
	if o == nil || common.IsNil(o.FundingBalanceAccountId) {
		return nil, false
	}
	return o.FundingBalanceAccountId, true
}

// HasFundingBalanceAccountId returns a boolean if a field has been set.
func (o *CapitalGrantAccount) HasFundingBalanceAccountId() bool {
	if o != nil && !common.IsNil(o.FundingBalanceAccountId) {
		return true
	}

	return false
}

// SetFundingBalanceAccountId gets a reference to the given string and assigns it to the FundingBalanceAccountId field.
func (o *CapitalGrantAccount) SetFundingBalanceAccountId(v string) {
	o.FundingBalanceAccountId = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *CapitalGrantAccount) GetId() string {
	if o == nil || common.IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CapitalGrantAccount) GetIdOk() (*string, bool) {
	if o == nil || common.IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *CapitalGrantAccount) HasId() bool {
	if o != nil && !common.IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *CapitalGrantAccount) SetId(v string) {
	o.Id = &v
}

// GetLimits returns the Limits field value if set, zero value otherwise.
func (o *CapitalGrantAccount) GetLimits() []GrantLimit {
	if o == nil || common.IsNil(o.Limits) {
		var ret []GrantLimit
		return ret
	}
	return o.Limits
}

// GetLimitsOk returns a tuple with the Limits field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CapitalGrantAccount) GetLimitsOk() ([]GrantLimit, bool) {
	if o == nil || common.IsNil(o.Limits) {
		return nil, false
	}
	return o.Limits, true
}

// HasLimits returns a boolean if a field has been set.
func (o *CapitalGrantAccount) HasLimits() bool {
	if o != nil && !common.IsNil(o.Limits) {
		return true
	}

	return false
}

// SetLimits gets a reference to the given []GrantLimit and assigns it to the Limits field.
func (o *CapitalGrantAccount) SetLimits(v []GrantLimit) {
	o.Limits = v
}

func (o CapitalGrantAccount) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CapitalGrantAccount) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.Balances) {
		toSerialize["balances"] = o.Balances
	}
	if !common.IsNil(o.FundingBalanceAccountId) {
		toSerialize["fundingBalanceAccountId"] = o.FundingBalanceAccountId
	}
	if !common.IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !common.IsNil(o.Limits) {
		toSerialize["limits"] = o.Limits
	}
	return toSerialize, nil
}

type NullableCapitalGrantAccount struct {
	value *CapitalGrantAccount
	isSet bool
}

func (v NullableCapitalGrantAccount) Get() *CapitalGrantAccount {
	return v.value
}

func (v *NullableCapitalGrantAccount) Set(val *CapitalGrantAccount) {
	v.value = val
	v.isSet = true
}

func (v NullableCapitalGrantAccount) IsSet() bool {
	return v.isSet
}

func (v *NullableCapitalGrantAccount) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCapitalGrantAccount(val *CapitalGrantAccount) *NullableCapitalGrantAccount {
	return &NullableCapitalGrantAccount{value: val, isSet: true}
}

func (v NullableCapitalGrantAccount) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCapitalGrantAccount) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
