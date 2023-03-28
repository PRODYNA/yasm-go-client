/*
YASM (Yet Another Skill Management) API

This is an example of using OAuth2 Implicit Flow in a specification to describe security to your API.

API version: 1.3.14
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

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
	if o == nil || o.ActiveProjects == nil {
		var ret bool
		return ret
	}
	return *o.ActiveProjects
}

// GetActiveProjectsOk returns a tuple with the ActiveProjects field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PersonIndustryFilter) GetActiveProjectsOk() (*bool, bool) {
	if o == nil || o.ActiveProjects == nil {
		return nil, false
	}
	return o.ActiveProjects, true
}

// HasActiveProjects returns a boolean if a field has been set.
func (o *PersonIndustryFilter) HasActiveProjects() bool {
	if o != nil && o.ActiveProjects != nil {
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
	if o == nil || o.AmountOfProjects == nil {
		var ret MinMax
		return ret
	}
	return *o.AmountOfProjects
}

// GetAmountOfProjectsOk returns a tuple with the AmountOfProjects field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PersonIndustryFilter) GetAmountOfProjectsOk() (*MinMax, bool) {
	if o == nil || o.AmountOfProjects == nil {
		return nil, false
	}
	return o.AmountOfProjects, true
}

// HasAmountOfProjects returns a boolean if a field has been set.
func (o *PersonIndustryFilter) HasAmountOfProjects() bool {
	if o != nil && o.AmountOfProjects != nil {
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
	if o == nil || o.ExperienceInMonth == nil {
		var ret MinMax
		return ret
	}
	return *o.ExperienceInMonth
}

// GetExperienceInMonthOk returns a tuple with the ExperienceInMonth field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PersonIndustryFilter) GetExperienceInMonthOk() (*MinMax, bool) {
	if o == nil || o.ExperienceInMonth == nil {
		return nil, false
	}
	return o.ExperienceInMonth, true
}

// HasExperienceInMonth returns a boolean if a field has been set.
func (o *PersonIndustryFilter) HasExperienceInMonth() bool {
	if o != nil && o.ExperienceInMonth != nil {
		return true
	}

	return false
}

// SetExperienceInMonth gets a reference to the given MinMax and assigns it to the ExperienceInMonth field.
func (o *PersonIndustryFilter) SetExperienceInMonth(v MinMax) {
	o.ExperienceInMonth = &v
}

func (o PersonIndustryFilter) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedAbstractEntityFilter, errAbstractEntityFilter := json.Marshal(o.AbstractEntityFilter)
	if errAbstractEntityFilter != nil {
		return []byte{}, errAbstractEntityFilter
	}
	errAbstractEntityFilter = json.Unmarshal([]byte(serializedAbstractEntityFilter), &toSerialize)
	if errAbstractEntityFilter != nil {
		return []byte{}, errAbstractEntityFilter
	}
	if o.ActiveProjects != nil {
		toSerialize["activeProjects"] = o.ActiveProjects
	}
	if o.AmountOfProjects != nil {
		toSerialize["amountOfProjects"] = o.AmountOfProjects
	}
	if o.ExperienceInMonth != nil {
		toSerialize["experienceInMonth"] = o.ExperienceInMonth
	}
	return json.Marshal(toSerialize)
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


