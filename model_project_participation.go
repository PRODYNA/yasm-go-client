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

// ProjectParticipation struct for ProjectParticipation
type ProjectParticipation struct {
	BasicDomainModel
	Timeframe *Timeframed `json:"timeframe,omitempty"`
	ProjectDetails *ProjectDetails `json:"projectDetails,omitempty"`
	DescriptionOverwrite *string `json:"descriptionOverwrite,omitempty"`
	PersonalDescription *string `json:"personalDescription,omitempty"`
	Experiences []Experience `json:"experiences"`
}

// NewProjectParticipation instantiates a new ProjectParticipation object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProjectParticipation(experiences []Experience, id string) *ProjectParticipation {
	this := ProjectParticipation{}
	this.Id = id
	this.Experiences = experiences
	return &this
}

// NewProjectParticipationWithDefaults instantiates a new ProjectParticipation object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProjectParticipationWithDefaults() *ProjectParticipation {
	this := ProjectParticipation{}
	return &this
}

// GetTimeframe returns the Timeframe field value if set, zero value otherwise.
func (o *ProjectParticipation) GetTimeframe() Timeframed {
	if o == nil || o.Timeframe == nil {
		var ret Timeframed
		return ret
	}
	return *o.Timeframe
}

// GetTimeframeOk returns a tuple with the Timeframe field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectParticipation) GetTimeframeOk() (*Timeframed, bool) {
	if o == nil || o.Timeframe == nil {
		return nil, false
	}
	return o.Timeframe, true
}

// HasTimeframe returns a boolean if a field has been set.
func (o *ProjectParticipation) HasTimeframe() bool {
	if o != nil && o.Timeframe != nil {
		return true
	}

	return false
}

// SetTimeframe gets a reference to the given Timeframed and assigns it to the Timeframe field.
func (o *ProjectParticipation) SetTimeframe(v Timeframed) {
	o.Timeframe = &v
}

// GetProjectDetails returns the ProjectDetails field value if set, zero value otherwise.
func (o *ProjectParticipation) GetProjectDetails() ProjectDetails {
	if o == nil || o.ProjectDetails == nil {
		var ret ProjectDetails
		return ret
	}
	return *o.ProjectDetails
}

// GetProjectDetailsOk returns a tuple with the ProjectDetails field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectParticipation) GetProjectDetailsOk() (*ProjectDetails, bool) {
	if o == nil || o.ProjectDetails == nil {
		return nil, false
	}
	return o.ProjectDetails, true
}

// HasProjectDetails returns a boolean if a field has been set.
func (o *ProjectParticipation) HasProjectDetails() bool {
	if o != nil && o.ProjectDetails != nil {
		return true
	}

	return false
}

// SetProjectDetails gets a reference to the given ProjectDetails and assigns it to the ProjectDetails field.
func (o *ProjectParticipation) SetProjectDetails(v ProjectDetails) {
	o.ProjectDetails = &v
}

// GetDescriptionOverwrite returns the DescriptionOverwrite field value if set, zero value otherwise.
func (o *ProjectParticipation) GetDescriptionOverwrite() string {
	if o == nil || o.DescriptionOverwrite == nil {
		var ret string
		return ret
	}
	return *o.DescriptionOverwrite
}

// GetDescriptionOverwriteOk returns a tuple with the DescriptionOverwrite field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectParticipation) GetDescriptionOverwriteOk() (*string, bool) {
	if o == nil || o.DescriptionOverwrite == nil {
		return nil, false
	}
	return o.DescriptionOverwrite, true
}

// HasDescriptionOverwrite returns a boolean if a field has been set.
func (o *ProjectParticipation) HasDescriptionOverwrite() bool {
	if o != nil && o.DescriptionOverwrite != nil {
		return true
	}

	return false
}

// SetDescriptionOverwrite gets a reference to the given string and assigns it to the DescriptionOverwrite field.
func (o *ProjectParticipation) SetDescriptionOverwrite(v string) {
	o.DescriptionOverwrite = &v
}

// GetPersonalDescription returns the PersonalDescription field value if set, zero value otherwise.
func (o *ProjectParticipation) GetPersonalDescription() string {
	if o == nil || o.PersonalDescription == nil {
		var ret string
		return ret
	}
	return *o.PersonalDescription
}

// GetPersonalDescriptionOk returns a tuple with the PersonalDescription field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectParticipation) GetPersonalDescriptionOk() (*string, bool) {
	if o == nil || o.PersonalDescription == nil {
		return nil, false
	}
	return o.PersonalDescription, true
}

// HasPersonalDescription returns a boolean if a field has been set.
func (o *ProjectParticipation) HasPersonalDescription() bool {
	if o != nil && o.PersonalDescription != nil {
		return true
	}

	return false
}

// SetPersonalDescription gets a reference to the given string and assigns it to the PersonalDescription field.
func (o *ProjectParticipation) SetPersonalDescription(v string) {
	o.PersonalDescription = &v
}

// GetExperiences returns the Experiences field value
func (o *ProjectParticipation) GetExperiences() []Experience {
	if o == nil {
		var ret []Experience
		return ret
	}

	return o.Experiences
}

// GetExperiencesOk returns a tuple with the Experiences field value
// and a boolean to check if the value has been set.
func (o *ProjectParticipation) GetExperiencesOk() ([]Experience, bool) {
	if o == nil {
		return nil, false
	}
	return o.Experiences, true
}

// SetExperiences sets field value
func (o *ProjectParticipation) SetExperiences(v []Experience) {
	o.Experiences = v
}

func (o ProjectParticipation) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedBasicDomainModel, errBasicDomainModel := json.Marshal(o.BasicDomainModel)
	if errBasicDomainModel != nil {
		return []byte{}, errBasicDomainModel
	}
	errBasicDomainModel = json.Unmarshal([]byte(serializedBasicDomainModel), &toSerialize)
	if errBasicDomainModel != nil {
		return []byte{}, errBasicDomainModel
	}
	if o.Timeframe != nil {
		toSerialize["timeframe"] = o.Timeframe
	}
	if o.ProjectDetails != nil {
		toSerialize["projectDetails"] = o.ProjectDetails
	}
	if o.DescriptionOverwrite != nil {
		toSerialize["descriptionOverwrite"] = o.DescriptionOverwrite
	}
	if o.PersonalDescription != nil {
		toSerialize["personalDescription"] = o.PersonalDescription
	}
	if true {
		toSerialize["experiences"] = o.Experiences
	}
	return json.Marshal(toSerialize)
}

type NullableProjectParticipation struct {
	value *ProjectParticipation
	isSet bool
}

func (v NullableProjectParticipation) Get() *ProjectParticipation {
	return v.value
}

func (v *NullableProjectParticipation) Set(val *ProjectParticipation) {
	v.value = val
	v.isSet = true
}

func (v NullableProjectParticipation) IsSet() bool {
	return v.isSet
}

func (v *NullableProjectParticipation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProjectParticipation(val *ProjectParticipation) *NullableProjectParticipation {
	return &NullableProjectParticipation{value: val, isSet: true}
}

func (v NullableProjectParticipation) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProjectParticipation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


