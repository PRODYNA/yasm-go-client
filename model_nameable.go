/*
YASM (Yet Another Skill Management) API

This is an example of using OAuth2 Implicit Flow in a specification to describe security to your API.

API version: 1.38.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// checks if the Nameable type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Nameable{}

// Nameable struct for Nameable
type Nameable struct {
	Name string `json:"name"`
}

// NewNameable instantiates a new Nameable object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNameable(name string) *Nameable {
	this := Nameable{}
	this.Name = name
	return &this
}

// NewNameableWithDefaults instantiates a new Nameable object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNameableWithDefaults() *Nameable {
	this := Nameable{}
	return &this
}

// GetName returns the Name field value
func (o *Nameable) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *Nameable) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *Nameable) SetName(v string) {
	o.Name = v
}

func (o Nameable) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Nameable) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	return toSerialize, nil
}

type NullableNameable struct {
	value *Nameable
	isSet bool
}

func (v NullableNameable) Get() *Nameable {
	return v.value
}

func (v *NullableNameable) Set(val *Nameable) {
	v.value = val
	v.isSet = true
}

func (v NullableNameable) IsSet() bool {
	return v.isSet
}

func (v *NullableNameable) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNameable(val *Nameable) *NullableNameable {
	return &NullableNameable{value: val, isSet: true}
}

func (v NullableNameable) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNameable) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


