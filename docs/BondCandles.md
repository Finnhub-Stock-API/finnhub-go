# BondCandles

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**C** | Pointer to **[]float32** | List of close prices for returned candles. | [optional] 
**T** | Pointer to **[]int64** | List of timestamp for returned candles. | [optional] 
**S** | Pointer to **string** | Status of the response. This field can either be ok or no_data. | [optional] 

## Methods

### NewBondCandles

`func NewBondCandles() *BondCandles`

NewBondCandles instantiates a new BondCandles object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBondCandlesWithDefaults

`func NewBondCandlesWithDefaults() *BondCandles`

NewBondCandlesWithDefaults instantiates a new BondCandles object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetC

`func (o *BondCandles) GetC() []float32`

GetC returns the C field if non-nil, zero value otherwise.

### GetCOk

`func (o *BondCandles) GetCOk() (*[]float32, bool)`

GetCOk returns a tuple with the C field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetC

`func (o *BondCandles) SetC(v []float32)`

SetC sets C field to given value.

### HasC

`func (o *BondCandles) HasC() bool`

HasC returns a boolean if a field has been set.

### GetT

`func (o *BondCandles) GetT() []int64`

GetT returns the T field if non-nil, zero value otherwise.

### GetTOk

`func (o *BondCandles) GetTOk() (*[]int64, bool)`

GetTOk returns a tuple with the T field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetT

`func (o *BondCandles) SetT(v []int64)`

SetT sets T field to given value.

### HasT

`func (o *BondCandles) HasT() bool`

HasT returns a boolean if a field has been set.

### GetS

`func (o *BondCandles) GetS() string`

GetS returns the S field if non-nil, zero value otherwise.

### GetSOk

`func (o *BondCandles) GetSOk() (*string, bool)`

GetSOk returns a tuple with the S field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetS

`func (o *BondCandles) SetS(v string)`

SetS sets S field to given value.

### HasS

`func (o *BondCandles) HasS() bool`

HasS returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


