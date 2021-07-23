# ForexCandles

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**O** | Pointer to **[]float32** | List of open prices for returned candles. | [optional] 
**H** | Pointer to **[]float32** | List of high prices for returned candles. | [optional] 
**L** | Pointer to **[]float32** | List of low prices for returned candles. | [optional] 
**C** | Pointer to **[]float32** | List of close prices for returned candles. | [optional] 
**V** | Pointer to **[]float32** | List of volume data for returned candles. | [optional] 
**T** | Pointer to **[]float32** | List of timestamp for returned candles. | [optional] 
**S** | Pointer to **string** | Status of the response. This field can either be ok or no_data. | [optional] 

## Methods

### NewForexCandles

`func NewForexCandles() *ForexCandles`

NewForexCandles instantiates a new ForexCandles object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewForexCandlesWithDefaults

`func NewForexCandlesWithDefaults() *ForexCandles`

NewForexCandlesWithDefaults instantiates a new ForexCandles object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetO

`func (o *ForexCandles) GetO() []float32`

GetO returns the O field if non-nil, zero value otherwise.

### GetOOk

`func (o *ForexCandles) GetOOk() (*[]float32, bool)`

GetOOk returns a tuple with the O field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetO

`func (o *ForexCandles) SetO(v []float32)`

SetO sets O field to given value.

### HasO

`func (o *ForexCandles) HasO() bool`

HasO returns a boolean if a field has been set.

### GetH

`func (o *ForexCandles) GetH() []float32`

GetH returns the H field if non-nil, zero value otherwise.

### GetHOk

`func (o *ForexCandles) GetHOk() (*[]float32, bool)`

GetHOk returns a tuple with the H field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetH

`func (o *ForexCandles) SetH(v []float32)`

SetH sets H field to given value.

### HasH

`func (o *ForexCandles) HasH() bool`

HasH returns a boolean if a field has been set.

### GetL

`func (o *ForexCandles) GetL() []float32`

GetL returns the L field if non-nil, zero value otherwise.

### GetLOk

`func (o *ForexCandles) GetLOk() (*[]float32, bool)`

GetLOk returns a tuple with the L field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetL

`func (o *ForexCandles) SetL(v []float32)`

SetL sets L field to given value.

### HasL

`func (o *ForexCandles) HasL() bool`

HasL returns a boolean if a field has been set.

### GetC

`func (o *ForexCandles) GetC() []float32`

GetC returns the C field if non-nil, zero value otherwise.

### GetCOk

`func (o *ForexCandles) GetCOk() (*[]float32, bool)`

GetCOk returns a tuple with the C field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetC

`func (o *ForexCandles) SetC(v []float32)`

SetC sets C field to given value.

### HasC

`func (o *ForexCandles) HasC() bool`

HasC returns a boolean if a field has been set.

### GetV

`func (o *ForexCandles) GetV() []float32`

GetV returns the V field if non-nil, zero value otherwise.

### GetVOk

`func (o *ForexCandles) GetVOk() (*[]float32, bool)`

GetVOk returns a tuple with the V field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetV

`func (o *ForexCandles) SetV(v []float32)`

SetV sets V field to given value.

### HasV

`func (o *ForexCandles) HasV() bool`

HasV returns a boolean if a field has been set.

### GetT

`func (o *ForexCandles) GetT() []float32`

GetT returns the T field if non-nil, zero value otherwise.

### GetTOk

`func (o *ForexCandles) GetTOk() (*[]float32, bool)`

GetTOk returns a tuple with the T field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetT

`func (o *ForexCandles) SetT(v []float32)`

SetT sets T field to given value.

### HasT

`func (o *ForexCandles) HasT() bool`

HasT returns a boolean if a field has been set.

### GetS

`func (o *ForexCandles) GetS() string`

GetS returns the S field if non-nil, zero value otherwise.

### GetSOk

`func (o *ForexCandles) GetSOk() (*string, bool)`

GetSOk returns a tuple with the S field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetS

`func (o *ForexCandles) SetS(v string)`

SetS sets S field to given value.

### HasS

`func (o *ForexCandles) HasS() bool`

HasS returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


