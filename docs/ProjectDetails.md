# ProjectDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Audit** | Pointer to [**Audit**](Audit.md) |  | [optional] 
**Project** | Pointer to [**Project**](Project.md) |  | [optional] 
**ExecutiveOrganizations** | Pointer to [**[]Organization**](Organization.md) |  | [optional] 
**Organization** | Pointer to [**Organization**](Organization.md) |  | [optional] 
**Industries** | Pointer to [**[]Industry**](Industry.md) |  | [optional] 
**Persons** | Pointer to [**[]Person**](Person.md) |  | [optional] 
**SkillGroups** | Pointer to [**[]ExperienceSkillGroup**](ExperienceSkillGroup.md) |  | [optional] 
**Awards** | Pointer to [**[]Award**](Award.md) |  | [optional] 
**Timeframe** | Pointer to [**Timeframed**](Timeframed.md) |  | [optional] 
**ObjectType** | Pointer to **string** |  | [optional] 

## Methods

### NewProjectDetails

`func NewProjectDetails() *ProjectDetails`

NewProjectDetails instantiates a new ProjectDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProjectDetailsWithDefaults

`func NewProjectDetailsWithDefaults() *ProjectDetails`

NewProjectDetailsWithDefaults instantiates a new ProjectDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAudit

`func (o *ProjectDetails) GetAudit() Audit`

GetAudit returns the Audit field if non-nil, zero value otherwise.

### GetAuditOk

`func (o *ProjectDetails) GetAuditOk() (*Audit, bool)`

GetAuditOk returns a tuple with the Audit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAudit

`func (o *ProjectDetails) SetAudit(v Audit)`

SetAudit sets Audit field to given value.

### HasAudit

`func (o *ProjectDetails) HasAudit() bool`

HasAudit returns a boolean if a field has been set.

### GetProject

`func (o *ProjectDetails) GetProject() Project`

GetProject returns the Project field if non-nil, zero value otherwise.

### GetProjectOk

`func (o *ProjectDetails) GetProjectOk() (*Project, bool)`

GetProjectOk returns a tuple with the Project field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProject

`func (o *ProjectDetails) SetProject(v Project)`

SetProject sets Project field to given value.

### HasProject

`func (o *ProjectDetails) HasProject() bool`

HasProject returns a boolean if a field has been set.

### GetExecutiveOrganizations

`func (o *ProjectDetails) GetExecutiveOrganizations() []Organization`

GetExecutiveOrganizations returns the ExecutiveOrganizations field if non-nil, zero value otherwise.

### GetExecutiveOrganizationsOk

`func (o *ProjectDetails) GetExecutiveOrganizationsOk() (*[]Organization, bool)`

GetExecutiveOrganizationsOk returns a tuple with the ExecutiveOrganizations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecutiveOrganizations

`func (o *ProjectDetails) SetExecutiveOrganizations(v []Organization)`

SetExecutiveOrganizations sets ExecutiveOrganizations field to given value.

### HasExecutiveOrganizations

`func (o *ProjectDetails) HasExecutiveOrganizations() bool`

HasExecutiveOrganizations returns a boolean if a field has been set.

### GetOrganization

`func (o *ProjectDetails) GetOrganization() Organization`

GetOrganization returns the Organization field if non-nil, zero value otherwise.

### GetOrganizationOk

`func (o *ProjectDetails) GetOrganizationOk() (*Organization, bool)`

GetOrganizationOk returns a tuple with the Organization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganization

`func (o *ProjectDetails) SetOrganization(v Organization)`

SetOrganization sets Organization field to given value.

### HasOrganization

`func (o *ProjectDetails) HasOrganization() bool`

HasOrganization returns a boolean if a field has been set.

### GetIndustries

`func (o *ProjectDetails) GetIndustries() []Industry`

GetIndustries returns the Industries field if non-nil, zero value otherwise.

### GetIndustriesOk

`func (o *ProjectDetails) GetIndustriesOk() (*[]Industry, bool)`

GetIndustriesOk returns a tuple with the Industries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndustries

`func (o *ProjectDetails) SetIndustries(v []Industry)`

SetIndustries sets Industries field to given value.

### HasIndustries

`func (o *ProjectDetails) HasIndustries() bool`

HasIndustries returns a boolean if a field has been set.

### GetPersons

`func (o *ProjectDetails) GetPersons() []Person`

GetPersons returns the Persons field if non-nil, zero value otherwise.

### GetPersonsOk

`func (o *ProjectDetails) GetPersonsOk() (*[]Person, bool)`

GetPersonsOk returns a tuple with the Persons field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPersons

`func (o *ProjectDetails) SetPersons(v []Person)`

SetPersons sets Persons field to given value.

### HasPersons

`func (o *ProjectDetails) HasPersons() bool`

HasPersons returns a boolean if a field has been set.

### GetSkillGroups

`func (o *ProjectDetails) GetSkillGroups() []ExperienceSkillGroup`

GetSkillGroups returns the SkillGroups field if non-nil, zero value otherwise.

### GetSkillGroupsOk

`func (o *ProjectDetails) GetSkillGroupsOk() (*[]ExperienceSkillGroup, bool)`

GetSkillGroupsOk returns a tuple with the SkillGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkillGroups

`func (o *ProjectDetails) SetSkillGroups(v []ExperienceSkillGroup)`

SetSkillGroups sets SkillGroups field to given value.

### HasSkillGroups

`func (o *ProjectDetails) HasSkillGroups() bool`

HasSkillGroups returns a boolean if a field has been set.

### GetAwards

`func (o *ProjectDetails) GetAwards() []Award`

GetAwards returns the Awards field if non-nil, zero value otherwise.

### GetAwardsOk

`func (o *ProjectDetails) GetAwardsOk() (*[]Award, bool)`

GetAwardsOk returns a tuple with the Awards field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAwards

`func (o *ProjectDetails) SetAwards(v []Award)`

SetAwards sets Awards field to given value.

### HasAwards

`func (o *ProjectDetails) HasAwards() bool`

HasAwards returns a boolean if a field has been set.

### GetTimeframe

`func (o *ProjectDetails) GetTimeframe() Timeframed`

GetTimeframe returns the Timeframe field if non-nil, zero value otherwise.

### GetTimeframeOk

`func (o *ProjectDetails) GetTimeframeOk() (*Timeframed, bool)`

GetTimeframeOk returns a tuple with the Timeframe field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeframe

`func (o *ProjectDetails) SetTimeframe(v Timeframed)`

SetTimeframe sets Timeframe field to given value.

### HasTimeframe

`func (o *ProjectDetails) HasTimeframe() bool`

HasTimeframe returns a boolean if a field has been set.

### GetObjectType

`func (o *ProjectDetails) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *ProjectDetails) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *ProjectDetails) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.

### HasObjectType

`func (o *ProjectDetails) HasObjectType() bool`

HasObjectType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


