/*
YASM (Yet Another Skill Management) API

This is an example of using OAuth2 Implicit Flow in a specification to describe security to your API.

API version: 1.34.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// checks if the SkillGroup type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SkillGroup{}

// SkillGroup struct for SkillGroup
type SkillGroup struct {
	Group *Skill `json:"group,omitempty"`
	Skills []Skill `json:"skills,omitempty"`
}

// NewSkillGroup instantiates a new SkillGroup object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSkillGroup() *SkillGroup {
	this := SkillGroup{}
	return &this
}

// NewSkillGroupWithDefaults instantiates a new SkillGroup object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSkillGroupWithDefaults() *SkillGroup {
	this := SkillGroup{}
	return &this
}

// GetGroup returns the Group field value if set, zero value otherwise.
func (o *SkillGroup) GetGroup() Skill {
	if o == nil || IsNil(o.Group) {
		var ret Skill
		return ret
	}
	return *o.Group
}

// GetGroupOk returns a tuple with the Group field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SkillGroup) GetGroupOk() (*Skill, bool) {
	if o == nil || IsNil(o.Group) {
		return nil, false
	}
	return o.Group, true
}

// HasGroup returns a boolean if a field has been set.
func (o *SkillGroup) HasGroup() bool {
	if o != nil && !IsNil(o.Group) {
		return true
	}

	return false
}

// SetGroup gets a reference to the given Skill and assigns it to the Group field.
func (o *SkillGroup) SetGroup(v Skill) {
	o.Group = &v
}

// GetSkills returns the Skills field value if set, zero value otherwise.
func (o *SkillGroup) GetSkills() []Skill {
	if o == nil || IsNil(o.Skills) {
		var ret []Skill
		return ret
	}
	return o.Skills
}

// GetSkillsOk returns a tuple with the Skills field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SkillGroup) GetSkillsOk() ([]Skill, bool) {
	if o == nil || IsNil(o.Skills) {
		return nil, false
	}
	return o.Skills, true
}

// HasSkills returns a boolean if a field has been set.
func (o *SkillGroup) HasSkills() bool {
	if o != nil && !IsNil(o.Skills) {
		return true
	}

	return false
}

// SetSkills gets a reference to the given []Skill and assigns it to the Skills field.
func (o *SkillGroup) SetSkills(v []Skill) {
	o.Skills = v
}

func (o SkillGroup) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SkillGroup) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Group) {
		toSerialize["group"] = o.Group
	}
	if !IsNil(o.Skills) {
		toSerialize["skills"] = o.Skills
	}
	return toSerialize, nil
}

type NullableSkillGroup struct {
	value *SkillGroup
	isSet bool
}

func (v NullableSkillGroup) Get() *SkillGroup {
	return v.value
}

func (v *NullableSkillGroup) Set(val *SkillGroup) {
	v.value = val
	v.isSet = true
}

func (v NullableSkillGroup) IsSet() bool {
	return v.isSet
}

func (v *NullableSkillGroup) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSkillGroup(val *SkillGroup) *NullableSkillGroup {
	return &NullableSkillGroup{value: val, isSet: true}
}

func (v NullableSkillGroup) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSkillGroup) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


