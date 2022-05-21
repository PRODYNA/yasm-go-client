/*
YASM (Yet Another Skill Management) API

This is an example of using OAuth2 Implicit Flow in a specification to describe security to your API.

API version: 1.0.13
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// IndustryDetails struct for IndustryDetails
type IndustryDetails struct {
	Industry *Industry `json:"industry,omitempty"`
	Organizations []Organization `json:"organizations,omitempty"`
}

// NewIndustryDetails instantiates a new IndustryDetails object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIndustryDetails() *IndustryDetails {
	this := IndustryDetails{}
	return &this
}

// NewIndustryDetailsWithDefaults instantiates a new IndustryDetails object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIndustryDetailsWithDefaults() *IndustryDetails {
	this := IndustryDetails{}
	return &this
}

// GetIndustry returns the Industry field value if set, zero value otherwise.
func (o *IndustryDetails) GetIndustry() Industry {
	if o == nil || o.Industry == nil {
		var ret Industry
		return ret
	}
	return *o.Industry
}

// GetIndustryOk returns a tuple with the Industry field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IndustryDetails) GetIndustryOk() (*Industry, bool) {
	if o == nil || o.Industry == nil {
		return nil, false
	}
	return o.Industry, true
}

// HasIndustry returns a boolean if a field has been set.
func (o *IndustryDetails) HasIndustry() bool {
	if o != nil && o.Industry != nil {
		return true
	}

	return false
}

// SetIndustry gets a reference to the given Industry and assigns it to the Industry field.
func (o *IndustryDetails) SetIndustry(v Industry) {
	o.Industry = &v
}

// GetOrganizations returns the Organizations field value if set, zero value otherwise.
func (o *IndustryDetails) GetOrganizations() []Organization {
	if o == nil || o.Organizations == nil {
		var ret []Organization
		return ret
	}
	return o.Organizations
}

// GetOrganizationsOk returns a tuple with the Organizations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IndustryDetails) GetOrganizationsOk() ([]Organization, bool) {
	if o == nil || o.Organizations == nil {
		return nil, false
	}
	return o.Organizations, true
}

// HasOrganizations returns a boolean if a field has been set.
func (o *IndustryDetails) HasOrganizations() bool {
	if o != nil && o.Organizations != nil {
		return true
	}

	return false
}

// SetOrganizations gets a reference to the given []Organization and assigns it to the Organizations field.
func (o *IndustryDetails) SetOrganizations(v []Organization) {
	o.Organizations = v
}

func (o IndustryDetails) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Industry != nil {
		toSerialize["industry"] = o.Industry
	}
	if o.Organizations != nil {
		toSerialize["organizations"] = o.Organizations
	}
	return json.Marshal(toSerialize)
}

type NullableIndustryDetails struct {
	value *IndustryDetails
	isSet bool
}

func (v NullableIndustryDetails) Get() *IndustryDetails {
	return v.value
}

func (v *NullableIndustryDetails) Set(val *IndustryDetails) {
	v.value = val
	v.isSet = true
}

func (v NullableIndustryDetails) IsSet() bool {
	return v.isSet
}

func (v *NullableIndustryDetails) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIndustryDetails(val *IndustryDetails) *NullableIndustryDetails {
	return &NullableIndustryDetails{value: val, isSet: true}
}

func (v NullableIndustryDetails) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIndustryDetails) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


