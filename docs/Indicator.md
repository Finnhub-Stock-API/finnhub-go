# Indicator

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Buy** | Pointer to **int64** | Number of buy signals | [optional] 
**Neutral** | Pointer to **int64** | Number of neutral signals | [optional] 
**Sell** | Pointer to **int64** | Number of sell signals | [optional] 

## Methods

### NewIndicator

`func NewIndicator() *Indicator`

NewIndicator instantiates a new Indicator object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIndicatorWithDefaults

`func NewIndicatorWithDefaults() *Indicator`

NewIndicatorWithDefaults instantiates a new Indicator object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBuy

`func (o *Indicator) GetBuy() int64`

GetBuy returns the Buy field if non-nil, zero value otherwise.

### GetBuyOk

`func (o *Indicator) GetBuyOk() (*int64, bool)`

GetBuyOk returns a tuple with the Buy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuy

`func (o *Indicator) SetBuy(v int64)`

SetBuy sets Buy field to given value.

### HasBuy

`func (o *Indicator) HasBuy() bool`

HasBuy returns a boolean if a field has been set.

### GetNeutral

`func (o *Indicator) GetNeutral() int64`

GetNeutral returns the Neutral field if non-nil, zero value otherwise.

### GetNeutralOk

`func (o *Indicator) GetNeutralOk() (*int64, bool)`

GetNeutralOk returns a tuple with the Neutral field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNeutral

`func (o *Indicator) SetNeutral(v int64)`

SetNeutral sets Neutral field to given value.

### HasNeutral

`func (o *Indicator) HasNeutral() bool`

HasNeutral returns a boolean if a field has been set.

### GetSell

`func (o *Indicator) GetSell() int64`

GetSell returns the Sell field if non-nil, zero value otherwise.

### GetSellOk

`func (o *Indicator) GetSellOk() (*int64, bool)`

GetSellOk returns a tuple with the Sell field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSell

`func (o *Indicator) SetSell(v int64)`

SetSell sets Sell field to given value.

### HasSell

`func (o *Indicator) HasSell() bool`

HasSell returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


