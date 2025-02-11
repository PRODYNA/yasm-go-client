/*
YASM (Yet Another Skill Management) API

This is an example of using OAuth2 Implicit Flow in a specification to describe security to your API.

API version: 1.63.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// checks if the DeviceType type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DeviceType{}

// DeviceType struct for DeviceType
type DeviceType struct {
	DeviceType *string `json:"deviceType,omitempty"`
}

// NewDeviceType instantiates a new DeviceType object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeviceType() *DeviceType {
	this := DeviceType{}
	return &this
}

// NewDeviceTypeWithDefaults instantiates a new DeviceType object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeviceTypeWithDefaults() *DeviceType {
	this := DeviceType{}
	return &this
}

// GetDeviceType returns the DeviceType field value if set, zero value otherwise.
func (o *DeviceType) GetDeviceType() string {
	if o == nil || IsNil(o.DeviceType) {
		var ret string
		return ret
	}
	return *o.DeviceType
}

// GetDeviceTypeOk returns a tuple with the DeviceType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceType) GetDeviceTypeOk() (*string, bool) {
	if o == nil || IsNil(o.DeviceType) {
		return nil, false
	}
	return o.DeviceType, true
}

// HasDeviceType returns a boolean if a field has been set.
func (o *DeviceType) HasDeviceType() bool {
	if o != nil && !IsNil(o.DeviceType) {
		return true
	}

	return false
}

// SetDeviceType gets a reference to the given string and assigns it to the DeviceType field.
func (o *DeviceType) SetDeviceType(v string) {
	o.DeviceType = &v
}

func (o DeviceType) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DeviceType) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.DeviceType) {
		toSerialize["deviceType"] = o.DeviceType
	}
	return toSerialize, nil
}

type NullableDeviceType struct {
	value *DeviceType
	isSet bool
}

func (v NullableDeviceType) Get() *DeviceType {
	return v.value
}

func (v *NullableDeviceType) Set(val *DeviceType) {
	v.value = val
	v.isSet = true
}

func (v NullableDeviceType) IsSet() bool {
	return v.isSet
}

func (v *NullableDeviceType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeviceType(val *DeviceType) *NullableDeviceType {
	return &NullableDeviceType{value: val, isSet: true}
}

func (v NullableDeviceType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeviceType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


