/*
YASM (Yet Another Skill Management) API

This is an example of using OAuth2 Implicit Flow in a specification to describe security to your API.

API version: 1.64.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// checks if the TimeframeFilter type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TimeframeFilter{}

// TimeframeFilter struct for TimeframeFilter
type TimeframeFilter struct {
	Startdate *string `json:"startdate,omitempty"`
	Enddate *string `json:"enddate,omitempty"`
	ObjectType *string `json:"objectType,omitempty"`
}

// NewTimeframeFilter instantiates a new TimeframeFilter object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTimeframeFilter() *TimeframeFilter {
	this := TimeframeFilter{}
	return &this
}

// NewTimeframeFilterWithDefaults instantiates a new TimeframeFilter object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTimeframeFilterWithDefaults() *TimeframeFilter {
	this := TimeframeFilter{}
	return &this
}

// GetStartdate returns the Startdate field value if set, zero value otherwise.
func (o *TimeframeFilter) GetStartdate() string {
	if o == nil || IsNil(o.Startdate) {
		var ret string
		return ret
	}
	return *o.Startdate
}

// GetStartdateOk returns a tuple with the Startdate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TimeframeFilter) GetStartdateOk() (*string, bool) {
	if o == nil || IsNil(o.Startdate) {
		return nil, false
	}
	return o.Startdate, true
}

// HasStartdate returns a boolean if a field has been set.
func (o *TimeframeFilter) HasStartdate() bool {
	if o != nil && !IsNil(o.Startdate) {
		return true
	}

	return false
}

// SetStartdate gets a reference to the given string and assigns it to the Startdate field.
func (o *TimeframeFilter) SetStartdate(v string) {
	o.Startdate = &v
}

// GetEnddate returns the Enddate field value if set, zero value otherwise.
func (o *TimeframeFilter) GetEnddate() string {
	if o == nil || IsNil(o.Enddate) {
		var ret string
		return ret
	}
	return *o.Enddate
}

// GetEnddateOk returns a tuple with the Enddate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TimeframeFilter) GetEnddateOk() (*string, bool) {
	if o == nil || IsNil(o.Enddate) {
		return nil, false
	}
	return o.Enddate, true
}

// HasEnddate returns a boolean if a field has been set.
func (o *TimeframeFilter) HasEnddate() bool {
	if o != nil && !IsNil(o.Enddate) {
		return true
	}

	return false
}

// SetEnddate gets a reference to the given string and assigns it to the Enddate field.
func (o *TimeframeFilter) SetEnddate(v string) {
	o.Enddate = &v
}

// GetObjectType returns the ObjectType field value if set, zero value otherwise.
func (o *TimeframeFilter) GetObjectType() string {
	if o == nil || IsNil(o.ObjectType) {
		var ret string
		return ret
	}
	return *o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TimeframeFilter) GetObjectTypeOk() (*string, bool) {
	if o == nil || IsNil(o.ObjectType) {
		return nil, false
	}
	return o.ObjectType, true
}

// HasObjectType returns a boolean if a field has been set.
func (o *TimeframeFilter) HasObjectType() bool {
	if o != nil && !IsNil(o.ObjectType) {
		return true
	}

	return false
}

// SetObjectType gets a reference to the given string and assigns it to the ObjectType field.
func (o *TimeframeFilter) SetObjectType(v string) {
	o.ObjectType = &v
}

func (o TimeframeFilter) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TimeframeFilter) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Startdate) {
		toSerialize["startdate"] = o.Startdate
	}
	if !IsNil(o.Enddate) {
		toSerialize["enddate"] = o.Enddate
	}
	if !IsNil(o.ObjectType) {
		toSerialize["objectType"] = o.ObjectType
	}
	return toSerialize, nil
}

type NullableTimeframeFilter struct {
	value *TimeframeFilter
	isSet bool
}

func (v NullableTimeframeFilter) Get() *TimeframeFilter {
	return v.value
}

func (v *NullableTimeframeFilter) Set(val *TimeframeFilter) {
	v.value = val
	v.isSet = true
}

func (v NullableTimeframeFilter) IsSet() bool {
	return v.isSet
}

func (v *NullableTimeframeFilter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTimeframeFilter(val *TimeframeFilter) *NullableTimeframeFilter {
	return &NullableTimeframeFilter{value: val, isSet: true}
}

func (v NullableTimeframeFilter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTimeframeFilter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


