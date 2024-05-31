/*
YASM (Yet Another Skill Management) API

This is an example of using OAuth2 Implicit Flow in a specification to describe security to your API.

API version: 1.27.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// checks if the Availability type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Availability{}

// Availability struct for Availability
type Availability struct {
	NamedDomainModel
	Startdate string `json:"startdate"`
	Enddate *string `json:"enddate,omitempty"`
	ObjectType *string `json:"objectType,omitempty"`
	WorkHours float32 `json:"workHours"`
	PlannedHours float32 `json:"plannedHours"`
	Descriptions []string `json:"descriptions,omitempty"`
}

// NewAvailability instantiates a new Availability object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAvailability(startdate string, workHours float32, plannedHours float32, id string, name string) *Availability {
	this := Availability{}
	this.Id = id
	this.Name = name
	this.Startdate = startdate
	this.WorkHours = workHours
	this.PlannedHours = plannedHours
	return &this
}

// NewAvailabilityWithDefaults instantiates a new Availability object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAvailabilityWithDefaults() *Availability {
	this := Availability{}
	return &this
}

// GetStartdate returns the Startdate field value
func (o *Availability) GetStartdate() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Startdate
}

// GetStartdateOk returns a tuple with the Startdate field value
// and a boolean to check if the value has been set.
func (o *Availability) GetStartdateOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Startdate, true
}

// SetStartdate sets field value
func (o *Availability) SetStartdate(v string) {
	o.Startdate = v
}

// GetEnddate returns the Enddate field value if set, zero value otherwise.
func (o *Availability) GetEnddate() string {
	if o == nil || IsNil(o.Enddate) {
		var ret string
		return ret
	}
	return *o.Enddate
}

// GetEnddateOk returns a tuple with the Enddate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Availability) GetEnddateOk() (*string, bool) {
	if o == nil || IsNil(o.Enddate) {
		return nil, false
	}
	return o.Enddate, true
}

// HasEnddate returns a boolean if a field has been set.
func (o *Availability) HasEnddate() bool {
	if o != nil && !IsNil(o.Enddate) {
		return true
	}

	return false
}

// SetEnddate gets a reference to the given string and assigns it to the Enddate field.
func (o *Availability) SetEnddate(v string) {
	o.Enddate = &v
}

// GetObjectType returns the ObjectType field value if set, zero value otherwise.
func (o *Availability) GetObjectType() string {
	if o == nil || IsNil(o.ObjectType) {
		var ret string
		return ret
	}
	return *o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Availability) GetObjectTypeOk() (*string, bool) {
	if o == nil || IsNil(o.ObjectType) {
		return nil, false
	}
	return o.ObjectType, true
}

// HasObjectType returns a boolean if a field has been set.
func (o *Availability) HasObjectType() bool {
	if o != nil && !IsNil(o.ObjectType) {
		return true
	}

	return false
}

// SetObjectType gets a reference to the given string and assigns it to the ObjectType field.
func (o *Availability) SetObjectType(v string) {
	o.ObjectType = &v
}

// GetWorkHours returns the WorkHours field value
func (o *Availability) GetWorkHours() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.WorkHours
}

// GetWorkHoursOk returns a tuple with the WorkHours field value
// and a boolean to check if the value has been set.
func (o *Availability) GetWorkHoursOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.WorkHours, true
}

// SetWorkHours sets field value
func (o *Availability) SetWorkHours(v float32) {
	o.WorkHours = v
}

// GetPlannedHours returns the PlannedHours field value
func (o *Availability) GetPlannedHours() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.PlannedHours
}

// GetPlannedHoursOk returns a tuple with the PlannedHours field value
// and a boolean to check if the value has been set.
func (o *Availability) GetPlannedHoursOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PlannedHours, true
}

// SetPlannedHours sets field value
func (o *Availability) SetPlannedHours(v float32) {
	o.PlannedHours = v
}

// GetDescriptions returns the Descriptions field value if set, zero value otherwise.
func (o *Availability) GetDescriptions() []string {
	if o == nil || IsNil(o.Descriptions) {
		var ret []string
		return ret
	}
	return o.Descriptions
}

// GetDescriptionsOk returns a tuple with the Descriptions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Availability) GetDescriptionsOk() ([]string, bool) {
	if o == nil || IsNil(o.Descriptions) {
		return nil, false
	}
	return o.Descriptions, true
}

// HasDescriptions returns a boolean if a field has been set.
func (o *Availability) HasDescriptions() bool {
	if o != nil && !IsNil(o.Descriptions) {
		return true
	}

	return false
}

// SetDescriptions gets a reference to the given []string and assigns it to the Descriptions field.
func (o *Availability) SetDescriptions(v []string) {
	o.Descriptions = v
}

func (o Availability) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Availability) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	serializedNamedDomainModel, errNamedDomainModel := json.Marshal(o.NamedDomainModel)
	if errNamedDomainModel != nil {
		return map[string]interface{}{}, errNamedDomainModel
	}
	errNamedDomainModel = json.Unmarshal([]byte(serializedNamedDomainModel), &toSerialize)
	if errNamedDomainModel != nil {
		return map[string]interface{}{}, errNamedDomainModel
	}
	toSerialize["startdate"] = o.Startdate
	if !IsNil(o.Enddate) {
		toSerialize["enddate"] = o.Enddate
	}
	if !IsNil(o.ObjectType) {
		toSerialize["objectType"] = o.ObjectType
	}
	toSerialize["workHours"] = o.WorkHours
	toSerialize["plannedHours"] = o.PlannedHours
	if !IsNil(o.Descriptions) {
		toSerialize["descriptions"] = o.Descriptions
	}
	return toSerialize, nil
}

type NullableAvailability struct {
	value *Availability
	isSet bool
}

func (v NullableAvailability) Get() *Availability {
	return v.value
}

func (v *NullableAvailability) Set(val *Availability) {
	v.value = val
	v.isSet = true
}

func (v NullableAvailability) IsSet() bool {
	return v.isSet
}

func (v *NullableAvailability) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAvailability(val *Availability) *NullableAvailability {
	return &NullableAvailability{value: val, isSet: true}
}

func (v NullableAvailability) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAvailability) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


