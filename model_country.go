/*
YASM (Yet Another Skill Management) API

This is an example of using OAuth2 Implicit Flow in a specification to describe security to your API.

API version: 0.9.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// Country struct for Country
type Country struct {
	Id string `json:"id"`
	Name string `json:"name"`
}

// NewCountry instantiates a new Country object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCountry(id string, name string) *Country {
	this := Country{}
	this.Id = id
	this.Name = name
	return &this
}

// NewCountryWithDefaults instantiates a new Country object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCountryWithDefaults() *Country {
	this := Country{}
	return &this
}

// GetId returns the Id field value
func (o *Country) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *Country) GetIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *Country) SetId(v string) {
	o.Id = v
}

// GetName returns the Name field value
func (o *Country) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *Country) GetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *Country) SetName(v string) {
	o.Name = v
}

func (o Country) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["name"] = o.Name
	}
	return json.Marshal(toSerialize)
}

type NullableCountry struct {
	value *Country
	isSet bool
}

func (v NullableCountry) Get() *Country {
	return v.value
}

func (v *NullableCountry) Set(val *Country) {
	v.value = val
	v.isSet = true
}

func (v NullableCountry) IsSet() bool {
	return v.isSet
}

func (v *NullableCountry) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCountry(val *Country) *NullableCountry {
	return &NullableCountry{value: val, isSet: true}
}

func (v NullableCountry) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCountry) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


