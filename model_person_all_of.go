/*
YASM (Yet Another Skill Management) API

This is an example of using OAuth2 Implicit Flow in a specification to describe security to your API.

API version: 0.9.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// PersonAllOf struct for PersonAllOf
type PersonAllOf struct {
	EmployeeId *string `json:"employeeId,omitempty"`
	JobTitle *string `json:"jobTitle,omitempty"`
	Company *string `json:"company,omitempty"`
	Department *string `json:"department,omitempty"`
	Mail *string `json:"mail,omitempty"`
	MobilePhone *string `json:"mobilePhone,omitempty"`
	// base64 encoded image
	Picture *string `json:"picture,omitempty"`
	// Marks persons not working for the company anymore
	Inactive *bool `json:"inactive,omitempty"`
}

// NewPersonAllOf instantiates a new PersonAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPersonAllOf() *PersonAllOf {
	this := PersonAllOf{}
	var inactive bool = false
	this.Inactive = &inactive
	return &this
}

// NewPersonAllOfWithDefaults instantiates a new PersonAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPersonAllOfWithDefaults() *PersonAllOf {
	this := PersonAllOf{}
	var inactive bool = false
	this.Inactive = &inactive
	return &this
}

// GetEmployeeId returns the EmployeeId field value if set, zero value otherwise.
func (o *PersonAllOf) GetEmployeeId() string {
	if o == nil || o.EmployeeId == nil {
		var ret string
		return ret
	}
	return *o.EmployeeId
}

// GetEmployeeIdOk returns a tuple with the EmployeeId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PersonAllOf) GetEmployeeIdOk() (*string, bool) {
	if o == nil || o.EmployeeId == nil {
		return nil, false
	}
	return o.EmployeeId, true
}

// HasEmployeeId returns a boolean if a field has been set.
func (o *PersonAllOf) HasEmployeeId() bool {
	if o != nil && o.EmployeeId != nil {
		return true
	}

	return false
}

// SetEmployeeId gets a reference to the given string and assigns it to the EmployeeId field.
func (o *PersonAllOf) SetEmployeeId(v string) {
	o.EmployeeId = &v
}

// GetJobTitle returns the JobTitle field value if set, zero value otherwise.
func (o *PersonAllOf) GetJobTitle() string {
	if o == nil || o.JobTitle == nil {
		var ret string
		return ret
	}
	return *o.JobTitle
}

// GetJobTitleOk returns a tuple with the JobTitle field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PersonAllOf) GetJobTitleOk() (*string, bool) {
	if o == nil || o.JobTitle == nil {
		return nil, false
	}
	return o.JobTitle, true
}

// HasJobTitle returns a boolean if a field has been set.
func (o *PersonAllOf) HasJobTitle() bool {
	if o != nil && o.JobTitle != nil {
		return true
	}

	return false
}

// SetJobTitle gets a reference to the given string and assigns it to the JobTitle field.
func (o *PersonAllOf) SetJobTitle(v string) {
	o.JobTitle = &v
}

// GetCompany returns the Company field value if set, zero value otherwise.
func (o *PersonAllOf) GetCompany() string {
	if o == nil || o.Company == nil {
		var ret string
		return ret
	}
	return *o.Company
}

// GetCompanyOk returns a tuple with the Company field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PersonAllOf) GetCompanyOk() (*string, bool) {
	if o == nil || o.Company == nil {
		return nil, false
	}
	return o.Company, true
}

// HasCompany returns a boolean if a field has been set.
func (o *PersonAllOf) HasCompany() bool {
	if o != nil && o.Company != nil {
		return true
	}

	return false
}

// SetCompany gets a reference to the given string and assigns it to the Company field.
func (o *PersonAllOf) SetCompany(v string) {
	o.Company = &v
}

// GetDepartment returns the Department field value if set, zero value otherwise.
func (o *PersonAllOf) GetDepartment() string {
	if o == nil || o.Department == nil {
		var ret string
		return ret
	}
	return *o.Department
}

// GetDepartmentOk returns a tuple with the Department field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PersonAllOf) GetDepartmentOk() (*string, bool) {
	if o == nil || o.Department == nil {
		return nil, false
	}
	return o.Department, true
}

// HasDepartment returns a boolean if a field has been set.
func (o *PersonAllOf) HasDepartment() bool {
	if o != nil && o.Department != nil {
		return true
	}

	return false
}

// SetDepartment gets a reference to the given string and assigns it to the Department field.
func (o *PersonAllOf) SetDepartment(v string) {
	o.Department = &v
}

// GetMail returns the Mail field value if set, zero value otherwise.
func (o *PersonAllOf) GetMail() string {
	if o == nil || o.Mail == nil {
		var ret string
		return ret
	}
	return *o.Mail
}

// GetMailOk returns a tuple with the Mail field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PersonAllOf) GetMailOk() (*string, bool) {
	if o == nil || o.Mail == nil {
		return nil, false
	}
	return o.Mail, true
}

// HasMail returns a boolean if a field has been set.
func (o *PersonAllOf) HasMail() bool {
	if o != nil && o.Mail != nil {
		return true
	}

	return false
}

// SetMail gets a reference to the given string and assigns it to the Mail field.
func (o *PersonAllOf) SetMail(v string) {
	o.Mail = &v
}

// GetMobilePhone returns the MobilePhone field value if set, zero value otherwise.
func (o *PersonAllOf) GetMobilePhone() string {
	if o == nil || o.MobilePhone == nil {
		var ret string
		return ret
	}
	return *o.MobilePhone
}

// GetMobilePhoneOk returns a tuple with the MobilePhone field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PersonAllOf) GetMobilePhoneOk() (*string, bool) {
	if o == nil || o.MobilePhone == nil {
		return nil, false
	}
	return o.MobilePhone, true
}

// HasMobilePhone returns a boolean if a field has been set.
func (o *PersonAllOf) HasMobilePhone() bool {
	if o != nil && o.MobilePhone != nil {
		return true
	}

	return false
}

// SetMobilePhone gets a reference to the given string and assigns it to the MobilePhone field.
func (o *PersonAllOf) SetMobilePhone(v string) {
	o.MobilePhone = &v
}

// GetPicture returns the Picture field value if set, zero value otherwise.
func (o *PersonAllOf) GetPicture() string {
	if o == nil || o.Picture == nil {
		var ret string
		return ret
	}
	return *o.Picture
}

// GetPictureOk returns a tuple with the Picture field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PersonAllOf) GetPictureOk() (*string, bool) {
	if o == nil || o.Picture == nil {
		return nil, false
	}
	return o.Picture, true
}

// HasPicture returns a boolean if a field has been set.
func (o *PersonAllOf) HasPicture() bool {
	if o != nil && o.Picture != nil {
		return true
	}

	return false
}

// SetPicture gets a reference to the given string and assigns it to the Picture field.
func (o *PersonAllOf) SetPicture(v string) {
	o.Picture = &v
}

// GetInactive returns the Inactive field value if set, zero value otherwise.
func (o *PersonAllOf) GetInactive() bool {
	if o == nil || o.Inactive == nil {
		var ret bool
		return ret
	}
	return *o.Inactive
}

// GetInactiveOk returns a tuple with the Inactive field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PersonAllOf) GetInactiveOk() (*bool, bool) {
	if o == nil || o.Inactive == nil {
		return nil, false
	}
	return o.Inactive, true
}

// HasInactive returns a boolean if a field has been set.
func (o *PersonAllOf) HasInactive() bool {
	if o != nil && o.Inactive != nil {
		return true
	}

	return false
}

// SetInactive gets a reference to the given bool and assigns it to the Inactive field.
func (o *PersonAllOf) SetInactive(v bool) {
	o.Inactive = &v
}

func (o PersonAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.EmployeeId != nil {
		toSerialize["employeeId"] = o.EmployeeId
	}
	if o.JobTitle != nil {
		toSerialize["jobTitle"] = o.JobTitle
	}
	if o.Company != nil {
		toSerialize["company"] = o.Company
	}
	if o.Department != nil {
		toSerialize["department"] = o.Department
	}
	if o.Mail != nil {
		toSerialize["mail"] = o.Mail
	}
	if o.MobilePhone != nil {
		toSerialize["mobilePhone"] = o.MobilePhone
	}
	if o.Picture != nil {
		toSerialize["picture"] = o.Picture
	}
	if o.Inactive != nil {
		toSerialize["inactive"] = o.Inactive
	}
	return json.Marshal(toSerialize)
}

type NullablePersonAllOf struct {
	value *PersonAllOf
	isSet bool
}

func (v NullablePersonAllOf) Get() *PersonAllOf {
	return v.value
}

func (v *NullablePersonAllOf) Set(val *PersonAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullablePersonAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullablePersonAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePersonAllOf(val *PersonAllOf) *NullablePersonAllOf {
	return &NullablePersonAllOf{value: val, isSet: true}
}

func (v NullablePersonAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePersonAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


