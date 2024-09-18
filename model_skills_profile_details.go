/*
YASM (Yet Another Skill Management) API

This is an example of using OAuth2 Implicit Flow in a specification to describe security to your API.

API version: 1.41.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// checks if the SkillsProfileDetails type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SkillsProfileDetails{}

// SkillsProfileDetails struct for SkillsProfileDetails
type SkillsProfileDetails struct {
	Audit *Audit `json:"audit,omitempty"`
	SkillsProfile *SkillsProfile `json:"skillsProfile,omitempty"`
	Person *Person `json:"person,omitempty"`
	Skills []ExperienceSkill `json:"skills,omitempty"`
}

// NewSkillsProfileDetails instantiates a new SkillsProfileDetails object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSkillsProfileDetails() *SkillsProfileDetails {
	this := SkillsProfileDetails{}
	return &this
}

// NewSkillsProfileDetailsWithDefaults instantiates a new SkillsProfileDetails object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSkillsProfileDetailsWithDefaults() *SkillsProfileDetails {
	this := SkillsProfileDetails{}
	return &this
}

// GetAudit returns the Audit field value if set, zero value otherwise.
func (o *SkillsProfileDetails) GetAudit() Audit {
	if o == nil || IsNil(o.Audit) {
		var ret Audit
		return ret
	}
	return *o.Audit
}

// GetAuditOk returns a tuple with the Audit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SkillsProfileDetails) GetAuditOk() (*Audit, bool) {
	if o == nil || IsNil(o.Audit) {
		return nil, false
	}
	return o.Audit, true
}

// HasAudit returns a boolean if a field has been set.
func (o *SkillsProfileDetails) HasAudit() bool {
	if o != nil && !IsNil(o.Audit) {
		return true
	}

	return false
}

// SetAudit gets a reference to the given Audit and assigns it to the Audit field.
func (o *SkillsProfileDetails) SetAudit(v Audit) {
	o.Audit = &v
}

// GetSkillsProfile returns the SkillsProfile field value if set, zero value otherwise.
func (o *SkillsProfileDetails) GetSkillsProfile() SkillsProfile {
	if o == nil || IsNil(o.SkillsProfile) {
		var ret SkillsProfile
		return ret
	}
	return *o.SkillsProfile
}

// GetSkillsProfileOk returns a tuple with the SkillsProfile field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SkillsProfileDetails) GetSkillsProfileOk() (*SkillsProfile, bool) {
	if o == nil || IsNil(o.SkillsProfile) {
		return nil, false
	}
	return o.SkillsProfile, true
}

// HasSkillsProfile returns a boolean if a field has been set.
func (o *SkillsProfileDetails) HasSkillsProfile() bool {
	if o != nil && !IsNil(o.SkillsProfile) {
		return true
	}

	return false
}

// SetSkillsProfile gets a reference to the given SkillsProfile and assigns it to the SkillsProfile field.
func (o *SkillsProfileDetails) SetSkillsProfile(v SkillsProfile) {
	o.SkillsProfile = &v
}

// GetPerson returns the Person field value if set, zero value otherwise.
func (o *SkillsProfileDetails) GetPerson() Person {
	if o == nil || IsNil(o.Person) {
		var ret Person
		return ret
	}
	return *o.Person
}

// GetPersonOk returns a tuple with the Person field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SkillsProfileDetails) GetPersonOk() (*Person, bool) {
	if o == nil || IsNil(o.Person) {
		return nil, false
	}
	return o.Person, true
}

// HasPerson returns a boolean if a field has been set.
func (o *SkillsProfileDetails) HasPerson() bool {
	if o != nil && !IsNil(o.Person) {
		return true
	}

	return false
}

// SetPerson gets a reference to the given Person and assigns it to the Person field.
func (o *SkillsProfileDetails) SetPerson(v Person) {
	o.Person = &v
}

// GetSkills returns the Skills field value if set, zero value otherwise.
func (o *SkillsProfileDetails) GetSkills() []ExperienceSkill {
	if o == nil || IsNil(o.Skills) {
		var ret []ExperienceSkill
		return ret
	}
	return o.Skills
}

// GetSkillsOk returns a tuple with the Skills field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SkillsProfileDetails) GetSkillsOk() ([]ExperienceSkill, bool) {
	if o == nil || IsNil(o.Skills) {
		return nil, false
	}
	return o.Skills, true
}

// HasSkills returns a boolean if a field has been set.
func (o *SkillsProfileDetails) HasSkills() bool {
	if o != nil && !IsNil(o.Skills) {
		return true
	}

	return false
}

// SetSkills gets a reference to the given []ExperienceSkill and assigns it to the Skills field.
func (o *SkillsProfileDetails) SetSkills(v []ExperienceSkill) {
	o.Skills = v
}

func (o SkillsProfileDetails) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SkillsProfileDetails) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Audit) {
		toSerialize["audit"] = o.Audit
	}
	if !IsNil(o.SkillsProfile) {
		toSerialize["skillsProfile"] = o.SkillsProfile
	}
	if !IsNil(o.Person) {
		toSerialize["person"] = o.Person
	}
	if !IsNil(o.Skills) {
		toSerialize["skills"] = o.Skills
	}
	return toSerialize, nil
}

type NullableSkillsProfileDetails struct {
	value *SkillsProfileDetails
	isSet bool
}

func (v NullableSkillsProfileDetails) Get() *SkillsProfileDetails {
	return v.value
}

func (v *NullableSkillsProfileDetails) Set(val *SkillsProfileDetails) {
	v.value = val
	v.isSet = true
}

func (v NullableSkillsProfileDetails) IsSet() bool {
	return v.isSet
}

func (v *NullableSkillsProfileDetails) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSkillsProfileDetails(val *SkillsProfileDetails) *NullableSkillsProfileDetails {
	return &NullableSkillsProfileDetails{value: val, isSet: true}
}

func (v NullableSkillsProfileDetails) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSkillsProfileDetails) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


