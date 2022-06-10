/*
YASM (Yet Another Skill Management) API

This is an example of using OAuth2 Implicit Flow in a specification to describe security to your API.

API version: 1.0.15
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// PersonOrganizationFilterAllOf struct for PersonOrganizationFilterAllOf
type PersonOrganizationFilterAllOf struct {
	AmountOfProjects *MinMax `json:"amountOfProjects,omitempty"`
}

// NewPersonOrganizationFilterAllOf instantiates a new PersonOrganizationFilterAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPersonOrganizationFilterAllOf() *PersonOrganizationFilterAllOf {
	this := PersonOrganizationFilterAllOf{}
	return &this
}

// NewPersonOrganizationFilterAllOfWithDefaults instantiates a new PersonOrganizationFilterAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPersonOrganizationFilterAllOfWithDefaults() *PersonOrganizationFilterAllOf {
	this := PersonOrganizationFilterAllOf{}
	return &this
}

// GetAmountOfProjects returns the AmountOfProjects field value if set, zero value otherwise.
func (o *PersonOrganizationFilterAllOf) GetAmountOfProjects() MinMax {
	if o == nil || o.AmountOfProjects == nil {
		var ret MinMax
		return ret
	}
	return *o.AmountOfProjects
}

// GetAmountOfProjectsOk returns a tuple with the AmountOfProjects field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PersonOrganizationFilterAllOf) GetAmountOfProjectsOk() (*MinMax, bool) {
	if o == nil || o.AmountOfProjects == nil {
		return nil, false
	}
	return o.AmountOfProjects, true
}

// HasAmountOfProjects returns a boolean if a field has been set.
func (o *PersonOrganizationFilterAllOf) HasAmountOfProjects() bool {
	if o != nil && o.AmountOfProjects != nil {
		return true
	}

	return false
}

// SetAmountOfProjects gets a reference to the given MinMax and assigns it to the AmountOfProjects field.
func (o *PersonOrganizationFilterAllOf) SetAmountOfProjects(v MinMax) {
	o.AmountOfProjects = &v
}

func (o PersonOrganizationFilterAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AmountOfProjects != nil {
		toSerialize["amountOfProjects"] = o.AmountOfProjects
	}
	return json.Marshal(toSerialize)
}

type NullablePersonOrganizationFilterAllOf struct {
	value *PersonOrganizationFilterAllOf
	isSet bool
}

func (v NullablePersonOrganizationFilterAllOf) Get() *PersonOrganizationFilterAllOf {
	return v.value
}

func (v *NullablePersonOrganizationFilterAllOf) Set(val *PersonOrganizationFilterAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullablePersonOrganizationFilterAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullablePersonOrganizationFilterAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePersonOrganizationFilterAllOf(val *PersonOrganizationFilterAllOf) *NullablePersonOrganizationFilterAllOf {
	return &NullablePersonOrganizationFilterAllOf{value: val, isSet: true}
}

func (v NullablePersonOrganizationFilterAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePersonOrganizationFilterAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


