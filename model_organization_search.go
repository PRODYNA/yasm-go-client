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

// checks if the OrganizationSearch type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OrganizationSearch{}

// OrganizationSearch struct for OrganizationSearch
type OrganizationSearch struct {
	Search
	PersonIds []string `json:"personIds,omitempty"`
	ProjectIds []string `json:"projectIds,omitempty"`
	OrganizationIds []string `json:"organizationIds,omitempty"`
	IndustryIds []string `json:"industryIds,omitempty"`
	CertificationIds []string `json:"certificationIds,omitempty"`
}

// NewOrganizationSearch instantiates a new OrganizationSearch object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrganizationSearch(skip int32, limit int32) *OrganizationSearch {
	this := OrganizationSearch{}
	this.Skip = skip
	this.Limit = limit
	return &this
}

// NewOrganizationSearchWithDefaults instantiates a new OrganizationSearch object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrganizationSearchWithDefaults() *OrganizationSearch {
	this := OrganizationSearch{}
	return &this
}

// GetPersonIds returns the PersonIds field value if set, zero value otherwise.
func (o *OrganizationSearch) GetPersonIds() []string {
	if o == nil || IsNil(o.PersonIds) {
		var ret []string
		return ret
	}
	return o.PersonIds
}

// GetPersonIdsOk returns a tuple with the PersonIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationSearch) GetPersonIdsOk() ([]string, bool) {
	if o == nil || IsNil(o.PersonIds) {
		return nil, false
	}
	return o.PersonIds, true
}

// HasPersonIds returns a boolean if a field has been set.
func (o *OrganizationSearch) HasPersonIds() bool {
	if o != nil && !IsNil(o.PersonIds) {
		return true
	}

	return false
}

// SetPersonIds gets a reference to the given []string and assigns it to the PersonIds field.
func (o *OrganizationSearch) SetPersonIds(v []string) {
	o.PersonIds = v
}

// GetProjectIds returns the ProjectIds field value if set, zero value otherwise.
func (o *OrganizationSearch) GetProjectIds() []string {
	if o == nil || IsNil(o.ProjectIds) {
		var ret []string
		return ret
	}
	return o.ProjectIds
}

// GetProjectIdsOk returns a tuple with the ProjectIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationSearch) GetProjectIdsOk() ([]string, bool) {
	if o == nil || IsNil(o.ProjectIds) {
		return nil, false
	}
	return o.ProjectIds, true
}

// HasProjectIds returns a boolean if a field has been set.
func (o *OrganizationSearch) HasProjectIds() bool {
	if o != nil && !IsNil(o.ProjectIds) {
		return true
	}

	return false
}

// SetProjectIds gets a reference to the given []string and assigns it to the ProjectIds field.
func (o *OrganizationSearch) SetProjectIds(v []string) {
	o.ProjectIds = v
}

// GetOrganizationIds returns the OrganizationIds field value if set, zero value otherwise.
func (o *OrganizationSearch) GetOrganizationIds() []string {
	if o == nil || IsNil(o.OrganizationIds) {
		var ret []string
		return ret
	}
	return o.OrganizationIds
}

// GetOrganizationIdsOk returns a tuple with the OrganizationIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationSearch) GetOrganizationIdsOk() ([]string, bool) {
	if o == nil || IsNil(o.OrganizationIds) {
		return nil, false
	}
	return o.OrganizationIds, true
}

// HasOrganizationIds returns a boolean if a field has been set.
func (o *OrganizationSearch) HasOrganizationIds() bool {
	if o != nil && !IsNil(o.OrganizationIds) {
		return true
	}

	return false
}

// SetOrganizationIds gets a reference to the given []string and assigns it to the OrganizationIds field.
func (o *OrganizationSearch) SetOrganizationIds(v []string) {
	o.OrganizationIds = v
}

// GetIndustryIds returns the IndustryIds field value if set, zero value otherwise.
func (o *OrganizationSearch) GetIndustryIds() []string {
	if o == nil || IsNil(o.IndustryIds) {
		var ret []string
		return ret
	}
	return o.IndustryIds
}

// GetIndustryIdsOk returns a tuple with the IndustryIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationSearch) GetIndustryIdsOk() ([]string, bool) {
	if o == nil || IsNil(o.IndustryIds) {
		return nil, false
	}
	return o.IndustryIds, true
}

// HasIndustryIds returns a boolean if a field has been set.
func (o *OrganizationSearch) HasIndustryIds() bool {
	if o != nil && !IsNil(o.IndustryIds) {
		return true
	}

	return false
}

// SetIndustryIds gets a reference to the given []string and assigns it to the IndustryIds field.
func (o *OrganizationSearch) SetIndustryIds(v []string) {
	o.IndustryIds = v
}

// GetCertificationIds returns the CertificationIds field value if set, zero value otherwise.
func (o *OrganizationSearch) GetCertificationIds() []string {
	if o == nil || IsNil(o.CertificationIds) {
		var ret []string
		return ret
	}
	return o.CertificationIds
}

// GetCertificationIdsOk returns a tuple with the CertificationIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationSearch) GetCertificationIdsOk() ([]string, bool) {
	if o == nil || IsNil(o.CertificationIds) {
		return nil, false
	}
	return o.CertificationIds, true
}

// HasCertificationIds returns a boolean if a field has been set.
func (o *OrganizationSearch) HasCertificationIds() bool {
	if o != nil && !IsNil(o.CertificationIds) {
		return true
	}

	return false
}

// SetCertificationIds gets a reference to the given []string and assigns it to the CertificationIds field.
func (o *OrganizationSearch) SetCertificationIds(v []string) {
	o.CertificationIds = v
}

func (o OrganizationSearch) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OrganizationSearch) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	serializedSearch, errSearch := json.Marshal(o.Search)
	if errSearch != nil {
		return map[string]interface{}{}, errSearch
	}
	errSearch = json.Unmarshal([]byte(serializedSearch), &toSerialize)
	if errSearch != nil {
		return map[string]interface{}{}, errSearch
	}
	if !IsNil(o.PersonIds) {
		toSerialize["personIds"] = o.PersonIds
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
	return toSerialize, nil
}

type NullableOrganizationSearch struct {
	value *OrganizationSearch
	isSet bool
}

func (v NullableOrganizationSearch) Get() *OrganizationSearch {
	return v.value
}

func (v *NullableOrganizationSearch) Set(val *OrganizationSearch) {
	v.value = val
	v.isSet = true
}

func (v NullableOrganizationSearch) IsSet() bool {
	return v.isSet
}

func (v *NullableOrganizationSearch) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrganizationSearch(val *OrganizationSearch) *NullableOrganizationSearch {
	return &NullableOrganizationSearch{value: val, isSet: true}
}

func (v NullableOrganizationSearch) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrganizationSearch) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


