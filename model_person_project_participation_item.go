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

// checks if the PersonProjectParticipationItem type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PersonProjectParticipationItem{}

// PersonProjectParticipationItem struct for PersonProjectParticipationItem
type PersonProjectParticipationItem struct {
	Participation *ProjectParticipation `json:"participation,omitempty"`
	Experiences []Experience `json:"experiences,omitempty"`
}

// NewPersonProjectParticipationItem instantiates a new PersonProjectParticipationItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPersonProjectParticipationItem() *PersonProjectParticipationItem {
	this := PersonProjectParticipationItem{}
	return &this
}

// NewPersonProjectParticipationItemWithDefaults instantiates a new PersonProjectParticipationItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPersonProjectParticipationItemWithDefaults() *PersonProjectParticipationItem {
	this := PersonProjectParticipationItem{}
	return &this
}

// GetParticipation returns the Participation field value if set, zero value otherwise.
func (o *PersonProjectParticipationItem) GetParticipation() ProjectParticipation {
	if o == nil || IsNil(o.Participation) {
		var ret ProjectParticipation
		return ret
	}
	return *o.Participation
}

// GetParticipationOk returns a tuple with the Participation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PersonProjectParticipationItem) GetParticipationOk() (*ProjectParticipation, bool) {
	if o == nil || IsNil(o.Participation) {
		return nil, false
	}
	return o.Participation, true
}

// HasParticipation returns a boolean if a field has been set.
func (o *PersonProjectParticipationItem) HasParticipation() bool {
	if o != nil && !IsNil(o.Participation) {
		return true
	}

	return false
}

// SetParticipation gets a reference to the given ProjectParticipation and assigns it to the Participation field.
func (o *PersonProjectParticipationItem) SetParticipation(v ProjectParticipation) {
	o.Participation = &v
}

// GetExperiences returns the Experiences field value if set, zero value otherwise.
func (o *PersonProjectParticipationItem) GetExperiences() []Experience {
	if o == nil || IsNil(o.Experiences) {
		var ret []Experience
		return ret
	}
	return o.Experiences
}

// GetExperiencesOk returns a tuple with the Experiences field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PersonProjectParticipationItem) GetExperiencesOk() ([]Experience, bool) {
	if o == nil || IsNil(o.Experiences) {
		return nil, false
	}
	return o.Experiences, true
}

// HasExperiences returns a boolean if a field has been set.
func (o *PersonProjectParticipationItem) HasExperiences() bool {
	if o != nil && !IsNil(o.Experiences) {
		return true
	}

	return false
}

// SetExperiences gets a reference to the given []Experience and assigns it to the Experiences field.
func (o *PersonProjectParticipationItem) SetExperiences(v []Experience) {
	o.Experiences = v
}

func (o PersonProjectParticipationItem) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PersonProjectParticipationItem) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Participation) {
		toSerialize["participation"] = o.Participation
	}
	if !IsNil(o.Experiences) {
		toSerialize["experiences"] = o.Experiences
	}
	return toSerialize, nil
}

type NullablePersonProjectParticipationItem struct {
	value *PersonProjectParticipationItem
	isSet bool
}

func (v NullablePersonProjectParticipationItem) Get() *PersonProjectParticipationItem {
	return v.value
}

func (v *NullablePersonProjectParticipationItem) Set(val *PersonProjectParticipationItem) {
	v.value = val
	v.isSet = true
}

func (v NullablePersonProjectParticipationItem) IsSet() bool {
	return v.isSet
}

func (v *NullablePersonProjectParticipationItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePersonProjectParticipationItem(val *PersonProjectParticipationItem) *NullablePersonProjectParticipationItem {
	return &NullablePersonProjectParticipationItem{value: val, isSet: true}
}

func (v NullablePersonProjectParticipationItem) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePersonProjectParticipationItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


