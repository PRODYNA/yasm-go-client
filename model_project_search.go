/*
YASM (Yet Another Skill Management) API

This is an example of using OAuth2 Implicit Flow in a specification to describe security to your API.

API version: 1.31.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// checks if the ProjectSearch type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ProjectSearch{}

// ProjectSearch struct for ProjectSearch
type ProjectSearch struct {
	Search
	ProjectIds []string `json:"projectIds,omitempty"`
	MinStartDate *string `json:"minStartDate,omitempty"`
	MaxEndDate *string `json:"maxEndDate,omitempty"`
	ProjectStatuses []ProjectStatus `json:"projectStatuses,omitempty"`
	ParticipationAmountInMonths *MinMax `json:"participationAmountInMonths,omitempty"`
	InvolvedCountryIds []string `json:"involvedCountryIds,omitempty"`
	OrganizationCountryIds []string `json:"organizationCountryIds,omitempty"`
	AmountOfInvolvedPersons *MinMax `json:"amountOfInvolvedPersons,omitempty"`
	SkillIds []string `json:"skillIds,omitempty"`
	ProjectType []ProjectType `json:"projectType,omitempty"`
	Confidentiality []Confidentiality `json:"confidentiality,omitempty"`
	ParticipantIds []string `json:"participantIds,omitempty"`
	IndustryIds []string `json:"industryIds,omitempty"`
	OrganizationIds []string `json:"organizationIds,omitempty"`
}

// NewProjectSearch instantiates a new ProjectSearch object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProjectSearch(skip int32, limit int32) *ProjectSearch {
	this := ProjectSearch{}
	this.Skip = skip
	this.Limit = limit
	return &this
}

// NewProjectSearchWithDefaults instantiates a new ProjectSearch object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProjectSearchWithDefaults() *ProjectSearch {
	this := ProjectSearch{}
	return &this
}

// GetProjectIds returns the ProjectIds field value if set, zero value otherwise.
func (o *ProjectSearch) GetProjectIds() []string {
	if o == nil || IsNil(o.ProjectIds) {
		var ret []string
		return ret
	}
	return o.ProjectIds
}

// GetProjectIdsOk returns a tuple with the ProjectIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectSearch) GetProjectIdsOk() ([]string, bool) {
	if o == nil || IsNil(o.ProjectIds) {
		return nil, false
	}
	return o.ProjectIds, true
}

// HasProjectIds returns a boolean if a field has been set.
func (o *ProjectSearch) HasProjectIds() bool {
	if o != nil && !IsNil(o.ProjectIds) {
		return true
	}

	return false
}

// SetProjectIds gets a reference to the given []string and assigns it to the ProjectIds field.
func (o *ProjectSearch) SetProjectIds(v []string) {
	o.ProjectIds = v
}

// GetMinStartDate returns the MinStartDate field value if set, zero value otherwise.
func (o *ProjectSearch) GetMinStartDate() string {
	if o == nil || IsNil(o.MinStartDate) {
		var ret string
		return ret
	}
	return *o.MinStartDate
}

// GetMinStartDateOk returns a tuple with the MinStartDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectSearch) GetMinStartDateOk() (*string, bool) {
	if o == nil || IsNil(o.MinStartDate) {
		return nil, false
	}
	return o.MinStartDate, true
}

// HasMinStartDate returns a boolean if a field has been set.
func (o *ProjectSearch) HasMinStartDate() bool {
	if o != nil && !IsNil(o.MinStartDate) {
		return true
	}

	return false
}

// SetMinStartDate gets a reference to the given string and assigns it to the MinStartDate field.
func (o *ProjectSearch) SetMinStartDate(v string) {
	o.MinStartDate = &v
}

// GetMaxEndDate returns the MaxEndDate field value if set, zero value otherwise.
func (o *ProjectSearch) GetMaxEndDate() string {
	if o == nil || IsNil(o.MaxEndDate) {
		var ret string
		return ret
	}
	return *o.MaxEndDate
}

// GetMaxEndDateOk returns a tuple with the MaxEndDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectSearch) GetMaxEndDateOk() (*string, bool) {
	if o == nil || IsNil(o.MaxEndDate) {
		return nil, false
	}
	return o.MaxEndDate, true
}

// HasMaxEndDate returns a boolean if a field has been set.
func (o *ProjectSearch) HasMaxEndDate() bool {
	if o != nil && !IsNil(o.MaxEndDate) {
		return true
	}

	return false
}

// SetMaxEndDate gets a reference to the given string and assigns it to the MaxEndDate field.
func (o *ProjectSearch) SetMaxEndDate(v string) {
	o.MaxEndDate = &v
}

// GetProjectStatuses returns the ProjectStatuses field value if set, zero value otherwise.
func (o *ProjectSearch) GetProjectStatuses() []ProjectStatus {
	if o == nil || IsNil(o.ProjectStatuses) {
		var ret []ProjectStatus
		return ret
	}
	return o.ProjectStatuses
}

// GetProjectStatusesOk returns a tuple with the ProjectStatuses field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectSearch) GetProjectStatusesOk() ([]ProjectStatus, bool) {
	if o == nil || IsNil(o.ProjectStatuses) {
		return nil, false
	}
	return o.ProjectStatuses, true
}

// HasProjectStatuses returns a boolean if a field has been set.
func (o *ProjectSearch) HasProjectStatuses() bool {
	if o != nil && !IsNil(o.ProjectStatuses) {
		return true
	}

	return false
}

// SetProjectStatuses gets a reference to the given []ProjectStatus and assigns it to the ProjectStatuses field.
func (o *ProjectSearch) SetProjectStatuses(v []ProjectStatus) {
	o.ProjectStatuses = v
}

// GetParticipationAmountInMonths returns the ParticipationAmountInMonths field value if set, zero value otherwise.
func (o *ProjectSearch) GetParticipationAmountInMonths() MinMax {
	if o == nil || IsNil(o.ParticipationAmountInMonths) {
		var ret MinMax
		return ret
	}
	return *o.ParticipationAmountInMonths
}

// GetParticipationAmountInMonthsOk returns a tuple with the ParticipationAmountInMonths field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectSearch) GetParticipationAmountInMonthsOk() (*MinMax, bool) {
	if o == nil || IsNil(o.ParticipationAmountInMonths) {
		return nil, false
	}
	return o.ParticipationAmountInMonths, true
}

// HasParticipationAmountInMonths returns a boolean if a field has been set.
func (o *ProjectSearch) HasParticipationAmountInMonths() bool {
	if o != nil && !IsNil(o.ParticipationAmountInMonths) {
		return true
	}

	return false
}

// SetParticipationAmountInMonths gets a reference to the given MinMax and assigns it to the ParticipationAmountInMonths field.
func (o *ProjectSearch) SetParticipationAmountInMonths(v MinMax) {
	o.ParticipationAmountInMonths = &v
}

// GetInvolvedCountryIds returns the InvolvedCountryIds field value if set, zero value otherwise.
func (o *ProjectSearch) GetInvolvedCountryIds() []string {
	if o == nil || IsNil(o.InvolvedCountryIds) {
		var ret []string
		return ret
	}
	return o.InvolvedCountryIds
}

// GetInvolvedCountryIdsOk returns a tuple with the InvolvedCountryIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectSearch) GetInvolvedCountryIdsOk() ([]string, bool) {
	if o == nil || IsNil(o.InvolvedCountryIds) {
		return nil, false
	}
	return o.InvolvedCountryIds, true
}

// HasInvolvedCountryIds returns a boolean if a field has been set.
func (o *ProjectSearch) HasInvolvedCountryIds() bool {
	if o != nil && !IsNil(o.InvolvedCountryIds) {
		return true
	}

	return false
}

// SetInvolvedCountryIds gets a reference to the given []string and assigns it to the InvolvedCountryIds field.
func (o *ProjectSearch) SetInvolvedCountryIds(v []string) {
	o.InvolvedCountryIds = v
}

// GetOrganizationCountryIds returns the OrganizationCountryIds field value if set, zero value otherwise.
func (o *ProjectSearch) GetOrganizationCountryIds() []string {
	if o == nil || IsNil(o.OrganizationCountryIds) {
		var ret []string
		return ret
	}
	return o.OrganizationCountryIds
}

// GetOrganizationCountryIdsOk returns a tuple with the OrganizationCountryIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectSearch) GetOrganizationCountryIdsOk() ([]string, bool) {
	if o == nil || IsNil(o.OrganizationCountryIds) {
		return nil, false
	}
	return o.OrganizationCountryIds, true
}

// HasOrganizationCountryIds returns a boolean if a field has been set.
func (o *ProjectSearch) HasOrganizationCountryIds() bool {
	if o != nil && !IsNil(o.OrganizationCountryIds) {
		return true
	}

	return false
}

// SetOrganizationCountryIds gets a reference to the given []string and assigns it to the OrganizationCountryIds field.
func (o *ProjectSearch) SetOrganizationCountryIds(v []string) {
	o.OrganizationCountryIds = v
}

// GetAmountOfInvolvedPersons returns the AmountOfInvolvedPersons field value if set, zero value otherwise.
func (o *ProjectSearch) GetAmountOfInvolvedPersons() MinMax {
	if o == nil || IsNil(o.AmountOfInvolvedPersons) {
		var ret MinMax
		return ret
	}
	return *o.AmountOfInvolvedPersons
}

// GetAmountOfInvolvedPersonsOk returns a tuple with the AmountOfInvolvedPersons field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectSearch) GetAmountOfInvolvedPersonsOk() (*MinMax, bool) {
	if o == nil || IsNil(o.AmountOfInvolvedPersons) {
		return nil, false
	}
	return o.AmountOfInvolvedPersons, true
}

// HasAmountOfInvolvedPersons returns a boolean if a field has been set.
func (o *ProjectSearch) HasAmountOfInvolvedPersons() bool {
	if o != nil && !IsNil(o.AmountOfInvolvedPersons) {
		return true
	}

	return false
}

// SetAmountOfInvolvedPersons gets a reference to the given MinMax and assigns it to the AmountOfInvolvedPersons field.
func (o *ProjectSearch) SetAmountOfInvolvedPersons(v MinMax) {
	o.AmountOfInvolvedPersons = &v
}

// GetSkillIds returns the SkillIds field value if set, zero value otherwise.
func (o *ProjectSearch) GetSkillIds() []string {
	if o == nil || IsNil(o.SkillIds) {
		var ret []string
		return ret
	}
	return o.SkillIds
}

// GetSkillIdsOk returns a tuple with the SkillIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectSearch) GetSkillIdsOk() ([]string, bool) {
	if o == nil || IsNil(o.SkillIds) {
		return nil, false
	}
	return o.SkillIds, true
}

// HasSkillIds returns a boolean if a field has been set.
func (o *ProjectSearch) HasSkillIds() bool {
	if o != nil && !IsNil(o.SkillIds) {
		return true
	}

	return false
}

// SetSkillIds gets a reference to the given []string and assigns it to the SkillIds field.
func (o *ProjectSearch) SetSkillIds(v []string) {
	o.SkillIds = v
}

// GetProjectType returns the ProjectType field value if set, zero value otherwise.
func (o *ProjectSearch) GetProjectType() []ProjectType {
	if o == nil || IsNil(o.ProjectType) {
		var ret []ProjectType
		return ret
	}
	return o.ProjectType
}

// GetProjectTypeOk returns a tuple with the ProjectType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectSearch) GetProjectTypeOk() ([]ProjectType, bool) {
	if o == nil || IsNil(o.ProjectType) {
		return nil, false
	}
	return o.ProjectType, true
}

// HasProjectType returns a boolean if a field has been set.
func (o *ProjectSearch) HasProjectType() bool {
	if o != nil && !IsNil(o.ProjectType) {
		return true
	}

	return false
}

// SetProjectType gets a reference to the given []ProjectType and assigns it to the ProjectType field.
func (o *ProjectSearch) SetProjectType(v []ProjectType) {
	o.ProjectType = v
}

// GetConfidentiality returns the Confidentiality field value if set, zero value otherwise.
func (o *ProjectSearch) GetConfidentiality() []Confidentiality {
	if o == nil || IsNil(o.Confidentiality) {
		var ret []Confidentiality
		return ret
	}
	return o.Confidentiality
}

// GetConfidentialityOk returns a tuple with the Confidentiality field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectSearch) GetConfidentialityOk() ([]Confidentiality, bool) {
	if o == nil || IsNil(o.Confidentiality) {
		return nil, false
	}
	return o.Confidentiality, true
}

// HasConfidentiality returns a boolean if a field has been set.
func (o *ProjectSearch) HasConfidentiality() bool {
	if o != nil && !IsNil(o.Confidentiality) {
		return true
	}

	return false
}

// SetConfidentiality gets a reference to the given []Confidentiality and assigns it to the Confidentiality field.
func (o *ProjectSearch) SetConfidentiality(v []Confidentiality) {
	o.Confidentiality = v
}

// GetParticipantIds returns the ParticipantIds field value if set, zero value otherwise.
func (o *ProjectSearch) GetParticipantIds() []string {
	if o == nil || IsNil(o.ParticipantIds) {
		var ret []string
		return ret
	}
	return o.ParticipantIds
}

// GetParticipantIdsOk returns a tuple with the ParticipantIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectSearch) GetParticipantIdsOk() ([]string, bool) {
	if o == nil || IsNil(o.ParticipantIds) {
		return nil, false
	}
	return o.ParticipantIds, true
}

// HasParticipantIds returns a boolean if a field has been set.
func (o *ProjectSearch) HasParticipantIds() bool {
	if o != nil && !IsNil(o.ParticipantIds) {
		return true
	}

	return false
}

// SetParticipantIds gets a reference to the given []string and assigns it to the ParticipantIds field.
func (o *ProjectSearch) SetParticipantIds(v []string) {
	o.ParticipantIds = v
}

// GetIndustryIds returns the IndustryIds field value if set, zero value otherwise.
func (o *ProjectSearch) GetIndustryIds() []string {
	if o == nil || IsNil(o.IndustryIds) {
		var ret []string
		return ret
	}
	return o.IndustryIds
}

// GetIndustryIdsOk returns a tuple with the IndustryIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectSearch) GetIndustryIdsOk() ([]string, bool) {
	if o == nil || IsNil(o.IndustryIds) {
		return nil, false
	}
	return o.IndustryIds, true
}

// HasIndustryIds returns a boolean if a field has been set.
func (o *ProjectSearch) HasIndustryIds() bool {
	if o != nil && !IsNil(o.IndustryIds) {
		return true
	}

	return false
}

// SetIndustryIds gets a reference to the given []string and assigns it to the IndustryIds field.
func (o *ProjectSearch) SetIndustryIds(v []string) {
	o.IndustryIds = v
}

// GetOrganizationIds returns the OrganizationIds field value if set, zero value otherwise.
func (o *ProjectSearch) GetOrganizationIds() []string {
	if o == nil || IsNil(o.OrganizationIds) {
		var ret []string
		return ret
	}
	return o.OrganizationIds
}

// GetOrganizationIdsOk returns a tuple with the OrganizationIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectSearch) GetOrganizationIdsOk() ([]string, bool) {
	if o == nil || IsNil(o.OrganizationIds) {
		return nil, false
	}
	return o.OrganizationIds, true
}

// HasOrganizationIds returns a boolean if a field has been set.
func (o *ProjectSearch) HasOrganizationIds() bool {
	if o != nil && !IsNil(o.OrganizationIds) {
		return true
	}

	return false
}

// SetOrganizationIds gets a reference to the given []string and assigns it to the OrganizationIds field.
func (o *ProjectSearch) SetOrganizationIds(v []string) {
	o.OrganizationIds = v
}

func (o ProjectSearch) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ProjectSearch) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	serializedSearch, errSearch := json.Marshal(o.Search)
	if errSearch != nil {
		return map[string]interface{}{}, errSearch
	}
	errSearch = json.Unmarshal([]byte(serializedSearch), &toSerialize)
	if errSearch != nil {
		return map[string]interface{}{}, errSearch
	}
	if !IsNil(o.ProjectIds) {
		toSerialize["projectIds"] = o.ProjectIds
	}
	if !IsNil(o.MinStartDate) {
		toSerialize["minStartDate"] = o.MinStartDate
	}
	if !IsNil(o.MaxEndDate) {
		toSerialize["maxEndDate"] = o.MaxEndDate
	}
	if !IsNil(o.ProjectStatuses) {
		toSerialize["projectStatuses"] = o.ProjectStatuses
	}
	if !IsNil(o.ParticipationAmountInMonths) {
		toSerialize["participationAmountInMonths"] = o.ParticipationAmountInMonths
	}
	if !IsNil(o.InvolvedCountryIds) {
		toSerialize["involvedCountryIds"] = o.InvolvedCountryIds
	}
	if !IsNil(o.OrganizationCountryIds) {
		toSerialize["organizationCountryIds"] = o.OrganizationCountryIds
	}
	if !IsNil(o.AmountOfInvolvedPersons) {
		toSerialize["amountOfInvolvedPersons"] = o.AmountOfInvolvedPersons
	}
	if !IsNil(o.SkillIds) {
		toSerialize["skillIds"] = o.SkillIds
	}
	if !IsNil(o.ProjectType) {
		toSerialize["projectType"] = o.ProjectType
	}
	if !IsNil(o.Confidentiality) {
		toSerialize["confidentiality"] = o.Confidentiality
	}
	if !IsNil(o.ParticipantIds) {
		toSerialize["participantIds"] = o.ParticipantIds
	}
	if !IsNil(o.IndustryIds) {
		toSerialize["industryIds"] = o.IndustryIds
	}
	if !IsNil(o.OrganizationIds) {
		toSerialize["organizationIds"] = o.OrganizationIds
	}
	return toSerialize, nil
}

type NullableProjectSearch struct {
	value *ProjectSearch
	isSet bool
}

func (v NullableProjectSearch) Get() *ProjectSearch {
	return v.value
}

func (v *NullableProjectSearch) Set(val *ProjectSearch) {
	v.value = val
	v.isSet = true
}

func (v NullableProjectSearch) IsSet() bool {
	return v.isSet
}

func (v *NullableProjectSearch) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProjectSearch(val *ProjectSearch) *NullableProjectSearch {
	return &NullableProjectSearch{value: val, isSet: true}
}

func (v NullableProjectSearch) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProjectSearch) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


