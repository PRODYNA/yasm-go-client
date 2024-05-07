# Go API client for client

This is an example of using OAuth2 Implicit Flow in a specification to describe security to your API.

## Overview
This API client was generated by the [OpenAPI Generator](https://openapi-generator.tech) project.  By using the [OpenAPI-spec](https://www.openapis.org/) from a remote server, you can easily generate an API client.

- API version: 1.25.0
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

All URIs are relative to *https://yasm.prodyna.com:443/api/v1*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*AvailabilityAPI* | [**CreateAvailability**](docs/AvailabilityAPI.md#createavailability) | **Post** /persons/{personId}/availabilities | Create a availability for a person
*AvailabilityAPI* | [**DeleteAvailability**](docs/AvailabilityAPI.md#deleteavailability) | **Delete** /persons/{personId}/availabilities/{availabilityId} | Delete a person availability
*AvailabilityAPI* | [**GetAvailabilities**](docs/AvailabilityAPI.md#getavailabilities) | **Get** /persons/{personId}/availabilities | Get a list of all activities for a person
*AvailabilityAPI* | [**UpdateAvailability**](docs/AvailabilityAPI.md#updateavailability) | **Put** /persons/{personId}/availabilities/{availabilityId} | Update a person availability
*CertificationAPI* | [**AddPersonCertification**](docs/CertificationAPI.md#addpersoncertification) | **Post** /persons/{personId}/certifications/{certificationId} | Add Certification to a Person
*CertificationAPI* | [**AddSkillToCertification**](docs/CertificationAPI.md#addskilltocertification) | **Post** /certifications/{certificationId}/skills/{skillId} | 
*CertificationAPI* | [**CreateCertification**](docs/CertificationAPI.md#createcertification) | **Post** /certifications | Create a Certification in an Organization
*CertificationAPI* | [**DeleteCertification**](docs/CertificationAPI.md#deletecertification) | **Delete** /certifications/{certificationId} | Delete a Certification
*CertificationAPI* | [**DeletePersonCertification**](docs/CertificationAPI.md#deletepersoncertification) | **Delete** /persons/{personId}/certifications/{certificationId} | Remove an Interest to a Person
*CertificationAPI* | [**DeleteSkillFromCertification**](docs/CertificationAPI.md#deleteskillfromcertification) | **Delete** /certifications/{certificationId}/skills/{skillId} | 
*CertificationAPI* | [**GetCertification**](docs/CertificationAPI.md#getcertification) | **Get** /certifications/{certificationId} | Get details about a Certification
*CertificationAPI* | [**MoveCertification**](docs/CertificationAPI.md#movecertification) | **Put** /certifications/{certificationId}/organizations/{organizationId} | Move a Certification to an Organization
*CertificationAPI* | [**SearchCertifications**](docs/CertificationAPI.md#searchcertifications) | **Post** /certifications/search | Search over certifications
*CertificationAPI* | [**UpdateCertification**](docs/CertificationAPI.md#updatecertification) | **Put** /certifications/{certificationId} | Update a Certification
*CertificationAPI* | [**UpdatePersonCertification**](docs/CertificationAPI.md#updatepersoncertification) | **Put** /persons/{personId}/certifications/{certificationId} | Update a Certification of a Person
*CertificationAPI* | [**UpdateSkillInCertification**](docs/CertificationAPI.md#updateskillincertification) | **Put** /certifications/{certificationId}/skills/{skillId} | 
*CountryAPI* | [**AddLanguageToCountry**](docs/CountryAPI.md#addlanguagetocountry) | **Post** /countries/{countryId}/languages/{languageId} | Assign a language to a country
*CountryAPI* | [**CreateCountry**](docs/CountryAPI.md#createcountry) | **Post** /countries | Create a new Country
*CountryAPI* | [**DeleteCountry**](docs/CountryAPI.md#deletecountry) | **Delete** /countries/{countryId} | Delete a Country
*CountryAPI* | [**GetCountry**](docs/CountryAPI.md#getcountry) | **Get** /countries/{countryId} | Get details about a Country
*CountryAPI* | [**RemoveLanguageFromCountry**](docs/CountryAPI.md#removelanguagefromcountry) | **Delete** /countries/{countryId}/languages/{languageId} | Assign a language to a country
*CountryAPI* | [**SearchCountries**](docs/CountryAPI.md#searchcountries) | **Post** /countries/search | Search over countries
*CountryAPI* | [**UpdateCountry**](docs/CountryAPI.md#updatecountry) | **Put** /countries/{countryId} | Update a Country
*CountryAPI* | [**UpdateOfficeCountry**](docs/CountryAPI.md#updateofficecountry) | **Put** /offices/{officeId}/countries/{countryId} | Assign a country to an office
*IndustryAPI* | [**AttachOrganizationToIndustry**](docs/IndustryAPI.md#attachorganizationtoindustry) | **Post** /organizations/{organizationId}/industries/{industryId} | Add an Organization to an Industry
*IndustryAPI* | [**CreateIndustry**](docs/IndustryAPI.md#createindustry) | **Post** /industries | Create an Industry
*IndustryAPI* | [**DeleteIndustry**](docs/IndustryAPI.md#deleteindustry) | **Delete** /industries/{industryId} | Delete an Industry
*IndustryAPI* | [**DetachOrganizationFromIndustry**](docs/IndustryAPI.md#detachorganizationfromindustry) | **Delete** /organizations/{organizationId}/industries/{industryId} | Remove an Organization to an Industry
*IndustryAPI* | [**GetIndustry**](docs/IndustryAPI.md#getindustry) | **Get** /industries/{industryId} | Get details about an Industry
*IndustryAPI* | [**SearchIndustries**](docs/IndustryAPI.md#searchindustries) | **Post** /industries/search | Search over industries
*IndustryAPI* | [**UpdateIndustry**](docs/IndustryAPI.md#updateindustry) | **Put** /industries/{industryId} | Update an Industry
*LanguageAPI* | [**AddLanguageToCountry**](docs/LanguageAPI.md#addlanguagetocountry) | **Post** /countries/{countryId}/languages/{languageId} | Assign a language to a country
*LanguageAPI* | [**AddPersonLanguage**](docs/LanguageAPI.md#addpersonlanguage) | **Post** /persons/{personId}/languages/{languageId} | Assign a language to the person
*LanguageAPI* | [**CreateLanguage**](docs/LanguageAPI.md#createlanguage) | **Post** /languages | Create a new language
*LanguageAPI* | [**DeleteLanguage**](docs/LanguageAPI.md#deletelanguage) | **Delete** /languages/{languageId} | Delete a language
*LanguageAPI* | [**GetLanguage**](docs/LanguageAPI.md#getlanguage) | **Get** /languages/{languageId} | Get details about a language
*LanguageAPI* | [**RemoveLanguageFromCountry**](docs/LanguageAPI.md#removelanguagefromcountry) | **Delete** /countries/{countryId}/languages/{languageId} | Assign a language to a country
*LanguageAPI* | [**RemovePersonLanguage**](docs/LanguageAPI.md#removepersonlanguage) | **Delete** /persons/{personId}/languages/{languageId} | Remove a language from a person
*LanguageAPI* | [**SearchLanguages**](docs/LanguageAPI.md#searchlanguages) | **Post** /languages/search | Search over languages
*LanguageAPI* | [**UpdatePersonLanguage**](docs/LanguageAPI.md#updatepersonlanguage) | **Put** /persons/{personId}/languages/{languageId} | Update a language of a person
*OfficeAPI* | [**AddOrganizationOffice**](docs/OfficeAPI.md#addorganizationoffice) | **Post** /organizations/{organizationId}/offices/{officeId} | Assign an office to an organization
*OfficeAPI* | [**CreateOffice**](docs/OfficeAPI.md#createoffice) | **Post** /offices | Create an Office in an Organization
*OfficeAPI* | [**DeleteOffice**](docs/OfficeAPI.md#deleteoffice) | **Delete** /offices/{officeId} | Delete an Office
*OfficeAPI* | [**DeleteOfficeCountry**](docs/OfficeAPI.md#deleteofficecountry) | **Delete** /offices/{officeId}/countries/{countryId} | Delete the office from a Person
*OfficeAPI* | [**DeleteOrganizationOffice**](docs/OfficeAPI.md#deleteorganizationoffice) | **Delete** /organizations/{organizationId}/offices/{officeId} | Delete an office from an organization
*OfficeAPI* | [**DeletePersonOffice**](docs/OfficeAPI.md#deletepersonoffice) | **Delete** /persons/{personId}/offices/{officeId} | Delete the office from a Person
*OfficeAPI* | [**GetOfficeDetails**](docs/OfficeAPI.md#getofficedetails) | **Get** /offices/{officeId} | Get details about an Office
*OfficeAPI* | [**SearchOffices**](docs/OfficeAPI.md#searchoffices) | **Post** /offices/search | Search over offices
*OfficeAPI* | [**UpdateOffice**](docs/OfficeAPI.md#updateoffice) | **Put** /offices/{officeId} | Update an Office
*OfficeAPI* | [**UpdateOfficeCountry**](docs/OfficeAPI.md#updateofficecountry) | **Put** /offices/{officeId}/countries/{countryId} | Assign a country to an office
*OfficeAPI* | [**UpdatePersonOffice**](docs/OfficeAPI.md#updatepersonoffice) | **Put** /persons/{personId}/offices/{officeId} | Assign a person to an office
*OrganizationAPI* | [**AddExecutiveOrganizationToProject**](docs/OrganizationAPI.md#addexecutiveorganizationtoproject) | **Post** /projects/{projectId}/executive-organizations/{organizationId} | Add an Organization to a Project as executive organization
*OrganizationAPI* | [**AddOrganizationOffice**](docs/OrganizationAPI.md#addorganizationoffice) | **Post** /organizations/{organizationId}/offices/{officeId} | Assign an office to an organization
*OrganizationAPI* | [**AddOrganizationToParentOrganization**](docs/OrganizationAPI.md#addorganizationtoparentorganization) | **Post** /organizations/{organizationId}/parents/{parentOrganizationId} | Attach an Organization to a parent Organization, returns the parent Organization
*OrganizationAPI* | [**AttachOrganizationToIndustry**](docs/OrganizationAPI.md#attachorganizationtoindustry) | **Post** /organizations/{organizationId}/industries/{industryId} | Add an Organization to an Industry
*OrganizationAPI* | [**CreateOffice**](docs/OrganizationAPI.md#createoffice) | **Post** /offices | Create an Office in an Organization
*OrganizationAPI* | [**CreateOrganization**](docs/OrganizationAPI.md#createorganization) | **Post** /organizations | Create an Organization
*OrganizationAPI* | [**CreateProject**](docs/OrganizationAPI.md#createproject) | **Post** /projects | Create a Project in an Organization
*OrganizationAPI* | [**DeleteOrganization**](docs/OrganizationAPI.md#deleteorganization) | **Delete** /organizations/{organizationId} | Delete an organization
*OrganizationAPI* | [**DeleteOrganizationOffice**](docs/OrganizationAPI.md#deleteorganizationoffice) | **Delete** /organizations/{organizationId}/offices/{officeId} | Delete an office from an organization
*OrganizationAPI* | [**DetachOrganizationFromIndustry**](docs/OrganizationAPI.md#detachorganizationfromindustry) | **Delete** /organizations/{organizationId}/industries/{industryId} | Remove an Organization to an Industry
*OrganizationAPI* | [**GetOrganization**](docs/OrganizationAPI.md#getorganization) | **Get** /organizations/{organizationId} | Get details about an Organization
*OrganizationAPI* | [**MoveCertification**](docs/OrganizationAPI.md#movecertification) | **Put** /certifications/{certificationId}/organizations/{organizationId} | Move a Certification to an Organization
*OrganizationAPI* | [**MoveProject**](docs/OrganizationAPI.md#moveproject) | **Put** /projects/{projectId}/organizations/{organizationId} | Move a Project to an Organization
*OrganizationAPI* | [**RemoveExecutiveOrganizationFromProject**](docs/OrganizationAPI.md#removeexecutiveorganizationfromproject) | **Delete** /projects/{projectId}/executive-organizations/{organizationId} | Remove an Organization from a Project as executive organization
*OrganizationAPI* | [**RemoveOrganizationFromParentOrganization**](docs/OrganizationAPI.md#removeorganizationfromparentorganization) | **Delete** /organizations/{organizationId}/parents/{parentOrganizationId} | Detaches an Organization from parent Organization, return the parent Organization
*OrganizationAPI* | [**RemoveOrganizationServiceManager**](docs/OrganizationAPI.md#removeorganizationservicemanager) | **Delete** /organizations/{organizationId}/service-manager/{personId} | Remove service manager from an Organization
*OrganizationAPI* | [**SearchOrganizations**](docs/OrganizationAPI.md#searchorganizations) | **Post** /organizations/search | Search over organizations
*OrganizationAPI* | [**UpdateOrganization**](docs/OrganizationAPI.md#updateorganization) | **Put** /organizations/{organizationId} | Update an Organization
*OrganizationAPI* | [**UpdateOrganizationServiceManager**](docs/OrganizationAPI.md#updateorganizationservicemanager) | **Put** /organizations/{organizationId}/service-manager/{personId} | Update service manager of an Organization
*PersonAPI* | [**AddPersonCertification**](docs/PersonAPI.md#addpersoncertification) | **Post** /persons/{personId}/certifications/{certificationId} | Add Certification to a Person
*PersonAPI* | [**AddPersonInterest**](docs/PersonAPI.md#addpersoninterest) | **Post** /persons/{personId}/interests/skills/{skillId} | Add an Interest to a Person
*PersonAPI* | [**AddPersonLanguage**](docs/PersonAPI.md#addpersonlanguage) | **Post** /persons/{personId}/languages/{languageId} | Assign a language to the person
*PersonAPI* | [**AddPersonProfile**](docs/PersonAPI.md#addpersonprofile) | **Post** /persons/{personId}/profiles/{profileId} | Add a Profile to a Person
*PersonAPI* | [**AddPersonSkillExperience**](docs/PersonAPI.md#addpersonskillexperience) | **Post** /persons/{personId}/experiences/skills/{skillId} | Add an Skill experience to a Person
*PersonAPI* | [**AddPersonSkillExperiences**](docs/PersonAPI.md#addpersonskillexperiences) | **Post** /persons/{personId}/experiences | Add an Skill experience to a Person (bulk)
*PersonAPI* | [**AddProjectParticipation**](docs/PersonAPI.md#addprojectparticipation) | **Post** /project-participations | Add Project to a Person
*PersonAPI* | [**AddSkillConfirmation**](docs/PersonAPI.md#addskillconfirmation) | **Post** /project-participations/{projectParticipationId}/skills/{skillId}/confirmation/{confirmingPersonId} | Confirm Skill
*PersonAPI* | [**CreateAvailability**](docs/PersonAPI.md#createavailability) | **Post** /persons/{personId}/availabilities | Create a availability for a person
*PersonAPI* | [**CreatePerson**](docs/PersonAPI.md#createperson) | **Post** /persons | Create a new Person
*PersonAPI* | [**DeleteAvailability**](docs/PersonAPI.md#deleteavailability) | **Delete** /persons/{personId}/availabilities/{availabilityId} | Delete a person availability
*PersonAPI* | [**DeleteOfficeCountry**](docs/PersonAPI.md#deleteofficecountry) | **Delete** /offices/{officeId}/countries/{countryId} | Delete the office from a Person
*PersonAPI* | [**DeletePerson**](docs/PersonAPI.md#deleteperson) | **Delete** /persons/{personId} | Delete an existing Person
*PersonAPI* | [**DeletePersonCertification**](docs/PersonAPI.md#deletepersoncertification) | **Delete** /persons/{personId}/certifications/{certificationId} | Remove an Interest to a Person
*PersonAPI* | [**DeletePersonInterest**](docs/PersonAPI.md#deletepersoninterest) | **Delete** /persons/{personId}/interests/skills/{skillId} | Remove an Interest to a Person
*PersonAPI* | [**DeletePersonOffice**](docs/PersonAPI.md#deletepersonoffice) | **Delete** /persons/{personId}/offices/{officeId} | Delete the office from a Person
*PersonAPI* | [**DeletePersonPicture**](docs/PersonAPI.md#deletepersonpicture) | **Delete** /persons/{personId}/picture | Delete person image
*PersonAPI* | [**DeletePersonProfile**](docs/PersonAPI.md#deletepersonprofile) | **Delete** /persons/{personId}/profiles/{profileId} | Remove a Profile from a Person
*PersonAPI* | [**DeletePersonSkillExperience**](docs/PersonAPI.md#deletepersonskillexperience) | **Delete** /persons/{personId}/experiences/skills/{skillId} | Remove an Skill Experience to a Person
*PersonAPI* | [**DeletePersonSkillExperiences**](docs/PersonAPI.md#deletepersonskillexperiences) | **Delete** /persons/{personId}/experiences | Remove an Skill Experience to a Person
*PersonAPI* | [**DeleteProjectParticipation**](docs/PersonAPI.md#deleteprojectparticipation) | **Delete** /project-participations/{projectParticipationId} | Remove an Project from a Person
*PersonAPI* | [**GeneratePersonProfile**](docs/PersonAPI.md#generatepersonprofile) | **Get** /persons/{personId}/pdf-profile | Generate a PDF profile from a Person
*PersonAPI* | [**GetAvailabilities**](docs/PersonAPI.md#getavailabilities) | **Get** /persons/{personId}/availabilities | Get a list of all activities for a person
*PersonAPI* | [**GetPerson**](docs/PersonAPI.md#getperson) | **Get** /persons/{personId} | Get basic info about a person
*PersonAPI* | [**ReadPersonPicture**](docs/PersonAPI.md#readpersonpicture) | **Get** /persons/{personId}/picture | Read person image
*PersonAPI* | [**ReadPersonProjectParticipation**](docs/PersonAPI.md#readpersonprojectparticipation) | **Get** /persons/{personId}/project-participation | Get a Project Participation of a Person
*PersonAPI* | [**ReadProjectParticipation**](docs/PersonAPI.md#readprojectparticipation) | **Get** /project-participations/{projectParticipationId} | Get a project participation
*PersonAPI* | [**RemoveManager**](docs/PersonAPI.md#removemanager) | **Delete** /persons/{personId}/manager | Remove a manager from a person
*PersonAPI* | [**RemoveOrganizationServiceManager**](docs/PersonAPI.md#removeorganizationservicemanager) | **Delete** /organizations/{organizationId}/service-manager/{personId} | Remove service manager from an Organization
*PersonAPI* | [**RemovePersonLanguage**](docs/PersonAPI.md#removepersonlanguage) | **Delete** /persons/{personId}/languages/{languageId} | Remove a language from a person
*PersonAPI* | [**RemoveSkillConfirmation**](docs/PersonAPI.md#removeskillconfirmation) | **Delete** /project-participations/{projectParticipationId}/skills/{skillId}/confirmation/{confirmingPersonId} | Remove a confirmation
*PersonAPI* | [**SearchPersons**](docs/PersonAPI.md#searchpersons) | **Post** /persons/search | Search over persons
*PersonAPI* | [**SearchProjectParticipations**](docs/PersonAPI.md#searchprojectparticipations) | **Post** /project-participations/search | Search over project participations
*PersonAPI* | [**SetManager**](docs/PersonAPI.md#setmanager) | **Put** /persons/{personId}/manager/{managerId} | Set a manager for a person
*PersonAPI* | [**UpdateAvailability**](docs/PersonAPI.md#updateavailability) | **Put** /persons/{personId}/availabilities/{availabilityId} | Update a person availability
*PersonAPI* | [**UpdateOrganizationServiceManager**](docs/PersonAPI.md#updateorganizationservicemanager) | **Put** /organizations/{organizationId}/service-manager/{personId} | Update service manager of an Organization
*PersonAPI* | [**UpdatePerson**](docs/PersonAPI.md#updateperson) | **Put** /persons/{personId} | Update an existing Person
*PersonAPI* | [**UpdatePersonCertification**](docs/PersonAPI.md#updatepersoncertification) | **Put** /persons/{personId}/certifications/{certificationId} | Update a Certification of a Person
*PersonAPI* | [**UpdatePersonLanguage**](docs/PersonAPI.md#updatepersonlanguage) | **Put** /persons/{personId}/languages/{languageId} | Update a language of a person
*PersonAPI* | [**UpdatePersonOffice**](docs/PersonAPI.md#updatepersonoffice) | **Put** /persons/{personId}/offices/{officeId} | Assign a person to an office
*PersonAPI* | [**UpdatePersonPicture**](docs/PersonAPI.md#updatepersonpicture) | **Put** /persons/{personId}/picture | Update person image
*PersonAPI* | [**UpdatePersonSkillExperience**](docs/PersonAPI.md#updatepersonskillexperience) | **Put** /persons/{personId}/experiences/skills/{skillId} | Edit an Skill experience to a Person
*PersonAPI* | [**UpdatePersonSkillExperiences**](docs/PersonAPI.md#updatepersonskillexperiences) | **Put** /persons/{personId}/experiences | Edit an Skill experience to a Person
*PersonAPI* | [**UpdateProjectParticipation**](docs/PersonAPI.md#updateprojectparticipation) | **Put** /project-participations/{projectParticipationId} | Update a Project of a Person
*ProfileAPI* | [**AddPersonProfile**](docs/ProfileAPI.md#addpersonprofile) | **Post** /persons/{personId}/profiles/{profileId} | Add a Profile to a Person
*ProfileAPI* | [**CreateProfile**](docs/ProfileAPI.md#createprofile) | **Post** /profiles | Create a new Profile
*ProfileAPI* | [**DeletePersonProfile**](docs/ProfileAPI.md#deletepersonprofile) | **Delete** /persons/{personId}/profiles/{profileId} | Remove a Profile from a Person
*ProfileAPI* | [**DeleteProfile**](docs/ProfileAPI.md#deleteprofile) | **Delete** /profiles/{profileId} | Delete a Profile
*ProfileAPI* | [**GetProfile**](docs/ProfileAPI.md#getprofile) | **Get** /profiles/{profileId} | Get details about a Profile
*ProfileAPI* | [**SearchProfiles**](docs/ProfileAPI.md#searchprofiles) | **Post** /profiles/search | Search over profiles
*ProfileAPI* | [**UpdateProfile**](docs/ProfileAPI.md#updateprofile) | **Put** /profiles/{profileId} | Update a Profile
*ProjectAPI* | [**AddExecutiveOrganizationToProject**](docs/ProjectAPI.md#addexecutiveorganizationtoproject) | **Post** /projects/{projectId}/executive-organizations/{organizationId} | Add an Organization to a Project as executive organization
*ProjectAPI* | [**AddProjectParticipation**](docs/ProjectAPI.md#addprojectparticipation) | **Post** /project-participations | Add Project to a Person
*ProjectAPI* | [**AddSkillConfirmation**](docs/ProjectAPI.md#addskillconfirmation) | **Post** /project-participations/{projectParticipationId}/skills/{skillId}/confirmation/{confirmingPersonId} | Confirm Skill
*ProjectAPI* | [**CreateProject**](docs/ProjectAPI.md#createproject) | **Post** /projects | Create a Project in an Organization
*ProjectAPI* | [**DeleteProject**](docs/ProjectAPI.md#deleteproject) | **Delete** /projects/{projectId} | Delete a project
*ProjectAPI* | [**DeleteProjectParticipation**](docs/ProjectAPI.md#deleteprojectparticipation) | **Delete** /project-participations/{projectParticipationId} | Remove an Project from a Person
*ProjectAPI* | [**GetProject**](docs/ProjectAPI.md#getproject) | **Get** /projects/{projectId} | Get details about a Project
*ProjectAPI* | [**MoveProject**](docs/ProjectAPI.md#moveproject) | **Put** /projects/{projectId}/organizations/{organizationId} | Move a Project to an Organization
*ProjectAPI* | [**ReadProjectParticipation**](docs/ProjectAPI.md#readprojectparticipation) | **Get** /project-participations/{projectParticipationId} | Get a project participation
*ProjectAPI* | [**RemoveExecutiveOrganizationFromProject**](docs/ProjectAPI.md#removeexecutiveorganizationfromproject) | **Delete** /projects/{projectId}/executive-organizations/{organizationId} | Remove an Organization from a Project as executive organization
*ProjectAPI* | [**RemoveSkillConfirmation**](docs/ProjectAPI.md#removeskillconfirmation) | **Delete** /project-participations/{projectParticipationId}/skills/{skillId}/confirmation/{confirmingPersonId} | Remove a confirmation
*ProjectAPI* | [**SearchProjectParticipations**](docs/ProjectAPI.md#searchprojectparticipations) | **Post** /project-participations/search | Search over project participations
*ProjectAPI* | [**SearchProjects**](docs/ProjectAPI.md#searchprojects) | **Post** /projects/search | Complex search over project entities
*ProjectAPI* | [**UpdateProject**](docs/ProjectAPI.md#updateproject) | **Put** /projects/{projectId} | Update a Project
*ProjectAPI* | [**UpdateProjectParticipation**](docs/ProjectAPI.md#updateprojectparticipation) | **Put** /project-participations/{projectParticipationId} | Update a Project of a Person
*ProjectParticipationAPI* | [**ReadPersonProjectParticipation**](docs/ProjectParticipationAPI.md#readpersonprojectparticipation) | **Get** /persons/{personId}/project-participation | Get a Project Participation of a Person
*SearchAPI* | [**Search**](docs/SearchAPI.md#search) | **Post** /search | Fulltext search on all kinds of objects
*SkillAPI* | [**AddPersonInterest**](docs/SkillAPI.md#addpersoninterest) | **Post** /persons/{personId}/interests/skills/{skillId} | Add an Interest to a Person
*SkillAPI* | [**AddPersonSkillExperience**](docs/SkillAPI.md#addpersonskillexperience) | **Post** /persons/{personId}/experiences/skills/{skillId} | Add an Skill experience to a Person
*SkillAPI* | [**AddPersonSkillExperiences**](docs/SkillAPI.md#addpersonskillexperiences) | **Post** /persons/{personId}/experiences | Add an Skill experience to a Person (bulk)
*SkillAPI* | [**AddSkillConfirmation**](docs/SkillAPI.md#addskillconfirmation) | **Post** /project-participations/{projectParticipationId}/skills/{skillId}/confirmation/{confirmingPersonId} | Confirm Skill
*SkillAPI* | [**AddSkillToCertification**](docs/SkillAPI.md#addskilltocertification) | **Post** /certifications/{certificationId}/skills/{skillId} | 
*SkillAPI* | [**AddSkillToParentSkill**](docs/SkillAPI.md#addskilltoparentskill) | **Post** /skills/{skillId}/parents/{parentSkillId} | Attach a Skill to a parent Skill, returns the parent Skill
*SkillAPI* | [**CreateSkill**](docs/SkillAPI.md#createskill) | **Post** /skills | Create a Skill
*SkillAPI* | [**DeletePersonInterest**](docs/SkillAPI.md#deletepersoninterest) | **Delete** /persons/{personId}/interests/skills/{skillId} | Remove an Interest to a Person
*SkillAPI* | [**DeletePersonSkillExperience**](docs/SkillAPI.md#deletepersonskillexperience) | **Delete** /persons/{personId}/experiences/skills/{skillId} | Remove an Skill Experience to a Person
*SkillAPI* | [**DeletePersonSkillExperiences**](docs/SkillAPI.md#deletepersonskillexperiences) | **Delete** /persons/{personId}/experiences | Remove an Skill Experience to a Person
*SkillAPI* | [**DeleteSkill**](docs/SkillAPI.md#deleteskill) | **Delete** /skills/{skillId} | Delete a Skill
*SkillAPI* | [**DeleteSkillFromCertification**](docs/SkillAPI.md#deleteskillfromcertification) | **Delete** /certifications/{certificationId}/skills/{skillId} | 
*SkillAPI* | [**GetSkill**](docs/SkillAPI.md#getskill) | **Get** /skills/{skillId} | Get details for a single skill
*SkillAPI* | [**RemoveSkillConfirmation**](docs/SkillAPI.md#removeskillconfirmation) | **Delete** /project-participations/{projectParticipationId}/skills/{skillId}/confirmation/{confirmingPersonId} | Remove a confirmation
*SkillAPI* | [**RemoveSkillFromParentSkill**](docs/SkillAPI.md#removeskillfromparentskill) | **Delete** /skills/{skillId}/parents/{parentSkillId} | Detaches a Skill from parent Skill, return the parent Skill
*SkillAPI* | [**SearchSkills**](docs/SkillAPI.md#searchskills) | **Post** /skills/search | Search over skills
*SkillAPI* | [**UpdatePersonSkillExperience**](docs/SkillAPI.md#updatepersonskillexperience) | **Put** /persons/{personId}/experiences/skills/{skillId} | Edit an Skill experience to a Person
*SkillAPI* | [**UpdatePersonSkillExperiences**](docs/SkillAPI.md#updatepersonskillexperiences) | **Put** /persons/{personId}/experiences | Edit an Skill experience to a Person
*SkillAPI* | [**UpdateSkill**](docs/SkillAPI.md#updateskill) | **Put** /skills/{skillId} | Update a Skill
*SkillAPI* | [**UpdateSkillInCertification**](docs/SkillAPI.md#updateskillincertification) | **Put** /certifications/{certificationId}/skills/{skillId} | 
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
 - [Confidentiality](docs/Confidentiality.md)
 - [Country](docs/Country.md)
 - [CountryDetails](docs/CountryDetails.md)
 - [Descriptable](docs/Descriptable.md)
 - [Error](docs/Error.md)
 - [Experience](docs/Experience.md)
 - [ExperienceSkill](docs/ExperienceSkill.md)
 - [ExperienceSkillGroup](docs/ExperienceSkillGroup.md)
 - [Geolocation](docs/Geolocation.md)
 - [Industry](docs/Industry.md)
 - [IndustryDetails](docs/IndustryDetails.md)
 - [IndustrySearch](docs/IndustrySearch.md)
 - [Language](docs/Language.md)
 - [LanguageDetails](docs/LanguageDetails.md)
 - [LanguageLevel](docs/LanguageLevel.md)
 - [Level](docs/Level.md)
 - [Locateable](docs/Locateable.md)
 - [MinMax](docs/MinMax.md)
 - [MinMaxPercent](docs/MinMaxPercent.md)
 - [Nameable](docs/Nameable.md)
 - [NamedDomainModel](docs/NamedDomainModel.md)
 - [Office](docs/Office.md)
 - [OfficeDetails](docs/OfficeDetails.md)
 - [OfficeSearch](docs/OfficeSearch.md)
 - [Organization](docs/Organization.md)
 - [OrganizationDetails](docs/OrganizationDetails.md)
 - [OrganizationSearch](docs/OrganizationSearch.md)
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
 - [PagedProjectParticipation](docs/PagedProjectParticipation.md)
 - [PagedProjects](docs/PagedProjects.md)
 - [PagedSkills](docs/PagedSkills.md)
 - [Person](docs/Person.md)
 - [PersonDetails](docs/PersonDetails.md)
 - [PersonProjectParticipationDetails](docs/PersonProjectParticipationDetails.md)
 - [PersonProjectParticipationItem](docs/PersonProjectParticipationItem.md)
 - [PersonScoreDetail](docs/PersonScoreDetail.md)
 - [PersonSearch](docs/PersonSearch.md)
 - [PersonSkillFilter](docs/PersonSkillFilter.md)
 - [Profile](docs/Profile.md)
 - [ProfileDetails](docs/ProfileDetails.md)
 - [Project](docs/Project.md)
 - [ProjectDetails](docs/ProjectDetails.md)
 - [ProjectParticipation](docs/ProjectParticipation.md)
 - [ProjectParticipationCreate](docs/ProjectParticipationCreate.md)
 - [ProjectParticipationDetails](docs/ProjectParticipationDetails.md)
 - [ProjectParticipationSearch](docs/ProjectParticipationSearch.md)
 - [ProjectParticipationUpdate](docs/ProjectParticipationUpdate.md)
 - [ProjectScoreDetail](docs/ProjectScoreDetail.md)
 - [ProjectSearch](docs/ProjectSearch.md)
 - [ProjectStatus](docs/ProjectStatus.md)
 - [ProjectType](docs/ProjectType.md)
 - [ScoreResult](docs/ScoreResult.md)
 - [Search](docs/Search.md)
 - [SearchResult](docs/SearchResult.md)
 - [SearchResultItem](docs/SearchResultItem.md)
 - [Seniority](docs/Seniority.md)
 - [Skill](docs/Skill.md)
 - [SkillDetails](docs/SkillDetails.md)
 - [SkillGroup](docs/SkillGroup.md)
 - [SkillLevel](docs/SkillLevel.md)
 - [SkillLevelUpdate](docs/SkillLevelUpdate.md)
 - [SkillLink](docs/SkillLink.md)
 - [SkillLinkUpdate](docs/SkillLinkUpdate.md)
 - [SkillSearch](docs/SkillSearch.md)
 - [Status](docs/Status.md)
 - [Synonymable](docs/Synonymable.md)
 - [TimeframeFilter](docs/TimeframeFilter.md)
 - [Timeframed](docs/Timeframed.md)
 - [Version](docs/Version.md)


## Documentation For Authorization


Authentication schemes defined for the API:
### oidcScheme

### bearerScheme

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



