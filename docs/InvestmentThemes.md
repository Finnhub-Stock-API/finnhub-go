# InvestmentThemes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Theme** | Pointer to **string** | Investment theme | [optional] 
**Data** | Pointer to [**[]InvestmentThemePortfolio**](InvestmentThemePortfolio.md) | Investment theme portfolio. | [optional] 

## Methods

### NewInvestmentThemes

`func NewInvestmentThemes() *InvestmentThemes`

NewInvestmentThemes instantiates a new InvestmentThemes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInvestmentThemesWithDefaults

`func NewInvestmentThemesWithDefaults() *InvestmentThemes`

NewInvestmentThemesWithDefaults instantiates a new InvestmentThemes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTheme

`func (o *InvestmentThemes) GetTheme() string`

GetTheme returns the Theme field if non-nil, zero value otherwise.

### GetThemeOk

`func (o *InvestmentThemes) GetThemeOk() (*string, bool)`

GetThemeOk returns a tuple with the Theme field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTheme

`func (o *InvestmentThemes) SetTheme(v string)`

SetTheme sets Theme field to given value.

### HasTheme

`func (o *InvestmentThemes) HasTheme() bool`

HasTheme returns a boolean if a field has been set.

### GetData

`func (o *InvestmentThemes) GetData() []InvestmentThemePortfolio`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *InvestmentThemes) GetDataOk() (*[]InvestmentThemePortfolio, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *InvestmentThemes) SetData(v []InvestmentThemePortfolio)`

SetData sets Data field to given value.

### HasData

`func (o *InvestmentThemes) HasData() bool`

HasData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


