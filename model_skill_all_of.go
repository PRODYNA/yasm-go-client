/*
YASM (Yet Another Skill Management) API

This is an example of using OAuth2 Implicit Flow in a specification to describe security to your API.

API version: 1.3.13
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// SkillAllOf struct for SkillAllOf
type SkillAllOf struct {
	Invest *bool `json:"invest,omitempty"`
	Kindgiver *bool `json:"kindgiver,omitempty"`
}

// NewSkillAllOf instantiates a new SkillAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSkillAllOf() *SkillAllOf {
	this := SkillAllOf{}
	var invest bool = false
	this.Invest = &invest
	var kindgiver bool = false
	this.Kindgiver = &kindgiver
	return &this
}

// NewSkillAllOfWithDefaults instantiates a new SkillAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSkillAllOfWithDefaults() *SkillAllOf {
	this := SkillAllOf{}
	var invest bool = false
	this.Invest = &invest
	var kindgiver bool = false
	this.Kindgiver = &kindgiver
	return &this
}

// GetInvest returns the Invest field value if set, zero value otherwise.
func (o *SkillAllOf) GetInvest() bool {
	if o == nil || o.Invest == nil {
		var ret bool
		return ret
	}
	return *o.Invest
}

// GetInvestOk returns a tuple with the Invest field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SkillAllOf) GetInvestOk() (*bool, bool) {
	if o == nil || o.Invest == nil {
		return nil, false
	}
	return o.Invest, true
}

// HasInvest returns a boolean if a field has been set.
func (o *SkillAllOf) HasInvest() bool {
	if o != nil && o.Invest != nil {
		return true
	}

	return false
}

// SetInvest gets a reference to the given bool and assigns it to the Invest field.
func (o *SkillAllOf) SetInvest(v bool) {
	o.Invest = &v
}

// GetKindgiver returns the Kindgiver field value if set, zero value otherwise.
func (o *SkillAllOf) GetKindgiver() bool {
	if o == nil || o.Kindgiver == nil {
		var ret bool
		return ret
	}
	return *o.Kindgiver
}

// GetKindgiverOk returns a tuple with the Kindgiver field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SkillAllOf) GetKindgiverOk() (*bool, bool) {
	if o == nil || o.Kindgiver == nil {
		return nil, false
	}
	return o.Kindgiver, true
}

// HasKindgiver returns a boolean if a field has been set.
func (o *SkillAllOf) HasKindgiver() bool {
	if o != nil && o.Kindgiver != nil {
		return true
	}

	return false
}

// SetKindgiver gets a reference to the given bool and assigns it to the Kindgiver field.
func (o *SkillAllOf) SetKindgiver(v bool) {
	o.Kindgiver = &v
}

func (o SkillAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Invest != nil {
		toSerialize["invest"] = o.Invest
	}
	if o.Kindgiver != nil {
		toSerialize["kindgiver"] = o.Kindgiver
	}
	return json.Marshal(toSerialize)
}

type NullableSkillAllOf struct {
	value *SkillAllOf
	isSet bool
}

func (v NullableSkillAllOf) Get() *SkillAllOf {
	return v.value
}

func (v *NullableSkillAllOf) Set(val *SkillAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableSkillAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableSkillAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSkillAllOf(val *SkillAllOf) *NullableSkillAllOf {
	return &NullableSkillAllOf{value: val, isSet: true}
}

func (v NullableSkillAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSkillAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


