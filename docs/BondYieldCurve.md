# BondYieldCurve

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**[]BondYieldCurveInfo**](BondYieldCurveInfo.md) | Array of data. | [optional] 
**Code** | Pointer to **string** | Bond&#39;s code | [optional] 

## Methods

### NewBondYieldCurve

`func NewBondYieldCurve() *BondYieldCurve`

NewBondYieldCurve instantiates a new BondYieldCurve object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBondYieldCurveWithDefaults

`func NewBondYieldCurveWithDefaults() *BondYieldCurve`

NewBondYieldCurveWithDefaults instantiates a new BondYieldCurve object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *BondYieldCurve) GetData() []BondYieldCurveInfo`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *BondYieldCurve) GetDataOk() (*[]BondYieldCurveInfo, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *BondYieldCurve) SetData(v []BondYieldCurveInfo)`

SetData sets Data field to given value.

### HasData

`func (o *BondYieldCurve) HasData() bool`

HasData returns a boolean if a field has been set.

### GetCode

`func (o *BondYieldCurve) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *BondYieldCurve) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *BondYieldCurve) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *BondYieldCurve) HasCode() bool`

HasCode returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


