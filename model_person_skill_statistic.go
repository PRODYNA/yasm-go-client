/*
YASM (Yet Another Skill Management) API

This is an example of using OAuth2 Implicit Flow in a specification to describe security to your API.

API version: 1.38.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// checks if the PersonSkillStatistic type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PersonSkillStatistic{}

// PersonSkillStatistic struct for PersonSkillStatistic
type PersonSkillStatistic struct {
	Skill *Skill `json:"skill,omitempty"`
	Duration *string `json:"duration,omitempty"`
	Interest *bool `json:"interest,omitempty"`
	UsedInProjects []SkillStatisticProject `json:"usedInProjects,omitempty"`
	TrainedByCertifications []CertificationView `json:"trainedByCertifications,omitempty"`
}

// NewPersonSkillStatistic instantiates a new PersonSkillStatistic object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPersonSkillStatistic() *PersonSkillStatistic {
	this := PersonSkillStatistic{}
	return &this
}

// NewPersonSkillStatisticWithDefaults instantiates a new PersonSkillStatistic object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPersonSkillStatisticWithDefaults() *PersonSkillStatistic {
	this := PersonSkillStatistic{}
	return &this
}

// GetSkill returns the Skill field value if set, zero value otherwise.
func (o *PersonSkillStatistic) GetSkill() Skill {
	if o == nil || IsNil(o.Skill) {
		var ret Skill
		return ret
	}
	return *o.Skill
}

// GetSkillOk returns a tuple with the Skill field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PersonSkillStatistic) GetSkillOk() (*Skill, bool) {
	if o == nil || IsNil(o.Skill) {
		return nil, false
	}
	return o.Skill, true
}

// HasSkill returns a boolean if a field has been set.
func (o *PersonSkillStatistic) HasSkill() bool {
	if o != nil && !IsNil(o.Skill) {
		return true
	}

	return false
}

// SetSkill gets a reference to the given Skill and assigns it to the Skill field.
func (o *PersonSkillStatistic) SetSkill(v Skill) {
	o.Skill = &v
}

// GetDuration returns the Duration field value if set, zero value otherwise.
func (o *PersonSkillStatistic) GetDuration() string {
	if o == nil || IsNil(o.Duration) {
		var ret string
		return ret
	}
	return *o.Duration
}

// GetDurationOk returns a tuple with the Duration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PersonSkillStatistic) GetDurationOk() (*string, bool) {
	if o == nil || IsNil(o.Duration) {
		return nil, false
	}
	return o.Duration, true
}

// HasDuration returns a boolean if a field has been set.
func (o *PersonSkillStatistic) HasDuration() bool {
	if o != nil && !IsNil(o.Duration) {
		return true
	}

	return false
}

// SetDuration gets a reference to the given string and assigns it to the Duration field.
func (o *PersonSkillStatistic) SetDuration(v string) {
	o.Duration = &v
}

// GetInterest returns the Interest field value if set, zero value otherwise.
func (o *PersonSkillStatistic) GetInterest() bool {
	if o == nil || IsNil(o.Interest) {
		var ret bool
		return ret
	}
	return *o.Interest
}

// GetInterestOk returns a tuple with the Interest field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PersonSkillStatistic) GetInterestOk() (*bool, bool) {
	if o == nil || IsNil(o.Interest) {
		return nil, false
	}
	return o.Interest, true
}

// HasInterest returns a boolean if a field has been set.
func (o *PersonSkillStatistic) HasInterest() bool {
	if o != nil && !IsNil(o.Interest) {
		return true
	}

	return false
}

// SetInterest gets a reference to the given bool and assigns it to the Interest field.
func (o *PersonSkillStatistic) SetInterest(v bool) {
	o.Interest = &v
}

// GetUsedInProjects returns the UsedInProjects field value if set, zero value otherwise.
func (o *PersonSkillStatistic) GetUsedInProjects() []SkillStatisticProject {
	if o == nil || IsNil(o.UsedInProjects) {
		var ret []SkillStatisticProject
		return ret
	}
	return o.UsedInProjects
}

// GetUsedInProjectsOk returns a tuple with the UsedInProjects field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PersonSkillStatistic) GetUsedInProjectsOk() ([]SkillStatisticProject, bool) {
	if o == nil || IsNil(o.UsedInProjects) {
		return nil, false
	}
	return o.UsedInProjects, true
}

// HasUsedInProjects returns a boolean if a field has been set.
func (o *PersonSkillStatistic) HasUsedInProjects() bool {
	if o != nil && !IsNil(o.UsedInProjects) {
		return true
	}

	return false
}

// SetUsedInProjects gets a reference to the given []SkillStatisticProject and assigns it to the UsedInProjects field.
func (o *PersonSkillStatistic) SetUsedInProjects(v []SkillStatisticProject) {
	o.UsedInProjects = v
}

// GetTrainedByCertifications returns the TrainedByCertifications field value if set, zero value otherwise.
func (o *PersonSkillStatistic) GetTrainedByCertifications() []CertificationView {
	if o == nil || IsNil(o.TrainedByCertifications) {
		var ret []CertificationView
		return ret
	}
	return o.TrainedByCertifications
}

// GetTrainedByCertificationsOk returns a tuple with the TrainedByCertifications field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PersonSkillStatistic) GetTrainedByCertificationsOk() ([]CertificationView, bool) {
	if o == nil || IsNil(o.TrainedByCertifications) {
		return nil, false
	}
	return o.TrainedByCertifications, true
}

// HasTrainedByCertifications returns a boolean if a field has been set.
func (o *PersonSkillStatistic) HasTrainedByCertifications() bool {
	if o != nil && !IsNil(o.TrainedByCertifications) {
		return true
	}

	return false
}

// SetTrainedByCertifications gets a reference to the given []CertificationView and assigns it to the TrainedByCertifications field.
func (o *PersonSkillStatistic) SetTrainedByCertifications(v []CertificationView) {
	o.TrainedByCertifications = v
}

func (o PersonSkillStatistic) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PersonSkillStatistic) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Skill) {
		toSerialize["skill"] = o.Skill
	}
	if !IsNil(o.Duration) {
		toSerialize["duration"] = o.Duration
	}
	if !IsNil(o.Interest) {
		toSerialize["interest"] = o.Interest
	}
	if !IsNil(o.UsedInProjects) {
		toSerialize["usedInProjects"] = o.UsedInProjects
	}
	if !IsNil(o.TrainedByCertifications) {
		toSerialize["trainedByCertifications"] = o.TrainedByCertifications
	}
	return toSerialize, nil
}

type NullablePersonSkillStatistic struct {
	value *PersonSkillStatistic
	isSet bool
}

func (v NullablePersonSkillStatistic) Get() *PersonSkillStatistic {
	return v.value
}

func (v *NullablePersonSkillStatistic) Set(val *PersonSkillStatistic) {
	v.value = val
	v.isSet = true
}

func (v NullablePersonSkillStatistic) IsSet() bool {
	return v.isSet
}

func (v *NullablePersonSkillStatistic) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePersonSkillStatistic(val *PersonSkillStatistic) *NullablePersonSkillStatistic {
	return &NullablePersonSkillStatistic{value: val, isSet: true}
}

func (v NullablePersonSkillStatistic) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePersonSkillStatistic) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


