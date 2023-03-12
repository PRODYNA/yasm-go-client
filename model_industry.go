/*
YASM (Yet Another Skill Management) API

This is an example of using OAuth2 Implicit Flow in a specification to describe security to your API.

API version: 1.3.11
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// Industry struct for Industry
type Industry struct {
	NamedDomainModel
	Suggestion bool `json:"suggestion"`
	Synonyms []string `json:"synonyms,omitempty"`
}

// NewIndustry instantiates a new Industry object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIndustry(suggestion bool, id string, name string) *Industry {
	this := Industry{}
	this.Id = id
	this.Name = name
	this.Suggestion = suggestion
	return &this
}

// NewIndustryWithDefaults instantiates a new Industry object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIndustryWithDefaults() *Industry {
	this := Industry{}
	var suggestion bool = false
	this.Suggestion = suggestion
	return &this
}

// GetSuggestion returns the Suggestion field value
func (o *Industry) GetSuggestion() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Suggestion
}

// GetSuggestionOk returns a tuple with the Suggestion field value
// and a boolean to check if the value has been set.
func (o *Industry) GetSuggestionOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Suggestion, true
}

// SetSuggestion sets field value
func (o *Industry) SetSuggestion(v bool) {
	o.Suggestion = v
}

// GetSynonyms returns the Synonyms field value if set, zero value otherwise.
func (o *Industry) GetSynonyms() []string {
	if o == nil || o.Synonyms == nil {
		var ret []string
		return ret
	}
	return o.Synonyms
}

// GetSynonymsOk returns a tuple with the Synonyms field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Industry) GetSynonymsOk() ([]string, bool) {
	if o == nil || o.Synonyms == nil {
		return nil, false
	}
	return o.Synonyms, true
}

// HasSynonyms returns a boolean if a field has been set.
func (o *Industry) HasSynonyms() bool {
	if o != nil && o.Synonyms != nil {
		return true
	}

	return false
}

// SetSynonyms gets a reference to the given []string and assigns it to the Synonyms field.
func (o *Industry) SetSynonyms(v []string) {
	o.Synonyms = v
}

func (o Industry) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedNamedDomainModel, errNamedDomainModel := json.Marshal(o.NamedDomainModel)
	if errNamedDomainModel != nil {
		return []byte{}, errNamedDomainModel
	}
	errNamedDomainModel = json.Unmarshal([]byte(serializedNamedDomainModel), &toSerialize)
	if errNamedDomainModel != nil {
		return []byte{}, errNamedDomainModel
	}
	if true {
		toSerialize["suggestion"] = o.Suggestion
	}
	if o.Synonyms != nil {
		toSerialize["synonyms"] = o.Synonyms
	}
	return json.Marshal(toSerialize)
}

type NullableIndustry struct {
	value *Industry
	isSet bool
}

func (v NullableIndustry) Get() *Industry {
	return v.value
}

func (v *NullableIndustry) Set(val *Industry) {
	v.value = val
	v.isSet = true
}

func (v NullableIndustry) IsSet() bool {
	return v.isSet
}

func (v *NullableIndustry) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIndustry(val *Industry) *NullableIndustry {
	return &NullableIndustry{value: val, isSet: true}
}

func (v NullableIndustry) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIndustry) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


