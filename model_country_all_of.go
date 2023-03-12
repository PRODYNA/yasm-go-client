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

// CountryAllOf struct for CountryAllOf
type CountryAllOf struct {
	// base64 encoded image
	Picture *string `json:"picture,omitempty"`
}

// NewCountryAllOf instantiates a new CountryAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCountryAllOf() *CountryAllOf {
	this := CountryAllOf{}
	return &this
}

// NewCountryAllOfWithDefaults instantiates a new CountryAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCountryAllOfWithDefaults() *CountryAllOf {
	this := CountryAllOf{}
	return &this
}

// GetPicture returns the Picture field value if set, zero value otherwise.
func (o *CountryAllOf) GetPicture() string {
	if o == nil || o.Picture == nil {
		var ret string
		return ret
	}
	return *o.Picture
}

// GetPictureOk returns a tuple with the Picture field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CountryAllOf) GetPictureOk() (*string, bool) {
	if o == nil || o.Picture == nil {
		return nil, false
	}
	return o.Picture, true
}

// HasPicture returns a boolean if a field has been set.
func (o *CountryAllOf) HasPicture() bool {
	if o != nil && o.Picture != nil {
		return true
	}

	return false
}

// SetPicture gets a reference to the given string and assigns it to the Picture field.
func (o *CountryAllOf) SetPicture(v string) {
	o.Picture = &v
}

func (o CountryAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Picture != nil {
		toSerialize["picture"] = o.Picture
	}
	return json.Marshal(toSerialize)
}

type NullableCountryAllOf struct {
	value *CountryAllOf
	isSet bool
}

func (v NullableCountryAllOf) Get() *CountryAllOf {
	return v.value
}

func (v *NullableCountryAllOf) Set(val *CountryAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableCountryAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableCountryAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCountryAllOf(val *CountryAllOf) *NullableCountryAllOf {
	return &NullableCountryAllOf{value: val, isSet: true}
}

func (v NullableCountryAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCountryAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


