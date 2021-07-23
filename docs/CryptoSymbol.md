# CryptoSymbol

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** | Symbol description | [optional] 
**DisplaySymbol** | Pointer to **string** | Display symbol name. | [optional] 
**Symbol** | Pointer to **string** | Unique symbol used to identify this symbol used in &lt;code&gt;/crypto/candle&lt;/code&gt; endpoint. | [optional] 

## Methods

### NewCryptoSymbol

`func NewCryptoSymbol() *CryptoSymbol`

NewCryptoSymbol instantiates a new CryptoSymbol object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCryptoSymbolWithDefaults

`func NewCryptoSymbolWithDefaults() *CryptoSymbol`

NewCryptoSymbolWithDefaults instantiates a new CryptoSymbol object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *CryptoSymbol) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CryptoSymbol) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CryptoSymbol) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CryptoSymbol) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDisplaySymbol

`func (o *CryptoSymbol) GetDisplaySymbol() string`

GetDisplaySymbol returns the DisplaySymbol field if non-nil, zero value otherwise.

### GetDisplaySymbolOk

`func (o *CryptoSymbol) GetDisplaySymbolOk() (*string, bool)`

GetDisplaySymbolOk returns a tuple with the DisplaySymbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplaySymbol

`func (o *CryptoSymbol) SetDisplaySymbol(v string)`

SetDisplaySymbol sets DisplaySymbol field to given value.

### HasDisplaySymbol

`func (o *CryptoSymbol) HasDisplaySymbol() bool`

HasDisplaySymbol returns a boolean if a field has been set.

### GetSymbol

`func (o *CryptoSymbol) GetSymbol() string`

GetSymbol returns the Symbol field if non-nil, zero value otherwise.

### GetSymbolOk

`func (o *CryptoSymbol) GetSymbolOk() (*string, bool)`

GetSymbolOk returns a tuple with the Symbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSymbol

`func (o *CryptoSymbol) SetSymbol(v string)`

SetSymbol sets Symbol field to given value.

### HasSymbol

`func (o *CryptoSymbol) HasSymbol() bool`

HasSymbol returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


