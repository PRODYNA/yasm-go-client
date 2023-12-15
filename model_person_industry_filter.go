/*
YASM (Yet Another Skill Management) API

This is an example of using OAuth2 Implicit Flow in a specification to describe security to your API.

API version: 1.10.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// checks if the PersonIndustryFilter type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PersonIndustryFilter{}

// PersonIndustryFilter struct for PersonIndustryFilter
type PersonIndustryFilter struct {
	AbstractEntityFilter
	ActiveProjects *bool `json:"activeProjects,omitempty"`
	AmountOfProjects *MinMax `json:"amountOfProjects,omitempty"`
	ExperienceInMonth *MinMax `json:"experienceInMonth,omitempty"`
}

// NewPersonIndustryFilter instantiates a new PersonIndustryFilter object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPersonIndustryFilter(id string) *PersonIndustryFilter {
	this := PersonIndustryFilter{}
	this.Id = id
	return &this
}

// NewPersonIndustryFilterWithDefaults instantiates a new PersonIndustryFilter object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPersonIndustryFilterWithDefaults() *PersonIndustryFilter {
	this := PersonIndustryFilter{}
	return &this
}

// GetActiveProjects returns the ActiveProjects field value if set, zero value otherwise.
func (o *PersonIndustryFilter) GetActiveProjects() bool {
	if o == nil || IsNil(o.ActiveProjects) {
		var ret bool
		return ret
	}
	return *o.ActiveProjects
}

// GetActiveProjectsOk returns a tuple with the ActiveProjects field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PersonIndustryFilter) GetActiveProjectsOk() (*bool, bool) {
	if o == nil || IsNil(o.ActiveProjects) {
		return nil, false
	}
	return o.ActiveProjects, true
}

// HasActiveProjects returns a boolean if a field has been set.
func (o *PersonIndustryFilter) HasActiveProjects() bool {
	if o != nil && !IsNil(o.ActiveProjects) {
		return true
	}

	return false
}

// SetActiveProjects gets a reference to the given bool and assigns it to the ActiveProjects field.
func (o *PersonIndustryFilter) SetActiveProjects(v bool) {
	o.ActiveProjects = &v
}

// GetAmountOfProjects returns the AmountOfProjects field value if set, zero value otherwise.
func (o *PersonIndustryFilter) GetAmountOfProjects() MinMax {
	if o == nil || IsNil(o.AmountOfProjects) {
		var ret MinMax
		return ret
	}
	return *o.AmountOfProjects
}

// GetAmountOfProjectsOk returns a tuple with the AmountOfProjects field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PersonIndustryFilter) GetAmountOfProjectsOk() (*MinMax, bool) {
	if o == nil || IsNil(o.AmountOfProjects) {
		return nil, false
	}
	return o.AmountOfProjects, true
}

// HasAmountOfProjects returns a boolean if a field has been set.
func (o *PersonIndustryFilter) HasAmountOfProjects() bool {
	if o != nil && !IsNil(o.AmountOfProjects) {
		return true
	}

	return false
}

// SetAmountOfProjects gets a reference to the given MinMax and assigns it to the AmountOfProjects field.
func (o *PersonIndustryFilter) SetAmountOfProjects(v MinMax) {
	o.AmountOfProjects = &v
}

// GetExperienceInMonth returns the ExperienceInMonth field value if set, zero value otherwise.
func (o *PersonIndustryFilter) GetExperienceInMonth() MinMax {
	if o == nil || IsNil(o.ExperienceInMonth) {
		var ret MinMax
		return ret
	}
	return *o.ExperienceInMonth
}

// GetExperienceInMonthOk returns a tuple with the ExperienceInMonth field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PersonIndustryFilter) GetExperienceInMonthOk() (*MinMax, bool) {
	if o == nil || IsNil(o.ExperienceInMonth) {
		return nil, false
	}
	return o.ExperienceInMonth, true
}

// HasExperienceInMonth returns a boolean if a field has been set.
func (o *PersonIndustryFilter) HasExperienceInMonth() bool {
	if o != nil && !IsNil(o.ExperienceInMonth) {
		return true
	}

	return false
}

// SetExperienceInMonth gets a reference to the given MinMax and assigns it to the ExperienceInMonth field.
func (o *PersonIndustryFilter) SetExperienceInMonth(v MinMax) {
	o.ExperienceInMonth = &v
}

func (o PersonIndustryFilter) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PersonIndustryFilter) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	serializedAbstractEntityFilter, errAbstractEntityFilter := json.Marshal(o.AbstractEntityFilter)
	if errAbstractEntityFilter != nil {
		return map[string]interface{}{}, errAbstractEntityFilter
	}
	errAbstractEntityFilter = json.Unmarshal([]byte(serializedAbstractEntityFilter), &toSerialize)
	if errAbstractEntityFilter != nil {
		return map[string]interface{}{}, errAbstractEntityFilter
	}
	if !IsNil(o.ActiveProjects) {
		toSerialize["activeProjects"] = o.ActiveProjects
	}
	if !IsNil(o.AmountOfProjects) {
		toSerialize["amountOfProjects"] = o.AmountOfProjects
	}
	if !IsNil(o.ExperienceInMonth) {
		toSerialize["experienceInMonth"] = o.ExperienceInMonth
	}
	return toSerialize, nil
}

type NullablePersonIndustryFilter struct {
	value *PersonIndustryFilter
	isSet bool
}

func (v NullablePersonIndustryFilter) Get() *PersonIndustryFilter {
	return v.value
}

func (v *NullablePersonIndustryFilter) Set(val *PersonIndustryFilter) {
	v.value = val
	v.isSet = true
}

func (v NullablePersonIndustryFilter) IsSet() bool {
	return v.isSet
}

func (v *NullablePersonIndustryFilter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePersonIndustryFilter(val *PersonIndustryFilter) *NullablePersonIndustryFilter {
	return &NullablePersonIndustryFilter{value: val, isSet: true}
}

func (v NullablePersonIndustryFilter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePersonIndustryFilter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


