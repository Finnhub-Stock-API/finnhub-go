# MarketCapData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AtDate** | Pointer to **string** | Date of the reading | [optional] 
**MarketCapitalization** | Pointer to **float32** | Value | [optional] 

## Methods

### NewMarketCapData

`func NewMarketCapData() *MarketCapData`

NewMarketCapData instantiates a new MarketCapData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMarketCapDataWithDefaults

`func NewMarketCapDataWithDefaults() *MarketCapData`

NewMarketCapDataWithDefaults instantiates a new MarketCapData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAtDate

`func (o *MarketCapData) GetAtDate() string`

GetAtDate returns the AtDate field if non-nil, zero value otherwise.

### GetAtDateOk

`func (o *MarketCapData) GetAtDateOk() (*string, bool)`

GetAtDateOk returns a tuple with the AtDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAtDate

`func (o *MarketCapData) SetAtDate(v string)`

SetAtDate sets AtDate field to given value.

### HasAtDate

`func (o *MarketCapData) HasAtDate() bool`

HasAtDate returns a boolean if a field has been set.

### GetMarketCapitalization

`func (o *MarketCapData) GetMarketCapitalization() float32`

GetMarketCapitalization returns the MarketCapitalization field if non-nil, zero value otherwise.

### GetMarketCapitalizationOk

`func (o *MarketCapData) GetMarketCapitalizationOk() (*float32, bool)`

GetMarketCapitalizationOk returns a tuple with the MarketCapitalization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarketCapitalization

`func (o *MarketCapData) SetMarketCapitalization(v float32)`

SetMarketCapitalization sets MarketCapitalization field to given value.

### HasMarketCapitalization

`func (o *MarketCapData) HasMarketCapitalization() bool`

HasMarketCapitalization returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


