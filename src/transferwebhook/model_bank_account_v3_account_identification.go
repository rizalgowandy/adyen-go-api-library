/*
Transfer webhooks

API version: 3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package transferwebhook

import (
	"encoding/json"
	"fmt"
)

// BankAccountV3AccountIdentification - Contains the bank account details. The fields required in this object depend on the country of the bank account and the currency of the transfer.
type BankAccountV3AccountIdentification struct {
	AULocalAccountIdentification      *AULocalAccountIdentification
	BRLocalAccountIdentification      *BRLocalAccountIdentification
	CALocalAccountIdentification      *CALocalAccountIdentification
	CZLocalAccountIdentification      *CZLocalAccountIdentification
	DKLocalAccountIdentification      *DKLocalAccountIdentification
	HULocalAccountIdentification      *HULocalAccountIdentification
	IbanAccountIdentification         *IbanAccountIdentification
	NOLocalAccountIdentification      *NOLocalAccountIdentification
	NumberAndBicAccountIdentification *NumberAndBicAccountIdentification
	PLLocalAccountIdentification      *PLLocalAccountIdentification
	SELocalAccountIdentification      *SELocalAccountIdentification
	SGLocalAccountIdentification      *SGLocalAccountIdentification
	UKLocalAccountIdentification      *UKLocalAccountIdentification
	USLocalAccountIdentification      *USLocalAccountIdentification
}

// AULocalAccountIdentificationAsBankAccountV3AccountIdentification is a convenience function that returns AULocalAccountIdentification wrapped in BankAccountV3AccountIdentification
func AULocalAccountIdentificationAsBankAccountV3AccountIdentification(v *AULocalAccountIdentification) BankAccountV3AccountIdentification {
	return BankAccountV3AccountIdentification{
		AULocalAccountIdentification: v,
	}
}

// BRLocalAccountIdentificationAsBankAccountV3AccountIdentification is a convenience function that returns BRLocalAccountIdentification wrapped in BankAccountV3AccountIdentification
func BRLocalAccountIdentificationAsBankAccountV3AccountIdentification(v *BRLocalAccountIdentification) BankAccountV3AccountIdentification {
	return BankAccountV3AccountIdentification{
		BRLocalAccountIdentification: v,
	}
}

// CALocalAccountIdentificationAsBankAccountV3AccountIdentification is a convenience function that returns CALocalAccountIdentification wrapped in BankAccountV3AccountIdentification
func CALocalAccountIdentificationAsBankAccountV3AccountIdentification(v *CALocalAccountIdentification) BankAccountV3AccountIdentification {
	return BankAccountV3AccountIdentification{
		CALocalAccountIdentification: v,
	}
}

// CZLocalAccountIdentificationAsBankAccountV3AccountIdentification is a convenience function that returns CZLocalAccountIdentification wrapped in BankAccountV3AccountIdentification
func CZLocalAccountIdentificationAsBankAccountV3AccountIdentification(v *CZLocalAccountIdentification) BankAccountV3AccountIdentification {
	return BankAccountV3AccountIdentification{
		CZLocalAccountIdentification: v,
	}
}

// DKLocalAccountIdentificationAsBankAccountV3AccountIdentification is a convenience function that returns DKLocalAccountIdentification wrapped in BankAccountV3AccountIdentification
func DKLocalAccountIdentificationAsBankAccountV3AccountIdentification(v *DKLocalAccountIdentification) BankAccountV3AccountIdentification {
	return BankAccountV3AccountIdentification{
		DKLocalAccountIdentification: v,
	}
}

// HULocalAccountIdentificationAsBankAccountV3AccountIdentification is a convenience function that returns HULocalAccountIdentification wrapped in BankAccountV3AccountIdentification
func HULocalAccountIdentificationAsBankAccountV3AccountIdentification(v *HULocalAccountIdentification) BankAccountV3AccountIdentification {
	return BankAccountV3AccountIdentification{
		HULocalAccountIdentification: v,
	}
}

// IbanAccountIdentificationAsBankAccountV3AccountIdentification is a convenience function that returns IbanAccountIdentification wrapped in BankAccountV3AccountIdentification
func IbanAccountIdentificationAsBankAccountV3AccountIdentification(v *IbanAccountIdentification) BankAccountV3AccountIdentification {
	return BankAccountV3AccountIdentification{
		IbanAccountIdentification: v,
	}
}

// NOLocalAccountIdentificationAsBankAccountV3AccountIdentification is a convenience function that returns NOLocalAccountIdentification wrapped in BankAccountV3AccountIdentification
func NOLocalAccountIdentificationAsBankAccountV3AccountIdentification(v *NOLocalAccountIdentification) BankAccountV3AccountIdentification {
	return BankAccountV3AccountIdentification{
		NOLocalAccountIdentification: v,
	}
}

// NumberAndBicAccountIdentificationAsBankAccountV3AccountIdentification is a convenience function that returns NumberAndBicAccountIdentification wrapped in BankAccountV3AccountIdentification
func NumberAndBicAccountIdentificationAsBankAccountV3AccountIdentification(v *NumberAndBicAccountIdentification) BankAccountV3AccountIdentification {
	return BankAccountV3AccountIdentification{
		NumberAndBicAccountIdentification: v,
	}
}

// PLLocalAccountIdentificationAsBankAccountV3AccountIdentification is a convenience function that returns PLLocalAccountIdentification wrapped in BankAccountV3AccountIdentification
func PLLocalAccountIdentificationAsBankAccountV3AccountIdentification(v *PLLocalAccountIdentification) BankAccountV3AccountIdentification {
	return BankAccountV3AccountIdentification{
		PLLocalAccountIdentification: v,
	}
}

// SELocalAccountIdentificationAsBankAccountV3AccountIdentification is a convenience function that returns SELocalAccountIdentification wrapped in BankAccountV3AccountIdentification
func SELocalAccountIdentificationAsBankAccountV3AccountIdentification(v *SELocalAccountIdentification) BankAccountV3AccountIdentification {
	return BankAccountV3AccountIdentification{
		SELocalAccountIdentification: v,
	}
}

// SGLocalAccountIdentificationAsBankAccountV3AccountIdentification is a convenience function that returns SGLocalAccountIdentification wrapped in BankAccountV3AccountIdentification
func SGLocalAccountIdentificationAsBankAccountV3AccountIdentification(v *SGLocalAccountIdentification) BankAccountV3AccountIdentification {
	return BankAccountV3AccountIdentification{
		SGLocalAccountIdentification: v,
	}
}

// UKLocalAccountIdentificationAsBankAccountV3AccountIdentification is a convenience function that returns UKLocalAccountIdentification wrapped in BankAccountV3AccountIdentification
func UKLocalAccountIdentificationAsBankAccountV3AccountIdentification(v *UKLocalAccountIdentification) BankAccountV3AccountIdentification {
	return BankAccountV3AccountIdentification{
		UKLocalAccountIdentification: v,
	}
}

// USLocalAccountIdentificationAsBankAccountV3AccountIdentification is a convenience function that returns USLocalAccountIdentification wrapped in BankAccountV3AccountIdentification
func USLocalAccountIdentificationAsBankAccountV3AccountIdentification(v *USLocalAccountIdentification) BankAccountV3AccountIdentification {
	return BankAccountV3AccountIdentification{
		USLocalAccountIdentification: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *BankAccountV3AccountIdentification) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into AULocalAccountIdentification
	err = json.Unmarshal(data, &dst.AULocalAccountIdentification)
	if err == nil {
		jsonAULocalAccountIdentification, _ := json.Marshal(dst.AULocalAccountIdentification)
		if string(jsonAULocalAccountIdentification) == "{}" || !dst.AULocalAccountIdentification.isValidType() { // empty struct
			dst.AULocalAccountIdentification = nil
		} else {
			match++
		}
	} else {
		dst.AULocalAccountIdentification = nil
	}

	// try to unmarshal data into BRLocalAccountIdentification
	err = json.Unmarshal(data, &dst.BRLocalAccountIdentification)
	if err == nil {
		jsonBRLocalAccountIdentification, _ := json.Marshal(dst.BRLocalAccountIdentification)
		if string(jsonBRLocalAccountIdentification) == "{}" || !dst.BRLocalAccountIdentification.isValidType() { // empty struct
			dst.BRLocalAccountIdentification = nil
		} else {
			match++
		}
	} else {
		dst.BRLocalAccountIdentification = nil
	}

	// try to unmarshal data into CALocalAccountIdentification
	err = json.Unmarshal(data, &dst.CALocalAccountIdentification)
	if err == nil {
		jsonCALocalAccountIdentification, _ := json.Marshal(dst.CALocalAccountIdentification)
		if string(jsonCALocalAccountIdentification) == "{}" || !dst.CALocalAccountIdentification.isValidType() { // empty struct
			dst.CALocalAccountIdentification = nil
		} else {
			match++
		}
	} else {
		dst.CALocalAccountIdentification = nil
	}

	// try to unmarshal data into CZLocalAccountIdentification
	err = json.Unmarshal(data, &dst.CZLocalAccountIdentification)
	if err == nil {
		jsonCZLocalAccountIdentification, _ := json.Marshal(dst.CZLocalAccountIdentification)
		if string(jsonCZLocalAccountIdentification) == "{}" || !dst.CZLocalAccountIdentification.isValidType() { // empty struct
			dst.CZLocalAccountIdentification = nil
		} else {
			match++
		}
	} else {
		dst.CZLocalAccountIdentification = nil
	}

	// try to unmarshal data into DKLocalAccountIdentification
	err = json.Unmarshal(data, &dst.DKLocalAccountIdentification)
	if err == nil {
		jsonDKLocalAccountIdentification, _ := json.Marshal(dst.DKLocalAccountIdentification)
		if string(jsonDKLocalAccountIdentification) == "{}" || !dst.DKLocalAccountIdentification.isValidType() { // empty struct
			dst.DKLocalAccountIdentification = nil
		} else {
			match++
		}
	} else {
		dst.DKLocalAccountIdentification = nil
	}

	// try to unmarshal data into HULocalAccountIdentification
	err = json.Unmarshal(data, &dst.HULocalAccountIdentification)
	if err == nil {
		jsonHULocalAccountIdentification, _ := json.Marshal(dst.HULocalAccountIdentification)
		if string(jsonHULocalAccountIdentification) == "{}" || !dst.HULocalAccountIdentification.isValidType() { // empty struct
			dst.HULocalAccountIdentification = nil
		} else {
			match++
		}
	} else {
		dst.HULocalAccountIdentification = nil
	}

	// try to unmarshal data into IbanAccountIdentification
	err = json.Unmarshal(data, &dst.IbanAccountIdentification)
	if err == nil {
		jsonIbanAccountIdentification, _ := json.Marshal(dst.IbanAccountIdentification)
		if string(jsonIbanAccountIdentification) == "{}" || !dst.IbanAccountIdentification.isValidType() { // empty struct
			dst.IbanAccountIdentification = nil
		} else {
			match++
		}
	} else {
		dst.IbanAccountIdentification = nil
	}

	// try to unmarshal data into NOLocalAccountIdentification
	err = json.Unmarshal(data, &dst.NOLocalAccountIdentification)
	if err == nil {
		jsonNOLocalAccountIdentification, _ := json.Marshal(dst.NOLocalAccountIdentification)
		if string(jsonNOLocalAccountIdentification) == "{}" || !dst.NOLocalAccountIdentification.isValidType() { // empty struct
			dst.NOLocalAccountIdentification = nil
		} else {
			match++
		}
	} else {
		dst.NOLocalAccountIdentification = nil
	}

	// try to unmarshal data into NumberAndBicAccountIdentification
	err = json.Unmarshal(data, &dst.NumberAndBicAccountIdentification)
	if err == nil {
		jsonNumberAndBicAccountIdentification, _ := json.Marshal(dst.NumberAndBicAccountIdentification)
		if string(jsonNumberAndBicAccountIdentification) == "{}" || !dst.NumberAndBicAccountIdentification.isValidType() { // empty struct
			dst.NumberAndBicAccountIdentification = nil
		} else {
			match++
		}
	} else {
		dst.NumberAndBicAccountIdentification = nil
	}

	// try to unmarshal data into PLLocalAccountIdentification
	err = json.Unmarshal(data, &dst.PLLocalAccountIdentification)
	if err == nil {
		jsonPLLocalAccountIdentification, _ := json.Marshal(dst.PLLocalAccountIdentification)
		if string(jsonPLLocalAccountIdentification) == "{}" || !dst.PLLocalAccountIdentification.isValidType() { // empty struct
			dst.PLLocalAccountIdentification = nil
		} else {
			match++
		}
	} else {
		dst.PLLocalAccountIdentification = nil
	}

	// try to unmarshal data into SELocalAccountIdentification
	err = json.Unmarshal(data, &dst.SELocalAccountIdentification)
	if err == nil {
		jsonSELocalAccountIdentification, _ := json.Marshal(dst.SELocalAccountIdentification)
		if string(jsonSELocalAccountIdentification) == "{}" || !dst.SELocalAccountIdentification.isValidType() { // empty struct
			dst.SELocalAccountIdentification = nil
		} else {
			match++
		}
	} else {
		dst.SELocalAccountIdentification = nil
	}

	// try to unmarshal data into SGLocalAccountIdentification
	err = json.Unmarshal(data, &dst.SGLocalAccountIdentification)
	if err == nil {
		jsonSGLocalAccountIdentification, _ := json.Marshal(dst.SGLocalAccountIdentification)
		if string(jsonSGLocalAccountIdentification) == "{}" || !dst.SGLocalAccountIdentification.isValidType() { // empty struct
			dst.SGLocalAccountIdentification = nil
		} else {
			match++
		}
	} else {
		dst.SGLocalAccountIdentification = nil
	}

	// try to unmarshal data into UKLocalAccountIdentification
	err = json.Unmarshal(data, &dst.UKLocalAccountIdentification)
	if err == nil {
		jsonUKLocalAccountIdentification, _ := json.Marshal(dst.UKLocalAccountIdentification)
		if string(jsonUKLocalAccountIdentification) == "{}" || !dst.UKLocalAccountIdentification.isValidType() { // empty struct
			dst.UKLocalAccountIdentification = nil
		} else {
			match++
		}
	} else {
		dst.UKLocalAccountIdentification = nil
	}

	// try to unmarshal data into USLocalAccountIdentification
	err = json.Unmarshal(data, &dst.USLocalAccountIdentification)
	if err == nil {
		jsonUSLocalAccountIdentification, _ := json.Marshal(dst.USLocalAccountIdentification)
		if string(jsonUSLocalAccountIdentification) == "{}" || !dst.USLocalAccountIdentification.isValidType() { // empty struct
			dst.USLocalAccountIdentification = nil
		} else {
			match++
		}
	} else {
		dst.USLocalAccountIdentification = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.AULocalAccountIdentification = nil
		dst.BRLocalAccountIdentification = nil
		dst.CALocalAccountIdentification = nil
		dst.CZLocalAccountIdentification = nil
		dst.DKLocalAccountIdentification = nil
		dst.HULocalAccountIdentification = nil
		dst.IbanAccountIdentification = nil
		dst.NOLocalAccountIdentification = nil
		dst.NumberAndBicAccountIdentification = nil
		dst.PLLocalAccountIdentification = nil
		dst.SELocalAccountIdentification = nil
		dst.SGLocalAccountIdentification = nil
		dst.UKLocalAccountIdentification = nil
		dst.USLocalAccountIdentification = nil

		return fmt.Errorf("data matches more than one schema in oneOf(BankAccountV3AccountIdentification)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(BankAccountV3AccountIdentification)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src BankAccountV3AccountIdentification) MarshalJSON() ([]byte, error) {
	if src.AULocalAccountIdentification != nil {
		return json.Marshal(&src.AULocalAccountIdentification)
	}

	if src.BRLocalAccountIdentification != nil {
		return json.Marshal(&src.BRLocalAccountIdentification)
	}

	if src.CALocalAccountIdentification != nil {
		return json.Marshal(&src.CALocalAccountIdentification)
	}

	if src.CZLocalAccountIdentification != nil {
		return json.Marshal(&src.CZLocalAccountIdentification)
	}

	if src.DKLocalAccountIdentification != nil {
		return json.Marshal(&src.DKLocalAccountIdentification)
	}

	if src.HULocalAccountIdentification != nil {
		return json.Marshal(&src.HULocalAccountIdentification)
	}

	if src.IbanAccountIdentification != nil {
		return json.Marshal(&src.IbanAccountIdentification)
	}

	if src.NOLocalAccountIdentification != nil {
		return json.Marshal(&src.NOLocalAccountIdentification)
	}

	if src.NumberAndBicAccountIdentification != nil {
		return json.Marshal(&src.NumberAndBicAccountIdentification)
	}

	if src.PLLocalAccountIdentification != nil {
		return json.Marshal(&src.PLLocalAccountIdentification)
	}

	if src.SELocalAccountIdentification != nil {
		return json.Marshal(&src.SELocalAccountIdentification)
	}

	if src.SGLocalAccountIdentification != nil {
		return json.Marshal(&src.SGLocalAccountIdentification)
	}

	if src.UKLocalAccountIdentification != nil {
		return json.Marshal(&src.UKLocalAccountIdentification)
	}

	if src.USLocalAccountIdentification != nil {
		return json.Marshal(&src.USLocalAccountIdentification)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *BankAccountV3AccountIdentification) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	if obj.AULocalAccountIdentification != nil {
		return obj.AULocalAccountIdentification
	}

	if obj.BRLocalAccountIdentification != nil {
		return obj.BRLocalAccountIdentification
	}

	if obj.CALocalAccountIdentification != nil {
		return obj.CALocalAccountIdentification
	}

	if obj.CZLocalAccountIdentification != nil {
		return obj.CZLocalAccountIdentification
	}

	if obj.DKLocalAccountIdentification != nil {
		return obj.DKLocalAccountIdentification
	}

	if obj.HULocalAccountIdentification != nil {
		return obj.HULocalAccountIdentification
	}

	if obj.IbanAccountIdentification != nil {
		return obj.IbanAccountIdentification
	}

	if obj.NOLocalAccountIdentification != nil {
		return obj.NOLocalAccountIdentification
	}

	if obj.NumberAndBicAccountIdentification != nil {
		return obj.NumberAndBicAccountIdentification
	}

	if obj.PLLocalAccountIdentification != nil {
		return obj.PLLocalAccountIdentification
	}

	if obj.SELocalAccountIdentification != nil {
		return obj.SELocalAccountIdentification
	}

	if obj.SGLocalAccountIdentification != nil {
		return obj.SGLocalAccountIdentification
	}

	if obj.UKLocalAccountIdentification != nil {
		return obj.UKLocalAccountIdentification
	}

	if obj.USLocalAccountIdentification != nil {
		return obj.USLocalAccountIdentification
	}

	// all schemas are nil
	return nil
}

type NullableBankAccountV3AccountIdentification struct {
	value *BankAccountV3AccountIdentification
	isSet bool
}

func (v NullableBankAccountV3AccountIdentification) Get() *BankAccountV3AccountIdentification {
	return v.value
}

func (v *NullableBankAccountV3AccountIdentification) Set(val *BankAccountV3AccountIdentification) {
	v.value = val
	v.isSet = true
}

func (v NullableBankAccountV3AccountIdentification) IsSet() bool {
	return v.isSet
}

func (v *NullableBankAccountV3AccountIdentification) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBankAccountV3AccountIdentification(val *BankAccountV3AccountIdentification) *NullableBankAccountV3AccountIdentification {
	return &NullableBankAccountV3AccountIdentification{value: val, isSet: true}
}

func (v NullableBankAccountV3AccountIdentification) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBankAccountV3AccountIdentification) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
