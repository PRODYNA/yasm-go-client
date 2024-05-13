/*
YASM (Yet Another Skill Management) API

This is an example of using OAuth2 Implicit Flow in a specification to describe security to your API.

API version: 1.26.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// checks if the SkillLinkUpdate type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SkillLinkUpdate{}

// SkillLinkUpdate struct for SkillLinkUpdate
type SkillLinkUpdate struct {
	KindGiver *bool `json:"kindGiver,omitempty"`
}

// NewSkillLinkUpdate instantiates a new SkillLinkUpdate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSkillLinkUpdate() *SkillLinkUpdate {
	this := SkillLinkUpdate{}
	var kindGiver bool = false
	this.KindGiver = &kindGiver
	return &this
}

// NewSkillLinkUpdateWithDefaults instantiates a new SkillLinkUpdate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSkillLinkUpdateWithDefaults() *SkillLinkUpdate {
	this := SkillLinkUpdate{}
	var kindGiver bool = false
	this.KindGiver = &kindGiver
	return &this
}

// GetKindGiver returns the KindGiver field value if set, zero value otherwise.
func (o *SkillLinkUpdate) GetKindGiver() bool {
	if o == nil || IsNil(o.KindGiver) {
		var ret bool
		return ret
	}
	return *o.KindGiver
}

// GetKindGiverOk returns a tuple with the KindGiver field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SkillLinkUpdate) GetKindGiverOk() (*bool, bool) {
	if o == nil || IsNil(o.KindGiver) {
		return nil, false
	}
	return o.KindGiver, true
}

// HasKindGiver returns a boolean if a field has been set.
func (o *SkillLinkUpdate) HasKindGiver() bool {
	if o != nil && !IsNil(o.KindGiver) {
		return true
	}

	return false
}

// SetKindGiver gets a reference to the given bool and assigns it to the KindGiver field.
func (o *SkillLinkUpdate) SetKindGiver(v bool) {
	o.KindGiver = &v
}

func (o SkillLinkUpdate) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SkillLinkUpdate) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.KindGiver) {
		toSerialize["kindGiver"] = o.KindGiver
	}
	return toSerialize, nil
}

type NullableSkillLinkUpdate struct {
	value *SkillLinkUpdate
	isSet bool
}

func (v NullableSkillLinkUpdate) Get() *SkillLinkUpdate {
	return v.value
}

func (v *NullableSkillLinkUpdate) Set(val *SkillLinkUpdate) {
	v.value = val
	v.isSet = true
}

func (v NullableSkillLinkUpdate) IsSet() bool {
	return v.isSet
}

func (v *NullableSkillLinkUpdate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSkillLinkUpdate(val *SkillLinkUpdate) *NullableSkillLinkUpdate {
	return &NullableSkillLinkUpdate{value: val, isSet: true}
}

func (v NullableSkillLinkUpdate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSkillLinkUpdate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


