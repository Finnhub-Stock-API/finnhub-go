# PriceMetrics

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Symbol** | Pointer to **string** | Symbol of the company. | [optional] 
**Data** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewPriceMetrics

`func NewPriceMetrics() *PriceMetrics`

NewPriceMetrics instantiates a new PriceMetrics object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPriceMetricsWithDefaults

`func NewPriceMetricsWithDefaults() *PriceMetrics`

NewPriceMetricsWithDefaults instantiates a new PriceMetrics object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSymbol

`func (o *PriceMetrics) GetSymbol() string`

GetSymbol returns the Symbol field if non-nil, zero value otherwise.

### GetSymbolOk

`func (o *PriceMetrics) GetSymbolOk() (*string, bool)`

GetSymbolOk returns a tuple with the Symbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSymbol

`func (o *PriceMetrics) SetSymbol(v string)`

SetSymbol sets Symbol field to given value.

### HasSymbol

`func (o *PriceMetrics) HasSymbol() bool`

HasSymbol returns a boolean if a field has been set.

### GetData

`func (o *PriceMetrics) GetData() map[string]interface{}`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *PriceMetrics) GetDataOk() (*map[string]interface{}, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *PriceMetrics) SetData(v map[string]interface{})`

SetData sets Data field to given value.

### HasData

`func (o *PriceMetrics) HasData() bool`

HasData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


