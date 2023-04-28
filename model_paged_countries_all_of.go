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

// PagedCountriesAllOf struct for PagedCountriesAllOf
type PagedCountriesAllOf struct {
	Countries []Country `json:"countries,omitempty"`
}

// NewPagedCountriesAllOf instantiates a new PagedCountriesAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPagedCountriesAllOf() *PagedCountriesAllOf {
	this := PagedCountriesAllOf{}
	return &this
}

// NewPagedCountriesAllOfWithDefaults instantiates a new PagedCountriesAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPagedCountriesAllOfWithDefaults() *PagedCountriesAllOf {
	this := PagedCountriesAllOf{}
	return &this
}

// GetCountries returns the Countries field value if set, zero value otherwise.
func (o *PagedCountriesAllOf) GetCountries() []Country {
	if o == nil || o.Countries == nil {
		var ret []Country
		return ret
	}
	return o.Countries
}

// GetCountriesOk returns a tuple with the Countries field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PagedCountriesAllOf) GetCountriesOk() ([]Country, bool) {
	if o == nil || o.Countries == nil {
		return nil, false
	}
	return o.Countries, true
}

// HasCountries returns a boolean if a field has been set.
func (o *PagedCountriesAllOf) HasCountries() bool {
	if o != nil && o.Countries != nil {
		return true
	}

	return false
}

// SetCountries gets a reference to the given []Country and assigns it to the Countries field.
func (o *PagedCountriesAllOf) SetCountries(v []Country) {
	o.Countries = v
}

func (o PagedCountriesAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Countries != nil {
		toSerialize["countries"] = o.Countries
	}
	return json.Marshal(toSerialize)
}

type NullablePagedCountriesAllOf struct {
	value *PagedCountriesAllOf
	isSet bool
}

func (v NullablePagedCountriesAllOf) Get() *PagedCountriesAllOf {
	return v.value
}

func (v *NullablePagedCountriesAllOf) Set(val *PagedCountriesAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullablePagedCountriesAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullablePagedCountriesAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePagedCountriesAllOf(val *PagedCountriesAllOf) *NullablePagedCountriesAllOf {
	return &NullablePagedCountriesAllOf{value: val, isSet: true}
}

func (v NullablePagedCountriesAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePagedCountriesAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


