# EarningRelease

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Symbol** | Pointer to **string** | Symbol. | [optional] 
**Date** | Pointer to **string** | Date. | [optional] 
**Hour** | Pointer to **string** | Indicates whether the earnings is announced before market open(&lt;code&gt;bmo&lt;/code&gt;), after market close(&lt;code&gt;amc&lt;/code&gt;), or during market hour(&lt;code&gt;dmh&lt;/code&gt;). | [optional] 
**Year** | Pointer to **int64** | Earnings year. | [optional] 
**Quarter** | Pointer to **int64** | Earnings quarter. | [optional] 
**EpsEstimate** | Pointer to **float32** | EPS estimate. | [optional] 
**EpsActual** | Pointer to **float32** | EPS actual. | [optional] 
**RevenueEstimate** | Pointer to **float32** | Revenue estimate including Finnhub&#39;s proprietary estimates. | [optional] 
**RevenueActual** | Pointer to **float32** | Revenue actual. | [optional] 

## Methods

### NewEarningRelease

`func NewEarningRelease() *EarningRelease`

NewEarningRelease instantiates a new EarningRelease object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEarningReleaseWithDefaults

`func NewEarningReleaseWithDefaults() *EarningRelease`

NewEarningReleaseWithDefaults instantiates a new EarningRelease object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSymbol

`func (o *EarningRelease) GetSymbol() string`

GetSymbol returns the Symbol field if non-nil, zero value otherwise.

### GetSymbolOk

`func (o *EarningRelease) GetSymbolOk() (*string, bool)`

GetSymbolOk returns a tuple with the Symbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSymbol

`func (o *EarningRelease) SetSymbol(v string)`

SetSymbol sets Symbol field to given value.

### HasSymbol

`func (o *EarningRelease) HasSymbol() bool`

HasSymbol returns a boolean if a field has been set.

### GetDate

`func (o *EarningRelease) GetDate() string`

GetDate returns the Date field if non-nil, zero value otherwise.

### GetDateOk

`func (o *EarningRelease) GetDateOk() (*string, bool)`

GetDateOk returns a tuple with the Date field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDate

`func (o *EarningRelease) SetDate(v string)`

SetDate sets Date field to given value.

### HasDate

`func (o *EarningRelease) HasDate() bool`

HasDate returns a boolean if a field has been set.

### GetHour

`func (o *EarningRelease) GetHour() string`

GetHour returns the Hour field if non-nil, zero value otherwise.

### GetHourOk

`func (o *EarningRelease) GetHourOk() (*string, bool)`

GetHourOk returns a tuple with the Hour field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHour

`func (o *EarningRelease) SetHour(v string)`

SetHour sets Hour field to given value.

### HasHour

`func (o *EarningRelease) HasHour() bool`

HasHour returns a boolean if a field has been set.

### GetYear

`func (o *EarningRelease) GetYear() int64`

GetYear returns the Year field if non-nil, zero value otherwise.

### GetYearOk

`func (o *EarningRelease) GetYearOk() (*int64, bool)`

GetYearOk returns a tuple with the Year field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetYear

`func (o *EarningRelease) SetYear(v int64)`

SetYear sets Year field to given value.

### HasYear

`func (o *EarningRelease) HasYear() bool`

HasYear returns a boolean if a field has been set.

### GetQuarter

`func (o *EarningRelease) GetQuarter() int64`

GetQuarter returns the Quarter field if non-nil, zero value otherwise.

### GetQuarterOk

`func (o *EarningRelease) GetQuarterOk() (*int64, bool)`

GetQuarterOk returns a tuple with the Quarter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuarter

`func (o *EarningRelease) SetQuarter(v int64)`

SetQuarter sets Quarter field to given value.

### HasQuarter

`func (o *EarningRelease) HasQuarter() bool`

HasQuarter returns a boolean if a field has been set.

### GetEpsEstimate

`func (o *EarningRelease) GetEpsEstimate() float32`

GetEpsEstimate returns the EpsEstimate field if non-nil, zero value otherwise.

### GetEpsEstimateOk

`func (o *EarningRelease) GetEpsEstimateOk() (*float32, bool)`

GetEpsEstimateOk returns a tuple with the EpsEstimate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEpsEstimate

`func (o *EarningRelease) SetEpsEstimate(v float32)`

SetEpsEstimate sets EpsEstimate field to given value.

### HasEpsEstimate

`func (o *EarningRelease) HasEpsEstimate() bool`

HasEpsEstimate returns a boolean if a field has been set.

### GetEpsActual

`func (o *EarningRelease) GetEpsActual() float32`

GetEpsActual returns the EpsActual field if non-nil, zero value otherwise.

### GetEpsActualOk

`func (o *EarningRelease) GetEpsActualOk() (*float32, bool)`

GetEpsActualOk returns a tuple with the EpsActual field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEpsActual

`func (o *EarningRelease) SetEpsActual(v float32)`

SetEpsActual sets EpsActual field to given value.

### HasEpsActual

`func (o *EarningRelease) HasEpsActual() bool`

HasEpsActual returns a boolean if a field has been set.

### GetRevenueEstimate

`func (o *EarningRelease) GetRevenueEstimate() float32`

GetRevenueEstimate returns the RevenueEstimate field if non-nil, zero value otherwise.

### GetRevenueEstimateOk

`func (o *EarningRelease) GetRevenueEstimateOk() (*float32, bool)`

GetRevenueEstimateOk returns a tuple with the RevenueEstimate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevenueEstimate

`func (o *EarningRelease) SetRevenueEstimate(v float32)`

SetRevenueEstimate sets RevenueEstimate field to given value.

### HasRevenueEstimate

`func (o *EarningRelease) HasRevenueEstimate() bool`

HasRevenueEstimate returns a boolean if a field has been set.

### GetRevenueActual

`func (o *EarningRelease) GetRevenueActual() float32`

GetRevenueActual returns the RevenueActual field if non-nil, zero value otherwise.

### GetRevenueActualOk

`func (o *EarningRelease) GetRevenueActualOk() (*float32, bool)`

GetRevenueActualOk returns a tuple with the RevenueActual field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevenueActual

`func (o *EarningRelease) SetRevenueActual(v float32)`

SetRevenueActual sets RevenueActual field to given value.

### HasRevenueActual

`func (o *EarningRelease) HasRevenueActual() bool`

HasRevenueActual returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


