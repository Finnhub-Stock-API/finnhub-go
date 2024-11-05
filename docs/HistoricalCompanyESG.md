# HistoricalCompanyESG

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Symbol** | Pointer to **string** | symbol | [optional] 
**Data** | Pointer to [**[]CompanyESG2**](CompanyESG2.md) | Historical ESG data points. | [optional] 

## Methods

### NewHistoricalCompanyESG

`func NewHistoricalCompanyESG() *HistoricalCompanyESG`

NewHistoricalCompanyESG instantiates a new HistoricalCompanyESG object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHistoricalCompanyESGWithDefaults

`func NewHistoricalCompanyESGWithDefaults() *HistoricalCompanyESG`

NewHistoricalCompanyESGWithDefaults instantiates a new HistoricalCompanyESG object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSymbol

`func (o *HistoricalCompanyESG) GetSymbol() string`

GetSymbol returns the Symbol field if non-nil, zero value otherwise.

### GetSymbolOk

`func (o *HistoricalCompanyESG) GetSymbolOk() (*string, bool)`

GetSymbolOk returns a tuple with the Symbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSymbol

`func (o *HistoricalCompanyESG) SetSymbol(v string)`

SetSymbol sets Symbol field to given value.

### HasSymbol

`func (o *HistoricalCompanyESG) HasSymbol() bool`

HasSymbol returns a boolean if a field has been set.

### GetData

`func (o *HistoricalCompanyESG) GetData() []CompanyESG2`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *HistoricalCompanyESG) GetDataOk() (*[]CompanyESG2, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *HistoricalCompanyESG) SetData(v []CompanyESG2)`

SetData sets Data field to given value.

### HasData

`func (o *HistoricalCompanyESG) HasData() bool`

HasData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


