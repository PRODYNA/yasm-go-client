/*
YASM (Yet Another Skill Management) API

This is an example of using OAuth2 Implicit Flow in a specification to describe security to your API.

API version: 1.64.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// checks if the SkillsProfileRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SkillsProfileRequest{}

// SkillsProfileRequest struct for SkillsProfileRequest
type SkillsProfileRequest struct {
	SkillsProfile *SkillsProfile `json:"skillsProfile,omitempty"`
	SkillIds []string `json:"skillIds,omitempty"`
	MethodologyIds []string `json:"methodologyIds,omitempty"`
}

// NewSkillsProfileRequest instantiates a new SkillsProfileRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSkillsProfileRequest() *SkillsProfileRequest {
	this := SkillsProfileRequest{}
	return &this
}

// NewSkillsProfileRequestWithDefaults instantiates a new SkillsProfileRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSkillsProfileRequestWithDefaults() *SkillsProfileRequest {
	this := SkillsProfileRequest{}
	return &this
}

// GetSkillsProfile returns the SkillsProfile field value if set, zero value otherwise.
func (o *SkillsProfileRequest) GetSkillsProfile() SkillsProfile {
	if o == nil || IsNil(o.SkillsProfile) {
		var ret SkillsProfile
		return ret
	}
	return *o.SkillsProfile
}

// GetSkillsProfileOk returns a tuple with the SkillsProfile field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SkillsProfileRequest) GetSkillsProfileOk() (*SkillsProfile, bool) {
	if o == nil || IsNil(o.SkillsProfile) {
		return nil, false
	}
	return o.SkillsProfile, true
}

// HasSkillsProfile returns a boolean if a field has been set.
func (o *SkillsProfileRequest) HasSkillsProfile() bool {
	if o != nil && !IsNil(o.SkillsProfile) {
		return true
	}

	return false
}

// SetSkillsProfile gets a reference to the given SkillsProfile and assigns it to the SkillsProfile field.
func (o *SkillsProfileRequest) SetSkillsProfile(v SkillsProfile) {
	o.SkillsProfile = &v
}

// GetSkillIds returns the SkillIds field value if set, zero value otherwise.
func (o *SkillsProfileRequest) GetSkillIds() []string {
	if o == nil || IsNil(o.SkillIds) {
		var ret []string
		return ret
	}
	return o.SkillIds
}

// GetSkillIdsOk returns a tuple with the SkillIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SkillsProfileRequest) GetSkillIdsOk() ([]string, bool) {
	if o == nil || IsNil(o.SkillIds) {
		return nil, false
	}
	return o.SkillIds, true
}

// HasSkillIds returns a boolean if a field has been set.
func (o *SkillsProfileRequest) HasSkillIds() bool {
	if o != nil && !IsNil(o.SkillIds) {
		return true
	}

	return false
}

// SetSkillIds gets a reference to the given []string and assigns it to the SkillIds field.
func (o *SkillsProfileRequest) SetSkillIds(v []string) {
	o.SkillIds = v
}

// GetMethodologyIds returns the MethodologyIds field value if set, zero value otherwise.
func (o *SkillsProfileRequest) GetMethodologyIds() []string {
	if o == nil || IsNil(o.MethodologyIds) {
		var ret []string
		return ret
	}
	return o.MethodologyIds
}

// GetMethodologyIdsOk returns a tuple with the MethodologyIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SkillsProfileRequest) GetMethodologyIdsOk() ([]string, bool) {
	if o == nil || IsNil(o.MethodologyIds) {
		return nil, false
	}
	return o.MethodologyIds, true
}

// HasMethodologyIds returns a boolean if a field has been set.
func (o *SkillsProfileRequest) HasMethodologyIds() bool {
	if o != nil && !IsNil(o.MethodologyIds) {
		return true
	}

	return false
}

// SetMethodologyIds gets a reference to the given []string and assigns it to the MethodologyIds field.
func (o *SkillsProfileRequest) SetMethodologyIds(v []string) {
	o.MethodologyIds = v
}

func (o SkillsProfileRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SkillsProfileRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.SkillsProfile) {
		toSerialize["skillsProfile"] = o.SkillsProfile
	}
	if !IsNil(o.SkillIds) {
		toSerialize["skillIds"] = o.SkillIds
	}
	if !IsNil(o.MethodologyIds) {
		toSerialize["methodologyIds"] = o.MethodologyIds
	}
	return toSerialize, nil
}

type NullableSkillsProfileRequest struct {
	value *SkillsProfileRequest
	isSet bool
}

func (v NullableSkillsProfileRequest) Get() *SkillsProfileRequest {
	return v.value
}

func (v *NullableSkillsProfileRequest) Set(val *SkillsProfileRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableSkillsProfileRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableSkillsProfileRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSkillsProfileRequest(val *SkillsProfileRequest) *NullableSkillsProfileRequest {
	return &NullableSkillsProfileRequest{value: val, isSet: true}
}

func (v NullableSkillsProfileRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSkillsProfileRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


