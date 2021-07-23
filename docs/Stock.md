# Stock

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** | Symbol description | [optional] 
**DisplaySymbol** | Pointer to **string** | Display symbol name. | [optional] 
**Symbol** | Pointer to **string** | Unique symbol used to identify this symbol used in &lt;code&gt;/stock/candle&lt;/code&gt; endpoint. | [optional] 
**Type** | Pointer to **string** | Security type. | [optional] 

## Methods

### NewStock

`func NewStock() *Stock`

NewStock instantiates a new Stock object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStockWithDefaults

`func NewStockWithDefaults() *Stock`

NewStockWithDefaults instantiates a new Stock object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *Stock) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Stock) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Stock) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Stock) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDisplaySymbol

`func (o *Stock) GetDisplaySymbol() string`

GetDisplaySymbol returns the DisplaySymbol field if non-nil, zero value otherwise.

### GetDisplaySymbolOk

`func (o *Stock) GetDisplaySymbolOk() (*string, bool)`

GetDisplaySymbolOk returns a tuple with the DisplaySymbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplaySymbol

`func (o *Stock) SetDisplaySymbol(v string)`

SetDisplaySymbol sets DisplaySymbol field to given value.

### HasDisplaySymbol

`func (o *Stock) HasDisplaySymbol() bool`

HasDisplaySymbol returns a boolean if a field has been set.

### GetSymbol

`func (o *Stock) GetSymbol() string`

GetSymbol returns the Symbol field if non-nil, zero value otherwise.

### GetSymbolOk

`func (o *Stock) GetSymbolOk() (*string, bool)`

GetSymbolOk returns a tuple with the Symbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSymbol

`func (o *Stock) SetSymbol(v string)`

SetSymbol sets Symbol field to given value.

### HasSymbol

`func (o *Stock) HasSymbol() bool`

HasSymbol returns a boolean if a field has been set.

### GetType

`func (o *Stock) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Stock) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Stock) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *Stock) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


