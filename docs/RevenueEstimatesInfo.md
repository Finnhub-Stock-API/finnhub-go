# RevenueEstimatesInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RevenueAvg** | Pointer to **float32** | Average revenue estimates including Finnhub&#39;s proprietary estimates. | [optional] 
**RevenueHigh** | Pointer to **float32** | Highest estimate. | [optional] 
**RevenueLow** | Pointer to **float32** | Lowest estimate. | [optional] 
**NumberAnalysts** | Pointer to **int64** | Number of Analysts. | [optional] 
**Period** | Pointer to **string** | Period. | [optional] 
**Year** | Pointer to **int64** | Fiscal year. | [optional] 
**Quarter** | Pointer to **int64** | Fiscal quarter. | [optional] 

## Methods

### NewRevenueEstimatesInfo

`func NewRevenueEstimatesInfo() *RevenueEstimatesInfo`

NewRevenueEstimatesInfo instantiates a new RevenueEstimatesInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRevenueEstimatesInfoWithDefaults

`func NewRevenueEstimatesInfoWithDefaults() *RevenueEstimatesInfo`

NewRevenueEstimatesInfoWithDefaults instantiates a new RevenueEstimatesInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRevenueAvg

`func (o *RevenueEstimatesInfo) GetRevenueAvg() float32`

GetRevenueAvg returns the RevenueAvg field if non-nil, zero value otherwise.

### GetRevenueAvgOk

`func (o *RevenueEstimatesInfo) GetRevenueAvgOk() (*float32, bool)`

GetRevenueAvgOk returns a tuple with the RevenueAvg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevenueAvg

`func (o *RevenueEstimatesInfo) SetRevenueAvg(v float32)`

SetRevenueAvg sets RevenueAvg field to given value.

### HasRevenueAvg

`func (o *RevenueEstimatesInfo) HasRevenueAvg() bool`

HasRevenueAvg returns a boolean if a field has been set.

### GetRevenueHigh

`func (o *RevenueEstimatesInfo) GetRevenueHigh() float32`

GetRevenueHigh returns the RevenueHigh field if non-nil, zero value otherwise.

### GetRevenueHighOk

`func (o *RevenueEstimatesInfo) GetRevenueHighOk() (*float32, bool)`

GetRevenueHighOk returns a tuple with the RevenueHigh field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevenueHigh

`func (o *RevenueEstimatesInfo) SetRevenueHigh(v float32)`

SetRevenueHigh sets RevenueHigh field to given value.

### HasRevenueHigh

`func (o *RevenueEstimatesInfo) HasRevenueHigh() bool`

HasRevenueHigh returns a boolean if a field has been set.

### GetRevenueLow

`func (o *RevenueEstimatesInfo) GetRevenueLow() float32`

GetRevenueLow returns the RevenueLow field if non-nil, zero value otherwise.

### GetRevenueLowOk

`func (o *RevenueEstimatesInfo) GetRevenueLowOk() (*float32, bool)`

GetRevenueLowOk returns a tuple with the RevenueLow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevenueLow

`func (o *RevenueEstimatesInfo) SetRevenueLow(v float32)`

SetRevenueLow sets RevenueLow field to given value.

### HasRevenueLow

`func (o *RevenueEstimatesInfo) HasRevenueLow() bool`

HasRevenueLow returns a boolean if a field has been set.

### GetNumberAnalysts

`func (o *RevenueEstimatesInfo) GetNumberAnalysts() int64`

GetNumberAnalysts returns the NumberAnalysts field if non-nil, zero value otherwise.

### GetNumberAnalystsOk

`func (o *RevenueEstimatesInfo) GetNumberAnalystsOk() (*int64, bool)`

GetNumberAnalystsOk returns a tuple with the NumberAnalysts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumberAnalysts

`func (o *RevenueEstimatesInfo) SetNumberAnalysts(v int64)`

SetNumberAnalysts sets NumberAnalysts field to given value.

### HasNumberAnalysts

`func (o *RevenueEstimatesInfo) HasNumberAnalysts() bool`

HasNumberAnalysts returns a boolean if a field has been set.

### GetPeriod

`func (o *RevenueEstimatesInfo) GetPeriod() string`

GetPeriod returns the Period field if non-nil, zero value otherwise.

### GetPeriodOk

`func (o *RevenueEstimatesInfo) GetPeriodOk() (*string, bool)`

GetPeriodOk returns a tuple with the Period field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeriod

`func (o *RevenueEstimatesInfo) SetPeriod(v string)`

SetPeriod sets Period field to given value.

### HasPeriod

`func (o *RevenueEstimatesInfo) HasPeriod() bool`

HasPeriod returns a boolean if a field has been set.

### GetYear

`func (o *RevenueEstimatesInfo) GetYear() int64`

GetYear returns the Year field if non-nil, zero value otherwise.

### GetYearOk

`func (o *RevenueEstimatesInfo) GetYearOk() (*int64, bool)`

GetYearOk returns a tuple with the Year field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetYear

`func (o *RevenueEstimatesInfo) SetYear(v int64)`

SetYear sets Year field to given value.

### HasYear

`func (o *RevenueEstimatesInfo) HasYear() bool`

HasYear returns a boolean if a field has been set.

### GetQuarter

`func (o *RevenueEstimatesInfo) GetQuarter() int64`

GetQuarter returns the Quarter field if non-nil, zero value otherwise.

### GetQuarterOk

`func (o *RevenueEstimatesInfo) GetQuarterOk() (*int64, bool)`

GetQuarterOk returns a tuple with the Quarter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuarter

`func (o *RevenueEstimatesInfo) SetQuarter(v int64)`

SetQuarter sets Quarter field to given value.

### HasQuarter

`func (o *RevenueEstimatesInfo) HasQuarter() bool`

HasQuarter returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


