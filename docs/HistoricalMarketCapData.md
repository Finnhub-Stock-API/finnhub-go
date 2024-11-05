# HistoricalMarketCapData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**[]MarketCapData**](MarketCapData.md) | Array of market data. | [optional] 
**Symbol** | Pointer to **string** | Symbol | [optional] 
**Currency** | Pointer to **string** | Currency | [optional] 

## Methods

### NewHistoricalMarketCapData

`func NewHistoricalMarketCapData() *HistoricalMarketCapData`

NewHistoricalMarketCapData instantiates a new HistoricalMarketCapData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHistoricalMarketCapDataWithDefaults

`func NewHistoricalMarketCapDataWithDefaults() *HistoricalMarketCapData`

NewHistoricalMarketCapDataWithDefaults instantiates a new HistoricalMarketCapData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *HistoricalMarketCapData) GetData() []MarketCapData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *HistoricalMarketCapData) GetDataOk() (*[]MarketCapData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *HistoricalMarketCapData) SetData(v []MarketCapData)`

SetData sets Data field to given value.

### HasData

`func (o *HistoricalMarketCapData) HasData() bool`

HasData returns a boolean if a field has been set.

### GetSymbol

`func (o *HistoricalMarketCapData) GetSymbol() string`

GetSymbol returns the Symbol field if non-nil, zero value otherwise.

### GetSymbolOk

`func (o *HistoricalMarketCapData) GetSymbolOk() (*string, bool)`

GetSymbolOk returns a tuple with the Symbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSymbol

`func (o *HistoricalMarketCapData) SetSymbol(v string)`

SetSymbol sets Symbol field to given value.

### HasSymbol

`func (o *HistoricalMarketCapData) HasSymbol() bool`

HasSymbol returns a boolean if a field has been set.

### GetCurrency

`func (o *HistoricalMarketCapData) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *HistoricalMarketCapData) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *HistoricalMarketCapData) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *HistoricalMarketCapData) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


