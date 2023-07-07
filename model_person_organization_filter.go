/*
YASM (Yet Another Skill Management) API

This is an example of using OAuth2 Implicit Flow in a specification to describe security to your API.

API version: 1.4.3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// PersonOrganizationFilter struct for PersonOrganizationFilter
type PersonOrganizationFilter struct {
	AbstractEntityFilter
	AmountOfProjects *MinMax `json:"amountOfProjects,omitempty"`
}

// NewPersonOrganizationFilter instantiates a new PersonOrganizationFilter object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPersonOrganizationFilter(id string) *PersonOrganizationFilter {
	this := PersonOrganizationFilter{}
	this.Id = id
	return &this
}

// NewPersonOrganizationFilterWithDefaults instantiates a new PersonOrganizationFilter object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPersonOrganizationFilterWithDefaults() *PersonOrganizationFilter {
	this := PersonOrganizationFilter{}
	return &this
}

// GetAmountOfProjects returns the AmountOfProjects field value if set, zero value otherwise.
func (o *PersonOrganizationFilter) GetAmountOfProjects() MinMax {
	if o == nil || o.AmountOfProjects == nil {
		var ret MinMax
		return ret
	}
	return *o.AmountOfProjects
}

// GetAmountOfProjectsOk returns a tuple with the AmountOfProjects field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PersonOrganizationFilter) GetAmountOfProjectsOk() (*MinMax, bool) {
	if o == nil || o.AmountOfProjects == nil {
		return nil, false
	}
	return o.AmountOfProjects, true
}

// HasAmountOfProjects returns a boolean if a field has been set.
func (o *PersonOrganizationFilter) HasAmountOfProjects() bool {
	if o != nil && o.AmountOfProjects != nil {
		return true
	}

	return false
}

// SetAmountOfProjects gets a reference to the given MinMax and assigns it to the AmountOfProjects field.
func (o *PersonOrganizationFilter) SetAmountOfProjects(v MinMax) {
	o.AmountOfProjects = &v
}

func (o PersonOrganizationFilter) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedAbstractEntityFilter, errAbstractEntityFilter := json.Marshal(o.AbstractEntityFilter)
	if errAbstractEntityFilter != nil {
		return []byte{}, errAbstractEntityFilter
	}
	errAbstractEntityFilter = json.Unmarshal([]byte(serializedAbstractEntityFilter), &toSerialize)
	if errAbstractEntityFilter != nil {
		return []byte{}, errAbstractEntityFilter
	}
	if o.AmountOfProjects != nil {
		toSerialize["amountOfProjects"] = o.AmountOfProjects
	}
	return json.Marshal(toSerialize)
}

type NullablePersonOrganizationFilter struct {
	value *PersonOrganizationFilter
	isSet bool
}

func (v NullablePersonOrganizationFilter) Get() *PersonOrganizationFilter {
	return v.value
}

func (v *NullablePersonOrganizationFilter) Set(val *PersonOrganizationFilter) {
	v.value = val
	v.isSet = true
}

func (v NullablePersonOrganizationFilter) IsSet() bool {
	return v.isSet
}

func (v *NullablePersonOrganizationFilter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePersonOrganizationFilter(val *PersonOrganizationFilter) *NullablePersonOrganizationFilter {
	return &NullablePersonOrganizationFilter{value: val, isSet: true}
}

func (v NullablePersonOrganizationFilter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePersonOrganizationFilter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


