/*
YASM (Yet Another Skill Management) API

This is an example of using OAuth2 Implicit Flow in a specification to describe security to your API.

API version: 1.64.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// checks if the SkillEvaluationFilter type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SkillEvaluationFilter{}

// SkillEvaluationFilter struct for SkillEvaluationFilter
type SkillEvaluationFilter struct {
	StartDate string `json:"startDate"`
	EndDate string `json:"endDate"`
	EmployeeIds []string `json:"employeeIds,omitempty"`
	PersonIds []string `json:"personIds,omitempty"`
	KindGiverIds []string `json:"kindGiverIds,omitempty"`
	IndustryIds []string `json:"industryIds,omitempty"`
	OrganizationIds []string `json:"organizationIds,omitempty"`
	CountryIds []string `json:"countryIds,omitempty"`
	ErpIds []string `json:"erpIds,omitempty"`
}

// NewSkillEvaluationFilter instantiates a new SkillEvaluationFilter object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSkillEvaluationFilter(startDate string, endDate string) *SkillEvaluationFilter {
	this := SkillEvaluationFilter{}
	this.StartDate = startDate
	this.EndDate = endDate
	return &this
}

// NewSkillEvaluationFilterWithDefaults instantiates a new SkillEvaluationFilter object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSkillEvaluationFilterWithDefaults() *SkillEvaluationFilter {
	this := SkillEvaluationFilter{}
	return &this
}

// GetStartDate returns the StartDate field value
func (o *SkillEvaluationFilter) GetStartDate() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.StartDate
}

// GetStartDateOk returns a tuple with the StartDate field value
// and a boolean to check if the value has been set.
func (o *SkillEvaluationFilter) GetStartDateOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.StartDate, true
}

// SetStartDate sets field value
func (o *SkillEvaluationFilter) SetStartDate(v string) {
	o.StartDate = v
}

// GetEndDate returns the EndDate field value
func (o *SkillEvaluationFilter) GetEndDate() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.EndDate
}

// GetEndDateOk returns a tuple with the EndDate field value
// and a boolean to check if the value has been set.
func (o *SkillEvaluationFilter) GetEndDateOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EndDate, true
}

// SetEndDate sets field value
func (o *SkillEvaluationFilter) SetEndDate(v string) {
	o.EndDate = v
}

// GetEmployeeIds returns the EmployeeIds field value if set, zero value otherwise.
func (o *SkillEvaluationFilter) GetEmployeeIds() []string {
	if o == nil || IsNil(o.EmployeeIds) {
		var ret []string
		return ret
	}
	return o.EmployeeIds
}

// GetEmployeeIdsOk returns a tuple with the EmployeeIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SkillEvaluationFilter) GetEmployeeIdsOk() ([]string, bool) {
	if o == nil || IsNil(o.EmployeeIds) {
		return nil, false
	}
	return o.EmployeeIds, true
}

// HasEmployeeIds returns a boolean if a field has been set.
func (o *SkillEvaluationFilter) HasEmployeeIds() bool {
	if o != nil && !IsNil(o.EmployeeIds) {
		return true
	}

	return false
}

// SetEmployeeIds gets a reference to the given []string and assigns it to the EmployeeIds field.
func (o *SkillEvaluationFilter) SetEmployeeIds(v []string) {
	o.EmployeeIds = v
}

// GetPersonIds returns the PersonIds field value if set, zero value otherwise.
func (o *SkillEvaluationFilter) GetPersonIds() []string {
	if o == nil || IsNil(o.PersonIds) {
		var ret []string
		return ret
	}
	return o.PersonIds
}

// GetPersonIdsOk returns a tuple with the PersonIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SkillEvaluationFilter) GetPersonIdsOk() ([]string, bool) {
	if o == nil || IsNil(o.PersonIds) {
		return nil, false
	}
	return o.PersonIds, true
}

// HasPersonIds returns a boolean if a field has been set.
func (o *SkillEvaluationFilter) HasPersonIds() bool {
	if o != nil && !IsNil(o.PersonIds) {
		return true
	}

	return false
}

// SetPersonIds gets a reference to the given []string and assigns it to the PersonIds field.
func (o *SkillEvaluationFilter) SetPersonIds(v []string) {
	o.PersonIds = v
}

// GetKindGiverIds returns the KindGiverIds field value if set, zero value otherwise.
func (o *SkillEvaluationFilter) GetKindGiverIds() []string {
	if o == nil || IsNil(o.KindGiverIds) {
		var ret []string
		return ret
	}
	return o.KindGiverIds
}

// GetKindGiverIdsOk returns a tuple with the KindGiverIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SkillEvaluationFilter) GetKindGiverIdsOk() ([]string, bool) {
	if o == nil || IsNil(o.KindGiverIds) {
		return nil, false
	}
	return o.KindGiverIds, true
}

// HasKindGiverIds returns a boolean if a field has been set.
func (o *SkillEvaluationFilter) HasKindGiverIds() bool {
	if o != nil && !IsNil(o.KindGiverIds) {
		return true
	}

	return false
}

// SetKindGiverIds gets a reference to the given []string and assigns it to the KindGiverIds field.
func (o *SkillEvaluationFilter) SetKindGiverIds(v []string) {
	o.KindGiverIds = v
}

// GetIndustryIds returns the IndustryIds field value if set, zero value otherwise.
func (o *SkillEvaluationFilter) GetIndustryIds() []string {
	if o == nil || IsNil(o.IndustryIds) {
		var ret []string
		return ret
	}
	return o.IndustryIds
}

// GetIndustryIdsOk returns a tuple with the IndustryIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SkillEvaluationFilter) GetIndustryIdsOk() ([]string, bool) {
	if o == nil || IsNil(o.IndustryIds) {
		return nil, false
	}
	return o.IndustryIds, true
}

// HasIndustryIds returns a boolean if a field has been set.
func (o *SkillEvaluationFilter) HasIndustryIds() bool {
	if o != nil && !IsNil(o.IndustryIds) {
		return true
	}

	return false
}

// SetIndustryIds gets a reference to the given []string and assigns it to the IndustryIds field.
func (o *SkillEvaluationFilter) SetIndustryIds(v []string) {
	o.IndustryIds = v
}

// GetOrganizationIds returns the OrganizationIds field value if set, zero value otherwise.
func (o *SkillEvaluationFilter) GetOrganizationIds() []string {
	if o == nil || IsNil(o.OrganizationIds) {
		var ret []string
		return ret
	}
	return o.OrganizationIds
}

// GetOrganizationIdsOk returns a tuple with the OrganizationIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SkillEvaluationFilter) GetOrganizationIdsOk() ([]string, bool) {
	if o == nil || IsNil(o.OrganizationIds) {
		return nil, false
	}
	return o.OrganizationIds, true
}

// HasOrganizationIds returns a boolean if a field has been set.
func (o *SkillEvaluationFilter) HasOrganizationIds() bool {
	if o != nil && !IsNil(o.OrganizationIds) {
		return true
	}

	return false
}

// SetOrganizationIds gets a reference to the given []string and assigns it to the OrganizationIds field.
func (o *SkillEvaluationFilter) SetOrganizationIds(v []string) {
	o.OrganizationIds = v
}

// GetCountryIds returns the CountryIds field value if set, zero value otherwise.
func (o *SkillEvaluationFilter) GetCountryIds() []string {
	if o == nil || IsNil(o.CountryIds) {
		var ret []string
		return ret
	}
	return o.CountryIds
}

// GetCountryIdsOk returns a tuple with the CountryIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SkillEvaluationFilter) GetCountryIdsOk() ([]string, bool) {
	if o == nil || IsNil(o.CountryIds) {
		return nil, false
	}
	return o.CountryIds, true
}

// HasCountryIds returns a boolean if a field has been set.
func (o *SkillEvaluationFilter) HasCountryIds() bool {
	if o != nil && !IsNil(o.CountryIds) {
		return true
	}

	return false
}

// SetCountryIds gets a reference to the given []string and assigns it to the CountryIds field.
func (o *SkillEvaluationFilter) SetCountryIds(v []string) {
	o.CountryIds = v
}

// GetErpIds returns the ErpIds field value if set, zero value otherwise.
func (o *SkillEvaluationFilter) GetErpIds() []string {
	if o == nil || IsNil(o.ErpIds) {
		var ret []string
		return ret
	}
	return o.ErpIds
}

// GetErpIdsOk returns a tuple with the ErpIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SkillEvaluationFilter) GetErpIdsOk() ([]string, bool) {
	if o == nil || IsNil(o.ErpIds) {
		return nil, false
	}
	return o.ErpIds, true
}

// HasErpIds returns a boolean if a field has been set.
func (o *SkillEvaluationFilter) HasErpIds() bool {
	if o != nil && !IsNil(o.ErpIds) {
		return true
	}

	return false
}

// SetErpIds gets a reference to the given []string and assigns it to the ErpIds field.
func (o *SkillEvaluationFilter) SetErpIds(v []string) {
	o.ErpIds = v
}

func (o SkillEvaluationFilter) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SkillEvaluationFilter) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["startDate"] = o.StartDate
	toSerialize["endDate"] = o.EndDate
	if !IsNil(o.EmployeeIds) {
		toSerialize["employeeIds"] = o.EmployeeIds
	}
	if !IsNil(o.PersonIds) {
		toSerialize["personIds"] = o.PersonIds
	}
	if !IsNil(o.KindGiverIds) {
		toSerialize["kindGiverIds"] = o.KindGiverIds
	}
	if !IsNil(o.IndustryIds) {
		toSerialize["industryIds"] = o.IndustryIds
	}
	if !IsNil(o.OrganizationIds) {
		toSerialize["organizationIds"] = o.OrganizationIds
	}
	if !IsNil(o.CountryIds) {
		toSerialize["countryIds"] = o.CountryIds
	}
	if !IsNil(o.ErpIds) {
		toSerialize["erpIds"] = o.ErpIds
	}
	return toSerialize, nil
}

type NullableSkillEvaluationFilter struct {
	value *SkillEvaluationFilter
	isSet bool
}

func (v NullableSkillEvaluationFilter) Get() *SkillEvaluationFilter {
	return v.value
}

func (v *NullableSkillEvaluationFilter) Set(val *SkillEvaluationFilter) {
	v.value = val
	v.isSet = true
}

func (v NullableSkillEvaluationFilter) IsSet() bool {
	return v.isSet
}

func (v *NullableSkillEvaluationFilter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSkillEvaluationFilter(val *SkillEvaluationFilter) *NullableSkillEvaluationFilter {
	return &NullableSkillEvaluationFilter{value: val, isSet: true}
}

func (v NullableSkillEvaluationFilter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSkillEvaluationFilter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


