/*
YASM (Yet Another Skill Management) API

This is an example of using OAuth2 Implicit Flow in a specification to describe security to your API.

API version: 1.50.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// checks if the ProjectParticipationUpdate type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ProjectParticipationUpdate{}

// ProjectParticipationUpdate struct for ProjectParticipationUpdate
type ProjectParticipationUpdate struct {
	Timeframe Timeframed `json:"timeframe"`
	Skills []SkillLevelUpdate `json:"skills,omitempty"`
	PersonalDescription *string `json:"personalDescription,omitempty"`
}

// NewProjectParticipationUpdate instantiates a new ProjectParticipationUpdate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProjectParticipationUpdate(timeframe Timeframed) *ProjectParticipationUpdate {
	this := ProjectParticipationUpdate{}
	this.Timeframe = timeframe
	return &this
}

// NewProjectParticipationUpdateWithDefaults instantiates a new ProjectParticipationUpdate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProjectParticipationUpdateWithDefaults() *ProjectParticipationUpdate {
	this := ProjectParticipationUpdate{}
	return &this
}

// GetTimeframe returns the Timeframe field value
func (o *ProjectParticipationUpdate) GetTimeframe() Timeframed {
	if o == nil {
		var ret Timeframed
		return ret
	}

	return o.Timeframe
}

// GetTimeframeOk returns a tuple with the Timeframe field value
// and a boolean to check if the value has been set.
func (o *ProjectParticipationUpdate) GetTimeframeOk() (*Timeframed, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Timeframe, true
}

// SetTimeframe sets field value
func (o *ProjectParticipationUpdate) SetTimeframe(v Timeframed) {
	o.Timeframe = v
}

// GetSkills returns the Skills field value if set, zero value otherwise.
func (o *ProjectParticipationUpdate) GetSkills() []SkillLevelUpdate {
	if o == nil || IsNil(o.Skills) {
		var ret []SkillLevelUpdate
		return ret
	}
	return o.Skills
}

// GetSkillsOk returns a tuple with the Skills field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectParticipationUpdate) GetSkillsOk() ([]SkillLevelUpdate, bool) {
	if o == nil || IsNil(o.Skills) {
		return nil, false
	}
	return o.Skills, true
}

// HasSkills returns a boolean if a field has been set.
func (o *ProjectParticipationUpdate) HasSkills() bool {
	if o != nil && !IsNil(o.Skills) {
		return true
	}

	return false
}

// SetSkills gets a reference to the given []SkillLevelUpdate and assigns it to the Skills field.
func (o *ProjectParticipationUpdate) SetSkills(v []SkillLevelUpdate) {
	o.Skills = v
}

// GetPersonalDescription returns the PersonalDescription field value if set, zero value otherwise.
func (o *ProjectParticipationUpdate) GetPersonalDescription() string {
	if o == nil || IsNil(o.PersonalDescription) {
		var ret string
		return ret
	}
	return *o.PersonalDescription
}

// GetPersonalDescriptionOk returns a tuple with the PersonalDescription field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectParticipationUpdate) GetPersonalDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.PersonalDescription) {
		return nil, false
	}
	return o.PersonalDescription, true
}

// HasPersonalDescription returns a boolean if a field has been set.
func (o *ProjectParticipationUpdate) HasPersonalDescription() bool {
	if o != nil && !IsNil(o.PersonalDescription) {
		return true
	}

	return false
}

// SetPersonalDescription gets a reference to the given string and assigns it to the PersonalDescription field.
func (o *ProjectParticipationUpdate) SetPersonalDescription(v string) {
	o.PersonalDescription = &v
}

func (o ProjectParticipationUpdate) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ProjectParticipationUpdate) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["timeframe"] = o.Timeframe
	if !IsNil(o.Skills) {
		toSerialize["skills"] = o.Skills
	}
	if !IsNil(o.PersonalDescription) {
		toSerialize["personalDescription"] = o.PersonalDescription
	}
	return toSerialize, nil
}

type NullableProjectParticipationUpdate struct {
	value *ProjectParticipationUpdate
	isSet bool
}

func (v NullableProjectParticipationUpdate) Get() *ProjectParticipationUpdate {
	return v.value
}

func (v *NullableProjectParticipationUpdate) Set(val *ProjectParticipationUpdate) {
	v.value = val
	v.isSet = true
}

func (v NullableProjectParticipationUpdate) IsSet() bool {
	return v.isSet
}

func (v *NullableProjectParticipationUpdate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProjectParticipationUpdate(val *ProjectParticipationUpdate) *NullableProjectParticipationUpdate {
	return &NullableProjectParticipationUpdate{value: val, isSet: true}
}

func (v NullableProjectParticipationUpdate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProjectParticipationUpdate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


