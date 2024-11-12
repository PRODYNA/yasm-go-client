/*
YASM (Yet Another Skill Management) API

This is an example of using OAuth2 Implicit Flow in a specification to describe security to your API.

API version: 1.46.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// checks if the SkillSearch type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SkillSearch{}

// SkillSearch struct for SkillSearch
type SkillSearch struct {
	Search
	// Gives you either all skills, only the root kills
	Types *string `json:"types,omitempty"`
	SkillIds []string `json:"skillIds,omitempty"`
	ProjectIds []string `json:"projectIds,omitempty"`
	OrganizationIds []string `json:"organizationIds,omitempty"`
	IndustryIds []string `json:"industryIds,omitempty"`
	CertificationIds []string `json:"certificationIds,omitempty"`
	OfficeIds []string `json:"officeIds,omitempty"`
}

// NewSkillSearch instantiates a new SkillSearch object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSkillSearch(skip int32, limit int32) *SkillSearch {
	this := SkillSearch{}
	this.Skip = skip
	this.Limit = limit
	var types string = "all"
	this.Types = &types
	return &this
}

// NewSkillSearchWithDefaults instantiates a new SkillSearch object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSkillSearchWithDefaults() *SkillSearch {
	this := SkillSearch{}
	var types string = "all"
	this.Types = &types
	return &this
}

// GetTypes returns the Types field value if set, zero value otherwise.
func (o *SkillSearch) GetTypes() string {
	if o == nil || IsNil(o.Types) {
		var ret string
		return ret
	}
	return *o.Types
}

// GetTypesOk returns a tuple with the Types field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SkillSearch) GetTypesOk() (*string, bool) {
	if o == nil || IsNil(o.Types) {
		return nil, false
	}
	return o.Types, true
}

// HasTypes returns a boolean if a field has been set.
func (o *SkillSearch) HasTypes() bool {
	if o != nil && !IsNil(o.Types) {
		return true
	}

	return false
}

// SetTypes gets a reference to the given string and assigns it to the Types field.
func (o *SkillSearch) SetTypes(v string) {
	o.Types = &v
}

// GetSkillIds returns the SkillIds field value if set, zero value otherwise.
func (o *SkillSearch) GetSkillIds() []string {
	if o == nil || IsNil(o.SkillIds) {
		var ret []string
		return ret
	}
	return o.SkillIds
}

// GetSkillIdsOk returns a tuple with the SkillIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SkillSearch) GetSkillIdsOk() ([]string, bool) {
	if o == nil || IsNil(o.SkillIds) {
		return nil, false
	}
	return o.SkillIds, true
}

// HasSkillIds returns a boolean if a field has been set.
func (o *SkillSearch) HasSkillIds() bool {
	if o != nil && !IsNil(o.SkillIds) {
		return true
	}

	return false
}

// SetSkillIds gets a reference to the given []string and assigns it to the SkillIds field.
func (o *SkillSearch) SetSkillIds(v []string) {
	o.SkillIds = v
}

// GetProjectIds returns the ProjectIds field value if set, zero value otherwise.
func (o *SkillSearch) GetProjectIds() []string {
	if o == nil || IsNil(o.ProjectIds) {
		var ret []string
		return ret
	}
	return o.ProjectIds
}

// GetProjectIdsOk returns a tuple with the ProjectIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SkillSearch) GetProjectIdsOk() ([]string, bool) {
	if o == nil || IsNil(o.ProjectIds) {
		return nil, false
	}
	return o.ProjectIds, true
}

// HasProjectIds returns a boolean if a field has been set.
func (o *SkillSearch) HasProjectIds() bool {
	if o != nil && !IsNil(o.ProjectIds) {
		return true
	}

	return false
}

// SetProjectIds gets a reference to the given []string and assigns it to the ProjectIds field.
func (o *SkillSearch) SetProjectIds(v []string) {
	o.ProjectIds = v
}

// GetOrganizationIds returns the OrganizationIds field value if set, zero value otherwise.
func (o *SkillSearch) GetOrganizationIds() []string {
	if o == nil || IsNil(o.OrganizationIds) {
		var ret []string
		return ret
	}
	return o.OrganizationIds
}

// GetOrganizationIdsOk returns a tuple with the OrganizationIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SkillSearch) GetOrganizationIdsOk() ([]string, bool) {
	if o == nil || IsNil(o.OrganizationIds) {
		return nil, false
	}
	return o.OrganizationIds, true
}

// HasOrganizationIds returns a boolean if a field has been set.
func (o *SkillSearch) HasOrganizationIds() bool {
	if o != nil && !IsNil(o.OrganizationIds) {
		return true
	}

	return false
}

// SetOrganizationIds gets a reference to the given []string and assigns it to the OrganizationIds field.
func (o *SkillSearch) SetOrganizationIds(v []string) {
	o.OrganizationIds = v
}

// GetIndustryIds returns the IndustryIds field value if set, zero value otherwise.
func (o *SkillSearch) GetIndustryIds() []string {
	if o == nil || IsNil(o.IndustryIds) {
		var ret []string
		return ret
	}
	return o.IndustryIds
}

// GetIndustryIdsOk returns a tuple with the IndustryIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SkillSearch) GetIndustryIdsOk() ([]string, bool) {
	if o == nil || IsNil(o.IndustryIds) {
		return nil, false
	}
	return o.IndustryIds, true
}

// HasIndustryIds returns a boolean if a field has been set.
func (o *SkillSearch) HasIndustryIds() bool {
	if o != nil && !IsNil(o.IndustryIds) {
		return true
	}

	return false
}

// SetIndustryIds gets a reference to the given []string and assigns it to the IndustryIds field.
func (o *SkillSearch) SetIndustryIds(v []string) {
	o.IndustryIds = v
}

// GetCertificationIds returns the CertificationIds field value if set, zero value otherwise.
func (o *SkillSearch) GetCertificationIds() []string {
	if o == nil || IsNil(o.CertificationIds) {
		var ret []string
		return ret
	}
	return o.CertificationIds
}

// GetCertificationIdsOk returns a tuple with the CertificationIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SkillSearch) GetCertificationIdsOk() ([]string, bool) {
	if o == nil || IsNil(o.CertificationIds) {
		return nil, false
	}
	return o.CertificationIds, true
}

// HasCertificationIds returns a boolean if a field has been set.
func (o *SkillSearch) HasCertificationIds() bool {
	if o != nil && !IsNil(o.CertificationIds) {
		return true
	}

	return false
}

// SetCertificationIds gets a reference to the given []string and assigns it to the CertificationIds field.
func (o *SkillSearch) SetCertificationIds(v []string) {
	o.CertificationIds = v
}

// GetOfficeIds returns the OfficeIds field value if set, zero value otherwise.
func (o *SkillSearch) GetOfficeIds() []string {
	if o == nil || IsNil(o.OfficeIds) {
		var ret []string
		return ret
	}
	return o.OfficeIds
}

// GetOfficeIdsOk returns a tuple with the OfficeIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SkillSearch) GetOfficeIdsOk() ([]string, bool) {
	if o == nil || IsNil(o.OfficeIds) {
		return nil, false
	}
	return o.OfficeIds, true
}

// HasOfficeIds returns a boolean if a field has been set.
func (o *SkillSearch) HasOfficeIds() bool {
	if o != nil && !IsNil(o.OfficeIds) {
		return true
	}

	return false
}

// SetOfficeIds gets a reference to the given []string and assigns it to the OfficeIds field.
func (o *SkillSearch) SetOfficeIds(v []string) {
	o.OfficeIds = v
}

func (o SkillSearch) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SkillSearch) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	serializedSearch, errSearch := json.Marshal(o.Search)
	if errSearch != nil {
		return map[string]interface{}{}, errSearch
	}
	errSearch = json.Unmarshal([]byte(serializedSearch), &toSerialize)
	if errSearch != nil {
		return map[string]interface{}{}, errSearch
	}
	if !IsNil(o.Types) {
		toSerialize["types"] = o.Types
	}
	if !IsNil(o.SkillIds) {
		toSerialize["skillIds"] = o.SkillIds
	}
	if !IsNil(o.ProjectIds) {
		toSerialize["projectIds"] = o.ProjectIds
	}
	if !IsNil(o.OrganizationIds) {
		toSerialize["organizationIds"] = o.OrganizationIds
	}
	if !IsNil(o.IndustryIds) {
		toSerialize["industryIds"] = o.IndustryIds
	}
	if !IsNil(o.CertificationIds) {
		toSerialize["certificationIds"] = o.CertificationIds
	}
	if !IsNil(o.OfficeIds) {
		toSerialize["officeIds"] = o.OfficeIds
	}
	return toSerialize, nil
}

type NullableSkillSearch struct {
	value *SkillSearch
	isSet bool
}

func (v NullableSkillSearch) Get() *SkillSearch {
	return v.value
}

func (v *NullableSkillSearch) Set(val *SkillSearch) {
	v.value = val
	v.isSet = true
}

func (v NullableSkillSearch) IsSet() bool {
	return v.isSet
}

func (v *NullableSkillSearch) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSkillSearch(val *SkillSearch) *NullableSkillSearch {
	return &NullableSkillSearch{value: val, isSet: true}
}

func (v NullableSkillSearch) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSkillSearch) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


