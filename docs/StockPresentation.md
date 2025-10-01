# StockPresentation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Symbol** | Pointer to **string** | Company symbol. | [optional] 
**Res** | Pointer to [**[]PresentationData**](PresentationData.md) | Presentation data. | [optional] 

## Methods

### NewStockPresentation

`func NewStockPresentation() *StockPresentation`

NewStockPresentation instantiates a new StockPresentation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStockPresentationWithDefaults

`func NewStockPresentationWithDefaults() *StockPresentation`

NewStockPresentationWithDefaults instantiates a new StockPresentation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSymbol

`func (o *StockPresentation) GetSymbol() string`

GetSymbol returns the Symbol field if non-nil, zero value otherwise.

### GetSymbolOk

`func (o *StockPresentation) GetSymbolOk() (*string, bool)`

GetSymbolOk returns a tuple with the Symbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSymbol

`func (o *StockPresentation) SetSymbol(v string)`

SetSymbol sets Symbol field to given value.

### HasSymbol

`func (o *StockPresentation) HasSymbol() bool`

HasSymbol returns a boolean if a field has been set.

### GetRes

`func (o *StockPresentation) GetRes() []PresentationData`

GetRes returns the Res field if non-nil, zero value otherwise.

### GetResOk

`func (o *StockPresentation) GetResOk() (*[]PresentationData, bool)`

GetResOk returns a tuple with the Res field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRes

`func (o *StockPresentation) SetRes(v []PresentationData)`

SetRes sets Res field to given value.

### HasRes

`func (o *StockPresentation) HasRes() bool`

HasRes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


