# StockCandles

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**O** | Pointer to **[]float32** | List of open prices for returned candles. | [optional] 
**H** | Pointer to **[]float32** | List of high prices for returned candles. | [optional] 
**L** | Pointer to **[]float32** | List of low prices for returned candles. | [optional] 
**C** | Pointer to **[]float32** | List of close prices for returned candles. | [optional] 
**V** | Pointer to **[]float32** | List of volume data for returned candles. | [optional] 
**T** | Pointer to **[]int64** | List of timestamp for returned candles. | [optional] 
**S** | Pointer to **string** | Status of the response. This field can either be ok or no_data. | [optional] 

## Methods

### NewStockCandles

`func NewStockCandles() *StockCandles`

NewStockCandles instantiates a new StockCandles object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStockCandlesWithDefaults

`func NewStockCandlesWithDefaults() *StockCandles`

NewStockCandlesWithDefaults instantiates a new StockCandles object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetO

`func (o *StockCandles) GetO() []float32`

GetO returns the O field if non-nil, zero value otherwise.

### GetOOk

`func (o *StockCandles) GetOOk() (*[]float32, bool)`

GetOOk returns a tuple with the O field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetO

`func (o *StockCandles) SetO(v []float32)`

SetO sets O field to given value.

### HasO

`func (o *StockCandles) HasO() bool`

HasO returns a boolean if a field has been set.

### GetH

`func (o *StockCandles) GetH() []float32`

GetH returns the H field if non-nil, zero value otherwise.

### GetHOk

`func (o *StockCandles) GetHOk() (*[]float32, bool)`

GetHOk returns a tuple with the H field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetH

`func (o *StockCandles) SetH(v []float32)`

SetH sets H field to given value.

### HasH

`func (o *StockCandles) HasH() bool`

HasH returns a boolean if a field has been set.

### GetL

`func (o *StockCandles) GetL() []float32`

GetL returns the L field if non-nil, zero value otherwise.

### GetLOk

`func (o *StockCandles) GetLOk() (*[]float32, bool)`

GetLOk returns a tuple with the L field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetL

`func (o *StockCandles) SetL(v []float32)`

SetL sets L field to given value.

### HasL

`func (o *StockCandles) HasL() bool`

HasL returns a boolean if a field has been set.

### GetC

`func (o *StockCandles) GetC() []float32`

GetC returns the C field if non-nil, zero value otherwise.

### GetCOk

`func (o *StockCandles) GetCOk() (*[]float32, bool)`

GetCOk returns a tuple with the C field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetC

`func (o *StockCandles) SetC(v []float32)`

SetC sets C field to given value.

### HasC

`func (o *StockCandles) HasC() bool`

HasC returns a boolean if a field has been set.

### GetV

`func (o *StockCandles) GetV() []float32`

GetV returns the V field if non-nil, zero value otherwise.

### GetVOk

`func (o *StockCandles) GetVOk() (*[]float32, bool)`

GetVOk returns a tuple with the V field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetV

`func (o *StockCandles) SetV(v []float32)`

SetV sets V field to given value.

### HasV

`func (o *StockCandles) HasV() bool`

HasV returns a boolean if a field has been set.

### GetT

`func (o *StockCandles) GetT() []int64`

GetT returns the T field if non-nil, zero value otherwise.

### GetTOk

`func (o *StockCandles) GetTOk() (*[]int64, bool)`

GetTOk returns a tuple with the T field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetT

`func (o *StockCandles) SetT(v []int64)`

SetT sets T field to given value.

### HasT

`func (o *StockCandles) HasT() bool`

HasT returns a boolean if a field has been set.

### GetS

`func (o *StockCandles) GetS() string`

GetS returns the S field if non-nil, zero value otherwise.

### GetSOk

`func (o *StockCandles) GetSOk() (*string, bool)`

GetSOk returns a tuple with the S field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetS

`func (o *StockCandles) SetS(v string)`

SetS sets S field to given value.

### HasS

`func (o *StockCandles) HasS() bool`

HasS returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


