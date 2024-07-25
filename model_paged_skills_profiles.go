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

// checks if the PagedSkillsProfiles type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PagedSkillsProfiles{}

// PagedSkillsProfiles struct for PagedSkillsProfiles
type PagedSkillsProfiles struct {
	Page
	SkillsProfiles []SkillsProfileDetails `json:"skillsProfiles,omitempty"`
}

// NewPagedSkillsProfiles instantiates a new PagedSkillsProfiles object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPagedSkillsProfiles() *PagedSkillsProfiles {
	this := PagedSkillsProfiles{}
	return &this
}

// NewPagedSkillsProfilesWithDefaults instantiates a new PagedSkillsProfiles object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPagedSkillsProfilesWithDefaults() *PagedSkillsProfiles {
	this := PagedSkillsProfiles{}
	return &this
}

// GetSkillsProfiles returns the SkillsProfiles field value if set, zero value otherwise.
func (o *PagedSkillsProfiles) GetSkillsProfiles() []SkillsProfileDetails {
	if o == nil || IsNil(o.SkillsProfiles) {
		var ret []SkillsProfileDetails
		return ret
	}
	return o.SkillsProfiles
}

// GetSkillsProfilesOk returns a tuple with the SkillsProfiles field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PagedSkillsProfiles) GetSkillsProfilesOk() ([]SkillsProfileDetails, bool) {
	if o == nil || IsNil(o.SkillsProfiles) {
		return nil, false
	}
	return o.SkillsProfiles, true
}

// HasSkillsProfiles returns a boolean if a field has been set.
func (o *PagedSkillsProfiles) HasSkillsProfiles() bool {
	if o != nil && !IsNil(o.SkillsProfiles) {
		return true
	}

	return false
}

// SetSkillsProfiles gets a reference to the given []SkillsProfileDetails and assigns it to the SkillsProfiles field.
func (o *PagedSkillsProfiles) SetSkillsProfiles(v []SkillsProfileDetails) {
	o.SkillsProfiles = v
}

func (o PagedSkillsProfiles) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PagedSkillsProfiles) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	serializedPage, errPage := json.Marshal(o.Page)
	if errPage != nil {
		return map[string]interface{}{}, errPage
	}
	errPage = json.Unmarshal([]byte(serializedPage), &toSerialize)
	if errPage != nil {
		return map[string]interface{}{}, errPage
	}
	if !IsNil(o.SkillsProfiles) {
		toSerialize["skillsProfiles"] = o.SkillsProfiles
	}
	return toSerialize, nil
}

type NullablePagedSkillsProfiles struct {
	value *PagedSkillsProfiles
	isSet bool
}

func (v NullablePagedSkillsProfiles) Get() *PagedSkillsProfiles {
	return v.value
}

func (v *NullablePagedSkillsProfiles) Set(val *PagedSkillsProfiles) {
	v.value = val
	v.isSet = true
}

func (v NullablePagedSkillsProfiles) IsSet() bool {
	return v.isSet
}

func (v *NullablePagedSkillsProfiles) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePagedSkillsProfiles(val *PagedSkillsProfiles) *NullablePagedSkillsProfiles {
	return &NullablePagedSkillsProfiles{value: val, isSet: true}
}

func (v NullablePagedSkillsProfiles) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePagedSkillsProfiles) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


