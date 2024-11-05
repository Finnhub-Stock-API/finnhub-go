# AirlinePriceIndex

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Date** | Pointer to **string** | Date | [optional] 
**PriceIndex** | Pointer to **float32** | Price Index | [optional] 
**DailyAvgPrice** | Pointer to **float32** | Daily average ticket price. | [optional] 

## Methods

### NewAirlinePriceIndex

`func NewAirlinePriceIndex() *AirlinePriceIndex`

NewAirlinePriceIndex instantiates a new AirlinePriceIndex object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAirlinePriceIndexWithDefaults

`func NewAirlinePriceIndexWithDefaults() *AirlinePriceIndex`

NewAirlinePriceIndexWithDefaults instantiates a new AirlinePriceIndex object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDate

`func (o *AirlinePriceIndex) GetDate() string`

GetDate returns the Date field if non-nil, zero value otherwise.

### GetDateOk

`func (o *AirlinePriceIndex) GetDateOk() (*string, bool)`

GetDateOk returns a tuple with the Date field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDate

`func (o *AirlinePriceIndex) SetDate(v string)`

SetDate sets Date field to given value.

### HasDate

`func (o *AirlinePriceIndex) HasDate() bool`

HasDate returns a boolean if a field has been set.

### GetPriceIndex

`func (o *AirlinePriceIndex) GetPriceIndex() float32`

GetPriceIndex returns the PriceIndex field if non-nil, zero value otherwise.

### GetPriceIndexOk

`func (o *AirlinePriceIndex) GetPriceIndexOk() (*float32, bool)`

GetPriceIndexOk returns a tuple with the PriceIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceIndex

`func (o *AirlinePriceIndex) SetPriceIndex(v float32)`

SetPriceIndex sets PriceIndex field to given value.

### HasPriceIndex

`func (o *AirlinePriceIndex) HasPriceIndex() bool`

HasPriceIndex returns a boolean if a field has been set.

### GetDailyAvgPrice

`func (o *AirlinePriceIndex) GetDailyAvgPrice() float32`

GetDailyAvgPrice returns the DailyAvgPrice field if non-nil, zero value otherwise.

### GetDailyAvgPriceOk

`func (o *AirlinePriceIndex) GetDailyAvgPriceOk() (*float32, bool)`

GetDailyAvgPriceOk returns a tuple with the DailyAvgPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDailyAvgPrice

`func (o *AirlinePriceIndex) SetDailyAvgPrice(v float32)`

SetDailyAvgPrice sets DailyAvgPrice field to given value.

### HasDailyAvgPrice

`func (o *AirlinePriceIndex) HasDailyAvgPrice() bool`

HasDailyAvgPrice returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


