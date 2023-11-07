/*
Adyen Payment API

API version: 68
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package payments

import (
	"encoding/json"

	"github.com/adyen/adyen-go-api-library/v8/src/common"
)

// checks if the AdditionalDataCommon type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &AdditionalDataCommon{}

// AdditionalDataCommon struct for AdditionalDataCommon
type AdditionalDataCommon struct {
	// Triggers test scenarios that allow to replicate certain communication errors.  Allowed values: * **NO_CONNECTION_AVAILABLE** – There wasn't a connection available to service the outgoing communication. This is a transient, retriable error since no messaging could be initiated to an issuing system (or third-party acquiring system). Therefore, the header Transient-Error: true is returned in the response. A subsequent request using the same idempotency key will be processed as if it was the first request. * **IOEXCEPTION_RECEIVED** – Something went wrong during transmission of the message or receiving the response. This is a classified as non-transient because the message could have been received by the issuing party and been acted upon. No transient error header is returned. If using idempotency, the (error) response is stored as the final result for the idempotency key. Subsequent messages with the same idempotency key not be processed beyond returning the stored response.
	RequestedTestErrorResponseCode *string `json:"RequestedTestErrorResponseCode,omitempty"`
	// Set to true to authorise a part of the requested amount in case the cardholder does not have enough funds on their account.  If a payment was partially authorised, the response includes resultCode: PartiallyAuthorised and the authorised amount in additionalData.authorisedAmountValue. To enable this functionality, contact our Support Team.
	AllowPartialAuth *string `json:"allowPartialAuth,omitempty"`
	// Flags a card payment request for either pre-authorisation or final authorisation. For more information, refer to [Authorisation types](https://docs.adyen.com/online-payments/adjust-authorisation#authorisation-types).  Allowed values: * **PreAuth** – flags the payment request to be handled as a pre-authorisation. * **FinalAuth** – flags the payment request to be handled as a final authorisation.
	AuthorisationType *string `json:"authorisationType,omitempty"`
	// Allows you to determine or override the acquirer account that should be used for the transaction.  If you need to process a payment with an acquirer different from a default one, you can set up a corresponding configuration on the Adyen payments platform. Then you can pass a custom routing flag in a payment request's additional data to target a specific acquirer.  To enable this functionality, contact [Support](https://www.adyen.help/hc/en-us/requests/new).
	CustomRoutingFlag *string `json:"customRoutingFlag,omitempty"`
	// In case of [asynchronous authorisation adjustment](https://docs.adyen.com/online-payments/adjust-authorisation#adjust-authorisation), this field denotes why the additional payment is made.  Possible values:   * **NoShow**: An incremental charge is carried out because of a no-show for a guaranteed reservation.   * **DelayedCharge**: An incremental charge is carried out to process an additional payment after the original services have been rendered and the respective payment has been processed.
	IndustryUsage *string `json:"industryUsage,omitempty"`
	// Set to **true** to require [manual capture](https://docs.adyen.com/online-payments/capture) for the transaction.
	ManualCapture *string `json:"manualCapture,omitempty"`
	// Allows you to link the transaction to the original or previous one in a subscription/card-on-file chain. This field is required for token-based transactions where Adyen does not tokenize the card.  Transaction identifier from card schemes, for example, Mastercard Trace ID or the Visa Transaction ID.  Submit the original transaction ID of the contract in your payment request if you are not tokenizing card details with Adyen and are making a merchant-initiated transaction (MIT) for subsequent charges.  Make sure you are sending `shopperInteraction` **ContAuth** and `recurringProcessingModel` **Subscription** or **UnscheduledCardOnFile** to ensure that the transaction is classified as MIT.
	NetworkTxReference *string `json:"networkTxReference,omitempty"`
	// Boolean indicator that can be optionally used for performing debit transactions on combo cards (for example, combo cards in Brazil). This is not mandatory but we recommend that you set this to true if you want to use the `selectedBrand` value to specify how to process the transaction.
	OverwriteBrand *string `json:"overwriteBrand,omitempty"`
	// This field is required if the transaction is performed by a registered payment facilitator. This field must contain the city of the actual merchant's address. * Format: alpha-numeric. * Maximum length: 13 characters.
	SubMerchantCity *string `json:"subMerchantCity,omitempty"`
	// This field is required if the transaction is performed by a registered payment facilitator. This field must contain the three-letter country code of the actual merchant's address. * Format: alpha-numeric. * Fixed length: 3 characters.
	SubMerchantCountry *string `json:"subMerchantCountry,omitempty"`
	// This field contains an identifier of the actual merchant when a transaction is submitted via a payment facilitator. The payment facilitator must send in this unique ID.  A unique identifier per submerchant that is required if the transaction is performed by a registered payment facilitator. * Format: alpha-numeric. * Fixed length: 15 characters.
	SubMerchantID *string `json:"subMerchantID,omitempty"`
	// This field is required if the transaction is performed by a registered payment facilitator. This field must contain the name of the actual merchant. * Format: alpha-numeric. * Maximum length: 22 characters.
	SubMerchantName *string `json:"subMerchantName,omitempty"`
	// This field is required if the transaction is performed by a registered payment facilitator. This field must contain the postal code of the actual merchant's address. * Format: alpha-numeric. * Maximum length: 10 characters.
	SubMerchantPostalCode *string `json:"subMerchantPostalCode,omitempty"`
	// This field is required if the transaction is performed by a registered payment facilitator, and if applicable to the country. This field must contain the state code of the actual merchant's address. * Format: alpha-numeric. * Maximum length: 3 characters.
	SubMerchantState *string `json:"subMerchantState,omitempty"`
	// This field is required if the transaction is performed by a registered payment facilitator. This field must contain the street of the actual merchant's address. * Format: alpha-numeric. * Maximum length: 60 characters.
	SubMerchantStreet *string `json:"subMerchantStreet,omitempty"`
	// This field is required if the transaction is performed by a registered payment facilitator. This field must contain the tax ID of the actual merchant. * Format: alpha-numeric. * Fixed length: 11 or 14 characters.
	SubMerchantTaxId *string `json:"subMerchantTaxId,omitempty"`
}

// NewAdditionalDataCommon instantiates a new AdditionalDataCommon object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAdditionalDataCommon() *AdditionalDataCommon {
	this := AdditionalDataCommon{}
	return &this
}

// NewAdditionalDataCommonWithDefaults instantiates a new AdditionalDataCommon object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAdditionalDataCommonWithDefaults() *AdditionalDataCommon {
	this := AdditionalDataCommon{}
	return &this
}

// GetRequestedTestErrorResponseCode returns the RequestedTestErrorResponseCode field value if set, zero value otherwise.
func (o *AdditionalDataCommon) GetRequestedTestErrorResponseCode() string {
	if o == nil || common.IsNil(o.RequestedTestErrorResponseCode) {
		var ret string
		return ret
	}
	return *o.RequestedTestErrorResponseCode
}

// GetRequestedTestErrorResponseCodeOk returns a tuple with the RequestedTestErrorResponseCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdditionalDataCommon) GetRequestedTestErrorResponseCodeOk() (*string, bool) {
	if o == nil || common.IsNil(o.RequestedTestErrorResponseCode) {
		return nil, false
	}
	return o.RequestedTestErrorResponseCode, true
}

// HasRequestedTestErrorResponseCode returns a boolean if a field has been set.
func (o *AdditionalDataCommon) HasRequestedTestErrorResponseCode() bool {
	if o != nil && !common.IsNil(o.RequestedTestErrorResponseCode) {
		return true
	}

	return false
}

// SetRequestedTestErrorResponseCode gets a reference to the given string and assigns it to the RequestedTestErrorResponseCode field.
func (o *AdditionalDataCommon) SetRequestedTestErrorResponseCode(v string) {
	o.RequestedTestErrorResponseCode = &v
}

// GetAllowPartialAuth returns the AllowPartialAuth field value if set, zero value otherwise.
func (o *AdditionalDataCommon) GetAllowPartialAuth() string {
	if o == nil || common.IsNil(o.AllowPartialAuth) {
		var ret string
		return ret
	}
	return *o.AllowPartialAuth
}

// GetAllowPartialAuthOk returns a tuple with the AllowPartialAuth field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdditionalDataCommon) GetAllowPartialAuthOk() (*string, bool) {
	if o == nil || common.IsNil(o.AllowPartialAuth) {
		return nil, false
	}
	return o.AllowPartialAuth, true
}

// HasAllowPartialAuth returns a boolean if a field has been set.
func (o *AdditionalDataCommon) HasAllowPartialAuth() bool {
	if o != nil && !common.IsNil(o.AllowPartialAuth) {
		return true
	}

	return false
}

// SetAllowPartialAuth gets a reference to the given string and assigns it to the AllowPartialAuth field.
func (o *AdditionalDataCommon) SetAllowPartialAuth(v string) {
	o.AllowPartialAuth = &v
}

// GetAuthorisationType returns the AuthorisationType field value if set, zero value otherwise.
func (o *AdditionalDataCommon) GetAuthorisationType() string {
	if o == nil || common.IsNil(o.AuthorisationType) {
		var ret string
		return ret
	}
	return *o.AuthorisationType
}

// GetAuthorisationTypeOk returns a tuple with the AuthorisationType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdditionalDataCommon) GetAuthorisationTypeOk() (*string, bool) {
	if o == nil || common.IsNil(o.AuthorisationType) {
		return nil, false
	}
	return o.AuthorisationType, true
}

// HasAuthorisationType returns a boolean if a field has been set.
func (o *AdditionalDataCommon) HasAuthorisationType() bool {
	if o != nil && !common.IsNil(o.AuthorisationType) {
		return true
	}

	return false
}

// SetAuthorisationType gets a reference to the given string and assigns it to the AuthorisationType field.
func (o *AdditionalDataCommon) SetAuthorisationType(v string) {
	o.AuthorisationType = &v
}

// GetCustomRoutingFlag returns the CustomRoutingFlag field value if set, zero value otherwise.
func (o *AdditionalDataCommon) GetCustomRoutingFlag() string {
	if o == nil || common.IsNil(o.CustomRoutingFlag) {
		var ret string
		return ret
	}
	return *o.CustomRoutingFlag
}

// GetCustomRoutingFlagOk returns a tuple with the CustomRoutingFlag field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdditionalDataCommon) GetCustomRoutingFlagOk() (*string, bool) {
	if o == nil || common.IsNil(o.CustomRoutingFlag) {
		return nil, false
	}
	return o.CustomRoutingFlag, true
}

// HasCustomRoutingFlag returns a boolean if a field has been set.
func (o *AdditionalDataCommon) HasCustomRoutingFlag() bool {
	if o != nil && !common.IsNil(o.CustomRoutingFlag) {
		return true
	}

	return false
}

// SetCustomRoutingFlag gets a reference to the given string and assigns it to the CustomRoutingFlag field.
func (o *AdditionalDataCommon) SetCustomRoutingFlag(v string) {
	o.CustomRoutingFlag = &v
}

// GetIndustryUsage returns the IndustryUsage field value if set, zero value otherwise.
func (o *AdditionalDataCommon) GetIndustryUsage() string {
	if o == nil || common.IsNil(o.IndustryUsage) {
		var ret string
		return ret
	}
	return *o.IndustryUsage
}

// GetIndustryUsageOk returns a tuple with the IndustryUsage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdditionalDataCommon) GetIndustryUsageOk() (*string, bool) {
	if o == nil || common.IsNil(o.IndustryUsage) {
		return nil, false
	}
	return o.IndustryUsage, true
}

// HasIndustryUsage returns a boolean if a field has been set.
func (o *AdditionalDataCommon) HasIndustryUsage() bool {
	if o != nil && !common.IsNil(o.IndustryUsage) {
		return true
	}

	return false
}

// SetIndustryUsage gets a reference to the given string and assigns it to the IndustryUsage field.
func (o *AdditionalDataCommon) SetIndustryUsage(v string) {
	o.IndustryUsage = &v
}

// GetManualCapture returns the ManualCapture field value if set, zero value otherwise.
func (o *AdditionalDataCommon) GetManualCapture() string {
	if o == nil || common.IsNil(o.ManualCapture) {
		var ret string
		return ret
	}
	return *o.ManualCapture
}

// GetManualCaptureOk returns a tuple with the ManualCapture field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdditionalDataCommon) GetManualCaptureOk() (*string, bool) {
	if o == nil || common.IsNil(o.ManualCapture) {
		return nil, false
	}
	return o.ManualCapture, true
}

// HasManualCapture returns a boolean if a field has been set.
func (o *AdditionalDataCommon) HasManualCapture() bool {
	if o != nil && !common.IsNil(o.ManualCapture) {
		return true
	}

	return false
}

// SetManualCapture gets a reference to the given string and assigns it to the ManualCapture field.
func (o *AdditionalDataCommon) SetManualCapture(v string) {
	o.ManualCapture = &v
}

// GetNetworkTxReference returns the NetworkTxReference field value if set, zero value otherwise.
func (o *AdditionalDataCommon) GetNetworkTxReference() string {
	if o == nil || common.IsNil(o.NetworkTxReference) {
		var ret string
		return ret
	}
	return *o.NetworkTxReference
}

// GetNetworkTxReferenceOk returns a tuple with the NetworkTxReference field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdditionalDataCommon) GetNetworkTxReferenceOk() (*string, bool) {
	if o == nil || common.IsNil(o.NetworkTxReference) {
		return nil, false
	}
	return o.NetworkTxReference, true
}

// HasNetworkTxReference returns a boolean if a field has been set.
func (o *AdditionalDataCommon) HasNetworkTxReference() bool {
	if o != nil && !common.IsNil(o.NetworkTxReference) {
		return true
	}

	return false
}

// SetNetworkTxReference gets a reference to the given string and assigns it to the NetworkTxReference field.
func (o *AdditionalDataCommon) SetNetworkTxReference(v string) {
	o.NetworkTxReference = &v
}

// GetOverwriteBrand returns the OverwriteBrand field value if set, zero value otherwise.
func (o *AdditionalDataCommon) GetOverwriteBrand() string {
	if o == nil || common.IsNil(o.OverwriteBrand) {
		var ret string
		return ret
	}
	return *o.OverwriteBrand
}

// GetOverwriteBrandOk returns a tuple with the OverwriteBrand field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdditionalDataCommon) GetOverwriteBrandOk() (*string, bool) {
	if o == nil || common.IsNil(o.OverwriteBrand) {
		return nil, false
	}
	return o.OverwriteBrand, true
}

// HasOverwriteBrand returns a boolean if a field has been set.
func (o *AdditionalDataCommon) HasOverwriteBrand() bool {
	if o != nil && !common.IsNil(o.OverwriteBrand) {
		return true
	}

	return false
}

// SetOverwriteBrand gets a reference to the given string and assigns it to the OverwriteBrand field.
func (o *AdditionalDataCommon) SetOverwriteBrand(v string) {
	o.OverwriteBrand = &v
}

// GetSubMerchantCity returns the SubMerchantCity field value if set, zero value otherwise.
func (o *AdditionalDataCommon) GetSubMerchantCity() string {
	if o == nil || common.IsNil(o.SubMerchantCity) {
		var ret string
		return ret
	}
	return *o.SubMerchantCity
}

// GetSubMerchantCityOk returns a tuple with the SubMerchantCity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdditionalDataCommon) GetSubMerchantCityOk() (*string, bool) {
	if o == nil || common.IsNil(o.SubMerchantCity) {
		return nil, false
	}
	return o.SubMerchantCity, true
}

// HasSubMerchantCity returns a boolean if a field has been set.
func (o *AdditionalDataCommon) HasSubMerchantCity() bool {
	if o != nil && !common.IsNil(o.SubMerchantCity) {
		return true
	}

	return false
}

// SetSubMerchantCity gets a reference to the given string and assigns it to the SubMerchantCity field.
func (o *AdditionalDataCommon) SetSubMerchantCity(v string) {
	o.SubMerchantCity = &v
}

// GetSubMerchantCountry returns the SubMerchantCountry field value if set, zero value otherwise.
func (o *AdditionalDataCommon) GetSubMerchantCountry() string {
	if o == nil || common.IsNil(o.SubMerchantCountry) {
		var ret string
		return ret
	}
	return *o.SubMerchantCountry
}

// GetSubMerchantCountryOk returns a tuple with the SubMerchantCountry field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdditionalDataCommon) GetSubMerchantCountryOk() (*string, bool) {
	if o == nil || common.IsNil(o.SubMerchantCountry) {
		return nil, false
	}
	return o.SubMerchantCountry, true
}

// HasSubMerchantCountry returns a boolean if a field has been set.
func (o *AdditionalDataCommon) HasSubMerchantCountry() bool {
	if o != nil && !common.IsNil(o.SubMerchantCountry) {
		return true
	}

	return false
}

// SetSubMerchantCountry gets a reference to the given string and assigns it to the SubMerchantCountry field.
func (o *AdditionalDataCommon) SetSubMerchantCountry(v string) {
	o.SubMerchantCountry = &v
}

// GetSubMerchantID returns the SubMerchantID field value if set, zero value otherwise.
func (o *AdditionalDataCommon) GetSubMerchantID() string {
	if o == nil || common.IsNil(o.SubMerchantID) {
		var ret string
		return ret
	}
	return *o.SubMerchantID
}

// GetSubMerchantIDOk returns a tuple with the SubMerchantID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdditionalDataCommon) GetSubMerchantIDOk() (*string, bool) {
	if o == nil || common.IsNil(o.SubMerchantID) {
		return nil, false
	}
	return o.SubMerchantID, true
}

// HasSubMerchantID returns a boolean if a field has been set.
func (o *AdditionalDataCommon) HasSubMerchantID() bool {
	if o != nil && !common.IsNil(o.SubMerchantID) {
		return true
	}

	return false
}

// SetSubMerchantID gets a reference to the given string and assigns it to the SubMerchantID field.
func (o *AdditionalDataCommon) SetSubMerchantID(v string) {
	o.SubMerchantID = &v
}

// GetSubMerchantName returns the SubMerchantName field value if set, zero value otherwise.
func (o *AdditionalDataCommon) GetSubMerchantName() string {
	if o == nil || common.IsNil(o.SubMerchantName) {
		var ret string
		return ret
	}
	return *o.SubMerchantName
}

// GetSubMerchantNameOk returns a tuple with the SubMerchantName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdditionalDataCommon) GetSubMerchantNameOk() (*string, bool) {
	if o == nil || common.IsNil(o.SubMerchantName) {
		return nil, false
	}
	return o.SubMerchantName, true
}

// HasSubMerchantName returns a boolean if a field has been set.
func (o *AdditionalDataCommon) HasSubMerchantName() bool {
	if o != nil && !common.IsNil(o.SubMerchantName) {
		return true
	}

	return false
}

// SetSubMerchantName gets a reference to the given string and assigns it to the SubMerchantName field.
func (o *AdditionalDataCommon) SetSubMerchantName(v string) {
	o.SubMerchantName = &v
}

// GetSubMerchantPostalCode returns the SubMerchantPostalCode field value if set, zero value otherwise.
func (o *AdditionalDataCommon) GetSubMerchantPostalCode() string {
	if o == nil || common.IsNil(o.SubMerchantPostalCode) {
		var ret string
		return ret
	}
	return *o.SubMerchantPostalCode
}

// GetSubMerchantPostalCodeOk returns a tuple with the SubMerchantPostalCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdditionalDataCommon) GetSubMerchantPostalCodeOk() (*string, bool) {
	if o == nil || common.IsNil(o.SubMerchantPostalCode) {
		return nil, false
	}
	return o.SubMerchantPostalCode, true
}

// HasSubMerchantPostalCode returns a boolean if a field has been set.
func (o *AdditionalDataCommon) HasSubMerchantPostalCode() bool {
	if o != nil && !common.IsNil(o.SubMerchantPostalCode) {
		return true
	}

	return false
}

// SetSubMerchantPostalCode gets a reference to the given string and assigns it to the SubMerchantPostalCode field.
func (o *AdditionalDataCommon) SetSubMerchantPostalCode(v string) {
	o.SubMerchantPostalCode = &v
}

// GetSubMerchantState returns the SubMerchantState field value if set, zero value otherwise.
func (o *AdditionalDataCommon) GetSubMerchantState() string {
	if o == nil || common.IsNil(o.SubMerchantState) {
		var ret string
		return ret
	}
	return *o.SubMerchantState
}

// GetSubMerchantStateOk returns a tuple with the SubMerchantState field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdditionalDataCommon) GetSubMerchantStateOk() (*string, bool) {
	if o == nil || common.IsNil(o.SubMerchantState) {
		return nil, false
	}
	return o.SubMerchantState, true
}

// HasSubMerchantState returns a boolean if a field has been set.
func (o *AdditionalDataCommon) HasSubMerchantState() bool {
	if o != nil && !common.IsNil(o.SubMerchantState) {
		return true
	}

	return false
}

// SetSubMerchantState gets a reference to the given string and assigns it to the SubMerchantState field.
func (o *AdditionalDataCommon) SetSubMerchantState(v string) {
	o.SubMerchantState = &v
}

// GetSubMerchantStreet returns the SubMerchantStreet field value if set, zero value otherwise.
func (o *AdditionalDataCommon) GetSubMerchantStreet() string {
	if o == nil || common.IsNil(o.SubMerchantStreet) {
		var ret string
		return ret
	}
	return *o.SubMerchantStreet
}

// GetSubMerchantStreetOk returns a tuple with the SubMerchantStreet field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdditionalDataCommon) GetSubMerchantStreetOk() (*string, bool) {
	if o == nil || common.IsNil(o.SubMerchantStreet) {
		return nil, false
	}
	return o.SubMerchantStreet, true
}

// HasSubMerchantStreet returns a boolean if a field has been set.
func (o *AdditionalDataCommon) HasSubMerchantStreet() bool {
	if o != nil && !common.IsNil(o.SubMerchantStreet) {
		return true
	}

	return false
}

// SetSubMerchantStreet gets a reference to the given string and assigns it to the SubMerchantStreet field.
func (o *AdditionalDataCommon) SetSubMerchantStreet(v string) {
	o.SubMerchantStreet = &v
}

// GetSubMerchantTaxId returns the SubMerchantTaxId field value if set, zero value otherwise.
func (o *AdditionalDataCommon) GetSubMerchantTaxId() string {
	if o == nil || common.IsNil(o.SubMerchantTaxId) {
		var ret string
		return ret
	}
	return *o.SubMerchantTaxId
}

// GetSubMerchantTaxIdOk returns a tuple with the SubMerchantTaxId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdditionalDataCommon) GetSubMerchantTaxIdOk() (*string, bool) {
	if o == nil || common.IsNil(o.SubMerchantTaxId) {
		return nil, false
	}
	return o.SubMerchantTaxId, true
}

// HasSubMerchantTaxId returns a boolean if a field has been set.
func (o *AdditionalDataCommon) HasSubMerchantTaxId() bool {
	if o != nil && !common.IsNil(o.SubMerchantTaxId) {
		return true
	}

	return false
}

// SetSubMerchantTaxId gets a reference to the given string and assigns it to the SubMerchantTaxId field.
func (o *AdditionalDataCommon) SetSubMerchantTaxId(v string) {
	o.SubMerchantTaxId = &v
}

func (o AdditionalDataCommon) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AdditionalDataCommon) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.RequestedTestErrorResponseCode) {
		toSerialize["RequestedTestErrorResponseCode"] = o.RequestedTestErrorResponseCode
	}
	if !common.IsNil(o.AllowPartialAuth) {
		toSerialize["allowPartialAuth"] = o.AllowPartialAuth
	}
	if !common.IsNil(o.AuthorisationType) {
		toSerialize["authorisationType"] = o.AuthorisationType
	}
	if !common.IsNil(o.CustomRoutingFlag) {
		toSerialize["customRoutingFlag"] = o.CustomRoutingFlag
	}
	if !common.IsNil(o.IndustryUsage) {
		toSerialize["industryUsage"] = o.IndustryUsage
	}
	if !common.IsNil(o.ManualCapture) {
		toSerialize["manualCapture"] = o.ManualCapture
	}
	if !common.IsNil(o.NetworkTxReference) {
		toSerialize["networkTxReference"] = o.NetworkTxReference
	}
	if !common.IsNil(o.OverwriteBrand) {
		toSerialize["overwriteBrand"] = o.OverwriteBrand
	}
	if !common.IsNil(o.SubMerchantCity) {
		toSerialize["subMerchantCity"] = o.SubMerchantCity
	}
	if !common.IsNil(o.SubMerchantCountry) {
		toSerialize["subMerchantCountry"] = o.SubMerchantCountry
	}
	if !common.IsNil(o.SubMerchantID) {
		toSerialize["subMerchantID"] = o.SubMerchantID
	}
	if !common.IsNil(o.SubMerchantName) {
		toSerialize["subMerchantName"] = o.SubMerchantName
	}
	if !common.IsNil(o.SubMerchantPostalCode) {
		toSerialize["subMerchantPostalCode"] = o.SubMerchantPostalCode
	}
	if !common.IsNil(o.SubMerchantState) {
		toSerialize["subMerchantState"] = o.SubMerchantState
	}
	if !common.IsNil(o.SubMerchantStreet) {
		toSerialize["subMerchantStreet"] = o.SubMerchantStreet
	}
	if !common.IsNil(o.SubMerchantTaxId) {
		toSerialize["subMerchantTaxId"] = o.SubMerchantTaxId
	}
	return toSerialize, nil
}

type NullableAdditionalDataCommon struct {
	value *AdditionalDataCommon
	isSet bool
}

func (v NullableAdditionalDataCommon) Get() *AdditionalDataCommon {
	return v.value
}

func (v *NullableAdditionalDataCommon) Set(val *AdditionalDataCommon) {
	v.value = val
	v.isSet = true
}

func (v NullableAdditionalDataCommon) IsSet() bool {
	return v.isSet
}

func (v *NullableAdditionalDataCommon) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAdditionalDataCommon(val *AdditionalDataCommon) *NullableAdditionalDataCommon {
	return &NullableAdditionalDataCommon{value: val, isSet: true}
}

func (v NullableAdditionalDataCommon) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAdditionalDataCommon) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

func (o *AdditionalDataCommon) isValidIndustryUsage() bool {
	var allowedEnumValues = []string{"NoShow", "DelayedCharge"}
	for _, allowed := range allowedEnumValues {
		if o.GetIndustryUsage() == allowed {
			return true
		}
	}
	return false
}
