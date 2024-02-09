/*
YASM (Yet Another Skill Management) API

This is an example of using OAuth2 Implicit Flow in a specification to describe security to your API.

API version: 1.12.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// checks if the ProfileDetails type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ProfileDetails{}

// ProfileDetails struct for ProfileDetails
type ProfileDetails struct {
	Profile *Profile `json:"profile,omitempty"`
}

// NewProfileDetails instantiates a new ProfileDetails object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProfileDetails() *ProfileDetails {
	this := ProfileDetails{}
	return &this
}

// NewProfileDetailsWithDefaults instantiates a new ProfileDetails object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProfileDetailsWithDefaults() *ProfileDetails {
	this := ProfileDetails{}
	return &this
}

// GetProfile returns the Profile field value if set, zero value otherwise.
func (o *ProfileDetails) GetProfile() Profile {
	if o == nil || IsNil(o.Profile) {
		var ret Profile
		return ret
	}
	return *o.Profile
}

// GetProfileOk returns a tuple with the Profile field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProfileDetails) GetProfileOk() (*Profile, bool) {
	if o == nil || IsNil(o.Profile) {
		return nil, false
	}
	return o.Profile, true
}

// HasProfile returns a boolean if a field has been set.
func (o *ProfileDetails) HasProfile() bool {
	if o != nil && !IsNil(o.Profile) {
		return true
	}

	return false
}

// SetProfile gets a reference to the given Profile and assigns it to the Profile field.
func (o *ProfileDetails) SetProfile(v Profile) {
	o.Profile = &v
}

func (o ProfileDetails) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ProfileDetails) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Profile) {
		toSerialize["profile"] = o.Profile
	}
	return toSerialize, nil
}

type NullableProfileDetails struct {
	value *ProfileDetails
	isSet bool
}

func (v NullableProfileDetails) Get() *ProfileDetails {
	return v.value
}

func (v *NullableProfileDetails) Set(val *ProfileDetails) {
	v.value = val
	v.isSet = true
}

func (v NullableProfileDetails) IsSet() bool {
	return v.isSet
}

func (v *NullableProfileDetails) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProfileDetails(val *ProfileDetails) *NullableProfileDetails {
	return &NullableProfileDetails{value: val, isSet: true}
}

func (v NullableProfileDetails) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProfileDetails) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


