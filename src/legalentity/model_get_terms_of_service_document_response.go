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

// checks if the GetTermsOfServiceDocumentResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetTermsOfServiceDocumentResponse{}

// GetTermsOfServiceDocumentResponse struct for GetTermsOfServiceDocumentResponse
type GetTermsOfServiceDocumentResponse struct {
	// The Terms of Service document in Base64-encoded format.
	Document *string `json:"document,omitempty"`
	// The unique identifier of the legal entity.
	Id *string `json:"id,omitempty"`
	// The language used for the Terms of Service document, specified by the two letter [ISO 639-1](https://en.wikipedia.org/wiki/List_of_ISO_639-1_codes) language code. For example, **nl** for Dutch.
	Language *string `json:"language,omitempty"`
	// The unique identifier of the Terms of Service document.
	TermsOfServiceDocumentId *string `json:"termsOfServiceDocumentId,omitempty"`
	// The type of Terms of Service.
	Type *string `json:"type,omitempty"`
}

// NewGetTermsOfServiceDocumentResponse instantiates a new GetTermsOfServiceDocumentResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetTermsOfServiceDocumentResponse() *GetTermsOfServiceDocumentResponse {
	this := GetTermsOfServiceDocumentResponse{}
	return &this
}

// NewGetTermsOfServiceDocumentResponseWithDefaults instantiates a new GetTermsOfServiceDocumentResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetTermsOfServiceDocumentResponseWithDefaults() *GetTermsOfServiceDocumentResponse {
	this := GetTermsOfServiceDocumentResponse{}
	return &this
}

// GetDocument returns the Document field value if set, zero value otherwise.
func (o *GetTermsOfServiceDocumentResponse) GetDocument() string {
	if o == nil || IsNil(o.Document) {
		var ret string
		return ret
	}
	return *o.Document
}

// GetDocumentOk returns a tuple with the Document field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetTermsOfServiceDocumentResponse) GetDocumentOk() (*string, bool) {
	if o == nil || IsNil(o.Document) {
		return nil, false
	}
	return o.Document, true
}

// HasDocument returns a boolean if a field has been set.
func (o *GetTermsOfServiceDocumentResponse) HasDocument() bool {
	if o != nil && !IsNil(o.Document) {
		return true
	}

	return false
}

// SetDocument gets a reference to the given string and assigns it to the Document field.
func (o *GetTermsOfServiceDocumentResponse) SetDocument(v string) {
	o.Document = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *GetTermsOfServiceDocumentResponse) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetTermsOfServiceDocumentResponse) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *GetTermsOfServiceDocumentResponse) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *GetTermsOfServiceDocumentResponse) SetId(v string) {
	o.Id = &v
}

// GetLanguage returns the Language field value if set, zero value otherwise.
func (o *GetTermsOfServiceDocumentResponse) GetLanguage() string {
	if o == nil || IsNil(o.Language) {
		var ret string
		return ret
	}
	return *o.Language
}

// GetLanguageOk returns a tuple with the Language field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetTermsOfServiceDocumentResponse) GetLanguageOk() (*string, bool) {
	if o == nil || IsNil(o.Language) {
		return nil, false
	}
	return o.Language, true
}

// HasLanguage returns a boolean if a field has been set.
func (o *GetTermsOfServiceDocumentResponse) HasLanguage() bool {
	if o != nil && !IsNil(o.Language) {
		return true
	}

	return false
}

// SetLanguage gets a reference to the given string and assigns it to the Language field.
func (o *GetTermsOfServiceDocumentResponse) SetLanguage(v string) {
	o.Language = &v
}

// GetTermsOfServiceDocumentId returns the TermsOfServiceDocumentId field value if set, zero value otherwise.
func (o *GetTermsOfServiceDocumentResponse) GetTermsOfServiceDocumentId() string {
	if o == nil || IsNil(o.TermsOfServiceDocumentId) {
		var ret string
		return ret
	}
	return *o.TermsOfServiceDocumentId
}

// GetTermsOfServiceDocumentIdOk returns a tuple with the TermsOfServiceDocumentId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetTermsOfServiceDocumentResponse) GetTermsOfServiceDocumentIdOk() (*string, bool) {
	if o == nil || IsNil(o.TermsOfServiceDocumentId) {
		return nil, false
	}
	return o.TermsOfServiceDocumentId, true
}

// HasTermsOfServiceDocumentId returns a boolean if a field has been set.
func (o *GetTermsOfServiceDocumentResponse) HasTermsOfServiceDocumentId() bool {
	if o != nil && !IsNil(o.TermsOfServiceDocumentId) {
		return true
	}

	return false
}

// SetTermsOfServiceDocumentId gets a reference to the given string and assigns it to the TermsOfServiceDocumentId field.
func (o *GetTermsOfServiceDocumentResponse) SetTermsOfServiceDocumentId(v string) {
	o.TermsOfServiceDocumentId = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *GetTermsOfServiceDocumentResponse) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetTermsOfServiceDocumentResponse) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *GetTermsOfServiceDocumentResponse) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *GetTermsOfServiceDocumentResponse) SetType(v string) {
	o.Type = &v
}

func (o GetTermsOfServiceDocumentResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetTermsOfServiceDocumentResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Document) {
		toSerialize["document"] = o.Document
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Language) {
		toSerialize["language"] = o.Language
	}
	if !IsNil(o.TermsOfServiceDocumentId) {
		toSerialize["termsOfServiceDocumentId"] = o.TermsOfServiceDocumentId
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	return toSerialize, nil
}

type NullableGetTermsOfServiceDocumentResponse struct {
	value *GetTermsOfServiceDocumentResponse
	isSet bool
}

func (v NullableGetTermsOfServiceDocumentResponse) Get() *GetTermsOfServiceDocumentResponse {
	return v.value
}

func (v *NullableGetTermsOfServiceDocumentResponse) Set(val *GetTermsOfServiceDocumentResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGetTermsOfServiceDocumentResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGetTermsOfServiceDocumentResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetTermsOfServiceDocumentResponse(val *GetTermsOfServiceDocumentResponse) *NullableGetTermsOfServiceDocumentResponse {
	return &NullableGetTermsOfServiceDocumentResponse{value: val, isSet: true}
}

func (v NullableGetTermsOfServiceDocumentResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetTermsOfServiceDocumentResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


