/*
YASM (Yet Another Skill Management) API

This is an example of using OAuth2 Implicit Flow in a specification to describe security to your API.

API version: 1.50.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// checks if the Industry type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Industry{}

// Industry struct for Industry
type Industry struct {
	NamedDomainModel
	Synonyms []string `json:"synonyms,omitempty"`
}

// NewIndustry instantiates a new Industry object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIndustry(id string, name string) *Industry {
	this := Industry{}
	this.Id = id
	this.Name = name
	return &this
}

// NewIndustryWithDefaults instantiates a new Industry object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIndustryWithDefaults() *Industry {
	this := Industry{}
	return &this
}

// GetSynonyms returns the Synonyms field value if set, zero value otherwise.
func (o *Industry) GetSynonyms() []string {
	if o == nil || IsNil(o.Synonyms) {
		var ret []string
		return ret
	}
	return o.Synonyms
}

// GetSynonymsOk returns a tuple with the Synonyms field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Industry) GetSynonymsOk() ([]string, bool) {
	if o == nil || IsNil(o.Synonyms) {
		return nil, false
	}
	return o.Synonyms, true
}

// HasSynonyms returns a boolean if a field has been set.
func (o *Industry) HasSynonyms() bool {
	if o != nil && !IsNil(o.Synonyms) {
		return true
	}

	return false
}

// SetSynonyms gets a reference to the given []string and assigns it to the Synonyms field.
func (o *Industry) SetSynonyms(v []string) {
	o.Synonyms = v
}

func (o Industry) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Industry) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	serializedNamedDomainModel, errNamedDomainModel := json.Marshal(o.NamedDomainModel)
	if errNamedDomainModel != nil {
		return map[string]interface{}{}, errNamedDomainModel
	}
	errNamedDomainModel = json.Unmarshal([]byte(serializedNamedDomainModel), &toSerialize)
	if errNamedDomainModel != nil {
		return map[string]interface{}{}, errNamedDomainModel
	}
	if !IsNil(o.Synonyms) {
		toSerialize["synonyms"] = o.Synonyms
	}
	return toSerialize, nil
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


