/*
 * YASM (Yet Another Skill Management) API
 *
 * This is an example of using OAuth2 Implicit Flow in a specification to describe security to your API.
 *
 * API version: VERSION
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// Date struct for Date
type Date struct {
	Year float32 `json:"year"`
	Month int32 `json:"month"`
	Day *int32 `json:"day,omitempty"`
}

// NewDate instantiates a new Date object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDate(year float32, month int32) *Date {
	this := Date{}
	this.Year = year
	this.Month = month
	return &this
}

// NewDateWithDefaults instantiates a new Date object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDateWithDefaults() *Date {
	this := Date{}
	return &this
}

// GetYear returns the Year field value
func (o *Date) GetYear() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Year
}

// GetYearOk returns a tuple with the Year field value
// and a boolean to check if the value has been set.
func (o *Date) GetYearOk() (*float32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Year, true
}

// SetYear sets field value
func (o *Date) SetYear(v float32) {
	o.Year = v
}

// GetMonth returns the Month field value
func (o *Date) GetMonth() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Month
}

// GetMonthOk returns a tuple with the Month field value
// and a boolean to check if the value has been set.
func (o *Date) GetMonthOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Month, true
}

// SetMonth sets field value
func (o *Date) SetMonth(v int32) {
	o.Month = v
}

// GetDay returns the Day field value if set, zero value otherwise.
func (o *Date) GetDay() int32 {
	if o == nil || o.Day == nil {
		var ret int32
		return ret
	}
	return *o.Day
}

// GetDayOk returns a tuple with the Day field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Date) GetDayOk() (*int32, bool) {
	if o == nil || o.Day == nil {
		return nil, false
	}
	return o.Day, true
}

// HasDay returns a boolean if a field has been set.
func (o *Date) HasDay() bool {
	if o != nil && o.Day != nil {
		return true
	}

	return false
}

// SetDay gets a reference to the given int32 and assigns it to the Day field.
func (o *Date) SetDay(v int32) {
	o.Day = &v
}

func (o Date) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["year"] = o.Year
	}
	if true {
		toSerialize["month"] = o.Month
	}
	if o.Day != nil {
		toSerialize["day"] = o.Day
	}
	return json.Marshal(toSerialize)
}

type NullableDate struct {
	value *Date
	isSet bool
}

func (v NullableDate) Get() *Date {
	return v.value
}

func (v *NullableDate) Set(val *Date) {
	v.value = val
	v.isSet = true
}

func (v NullableDate) IsSet() bool {
	return v.isSet
}

func (v *NullableDate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDate(val *Date) *NullableDate {
	return &NullableDate{value: val, isSet: true}
}

func (v NullableDate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


