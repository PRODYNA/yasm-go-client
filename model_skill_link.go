/*
YASM (Yet Another Skill Management) API

This is an example of using OAuth2 Implicit Flow in a specification to describe security to your API.

API version: 1.65.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// checks if the SkillLink type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SkillLink{}

// SkillLink struct for SkillLink
type SkillLink struct {
	Skill *Skill `json:"skill,omitempty"`
	KindGiver *bool `json:"kindGiver,omitempty"`
	ChildCount *int32 `json:"childCount,omitempty"`
}

// NewSkillLink instantiates a new SkillLink object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSkillLink() *SkillLink {
	this := SkillLink{}
	var kindGiver bool = false
	this.KindGiver = &kindGiver
	var childCount int32 = 0
	this.ChildCount = &childCount
	return &this
}

// NewSkillLinkWithDefaults instantiates a new SkillLink object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSkillLinkWithDefaults() *SkillLink {
	this := SkillLink{}
	var kindGiver bool = false
	this.KindGiver = &kindGiver
	var childCount int32 = 0
	this.ChildCount = &childCount
	return &this
}

// GetSkill returns the Skill field value if set, zero value otherwise.
func (o *SkillLink) GetSkill() Skill {
	if o == nil || IsNil(o.Skill) {
		var ret Skill
		return ret
	}
	return *o.Skill
}

// GetSkillOk returns a tuple with the Skill field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SkillLink) GetSkillOk() (*Skill, bool) {
	if o == nil || IsNil(o.Skill) {
		return nil, false
	}
	return o.Skill, true
}

// HasSkill returns a boolean if a field has been set.
func (o *SkillLink) HasSkill() bool {
	if o != nil && !IsNil(o.Skill) {
		return true
	}

	return false
}

// SetSkill gets a reference to the given Skill and assigns it to the Skill field.
func (o *SkillLink) SetSkill(v Skill) {
	o.Skill = &v
}

// GetKindGiver returns the KindGiver field value if set, zero value otherwise.
func (o *SkillLink) GetKindGiver() bool {
	if o == nil || IsNil(o.KindGiver) {
		var ret bool
		return ret
	}
	return *o.KindGiver
}

// GetKindGiverOk returns a tuple with the KindGiver field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SkillLink) GetKindGiverOk() (*bool, bool) {
	if o == nil || IsNil(o.KindGiver) {
		return nil, false
	}
	return o.KindGiver, true
}

// HasKindGiver returns a boolean if a field has been set.
func (o *SkillLink) HasKindGiver() bool {
	if o != nil && !IsNil(o.KindGiver) {
		return true
	}

	return false
}

// SetKindGiver gets a reference to the given bool and assigns it to the KindGiver field.
func (o *SkillLink) SetKindGiver(v bool) {
	o.KindGiver = &v
}

// GetChildCount returns the ChildCount field value if set, zero value otherwise.
func (o *SkillLink) GetChildCount() int32 {
	if o == nil || IsNil(o.ChildCount) {
		var ret int32
		return ret
	}
	return *o.ChildCount
}

// GetChildCountOk returns a tuple with the ChildCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SkillLink) GetChildCountOk() (*int32, bool) {
	if o == nil || IsNil(o.ChildCount) {
		return nil, false
	}
	return o.ChildCount, true
}

// HasChildCount returns a boolean if a field has been set.
func (o *SkillLink) HasChildCount() bool {
	if o != nil && !IsNil(o.ChildCount) {
		return true
	}

	return false
}

// SetChildCount gets a reference to the given int32 and assigns it to the ChildCount field.
func (o *SkillLink) SetChildCount(v int32) {
	o.ChildCount = &v
}

func (o SkillLink) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SkillLink) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Skill) {
		toSerialize["skill"] = o.Skill
	}
	if !IsNil(o.KindGiver) {
		toSerialize["kindGiver"] = o.KindGiver
	}
	if !IsNil(o.ChildCount) {
		toSerialize["childCount"] = o.ChildCount
	}
	return toSerialize, nil
}

type NullableSkillLink struct {
	value *SkillLink
	isSet bool
}

func (v NullableSkillLink) Get() *SkillLink {
	return v.value
}

func (v *NullableSkillLink) Set(val *SkillLink) {
	v.value = val
	v.isSet = true
}

func (v NullableSkillLink) IsSet() bool {
	return v.isSet
}

func (v *NullableSkillLink) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSkillLink(val *SkillLink) *NullableSkillLink {
	return &NullableSkillLink{value: val, isSet: true}
}

func (v NullableSkillLink) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSkillLink) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


