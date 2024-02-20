/*
YASM (Yet Another Skill Management) API

This is an example of using OAuth2 Implicit Flow in a specification to describe security to your API.

API version: 1.13.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// checks if the SkillLevelUpdate type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SkillLevelUpdate{}

// SkillLevelUpdate struct for SkillLevelUpdate
type SkillLevelUpdate struct {
	SkillId *string `json:"skillId,omitempty"`
	Level *Level `json:"level,omitempty"`
}

// NewSkillLevelUpdate instantiates a new SkillLevelUpdate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSkillLevelUpdate() *SkillLevelUpdate {
	this := SkillLevelUpdate{}
	return &this
}

// NewSkillLevelUpdateWithDefaults instantiates a new SkillLevelUpdate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSkillLevelUpdateWithDefaults() *SkillLevelUpdate {
	this := SkillLevelUpdate{}
	return &this
}

// GetSkillId returns the SkillId field value if set, zero value otherwise.
func (o *SkillLevelUpdate) GetSkillId() string {
	if o == nil || IsNil(o.SkillId) {
		var ret string
		return ret
	}
	return *o.SkillId
}

// GetSkillIdOk returns a tuple with the SkillId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SkillLevelUpdate) GetSkillIdOk() (*string, bool) {
	if o == nil || IsNil(o.SkillId) {
		return nil, false
	}
	return o.SkillId, true
}

// HasSkillId returns a boolean if a field has been set.
func (o *SkillLevelUpdate) HasSkillId() bool {
	if o != nil && !IsNil(o.SkillId) {
		return true
	}

	return false
}

// SetSkillId gets a reference to the given string and assigns it to the SkillId field.
func (o *SkillLevelUpdate) SetSkillId(v string) {
	o.SkillId = &v
}

// GetLevel returns the Level field value if set, zero value otherwise.
func (o *SkillLevelUpdate) GetLevel() Level {
	if o == nil || IsNil(o.Level) {
		var ret Level
		return ret
	}
	return *o.Level
}

// GetLevelOk returns a tuple with the Level field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SkillLevelUpdate) GetLevelOk() (*Level, bool) {
	if o == nil || IsNil(o.Level) {
		return nil, false
	}
	return o.Level, true
}

// HasLevel returns a boolean if a field has been set.
func (o *SkillLevelUpdate) HasLevel() bool {
	if o != nil && !IsNil(o.Level) {
		return true
	}

	return false
}

// SetLevel gets a reference to the given Level and assigns it to the Level field.
func (o *SkillLevelUpdate) SetLevel(v Level) {
	o.Level = &v
}

func (o SkillLevelUpdate) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SkillLevelUpdate) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.SkillId) {
		toSerialize["skillId"] = o.SkillId
	}
	if !IsNil(o.Level) {
		toSerialize["level"] = o.Level
	}
	return toSerialize, nil
}

type NullableSkillLevelUpdate struct {
	value *SkillLevelUpdate
	isSet bool
}

func (v NullableSkillLevelUpdate) Get() *SkillLevelUpdate {
	return v.value
}

func (v *NullableSkillLevelUpdate) Set(val *SkillLevelUpdate) {
	v.value = val
	v.isSet = true
}

func (v NullableSkillLevelUpdate) IsSet() bool {
	return v.isSet
}

func (v *NullableSkillLevelUpdate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSkillLevelUpdate(val *SkillLevelUpdate) *NullableSkillLevelUpdate {
	return &NullableSkillLevelUpdate{value: val, isSet: true}
}

func (v NullableSkillLevelUpdate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSkillLevelUpdate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


