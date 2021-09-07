/*
YASM (Yet Another Skill Management) API

This is an example of using OAuth2 Implicit Flow in a specification to describe security to your API.

API version: 0.8.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// AvailabilityAllOf struct for AvailabilityAllOf
type AvailabilityAllOf struct {
	WorkHours *float32 `json:"workHours,omitempty"`
	PlannedHours *float32 `json:"plannedHours,omitempty"`
}

// NewAvailabilityAllOf instantiates a new AvailabilityAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAvailabilityAllOf() *AvailabilityAllOf {
	this := AvailabilityAllOf{}
	return &this
}

// NewAvailabilityAllOfWithDefaults instantiates a new AvailabilityAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAvailabilityAllOfWithDefaults() *AvailabilityAllOf {
	this := AvailabilityAllOf{}
	return &this
}

// GetWorkHours returns the WorkHours field value if set, zero value otherwise.
func (o *AvailabilityAllOf) GetWorkHours() float32 {
	if o == nil || o.WorkHours == nil {
		var ret float32
		return ret
	}
	return *o.WorkHours
}

// GetWorkHoursOk returns a tuple with the WorkHours field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AvailabilityAllOf) GetWorkHoursOk() (*float32, bool) {
	if o == nil || o.WorkHours == nil {
		return nil, false
	}
	return o.WorkHours, true
}

// HasWorkHours returns a boolean if a field has been set.
func (o *AvailabilityAllOf) HasWorkHours() bool {
	if o != nil && o.WorkHours != nil {
		return true
	}

	return false
}

// SetWorkHours gets a reference to the given float32 and assigns it to the WorkHours field.
func (o *AvailabilityAllOf) SetWorkHours(v float32) {
	o.WorkHours = &v
}

// GetPlannedHours returns the PlannedHours field value if set, zero value otherwise.
func (o *AvailabilityAllOf) GetPlannedHours() float32 {
	if o == nil || o.PlannedHours == nil {
		var ret float32
		return ret
	}
	return *o.PlannedHours
}

// GetPlannedHoursOk returns a tuple with the PlannedHours field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AvailabilityAllOf) GetPlannedHoursOk() (*float32, bool) {
	if o == nil || o.PlannedHours == nil {
		return nil, false
	}
	return o.PlannedHours, true
}

// HasPlannedHours returns a boolean if a field has been set.
func (o *AvailabilityAllOf) HasPlannedHours() bool {
	if o != nil && o.PlannedHours != nil {
		return true
	}

	return false
}

// SetPlannedHours gets a reference to the given float32 and assigns it to the PlannedHours field.
func (o *AvailabilityAllOf) SetPlannedHours(v float32) {
	o.PlannedHours = &v
}

func (o AvailabilityAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.WorkHours != nil {
		toSerialize["workHours"] = o.WorkHours
	}
	if o.PlannedHours != nil {
		toSerialize["plannedHours"] = o.PlannedHours
	}
	return json.Marshal(toSerialize)
}

type NullableAvailabilityAllOf struct {
	value *AvailabilityAllOf
	isSet bool
}

func (v NullableAvailabilityAllOf) Get() *AvailabilityAllOf {
	return v.value
}

func (v *NullableAvailabilityAllOf) Set(val *AvailabilityAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableAvailabilityAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableAvailabilityAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAvailabilityAllOf(val *AvailabilityAllOf) *NullableAvailabilityAllOf {
	return &NullableAvailabilityAllOf{value: val, isSet: true}
}

func (v NullableAvailabilityAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAvailabilityAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


