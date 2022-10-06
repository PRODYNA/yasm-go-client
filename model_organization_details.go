/*
YASM (Yet Another Skill Management) API

This is an example of using OAuth2 Implicit Flow in a specification to describe security to your API.

API version: 1.2.5
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// OrganizationDetails struct for OrganizationDetails
type OrganizationDetails struct {
	Organization *Organization `json:"organization,omitempty"`
	Projects []Project `json:"projects,omitempty"`
	Industries []Industry `json:"industries,omitempty"`
	Certifications []Certification `json:"certifications,omitempty"`
	Offices []Office `json:"offices,omitempty"`
}

// NewOrganizationDetails instantiates a new OrganizationDetails object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrganizationDetails() *OrganizationDetails {
	this := OrganizationDetails{}
	return &this
}

// NewOrganizationDetailsWithDefaults instantiates a new OrganizationDetails object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrganizationDetailsWithDefaults() *OrganizationDetails {
	this := OrganizationDetails{}
	return &this
}

// GetOrganization returns the Organization field value if set, zero value otherwise.
func (o *OrganizationDetails) GetOrganization() Organization {
	if o == nil || o.Organization == nil {
		var ret Organization
		return ret
	}
	return *o.Organization
}

// GetOrganizationOk returns a tuple with the Organization field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationDetails) GetOrganizationOk() (*Organization, bool) {
	if o == nil || o.Organization == nil {
		return nil, false
	}
	return o.Organization, true
}

// HasOrganization returns a boolean if a field has been set.
func (o *OrganizationDetails) HasOrganization() bool {
	if o != nil && o.Organization != nil {
		return true
	}

	return false
}

// SetOrganization gets a reference to the given Organization and assigns it to the Organization field.
func (o *OrganizationDetails) SetOrganization(v Organization) {
	o.Organization = &v
}

// GetProjects returns the Projects field value if set, zero value otherwise.
func (o *OrganizationDetails) GetProjects() []Project {
	if o == nil || o.Projects == nil {
		var ret []Project
		return ret
	}
	return o.Projects
}

// GetProjectsOk returns a tuple with the Projects field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationDetails) GetProjectsOk() ([]Project, bool) {
	if o == nil || o.Projects == nil {
		return nil, false
	}
	return o.Projects, true
}

// HasProjects returns a boolean if a field has been set.
func (o *OrganizationDetails) HasProjects() bool {
	if o != nil && o.Projects != nil {
		return true
	}

	return false
}

// SetProjects gets a reference to the given []Project and assigns it to the Projects field.
func (o *OrganizationDetails) SetProjects(v []Project) {
	o.Projects = v
}

// GetIndustries returns the Industries field value if set, zero value otherwise.
func (o *OrganizationDetails) GetIndustries() []Industry {
	if o == nil || o.Industries == nil {
		var ret []Industry
		return ret
	}
	return o.Industries
}

// GetIndustriesOk returns a tuple with the Industries field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationDetails) GetIndustriesOk() ([]Industry, bool) {
	if o == nil || o.Industries == nil {
		return nil, false
	}
	return o.Industries, true
}

// HasIndustries returns a boolean if a field has been set.
func (o *OrganizationDetails) HasIndustries() bool {
	if o != nil && o.Industries != nil {
		return true
	}

	return false
}

// SetIndustries gets a reference to the given []Industry and assigns it to the Industries field.
func (o *OrganizationDetails) SetIndustries(v []Industry) {
	o.Industries = v
}

// GetCertifications returns the Certifications field value if set, zero value otherwise.
func (o *OrganizationDetails) GetCertifications() []Certification {
	if o == nil || o.Certifications == nil {
		var ret []Certification
		return ret
	}
	return o.Certifications
}

// GetCertificationsOk returns a tuple with the Certifications field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationDetails) GetCertificationsOk() ([]Certification, bool) {
	if o == nil || o.Certifications == nil {
		return nil, false
	}
	return o.Certifications, true
}

// HasCertifications returns a boolean if a field has been set.
func (o *OrganizationDetails) HasCertifications() bool {
	if o != nil && o.Certifications != nil {
		return true
	}

	return false
}

// SetCertifications gets a reference to the given []Certification and assigns it to the Certifications field.
func (o *OrganizationDetails) SetCertifications(v []Certification) {
	o.Certifications = v
}

// GetOffices returns the Offices field value if set, zero value otherwise.
func (o *OrganizationDetails) GetOffices() []Office {
	if o == nil || o.Offices == nil {
		var ret []Office
		return ret
	}
	return o.Offices
}

// GetOfficesOk returns a tuple with the Offices field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationDetails) GetOfficesOk() ([]Office, bool) {
	if o == nil || o.Offices == nil {
		return nil, false
	}
	return o.Offices, true
}

// HasOffices returns a boolean if a field has been set.
func (o *OrganizationDetails) HasOffices() bool {
	if o != nil && o.Offices != nil {
		return true
	}

	return false
}

// SetOffices gets a reference to the given []Office and assigns it to the Offices field.
func (o *OrganizationDetails) SetOffices(v []Office) {
	o.Offices = v
}

func (o OrganizationDetails) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Organization != nil {
		toSerialize["organization"] = o.Organization
	}
	if o.Projects != nil {
		toSerialize["projects"] = o.Projects
	}
	if o.Industries != nil {
		toSerialize["industries"] = o.Industries
	}
	if o.Certifications != nil {
		toSerialize["certifications"] = o.Certifications
	}
	if o.Offices != nil {
		toSerialize["offices"] = o.Offices
	}
	return json.Marshal(toSerialize)
}

type NullableOrganizationDetails struct {
	value *OrganizationDetails
	isSet bool
}

func (v NullableOrganizationDetails) Get() *OrganizationDetails {
	return v.value
}

func (v *NullableOrganizationDetails) Set(val *OrganizationDetails) {
	v.value = val
	v.isSet = true
}

func (v NullableOrganizationDetails) IsSet() bool {
	return v.isSet
}

func (v *NullableOrganizationDetails) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrganizationDetails(val *OrganizationDetails) *NullableOrganizationDetails {
	return &NullableOrganizationDetails{value: val, isSet: true}
}

func (v NullableOrganizationDetails) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrganizationDetails) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


