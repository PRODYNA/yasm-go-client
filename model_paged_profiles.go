/*
YASM (Yet Another Skill Management) API

This is an example of using OAuth2 Implicit Flow in a specification to describe security to your API.

API version: 1.26.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// checks if the PagedProfiles type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PagedProfiles{}

// PagedProfiles struct for PagedProfiles
type PagedProfiles struct {
	Page
	Profiles []ProfileDetails `json:"profiles,omitempty"`
}

// NewPagedProfiles instantiates a new PagedProfiles object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPagedProfiles() *PagedProfiles {
	this := PagedProfiles{}
	return &this
}

// NewPagedProfilesWithDefaults instantiates a new PagedProfiles object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPagedProfilesWithDefaults() *PagedProfiles {
	this := PagedProfiles{}
	return &this
}

// GetProfiles returns the Profiles field value if set, zero value otherwise.
func (o *PagedProfiles) GetProfiles() []ProfileDetails {
	if o == nil || IsNil(o.Profiles) {
		var ret []ProfileDetails
		return ret
	}
	return o.Profiles
}

// GetProfilesOk returns a tuple with the Profiles field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PagedProfiles) GetProfilesOk() ([]ProfileDetails, bool) {
	if o == nil || IsNil(o.Profiles) {
		return nil, false
	}
	return o.Profiles, true
}

// HasProfiles returns a boolean if a field has been set.
func (o *PagedProfiles) HasProfiles() bool {
	if o != nil && !IsNil(o.Profiles) {
		return true
	}

	return false
}

// SetProfiles gets a reference to the given []ProfileDetails and assigns it to the Profiles field.
func (o *PagedProfiles) SetProfiles(v []ProfileDetails) {
	o.Profiles = v
}

func (o PagedProfiles) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PagedProfiles) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	serializedPage, errPage := json.Marshal(o.Page)
	if errPage != nil {
		return map[string]interface{}{}, errPage
	}
	errPage = json.Unmarshal([]byte(serializedPage), &toSerialize)
	if errPage != nil {
		return map[string]interface{}{}, errPage
	}
	if !IsNil(o.Profiles) {
		toSerialize["profiles"] = o.Profiles
	}
	return toSerialize, nil
}

type NullablePagedProfiles struct {
	value *PagedProfiles
	isSet bool
}

func (v NullablePagedProfiles) Get() *PagedProfiles {
	return v.value
}

func (v *NullablePagedProfiles) Set(val *PagedProfiles) {
	v.value = val
	v.isSet = true
}

func (v NullablePagedProfiles) IsSet() bool {
	return v.isSet
}

func (v *NullablePagedProfiles) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePagedProfiles(val *PagedProfiles) *NullablePagedProfiles {
	return &NullablePagedProfiles{value: val, isSet: true}
}

func (v NullablePagedProfiles) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePagedProfiles) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


