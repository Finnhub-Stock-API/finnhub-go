# MarketHolidayData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EventName** | Pointer to **string** | Holiday&#39;s name. | [optional] 
**AtDate** | Pointer to **string** | Date. | [optional] 
**TradingHour** | Pointer to **string** | Trading hours for this day if the market is partially closed only. | [optional] 

## Methods

### NewMarketHolidayData

`func NewMarketHolidayData() *MarketHolidayData`

NewMarketHolidayData instantiates a new MarketHolidayData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMarketHolidayDataWithDefaults

`func NewMarketHolidayDataWithDefaults() *MarketHolidayData`

NewMarketHolidayDataWithDefaults instantiates a new MarketHolidayData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEventName

`func (o *MarketHolidayData) GetEventName() string`

GetEventName returns the EventName field if non-nil, zero value otherwise.

### GetEventNameOk

`func (o *MarketHolidayData) GetEventNameOk() (*string, bool)`

GetEventNameOk returns a tuple with the EventName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventName

`func (o *MarketHolidayData) SetEventName(v string)`

SetEventName sets EventName field to given value.

### HasEventName

`func (o *MarketHolidayData) HasEventName() bool`

HasEventName returns a boolean if a field has been set.

### GetAtDate

`func (o *MarketHolidayData) GetAtDate() string`

GetAtDate returns the AtDate field if non-nil, zero value otherwise.

### GetAtDateOk

`func (o *MarketHolidayData) GetAtDateOk() (*string, bool)`

GetAtDateOk returns a tuple with the AtDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAtDate

`func (o *MarketHolidayData) SetAtDate(v string)`

SetAtDate sets AtDate field to given value.

### HasAtDate

`func (o *MarketHolidayData) HasAtDate() bool`

HasAtDate returns a boolean if a field has been set.

### GetTradingHour

`func (o *MarketHolidayData) GetTradingHour() string`

GetTradingHour returns the TradingHour field if non-nil, zero value otherwise.

### GetTradingHourOk

`func (o *MarketHolidayData) GetTradingHourOk() (*string, bool)`

GetTradingHourOk returns a tuple with the TradingHour field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTradingHour

`func (o *MarketHolidayData) SetTradingHour(v string)`

SetTradingHour sets TradingHour field to given value.

### HasTradingHour

`func (o *MarketHolidayData) HasTradingHour() bool`

HasTradingHour returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


