# RecommendationTrend

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Symbol** | Pointer to **string** | Company symbol. | [optional] 
**Buy** | Pointer to **int64** | Number of recommendations that fall into the Buy category | [optional] 
**Hold** | Pointer to **int64** | Number of recommendations that fall into the Hold category | [optional] 
**Period** | Pointer to **string** | Updated period | [optional] 
**Sell** | Pointer to **int64** | Number of recommendations that fall into the Sell category | [optional] 
**StrongBuy** | Pointer to **int64** | Number of recommendations that fall into the Strong Buy category | [optional] 
**StrongSell** | Pointer to **int64** | Number of recommendations that fall into the Strong Sell category | [optional] 

## Methods

### NewRecommendationTrend

`func NewRecommendationTrend() *RecommendationTrend`

NewRecommendationTrend instantiates a new RecommendationTrend object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRecommendationTrendWithDefaults

`func NewRecommendationTrendWithDefaults() *RecommendationTrend`

NewRecommendationTrendWithDefaults instantiates a new RecommendationTrend object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSymbol

`func (o *RecommendationTrend) GetSymbol() string`

GetSymbol returns the Symbol field if non-nil, zero value otherwise.

### GetSymbolOk

`func (o *RecommendationTrend) GetSymbolOk() (*string, bool)`

GetSymbolOk returns a tuple with the Symbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSymbol

`func (o *RecommendationTrend) SetSymbol(v string)`

SetSymbol sets Symbol field to given value.

### HasSymbol

`func (o *RecommendationTrend) HasSymbol() bool`

HasSymbol returns a boolean if a field has been set.

### GetBuy

`func (o *RecommendationTrend) GetBuy() int64`

GetBuy returns the Buy field if non-nil, zero value otherwise.

### GetBuyOk

`func (o *RecommendationTrend) GetBuyOk() (*int64, bool)`

GetBuyOk returns a tuple with the Buy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuy

`func (o *RecommendationTrend) SetBuy(v int64)`

SetBuy sets Buy field to given value.

### HasBuy

`func (o *RecommendationTrend) HasBuy() bool`

HasBuy returns a boolean if a field has been set.

### GetHold

`func (o *RecommendationTrend) GetHold() int64`

GetHold returns the Hold field if non-nil, zero value otherwise.

### GetHoldOk

`func (o *RecommendationTrend) GetHoldOk() (*int64, bool)`

GetHoldOk returns a tuple with the Hold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHold

`func (o *RecommendationTrend) SetHold(v int64)`

SetHold sets Hold field to given value.

### HasHold

`func (o *RecommendationTrend) HasHold() bool`

HasHold returns a boolean if a field has been set.

### GetPeriod

`func (o *RecommendationTrend) GetPeriod() string`

GetPeriod returns the Period field if non-nil, zero value otherwise.

### GetPeriodOk

`func (o *RecommendationTrend) GetPeriodOk() (*string, bool)`

GetPeriodOk returns a tuple with the Period field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeriod

`func (o *RecommendationTrend) SetPeriod(v string)`

SetPeriod sets Period field to given value.

### HasPeriod

`func (o *RecommendationTrend) HasPeriod() bool`

HasPeriod returns a boolean if a field has been set.

### GetSell

`func (o *RecommendationTrend) GetSell() int64`

GetSell returns the Sell field if non-nil, zero value otherwise.

### GetSellOk

`func (o *RecommendationTrend) GetSellOk() (*int64, bool)`

GetSellOk returns a tuple with the Sell field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSell

`func (o *RecommendationTrend) SetSell(v int64)`

SetSell sets Sell field to given value.

### HasSell

`func (o *RecommendationTrend) HasSell() bool`

HasSell returns a boolean if a field has been set.

### GetStrongBuy

`func (o *RecommendationTrend) GetStrongBuy() int64`

GetStrongBuy returns the StrongBuy field if non-nil, zero value otherwise.

### GetStrongBuyOk

`func (o *RecommendationTrend) GetStrongBuyOk() (*int64, bool)`

GetStrongBuyOk returns a tuple with the StrongBuy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStrongBuy

`func (o *RecommendationTrend) SetStrongBuy(v int64)`

SetStrongBuy sets StrongBuy field to given value.

### HasStrongBuy

`func (o *RecommendationTrend) HasStrongBuy() bool`

HasStrongBuy returns a boolean if a field has been set.

### GetStrongSell

`func (o *RecommendationTrend) GetStrongSell() int64`

GetStrongSell returns the StrongSell field if non-nil, zero value otherwise.

### GetStrongSellOk

`func (o *RecommendationTrend) GetStrongSellOk() (*int64, bool)`

GetStrongSellOk returns a tuple with the StrongSell field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStrongSell

`func (o *RecommendationTrend) SetStrongSell(v int64)`

SetStrongSell sets StrongSell field to given value.

### HasStrongSell

`func (o *RecommendationTrend) HasStrongSell() bool`

HasStrongSell returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


