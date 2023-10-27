/*
YASM (Yet Another Skill Management) API

This is an example of using OAuth2 Implicit Flow in a specification to describe security to your API.

API version: 1.5.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// checks if the CertificationSearch type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CertificationSearch{}

// CertificationSearch struct for CertificationSearch
type CertificationSearch struct {
	CertificationIds []string `json:"certificationIds,omitempty"`
	SkillIds []string `json:"skillIds,omitempty"`
	OrganizationIds []string `json:"organizationIds,omitempty"`
}

// NewCertificationSearch instantiates a new CertificationSearch object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCertificationSearch() *CertificationSearch {
	this := CertificationSearch{}
	return &this
}

// NewCertificationSearchWithDefaults instantiates a new CertificationSearch object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCertificationSearchWithDefaults() *CertificationSearch {
	this := CertificationSearch{}
	return &this
}

// GetCertificationIds returns the CertificationIds field value if set, zero value otherwise.
func (o *CertificationSearch) GetCertificationIds() []string {
	if o == nil || IsNil(o.CertificationIds) {
		var ret []string
		return ret
	}
	return o.CertificationIds
}

// GetCertificationIdsOk returns a tuple with the CertificationIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CertificationSearch) GetCertificationIdsOk() ([]string, bool) {
	if o == nil || IsNil(o.CertificationIds) {
		return nil, false
	}
	return o.CertificationIds, true
}

// HasCertificationIds returns a boolean if a field has been set.
func (o *CertificationSearch) HasCertificationIds() bool {
	if o != nil && !IsNil(o.CertificationIds) {
		return true
	}

	return false
}

// SetCertificationIds gets a reference to the given []string and assigns it to the CertificationIds field.
func (o *CertificationSearch) SetCertificationIds(v []string) {
	o.CertificationIds = v
}

// GetSkillIds returns the SkillIds field value if set, zero value otherwise.
func (o *CertificationSearch) GetSkillIds() []string {
	if o == nil || IsNil(o.SkillIds) {
		var ret []string
		return ret
	}
	return o.SkillIds
}

// GetSkillIdsOk returns a tuple with the SkillIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CertificationSearch) GetSkillIdsOk() ([]string, bool) {
	if o == nil || IsNil(o.SkillIds) {
		return nil, false
	}
	return o.SkillIds, true
}

// HasSkillIds returns a boolean if a field has been set.
func (o *CertificationSearch) HasSkillIds() bool {
	if o != nil && !IsNil(o.SkillIds) {
		return true
	}

	return false
}

// SetSkillIds gets a reference to the given []string and assigns it to the SkillIds field.
func (o *CertificationSearch) SetSkillIds(v []string) {
	o.SkillIds = v
}

// GetOrganizationIds returns the OrganizationIds field value if set, zero value otherwise.
func (o *CertificationSearch) GetOrganizationIds() []string {
	if o == nil || IsNil(o.OrganizationIds) {
		var ret []string
		return ret
	}
	return o.OrganizationIds
}

// GetOrganizationIdsOk returns a tuple with the OrganizationIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CertificationSearch) GetOrganizationIdsOk() ([]string, bool) {
	if o == nil || IsNil(o.OrganizationIds) {
		return nil, false
	}
	return o.OrganizationIds, true
}

// HasOrganizationIds returns a boolean if a field has been set.
func (o *CertificationSearch) HasOrganizationIds() bool {
	if o != nil && !IsNil(o.OrganizationIds) {
		return true
	}

	return false
}

// SetOrganizationIds gets a reference to the given []string and assigns it to the OrganizationIds field.
func (o *CertificationSearch) SetOrganizationIds(v []string) {
	o.OrganizationIds = v
}

func (o CertificationSearch) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CertificationSearch) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CertificationIds) {
		toSerialize["certificationIds"] = o.CertificationIds
	}
	if !IsNil(o.SkillIds) {
		toSerialize["skillIds"] = o.SkillIds
	}
	if !IsNil(o.OrganizationIds) {
		toSerialize["organizationIds"] = o.OrganizationIds
	}
	return toSerialize, nil
}

type NullableCertificationSearch struct {
	value *CertificationSearch
	isSet bool
}

func (v NullableCertificationSearch) Get() *CertificationSearch {
	return v.value
}

func (v *NullableCertificationSearch) Set(val *CertificationSearch) {
	v.value = val
	v.isSet = true
}

func (v NullableCertificationSearch) IsSet() bool {
	return v.isSet
}

func (v *NullableCertificationSearch) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCertificationSearch(val *CertificationSearch) *NullableCertificationSearch {
	return &NullableCertificationSearch{value: val, isSet: true}
}

func (v NullableCertificationSearch) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCertificationSearch) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


