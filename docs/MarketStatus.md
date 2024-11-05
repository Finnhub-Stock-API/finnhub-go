# MarketStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Exchange** | Pointer to **string** | Exchange. | [optional] 
**Timezone** | Pointer to **string** | Timezone. | [optional] 
**Session** | Pointer to **string** | Market session. Can be 1 of the following values: &lt;code&gt;pre-market&lt;/code&gt;,&lt;code&gt;regular&lt;/code&gt;,&lt;code&gt;post-market&lt;/code&gt; or &lt;code&gt;null&lt;/code&gt; if the market is closed. | [optional] 
**Holiday** | Pointer to **string** | Holiday event. | [optional] 
**IsOpen** | Pointer to **bool** | Whether the market is open at the moment. | [optional] 
**T** | Pointer to **int64** | Current timestamp. | [optional] 

## Methods

### NewMarketStatus

`func NewMarketStatus() *MarketStatus`

NewMarketStatus instantiates a new MarketStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMarketStatusWithDefaults

`func NewMarketStatusWithDefaults() *MarketStatus`

NewMarketStatusWithDefaults instantiates a new MarketStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExchange

`func (o *MarketStatus) GetExchange() string`

GetExchange returns the Exchange field if non-nil, zero value otherwise.

### GetExchangeOk

`func (o *MarketStatus) GetExchangeOk() (*string, bool)`

GetExchangeOk returns a tuple with the Exchange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExchange

`func (o *MarketStatus) SetExchange(v string)`

SetExchange sets Exchange field to given value.

### HasExchange

`func (o *MarketStatus) HasExchange() bool`

HasExchange returns a boolean if a field has been set.

### GetTimezone

`func (o *MarketStatus) GetTimezone() string`

GetTimezone returns the Timezone field if non-nil, zero value otherwise.

### GetTimezoneOk

`func (o *MarketStatus) GetTimezoneOk() (*string, bool)`

GetTimezoneOk returns a tuple with the Timezone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimezone

`func (o *MarketStatus) SetTimezone(v string)`

SetTimezone sets Timezone field to given value.

### HasTimezone

`func (o *MarketStatus) HasTimezone() bool`

HasTimezone returns a boolean if a field has been set.

### GetSession

`func (o *MarketStatus) GetSession() string`

GetSession returns the Session field if non-nil, zero value otherwise.

### GetSessionOk

`func (o *MarketStatus) GetSessionOk() (*string, bool)`

GetSessionOk returns a tuple with the Session field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSession

`func (o *MarketStatus) SetSession(v string)`

SetSession sets Session field to given value.

### HasSession

`func (o *MarketStatus) HasSession() bool`

HasSession returns a boolean if a field has been set.

### GetHoliday

`func (o *MarketStatus) GetHoliday() string`

GetHoliday returns the Holiday field if non-nil, zero value otherwise.

### GetHolidayOk

`func (o *MarketStatus) GetHolidayOk() (*string, bool)`

GetHolidayOk returns a tuple with the Holiday field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHoliday

`func (o *MarketStatus) SetHoliday(v string)`

SetHoliday sets Holiday field to given value.

### HasHoliday

`func (o *MarketStatus) HasHoliday() bool`

HasHoliday returns a boolean if a field has been set.

### GetIsOpen

`func (o *MarketStatus) GetIsOpen() bool`

GetIsOpen returns the IsOpen field if non-nil, zero value otherwise.

### GetIsOpenOk

`func (o *MarketStatus) GetIsOpenOk() (*bool, bool)`

GetIsOpenOk returns a tuple with the IsOpen field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsOpen

`func (o *MarketStatus) SetIsOpen(v bool)`

SetIsOpen sets IsOpen field to given value.

### HasIsOpen

`func (o *MarketStatus) HasIsOpen() bool`

HasIsOpen returns a boolean if a field has been set.

### GetT

`func (o *MarketStatus) GetT() int64`

GetT returns the T field if non-nil, zero value otherwise.

### GetTOk

`func (o *MarketStatus) GetTOk() (*int64, bool)`

GetTOk returns a tuple with the T field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetT

`func (o *MarketStatus) SetT(v int64)`

SetT sets T field to given value.

### HasT

`func (o *MarketStatus) HasT() bool`

HasT returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


