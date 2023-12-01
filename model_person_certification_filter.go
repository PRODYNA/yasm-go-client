/*
YASM (Yet Another Skill Management) API

This is an example of using OAuth2 Implicit Flow in a specification to describe security to your API.

API version: 1.7.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// checks if the PersonCertificationFilter type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PersonCertificationFilter{}

// PersonCertificationFilter struct for PersonCertificationFilter
type PersonCertificationFilter struct {
	AbstractEntityFilter
	// Include employees who started the certification
	StartedCertificaiton *bool `json:"startedCertificaiton,omitempty"`
}

// NewPersonCertificationFilter instantiates a new PersonCertificationFilter object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPersonCertificationFilter(id string) *PersonCertificationFilter {
	this := PersonCertificationFilter{}
	this.Id = id
	return &this
}

// NewPersonCertificationFilterWithDefaults instantiates a new PersonCertificationFilter object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPersonCertificationFilterWithDefaults() *PersonCertificationFilter {
	this := PersonCertificationFilter{}
	return &this
}

// GetStartedCertificaiton returns the StartedCertificaiton field value if set, zero value otherwise.
func (o *PersonCertificationFilter) GetStartedCertificaiton() bool {
	if o == nil || IsNil(o.StartedCertificaiton) {
		var ret bool
		return ret
	}
	return *o.StartedCertificaiton
}

// GetStartedCertificaitonOk returns a tuple with the StartedCertificaiton field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PersonCertificationFilter) GetStartedCertificaitonOk() (*bool, bool) {
	if o == nil || IsNil(o.StartedCertificaiton) {
		return nil, false
	}
	return o.StartedCertificaiton, true
}

// HasStartedCertificaiton returns a boolean if a field has been set.
func (o *PersonCertificationFilter) HasStartedCertificaiton() bool {
	if o != nil && !IsNil(o.StartedCertificaiton) {
		return true
	}

	return false
}

// SetStartedCertificaiton gets a reference to the given bool and assigns it to the StartedCertificaiton field.
func (o *PersonCertificationFilter) SetStartedCertificaiton(v bool) {
	o.StartedCertificaiton = &v
}

func (o PersonCertificationFilter) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PersonCertificationFilter) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	serializedAbstractEntityFilter, errAbstractEntityFilter := json.Marshal(o.AbstractEntityFilter)
	if errAbstractEntityFilter != nil {
		return map[string]interface{}{}, errAbstractEntityFilter
	}
	errAbstractEntityFilter = json.Unmarshal([]byte(serializedAbstractEntityFilter), &toSerialize)
	if errAbstractEntityFilter != nil {
		return map[string]interface{}{}, errAbstractEntityFilter
	}
	if !IsNil(o.StartedCertificaiton) {
		toSerialize["startedCertificaiton"] = o.StartedCertificaiton
	}
	return toSerialize, nil
}

type NullablePersonCertificationFilter struct {
	value *PersonCertificationFilter
	isSet bool
}

func (v NullablePersonCertificationFilter) Get() *PersonCertificationFilter {
	return v.value
}

func (v *NullablePersonCertificationFilter) Set(val *PersonCertificationFilter) {
	v.value = val
	v.isSet = true
}

func (v NullablePersonCertificationFilter) IsSet() bool {
	return v.isSet
}

func (v *NullablePersonCertificationFilter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePersonCertificationFilter(val *PersonCertificationFilter) *NullablePersonCertificationFilter {
	return &NullablePersonCertificationFilter{value: val, isSet: true}
}

func (v NullablePersonCertificationFilter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePersonCertificationFilter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


