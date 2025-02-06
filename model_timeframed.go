/*
YASM (Yet Another Skill Management) API

This is an example of using OAuth2 Implicit Flow in a specification to describe security to your API.

API version: 1.61.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// checks if the Timeframed type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Timeframed{}

// Timeframed struct for Timeframed
type Timeframed struct {
	Startdate string `json:"startdate"`
	Enddate *string `json:"enddate,omitempty"`
	ObjectType *string `json:"objectType,omitempty"`
}

// NewTimeframed instantiates a new Timeframed object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTimeframed(startdate string) *Timeframed {
	this := Timeframed{}
	this.Startdate = startdate
	return &this
}

// NewTimeframedWithDefaults instantiates a new Timeframed object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTimeframedWithDefaults() *Timeframed {
	this := Timeframed{}
	return &this
}

// GetStartdate returns the Startdate field value
func (o *Timeframed) GetStartdate() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Startdate
}

// GetStartdateOk returns a tuple with the Startdate field value
// and a boolean to check if the value has been set.
func (o *Timeframed) GetStartdateOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Startdate, true
}

// SetStartdate sets field value
func (o *Timeframed) SetStartdate(v string) {
	o.Startdate = v
}

// GetEnddate returns the Enddate field value if set, zero value otherwise.
func (o *Timeframed) GetEnddate() string {
	if o == nil || IsNil(o.Enddate) {
		var ret string
		return ret
	}
	return *o.Enddate
}

// GetEnddateOk returns a tuple with the Enddate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Timeframed) GetEnddateOk() (*string, bool) {
	if o == nil || IsNil(o.Enddate) {
		return nil, false
	}
	return o.Enddate, true
}

// HasEnddate returns a boolean if a field has been set.
func (o *Timeframed) HasEnddate() bool {
	if o != nil && !IsNil(o.Enddate) {
		return true
	}

	return false
}

// SetEnddate gets a reference to the given string and assigns it to the Enddate field.
func (o *Timeframed) SetEnddate(v string) {
	o.Enddate = &v
}

// GetObjectType returns the ObjectType field value if set, zero value otherwise.
func (o *Timeframed) GetObjectType() string {
	if o == nil || IsNil(o.ObjectType) {
		var ret string
		return ret
	}
	return *o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Timeframed) GetObjectTypeOk() (*string, bool) {
	if o == nil || IsNil(o.ObjectType) {
		return nil, false
	}
	return o.ObjectType, true
}

// HasObjectType returns a boolean if a field has been set.
func (o *Timeframed) HasObjectType() bool {
	if o != nil && !IsNil(o.ObjectType) {
		return true
	}

	return false
}

// SetObjectType gets a reference to the given string and assigns it to the ObjectType field.
func (o *Timeframed) SetObjectType(v string) {
	o.ObjectType = &v
}

func (o Timeframed) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Timeframed) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["startdate"] = o.Startdate
	if !IsNil(o.Enddate) {
		toSerialize["enddate"] = o.Enddate
	}
	if !IsNil(o.ObjectType) {
		toSerialize["objectType"] = o.ObjectType
	}
	return toSerialize, nil
}

type NullableTimeframed struct {
	value *Timeframed
	isSet bool
}

func (v NullableTimeframed) Get() *Timeframed {
	return v.value
}

func (v *NullableTimeframed) Set(val *Timeframed) {
	v.value = val
	v.isSet = true
}

func (v NullableTimeframed) IsSet() bool {
	return v.isSet
}

func (v *NullableTimeframed) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTimeframed(val *Timeframed) *NullableTimeframed {
	return &NullableTimeframed{value: val, isSet: true}
}

func (v NullableTimeframed) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTimeframed) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


