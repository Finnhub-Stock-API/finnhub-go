# EarningResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Actual** | Pointer to **float32** | Actual earning result. | [optional] 
**Estimate** | Pointer to **float32** | Estimated earning. | [optional] 
**Surprise** | Pointer to **float32** | Surprise - The difference between actual and estimate. | [optional] 
**SurprisePercent** | Pointer to **float32** | Surprise percent. | [optional] 
**Period** | Pointer to **string** | Reported period. | [optional] 
**Symbol** | Pointer to **string** | Company symbol. | [optional] 
**Year** | Pointer to **int64** | Earnings year. | [optional] 
**Quarter** | Pointer to **int64** | Earnings quarter. | [optional] 

## Methods

### NewEarningResult

`func NewEarningResult() *EarningResult`

NewEarningResult instantiates a new EarningResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEarningResultWithDefaults

`func NewEarningResultWithDefaults() *EarningResult`

NewEarningResultWithDefaults instantiates a new EarningResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActual

`func (o *EarningResult) GetActual() float32`

GetActual returns the Actual field if non-nil, zero value otherwise.

### GetActualOk

`func (o *EarningResult) GetActualOk() (*float32, bool)`

GetActualOk returns a tuple with the Actual field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActual

`func (o *EarningResult) SetActual(v float32)`

SetActual sets Actual field to given value.

### HasActual

`func (o *EarningResult) HasActual() bool`

HasActual returns a boolean if a field has been set.

### GetEstimate

`func (o *EarningResult) GetEstimate() float32`

GetEstimate returns the Estimate field if non-nil, zero value otherwise.

### GetEstimateOk

`func (o *EarningResult) GetEstimateOk() (*float32, bool)`

GetEstimateOk returns a tuple with the Estimate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEstimate

`func (o *EarningResult) SetEstimate(v float32)`

SetEstimate sets Estimate field to given value.

### HasEstimate

`func (o *EarningResult) HasEstimate() bool`

HasEstimate returns a boolean if a field has been set.

### GetSurprise

`func (o *EarningResult) GetSurprise() float32`

GetSurprise returns the Surprise field if non-nil, zero value otherwise.

### GetSurpriseOk

`func (o *EarningResult) GetSurpriseOk() (*float32, bool)`

GetSurpriseOk returns a tuple with the Surprise field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSurprise

`func (o *EarningResult) SetSurprise(v float32)`

SetSurprise sets Surprise field to given value.

### HasSurprise

`func (o *EarningResult) HasSurprise() bool`

HasSurprise returns a boolean if a field has been set.

### GetSurprisePercent

`func (o *EarningResult) GetSurprisePercent() float32`

GetSurprisePercent returns the SurprisePercent field if non-nil, zero value otherwise.

### GetSurprisePercentOk

`func (o *EarningResult) GetSurprisePercentOk() (*float32, bool)`

GetSurprisePercentOk returns a tuple with the SurprisePercent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSurprisePercent

`func (o *EarningResult) SetSurprisePercent(v float32)`

SetSurprisePercent sets SurprisePercent field to given value.

### HasSurprisePercent

`func (o *EarningResult) HasSurprisePercent() bool`

HasSurprisePercent returns a boolean if a field has been set.

### GetPeriod

`func (o *EarningResult) GetPeriod() string`

GetPeriod returns the Period field if non-nil, zero value otherwise.

### GetPeriodOk

`func (o *EarningResult) GetPeriodOk() (*string, bool)`

GetPeriodOk returns a tuple with the Period field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeriod

`func (o *EarningResult) SetPeriod(v string)`

SetPeriod sets Period field to given value.

### HasPeriod

`func (o *EarningResult) HasPeriod() bool`

HasPeriod returns a boolean if a field has been set.

### GetSymbol

`func (o *EarningResult) GetSymbol() string`

GetSymbol returns the Symbol field if non-nil, zero value otherwise.

### GetSymbolOk

`func (o *EarningResult) GetSymbolOk() (*string, bool)`

GetSymbolOk returns a tuple with the Symbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSymbol

`func (o *EarningResult) SetSymbol(v string)`

SetSymbol sets Symbol field to given value.

### HasSymbol

`func (o *EarningResult) HasSymbol() bool`

HasSymbol returns a boolean if a field has been set.

### GetYear

`func (o *EarningResult) GetYear() int64`

GetYear returns the Year field if non-nil, zero value otherwise.

### GetYearOk

`func (o *EarningResult) GetYearOk() (*int64, bool)`

GetYearOk returns a tuple with the Year field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetYear

`func (o *EarningResult) SetYear(v int64)`

SetYear sets Year field to given value.

### HasYear

`func (o *EarningResult) HasYear() bool`

HasYear returns a boolean if a field has been set.

### GetQuarter

`func (o *EarningResult) GetQuarter() int64`

GetQuarter returns the Quarter field if non-nil, zero value otherwise.

### GetQuarterOk

`func (o *EarningResult) GetQuarterOk() (*int64, bool)`

GetQuarterOk returns a tuple with the Quarter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuarter

`func (o *EarningResult) SetQuarter(v int64)`

SetQuarter sets Quarter field to given value.

### HasQuarter

`func (o *EarningResult) HasQuarter() bool`

HasQuarter returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


