/*
YASM (Yet Another Skill Management) API

This is an example of using OAuth2 Implicit Flow in a specification to describe security to your API.

API version: 1.3.18
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// SkillDetails struct for SkillDetails
type SkillDetails struct {
	Skill *Skill `json:"skill,omitempty"`
	Children []Skill `json:"children,omitempty"`
	Parents []Skill `json:"parents,omitempty"`
	Kinds []Skill `json:"kinds,omitempty"`
	// The display name of the skill, e.g. \"Communication (Computing)\" or \"Communication (Design)\" or \"Protcool (Communication,IT)\"
	DisplayName *string `json:"displayName,omitempty"`
}

// NewSkillDetails instantiates a new SkillDetails object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSkillDetails() *SkillDetails {
	this := SkillDetails{}
	return &this
}

// NewSkillDetailsWithDefaults instantiates a new SkillDetails object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSkillDetailsWithDefaults() *SkillDetails {
	this := SkillDetails{}
	return &this
}

// GetSkill returns the Skill field value if set, zero value otherwise.
func (o *SkillDetails) GetSkill() Skill {
	if o == nil || o.Skill == nil {
		var ret Skill
		return ret
	}
	return *o.Skill
}

// GetSkillOk returns a tuple with the Skill field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SkillDetails) GetSkillOk() (*Skill, bool) {
	if o == nil || o.Skill == nil {
		return nil, false
	}
	return o.Skill, true
}

// HasSkill returns a boolean if a field has been set.
func (o *SkillDetails) HasSkill() bool {
	if o != nil && o.Skill != nil {
		return true
	}

	return false
}

// SetSkill gets a reference to the given Skill and assigns it to the Skill field.
func (o *SkillDetails) SetSkill(v Skill) {
	o.Skill = &v
}

// GetChildren returns the Children field value if set, zero value otherwise.
func (o *SkillDetails) GetChildren() []Skill {
	if o == nil || o.Children == nil {
		var ret []Skill
		return ret
	}
	return o.Children
}

// GetChildrenOk returns a tuple with the Children field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SkillDetails) GetChildrenOk() ([]Skill, bool) {
	if o == nil || o.Children == nil {
		return nil, false
	}
	return o.Children, true
}

// HasChildren returns a boolean if a field has been set.
func (o *SkillDetails) HasChildren() bool {
	if o != nil && o.Children != nil {
		return true
	}

	return false
}

// SetChildren gets a reference to the given []Skill and assigns it to the Children field.
func (o *SkillDetails) SetChildren(v []Skill) {
	o.Children = v
}

// GetParents returns the Parents field value if set, zero value otherwise.
func (o *SkillDetails) GetParents() []Skill {
	if o == nil || o.Parents == nil {
		var ret []Skill
		return ret
	}
	return o.Parents
}

// GetParentsOk returns a tuple with the Parents field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SkillDetails) GetParentsOk() ([]Skill, bool) {
	if o == nil || o.Parents == nil {
		return nil, false
	}
	return o.Parents, true
}

// HasParents returns a boolean if a field has been set.
func (o *SkillDetails) HasParents() bool {
	if o != nil && o.Parents != nil {
		return true
	}

	return false
}

// SetParents gets a reference to the given []Skill and assigns it to the Parents field.
func (o *SkillDetails) SetParents(v []Skill) {
	o.Parents = v
}

// GetKinds returns the Kinds field value if set, zero value otherwise.
func (o *SkillDetails) GetKinds() []Skill {
	if o == nil || o.Kinds == nil {
		var ret []Skill
		return ret
	}
	return o.Kinds
}

// GetKindsOk returns a tuple with the Kinds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SkillDetails) GetKindsOk() ([]Skill, bool) {
	if o == nil || o.Kinds == nil {
		return nil, false
	}
	return o.Kinds, true
}

// HasKinds returns a boolean if a field has been set.
func (o *SkillDetails) HasKinds() bool {
	if o != nil && o.Kinds != nil {
		return true
	}

	return false
}

// SetKinds gets a reference to the given []Skill and assigns it to the Kinds field.
func (o *SkillDetails) SetKinds(v []Skill) {
	o.Kinds = v
}

// GetDisplayName returns the DisplayName field value if set, zero value otherwise.
func (o *SkillDetails) GetDisplayName() string {
	if o == nil || o.DisplayName == nil {
		var ret string
		return ret
	}
	return *o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SkillDetails) GetDisplayNameOk() (*string, bool) {
	if o == nil || o.DisplayName == nil {
		return nil, false
	}
	return o.DisplayName, true
}

// HasDisplayName returns a boolean if a field has been set.
func (o *SkillDetails) HasDisplayName() bool {
	if o != nil && o.DisplayName != nil {
		return true
	}

	return false
}

// SetDisplayName gets a reference to the given string and assigns it to the DisplayName field.
func (o *SkillDetails) SetDisplayName(v string) {
	o.DisplayName = &v
}

func (o SkillDetails) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Skill != nil {
		toSerialize["skill"] = o.Skill
	}
	if o.Children != nil {
		toSerialize["children"] = o.Children
	}
	if o.Parents != nil {
		toSerialize["parents"] = o.Parents
	}
	if o.Kinds != nil {
		toSerialize["kinds"] = o.Kinds
	}
	if o.DisplayName != nil {
		toSerialize["displayName"] = o.DisplayName
	}
	return json.Marshal(toSerialize)
}

type NullableSkillDetails struct {
	value *SkillDetails
	isSet bool
}

func (v NullableSkillDetails) Get() *SkillDetails {
	return v.value
}

func (v *NullableSkillDetails) Set(val *SkillDetails) {
	v.value = val
	v.isSet = true
}

func (v NullableSkillDetails) IsSet() bool {
	return v.isSet
}

func (v *NullableSkillDetails) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSkillDetails(val *SkillDetails) *NullableSkillDetails {
	return &NullableSkillDetails{value: val, isSet: true}
}

func (v NullableSkillDetails) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSkillDetails) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


