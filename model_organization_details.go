/*
YASM (Yet Another Skill Management) API

This is an example of using OAuth2 Implicit Flow in a specification to describe security to your API.

API version: 1.55.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// checks if the OrganizationDetails type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OrganizationDetails{}

// OrganizationDetails struct for OrganizationDetails
type OrganizationDetails struct {
	Audit *Audit `json:"audit,omitempty"`
	Organization *Organization `json:"organization,omitempty"`
	ServiceManager *Person `json:"serviceManager,omitempty"`
	Children []Organization `json:"children,omitempty"`
	Parents []Organization `json:"parents,omitempty"`
	Projects []Project `json:"projects,omitempty"`
	Industries []Industry `json:"industries,omitempty"`
	Certifications []Certification `json:"certifications,omitempty"`
	Awards []Award `json:"awards,omitempty"`
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

// GetAudit returns the Audit field value if set, zero value otherwise.
func (o *OrganizationDetails) GetAudit() Audit {
	if o == nil || IsNil(o.Audit) {
		var ret Audit
		return ret
	}
	return *o.Audit
}

// GetAuditOk returns a tuple with the Audit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationDetails) GetAuditOk() (*Audit, bool) {
	if o == nil || IsNil(o.Audit) {
		return nil, false
	}
	return o.Audit, true
}

// HasAudit returns a boolean if a field has been set.
func (o *OrganizationDetails) HasAudit() bool {
	if o != nil && !IsNil(o.Audit) {
		return true
	}

	return false
}

// SetAudit gets a reference to the given Audit and assigns it to the Audit field.
func (o *OrganizationDetails) SetAudit(v Audit) {
	o.Audit = &v
}

// GetOrganization returns the Organization field value if set, zero value otherwise.
func (o *OrganizationDetails) GetOrganization() Organization {
	if o == nil || IsNil(o.Organization) {
		var ret Organization
		return ret
	}
	return *o.Organization
}

// GetOrganizationOk returns a tuple with the Organization field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationDetails) GetOrganizationOk() (*Organization, bool) {
	if o == nil || IsNil(o.Organization) {
		return nil, false
	}
	return o.Organization, true
}

// HasOrganization returns a boolean if a field has been set.
func (o *OrganizationDetails) HasOrganization() bool {
	if o != nil && !IsNil(o.Organization) {
		return true
	}

	return false
}

// SetOrganization gets a reference to the given Organization and assigns it to the Organization field.
func (o *OrganizationDetails) SetOrganization(v Organization) {
	o.Organization = &v
}

// GetServiceManager returns the ServiceManager field value if set, zero value otherwise.
func (o *OrganizationDetails) GetServiceManager() Person {
	if o == nil || IsNil(o.ServiceManager) {
		var ret Person
		return ret
	}
	return *o.ServiceManager
}

// GetServiceManagerOk returns a tuple with the ServiceManager field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationDetails) GetServiceManagerOk() (*Person, bool) {
	if o == nil || IsNil(o.ServiceManager) {
		return nil, false
	}
	return o.ServiceManager, true
}

// HasServiceManager returns a boolean if a field has been set.
func (o *OrganizationDetails) HasServiceManager() bool {
	if o != nil && !IsNil(o.ServiceManager) {
		return true
	}

	return false
}

// SetServiceManager gets a reference to the given Person and assigns it to the ServiceManager field.
func (o *OrganizationDetails) SetServiceManager(v Person) {
	o.ServiceManager = &v
}

// GetChildren returns the Children field value if set, zero value otherwise.
func (o *OrganizationDetails) GetChildren() []Organization {
	if o == nil || IsNil(o.Children) {
		var ret []Organization
		return ret
	}
	return o.Children
}

// GetChildrenOk returns a tuple with the Children field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationDetails) GetChildrenOk() ([]Organization, bool) {
	if o == nil || IsNil(o.Children) {
		return nil, false
	}
	return o.Children, true
}

// HasChildren returns a boolean if a field has been set.
func (o *OrganizationDetails) HasChildren() bool {
	if o != nil && !IsNil(o.Children) {
		return true
	}

	return false
}

// SetChildren gets a reference to the given []Organization and assigns it to the Children field.
func (o *OrganizationDetails) SetChildren(v []Organization) {
	o.Children = v
}

// GetParents returns the Parents field value if set, zero value otherwise.
func (o *OrganizationDetails) GetParents() []Organization {
	if o == nil || IsNil(o.Parents) {
		var ret []Organization
		return ret
	}
	return o.Parents
}

// GetParentsOk returns a tuple with the Parents field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationDetails) GetParentsOk() ([]Organization, bool) {
	if o == nil || IsNil(o.Parents) {
		return nil, false
	}
	return o.Parents, true
}

// HasParents returns a boolean if a field has been set.
func (o *OrganizationDetails) HasParents() bool {
	if o != nil && !IsNil(o.Parents) {
		return true
	}

	return false
}

// SetParents gets a reference to the given []Organization and assigns it to the Parents field.
func (o *OrganizationDetails) SetParents(v []Organization) {
	o.Parents = v
}

// GetProjects returns the Projects field value if set, zero value otherwise.
func (o *OrganizationDetails) GetProjects() []Project {
	if o == nil || IsNil(o.Projects) {
		var ret []Project
		return ret
	}
	return o.Projects
}

// GetProjectsOk returns a tuple with the Projects field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationDetails) GetProjectsOk() ([]Project, bool) {
	if o == nil || IsNil(o.Projects) {
		return nil, false
	}
	return o.Projects, true
}

// HasProjects returns a boolean if a field has been set.
func (o *OrganizationDetails) HasProjects() bool {
	if o != nil && !IsNil(o.Projects) {
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
	if o == nil || IsNil(o.Industries) {
		var ret []Industry
		return ret
	}
	return o.Industries
}

// GetIndustriesOk returns a tuple with the Industries field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationDetails) GetIndustriesOk() ([]Industry, bool) {
	if o == nil || IsNil(o.Industries) {
		return nil, false
	}
	return o.Industries, true
}

// HasIndustries returns a boolean if a field has been set.
func (o *OrganizationDetails) HasIndustries() bool {
	if o != nil && !IsNil(o.Industries) {
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
	if o == nil || IsNil(o.Certifications) {
		var ret []Certification
		return ret
	}
	return o.Certifications
}

// GetCertificationsOk returns a tuple with the Certifications field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationDetails) GetCertificationsOk() ([]Certification, bool) {
	if o == nil || IsNil(o.Certifications) {
		return nil, false
	}
	return o.Certifications, true
}

// HasCertifications returns a boolean if a field has been set.
func (o *OrganizationDetails) HasCertifications() bool {
	if o != nil && !IsNil(o.Certifications) {
		return true
	}

	return false
}

// SetCertifications gets a reference to the given []Certification and assigns it to the Certifications field.
func (o *OrganizationDetails) SetCertifications(v []Certification) {
	o.Certifications = v
}

// GetAwards returns the Awards field value if set, zero value otherwise.
func (o *OrganizationDetails) GetAwards() []Award {
	if o == nil || IsNil(o.Awards) {
		var ret []Award
		return ret
	}
	return o.Awards
}

// GetAwardsOk returns a tuple with the Awards field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationDetails) GetAwardsOk() ([]Award, bool) {
	if o == nil || IsNil(o.Awards) {
		return nil, false
	}
	return o.Awards, true
}

// HasAwards returns a boolean if a field has been set.
func (o *OrganizationDetails) HasAwards() bool {
	if o != nil && !IsNil(o.Awards) {
		return true
	}

	return false
}

// SetAwards gets a reference to the given []Award and assigns it to the Awards field.
func (o *OrganizationDetails) SetAwards(v []Award) {
	o.Awards = v
}

// GetOffices returns the Offices field value if set, zero value otherwise.
func (o *OrganizationDetails) GetOffices() []Office {
	if o == nil || IsNil(o.Offices) {
		var ret []Office
		return ret
	}
	return o.Offices
}

// GetOfficesOk returns a tuple with the Offices field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationDetails) GetOfficesOk() ([]Office, bool) {
	if o == nil || IsNil(o.Offices) {
		return nil, false
	}
	return o.Offices, true
}

// HasOffices returns a boolean if a field has been set.
func (o *OrganizationDetails) HasOffices() bool {
	if o != nil && !IsNil(o.Offices) {
		return true
	}

	return false
}

// SetOffices gets a reference to the given []Office and assigns it to the Offices field.
func (o *OrganizationDetails) SetOffices(v []Office) {
	o.Offices = v
}

func (o OrganizationDetails) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OrganizationDetails) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Audit) {
		toSerialize["audit"] = o.Audit
	}
	if !IsNil(o.Organization) {
		toSerialize["organization"] = o.Organization
	}
	if !IsNil(o.ServiceManager) {
		toSerialize["serviceManager"] = o.ServiceManager
	}
	if !IsNil(o.Children) {
		toSerialize["children"] = o.Children
	}
	if !IsNil(o.Parents) {
		toSerialize["parents"] = o.Parents
	}
	if !IsNil(o.Projects) {
		toSerialize["projects"] = o.Projects
	}
	if !IsNil(o.Industries) {
		toSerialize["industries"] = o.Industries
	}
	if !IsNil(o.Certifications) {
		toSerialize["certifications"] = o.Certifications
	}
	if !IsNil(o.Awards) {
		toSerialize["awards"] = o.Awards
	}
	if !IsNil(o.Offices) {
		toSerialize["offices"] = o.Offices
	}
	return toSerialize, nil
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


