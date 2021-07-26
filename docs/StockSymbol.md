# StockSymbol

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** | Symbol description | [optional] 
**DisplaySymbol** | Pointer to **string** | Display symbol name. | [optional] 
**Symbol** | Pointer to **string** | Unique symbol used to identify this symbol used in &lt;code&gt;/stock/candle&lt;/code&gt; endpoint. | [optional] 
**Type** | Pointer to **string** | Security type. | [optional] 
**Mic** | Pointer to **string** | Primary exchange&#39;s MIC. | [optional] 
**Figi** | Pointer to **string** | FIGI identifier. | [optional] 
**Currency** | Pointer to **string** | Price&#39;s currency. This might be different from the reporting currency of fundamental data. | [optional] 

## Methods

### NewStockSymbol

`func NewStockSymbol() *StockSymbol`

NewStockSymbol instantiates a new StockSymbol object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStockSymbolWithDefaults

`func NewStockSymbolWithDefaults() *StockSymbol`

NewStockSymbolWithDefaults instantiates a new StockSymbol object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *StockSymbol) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *StockSymbol) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *StockSymbol) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *StockSymbol) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDisplaySymbol

`func (o *StockSymbol) GetDisplaySymbol() string`

GetDisplaySymbol returns the DisplaySymbol field if non-nil, zero value otherwise.

### GetDisplaySymbolOk

`func (o *StockSymbol) GetDisplaySymbolOk() (*string, bool)`

GetDisplaySymbolOk returns a tuple with the DisplaySymbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplaySymbol

`func (o *StockSymbol) SetDisplaySymbol(v string)`

SetDisplaySymbol sets DisplaySymbol field to given value.

### HasDisplaySymbol

`func (o *StockSymbol) HasDisplaySymbol() bool`

HasDisplaySymbol returns a boolean if a field has been set.

### GetSymbol

`func (o *StockSymbol) GetSymbol() string`

GetSymbol returns the Symbol field if non-nil, zero value otherwise.

### GetSymbolOk

`func (o *StockSymbol) GetSymbolOk() (*string, bool)`

GetSymbolOk returns a tuple with the Symbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSymbol

`func (o *StockSymbol) SetSymbol(v string)`

SetSymbol sets Symbol field to given value.

### HasSymbol

`func (o *StockSymbol) HasSymbol() bool`

HasSymbol returns a boolean if a field has been set.

### GetType

`func (o *StockSymbol) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *StockSymbol) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *StockSymbol) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *StockSymbol) HasType() bool`

HasType returns a boolean if a field has been set.

### GetMic

`func (o *StockSymbol) GetMic() string`

GetMic returns the Mic field if non-nil, zero value otherwise.

### GetMicOk

`func (o *StockSymbol) GetMicOk() (*string, bool)`

GetMicOk returns a tuple with the Mic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMic

`func (o *StockSymbol) SetMic(v string)`

SetMic sets Mic field to given value.

### HasMic

`func (o *StockSymbol) HasMic() bool`

HasMic returns a boolean if a field has been set.

### GetFigi

`func (o *StockSymbol) GetFigi() string`

GetFigi returns the Figi field if non-nil, zero value otherwise.

### GetFigiOk

`func (o *StockSymbol) GetFigiOk() (*string, bool)`

GetFigiOk returns a tuple with the Figi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFigi

`func (o *StockSymbol) SetFigi(v string)`

SetFigi sets Figi field to given value.

### HasFigi

`func (o *StockSymbol) HasFigi() bool`

HasFigi returns a boolean if a field has been set.

### GetCurrency

`func (o *StockSymbol) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *StockSymbol) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *StockSymbol) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *StockSymbol) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


