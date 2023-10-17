# MarketHoliday

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Timezone** | Pointer to **string** | Timezone. | [optional] 
**Exchange** | Pointer to **string** | Exchange. | [optional] 
**Data** | Pointer to [**[]MarketHolidayData**](MarketHolidayData.md) | Array of holidays. | [optional] 

## Methods

### NewMarketHoliday

`func NewMarketHoliday() *MarketHoliday`

NewMarketHoliday instantiates a new MarketHoliday object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMarketHolidayWithDefaults

`func NewMarketHolidayWithDefaults() *MarketHoliday`

NewMarketHolidayWithDefaults instantiates a new MarketHoliday object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTimezone

`func (o *MarketHoliday) GetTimezone() string`

GetTimezone returns the Timezone field if non-nil, zero value otherwise.

### GetTimezoneOk

`func (o *MarketHoliday) GetTimezoneOk() (*string, bool)`

GetTimezoneOk returns a tuple with the Timezone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimezone

`func (o *MarketHoliday) SetTimezone(v string)`

SetTimezone sets Timezone field to given value.

### HasTimezone

`func (o *MarketHoliday) HasTimezone() bool`

HasTimezone returns a boolean if a field has been set.

### GetExchange

`func (o *MarketHoliday) GetExchange() string`

GetExchange returns the Exchange field if non-nil, zero value otherwise.

### GetExchangeOk

`func (o *MarketHoliday) GetExchangeOk() (*string, bool)`

GetExchangeOk returns a tuple with the Exchange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExchange

`func (o *MarketHoliday) SetExchange(v string)`

SetExchange sets Exchange field to given value.

### HasExchange

`func (o *MarketHoliday) HasExchange() bool`

HasExchange returns a boolean if a field has been set.

### GetData

`func (o *MarketHoliday) GetData() []MarketHolidayData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *MarketHoliday) GetDataOk() (*[]MarketHolidayData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *MarketHoliday) SetData(v []MarketHolidayData)`

SetData sets Data field to given value.

### HasData

`func (o *MarketHoliday) HasData() bool`

HasData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


