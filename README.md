# Go API client for client

This is an example of using OAuth2 Implicit Flow in a specification to describe security to your API.

## Overview
This API client was generated by the [OpenAPI Generator](https://openapi-generator.tech) project.  By using the [OpenAPI-spec](https://www.openapis.org/) from a remote server, you can easily generate an API client.

- API version: 1.6.1
- Package version: 1.0.0
- Build package: org.openapitools.codegen.languages.GoClientCodegen

## Installation

Install the following dependencies:

```shell
go get github.com/stretchr/testify/assert
go get golang.org/x/net/context
```

Put the package under your project folder and add the following in import:

```golang
import client "github.com/prodyna-yasm/yasm-api-go"
```

To use a proxy, set the environment variable `HTTP_PROXY`:

```golang
os.Setenv("HTTP_PROXY", "http://proxy_name:proxy_port")
```

## Configuration of Server URL

Default configuration comes with `Servers` field that contains server objects as defined in the OpenAPI specification.

### Select Server Configuration

For using other server than the one defined on index 0 set context value `sw.ContextServerIndex` of type `int`.

```golang
ctx := context.WithValue(context.Background(), client.ContextServerIndex, 1)
```

### Templated Server URL

Templated server URL is formatted using default variables from configuration or from context value `sw.ContextServerVariables` of type `map[string]string`.

```golang
ctx := context.WithValue(context.Background(), client.ContextServerVariables, map[string]string{
	"basePath": "v2",
})
```

Note, enum values are always validated and all unused variables are silently ignored.

### URLs Configuration per Operation

Each operation can use different server URL defined using `OperationServers` map in the `Configuration`.
An operation is uniquely identified by `"{classname}Service.{nickname}"` string.
Similar rules for overriding default operation server index and variables applies by using `sw.ContextOperationServerIndices` and `sw.ContextOperationServerVariables` context maps.

```golang
ctx := context.WithValue(context.Background(), client.ContextOperationServerIndices, map[string]int{
	"{classname}Service.{nickname}": 2,
})
ctx = context.WithValue(context.Background(), client.ContextOperationServerVariables, map[string]map[string]string{
	"{classname}Service.{nickname}": {
		"port": "8443",
	},
})
```

## Documentation for API Endpoints

All URIs are relative to *http://localhost:8080/api/v1*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*AvailabilityAPI* | [**CreateAvailability**](docs/AvailabilityAPI.md#createavailability) | **Post** /persons/{personId}/availabilities | Create a availability for a person
*AvailabilityAPI* | [**DeleteAvailability**](docs/AvailabilityAPI.md#deleteavailability) | **Delete** /persons/{personId}/availabilities/{availabilityId} | Delete a person availability
*AvailabilityAPI* | [**GetAvailabilities**](docs/AvailabilityAPI.md#getavailabilities) | **Get** /persons/{personId}/availabilities | Get a list of all activities for a person
*AvailabilityAPI* | [**UpdateAvailability**](docs/AvailabilityAPI.md#updateavailability) | **Put** /persons/{personId}/availabilities/{availabilityId} | Update a person availability
*CertificationAPI* | [**AddPersonCertification**](docs/CertificationAPI.md#addpersoncertification) | **Post** /persons/{personId}/certifications/{certificationId} | Add Certification to a Person
*CertificationAPI* | [**AddSkillToCertification**](docs/CertificationAPI.md#addskilltocertification) | **Post** /certifications/{certificationId}/skills/{skillId} | 
*CertificationAPI* | [**CreateCertification**](docs/CertificationAPI.md#createcertification) | **Post** /organizations/{organizationId}/certifications | Create a Certification in an Organization
*CertificationAPI* | [**DeleteCertification**](docs/CertificationAPI.md#deletecertification) | **Delete** /certifications/{certificationId} | Delete a Certification
*CertificationAPI* | [**DeletePersonCertification**](docs/CertificationAPI.md#deletepersoncertification) | **Delete** /persons/{personId}/certifications/{certificationId} | Remove an Interest to a Person
*CertificationAPI* | [**DeleteSkillFromCertification**](docs/CertificationAPI.md#deleteskillfromcertification) | **Delete** /certifications/{certificationId}/skills/{skillId} | 
*CertificationAPI* | [**GetCertification**](docs/CertificationAPI.md#getcertification) | **Get** /certifications/{certificationId} | Get details about a Certification
*CertificationAPI* | [**GetCertifications**](docs/CertificationAPI.md#getcertifications) | **Get** /certifications | Get a list of all Certifications independent of the Organization
*CertificationAPI* | [**GetCertificationsForOrganization**](docs/CertificationAPI.md#getcertificationsfororganization) | **Get** /organizations/{organizationId}/certifications | Get a list of all certifications for a organization
*CertificationAPI* | [**MoveCertification**](docs/CertificationAPI.md#movecertification) | **Put** /organizations/{organizationId}/certificates/{certificateId} | Move a Certification to an Organization
*CertificationAPI* | [**SearchCertifications**](docs/CertificationAPI.md#searchcertifications) | **Post** /certifications/search | Complex search over certification entities
*CertificationAPI* | [**UpdateCertification**](docs/CertificationAPI.md#updatecertification) | **Put** /certifications/{certificationId} | Update a Certification
*CertificationAPI* | [**UpdatePersonCertification**](docs/CertificationAPI.md#updatepersoncertification) | **Put** /persons/{personId}/certifications/{certificationId} | Update a Certification of a Person
*CertificationAPI* | [**UpdateSkillInCertification**](docs/CertificationAPI.md#updateskillincertification) | **Put** /certifications/{certificationId}/skills/{skillId} | 
*CountryAPI* | [**AddLanguageToCountry**](docs/CountryAPI.md#addlanguagetocountry) | **Post** /countries/{countryId}/languages/{languageId} | Assign a language to a country
*CountryAPI* | [**CreateCountry**](docs/CountryAPI.md#createcountry) | **Post** /countries | Create a new Country
*CountryAPI* | [**DeleteCountry**](docs/CountryAPI.md#deletecountry) | **Delete** /countries/{countryId} | Delete a Country
*CountryAPI* | [**GetCountries**](docs/CountryAPI.md#getcountries) | **Get** /countries | Get all Countries
*CountryAPI* | [**GetCountry**](docs/CountryAPI.md#getcountry) | **Get** /countries/{countryId} | Get details about a Country
*CountryAPI* | [**RemoveLanguageFromCountry**](docs/CountryAPI.md#removelanguagefromcountry) | **Delete** /countries/{countryId}/languages/{languageId} | Assign a language to a country
*CountryAPI* | [**UpdateCountry**](docs/CountryAPI.md#updatecountry) | **Put** /countries/{countryId} | Update a Country
*IndustryAPI* | [**AttachOrganizationToIndustry**](docs/IndustryAPI.md#attachorganizationtoindustry) | **Post** /organizations/{organizationId}/industries/{industryId} | Add an Organization to an Industry
*IndustryAPI* | [**CreateIndustry**](docs/IndustryAPI.md#createindustry) | **Post** /industries | Create an Industry
*IndustryAPI* | [**DeleteIndustry**](docs/IndustryAPI.md#deleteindustry) | **Delete** /industries/{industryId} | Delete an Industry
*IndustryAPI* | [**DetachOrganizationFromIndustry**](docs/IndustryAPI.md#detachorganizationfromindustry) | **Delete** /organizations/{organizationId}/industries/{industryId} | Remove an Organization to an Industry
*IndustryAPI* | [**GetIndustries**](docs/IndustryAPI.md#getindustries) | **Get** /industries | Get all Industries
*IndustryAPI* | [**GetIndustry**](docs/IndustryAPI.md#getindustry) | **Get** /industries/{industryId} | Get details about an Industry
*IndustryAPI* | [**UpdateIndustry**](docs/IndustryAPI.md#updateindustry) | **Put** /industries/{industryId} | Update an Industry
*LanguageAPI* | [**AddLanguageToCountry**](docs/LanguageAPI.md#addlanguagetocountry) | **Post** /countries/{countryId}/languages/{languageId} | Assign a language to a country
*LanguageAPI* | [**AddPersonLanguage**](docs/LanguageAPI.md#addpersonlanguage) | **Post** /persons/{personId}/languages/{languageId} | Assign a language to the person
*LanguageAPI* | [**CreateLanguage**](docs/LanguageAPI.md#createlanguage) | **Post** /languages | Create a new language
*LanguageAPI* | [**DeleteLanguage**](docs/LanguageAPI.md#deletelanguage) | **Delete** /languages/{languageId} | Delete a language
*LanguageAPI* | [**GetLanguage**](docs/LanguageAPI.md#getlanguage) | **Get** /languages/{languageId} | Get details about a language
*LanguageAPI* | [**GetLanguages**](docs/LanguageAPI.md#getlanguages) | **Get** /languages | Get a list of Languages
*LanguageAPI* | [**RemoveLanguageFromCountry**](docs/LanguageAPI.md#removelanguagefromcountry) | **Delete** /countries/{countryId}/languages/{languageId} | Assign a language to a country
*LanguageAPI* | [**RemovePersonLanguage**](docs/LanguageAPI.md#removepersonlanguage) | **Delete** /persons/{personId}/languages/{languageId} | Remove a language from a person
*LanguageAPI* | [**UupdatePersonLanguage**](docs/LanguageAPI.md#uupdatepersonlanguage) | **Put** /persons/{personId}/languages/{languageId} | Update a language of a person
*OfficeAPI* | [**AddPersonOffice**](docs/OfficeAPI.md#addpersonoffice) | **Post** /persons/{personId}/offices/{officeId} | Assign a person to an office
*OfficeAPI* | [**CreateOffice**](docs/OfficeAPI.md#createoffice) | **Post** /organizations/{organizationId}/offices | Create an Office in an Organization
*OfficeAPI* | [**DeleteOffice**](docs/OfficeAPI.md#deleteoffice) | **Delete** /organizations/{organizationId}/offices/{officeId} | Delete an Office from an Organization
*OfficeAPI* | [**DeletePersonOffice**](docs/OfficeAPI.md#deletepersonoffice) | **Delete** /persons/{personId}/offices/{officeId} | Delete the office from a Person
*OfficeAPI* | [**GetOffice**](docs/OfficeAPI.md#getoffice) | **Get** /organizations/{organizationId}/offices/{officeId} | Get an Office for an Organization
*OfficeAPI* | [**GetOfficeDetails**](docs/OfficeAPI.md#getofficedetails) | **Get** /offices/{officeId} | Get details about an Office independent of Organization
*OfficeAPI* | [**GetOffices**](docs/OfficeAPI.md#getoffices) | **Get** /offices | Get a list of all Offices
*OfficeAPI* | [**UpdateOffice**](docs/OfficeAPI.md#updateoffice) | **Put** /organizations/{organizationId}/offices/{officeId} | Update an Office for an Organization
*OrganizationAPI* | [**AttachOrganizationToIndustry**](docs/OrganizationAPI.md#attachorganizationtoindustry) | **Post** /organizations/{organizationId}/industries/{industryId} | Add an Organization to an Industry
*OrganizationAPI* | [**CreateCertification**](docs/OrganizationAPI.md#createcertification) | **Post** /organizations/{organizationId}/certifications | Create a Certification in an Organization
*OrganizationAPI* | [**CreateOffice**](docs/OrganizationAPI.md#createoffice) | **Post** /organizations/{organizationId}/offices | Create an Office in an Organization
*OrganizationAPI* | [**CreateOrganization**](docs/OrganizationAPI.md#createorganization) | **Post** /organizations | Create an Organization
*OrganizationAPI* | [**CreateProject**](docs/OrganizationAPI.md#createproject) | **Post** /organizations/{organizationId}/projects | Create a Project in an Organization
*OrganizationAPI* | [**DeleteOffice**](docs/OrganizationAPI.md#deleteoffice) | **Delete** /organizations/{organizationId}/offices/{officeId} | Delete an Office from an Organization
*OrganizationAPI* | [**DeleteOrganization**](docs/OrganizationAPI.md#deleteorganization) | **Delete** /organizations/{organizationId} | Delete an organization
*OrganizationAPI* | [**DetachOrganizationFromIndustry**](docs/OrganizationAPI.md#detachorganizationfromindustry) | **Delete** /organizations/{organizationId}/industries/{industryId} | Remove an Organization to an Industry
*OrganizationAPI* | [**GetCertificationsForOrganization**](docs/OrganizationAPI.md#getcertificationsfororganization) | **Get** /organizations/{organizationId}/certifications | Get a list of all certifications for a organization
*OrganizationAPI* | [**GetOffice**](docs/OrganizationAPI.md#getoffice) | **Get** /organizations/{organizationId}/offices/{officeId} | Get an Office for an Organization
*OrganizationAPI* | [**GetOrganization**](docs/OrganizationAPI.md#getorganization) | **Get** /organizations/{organizationId} | Get details about an Organization
*OrganizationAPI* | [**GetOrganizationProjects**](docs/OrganizationAPI.md#getorganizationprojects) | **Get** /organizations/{organizationId}/projects | Get a list of all Projects for an Organization
*OrganizationAPI* | [**GetOrganizations**](docs/OrganizationAPI.md#getorganizations) | **Get** /organizations | Get a list of all Organizations
*OrganizationAPI* | [**MergeOrganizations**](docs/OrganizationAPI.md#mergeorganizations) | **Put** /organizations/{organizationId}/merge/{otherOrganizationId} | Merge two organizations
*OrganizationAPI* | [**MoveCertification**](docs/OrganizationAPI.md#movecertification) | **Put** /organizations/{organizationId}/certificates/{certificateId} | Move a Certification to an Organization
*OrganizationAPI* | [**UpdateOffice**](docs/OrganizationAPI.md#updateoffice) | **Put** /organizations/{organizationId}/offices/{officeId} | Update an Office for an Organization
*OrganizationAPI* | [**UpdateOrganization**](docs/OrganizationAPI.md#updateorganization) | **Put** /organizations/{organizationId} | Update an Organization
*OrganizationAPI* | [**UpdateProjectOrganization**](docs/OrganizationAPI.md#updateprojectorganization) | **Put** /organizations/{organizationId}/projects/{projectId} | project is now point to the new organization
*PersonAPI* | [**AddPersonCertification**](docs/PersonAPI.md#addpersoncertification) | **Post** /persons/{personId}/certifications/{certificationId} | Add Certification to a Person
*PersonAPI* | [**AddPersonInterest**](docs/PersonAPI.md#addpersoninterest) | **Post** /persons/{personId}/interests/skills/{skillId} | Add an Interest to a Person
*PersonAPI* | [**AddPersonLanguage**](docs/PersonAPI.md#addpersonlanguage) | **Post** /persons/{personId}/languages/{languageId} | Assign a language to the person
*PersonAPI* | [**AddPersonOffice**](docs/PersonAPI.md#addpersonoffice) | **Post** /persons/{personId}/offices/{officeId} | Assign a person to an office
*PersonAPI* | [**AddPersonProfile**](docs/PersonAPI.md#addpersonprofile) | **Post** /persons/{personId}/profiles/{profileId} | Add a Profile to a Person
*PersonAPI* | [**AddPersonProject**](docs/PersonAPI.md#addpersonproject) | **Post** /persons/{personId}/projects/{projectId} | Add Project to a Person
*PersonAPI* | [**AddPersonProjectSkill**](docs/PersonAPI.md#addpersonprojectskill) | **Post** /persons/{personId}/projects/{projectId}/skills/{skillId} | Add Skill to a Project participation
*PersonAPI* | [**AddPersonSkillExperience**](docs/PersonAPI.md#addpersonskillexperience) | **Post** /persons/{personId}/experiences/skills/{skillId} | Add an Skill experience to a Person
*PersonAPI* | [**AddPersonSkillExperiences**](docs/PersonAPI.md#addpersonskillexperiences) | **Post** /persons/{personId}/experiences | Add an Skill experience to a Person (bulk)
*PersonAPI* | [**ConfirmSkill**](docs/PersonAPI.md#confirmskill) | **Post** /persons/{personId}/projects/{projectId}/skills/{skillId}/confirmation/{confirmingPersonId} | Confirm Skill
*PersonAPI* | [**CreateAvailability**](docs/PersonAPI.md#createavailability) | **Post** /persons/{personId}/availabilities | Create a availability for a person
*PersonAPI* | [**CreatePerson**](docs/PersonAPI.md#createperson) | **Post** /persons | Create a new Person
*PersonAPI* | [**DeleteAvailability**](docs/PersonAPI.md#deleteavailability) | **Delete** /persons/{personId}/availabilities/{availabilityId} | Delete a person availability
*PersonAPI* | [**DeleteConfirmation**](docs/PersonAPI.md#deleteconfirmation) | **Delete** /persons/{personId}/projects/{projectId}/skills/{skillId}/confirmation/{confirmingPersonId} | Remove a confirmation
*PersonAPI* | [**DeletePerson**](docs/PersonAPI.md#deleteperson) | **Delete** /persons/{personId} | Delete an existing Person
*PersonAPI* | [**DeletePersonCertification**](docs/PersonAPI.md#deletepersoncertification) | **Delete** /persons/{personId}/certifications/{certificationId} | Remove an Interest to a Person
*PersonAPI* | [**DeletePersonInterest**](docs/PersonAPI.md#deletepersoninterest) | **Delete** /persons/{personId}/interests/skills/{skillId} | Remove an Interest to a Person
*PersonAPI* | [**DeletePersonOffice**](docs/PersonAPI.md#deletepersonoffice) | **Delete** /persons/{personId}/offices/{officeId} | Delete the office from a Person
*PersonAPI* | [**DeletePersonProfile**](docs/PersonAPI.md#deletepersonprofile) | **Delete** /persons/{personId}/profiles/{profileId} | Remove a Profile from a Person
*PersonAPI* | [**DeletePersonProject**](docs/PersonAPI.md#deletepersonproject) | **Delete** /persons/{personId}/projects/{projectId} | Remove an Project from a Person
*PersonAPI* | [**DeletePersonProjectSkill**](docs/PersonAPI.md#deletepersonprojectskill) | **Delete** /persons/{personId}/projects/{projectId}/skills/{skillId} | Remove a Skill from a Project participation
*PersonAPI* | [**DeletePersonSkillExperience**](docs/PersonAPI.md#deletepersonskillexperience) | **Delete** /persons/{personId}/experiences/skills/{skillId} | Remove an Skill Experience to a Person
*PersonAPI* | [**DeletePersonSkillExperiences**](docs/PersonAPI.md#deletepersonskillexperiences) | **Delete** /persons/{personId}/experiences | Remove an Skill Experience to a Person
*PersonAPI* | [**GeneratePersonProfile**](docs/PersonAPI.md#generatepersonprofile) | **Get** /persons/{personId}/pdf-profile | Generate a PDF profile from a Person
*PersonAPI* | [**GetAllBusinessDepartments**](docs/PersonAPI.md#getallbusinessdepartments) | **Get** /persons/departments | Get all unique business departments
*PersonAPI* | [**GetAvailabilities**](docs/PersonAPI.md#getavailabilities) | **Get** /persons/{personId}/availabilities | Get a list of all activities for a person
*PersonAPI* | [**GetPerson**](docs/PersonAPI.md#getperson) | **Get** /persons/{personId} | Get basic info about a person
*PersonAPI* | [**ReadPersonPicture**](docs/PersonAPI.md#readpersonpicture) | **Get** /persons/{personId}/picture | Read person image
*PersonAPI* | [**ReadPersonProject**](docs/PersonAPI.md#readpersonproject) | **Get** /persons/{personId}/projects/{projectId} | Get a Project Partifipation of a Person
*PersonAPI* | [**RemovePersonLanguage**](docs/PersonAPI.md#removepersonlanguage) | **Delete** /persons/{personId}/languages/{languageId} | Remove a language from a person
*PersonAPI* | [**SearchPersons**](docs/PersonAPI.md#searchpersons) | **Post** /persons/search | Complex search over person entities
*PersonAPI* | [**UpdateAvailability**](docs/PersonAPI.md#updateavailability) | **Put** /persons/{personId}/availabilities/{availabilityId} | Update a person availability
*PersonAPI* | [**UpdatePerson**](docs/PersonAPI.md#updateperson) | **Put** /persons/{personId} | Update an existing Person
*PersonAPI* | [**UpdatePersonCertification**](docs/PersonAPI.md#updatepersoncertification) | **Put** /persons/{personId}/certifications/{certificationId} | Update a Certification of a Person
*PersonAPI* | [**UpdatePersonPicture**](docs/PersonAPI.md#updatepersonpicture) | **Put** /persons/{personId}/picture | Update person image
*PersonAPI* | [**UpdatePersonProject**](docs/PersonAPI.md#updatepersonproject) | **Put** /persons/{personId}/projects/{projectId} | Update a Project of a Person
*PersonAPI* | [**UpdatePersonProjectSkill**](docs/PersonAPI.md#updatepersonprojectskill) | **Put** /persons/{personId}/projects/{projectId}/skills/{skillId} | Update the level of a Skill in a Project participation
*PersonAPI* | [**UpdatePersonSkillExperience**](docs/PersonAPI.md#updatepersonskillexperience) | **Put** /persons/{personId}/experiences/skills/{skillId} | Edit an Skill experience to a Person
*PersonAPI* | [**UpdatePersonSkillExperiences**](docs/PersonAPI.md#updatepersonskillexperiences) | **Put** /persons/{personId}/experiences | Edit an Skill experience to a Person
*PersonAPI* | [**UupdatePersonLanguage**](docs/PersonAPI.md#uupdatepersonlanguage) | **Put** /persons/{personId}/languages/{languageId} | Update a language of a person
*ProfileAPI* | [**AddPersonProfile**](docs/ProfileAPI.md#addpersonprofile) | **Post** /persons/{personId}/profiles/{profileId} | Add a Profile to a Person
*ProfileAPI* | [**CreateProfile**](docs/ProfileAPI.md#createprofile) | **Post** /profiles | Create a new Profile
*ProfileAPI* | [**DeletePersonProfile**](docs/ProfileAPI.md#deletepersonprofile) | **Delete** /persons/{personId}/profiles/{profileId} | Remove a Profile from a Person
*ProfileAPI* | [**DeleteProfile**](docs/ProfileAPI.md#deleteprofile) | **Delete** /profiles/{profileId} | Delete a Profile
*ProfileAPI* | [**GetProfile**](docs/ProfileAPI.md#getprofile) | **Get** /profiles/{profileId} | Get details about a Profile
*ProfileAPI* | [**GetProfiles**](docs/ProfileAPI.md#getprofiles) | **Get** /profiles | Get all Profiles
*ProfileAPI* | [**UpdateProfile**](docs/ProfileAPI.md#updateprofile) | **Put** /profiles/{profileId} | Update a Profile
*ProjectAPI* | [**AddPersonProject**](docs/ProjectAPI.md#addpersonproject) | **Post** /persons/{personId}/projects/{projectId} | Add Project to a Person
*ProjectAPI* | [**AddPersonProjectSkill**](docs/ProjectAPI.md#addpersonprojectskill) | **Post** /persons/{personId}/projects/{projectId}/skills/{skillId} | Add Skill to a Project participation
*ProjectAPI* | [**ConfirmSkill**](docs/ProjectAPI.md#confirmskill) | **Post** /persons/{personId}/projects/{projectId}/skills/{skillId}/confirmation/{confirmingPersonId} | Confirm Skill
*ProjectAPI* | [**CreateProject**](docs/ProjectAPI.md#createproject) | **Post** /organizations/{organizationId}/projects | Create a Project in an Organization
*ProjectAPI* | [**DeleteConfirmation**](docs/ProjectAPI.md#deleteconfirmation) | **Delete** /persons/{personId}/projects/{projectId}/skills/{skillId}/confirmation/{confirmingPersonId} | Remove a confirmation
*ProjectAPI* | [**DeletePersonProject**](docs/ProjectAPI.md#deletepersonproject) | **Delete** /persons/{personId}/projects/{projectId} | Remove an Project from a Person
*ProjectAPI* | [**DeletePersonProjectSkill**](docs/ProjectAPI.md#deletepersonprojectskill) | **Delete** /persons/{personId}/projects/{projectId}/skills/{skillId} | Remove a Skill from a Project participation
*ProjectAPI* | [**DeleteProject**](docs/ProjectAPI.md#deleteproject) | **Delete** /projects/{projectId} | Delete a project
*ProjectAPI* | [**GetOrganizationProjects**](docs/ProjectAPI.md#getorganizationprojects) | **Get** /organizations/{organizationId}/projects | Get a list of all Projects for an Organization
*ProjectAPI* | [**GetProject**](docs/ProjectAPI.md#getproject) | **Get** /projects/{projectId} | Get details about a Project
*ProjectAPI* | [**MergeProjects**](docs/ProjectAPI.md#mergeprojects) | **Put** /projects/{projectId}/merge/{otherProjectId} | Merge to projects
*ProjectAPI* | [**ReadPersonProject**](docs/ProjectAPI.md#readpersonproject) | **Get** /persons/{personId}/projects/{projectId} | Get a Project Partifipation of a Person
*ProjectAPI* | [**SearchProjectParticipations**](docs/ProjectAPI.md#searchprojectparticipations) | **Post** /project-participations/search | Complex search over project entities
*ProjectAPI* | [**SearchProjects**](docs/ProjectAPI.md#searchprojects) | **Post** /projects/search | Complex search over project entities
*ProjectAPI* | [**UpdatePersonProject**](docs/ProjectAPI.md#updatepersonproject) | **Put** /persons/{personId}/projects/{projectId} | Update a Project of a Person
*ProjectAPI* | [**UpdatePersonProjectSkill**](docs/ProjectAPI.md#updatepersonprojectskill) | **Put** /persons/{personId}/projects/{projectId}/skills/{skillId} | Update the level of a Skill in a Project participation
*ProjectAPI* | [**UpdateProject**](docs/ProjectAPI.md#updateproject) | **Put** /projects/{projectId} | Update a Project
*ProjectAPI* | [**UpdateProjectOrganization**](docs/ProjectAPI.md#updateprojectorganization) | **Put** /organizations/{organizationId}/projects/{projectId} | project is now point to the new organization
*SearchAPI* | [**SearchAll**](docs/SearchAPI.md#searchall) | **Get** /search/all/{text} | Fulltext search on all kinds of objects
*SkillAPI* | [**AddPersonInterest**](docs/SkillAPI.md#addpersoninterest) | **Post** /persons/{personId}/interests/skills/{skillId} | Add an Interest to a Person
*SkillAPI* | [**AddPersonProjectSkill**](docs/SkillAPI.md#addpersonprojectskill) | **Post** /persons/{personId}/projects/{projectId}/skills/{skillId} | Add Skill to a Project participation
*SkillAPI* | [**AddPersonSkillExperience**](docs/SkillAPI.md#addpersonskillexperience) | **Post** /persons/{personId}/experiences/skills/{skillId} | Add an Skill experience to a Person
*SkillAPI* | [**AddPersonSkillExperiences**](docs/SkillAPI.md#addpersonskillexperiences) | **Post** /persons/{personId}/experiences | Add an Skill experience to a Person (bulk)
*SkillAPI* | [**AddSkillToCertification**](docs/SkillAPI.md#addskilltocertification) | **Post** /certifications/{certificationId}/skills/{skillId} | 
*SkillAPI* | [**AddSkillToParentSkill**](docs/SkillAPI.md#addskilltoparentskill) | **Post** /skills/{skillId}/parents/{parentSkillId} | Attach a Skill to a parent Skill, returns the parent Skill
*SkillAPI* | [**ConfirmSkill**](docs/SkillAPI.md#confirmskill) | **Post** /persons/{personId}/projects/{projectId}/skills/{skillId}/confirmation/{confirmingPersonId} | Confirm Skill
*SkillAPI* | [**CreateSkill**](docs/SkillAPI.md#createskill) | **Post** /skills | Create a Skill
*SkillAPI* | [**DeleteConfirmation**](docs/SkillAPI.md#deleteconfirmation) | **Delete** /persons/{personId}/projects/{projectId}/skills/{skillId}/confirmation/{confirmingPersonId} | Remove a confirmation
*SkillAPI* | [**DeletePersonInterest**](docs/SkillAPI.md#deletepersoninterest) | **Delete** /persons/{personId}/interests/skills/{skillId} | Remove an Interest to a Person
*SkillAPI* | [**DeletePersonProjectSkill**](docs/SkillAPI.md#deletepersonprojectskill) | **Delete** /persons/{personId}/projects/{projectId}/skills/{skillId} | Remove a Skill from a Project participation
*SkillAPI* | [**DeletePersonSkillExperience**](docs/SkillAPI.md#deletepersonskillexperience) | **Delete** /persons/{personId}/experiences/skills/{skillId} | Remove an Skill Experience to a Person
*SkillAPI* | [**DeletePersonSkillExperiences**](docs/SkillAPI.md#deletepersonskillexperiences) | **Delete** /persons/{personId}/experiences | Remove an Skill Experience to a Person
*SkillAPI* | [**DeleteSkill**](docs/SkillAPI.md#deleteskill) | **Delete** /skills/{skillId} | Delete a Skill
*SkillAPI* | [**DeleteSkillFromCertification**](docs/SkillAPI.md#deleteskillfromcertification) | **Delete** /certifications/{certificationId}/skills/{skillId} | 
*SkillAPI* | [**GetSkill**](docs/SkillAPI.md#getskill) | **Get** /skills/{skillId} | Get details for a single skill
*SkillAPI* | [**GetSkillParents**](docs/SkillAPI.md#getskillparents) | **Get** /skills/{skillId}/parents | Get ghe list of parents for a skill
*SkillAPI* | [**GetSkills**](docs/SkillAPI.md#getskills) | **Get** /skills | Get a list of all skills, optionally only root, optionally only kinds
*SkillAPI* | [**MergeSkills**](docs/SkillAPI.md#mergeskills) | **Put** /skills/{skillId}/merge/{otherSkillId} | Merge two skills
*SkillAPI* | [**RemoveSkillFromParentSkill**](docs/SkillAPI.md#removeskillfromparentskill) | **Delete** /skills/{skillId}/parents/{parentSkillId} | Detaches a Skill from parent Skill, return the parent Skill
*SkillAPI* | [**UpdatePersonProjectSkill**](docs/SkillAPI.md#updatepersonprojectskill) | **Put** /persons/{personId}/projects/{projectId}/skills/{skillId} | Update the level of a Skill in a Project participation
*SkillAPI* | [**UpdatePersonSkillExperience**](docs/SkillAPI.md#updatepersonskillexperience) | **Put** /persons/{personId}/experiences/skills/{skillId} | Edit an Skill experience to a Person
*SkillAPI* | [**UpdatePersonSkillExperiences**](docs/SkillAPI.md#updatepersonskillexperiences) | **Put** /persons/{personId}/experiences | Edit an Skill experience to a Person
*SkillAPI* | [**UpdateSkill**](docs/SkillAPI.md#updateskill) | **Put** /skills/{skillId} | Update a Skill
*SkillAPI* | [**UpdateSkillInCertification**](docs/SkillAPI.md#updateskillincertification) | **Put** /certifications/{certificationId}/skills/{skillId} | 
*StatusAPI* | [**GetPing**](docs/StatusAPI.md#getping) | **Get** /ping | Server heartbeat operation
*StatusAPI* | [**GetVersion**](docs/StatusAPI.md#getversion) | **Get** /version | Information about the server
*TemplatesAPI* | [**GetTemplates**](docs/TemplatesAPI.md#gettemplates) | **Get** /templates | list of existing templates


## Documentation For Models

 - [AbstractEntityFilter](docs/AbstractEntityFilter.md)
 - [Availability](docs/Availability.md)
 - [AvailabilityDetail](docs/AvailabilityDetail.md)
 - [AvailabilityFilter](docs/AvailabilityFilter.md)
 - [BasicDomainModel](docs/BasicDomainModel.md)
 - [Certification](docs/Certification.md)
 - [CertificationDetails](docs/CertificationDetails.md)
 - [CertificationSearch](docs/CertificationSearch.md)
 - [Country](docs/Country.md)
 - [CountryDetails](docs/CountryDetails.md)
 - [Descriptable](docs/Descriptable.md)
 - [EntityFilter](docs/EntityFilter.md)
 - [Error](docs/Error.md)
 - [Experience](docs/Experience.md)
 - [ExperienceSkill](docs/ExperienceSkill.md)
 - [ExperienceSkillGroup](docs/ExperienceSkillGroup.md)
 - [Geolocation](docs/Geolocation.md)
 - [Industry](docs/Industry.md)
 - [IndustryDetails](docs/IndustryDetails.md)
 - [Language](docs/Language.md)
 - [LanguageDetails](docs/LanguageDetails.md)
 - [LanguageLevel](docs/LanguageLevel.md)
 - [Level](docs/Level.md)
 - [Linkable](docs/Linkable.md)
 - [Locateable](docs/Locateable.md)
 - [MinMax](docs/MinMax.md)
 - [MinMaxPercent](docs/MinMaxPercent.md)
 - [Nameable](docs/Nameable.md)
 - [NamedDomainModel](docs/NamedDomainModel.md)
 - [Office](docs/Office.md)
 - [Organization](docs/Organization.md)
 - [OrganizationDetails](docs/OrganizationDetails.md)
 - [Page](docs/Page.md)
 - [PagedAvailabilities](docs/PagedAvailabilities.md)
 - [PagedCertifications](docs/PagedCertifications.md)
 - [PagedCountries](docs/PagedCountries.md)
 - [PagedIndustries](docs/PagedIndustries.md)
 - [PagedLanguages](docs/PagedLanguages.md)
 - [PagedOffices](docs/PagedOffices.md)
 - [PagedOrganizations](docs/PagedOrganizations.md)
 - [PagedPersons](docs/PagedPersons.md)
 - [PagedProfiles](docs/PagedProfiles.md)
 - [PagedProjectParticipations](docs/PagedProjectParticipations.md)
 - [PagedProjects](docs/PagedProjects.md)
 - [PagedSkills](docs/PagedSkills.md)
 - [Person](docs/Person.md)
 - [PersonCertificationFilter](docs/PersonCertificationFilter.md)
 - [PersonDetails](docs/PersonDetails.md)
 - [PersonIndustryFilter](docs/PersonIndustryFilter.md)
 - [PersonOrganizationFilter](docs/PersonOrganizationFilter.md)
 - [PersonProjectFilter](docs/PersonProjectFilter.md)
 - [PersonScoreDetail](docs/PersonScoreDetail.md)
 - [PersonSearch](docs/PersonSearch.md)
 - [PersonSkillFilter](docs/PersonSkillFilter.md)
 - [Profile](docs/Profile.md)
 - [ProfileDetails](docs/ProfileDetails.md)
 - [Project](docs/Project.md)
 - [ProjectDetails](docs/ProjectDetails.md)
 - [ProjectParticipation](docs/ProjectParticipation.md)
 - [ProjectParticipationSearch](docs/ProjectParticipationSearch.md)
 - [ProjectParticipationUpdate](docs/ProjectParticipationUpdate.md)
 - [ProjectScoreDetail](docs/ProjectScoreDetail.md)
 - [ProjectSearch](docs/ProjectSearch.md)
 - [ProjectStatus](docs/ProjectStatus.md)
 - [ProjectType](docs/ProjectType.md)
 - [ScoreResult](docs/ScoreResult.md)
 - [SearchResult](docs/SearchResult.md)
 - [SearchResultItem](docs/SearchResultItem.md)
 - [Seniority](docs/Seniority.md)
 - [Skill](docs/Skill.md)
 - [SkillDetails](docs/SkillDetails.md)
 - [SkillGroup](docs/SkillGroup.md)
 - [SkillLevel](docs/SkillLevel.md)
 - [SkillLevelUpdate](docs/SkillLevelUpdate.md)
 - [Status](docs/Status.md)
 - [Suggestable](docs/Suggestable.md)
 - [Synonymable](docs/Synonymable.md)
 - [Timeframed](docs/Timeframed.md)
 - [Version](docs/Version.md)


## Documentation For Authorization


Authentication schemes defined for the API:
### bearerAuth

- **Type**: HTTP Bearer token authentication

Example

```golang
auth := context.WithValue(context.Background(), sw.ContextAccessToken, "BEARER_TOKEN_STRING")
r, err := client.Service.Operation(auth, args)
```


## Documentation for Utility Methods

Due to the fact that model structure members are all pointers, this package contains
a number of utility functions to easily obtain pointers to values of basic types.
Each of these functions takes a value of the given basic type and returns a pointer to it:

* `PtrBool`
* `PtrInt`
* `PtrInt32`
* `PtrInt64`
* `PtrFloat`
* `PtrFloat32`
* `PtrFloat64`
* `PtrString`
* `PtrTime`

## Author



