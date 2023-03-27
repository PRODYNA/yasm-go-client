/*
YASM (Yet Another Skill Management) API

This is an example of using OAuth2 Implicit Flow in a specification to describe security to your API.

API version: 1.3.13
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// ProjectDetails struct for ProjectDetails
type ProjectDetails struct {
	Project *Project `json:"project,omitempty"`
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
	if o == nil || o.Project == nil {
		var ret Project
		return ret
	}
	return *o.Project
}

// GetProjectOk returns a tuple with the Project field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectDetails) GetProjectOk() (*Project, bool) {
	if o == nil || o.Project == nil {
		return nil, false
	}
	return o.Project, true
}

// HasProject returns a boolean if a field has been set.
func (o *ProjectDetails) HasProject() bool {
	if o != nil && o.Project != nil {
		return true
	}

	return false
}

// SetProject gets a reference to the given Project and assigns it to the Project field.
func (o *ProjectDetails) SetProject(v Project) {
	o.Project = &v
}

// GetOrganization returns the Organization field value if set, zero value otherwise.
func (o *ProjectDetails) GetOrganization() Organization {
	if o == nil || o.Organization == nil {
		var ret Organization
		return ret
	}
	return *o.Organization
}

// GetOrganizationOk returns a tuple with the Organization field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectDetails) GetOrganizationOk() (*Organization, bool) {
	if o == nil || o.Organization == nil {
		return nil, false
	}
	return o.Organization, true
}

// HasOrganization returns a boolean if a field has been set.
func (o *ProjectDetails) HasOrganization() bool {
	if o != nil && o.Organization != nil {
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
	if o == nil || o.Industries == nil {
		var ret []Industry
		return ret
	}
	return o.Industries
}

// GetIndustriesOk returns a tuple with the Industries field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectDetails) GetIndustriesOk() ([]Industry, bool) {
	if o == nil || o.Industries == nil {
		return nil, false
	}
	return o.Industries, true
}

// HasIndustries returns a boolean if a field has been set.
func (o *ProjectDetails) HasIndustries() bool {
	if o != nil && o.Industries != nil {
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
	if o == nil || o.Persons == nil {
		var ret []Person
		return ret
	}
	return o.Persons
}

// GetPersonsOk returns a tuple with the Persons field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectDetails) GetPersonsOk() ([]Person, bool) {
	if o == nil || o.Persons == nil {
		return nil, false
	}
	return o.Persons, true
}

// HasPersons returns a boolean if a field has been set.
func (o *ProjectDetails) HasPersons() bool {
	if o != nil && o.Persons != nil {
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
	if o == nil || o.SkillGroups == nil {
		var ret []SkillGroup
		return ret
	}
	return o.SkillGroups
}

// GetSkillGroupsOk returns a tuple with the SkillGroups field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectDetails) GetSkillGroupsOk() ([]SkillGroup, bool) {
	if o == nil || o.SkillGroups == nil {
		return nil, false
	}
	return o.SkillGroups, true
}

// HasSkillGroups returns a boolean if a field has been set.
func (o *ProjectDetails) HasSkillGroups() bool {
	if o != nil && o.SkillGroups != nil {
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
	if o == nil || o.Timeframe == nil {
		var ret Timeframed
		return ret
	}
	return *o.Timeframe
}

// GetTimeframeOk returns a tuple with the Timeframe field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectDetails) GetTimeframeOk() (*Timeframed, bool) {
	if o == nil || o.Timeframe == nil {
		return nil, false
	}
	return o.Timeframe, true
}

// HasTimeframe returns a boolean if a field has been set.
func (o *ProjectDetails) HasTimeframe() bool {
	if o != nil && o.Timeframe != nil {
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
	if o == nil || o.ObjectType == nil {
		var ret string
		return ret
	}
	return *o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectDetails) GetObjectTypeOk() (*string, bool) {
	if o == nil || o.ObjectType == nil {
		return nil, false
	}
	return o.ObjectType, true
}

// HasObjectType returns a boolean if a field has been set.
func (o *ProjectDetails) HasObjectType() bool {
	if o != nil && o.ObjectType != nil {
		return true
	}

	return false
}

// SetObjectType gets a reference to the given string and assigns it to the ObjectType field.
func (o *ProjectDetails) SetObjectType(v string) {
	o.ObjectType = &v
}

func (o ProjectDetails) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Project != nil {
		toSerialize["project"] = o.Project
	}
	if o.Organization != nil {
		toSerialize["organization"] = o.Organization
	}
	if o.Industries != nil {
		toSerialize["industries"] = o.Industries
	}
	if o.Persons != nil {
		toSerialize["persons"] = o.Persons
	}
	if o.SkillGroups != nil {
		toSerialize["skillGroups"] = o.SkillGroups
	}
	if o.Timeframe != nil {
		toSerialize["timeframe"] = o.Timeframe
	}
	if o.ObjectType != nil {
		toSerialize["objectType"] = o.ObjectType
	}
	return json.Marshal(toSerialize)
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


