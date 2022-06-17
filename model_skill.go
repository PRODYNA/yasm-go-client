/*
YASM (Yet Another Skill Management) API

This is an example of using OAuth2 Implicit Flow in a specification to describe security to your API.

API version: 1.0.16
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// Skill struct for Skill
type Skill struct {
	Id string `json:"id"`
	Name string `json:"name"`
	Suggestion bool `json:"suggestion"`
	// The entity can be linked
	Linkable *bool `json:"linkable,omitempty"`
	Synonyms []string `json:"synonyms,omitempty"`
	Description *string `json:"description,omitempty"`
	Invest *bool `json:"invest,omitempty"`
	Kindgiver *bool `json:"kindgiver,omitempty"`
}

// NewSkill instantiates a new Skill object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSkill(id string, name string, suggestion bool) *Skill {
	this := Skill{}
	this.Id = id
	this.Name = name
	this.Suggestion = suggestion
	var linkable bool = false
	this.Linkable = &linkable
	var invest bool = false
	this.Invest = &invest
	var kindgiver bool = false
	this.Kindgiver = &kindgiver
	return &this
}

// NewSkillWithDefaults instantiates a new Skill object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSkillWithDefaults() *Skill {
	this := Skill{}
	var suggestion bool = false
	this.Suggestion = suggestion
	var linkable bool = false
	this.Linkable = &linkable
	var invest bool = false
	this.Invest = &invest
	var kindgiver bool = false
	this.Kindgiver = &kindgiver
	return &this
}

// GetId returns the Id field value
func (o *Skill) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *Skill) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *Skill) SetId(v string) {
	o.Id = v
}

// GetName returns the Name field value
func (o *Skill) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *Skill) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *Skill) SetName(v string) {
	o.Name = v
}

// GetSuggestion returns the Suggestion field value
func (o *Skill) GetSuggestion() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Suggestion
}

// GetSuggestionOk returns a tuple with the Suggestion field value
// and a boolean to check if the value has been set.
func (o *Skill) GetSuggestionOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Suggestion, true
}

// SetSuggestion sets field value
func (o *Skill) SetSuggestion(v bool) {
	o.Suggestion = v
}

// GetLinkable returns the Linkable field value if set, zero value otherwise.
func (o *Skill) GetLinkable() bool {
	if o == nil || o.Linkable == nil {
		var ret bool
		return ret
	}
	return *o.Linkable
}

// GetLinkableOk returns a tuple with the Linkable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Skill) GetLinkableOk() (*bool, bool) {
	if o == nil || o.Linkable == nil {
		return nil, false
	}
	return o.Linkable, true
}

// HasLinkable returns a boolean if a field has been set.
func (o *Skill) HasLinkable() bool {
	if o != nil && o.Linkable != nil {
		return true
	}

	return false
}

// SetLinkable gets a reference to the given bool and assigns it to the Linkable field.
func (o *Skill) SetLinkable(v bool) {
	o.Linkable = &v
}

// GetSynonyms returns the Synonyms field value if set, zero value otherwise.
func (o *Skill) GetSynonyms() []string {
	if o == nil || o.Synonyms == nil {
		var ret []string
		return ret
	}
	return o.Synonyms
}

// GetSynonymsOk returns a tuple with the Synonyms field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Skill) GetSynonymsOk() ([]string, bool) {
	if o == nil || o.Synonyms == nil {
		return nil, false
	}
	return o.Synonyms, true
}

// HasSynonyms returns a boolean if a field has been set.
func (o *Skill) HasSynonyms() bool {
	if o != nil && o.Synonyms != nil {
		return true
	}

	return false
}

// SetSynonyms gets a reference to the given []string and assigns it to the Synonyms field.
func (o *Skill) SetSynonyms(v []string) {
	o.Synonyms = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *Skill) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Skill) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *Skill) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *Skill) SetDescription(v string) {
	o.Description = &v
}

// GetInvest returns the Invest field value if set, zero value otherwise.
func (o *Skill) GetInvest() bool {
	if o == nil || o.Invest == nil {
		var ret bool
		return ret
	}
	return *o.Invest
}

// GetInvestOk returns a tuple with the Invest field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Skill) GetInvestOk() (*bool, bool) {
	if o == nil || o.Invest == nil {
		return nil, false
	}
	return o.Invest, true
}

// HasInvest returns a boolean if a field has been set.
func (o *Skill) HasInvest() bool {
	if o != nil && o.Invest != nil {
		return true
	}

	return false
}

// SetInvest gets a reference to the given bool and assigns it to the Invest field.
func (o *Skill) SetInvest(v bool) {
	o.Invest = &v
}

// GetKindgiver returns the Kindgiver field value if set, zero value otherwise.
func (o *Skill) GetKindgiver() bool {
	if o == nil || o.Kindgiver == nil {
		var ret bool
		return ret
	}
	return *o.Kindgiver
}

// GetKindgiverOk returns a tuple with the Kindgiver field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Skill) GetKindgiverOk() (*bool, bool) {
	if o == nil || o.Kindgiver == nil {
		return nil, false
	}
	return o.Kindgiver, true
}

// HasKindgiver returns a boolean if a field has been set.
func (o *Skill) HasKindgiver() bool {
	if o != nil && o.Kindgiver != nil {
		return true
	}

	return false
}

// SetKindgiver gets a reference to the given bool and assigns it to the Kindgiver field.
func (o *Skill) SetKindgiver(v bool) {
	o.Kindgiver = &v
}

func (o Skill) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["name"] = o.Name
	}
	if true {
		toSerialize["suggestion"] = o.Suggestion
	}
	if o.Linkable != nil {
		toSerialize["linkable"] = o.Linkable
	}
	if o.Synonyms != nil {
		toSerialize["synonyms"] = o.Synonyms
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.Invest != nil {
		toSerialize["invest"] = o.Invest
	}
	if o.Kindgiver != nil {
		toSerialize["kindgiver"] = o.Kindgiver
	}
	return json.Marshal(toSerialize)
}

type NullableSkill struct {
	value *Skill
	isSet bool
}

func (v NullableSkill) Get() *Skill {
	return v.value
}

func (v *NullableSkill) Set(val *Skill) {
	v.value = val
	v.isSet = true
}

func (v NullableSkill) IsSet() bool {
	return v.isSet
}

func (v *NullableSkill) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSkill(val *Skill) *NullableSkill {
	return &NullableSkill{value: val, isSet: true}
}

func (v NullableSkill) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSkill) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


