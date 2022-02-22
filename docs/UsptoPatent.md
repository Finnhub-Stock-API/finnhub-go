# UsptoPatent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApplicationNumber** | Pointer to **string** | Application Number. | [optional] 
**CompanyFilingName** | Pointer to **[]string** | Array of companies&#39; name on the patent. | [optional] 
**FilingDate** | Pointer to **string** | Filing date. | [optional] 
**Description** | Pointer to **string** | Description. | [optional] 
**FilingStatus** | Pointer to **string** | Filing status. | [optional] 
**PatentNumber** | Pointer to **string** | Patent number. | [optional] 
**PublicationDate** | Pointer to **string** | Publication date. | [optional] 
**PatentType** | Pointer to **string** | Patent&#39;s type. | [optional] 
**Url** | Pointer to **string** | URL of the original article. | [optional] 

## Methods

### NewUsptoPatent

`func NewUsptoPatent() *UsptoPatent`

NewUsptoPatent instantiates a new UsptoPatent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUsptoPatentWithDefaults

`func NewUsptoPatentWithDefaults() *UsptoPatent`

NewUsptoPatentWithDefaults instantiates a new UsptoPatent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApplicationNumber

`func (o *UsptoPatent) GetApplicationNumber() string`

GetApplicationNumber returns the ApplicationNumber field if non-nil, zero value otherwise.

### GetApplicationNumberOk

`func (o *UsptoPatent) GetApplicationNumberOk() (*string, bool)`

GetApplicationNumberOk returns a tuple with the ApplicationNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationNumber

`func (o *UsptoPatent) SetApplicationNumber(v string)`

SetApplicationNumber sets ApplicationNumber field to given value.

### HasApplicationNumber

`func (o *UsptoPatent) HasApplicationNumber() bool`

HasApplicationNumber returns a boolean if a field has been set.

### GetCompanyFilingName

`func (o *UsptoPatent) GetCompanyFilingName() []string`

GetCompanyFilingName returns the CompanyFilingName field if non-nil, zero value otherwise.

### GetCompanyFilingNameOk

`func (o *UsptoPatent) GetCompanyFilingNameOk() (*[]string, bool)`

GetCompanyFilingNameOk returns a tuple with the CompanyFilingName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyFilingName

`func (o *UsptoPatent) SetCompanyFilingName(v []string)`

SetCompanyFilingName sets CompanyFilingName field to given value.

### HasCompanyFilingName

`func (o *UsptoPatent) HasCompanyFilingName() bool`

HasCompanyFilingName returns a boolean if a field has been set.

### GetFilingDate

`func (o *UsptoPatent) GetFilingDate() string`

GetFilingDate returns the FilingDate field if non-nil, zero value otherwise.

### GetFilingDateOk

`func (o *UsptoPatent) GetFilingDateOk() (*string, bool)`

GetFilingDateOk returns a tuple with the FilingDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilingDate

`func (o *UsptoPatent) SetFilingDate(v string)`

SetFilingDate sets FilingDate field to given value.

### HasFilingDate

`func (o *UsptoPatent) HasFilingDate() bool`

HasFilingDate returns a boolean if a field has been set.

### GetDescription

`func (o *UsptoPatent) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *UsptoPatent) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *UsptoPatent) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *UsptoPatent) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetFilingStatus

`func (o *UsptoPatent) GetFilingStatus() string`

GetFilingStatus returns the FilingStatus field if non-nil, zero value otherwise.

### GetFilingStatusOk

`func (o *UsptoPatent) GetFilingStatusOk() (*string, bool)`

GetFilingStatusOk returns a tuple with the FilingStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilingStatus

`func (o *UsptoPatent) SetFilingStatus(v string)`

SetFilingStatus sets FilingStatus field to given value.

### HasFilingStatus

`func (o *UsptoPatent) HasFilingStatus() bool`

HasFilingStatus returns a boolean if a field has been set.

### GetPatentNumber

`func (o *UsptoPatent) GetPatentNumber() string`

GetPatentNumber returns the PatentNumber field if non-nil, zero value otherwise.

### GetPatentNumberOk

`func (o *UsptoPatent) GetPatentNumberOk() (*string, bool)`

GetPatentNumberOk returns a tuple with the PatentNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPatentNumber

`func (o *UsptoPatent) SetPatentNumber(v string)`

SetPatentNumber sets PatentNumber field to given value.

### HasPatentNumber

`func (o *UsptoPatent) HasPatentNumber() bool`

HasPatentNumber returns a boolean if a field has been set.

### GetPublicationDate

`func (o *UsptoPatent) GetPublicationDate() string`

GetPublicationDate returns the PublicationDate field if non-nil, zero value otherwise.

### GetPublicationDateOk

`func (o *UsptoPatent) GetPublicationDateOk() (*string, bool)`

GetPublicationDateOk returns a tuple with the PublicationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicationDate

`func (o *UsptoPatent) SetPublicationDate(v string)`

SetPublicationDate sets PublicationDate field to given value.

### HasPublicationDate

`func (o *UsptoPatent) HasPublicationDate() bool`

HasPublicationDate returns a boolean if a field has been set.

### GetPatentType

`func (o *UsptoPatent) GetPatentType() string`

GetPatentType returns the PatentType field if non-nil, zero value otherwise.

### GetPatentTypeOk

`func (o *UsptoPatent) GetPatentTypeOk() (*string, bool)`

GetPatentTypeOk returns a tuple with the PatentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPatentType

`func (o *UsptoPatent) SetPatentType(v string)`

SetPatentType sets PatentType field to given value.

### HasPatentType

`func (o *UsptoPatent) HasPatentType() bool`

HasPatentType returns a boolean if a field has been set.

### GetUrl

`func (o *UsptoPatent) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *UsptoPatent) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *UsptoPatent) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *UsptoPatent) HasUrl() bool`

HasUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


