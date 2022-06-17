/*
YASM (Yet Another Skill Management) API

This is an example of using OAuth2 Implicit Flow in a specification to describe security to your API.

API version: 1.0.16
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// PersonSearchOrganizationsInner struct for PersonSearchOrganizationsInner
type PersonSearchOrganizationsInner struct {
	Id string `json:"id"`
	AmountOfProjects *MinMax `json:"amountOfProjects,omitempty"`
}

// NewPersonSearchOrganizationsInner instantiates a new PersonSearchOrganizationsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPersonSearchOrganizationsInner(id string) *PersonSearchOrganizationsInner {
	this := PersonSearchOrganizationsInner{}
	this.Id = id
	return &this
}

// NewPersonSearchOrganizationsInnerWithDefaults instantiates a new PersonSearchOrganizationsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPersonSearchOrganizationsInnerWithDefaults() *PersonSearchOrganizationsInner {
	this := PersonSearchOrganizationsInner{}
	return &this
}

// GetId returns the Id field value
func (o *PersonSearchOrganizationsInner) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *PersonSearchOrganizationsInner) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *PersonSearchOrganizationsInner) SetId(v string) {
	o.Id = v
}

// GetAmountOfProjects returns the AmountOfProjects field value if set, zero value otherwise.
func (o *PersonSearchOrganizationsInner) GetAmountOfProjects() MinMax {
	if o == nil || o.AmountOfProjects == nil {
		var ret MinMax
		return ret
	}
	return *o.AmountOfProjects
}

// GetAmountOfProjectsOk returns a tuple with the AmountOfProjects field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PersonSearchOrganizationsInner) GetAmountOfProjectsOk() (*MinMax, bool) {
	if o == nil || o.AmountOfProjects == nil {
		return nil, false
	}
	return o.AmountOfProjects, true
}

// HasAmountOfProjects returns a boolean if a field has been set.
func (o *PersonSearchOrganizationsInner) HasAmountOfProjects() bool {
	if o != nil && o.AmountOfProjects != nil {
		return true
	}

	return false
}

// SetAmountOfProjects gets a reference to the given MinMax and assigns it to the AmountOfProjects field.
func (o *PersonSearchOrganizationsInner) SetAmountOfProjects(v MinMax) {
	o.AmountOfProjects = &v
}

func (o PersonSearchOrganizationsInner) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["id"] = o.Id
	}
	if o.AmountOfProjects != nil {
		toSerialize["amountOfProjects"] = o.AmountOfProjects
	}
	return json.Marshal(toSerialize)
}

type NullablePersonSearchOrganizationsInner struct {
	value *PersonSearchOrganizationsInner
	isSet bool
}

func (v NullablePersonSearchOrganizationsInner) Get() *PersonSearchOrganizationsInner {
	return v.value
}

func (v *NullablePersonSearchOrganizationsInner) Set(val *PersonSearchOrganizationsInner) {
	v.value = val
	v.isSet = true
}

func (v NullablePersonSearchOrganizationsInner) IsSet() bool {
	return v.isSet
}

func (v *NullablePersonSearchOrganizationsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePersonSearchOrganizationsInner(val *PersonSearchOrganizationsInner) *NullablePersonSearchOrganizationsInner {
	return &NullablePersonSearchOrganizationsInner{value: val, isSet: true}
}

func (v NullablePersonSearchOrganizationsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePersonSearchOrganizationsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


