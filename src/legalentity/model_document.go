/*
Legal Entity Management API

API version: 3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package legalentity

import (
	"encoding/json"
	"time"

	"github.com/adyen/adyen-go-api-library/v7/src/common"
)

// checks if the Document type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &Document{}

// Document struct for Document
type Document struct {
	Attachment *Attachment `json:"attachment,omitempty"`
	// Array that contains the document. The array supports multiple attachments for uploading different sides or pages of a document.
	Attachments []Attachment `json:"attachments"`
	// The creation date of the document.
	CreationDate *time.Time `json:"creationDate,omitempty"`
	// Your description for the document.
	Description string `json:"description"`
	// The expiry date of the document, in YYYY-MM-DD format.
	// Deprecated
	ExpiryDate *string `json:"expiryDate,omitempty"`
	// The filename of the document.
	FileName *string `json:"fileName,omitempty"`
	// The unique identifier of the document.
	Id *string `json:"id,omitempty"`
	// The two-character [ISO 3166-1 alpha-2](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2) country code where the document was issued. For example, **US**.
	// Deprecated
	IssuerCountry *string `json:"issuerCountry,omitempty"`
	// The state or province where the document was issued (AU only).
	// Deprecated
	IssuerState *string `json:"issuerState,omitempty"`
	// The modification date of the document.
	ModificationDate *time.Time `json:"modificationDate,omitempty"`
	// The number in the document.
	Number *string     `json:"number,omitempty"`
	Owner  OwnerEntity `json:"owner"`
	// Type of document, used when providing an ID number or uploading a document. The possible values depend on the legal entity type.  When providing ID numbers: * For **individual**, the `type` values can be **driversLicense**, **identityCard**, **nationalIdNumber**, or **passport**.  When uploading photo IDs: * For **individual**, the `type` values can be **identityCard**, **driversLicense**, or **passport**.  When uploading other documents: * For **organization**, the `type` values can be **proofOfAddress**, **registrationDocument**, **vatDocument**, **proofOfOrganizationTaxInfo**, **proofOfOwnership**, or **proofOfIndustry**.   * For **individual**, the `type` values can be **identityCard**, **driversLicense**, **passport**, **proofOfNationalIdNumber**, **proofOfResidency**, **proofOfIndustry**, or **proofOfIndividualTaxId**.  * For **soleProprietorship**, the `type` values can be **constitutionalDocument**, **proofOfAddress**, or **proofOfIndustry**.  * Use **bankStatement** to upload documents for a [transfer instrument](https://docs.adyen.com/api-explorer/#/legalentity/latest/post/transferInstruments__resParam_id).
	Type string `json:"type"`
}

// NewDocument instantiates a new Document object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDocument(attachments []Attachment, description string, owner OwnerEntity, type_ string) *Document {
	this := Document{}
	this.Attachments = attachments
	this.Description = description
	this.Owner = owner
	this.Type = type_
	return &this
}

// NewDocumentWithDefaults instantiates a new Document object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDocumentWithDefaults() *Document {
	this := Document{}
	return &this
}

// GetAttachment returns the Attachment field value if set, zero value otherwise.
func (o *Document) GetAttachment() Attachment {
	if o == nil || common.IsNil(o.Attachment) {
		var ret Attachment
		return ret
	}
	return *o.Attachment
}

// GetAttachmentOk returns a tuple with the Attachment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Document) GetAttachmentOk() (*Attachment, bool) {
	if o == nil || common.IsNil(o.Attachment) {
		return nil, false
	}
	return o.Attachment, true
}

// HasAttachment returns a boolean if a field has been set.
func (o *Document) HasAttachment() bool {
	if o != nil && !common.IsNil(o.Attachment) {
		return true
	}

	return false
}

// SetAttachment gets a reference to the given Attachment and assigns it to the Attachment field.
func (o *Document) SetAttachment(v Attachment) {
	o.Attachment = &v
}

// GetAttachments returns the Attachments field value
func (o *Document) GetAttachments() []Attachment {
	if o == nil {
		var ret []Attachment
		return ret
	}

	return o.Attachments
}

// GetAttachmentsOk returns a tuple with the Attachments field value
// and a boolean to check if the value has been set.
func (o *Document) GetAttachmentsOk() ([]Attachment, bool) {
	if o == nil {
		return nil, false
	}
	return o.Attachments, true
}

// SetAttachments sets field value
func (o *Document) SetAttachments(v []Attachment) {
	o.Attachments = v
}

// GetCreationDate returns the CreationDate field value if set, zero value otherwise.
func (o *Document) GetCreationDate() time.Time {
	if o == nil || common.IsNil(o.CreationDate) {
		var ret time.Time
		return ret
	}
	return *o.CreationDate
}

// GetCreationDateOk returns a tuple with the CreationDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Document) GetCreationDateOk() (*time.Time, bool) {
	if o == nil || common.IsNil(o.CreationDate) {
		return nil, false
	}
	return o.CreationDate, true
}

// HasCreationDate returns a boolean if a field has been set.
func (o *Document) HasCreationDate() bool {
	if o != nil && !common.IsNil(o.CreationDate) {
		return true
	}

	return false
}

// SetCreationDate gets a reference to the given time.Time and assigns it to the CreationDate field.
func (o *Document) SetCreationDate(v time.Time) {
	o.CreationDate = &v
}

// GetDescription returns the Description field value
func (o *Document) GetDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *Document) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Description, true
}

// SetDescription sets field value
func (o *Document) SetDescription(v string) {
	o.Description = v
}

// GetExpiryDate returns the ExpiryDate field value if set, zero value otherwise.
// Deprecated
func (o *Document) GetExpiryDate() string {
	if o == nil || common.IsNil(o.ExpiryDate) {
		var ret string
		return ret
	}
	return *o.ExpiryDate
}

// GetExpiryDateOk returns a tuple with the ExpiryDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
// Deprecated
func (o *Document) GetExpiryDateOk() (*string, bool) {
	if o == nil || common.IsNil(o.ExpiryDate) {
		return nil, false
	}
	return o.ExpiryDate, true
}

// HasExpiryDate returns a boolean if a field has been set.
func (o *Document) HasExpiryDate() bool {
	if o != nil && !common.IsNil(o.ExpiryDate) {
		return true
	}

	return false
}

// SetExpiryDate gets a reference to the given string and assigns it to the ExpiryDate field.
// Deprecated
func (o *Document) SetExpiryDate(v string) {
	o.ExpiryDate = &v
}

// GetFileName returns the FileName field value if set, zero value otherwise.
func (o *Document) GetFileName() string {
	if o == nil || common.IsNil(o.FileName) {
		var ret string
		return ret
	}
	return *o.FileName
}

// GetFileNameOk returns a tuple with the FileName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Document) GetFileNameOk() (*string, bool) {
	if o == nil || common.IsNil(o.FileName) {
		return nil, false
	}
	return o.FileName, true
}

// HasFileName returns a boolean if a field has been set.
func (o *Document) HasFileName() bool {
	if o != nil && !common.IsNil(o.FileName) {
		return true
	}

	return false
}

// SetFileName gets a reference to the given string and assigns it to the FileName field.
func (o *Document) SetFileName(v string) {
	o.FileName = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *Document) GetId() string {
	if o == nil || common.IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Document) GetIdOk() (*string, bool) {
	if o == nil || common.IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *Document) HasId() bool {
	if o != nil && !common.IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *Document) SetId(v string) {
	o.Id = &v
}

// GetIssuerCountry returns the IssuerCountry field value if set, zero value otherwise.
// Deprecated
func (o *Document) GetIssuerCountry() string {
	if o == nil || common.IsNil(o.IssuerCountry) {
		var ret string
		return ret
	}
	return *o.IssuerCountry
}

// GetIssuerCountryOk returns a tuple with the IssuerCountry field value if set, nil otherwise
// and a boolean to check if the value has been set.
// Deprecated
func (o *Document) GetIssuerCountryOk() (*string, bool) {
	if o == nil || common.IsNil(o.IssuerCountry) {
		return nil, false
	}
	return o.IssuerCountry, true
}

// HasIssuerCountry returns a boolean if a field has been set.
func (o *Document) HasIssuerCountry() bool {
	if o != nil && !common.IsNil(o.IssuerCountry) {
		return true
	}

	return false
}

// SetIssuerCountry gets a reference to the given string and assigns it to the IssuerCountry field.
// Deprecated
func (o *Document) SetIssuerCountry(v string) {
	o.IssuerCountry = &v
}

// GetIssuerState returns the IssuerState field value if set, zero value otherwise.
// Deprecated
func (o *Document) GetIssuerState() string {
	if o == nil || common.IsNil(o.IssuerState) {
		var ret string
		return ret
	}
	return *o.IssuerState
}

// GetIssuerStateOk returns a tuple with the IssuerState field value if set, nil otherwise
// and a boolean to check if the value has been set.
// Deprecated
func (o *Document) GetIssuerStateOk() (*string, bool) {
	if o == nil || common.IsNil(o.IssuerState) {
		return nil, false
	}
	return o.IssuerState, true
}

// HasIssuerState returns a boolean if a field has been set.
func (o *Document) HasIssuerState() bool {
	if o != nil && !common.IsNil(o.IssuerState) {
		return true
	}

	return false
}

// SetIssuerState gets a reference to the given string and assigns it to the IssuerState field.
// Deprecated
func (o *Document) SetIssuerState(v string) {
	o.IssuerState = &v
}

// GetModificationDate returns the ModificationDate field value if set, zero value otherwise.
func (o *Document) GetModificationDate() time.Time {
	if o == nil || common.IsNil(o.ModificationDate) {
		var ret time.Time
		return ret
	}
	return *o.ModificationDate
}

// GetModificationDateOk returns a tuple with the ModificationDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Document) GetModificationDateOk() (*time.Time, bool) {
	if o == nil || common.IsNil(o.ModificationDate) {
		return nil, false
	}
	return o.ModificationDate, true
}

// HasModificationDate returns a boolean if a field has been set.
func (o *Document) HasModificationDate() bool {
	if o != nil && !common.IsNil(o.ModificationDate) {
		return true
	}

	return false
}

// SetModificationDate gets a reference to the given time.Time and assigns it to the ModificationDate field.
func (o *Document) SetModificationDate(v time.Time) {
	o.ModificationDate = &v
}

// GetNumber returns the Number field value if set, zero value otherwise.
func (o *Document) GetNumber() string {
	if o == nil || common.IsNil(o.Number) {
		var ret string
		return ret
	}
	return *o.Number
}

// GetNumberOk returns a tuple with the Number field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Document) GetNumberOk() (*string, bool) {
	if o == nil || common.IsNil(o.Number) {
		return nil, false
	}
	return o.Number, true
}

// HasNumber returns a boolean if a field has been set.
func (o *Document) HasNumber() bool {
	if o != nil && !common.IsNil(o.Number) {
		return true
	}

	return false
}

// SetNumber gets a reference to the given string and assigns it to the Number field.
func (o *Document) SetNumber(v string) {
	o.Number = &v
}

// GetOwner returns the Owner field value
func (o *Document) GetOwner() OwnerEntity {
	if o == nil {
		var ret OwnerEntity
		return ret
	}

	return o.Owner
}

// GetOwnerOk returns a tuple with the Owner field value
// and a boolean to check if the value has been set.
func (o *Document) GetOwnerOk() (*OwnerEntity, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Owner, true
}

// SetOwner sets field value
func (o *Document) SetOwner(v OwnerEntity) {
	o.Owner = v
}

// GetType returns the Type field value
func (o *Document) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *Document) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *Document) SetType(v string) {
	o.Type = v
}

func (o Document) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Document) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.Attachment) {
		toSerialize["attachment"] = o.Attachment
	}
	toSerialize["attachments"] = o.Attachments
	if !common.IsNil(o.CreationDate) {
		toSerialize["creationDate"] = o.CreationDate
	}
	toSerialize["description"] = o.Description
	if !common.IsNil(o.ExpiryDate) {
		toSerialize["expiryDate"] = o.ExpiryDate
	}
	if !common.IsNil(o.FileName) {
		toSerialize["fileName"] = o.FileName
	}
	if !common.IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !common.IsNil(o.IssuerCountry) {
		toSerialize["issuerCountry"] = o.IssuerCountry
	}
	if !common.IsNil(o.IssuerState) {
		toSerialize["issuerState"] = o.IssuerState
	}
	if !common.IsNil(o.ModificationDate) {
		toSerialize["modificationDate"] = o.ModificationDate
	}
	if !common.IsNil(o.Number) {
		toSerialize["number"] = o.Number
	}
	toSerialize["owner"] = o.Owner
	toSerialize["type"] = o.Type
	return toSerialize, nil
}

type NullableDocument struct {
	value *Document
	isSet bool
}

func (v NullableDocument) Get() *Document {
	return v.value
}

func (v *NullableDocument) Set(val *Document) {
	v.value = val
	v.isSet = true
}

func (v NullableDocument) IsSet() bool {
	return v.isSet
}

func (v *NullableDocument) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDocument(val *Document) *NullableDocument {
	return &NullableDocument{value: val, isSet: true}
}

func (v NullableDocument) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDocument) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

func (o *Document) isValidType() bool {
	var allowedEnumValues = []string{"bankStatement", "driversLicense", "identityCard", "nationalIdNumber", "passport", "proofOfAddress", "proofOfNationalIdNumber", "proofOfResidency", "registrationDocument", "vatDocument", "proofOfOrganizationTaxInfo", "proofOfIndustry", "constitutionalDocument"}
	for _, allowed := range allowedEnumValues {
		if o.GetType() == allowed {
			return true
		}
	}
	return false
}
