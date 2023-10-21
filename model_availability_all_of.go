/*
YASM (Yet Another Skill Management) API

This is an example of using OAuth2 Implicit Flow in a specification to describe security to your API.

API version: 1.4.11
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// AvailabilityAllOf struct for AvailabilityAllOf
type AvailabilityAllOf struct {
	WorkHours float32 `json:"workHours"`
	PlannedHours float32 `json:"plannedHours"`
	Descriptions []string `json:"descriptions,omitempty"`
}

// NewAvailabilityAllOf instantiates a new AvailabilityAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAvailabilityAllOf(workHours float32, plannedHours float32) *AvailabilityAllOf {
	this := AvailabilityAllOf{}
	this.WorkHours = workHours
	this.PlannedHours = plannedHours
	return &this
}

// NewAvailabilityAllOfWithDefaults instantiates a new AvailabilityAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAvailabilityAllOfWithDefaults() *AvailabilityAllOf {
	this := AvailabilityAllOf{}
	return &this
}

// GetWorkHours returns the WorkHours field value
func (o *AvailabilityAllOf) GetWorkHours() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.WorkHours
}

// GetWorkHoursOk returns a tuple with the WorkHours field value
// and a boolean to check if the value has been set.
func (o *AvailabilityAllOf) GetWorkHoursOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.WorkHours, true
}

// SetWorkHours sets field value
func (o *AvailabilityAllOf) SetWorkHours(v float32) {
	o.WorkHours = v
}

// GetPlannedHours returns the PlannedHours field value
func (o *AvailabilityAllOf) GetPlannedHours() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.PlannedHours
}

// GetPlannedHoursOk returns a tuple with the PlannedHours field value
// and a boolean to check if the value has been set.
func (o *AvailabilityAllOf) GetPlannedHoursOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PlannedHours, true
}

// SetPlannedHours sets field value
func (o *AvailabilityAllOf) SetPlannedHours(v float32) {
	o.PlannedHours = v
}

// GetDescriptions returns the Descriptions field value if set, zero value otherwise.
func (o *AvailabilityAllOf) GetDescriptions() []string {
	if o == nil || o.Descriptions == nil {
		var ret []string
		return ret
	}
	return o.Descriptions
}

// GetDescriptionsOk returns a tuple with the Descriptions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AvailabilityAllOf) GetDescriptionsOk() ([]string, bool) {
	if o == nil || o.Descriptions == nil {
		return nil, false
	}
	return o.Descriptions, true
}

// HasDescriptions returns a boolean if a field has been set.
func (o *AvailabilityAllOf) HasDescriptions() bool {
	if o != nil && o.Descriptions != nil {
		return true
	}

	return false
}

// SetDescriptions gets a reference to the given []string and assigns it to the Descriptions field.
func (o *AvailabilityAllOf) SetDescriptions(v []string) {
	o.Descriptions = v
}

func (o AvailabilityAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["workHours"] = o.WorkHours
	}
	if true {
		toSerialize["plannedHours"] = o.PlannedHours
	}
	if o.Descriptions != nil {
		toSerialize["descriptions"] = o.Descriptions
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


