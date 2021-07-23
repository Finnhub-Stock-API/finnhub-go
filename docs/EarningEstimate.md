# EarningEstimate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EpsAvg** | Pointer to **float32** | Average EPS estimates including Finnhub&#39;s proprietary estimates. | [optional] 
**EpsHigh** | Pointer to **float32** | Highest estimate. | [optional] 
**EpsLow** | Pointer to **float32** | Lowest estimate. | [optional] 
**NumberAnalysts** | Pointer to **int64** | Number of Analysts. | [optional] 
**Period** | Pointer to **string** | Period. | [optional] 

## Methods

### NewEarningEstimate

`func NewEarningEstimate() *EarningEstimate`

NewEarningEstimate instantiates a new EarningEstimate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEarningEstimateWithDefaults

`func NewEarningEstimateWithDefaults() *EarningEstimate`

NewEarningEstimateWithDefaults instantiates a new EarningEstimate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEpsAvg

`func (o *EarningEstimate) GetEpsAvg() float32`

GetEpsAvg returns the EpsAvg field if non-nil, zero value otherwise.

### GetEpsAvgOk

`func (o *EarningEstimate) GetEpsAvgOk() (*float32, bool)`

GetEpsAvgOk returns a tuple with the EpsAvg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEpsAvg

`func (o *EarningEstimate) SetEpsAvg(v float32)`

SetEpsAvg sets EpsAvg field to given value.

### HasEpsAvg

`func (o *EarningEstimate) HasEpsAvg() bool`

HasEpsAvg returns a boolean if a field has been set.

### GetEpsHigh

`func (o *EarningEstimate) GetEpsHigh() float32`

GetEpsHigh returns the EpsHigh field if non-nil, zero value otherwise.

### GetEpsHighOk

`func (o *EarningEstimate) GetEpsHighOk() (*float32, bool)`

GetEpsHighOk returns a tuple with the EpsHigh field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEpsHigh

`func (o *EarningEstimate) SetEpsHigh(v float32)`

SetEpsHigh sets EpsHigh field to given value.

### HasEpsHigh

`func (o *EarningEstimate) HasEpsHigh() bool`

HasEpsHigh returns a boolean if a field has been set.

### GetEpsLow

`func (o *EarningEstimate) GetEpsLow() float32`

GetEpsLow returns the EpsLow field if non-nil, zero value otherwise.

### GetEpsLowOk

`func (o *EarningEstimate) GetEpsLowOk() (*float32, bool)`

GetEpsLowOk returns a tuple with the EpsLow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEpsLow

`func (o *EarningEstimate) SetEpsLow(v float32)`

SetEpsLow sets EpsLow field to given value.

### HasEpsLow

`func (o *EarningEstimate) HasEpsLow() bool`

HasEpsLow returns a boolean if a field has been set.

### GetNumberAnalysts

`func (o *EarningEstimate) GetNumberAnalysts() int64`

GetNumberAnalysts returns the NumberAnalysts field if non-nil, zero value otherwise.

### GetNumberAnalystsOk

`func (o *EarningEstimate) GetNumberAnalystsOk() (*int64, bool)`

GetNumberAnalystsOk returns a tuple with the NumberAnalysts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumberAnalysts

`func (o *EarningEstimate) SetNumberAnalysts(v int64)`

SetNumberAnalysts sets NumberAnalysts field to given value.

### HasNumberAnalysts

`func (o *EarningEstimate) HasNumberAnalysts() bool`

HasNumberAnalysts returns a boolean if a field has been set.

### GetPeriod

`func (o *EarningEstimate) GetPeriod() string`

GetPeriod returns the Period field if non-nil, zero value otherwise.

### GetPeriodOk

`func (o *EarningEstimate) GetPeriodOk() (*string, bool)`

GetPeriodOk returns a tuple with the Period field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeriod

`func (o *EarningEstimate) SetPeriod(v string)`

SetPeriod sets Period field to given value.

### HasPeriod

`func (o *EarningEstimate) HasPeriod() bool`

HasPeriod returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


