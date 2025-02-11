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

// checks if the PersonProjectParticipationDetails type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PersonProjectParticipationDetails{}

// PersonProjectParticipationDetails struct for PersonProjectParticipationDetails
type PersonProjectParticipationDetails struct {
	ProjectDetails *ProjectDetails `json:"projectDetails,omitempty"`
	ParticipationItems []PersonProjectParticipationItem `json:"participationItems,omitempty"`
}

// NewPersonProjectParticipationDetails instantiates a new PersonProjectParticipationDetails object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPersonProjectParticipationDetails() *PersonProjectParticipationDetails {
	this := PersonProjectParticipationDetails{}
	return &this
}

// NewPersonProjectParticipationDetailsWithDefaults instantiates a new PersonProjectParticipationDetails object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPersonProjectParticipationDetailsWithDefaults() *PersonProjectParticipationDetails {
	this := PersonProjectParticipationDetails{}
	return &this
}

// GetProjectDetails returns the ProjectDetails field value if set, zero value otherwise.
func (o *PersonProjectParticipationDetails) GetProjectDetails() ProjectDetails {
	if o == nil || IsNil(o.ProjectDetails) {
		var ret ProjectDetails
		return ret
	}
	return *o.ProjectDetails
}

// GetProjectDetailsOk returns a tuple with the ProjectDetails field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PersonProjectParticipationDetails) GetProjectDetailsOk() (*ProjectDetails, bool) {
	if o == nil || IsNil(o.ProjectDetails) {
		return nil, false
	}
	return o.ProjectDetails, true
}

// HasProjectDetails returns a boolean if a field has been set.
func (o *PersonProjectParticipationDetails) HasProjectDetails() bool {
	if o != nil && !IsNil(o.ProjectDetails) {
		return true
	}

	return false
}

// SetProjectDetails gets a reference to the given ProjectDetails and assigns it to the ProjectDetails field.
func (o *PersonProjectParticipationDetails) SetProjectDetails(v ProjectDetails) {
	o.ProjectDetails = &v
}

// GetParticipationItems returns the ParticipationItems field value if set, zero value otherwise.
func (o *PersonProjectParticipationDetails) GetParticipationItems() []PersonProjectParticipationItem {
	if o == nil || IsNil(o.ParticipationItems) {
		var ret []PersonProjectParticipationItem
		return ret
	}
	return o.ParticipationItems
}

// GetParticipationItemsOk returns a tuple with the ParticipationItems field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PersonProjectParticipationDetails) GetParticipationItemsOk() ([]PersonProjectParticipationItem, bool) {
	if o == nil || IsNil(o.ParticipationItems) {
		return nil, false
	}
	return o.ParticipationItems, true
}

// HasParticipationItems returns a boolean if a field has been set.
func (o *PersonProjectParticipationDetails) HasParticipationItems() bool {
	if o != nil && !IsNil(o.ParticipationItems) {
		return true
	}

	return false
}

// SetParticipationItems gets a reference to the given []PersonProjectParticipationItem and assigns it to the ParticipationItems field.
func (o *PersonProjectParticipationDetails) SetParticipationItems(v []PersonProjectParticipationItem) {
	o.ParticipationItems = v
}

func (o PersonProjectParticipationDetails) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PersonProjectParticipationDetails) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ProjectDetails) {
		toSerialize["projectDetails"] = o.ProjectDetails
	}
	if !IsNil(o.ParticipationItems) {
		toSerialize["participationItems"] = o.ParticipationItems
	}
	return toSerialize, nil
}

type NullablePersonProjectParticipationDetails struct {
	value *PersonProjectParticipationDetails
	isSet bool
}

func (v NullablePersonProjectParticipationDetails) Get() *PersonProjectParticipationDetails {
	return v.value
}

func (v *NullablePersonProjectParticipationDetails) Set(val *PersonProjectParticipationDetails) {
	v.value = val
	v.isSet = true
}

func (v NullablePersonProjectParticipationDetails) IsSet() bool {
	return v.isSet
}

func (v *NullablePersonProjectParticipationDetails) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePersonProjectParticipationDetails(val *PersonProjectParticipationDetails) *NullablePersonProjectParticipationDetails {
	return &NullablePersonProjectParticipationDetails{value: val, isSet: true}
}

func (v NullablePersonProjectParticipationDetails) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePersonProjectParticipationDetails) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


