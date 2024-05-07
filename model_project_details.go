/*
YASM (Yet Another Skill Management) API

This is an example of using OAuth2 Implicit Flow in a specification to describe security to your API.

API version: 1.25.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// checks if the ProjectDetails type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ProjectDetails{}

// ProjectDetails struct for ProjectDetails
type ProjectDetails struct {
	Project *Project `json:"project,omitempty"`
	ExecutiveOrganizations []Organization `json:"executiveOrganizations,omitempty"`
	Organization *Organization `json:"organization,omitempty"`
	Industries []Industry `json:"industries,omitempty"`
	Persons []Person `json:"persons,omitempty"`
	SkillGroups []SkillGroup `json:"skillGroups,omitempty"`
	Timeframe *Timeframed `json:"timeframe,omitempty"`
	ObjectType *string `json:"objectType,omitempty"`
}

// NewProjectDetails instantiates a new ProjectDetails object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProjectDetails() *ProjectDetails {
	this := ProjectDetails{}
	return &this
}

// NewProjectDetailsWithDefaults instantiates a new ProjectDetails object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProjectDetailsWithDefaults() *ProjectDetails {
	this := ProjectDetails{}
	return &this
}

// GetProject returns the Project field value if set, zero value otherwise.
func (o *ProjectDetails) GetProject() Project {
	if o == nil || IsNil(o.Project) {
		var ret Project
		return ret
	}
	return *o.Project
}

// GetProjectOk returns a tuple with the Project field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectDetails) GetProjectOk() (*Project, bool) {
	if o == nil || IsNil(o.Project) {
		return nil, false
	}
	return o.Project, true
}

// HasProject returns a boolean if a field has been set.
func (o *ProjectDetails) HasProject() bool {
	if o != nil && !IsNil(o.Project) {
		return true
	}

	return false
}

// SetProject gets a reference to the given Project and assigns it to the Project field.
func (o *ProjectDetails) SetProject(v Project) {
	o.Project = &v
}

// GetExecutiveOrganizations returns the ExecutiveOrganizations field value if set, zero value otherwise.
func (o *ProjectDetails) GetExecutiveOrganizations() []Organization {
	if o == nil || IsNil(o.ExecutiveOrganizations) {
		var ret []Organization
		return ret
	}
	return o.ExecutiveOrganizations
}

// GetExecutiveOrganizationsOk returns a tuple with the ExecutiveOrganizations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectDetails) GetExecutiveOrganizationsOk() ([]Organization, bool) {
	if o == nil || IsNil(o.ExecutiveOrganizations) {
		return nil, false
	}
	return o.ExecutiveOrganizations, true
}

// HasExecutiveOrganizations returns a boolean if a field has been set.
func (o *ProjectDetails) HasExecutiveOrganizations() bool {
	if o != nil && !IsNil(o.ExecutiveOrganizations) {
		return true
	}

	return false
}

// SetExecutiveOrganizations gets a reference to the given []Organization and assigns it to the ExecutiveOrganizations field.
func (o *ProjectDetails) SetExecutiveOrganizations(v []Organization) {
	o.ExecutiveOrganizations = v
}

// GetOrganization returns the Organization field value if set, zero value otherwise.
func (o *ProjectDetails) GetOrganization() Organization {
	if o == nil || IsNil(o.Organization) {
		var ret Organization
		return ret
	}
	return *o.Organization
}

// GetOrganizationOk returns a tuple with the Organization field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectDetails) GetOrganizationOk() (*Organization, bool) {
	if o == nil || IsNil(o.Organization) {
		return nil, false
	}
	return o.Organization, true
}

// HasOrganization returns a boolean if a field has been set.
func (o *ProjectDetails) HasOrganization() bool {
	if o != nil && !IsNil(o.Organization) {
		return true
	}

	return false
}

// SetOrganization gets a reference to the given Organization and assigns it to the Organization field.
func (o *ProjectDetails) SetOrganization(v Organization) {
	o.Organization = &v
}

// GetIndustries returns the Industries field value if set, zero value otherwise.
func (o *ProjectDetails) GetIndustries() []Industry {
	if o == nil || IsNil(o.Industries) {
		var ret []Industry
		return ret
	}
	return o.Industries
}

// GetIndustriesOk returns a tuple with the Industries field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectDetails) GetIndustriesOk() ([]Industry, bool) {
	if o == nil || IsNil(o.Industries) {
		return nil, false
	}
	return o.Industries, true
}

// HasIndustries returns a boolean if a field has been set.
func (o *ProjectDetails) HasIndustries() bool {
	if o != nil && !IsNil(o.Industries) {
		return true
	}

	return false
}

// SetIndustries gets a reference to the given []Industry and assigns it to the Industries field.
func (o *ProjectDetails) SetIndustries(v []Industry) {
	o.Industries = v
}

// GetPersons returns the Persons field value if set, zero value otherwise.
func (o *ProjectDetails) GetPersons() []Person {
	if o == nil || IsNil(o.Persons) {
		var ret []Person
		return ret
	}
	return o.Persons
}

// GetPersonsOk returns a tuple with the Persons field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectDetails) GetPersonsOk() ([]Person, bool) {
	if o == nil || IsNil(o.Persons) {
		return nil, false
	}
	return o.Persons, true
}

// HasPersons returns a boolean if a field has been set.
func (o *ProjectDetails) HasPersons() bool {
	if o != nil && !IsNil(o.Persons) {
		return true
	}

	return false
}

// SetPersons gets a reference to the given []Person and assigns it to the Persons field.
func (o *ProjectDetails) SetPersons(v []Person) {
	o.Persons = v
}

// GetSkillGroups returns the SkillGroups field value if set, zero value otherwise.
func (o *ProjectDetails) GetSkillGroups() []SkillGroup {
	if o == nil || IsNil(o.SkillGroups) {
		var ret []SkillGroup
		return ret
	}
	return o.SkillGroups
}

// GetSkillGroupsOk returns a tuple with the SkillGroups field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectDetails) GetSkillGroupsOk() ([]SkillGroup, bool) {
	if o == nil || IsNil(o.SkillGroups) {
		return nil, false
	}
	return o.SkillGroups, true
}

// HasSkillGroups returns a boolean if a field has been set.
func (o *ProjectDetails) HasSkillGroups() bool {
	if o != nil && !IsNil(o.SkillGroups) {
		return true
	}

	return false
}

// SetSkillGroups gets a reference to the given []SkillGroup and assigns it to the SkillGroups field.
func (o *ProjectDetails) SetSkillGroups(v []SkillGroup) {
	o.SkillGroups = v
}

// GetTimeframe returns the Timeframe field value if set, zero value otherwise.
func (o *ProjectDetails) GetTimeframe() Timeframed {
	if o == nil || IsNil(o.Timeframe) {
		var ret Timeframed
		return ret
	}
	return *o.Timeframe
}

// GetTimeframeOk returns a tuple with the Timeframe field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectDetails) GetTimeframeOk() (*Timeframed, bool) {
	if o == nil || IsNil(o.Timeframe) {
		return nil, false
	}
	return o.Timeframe, true
}

// HasTimeframe returns a boolean if a field has been set.
func (o *ProjectDetails) HasTimeframe() bool {
	if o != nil && !IsNil(o.Timeframe) {
		return true
	}

	return false
}

// SetTimeframe gets a reference to the given Timeframed and assigns it to the Timeframe field.
func (o *ProjectDetails) SetTimeframe(v Timeframed) {
	o.Timeframe = &v
}

// GetObjectType returns the ObjectType field value if set, zero value otherwise.
func (o *ProjectDetails) GetObjectType() string {
	if o == nil || IsNil(o.ObjectType) {
		var ret string
		return ret
	}
	return *o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectDetails) GetObjectTypeOk() (*string, bool) {
	if o == nil || IsNil(o.ObjectType) {
		return nil, false
	}
	return o.ObjectType, true
}

// HasObjectType returns a boolean if a field has been set.
func (o *ProjectDetails) HasObjectType() bool {
	if o != nil && !IsNil(o.ObjectType) {
		return true
	}

	return false
}

// SetObjectType gets a reference to the given string and assigns it to the ObjectType field.
func (o *ProjectDetails) SetObjectType(v string) {
	o.ObjectType = &v
}

func (o ProjectDetails) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ProjectDetails) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Project) {
		toSerialize["project"] = o.Project
	}
	if !IsNil(o.ExecutiveOrganizations) {
		toSerialize["executiveOrganizations"] = o.ExecutiveOrganizations
	}
	if !IsNil(o.Organization) {
		toSerialize["organization"] = o.Organization
	}
	if !IsNil(o.Industries) {
		toSerialize["industries"] = o.Industries
	}
	if !IsNil(o.Persons) {
		toSerialize["persons"] = o.Persons
	}
	if !IsNil(o.SkillGroups) {
		toSerialize["skillGroups"] = o.SkillGroups
	}
	if !IsNil(o.Timeframe) {
		toSerialize["timeframe"] = o.Timeframe
	}
	if !IsNil(o.ObjectType) {
		toSerialize["objectType"] = o.ObjectType
	}
	return toSerialize, nil
}

type NullableProjectDetails struct {
	value *ProjectDetails
	isSet bool
}

func (v NullableProjectDetails) Get() *ProjectDetails {
	return v.value
}

func (v *NullableProjectDetails) Set(val *ProjectDetails) {
	v.value = val
	v.isSet = true
}

func (v NullableProjectDetails) IsSet() bool {
	return v.isSet
}

func (v *NullableProjectDetails) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProjectDetails(val *ProjectDetails) *NullableProjectDetails {
	return &NullableProjectDetails{value: val, isSet: true}
}

func (v NullableProjectDetails) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProjectDetails) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


