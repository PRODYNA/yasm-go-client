/*
YASM (Yet Another Skill Management) API

This is an example of using OAuth2 Implicit Flow in a specification to describe security to your API.

API version: 1.6.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// checks if the MinMax type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MinMax{}

// MinMax struct for MinMax
type MinMax struct {
	Min *int32 `json:"min,omitempty"`
	Max *int32 `json:"max,omitempty"`
}

// NewMinMax instantiates a new MinMax object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMinMax() *MinMax {
	this := MinMax{}
	return &this
}

// NewMinMaxWithDefaults instantiates a new MinMax object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMinMaxWithDefaults() *MinMax {
	this := MinMax{}
	return &this
}

// GetMin returns the Min field value if set, zero value otherwise.
func (o *MinMax) GetMin() int32 {
	if o == nil || IsNil(o.Min) {
		var ret int32
		return ret
	}
	return *o.Min
}

// GetMinOk returns a tuple with the Min field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MinMax) GetMinOk() (*int32, bool) {
	if o == nil || IsNil(o.Min) {
		return nil, false
	}
	return o.Min, true
}

// HasMin returns a boolean if a field has been set.
func (o *MinMax) HasMin() bool {
	if o != nil && !IsNil(o.Min) {
		return true
	}

	return false
}

// SetMin gets a reference to the given int32 and assigns it to the Min field.
func (o *MinMax) SetMin(v int32) {
	o.Min = &v
}

// GetMax returns the Max field value if set, zero value otherwise.
func (o *MinMax) GetMax() int32 {
	if o == nil || IsNil(o.Max) {
		var ret int32
		return ret
	}
	return *o.Max
}

// GetMaxOk returns a tuple with the Max field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MinMax) GetMaxOk() (*int32, bool) {
	if o == nil || IsNil(o.Max) {
		return nil, false
	}
	return o.Max, true
}

// HasMax returns a boolean if a field has been set.
func (o *MinMax) HasMax() bool {
	if o != nil && !IsNil(o.Max) {
		return true
	}

	return false
}

// SetMax gets a reference to the given int32 and assigns it to the Max field.
func (o *MinMax) SetMax(v int32) {
	o.Max = &v
}

func (o MinMax) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MinMax) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Min) {
		toSerialize["min"] = o.Min
	}
	if !IsNil(o.Max) {
		toSerialize["max"] = o.Max
	}
	return toSerialize, nil
}

type NullableMinMax struct {
	value *MinMax
	isSet bool
}

func (v NullableMinMax) Get() *MinMax {
	return v.value
}

func (v *NullableMinMax) Set(val *MinMax) {
	v.value = val
	v.isSet = true
}

func (v NullableMinMax) IsSet() bool {
	return v.isSet
}

func (v *NullableMinMax) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMinMax(val *MinMax) *NullableMinMax {
	return &NullableMinMax{value: val, isSet: true}
}

func (v NullableMinMax) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMinMax) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


