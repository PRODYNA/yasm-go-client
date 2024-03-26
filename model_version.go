/*
YASM (Yet Another Skill Management) API

This is an example of using OAuth2 Implicit Flow in a specification to describe security to your API.

API version: 1.16.5
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// checks if the Version type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Version{}

// Version struct for Version
type Version struct {
	ApiVersion string `json:"apiVersion"`
	ServerVersion string `json:"serverVersion"`
}

// NewVersion instantiates a new Version object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVersion(apiVersion string, serverVersion string) *Version {
	this := Version{}
	this.ApiVersion = apiVersion
	this.ServerVersion = serverVersion
	return &this
}

// NewVersionWithDefaults instantiates a new Version object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVersionWithDefaults() *Version {
	this := Version{}
	return &this
}

// GetApiVersion returns the ApiVersion field value
func (o *Version) GetApiVersion() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ApiVersion
}

// GetApiVersionOk returns a tuple with the ApiVersion field value
// and a boolean to check if the value has been set.
func (o *Version) GetApiVersionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ApiVersion, true
}

// SetApiVersion sets field value
func (o *Version) SetApiVersion(v string) {
	o.ApiVersion = v
}

// GetServerVersion returns the ServerVersion field value
func (o *Version) GetServerVersion() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ServerVersion
}

// GetServerVersionOk returns a tuple with the ServerVersion field value
// and a boolean to check if the value has been set.
func (o *Version) GetServerVersionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ServerVersion, true
}

// SetServerVersion sets field value
func (o *Version) SetServerVersion(v string) {
	o.ServerVersion = v
}

func (o Version) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Version) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["apiVersion"] = o.ApiVersion
	toSerialize["serverVersion"] = o.ServerVersion
	return toSerialize, nil
}

type NullableVersion struct {
	value *Version
	isSet bool
}

func (v NullableVersion) Get() *Version {
	return v.value
}

func (v *NullableVersion) Set(val *Version) {
	v.value = val
	v.isSet = true
}

func (v NullableVersion) IsSet() bool {
	return v.isSet
}

func (v *NullableVersion) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVersion(val *Version) *NullableVersion {
	return &NullableVersion{value: val, isSet: true}
}

func (v NullableVersion) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVersion) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


