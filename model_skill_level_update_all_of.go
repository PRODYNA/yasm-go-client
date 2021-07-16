/*
 * YASM (Yet Another Skill Management) API
 *
 * This is an example of using OAuth2 Implicit Flow in a specification to describe security to your API.
 *
 * API version: 0.7.9
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// SkillLevelUpdateAllOf struct for SkillLevelUpdateAllOf
type SkillLevelUpdateAllOf struct {
	SkillId *string `json:"skillId,omitempty"`
	Level *Level `json:"level,omitempty"`
}

// NewSkillLevelUpdateAllOf instantiates a new SkillLevelUpdateAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSkillLevelUpdateAllOf() *SkillLevelUpdateAllOf {
	this := SkillLevelUpdateAllOf{}
	return &this
}

// NewSkillLevelUpdateAllOfWithDefaults instantiates a new SkillLevelUpdateAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSkillLevelUpdateAllOfWithDefaults() *SkillLevelUpdateAllOf {
	this := SkillLevelUpdateAllOf{}
	return &this
}

// GetSkillId returns the SkillId field value if set, zero value otherwise.
func (o *SkillLevelUpdateAllOf) GetSkillId() string {
	if o == nil || o.SkillId == nil {
		var ret string
		return ret
	}
	return *o.SkillId
}

// GetSkillIdOk returns a tuple with the SkillId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SkillLevelUpdateAllOf) GetSkillIdOk() (*string, bool) {
	if o == nil || o.SkillId == nil {
		return nil, false
	}
	return o.SkillId, true
}

// HasSkillId returns a boolean if a field has been set.
func (o *SkillLevelUpdateAllOf) HasSkillId() bool {
	if o != nil && o.SkillId != nil {
		return true
	}

	return false
}

// SetSkillId gets a reference to the given string and assigns it to the SkillId field.
func (o *SkillLevelUpdateAllOf) SetSkillId(v string) {
	o.SkillId = &v
}

// GetLevel returns the Level field value if set, zero value otherwise.
func (o *SkillLevelUpdateAllOf) GetLevel() Level {
	if o == nil || o.Level == nil {
		var ret Level
		return ret
	}
	return *o.Level
}

// GetLevelOk returns a tuple with the Level field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SkillLevelUpdateAllOf) GetLevelOk() (*Level, bool) {
	if o == nil || o.Level == nil {
		return nil, false
	}
	return o.Level, true
}

// HasLevel returns a boolean if a field has been set.
func (o *SkillLevelUpdateAllOf) HasLevel() bool {
	if o != nil && o.Level != nil {
		return true
	}

	return false
}

// SetLevel gets a reference to the given Level and assigns it to the Level field.
func (o *SkillLevelUpdateAllOf) SetLevel(v Level) {
	o.Level = &v
}

func (o SkillLevelUpdateAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.SkillId != nil {
		toSerialize["skillId"] = o.SkillId
	}
	if o.Level != nil {
		toSerialize["level"] = o.Level
	}
	return json.Marshal(toSerialize)
}

type NullableSkillLevelUpdateAllOf struct {
	value *SkillLevelUpdateAllOf
	isSet bool
}

func (v NullableSkillLevelUpdateAllOf) Get() *SkillLevelUpdateAllOf {
	return v.value
}

func (v *NullableSkillLevelUpdateAllOf) Set(val *SkillLevelUpdateAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableSkillLevelUpdateAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableSkillLevelUpdateAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSkillLevelUpdateAllOf(val *SkillLevelUpdateAllOf) *NullableSkillLevelUpdateAllOf {
	return &NullableSkillLevelUpdateAllOf{value: val, isSet: true}
}

func (v NullableSkillLevelUpdateAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSkillLevelUpdateAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


