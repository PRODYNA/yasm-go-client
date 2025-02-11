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

// checks if the Dimension type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Dimension{}

// Dimension struct for Dimension
type Dimension struct {
	Width *float32 `json:"width,omitempty"`
	Height *float32 `json:"height,omitempty"`
}

// NewDimension instantiates a new Dimension object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDimension() *Dimension {
	this := Dimension{}
	return &this
}

// NewDimensionWithDefaults instantiates a new Dimension object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDimensionWithDefaults() *Dimension {
	this := Dimension{}
	return &this
}

// GetWidth returns the Width field value if set, zero value otherwise.
func (o *Dimension) GetWidth() float32 {
	if o == nil || IsNil(o.Width) {
		var ret float32
		return ret
	}
	return *o.Width
}

// GetWidthOk returns a tuple with the Width field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Dimension) GetWidthOk() (*float32, bool) {
	if o == nil || IsNil(o.Width) {
		return nil, false
	}
	return o.Width, true
}

// HasWidth returns a boolean if a field has been set.
func (o *Dimension) HasWidth() bool {
	if o != nil && !IsNil(o.Width) {
		return true
	}

	return false
}

// SetWidth gets a reference to the given float32 and assigns it to the Width field.
func (o *Dimension) SetWidth(v float32) {
	o.Width = &v
}

// GetHeight returns the Height field value if set, zero value otherwise.
func (o *Dimension) GetHeight() float32 {
	if o == nil || IsNil(o.Height) {
		var ret float32
		return ret
	}
	return *o.Height
}

// GetHeightOk returns a tuple with the Height field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Dimension) GetHeightOk() (*float32, bool) {
	if o == nil || IsNil(o.Height) {
		return nil, false
	}
	return o.Height, true
}

// HasHeight returns a boolean if a field has been set.
func (o *Dimension) HasHeight() bool {
	if o != nil && !IsNil(o.Height) {
		return true
	}

	return false
}

// SetHeight gets a reference to the given float32 and assigns it to the Height field.
func (o *Dimension) SetHeight(v float32) {
	o.Height = &v
}

func (o Dimension) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Dimension) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Width) {
		toSerialize["width"] = o.Width
	}
	if !IsNil(o.Height) {
		toSerialize["height"] = o.Height
	}
	return toSerialize, nil
}

type NullableDimension struct {
	value *Dimension
	isSet bool
}

func (v NullableDimension) Get() *Dimension {
	return v.value
}

func (v *NullableDimension) Set(val *Dimension) {
	v.value = val
	v.isSet = true
}

func (v NullableDimension) IsSet() bool {
	return v.isSet
}

func (v *NullableDimension) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDimension(val *Dimension) *NullableDimension {
	return &NullableDimension{value: val, isSet: true}
}

func (v NullableDimension) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDimension) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


