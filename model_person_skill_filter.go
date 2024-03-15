/*
YASM (Yet Another Skill Management) API

This is an example of using OAuth2 Implicit Flow in a specification to describe security to your API.

API version: 1.15.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// checks if the PersonSkillFilter type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PersonSkillFilter{}

// PersonSkillFilter struct for PersonSkillFilter
type PersonSkillFilter struct {
	Id string `json:"id"`
	ObjectType *string `json:"objectType,omitempty"`
	ExperienceInMonth *MinMax `json:"experienceInMonth,omitempty"`
	// filters the last time the employee used the skill in a project
	LastAssignment *string `json:"lastAssignment,omitempty"`
}

// NewPersonSkillFilter instantiates a new PersonSkillFilter object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPersonSkillFilter(id string) *PersonSkillFilter {
	this := PersonSkillFilter{}
	this.Id = id
	return &this
}

// NewPersonSkillFilterWithDefaults instantiates a new PersonSkillFilter object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPersonSkillFilterWithDefaults() *PersonSkillFilter {
	this := PersonSkillFilter{}
	return &this
}

// GetId returns the Id field value
func (o *PersonSkillFilter) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *PersonSkillFilter) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *PersonSkillFilter) SetId(v string) {
	o.Id = v
}

// GetObjectType returns the ObjectType field value if set, zero value otherwise.
func (o *PersonSkillFilter) GetObjectType() string {
	if o == nil || IsNil(o.ObjectType) {
		var ret string
		return ret
	}
	return *o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PersonSkillFilter) GetObjectTypeOk() (*string, bool) {
	if o == nil || IsNil(o.ObjectType) {
		return nil, false
	}
	return o.ObjectType, true
}

// HasObjectType returns a boolean if a field has been set.
func (o *PersonSkillFilter) HasObjectType() bool {
	if o != nil && !IsNil(o.ObjectType) {
		return true
	}

	return false
}

// SetObjectType gets a reference to the given string and assigns it to the ObjectType field.
func (o *PersonSkillFilter) SetObjectType(v string) {
	o.ObjectType = &v
}

// GetExperienceInMonth returns the ExperienceInMonth field value if set, zero value otherwise.
func (o *PersonSkillFilter) GetExperienceInMonth() MinMax {
	if o == nil || IsNil(o.ExperienceInMonth) {
		var ret MinMax
		return ret
	}
	return *o.ExperienceInMonth
}

// GetExperienceInMonthOk returns a tuple with the ExperienceInMonth field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PersonSkillFilter) GetExperienceInMonthOk() (*MinMax, bool) {
	if o == nil || IsNil(o.ExperienceInMonth) {
		return nil, false
	}
	return o.ExperienceInMonth, true
}

// HasExperienceInMonth returns a boolean if a field has been set.
func (o *PersonSkillFilter) HasExperienceInMonth() bool {
	if o != nil && !IsNil(o.ExperienceInMonth) {
		return true
	}

	return false
}

// SetExperienceInMonth gets a reference to the given MinMax and assigns it to the ExperienceInMonth field.
func (o *PersonSkillFilter) SetExperienceInMonth(v MinMax) {
	o.ExperienceInMonth = &v
}

// GetLastAssignment returns the LastAssignment field value if set, zero value otherwise.
func (o *PersonSkillFilter) GetLastAssignment() string {
	if o == nil || IsNil(o.LastAssignment) {
		var ret string
		return ret
	}
	return *o.LastAssignment
}

// GetLastAssignmentOk returns a tuple with the LastAssignment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PersonSkillFilter) GetLastAssignmentOk() (*string, bool) {
	if o == nil || IsNil(o.LastAssignment) {
		return nil, false
	}
	return o.LastAssignment, true
}

// HasLastAssignment returns a boolean if a field has been set.
func (o *PersonSkillFilter) HasLastAssignment() bool {
	if o != nil && !IsNil(o.LastAssignment) {
		return true
	}

	return false
}

// SetLastAssignment gets a reference to the given string and assigns it to the LastAssignment field.
func (o *PersonSkillFilter) SetLastAssignment(v string) {
	o.LastAssignment = &v
}

func (o PersonSkillFilter) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PersonSkillFilter) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	if !IsNil(o.ObjectType) {
		toSerialize["objectType"] = o.ObjectType
	}
	if !IsNil(o.ExperienceInMonth) {
		toSerialize["experienceInMonth"] = o.ExperienceInMonth
	}
	if !IsNil(o.LastAssignment) {
		toSerialize["lastAssignment"] = o.LastAssignment
	}
	return toSerialize, nil
}

type NullablePersonSkillFilter struct {
	value *PersonSkillFilter
	isSet bool
}

func (v NullablePersonSkillFilter) Get() *PersonSkillFilter {
	return v.value
}

func (v *NullablePersonSkillFilter) Set(val *PersonSkillFilter) {
	v.value = val
	v.isSet = true
}

func (v NullablePersonSkillFilter) IsSet() bool {
	return v.isSet
}

func (v *NullablePersonSkillFilter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePersonSkillFilter(val *PersonSkillFilter) *NullablePersonSkillFilter {
	return &NullablePersonSkillFilter{value: val, isSet: true}
}

func (v NullablePersonSkillFilter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePersonSkillFilter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


