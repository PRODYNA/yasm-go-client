/*
YASM (Yet Another Skill Management) API

This is an example of using OAuth2 Implicit Flow in a specification to describe security to your API.

API version: 1.4.13
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// PagedCountries struct for PagedCountries
type PagedCountries struct {
	Page
	Countries []Country `json:"countries,omitempty"`
}

// NewPagedCountries instantiates a new PagedCountries object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPagedCountries() *PagedCountries {
	this := PagedCountries{}
	return &this
}

// NewPagedCountriesWithDefaults instantiates a new PagedCountries object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPagedCountriesWithDefaults() *PagedCountries {
	this := PagedCountries{}
	return &this
}

// GetCountries returns the Countries field value if set, zero value otherwise.
func (o *PagedCountries) GetCountries() []Country {
	if o == nil || o.Countries == nil {
		var ret []Country
		return ret
	}
	return o.Countries
}

// GetCountriesOk returns a tuple with the Countries field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PagedCountries) GetCountriesOk() ([]Country, bool) {
	if o == nil || o.Countries == nil {
		return nil, false
	}
	return o.Countries, true
}

// HasCountries returns a boolean if a field has been set.
func (o *PagedCountries) HasCountries() bool {
	if o != nil && o.Countries != nil {
		return true
	}

	return false
}

// SetCountries gets a reference to the given []Country and assigns it to the Countries field.
func (o *PagedCountries) SetCountries(v []Country) {
	o.Countries = v
}

func (o PagedCountries) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedPage, errPage := json.Marshal(o.Page)
	if errPage != nil {
		return []byte{}, errPage
	}
	errPage = json.Unmarshal([]byte(serializedPage), &toSerialize)
	if errPage != nil {
		return []byte{}, errPage
	}
	if o.Countries != nil {
		toSerialize["countries"] = o.Countries
	}
	return json.Marshal(toSerialize)
}

type NullablePagedCountries struct {
	value *PagedCountries
	isSet bool
}

func (v NullablePagedCountries) Get() *PagedCountries {
	return v.value
}

func (v *NullablePagedCountries) Set(val *PagedCountries) {
	v.value = val
	v.isSet = true
}

func (v NullablePagedCountries) IsSet() bool {
	return v.isSet
}

func (v *NullablePagedCountries) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePagedCountries(val *PagedCountries) *NullablePagedCountries {
	return &NullablePagedCountries{value: val, isSet: true}
}

func (v NullablePagedCountries) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePagedCountries) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


