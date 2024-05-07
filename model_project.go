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

// checks if the Project type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Project{}

// Project struct for Project
type Project struct {
	NamedDomainModel
	Synonyms []string `json:"synonyms,omitempty"`
	Location *string `json:"location,omitempty"`
	Geolocation *Geolocation `json:"geolocation,omitempty"`
	Description *string `json:"description,omitempty"`
	// true if project was done outside of the organization
	External *bool `json:"external,omitempty"`
	ProjectType *ProjectType `json:"projectType,omitempty"`
	Confidentiality *Confidentiality `json:"confidentiality,omitempty"`
}

// NewProject instantiates a new Project object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProject(id string, name string) *Project {
	this := Project{}
	this.Id = id
	this.Name = name
	var external bool = false
	this.External = &external
	return &this
}

// NewProjectWithDefaults instantiates a new Project object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProjectWithDefaults() *Project {
	this := Project{}
	var external bool = false
	this.External = &external
	return &this
}

// GetSynonyms returns the Synonyms field value if set, zero value otherwise.
func (o *Project) GetSynonyms() []string {
	if o == nil || IsNil(o.Synonyms) {
		var ret []string
		return ret
	}
	return o.Synonyms
}

// GetSynonymsOk returns a tuple with the Synonyms field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Project) GetSynonymsOk() ([]string, bool) {
	if o == nil || IsNil(o.Synonyms) {
		return nil, false
	}
	return o.Synonyms, true
}

// HasSynonyms returns a boolean if a field has been set.
func (o *Project) HasSynonyms() bool {
	if o != nil && !IsNil(o.Synonyms) {
		return true
	}

	return false
}

// SetSynonyms gets a reference to the given []string and assigns it to the Synonyms field.
func (o *Project) SetSynonyms(v []string) {
	o.Synonyms = v
}

// GetLocation returns the Location field value if set, zero value otherwise.
func (o *Project) GetLocation() string {
	if o == nil || IsNil(o.Location) {
		var ret string
		return ret
	}
	return *o.Location
}

// GetLocationOk returns a tuple with the Location field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Project) GetLocationOk() (*string, bool) {
	if o == nil || IsNil(o.Location) {
		return nil, false
	}
	return o.Location, true
}

// HasLocation returns a boolean if a field has been set.
func (o *Project) HasLocation() bool {
	if o != nil && !IsNil(o.Location) {
		return true
	}

	return false
}

// SetLocation gets a reference to the given string and assigns it to the Location field.
func (o *Project) SetLocation(v string) {
	o.Location = &v
}

// GetGeolocation returns the Geolocation field value if set, zero value otherwise.
func (o *Project) GetGeolocation() Geolocation {
	if o == nil || IsNil(o.Geolocation) {
		var ret Geolocation
		return ret
	}
	return *o.Geolocation
}

// GetGeolocationOk returns a tuple with the Geolocation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Project) GetGeolocationOk() (*Geolocation, bool) {
	if o == nil || IsNil(o.Geolocation) {
		return nil, false
	}
	return o.Geolocation, true
}

// HasGeolocation returns a boolean if a field has been set.
func (o *Project) HasGeolocation() bool {
	if o != nil && !IsNil(o.Geolocation) {
		return true
	}

	return false
}

// SetGeolocation gets a reference to the given Geolocation and assigns it to the Geolocation field.
func (o *Project) SetGeolocation(v Geolocation) {
	o.Geolocation = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *Project) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Project) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *Project) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *Project) SetDescription(v string) {
	o.Description = &v
}

// GetExternal returns the External field value if set, zero value otherwise.
func (o *Project) GetExternal() bool {
	if o == nil || IsNil(o.External) {
		var ret bool
		return ret
	}
	return *o.External
}

// GetExternalOk returns a tuple with the External field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Project) GetExternalOk() (*bool, bool) {
	if o == nil || IsNil(o.External) {
		return nil, false
	}
	return o.External, true
}

// HasExternal returns a boolean if a field has been set.
func (o *Project) HasExternal() bool {
	if o != nil && !IsNil(o.External) {
		return true
	}

	return false
}

// SetExternal gets a reference to the given bool and assigns it to the External field.
func (o *Project) SetExternal(v bool) {
	o.External = &v
}

// GetProjectType returns the ProjectType field value if set, zero value otherwise.
func (o *Project) GetProjectType() ProjectType {
	if o == nil || IsNil(o.ProjectType) {
		var ret ProjectType
		return ret
	}
	return *o.ProjectType
}

// GetProjectTypeOk returns a tuple with the ProjectType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Project) GetProjectTypeOk() (*ProjectType, bool) {
	if o == nil || IsNil(o.ProjectType) {
		return nil, false
	}
	return o.ProjectType, true
}

// HasProjectType returns a boolean if a field has been set.
func (o *Project) HasProjectType() bool {
	if o != nil && !IsNil(o.ProjectType) {
		return true
	}

	return false
}

// SetProjectType gets a reference to the given ProjectType and assigns it to the ProjectType field.
func (o *Project) SetProjectType(v ProjectType) {
	o.ProjectType = &v
}

// GetConfidentiality returns the Confidentiality field value if set, zero value otherwise.
func (o *Project) GetConfidentiality() Confidentiality {
	if o == nil || IsNil(o.Confidentiality) {
		var ret Confidentiality
		return ret
	}
	return *o.Confidentiality
}

// GetConfidentialityOk returns a tuple with the Confidentiality field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Project) GetConfidentialityOk() (*Confidentiality, bool) {
	if o == nil || IsNil(o.Confidentiality) {
		return nil, false
	}
	return o.Confidentiality, true
}

// HasConfidentiality returns a boolean if a field has been set.
func (o *Project) HasConfidentiality() bool {
	if o != nil && !IsNil(o.Confidentiality) {
		return true
	}

	return false
}

// SetConfidentiality gets a reference to the given Confidentiality and assigns it to the Confidentiality field.
func (o *Project) SetConfidentiality(v Confidentiality) {
	o.Confidentiality = &v
}

func (o Project) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Project) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	serializedNamedDomainModel, errNamedDomainModel := json.Marshal(o.NamedDomainModel)
	if errNamedDomainModel != nil {
		return map[string]interface{}{}, errNamedDomainModel
	}
	errNamedDomainModel = json.Unmarshal([]byte(serializedNamedDomainModel), &toSerialize)
	if errNamedDomainModel != nil {
		return map[string]interface{}{}, errNamedDomainModel
	}
	if !IsNil(o.Synonyms) {
		toSerialize["synonyms"] = o.Synonyms
	}
	if !IsNil(o.Location) {
		toSerialize["location"] = o.Location
	}
	if !IsNil(o.Geolocation) {
		toSerialize["geolocation"] = o.Geolocation
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.External) {
		toSerialize["external"] = o.External
	}
	if !IsNil(o.ProjectType) {
		toSerialize["projectType"] = o.ProjectType
	}
	if !IsNil(o.Confidentiality) {
		toSerialize["confidentiality"] = o.Confidentiality
	}
	return toSerialize, nil
}

type NullableProject struct {
	value *Project
	isSet bool
}

func (v NullableProject) Get() *Project {
	return v.value
}

func (v *NullableProject) Set(val *Project) {
	v.value = val
	v.isSet = true
}

func (v NullableProject) IsSet() bool {
	return v.isSet
}

func (v *NullableProject) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProject(val *Project) *NullableProject {
	return &NullableProject{value: val, isSet: true}
}

func (v NullableProject) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProject) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


