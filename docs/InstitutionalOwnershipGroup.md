# InstitutionalOwnershipGroup

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ReportDate** | Pointer to **string** | Report date. | [optional] 
**Ownership** | Pointer to [**[]InstitutionalOwnershipInfo**](InstitutionalOwnershipInfo.md) | Array of institutional investors. | [optional] 

## Methods

### NewInstitutionalOwnershipGroup

`func NewInstitutionalOwnershipGroup() *InstitutionalOwnershipGroup`

NewInstitutionalOwnershipGroup instantiates a new InstitutionalOwnershipGroup object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInstitutionalOwnershipGroupWithDefaults

`func NewInstitutionalOwnershipGroupWithDefaults() *InstitutionalOwnershipGroup`

NewInstitutionalOwnershipGroupWithDefaults instantiates a new InstitutionalOwnershipGroup object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetReportDate

`func (o *InstitutionalOwnershipGroup) GetReportDate() string`

GetReportDate returns the ReportDate field if non-nil, zero value otherwise.

### GetReportDateOk

`func (o *InstitutionalOwnershipGroup) GetReportDateOk() (*string, bool)`

GetReportDateOk returns a tuple with the ReportDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReportDate

`func (o *InstitutionalOwnershipGroup) SetReportDate(v string)`

SetReportDate sets ReportDate field to given value.

### HasReportDate

`func (o *InstitutionalOwnershipGroup) HasReportDate() bool`

HasReportDate returns a boolean if a field has been set.

### GetOwnership

`func (o *InstitutionalOwnershipGroup) GetOwnership() []InstitutionalOwnershipInfo`

GetOwnership returns the Ownership field if non-nil, zero value otherwise.

### GetOwnershipOk

`func (o *InstitutionalOwnershipGroup) GetOwnershipOk() (*[]InstitutionalOwnershipInfo, bool)`

GetOwnershipOk returns a tuple with the Ownership field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnership

`func (o *InstitutionalOwnershipGroup) SetOwnership(v []InstitutionalOwnershipInfo)`

SetOwnership sets Ownership field to given value.

### HasOwnership

`func (o *InstitutionalOwnershipGroup) HasOwnership() bool`

HasOwnership returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


