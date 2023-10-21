/*
YASM (Yet Another Skill Management) API

This is an example of using OAuth2 Implicit Flow in a specification to describe security to your API.

API version: 1.4.11
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// PagedLanguagesAllOf struct for PagedLanguagesAllOf
type PagedLanguagesAllOf struct {
	Languages []Language `json:"languages,omitempty"`
}

// NewPagedLanguagesAllOf instantiates a new PagedLanguagesAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPagedLanguagesAllOf() *PagedLanguagesAllOf {
	this := PagedLanguagesAllOf{}
	return &this
}

// NewPagedLanguagesAllOfWithDefaults instantiates a new PagedLanguagesAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPagedLanguagesAllOfWithDefaults() *PagedLanguagesAllOf {
	this := PagedLanguagesAllOf{}
	return &this
}

// GetLanguages returns the Languages field value if set, zero value otherwise.
func (o *PagedLanguagesAllOf) GetLanguages() []Language {
	if o == nil || o.Languages == nil {
		var ret []Language
		return ret
	}
	return o.Languages
}

// GetLanguagesOk returns a tuple with the Languages field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PagedLanguagesAllOf) GetLanguagesOk() ([]Language, bool) {
	if o == nil || o.Languages == nil {
		return nil, false
	}
	return o.Languages, true
}

// HasLanguages returns a boolean if a field has been set.
func (o *PagedLanguagesAllOf) HasLanguages() bool {
	if o != nil && o.Languages != nil {
		return true
	}

	return false
}

// SetLanguages gets a reference to the given []Language and assigns it to the Languages field.
func (o *PagedLanguagesAllOf) SetLanguages(v []Language) {
	o.Languages = v
}

func (o PagedLanguagesAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Languages != nil {
		toSerialize["languages"] = o.Languages
	}
	return json.Marshal(toSerialize)
}

type NullablePagedLanguagesAllOf struct {
	value *PagedLanguagesAllOf
	isSet bool
}

func (v NullablePagedLanguagesAllOf) Get() *PagedLanguagesAllOf {
	return v.value
}

func (v *NullablePagedLanguagesAllOf) Set(val *PagedLanguagesAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullablePagedLanguagesAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullablePagedLanguagesAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePagedLanguagesAllOf(val *PagedLanguagesAllOf) *NullablePagedLanguagesAllOf {
	return &NullablePagedLanguagesAllOf{value: val, isSet: true}
}

func (v NullablePagedLanguagesAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePagedLanguagesAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


