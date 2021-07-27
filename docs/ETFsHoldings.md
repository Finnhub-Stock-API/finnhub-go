# ETFsHoldings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Symbol** | Pointer to **string** | ETF symbol. | [optional] 
**AtDate** | Pointer to **string** | Holdings update date. | [optional] 
**NumberOfHoldings** | Pointer to **int64** | Number of holdings. | [optional] 
**Holdings** | Pointer to [**[]ETFHoldingsData**](ETFHoldingsData.md) | Array of holdings. | [optional] 

## Methods

### NewETFsHoldings

`func NewETFsHoldings() *ETFsHoldings`

NewETFsHoldings instantiates a new ETFsHoldings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewETFsHoldingsWithDefaults

`func NewETFsHoldingsWithDefaults() *ETFsHoldings`

NewETFsHoldingsWithDefaults instantiates a new ETFsHoldings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSymbol

`func (o *ETFsHoldings) GetSymbol() string`

GetSymbol returns the Symbol field if non-nil, zero value otherwise.

### GetSymbolOk

`func (o *ETFsHoldings) GetSymbolOk() (*string, bool)`

GetSymbolOk returns a tuple with the Symbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSymbol

`func (o *ETFsHoldings) SetSymbol(v string)`

SetSymbol sets Symbol field to given value.

### HasSymbol

`func (o *ETFsHoldings) HasSymbol() bool`

HasSymbol returns a boolean if a field has been set.

### GetAtDate

`func (o *ETFsHoldings) GetAtDate() string`

GetAtDate returns the AtDate field if non-nil, zero value otherwise.

### GetAtDateOk

`func (o *ETFsHoldings) GetAtDateOk() (*string, bool)`

GetAtDateOk returns a tuple with the AtDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAtDate

`func (o *ETFsHoldings) SetAtDate(v string)`

SetAtDate sets AtDate field to given value.

### HasAtDate

`func (o *ETFsHoldings) HasAtDate() bool`

HasAtDate returns a boolean if a field has been set.

### GetNumberOfHoldings

`func (o *ETFsHoldings) GetNumberOfHoldings() int64`

GetNumberOfHoldings returns the NumberOfHoldings field if non-nil, zero value otherwise.

### GetNumberOfHoldingsOk

`func (o *ETFsHoldings) GetNumberOfHoldingsOk() (*int64, bool)`

GetNumberOfHoldingsOk returns a tuple with the NumberOfHoldings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumberOfHoldings

`func (o *ETFsHoldings) SetNumberOfHoldings(v int64)`

SetNumberOfHoldings sets NumberOfHoldings field to given value.

### HasNumberOfHoldings

`func (o *ETFsHoldings) HasNumberOfHoldings() bool`

HasNumberOfHoldings returns a boolean if a field has been set.

### GetHoldings

`func (o *ETFsHoldings) GetHoldings() []ETFHoldingsData`

GetHoldings returns the Holdings field if non-nil, zero value otherwise.

### GetHoldingsOk

`func (o *ETFsHoldings) GetHoldingsOk() (*[]ETFHoldingsData, bool)`

GetHoldingsOk returns a tuple with the Holdings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHoldings

`func (o *ETFsHoldings) SetHoldings(v []ETFHoldingsData)`

SetHoldings sets Holdings field to given value.

### HasHoldings

`func (o *ETFsHoldings) HasHoldings() bool`

HasHoldings returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


