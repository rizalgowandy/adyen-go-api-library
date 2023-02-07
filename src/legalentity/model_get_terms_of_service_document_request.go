/*
Legal Entity Management API

The Legal Entity Management API enables you to manage legal entities that contain information required for verification.  ## Authentication To connect to the Legal Entity Management API, you must use the basic authentication credentials of your web service user. If you don't have one, contact the [Adyen Support Team](https://www.adyen.help/hc/en-us/requests/new). Use the web service user credentials to authenticate your request, for example:  ``` curl -U \"ws12345@Scope.BalancePlatform_YourBalancePlatform\":\"YourWsPassword\" \\ -H \"Content-Type: application/json\" \\ ... ``` Note that when going live, you need to generate new web service user credentials to access the [live endpoints](https://docs.adyen.com/development-resources/live-endpoints).  ## Versioning The Legal Entity Management API supports versioning of its endpoints through a version suffix in the endpoint URL. This suffix has the following format: \"vXX\", where XX is the version number.  For example: ``` https://kyc-test.adyen.com/lem/v2/legalEntities ``` ## Going live When going live, your Adyen contact will provide your API credential for the live environment. You can then use the username and password to send requests to `https://kyc-live.adyen.com/lem/v2`.  

API version: 2
Contact: developer-experience@adyen.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package legalentity

import (
	"encoding/json"
)

// checks if the GetTermsOfServiceDocumentRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetTermsOfServiceDocumentRequest{}

// GetTermsOfServiceDocumentRequest struct for GetTermsOfServiceDocumentRequest
type GetTermsOfServiceDocumentRequest struct {
	// The language to be used for the Terms of Service document, specified by the two letter [ISO 639-1](https://en.wikipedia.org/wiki/List_of_ISO_639-1_codes) language code. For example, **nl** for Dutch.
	Language *string `json:"language,omitempty"`
	// The type of Terms of Service.
	Type *string `json:"type,omitempty"`
}

// NewGetTermsOfServiceDocumentRequest instantiates a new GetTermsOfServiceDocumentRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetTermsOfServiceDocumentRequest() *GetTermsOfServiceDocumentRequest {
	this := GetTermsOfServiceDocumentRequest{}
	return &this
}

// NewGetTermsOfServiceDocumentRequestWithDefaults instantiates a new GetTermsOfServiceDocumentRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetTermsOfServiceDocumentRequestWithDefaults() *GetTermsOfServiceDocumentRequest {
	this := GetTermsOfServiceDocumentRequest{}
	return &this
}

// GetLanguage returns the Language field value if set, zero value otherwise.
func (o *GetTermsOfServiceDocumentRequest) GetLanguage() string {
	if o == nil || IsNil(o.Language) {
		var ret string
		return ret
	}
	return *o.Language
}

// GetLanguageOk returns a tuple with the Language field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetTermsOfServiceDocumentRequest) GetLanguageOk() (*string, bool) {
	if o == nil || IsNil(o.Language) {
		return nil, false
	}
	return o.Language, true
}

// HasLanguage returns a boolean if a field has been set.
func (o *GetTermsOfServiceDocumentRequest) HasLanguage() bool {
	if o != nil && !IsNil(o.Language) {
		return true
	}

	return false
}

// SetLanguage gets a reference to the given string and assigns it to the Language field.
func (o *GetTermsOfServiceDocumentRequest) SetLanguage(v string) {
	o.Language = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *GetTermsOfServiceDocumentRequest) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetTermsOfServiceDocumentRequest) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *GetTermsOfServiceDocumentRequest) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *GetTermsOfServiceDocumentRequest) SetType(v string) {
	o.Type = &v
}

func (o GetTermsOfServiceDocumentRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetTermsOfServiceDocumentRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Language) {
		toSerialize["language"] = o.Language
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	return toSerialize, nil
}

type NullableGetTermsOfServiceDocumentRequest struct {
	value *GetTermsOfServiceDocumentRequest
	isSet bool
}

func (v NullableGetTermsOfServiceDocumentRequest) Get() *GetTermsOfServiceDocumentRequest {
	return v.value
}

func (v *NullableGetTermsOfServiceDocumentRequest) Set(val *GetTermsOfServiceDocumentRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableGetTermsOfServiceDocumentRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableGetTermsOfServiceDocumentRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetTermsOfServiceDocumentRequest(val *GetTermsOfServiceDocumentRequest) *NullableGetTermsOfServiceDocumentRequest {
	return &NullableGetTermsOfServiceDocumentRequest{value: val, isSet: true}
}

func (v NullableGetTermsOfServiceDocumentRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetTermsOfServiceDocumentRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


