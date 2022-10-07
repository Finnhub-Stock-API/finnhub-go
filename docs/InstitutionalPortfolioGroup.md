# InstitutionalPortfolioGroup

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ReportDate** | Pointer to **string** | Report date. | [optional] 
**FilingDate** | Pointer to **string** | Filing date. | [optional] 
**Portfolio** | Pointer to [**[]InstitutionalPortfolioInfo**](InstitutionalPortfolioInfo.md) | Array of positions. | [optional] 

## Methods

### NewInstitutionalPortfolioGroup

`func NewInstitutionalPortfolioGroup() *InstitutionalPortfolioGroup`

NewInstitutionalPortfolioGroup instantiates a new InstitutionalPortfolioGroup object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInstitutionalPortfolioGroupWithDefaults

`func NewInstitutionalPortfolioGroupWithDefaults() *InstitutionalPortfolioGroup`

NewInstitutionalPortfolioGroupWithDefaults instantiates a new InstitutionalPortfolioGroup object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetReportDate

`func (o *InstitutionalPortfolioGroup) GetReportDate() string`

GetReportDate returns the ReportDate field if non-nil, zero value otherwise.

### GetReportDateOk

`func (o *InstitutionalPortfolioGroup) GetReportDateOk() (*string, bool)`

GetReportDateOk returns a tuple with the ReportDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReportDate

`func (o *InstitutionalPortfolioGroup) SetReportDate(v string)`

SetReportDate sets ReportDate field to given value.

### HasReportDate

`func (o *InstitutionalPortfolioGroup) HasReportDate() bool`

HasReportDate returns a boolean if a field has been set.

### GetFilingDate

`func (o *InstitutionalPortfolioGroup) GetFilingDate() string`

GetFilingDate returns the FilingDate field if non-nil, zero value otherwise.

### GetFilingDateOk

`func (o *InstitutionalPortfolioGroup) GetFilingDateOk() (*string, bool)`

GetFilingDateOk returns a tuple with the FilingDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilingDate

`func (o *InstitutionalPortfolioGroup) SetFilingDate(v string)`

SetFilingDate sets FilingDate field to given value.

### HasFilingDate

`func (o *InstitutionalPortfolioGroup) HasFilingDate() bool`

HasFilingDate returns a boolean if a field has been set.

### GetPortfolio

`func (o *InstitutionalPortfolioGroup) GetPortfolio() []InstitutionalPortfolioInfo`

GetPortfolio returns the Portfolio field if non-nil, zero value otherwise.

### GetPortfolioOk

`func (o *InstitutionalPortfolioGroup) GetPortfolioOk() (*[]InstitutionalPortfolioInfo, bool)`

GetPortfolioOk returns a tuple with the Portfolio field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortfolio

`func (o *InstitutionalPortfolioGroup) SetPortfolio(v []InstitutionalPortfolioInfo)`

SetPortfolio sets Portfolio field to given value.

### HasPortfolio

`func (o *InstitutionalPortfolioGroup) HasPortfolio() bool`

HasPortfolio returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


