/*
YASM (Yet Another Skill Management) API

This is an example of using OAuth2 Implicit Flow in a specification to describe security to your API.

API version: 1.3.5
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// PersonIndustryFilterAllOf struct for PersonIndustryFilterAllOf
type PersonIndustryFilterAllOf struct {
	ActiveProjects *bool `json:"activeProjects,omitempty"`
	AmountOfProjects *MinMax `json:"amountOfProjects,omitempty"`
	ExperienceInMonth *MinMax `json:"experienceInMonth,omitempty"`
}

// NewPersonIndustryFilterAllOf instantiates a new PersonIndustryFilterAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPersonIndustryFilterAllOf() *PersonIndustryFilterAllOf {
	this := PersonIndustryFilterAllOf{}
	return &this
}

// NewPersonIndustryFilterAllOfWithDefaults instantiates a new PersonIndustryFilterAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPersonIndustryFilterAllOfWithDefaults() *PersonIndustryFilterAllOf {
	this := PersonIndustryFilterAllOf{}
	return &this
}

// GetActiveProjects returns the ActiveProjects field value if set, zero value otherwise.
func (o *PersonIndustryFilterAllOf) GetActiveProjects() bool {
	if o == nil || o.ActiveProjects == nil {
		var ret bool
		return ret
	}
	return *o.ActiveProjects
}

// GetActiveProjectsOk returns a tuple with the ActiveProjects field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PersonIndustryFilterAllOf) GetActiveProjectsOk() (*bool, bool) {
	if o == nil || o.ActiveProjects == nil {
		return nil, false
	}
	return o.ActiveProjects, true
}

// HasActiveProjects returns a boolean if a field has been set.
func (o *PersonIndustryFilterAllOf) HasActiveProjects() bool {
	if o != nil && o.ActiveProjects != nil {
		return true
	}

	return false
}

// SetActiveProjects gets a reference to the given bool and assigns it to the ActiveProjects field.
func (o *PersonIndustryFilterAllOf) SetActiveProjects(v bool) {
	o.ActiveProjects = &v
}

// GetAmountOfProjects returns the AmountOfProjects field value if set, zero value otherwise.
func (o *PersonIndustryFilterAllOf) GetAmountOfProjects() MinMax {
	if o == nil || o.AmountOfProjects == nil {
		var ret MinMax
		return ret
	}
	return *o.AmountOfProjects
}

// GetAmountOfProjectsOk returns a tuple with the AmountOfProjects field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PersonIndustryFilterAllOf) GetAmountOfProjectsOk() (*MinMax, bool) {
	if o == nil || o.AmountOfProjects == nil {
		return nil, false
	}
	return o.AmountOfProjects, true
}

// HasAmountOfProjects returns a boolean if a field has been set.
func (o *PersonIndustryFilterAllOf) HasAmountOfProjects() bool {
	if o != nil && o.AmountOfProjects != nil {
		return true
	}

	return false
}

// SetAmountOfProjects gets a reference to the given MinMax and assigns it to the AmountOfProjects field.
func (o *PersonIndustryFilterAllOf) SetAmountOfProjects(v MinMax) {
	o.AmountOfProjects = &v
}

// GetExperienceInMonth returns the ExperienceInMonth field value if set, zero value otherwise.
func (o *PersonIndustryFilterAllOf) GetExperienceInMonth() MinMax {
	if o == nil || o.ExperienceInMonth == nil {
		var ret MinMax
		return ret
	}
	return *o.ExperienceInMonth
}

// GetExperienceInMonthOk returns a tuple with the ExperienceInMonth field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PersonIndustryFilterAllOf) GetExperienceInMonthOk() (*MinMax, bool) {
	if o == nil || o.ExperienceInMonth == nil {
		return nil, false
	}
	return o.ExperienceInMonth, true
}

// HasExperienceInMonth returns a boolean if a field has been set.
func (o *PersonIndustryFilterAllOf) HasExperienceInMonth() bool {
	if o != nil && o.ExperienceInMonth != nil {
		return true
	}

	return false
}

// SetExperienceInMonth gets a reference to the given MinMax and assigns it to the ExperienceInMonth field.
func (o *PersonIndustryFilterAllOf) SetExperienceInMonth(v MinMax) {
	o.ExperienceInMonth = &v
}

func (o PersonIndustryFilterAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
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

type NullablePersonIndustryFilterAllOf struct {
	value *PersonIndustryFilterAllOf
	isSet bool
}

func (v NullablePersonIndustryFilterAllOf) Get() *PersonIndustryFilterAllOf {
	return v.value
}

func (v *NullablePersonIndustryFilterAllOf) Set(val *PersonIndustryFilterAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullablePersonIndustryFilterAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullablePersonIndustryFilterAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePersonIndustryFilterAllOf(val *PersonIndustryFilterAllOf) *NullablePersonIndustryFilterAllOf {
	return &NullablePersonIndustryFilterAllOf{value: val, isSet: true}
}

func (v NullablePersonIndustryFilterAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePersonIndustryFilterAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


