# InstitutionalPortfolio

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | Investor&#39;s name. | [optional] 
**Cik** | Pointer to **string** | CIK. | [optional] 
**Data** | Pointer to [**[]InstitutionalPortfolioGroup**](InstitutionalPortfolioGroup.md) | Array of positions. | [optional] 

## Methods

### NewInstitutionalPortfolio

`func NewInstitutionalPortfolio() *InstitutionalPortfolio`

NewInstitutionalPortfolio instantiates a new InstitutionalPortfolio object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInstitutionalPortfolioWithDefaults

`func NewInstitutionalPortfolioWithDefaults() *InstitutionalPortfolio`

NewInstitutionalPortfolioWithDefaults instantiates a new InstitutionalPortfolio object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *InstitutionalPortfolio) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *InstitutionalPortfolio) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *InstitutionalPortfolio) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *InstitutionalPortfolio) HasName() bool`

HasName returns a boolean if a field has been set.

### GetCik

`func (o *InstitutionalPortfolio) GetCik() string`

GetCik returns the Cik field if non-nil, zero value otherwise.

### GetCikOk

`func (o *InstitutionalPortfolio) GetCikOk() (*string, bool)`

GetCikOk returns a tuple with the Cik field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCik

`func (o *InstitutionalPortfolio) SetCik(v string)`

SetCik sets Cik field to given value.

### HasCik

`func (o *InstitutionalPortfolio) HasCik() bool`

HasCik returns a boolean if a field has been set.

### GetData

`func (o *InstitutionalPortfolio) GetData() []InstitutionalPortfolioGroup`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *InstitutionalPortfolio) GetDataOk() (*[]InstitutionalPortfolioGroup, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *InstitutionalPortfolio) SetData(v []InstitutionalPortfolioGroup)`

SetData sets Data field to given value.

### HasData

`func (o *InstitutionalPortfolio) HasData() bool`

HasData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


