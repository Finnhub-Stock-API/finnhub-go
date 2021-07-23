# Trend

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Adx** | Pointer to **float32** | ADX reading | [optional] 
**Trending** | Pointer to **bool** | Whether market is trending or going sideway | [optional] 

## Methods

### NewTrend

`func NewTrend() *Trend`

NewTrend instantiates a new Trend object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTrendWithDefaults

`func NewTrendWithDefaults() *Trend`

NewTrendWithDefaults instantiates a new Trend object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAdx

`func (o *Trend) GetAdx() float32`

GetAdx returns the Adx field if non-nil, zero value otherwise.

### GetAdxOk

`func (o *Trend) GetAdxOk() (*float32, bool)`

GetAdxOk returns a tuple with the Adx field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdx

`func (o *Trend) SetAdx(v float32)`

SetAdx sets Adx field to given value.

### HasAdx

`func (o *Trend) HasAdx() bool`

HasAdx returns a boolean if a field has been set.

### GetTrending

`func (o *Trend) GetTrending() bool`

GetTrending returns the Trending field if non-nil, zero value otherwise.

### GetTrendingOk

`func (o *Trend) GetTrendingOk() (*bool, bool)`

GetTrendingOk returns a tuple with the Trending field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrending

`func (o *Trend) SetTrending(v bool)`

SetTrending sets Trending field to given value.

### HasTrending

`func (o *Trend) HasTrending() bool`

HasTrending returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


