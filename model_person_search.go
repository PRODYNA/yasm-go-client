/*
YASM (Yet Another Skill Management) API

This is an example of using OAuth2 Implicit Flow in a specification to describe security to your API.

API version: 1.6.6
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// checks if the PersonSearch type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PersonSearch{}

// PersonSearch struct for PersonSearch
type PersonSearch struct {
	PersonIds []string `json:"personIds,omitempty"`
	EmployeeIds []string `json:"employeeIds,omitempty"`
	ProfileIds []string `json:"profileIds,omitempty"`
	OfficeIds []string `json:"officeIds,omitempty"`
	DepartmentNames []string `json:"departmentNames,omitempty"`
	LanguageIds []string `json:"languageIds,omitempty"`
	Availability *AvailabilityFilter `json:"availability,omitempty"`
	OnsiteRatio *MinMaxPercent `json:"onsiteRatio,omitempty"`
	Seniority []Seniority `json:"seniority,omitempty"`
	Skills []PersonSkillFilter `json:"skills,omitempty"`
	Projects []PersonProjectFilter `json:"projects,omitempty"`
	Organizations []PersonOrganizationFilter `json:"organizations,omitempty"`
	Industries []PersonIndustryFilter `json:"industries,omitempty"`
	Certifications []PersonCertificationFilter `json:"certifications,omitempty"`
}

// NewPersonSearch instantiates a new PersonSearch object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPersonSearch() *PersonSearch {
	this := PersonSearch{}
	return &this
}

// NewPersonSearchWithDefaults instantiates a new PersonSearch object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPersonSearchWithDefaults() *PersonSearch {
	this := PersonSearch{}
	return &this
}

// GetPersonIds returns the PersonIds field value if set, zero value otherwise.
func (o *PersonSearch) GetPersonIds() []string {
	if o == nil || IsNil(o.PersonIds) {
		var ret []string
		return ret
	}
	return o.PersonIds
}

// GetPersonIdsOk returns a tuple with the PersonIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PersonSearch) GetPersonIdsOk() ([]string, bool) {
	if o == nil || IsNil(o.PersonIds) {
		return nil, false
	}
	return o.PersonIds, true
}

// HasPersonIds returns a boolean if a field has been set.
func (o *PersonSearch) HasPersonIds() bool {
	if o != nil && !IsNil(o.PersonIds) {
		return true
	}

	return false
}

// SetPersonIds gets a reference to the given []string and assigns it to the PersonIds field.
func (o *PersonSearch) SetPersonIds(v []string) {
	o.PersonIds = v
}

// GetEmployeeIds returns the EmployeeIds field value if set, zero value otherwise.
func (o *PersonSearch) GetEmployeeIds() []string {
	if o == nil || IsNil(o.EmployeeIds) {
		var ret []string
		return ret
	}
	return o.EmployeeIds
}

// GetEmployeeIdsOk returns a tuple with the EmployeeIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PersonSearch) GetEmployeeIdsOk() ([]string, bool) {
	if o == nil || IsNil(o.EmployeeIds) {
		return nil, false
	}
	return o.EmployeeIds, true
}

// HasEmployeeIds returns a boolean if a field has been set.
func (o *PersonSearch) HasEmployeeIds() bool {
	if o != nil && !IsNil(o.EmployeeIds) {
		return true
	}

	return false
}

// SetEmployeeIds gets a reference to the given []string and assigns it to the EmployeeIds field.
func (o *PersonSearch) SetEmployeeIds(v []string) {
	o.EmployeeIds = v
}

// GetProfileIds returns the ProfileIds field value if set, zero value otherwise.
func (o *PersonSearch) GetProfileIds() []string {
	if o == nil || IsNil(o.ProfileIds) {
		var ret []string
		return ret
	}
	return o.ProfileIds
}

// GetProfileIdsOk returns a tuple with the ProfileIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PersonSearch) GetProfileIdsOk() ([]string, bool) {
	if o == nil || IsNil(o.ProfileIds) {
		return nil, false
	}
	return o.ProfileIds, true
}

// HasProfileIds returns a boolean if a field has been set.
func (o *PersonSearch) HasProfileIds() bool {
	if o != nil && !IsNil(o.ProfileIds) {
		return true
	}

	return false
}

// SetProfileIds gets a reference to the given []string and assigns it to the ProfileIds field.
func (o *PersonSearch) SetProfileIds(v []string) {
	o.ProfileIds = v
}

// GetOfficeIds returns the OfficeIds field value if set, zero value otherwise.
func (o *PersonSearch) GetOfficeIds() []string {
	if o == nil || IsNil(o.OfficeIds) {
		var ret []string
		return ret
	}
	return o.OfficeIds
}

// GetOfficeIdsOk returns a tuple with the OfficeIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PersonSearch) GetOfficeIdsOk() ([]string, bool) {
	if o == nil || IsNil(o.OfficeIds) {
		return nil, false
	}
	return o.OfficeIds, true
}

// HasOfficeIds returns a boolean if a field has been set.
func (o *PersonSearch) HasOfficeIds() bool {
	if o != nil && !IsNil(o.OfficeIds) {
		return true
	}

	return false
}

// SetOfficeIds gets a reference to the given []string and assigns it to the OfficeIds field.
func (o *PersonSearch) SetOfficeIds(v []string) {
	o.OfficeIds = v
}

// GetDepartmentNames returns the DepartmentNames field value if set, zero value otherwise.
func (o *PersonSearch) GetDepartmentNames() []string {
	if o == nil || IsNil(o.DepartmentNames) {
		var ret []string
		return ret
	}
	return o.DepartmentNames
}

// GetDepartmentNamesOk returns a tuple with the DepartmentNames field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PersonSearch) GetDepartmentNamesOk() ([]string, bool) {
	if o == nil || IsNil(o.DepartmentNames) {
		return nil, false
	}
	return o.DepartmentNames, true
}

// HasDepartmentNames returns a boolean if a field has been set.
func (o *PersonSearch) HasDepartmentNames() bool {
	if o != nil && !IsNil(o.DepartmentNames) {
		return true
	}

	return false
}

// SetDepartmentNames gets a reference to the given []string and assigns it to the DepartmentNames field.
func (o *PersonSearch) SetDepartmentNames(v []string) {
	o.DepartmentNames = v
}

// GetLanguageIds returns the LanguageIds field value if set, zero value otherwise.
func (o *PersonSearch) GetLanguageIds() []string {
	if o == nil || IsNil(o.LanguageIds) {
		var ret []string
		return ret
	}
	return o.LanguageIds
}

// GetLanguageIdsOk returns a tuple with the LanguageIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PersonSearch) GetLanguageIdsOk() ([]string, bool) {
	if o == nil || IsNil(o.LanguageIds) {
		return nil, false
	}
	return o.LanguageIds, true
}

// HasLanguageIds returns a boolean if a field has been set.
func (o *PersonSearch) HasLanguageIds() bool {
	if o != nil && !IsNil(o.LanguageIds) {
		return true
	}

	return false
}

// SetLanguageIds gets a reference to the given []string and assigns it to the LanguageIds field.
func (o *PersonSearch) SetLanguageIds(v []string) {
	o.LanguageIds = v
}

// GetAvailability returns the Availability field value if set, zero value otherwise.
func (o *PersonSearch) GetAvailability() AvailabilityFilter {
	if o == nil || IsNil(o.Availability) {
		var ret AvailabilityFilter
		return ret
	}
	return *o.Availability
}

// GetAvailabilityOk returns a tuple with the Availability field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PersonSearch) GetAvailabilityOk() (*AvailabilityFilter, bool) {
	if o == nil || IsNil(o.Availability) {
		return nil, false
	}
	return o.Availability, true
}

// HasAvailability returns a boolean if a field has been set.
func (o *PersonSearch) HasAvailability() bool {
	if o != nil && !IsNil(o.Availability) {
		return true
	}

	return false
}

// SetAvailability gets a reference to the given AvailabilityFilter and assigns it to the Availability field.
func (o *PersonSearch) SetAvailability(v AvailabilityFilter) {
	o.Availability = &v
}

// GetOnsiteRatio returns the OnsiteRatio field value if set, zero value otherwise.
func (o *PersonSearch) GetOnsiteRatio() MinMaxPercent {
	if o == nil || IsNil(o.OnsiteRatio) {
		var ret MinMaxPercent
		return ret
	}
	return *o.OnsiteRatio
}

// GetOnsiteRatioOk returns a tuple with the OnsiteRatio field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PersonSearch) GetOnsiteRatioOk() (*MinMaxPercent, bool) {
	if o == nil || IsNil(o.OnsiteRatio) {
		return nil, false
	}
	return o.OnsiteRatio, true
}

// HasOnsiteRatio returns a boolean if a field has been set.
func (o *PersonSearch) HasOnsiteRatio() bool {
	if o != nil && !IsNil(o.OnsiteRatio) {
		return true
	}

	return false
}

// SetOnsiteRatio gets a reference to the given MinMaxPercent and assigns it to the OnsiteRatio field.
func (o *PersonSearch) SetOnsiteRatio(v MinMaxPercent) {
	o.OnsiteRatio = &v
}

// GetSeniority returns the Seniority field value if set, zero value otherwise.
func (o *PersonSearch) GetSeniority() []Seniority {
	if o == nil || IsNil(o.Seniority) {
		var ret []Seniority
		return ret
	}
	return o.Seniority
}

// GetSeniorityOk returns a tuple with the Seniority field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PersonSearch) GetSeniorityOk() ([]Seniority, bool) {
	if o == nil || IsNil(o.Seniority) {
		return nil, false
	}
	return o.Seniority, true
}

// HasSeniority returns a boolean if a field has been set.
func (o *PersonSearch) HasSeniority() bool {
	if o != nil && !IsNil(o.Seniority) {
		return true
	}

	return false
}

// SetSeniority gets a reference to the given []Seniority and assigns it to the Seniority field.
func (o *PersonSearch) SetSeniority(v []Seniority) {
	o.Seniority = v
}

// GetSkills returns the Skills field value if set, zero value otherwise.
func (o *PersonSearch) GetSkills() []PersonSkillFilter {
	if o == nil || IsNil(o.Skills) {
		var ret []PersonSkillFilter
		return ret
	}
	return o.Skills
}

// GetSkillsOk returns a tuple with the Skills field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PersonSearch) GetSkillsOk() ([]PersonSkillFilter, bool) {
	if o == nil || IsNil(o.Skills) {
		return nil, false
	}
	return o.Skills, true
}

// HasSkills returns a boolean if a field has been set.
func (o *PersonSearch) HasSkills() bool {
	if o != nil && !IsNil(o.Skills) {
		return true
	}

	return false
}

// SetSkills gets a reference to the given []PersonSkillFilter and assigns it to the Skills field.
func (o *PersonSearch) SetSkills(v []PersonSkillFilter) {
	o.Skills = v
}

// GetProjects returns the Projects field value if set, zero value otherwise.
func (o *PersonSearch) GetProjects() []PersonProjectFilter {
	if o == nil || IsNil(o.Projects) {
		var ret []PersonProjectFilter
		return ret
	}
	return o.Projects
}

// GetProjectsOk returns a tuple with the Projects field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PersonSearch) GetProjectsOk() ([]PersonProjectFilter, bool) {
	if o == nil || IsNil(o.Projects) {
		return nil, false
	}
	return o.Projects, true
}

// HasProjects returns a boolean if a field has been set.
func (o *PersonSearch) HasProjects() bool {
	if o != nil && !IsNil(o.Projects) {
		return true
	}

	return false
}

// SetProjects gets a reference to the given []PersonProjectFilter and assigns it to the Projects field.
func (o *PersonSearch) SetProjects(v []PersonProjectFilter) {
	o.Projects = v
}

// GetOrganizations returns the Organizations field value if set, zero value otherwise.
func (o *PersonSearch) GetOrganizations() []PersonOrganizationFilter {
	if o == nil || IsNil(o.Organizations) {
		var ret []PersonOrganizationFilter
		return ret
	}
	return o.Organizations
}

// GetOrganizationsOk returns a tuple with the Organizations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PersonSearch) GetOrganizationsOk() ([]PersonOrganizationFilter, bool) {
	if o == nil || IsNil(o.Organizations) {
		return nil, false
	}
	return o.Organizations, true
}

// HasOrganizations returns a boolean if a field has been set.
func (o *PersonSearch) HasOrganizations() bool {
	if o != nil && !IsNil(o.Organizations) {
		return true
	}

	return false
}

// SetOrganizations gets a reference to the given []PersonOrganizationFilter and assigns it to the Organizations field.
func (o *PersonSearch) SetOrganizations(v []PersonOrganizationFilter) {
	o.Organizations = v
}

// GetIndustries returns the Industries field value if set, zero value otherwise.
func (o *PersonSearch) GetIndustries() []PersonIndustryFilter {
	if o == nil || IsNil(o.Industries) {
		var ret []PersonIndustryFilter
		return ret
	}
	return o.Industries
}

// GetIndustriesOk returns a tuple with the Industries field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PersonSearch) GetIndustriesOk() ([]PersonIndustryFilter, bool) {
	if o == nil || IsNil(o.Industries) {
		return nil, false
	}
	return o.Industries, true
}

// HasIndustries returns a boolean if a field has been set.
func (o *PersonSearch) HasIndustries() bool {
	if o != nil && !IsNil(o.Industries) {
		return true
	}

	return false
}

// SetIndustries gets a reference to the given []PersonIndustryFilter and assigns it to the Industries field.
func (o *PersonSearch) SetIndustries(v []PersonIndustryFilter) {
	o.Industries = v
}

// GetCertifications returns the Certifications field value if set, zero value otherwise.
func (o *PersonSearch) GetCertifications() []PersonCertificationFilter {
	if o == nil || IsNil(o.Certifications) {
		var ret []PersonCertificationFilter
		return ret
	}
	return o.Certifications
}

// GetCertificationsOk returns a tuple with the Certifications field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PersonSearch) GetCertificationsOk() ([]PersonCertificationFilter, bool) {
	if o == nil || IsNil(o.Certifications) {
		return nil, false
	}
	return o.Certifications, true
}

// HasCertifications returns a boolean if a field has been set.
func (o *PersonSearch) HasCertifications() bool {
	if o != nil && !IsNil(o.Certifications) {
		return true
	}

	return false
}

// SetCertifications gets a reference to the given []PersonCertificationFilter and assigns it to the Certifications field.
func (o *PersonSearch) SetCertifications(v []PersonCertificationFilter) {
	o.Certifications = v
}

func (o PersonSearch) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PersonSearch) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.PersonIds) {
		toSerialize["personIds"] = o.PersonIds
	}
	if !IsNil(o.EmployeeIds) {
		toSerialize["employeeIds"] = o.EmployeeIds
	}
	if !IsNil(o.ProfileIds) {
		toSerialize["profileIds"] = o.ProfileIds
	}
	if !IsNil(o.OfficeIds) {
		toSerialize["officeIds"] = o.OfficeIds
	}
	if !IsNil(o.DepartmentNames) {
		toSerialize["departmentNames"] = o.DepartmentNames
	}
	if !IsNil(o.LanguageIds) {
		toSerialize["languageIds"] = o.LanguageIds
	}
	if !IsNil(o.Availability) {
		toSerialize["availability"] = o.Availability
	}
	if !IsNil(o.OnsiteRatio) {
		toSerialize["onsiteRatio"] = o.OnsiteRatio
	}
	if !IsNil(o.Seniority) {
		toSerialize["seniority"] = o.Seniority
	}
	if !IsNil(o.Skills) {
		toSerialize["skills"] = o.Skills
	}
	if !IsNil(o.Projects) {
		toSerialize["projects"] = o.Projects
	}
	if !IsNil(o.Organizations) {
		toSerialize["organizations"] = o.Organizations
	}
	if !IsNil(o.Industries) {
		toSerialize["industries"] = o.Industries
	}
	if !IsNil(o.Certifications) {
		toSerialize["certifications"] = o.Certifications
	}
	return toSerialize, nil
}

type NullablePersonSearch struct {
	value *PersonSearch
	isSet bool
}

func (v NullablePersonSearch) Get() *PersonSearch {
	return v.value
}

func (v *NullablePersonSearch) Set(val *PersonSearch) {
	v.value = val
	v.isSet = true
}

func (v NullablePersonSearch) IsSet() bool {
	return v.isSet
}

func (v *NullablePersonSearch) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePersonSearch(val *PersonSearch) *NullablePersonSearch {
	return &NullablePersonSearch{value: val, isSet: true}
}

func (v NullablePersonSearch) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePersonSearch) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


