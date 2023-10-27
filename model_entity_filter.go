/*
YASM (Yet Another Skill Management) API

This is an example of using OAuth2 Implicit Flow in a specification to describe security to your API.

API version: 1.5.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// checks if the EntityFilter type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EntityFilter{}

// EntityFilter struct for EntityFilter
type EntityFilter struct {
	AbstractEntityFilter
}

// NewEntityFilter instantiates a new EntityFilter object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEntityFilter(id string) *EntityFilter {
	this := EntityFilter{}
	this.Id = id
	return &this
}

// NewEntityFilterWithDefaults instantiates a new EntityFilter object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEntityFilterWithDefaults() *EntityFilter {
	this := EntityFilter{}
	return &this
}

func (o EntityFilter) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EntityFilter) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	serializedAbstractEntityFilter, errAbstractEntityFilter := json.Marshal(o.AbstractEntityFilter)
	if errAbstractEntityFilter != nil {
		return map[string]interface{}{}, errAbstractEntityFilter
	}
	errAbstractEntityFilter = json.Unmarshal([]byte(serializedAbstractEntityFilter), &toSerialize)
	if errAbstractEntityFilter != nil {
		return map[string]interface{}{}, errAbstractEntityFilter
	}
	return toSerialize, nil
}

type NullableEntityFilter struct {
	value *EntityFilter
	isSet bool
}

func (v NullableEntityFilter) Get() *EntityFilter {
	return v.value
}

func (v *NullableEntityFilter) Set(val *EntityFilter) {
	v.value = val
	v.isSet = true
}

func (v NullableEntityFilter) IsSet() bool {
	return v.isSet
}

func (v *NullableEntityFilter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEntityFilter(val *EntityFilter) *NullableEntityFilter {
	return &NullableEntityFilter{value: val, isSet: true}
}

func (v NullableEntityFilter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEntityFilter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


