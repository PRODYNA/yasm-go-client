/*
YASM (Yet Another Skill Management) API

This is an example of using OAuth2 Implicit Flow in a specification to describe security to your API.

API version: 1.4.8
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// ExperienceSkillGroup struct for ExperienceSkillGroup
type ExperienceSkillGroup struct {
	Group *Skill `json:"group,omitempty"`
	Skills []ExperienceSkill `json:"skills,omitempty"`
}

// NewExperienceSkillGroup instantiates a new ExperienceSkillGroup object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewExperienceSkillGroup() *ExperienceSkillGroup {
	this := ExperienceSkillGroup{}
	return &this
}

// NewExperienceSkillGroupWithDefaults instantiates a new ExperienceSkillGroup object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewExperienceSkillGroupWithDefaults() *ExperienceSkillGroup {
	this := ExperienceSkillGroup{}
	return &this
}

// GetGroup returns the Group field value if set, zero value otherwise.
func (o *ExperienceSkillGroup) GetGroup() Skill {
	if o == nil || o.Group == nil {
		var ret Skill
		return ret
	}
	return *o.Group
}

// GetGroupOk returns a tuple with the Group field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExperienceSkillGroup) GetGroupOk() (*Skill, bool) {
	if o == nil || o.Group == nil {
		return nil, false
	}
	return o.Group, true
}

// HasGroup returns a boolean if a field has been set.
func (o *ExperienceSkillGroup) HasGroup() bool {
	if o != nil && o.Group != nil {
		return true
	}

	return false
}

// SetGroup gets a reference to the given Skill and assigns it to the Group field.
func (o *ExperienceSkillGroup) SetGroup(v Skill) {
	o.Group = &v
}

// GetSkills returns the Skills field value if set, zero value otherwise.
func (o *ExperienceSkillGroup) GetSkills() []ExperienceSkill {
	if o == nil || o.Skills == nil {
		var ret []ExperienceSkill
		return ret
	}
	return o.Skills
}

// GetSkillsOk returns a tuple with the Skills field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExperienceSkillGroup) GetSkillsOk() ([]ExperienceSkill, bool) {
	if o == nil || o.Skills == nil {
		return nil, false
	}
	return o.Skills, true
}

// HasSkills returns a boolean if a field has been set.
func (o *ExperienceSkillGroup) HasSkills() bool {
	if o != nil && o.Skills != nil {
		return true
	}

	return false
}

// SetSkills gets a reference to the given []ExperienceSkill and assigns it to the Skills field.
func (o *ExperienceSkillGroup) SetSkills(v []ExperienceSkill) {
	o.Skills = v
}

func (o ExperienceSkillGroup) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Group != nil {
		toSerialize["group"] = o.Group
	}
	if o.Skills != nil {
		toSerialize["skills"] = o.Skills
	}
	return json.Marshal(toSerialize)
}

type NullableExperienceSkillGroup struct {
	value *ExperienceSkillGroup
	isSet bool
}

func (v NullableExperienceSkillGroup) Get() *ExperienceSkillGroup {
	return v.value
}

func (v *NullableExperienceSkillGroup) Set(val *ExperienceSkillGroup) {
	v.value = val
	v.isSet = true
}

func (v NullableExperienceSkillGroup) IsSet() bool {
	return v.isSet
}

func (v *NullableExperienceSkillGroup) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableExperienceSkillGroup(val *ExperienceSkillGroup) *NullableExperienceSkillGroup {
	return &NullableExperienceSkillGroup{value: val, isSet: true}
}

func (v NullableExperienceSkillGroup) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableExperienceSkillGroup) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


