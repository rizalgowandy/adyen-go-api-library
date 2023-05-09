/*
Configuration API

API version: 2
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package balanceplatform

import (
	"encoding/json"
	"fmt"
)

// SweepConfigurationV2Schedule - The schedule when the `triggerAmount` is evaluated. If the balance meets the threshold, funds are pushed out of or pulled in to the balance account.
type SweepConfigurationV2Schedule struct {
	CronSweepSchedule *CronSweepSchedule
	SweepSchedule     *SweepSchedule
}

// CronSweepScheduleAsSweepConfigurationV2Schedule is a convenience function that returns CronSweepSchedule wrapped in SweepConfigurationV2Schedule
func CronSweepScheduleAsSweepConfigurationV2Schedule(v *CronSweepSchedule) SweepConfigurationV2Schedule {
	return SweepConfigurationV2Schedule{
		CronSweepSchedule: v,
	}
}

// SweepScheduleAsSweepConfigurationV2Schedule is a convenience function that returns SweepSchedule wrapped in SweepConfigurationV2Schedule
func SweepScheduleAsSweepConfigurationV2Schedule(v *SweepSchedule) SweepConfigurationV2Schedule {
	return SweepConfigurationV2Schedule{
		SweepSchedule: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *SweepConfigurationV2Schedule) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into CronSweepSchedule
	err = json.Unmarshal(data, &dst.CronSweepSchedule)
	if err == nil {
		jsonCronSweepSchedule, _ := json.Marshal(dst.CronSweepSchedule)
		if string(jsonCronSweepSchedule) == "{}" || !dst.CronSweepSchedule.isValidType() { // empty struct
			dst.CronSweepSchedule = nil
		} else {
			match++
		}
	} else {
		dst.CronSweepSchedule = nil
	}

	// try to unmarshal data into SweepSchedule
	err = json.Unmarshal(data, &dst.SweepSchedule)
	if err == nil {
		jsonSweepSchedule, _ := json.Marshal(dst.SweepSchedule)
		if string(jsonSweepSchedule) == "{}" || !dst.SweepSchedule.isValidType() { // empty struct
			dst.SweepSchedule = nil
		} else {
			match++
		}
	} else {
		dst.SweepSchedule = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.CronSweepSchedule = nil
		dst.SweepSchedule = nil

		return fmt.Errorf("data matches more than one schema in oneOf(SweepConfigurationV2Schedule)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(SweepConfigurationV2Schedule)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src SweepConfigurationV2Schedule) MarshalJSON() ([]byte, error) {
	if src.CronSweepSchedule != nil {
		return json.Marshal(&src.CronSweepSchedule)
	}

	if src.SweepSchedule != nil {
		return json.Marshal(&src.SweepSchedule)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *SweepConfigurationV2Schedule) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	if obj.CronSweepSchedule != nil {
		return obj.CronSweepSchedule
	}

	if obj.SweepSchedule != nil {
		return obj.SweepSchedule
	}

	// all schemas are nil
	return nil
}

type NullableSweepConfigurationV2Schedule struct {
	value *SweepConfigurationV2Schedule
	isSet bool
}

func (v NullableSweepConfigurationV2Schedule) Get() *SweepConfigurationV2Schedule {
	return v.value
}

func (v *NullableSweepConfigurationV2Schedule) Set(val *SweepConfigurationV2Schedule) {
	v.value = val
	v.isSet = true
}

func (v NullableSweepConfigurationV2Schedule) IsSet() bool {
	return v.isSet
}

func (v *NullableSweepConfigurationV2Schedule) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSweepConfigurationV2Schedule(val *SweepConfigurationV2Schedule) *NullableSweepConfigurationV2Schedule {
	return &NullableSweepConfigurationV2Schedule{value: val, isSet: true}
}

func (v NullableSweepConfigurationV2Schedule) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSweepConfigurationV2Schedule) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
