# CongressionalTrading

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Symbol** | Pointer to **string** | Symbol of the company. | [optional] 
**Data** | Pointer to [**[]CongressionalTransaction**](CongressionalTransaction.md) | Array of stock trades. | [optional] 

## Methods

### NewCongressionalTrading

`func NewCongressionalTrading() *CongressionalTrading`

NewCongressionalTrading instantiates a new CongressionalTrading object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCongressionalTradingWithDefaults

`func NewCongressionalTradingWithDefaults() *CongressionalTrading`

NewCongressionalTradingWithDefaults instantiates a new CongressionalTrading object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSymbol

`func (o *CongressionalTrading) GetSymbol() string`

GetSymbol returns the Symbol field if non-nil, zero value otherwise.

### GetSymbolOk

`func (o *CongressionalTrading) GetSymbolOk() (*string, bool)`

GetSymbolOk returns a tuple with the Symbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSymbol

`func (o *CongressionalTrading) SetSymbol(v string)`

SetSymbol sets Symbol field to given value.

### HasSymbol

`func (o *CongressionalTrading) HasSymbol() bool`

HasSymbol returns a boolean if a field has been set.

### GetData

`func (o *CongressionalTrading) GetData() []CongressionalTransaction`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *CongressionalTrading) GetDataOk() (*[]CongressionalTransaction, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *CongressionalTrading) SetData(v []CongressionalTransaction)`

SetData sets Data field to given value.

### HasData

`func (o *CongressionalTrading) HasData() bool`

HasData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


