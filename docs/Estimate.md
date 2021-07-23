# Estimate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RevenueAvg** | Pointer to **int64** | Average revenue estimates including Finnhub&#39;s proprietary estimates. | [optional] 
**RevenueHigh** | Pointer to **int64** | Highest estimate. | [optional] 
**RevenueLow** | Pointer to **int64** | Lowest estimate. | [optional] 
**NumberAnalysts** | Pointer to **int64** | Number of Analysts. | [optional] 
**Period** | Pointer to **string** | Period. | [optional] 

## Methods

### NewEstimate

`func NewEstimate() *Estimate`

NewEstimate instantiates a new Estimate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEstimateWithDefaults

`func NewEstimateWithDefaults() *Estimate`

NewEstimateWithDefaults instantiates a new Estimate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRevenueAvg

`func (o *Estimate) GetRevenueAvg() int64`

GetRevenueAvg returns the RevenueAvg field if non-nil, zero value otherwise.

### GetRevenueAvgOk

`func (o *Estimate) GetRevenueAvgOk() (*int64, bool)`

GetRevenueAvgOk returns a tuple with the RevenueAvg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevenueAvg

`func (o *Estimate) SetRevenueAvg(v int64)`

SetRevenueAvg sets RevenueAvg field to given value.

### HasRevenueAvg

`func (o *Estimate) HasRevenueAvg() bool`

HasRevenueAvg returns a boolean if a field has been set.

### GetRevenueHigh

`func (o *Estimate) GetRevenueHigh() int64`

GetRevenueHigh returns the RevenueHigh field if non-nil, zero value otherwise.

### GetRevenueHighOk

`func (o *Estimate) GetRevenueHighOk() (*int64, bool)`

GetRevenueHighOk returns a tuple with the RevenueHigh field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevenueHigh

`func (o *Estimate) SetRevenueHigh(v int64)`

SetRevenueHigh sets RevenueHigh field to given value.

### HasRevenueHigh

`func (o *Estimate) HasRevenueHigh() bool`

HasRevenueHigh returns a boolean if a field has been set.

### GetRevenueLow

`func (o *Estimate) GetRevenueLow() int64`

GetRevenueLow returns the RevenueLow field if non-nil, zero value otherwise.

### GetRevenueLowOk

`func (o *Estimate) GetRevenueLowOk() (*int64, bool)`

GetRevenueLowOk returns a tuple with the RevenueLow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevenueLow

`func (o *Estimate) SetRevenueLow(v int64)`

SetRevenueLow sets RevenueLow field to given value.

### HasRevenueLow

`func (o *Estimate) HasRevenueLow() bool`

HasRevenueLow returns a boolean if a field has been set.

### GetNumberAnalysts

`func (o *Estimate) GetNumberAnalysts() int64`

GetNumberAnalysts returns the NumberAnalysts field if non-nil, zero value otherwise.

### GetNumberAnalystsOk

`func (o *Estimate) GetNumberAnalystsOk() (*int64, bool)`

GetNumberAnalystsOk returns a tuple with the NumberAnalysts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumberAnalysts

`func (o *Estimate) SetNumberAnalysts(v int64)`

SetNumberAnalysts sets NumberAnalysts field to given value.

### HasNumberAnalysts

`func (o *Estimate) HasNumberAnalysts() bool`

HasNumberAnalysts returns a boolean if a field has been set.

### GetPeriod

`func (o *Estimate) GetPeriod() string`

GetPeriod returns the Period field if non-nil, zero value otherwise.

### GetPeriodOk

`func (o *Estimate) GetPeriodOk() (*string, bool)`

GetPeriodOk returns a tuple with the Period field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeriod

`func (o *Estimate) SetPeriod(v string)`

SetPeriod sets Period field to given value.

### HasPeriod

`func (o *Estimate) HasPeriod() bool`

HasPeriod returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


