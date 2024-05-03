/*
YASM (Yet Another Skill Management) API

This is an example of using OAuth2 Implicit Flow in a specification to describe security to your API.

API version: 1.23.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// checks if the Suggestable type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Suggestable{}

// Suggestable struct for Suggestable
type Suggestable struct {
	Suggestion bool `json:"suggestion"`
}

// NewSuggestable instantiates a new Suggestable object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSuggestable(suggestion bool) *Suggestable {
	this := Suggestable{}
	this.Suggestion = suggestion
	return &this
}

// NewSuggestableWithDefaults instantiates a new Suggestable object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSuggestableWithDefaults() *Suggestable {
	this := Suggestable{}
	var suggestion bool = false
	this.Suggestion = suggestion
	return &this
}

// GetSuggestion returns the Suggestion field value
func (o *Suggestable) GetSuggestion() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Suggestion
}

// GetSuggestionOk returns a tuple with the Suggestion field value
// and a boolean to check if the value has been set.
func (o *Suggestable) GetSuggestionOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Suggestion, true
}

// SetSuggestion sets field value
func (o *Suggestable) SetSuggestion(v bool) {
	o.Suggestion = v
}

func (o Suggestable) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Suggestable) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["suggestion"] = o.Suggestion
	return toSerialize, nil
}

type NullableSuggestable struct {
	value *Suggestable
	isSet bool
}

func (v NullableSuggestable) Get() *Suggestable {
	return v.value
}

func (v *NullableSuggestable) Set(val *Suggestable) {
	v.value = val
	v.isSet = true
}

func (v NullableSuggestable) IsSet() bool {
	return v.isSet
}

func (v *NullableSuggestable) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSuggestable(val *Suggestable) *NullableSuggestable {
	return &NullableSuggestable{value: val, isSet: true}
}

func (v NullableSuggestable) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSuggestable) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


