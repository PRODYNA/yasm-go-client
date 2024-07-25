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

// checks if the SkillDetails type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SkillDetails{}

// SkillDetails struct for SkillDetails
type SkillDetails struct {
	Audit *Audit `json:"audit,omitempty"`
	Skill *Skill `json:"skill,omitempty"`
	ParentPath []Skill `json:"parentPath,omitempty"`
	Children []SkillLink `json:"children,omitempty"`
	AdditionalParents []Skill `json:"additionalParents,omitempty"`
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

// GetAudit returns the Audit field value if set, zero value otherwise.
func (o *SkillDetails) GetAudit() Audit {
	if o == nil || IsNil(o.Audit) {
		var ret Audit
		return ret
	}
	return *o.Audit
}

// GetAuditOk returns a tuple with the Audit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SkillDetails) GetAuditOk() (*Audit, bool) {
	if o == nil || IsNil(o.Audit) {
		return nil, false
	}
	return o.Audit, true
}

// HasAudit returns a boolean if a field has been set.
func (o *SkillDetails) HasAudit() bool {
	if o != nil && !IsNil(o.Audit) {
		return true
	}

	return false
}

// SetAudit gets a reference to the given Audit and assigns it to the Audit field.
func (o *SkillDetails) SetAudit(v Audit) {
	o.Audit = &v
}

// GetSkill returns the Skill field value if set, zero value otherwise.
func (o *SkillDetails) GetSkill() Skill {
	if o == nil || IsNil(o.Skill) {
		var ret Skill
		return ret
	}
	return *o.Skill
}

// GetSkillOk returns a tuple with the Skill field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SkillDetails) GetSkillOk() (*Skill, bool) {
	if o == nil || IsNil(o.Skill) {
		return nil, false
	}
	return o.Skill, true
}

// HasSkill returns a boolean if a field has been set.
func (o *SkillDetails) HasSkill() bool {
	if o != nil && !IsNil(o.Skill) {
		return true
	}

	return false
}

// SetSkill gets a reference to the given Skill and assigns it to the Skill field.
func (o *SkillDetails) SetSkill(v Skill) {
	o.Skill = &v
}

// GetParentPath returns the ParentPath field value if set, zero value otherwise.
func (o *SkillDetails) GetParentPath() []Skill {
	if o == nil || IsNil(o.ParentPath) {
		var ret []Skill
		return ret
	}
	return o.ParentPath
}

// GetParentPathOk returns a tuple with the ParentPath field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SkillDetails) GetParentPathOk() ([]Skill, bool) {
	if o == nil || IsNil(o.ParentPath) {
		return nil, false
	}
	return o.ParentPath, true
}

// HasParentPath returns a boolean if a field has been set.
func (o *SkillDetails) HasParentPath() bool {
	if o != nil && !IsNil(o.ParentPath) {
		return true
	}

	return false
}

// SetParentPath gets a reference to the given []Skill and assigns it to the ParentPath field.
func (o *SkillDetails) SetParentPath(v []Skill) {
	o.ParentPath = v
}

// GetChildren returns the Children field value if set, zero value otherwise.
func (o *SkillDetails) GetChildren() []SkillLink {
	if o == nil || IsNil(o.Children) {
		var ret []SkillLink
		return ret
	}
	return o.Children
}

// GetChildrenOk returns a tuple with the Children field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SkillDetails) GetChildrenOk() ([]SkillLink, bool) {
	if o == nil || IsNil(o.Children) {
		return nil, false
	}
	return o.Children, true
}

// HasChildren returns a boolean if a field has been set.
func (o *SkillDetails) HasChildren() bool {
	if o != nil && !IsNil(o.Children) {
		return true
	}

	return false
}

// SetChildren gets a reference to the given []SkillLink and assigns it to the Children field.
func (o *SkillDetails) SetChildren(v []SkillLink) {
	o.Children = v
}

// GetAdditionalParents returns the AdditionalParents field value if set, zero value otherwise.
func (o *SkillDetails) GetAdditionalParents() []Skill {
	if o == nil || IsNil(o.AdditionalParents) {
		var ret []Skill
		return ret
	}
	return o.AdditionalParents
}

// GetAdditionalParentsOk returns a tuple with the AdditionalParents field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SkillDetails) GetAdditionalParentsOk() ([]Skill, bool) {
	if o == nil || IsNil(o.AdditionalParents) {
		return nil, false
	}
	return o.AdditionalParents, true
}

// HasAdditionalParents returns a boolean if a field has been set.
func (o *SkillDetails) HasAdditionalParents() bool {
	if o != nil && !IsNil(o.AdditionalParents) {
		return true
	}

	return false
}

// SetAdditionalParents gets a reference to the given []Skill and assigns it to the AdditionalParents field.
func (o *SkillDetails) SetAdditionalParents(v []Skill) {
	o.AdditionalParents = v
}

func (o SkillDetails) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SkillDetails) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Audit) {
		toSerialize["audit"] = o.Audit
	}
	if !IsNil(o.Skill) {
		toSerialize["skill"] = o.Skill
	}
	if !IsNil(o.ParentPath) {
		toSerialize["parentPath"] = o.ParentPath
	}
	if !IsNil(o.Children) {
		toSerialize["children"] = o.Children
	}
	if !IsNil(o.AdditionalParents) {
		toSerialize["additionalParents"] = o.AdditionalParents
	}
	return toSerialize, nil
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


