# MutualFundHoldings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Symbol** | Pointer to **string** | Symbol. | [optional] 
**AtDate** | Pointer to **string** | Holdings update date. | [optional] 
**NumberOfHoldings** | Pointer to **int64** | Number of holdings. | [optional] 
**Holdings** | Pointer to [**[]MutualFundHoldingsData**](MutualFundHoldingsData.md) | Array of holdings. | [optional] 

## Methods

### NewMutualFundHoldings

`func NewMutualFundHoldings() *MutualFundHoldings`

NewMutualFundHoldings instantiates a new MutualFundHoldings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMutualFundHoldingsWithDefaults

`func NewMutualFundHoldingsWithDefaults() *MutualFundHoldings`

NewMutualFundHoldingsWithDefaults instantiates a new MutualFundHoldings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSymbol

`func (o *MutualFundHoldings) GetSymbol() string`

GetSymbol returns the Symbol field if non-nil, zero value otherwise.

### GetSymbolOk

`func (o *MutualFundHoldings) GetSymbolOk() (*string, bool)`

GetSymbolOk returns a tuple with the Symbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSymbol

`func (o *MutualFundHoldings) SetSymbol(v string)`

SetSymbol sets Symbol field to given value.

### HasSymbol

`func (o *MutualFundHoldings) HasSymbol() bool`

HasSymbol returns a boolean if a field has been set.

### GetAtDate

`func (o *MutualFundHoldings) GetAtDate() string`

GetAtDate returns the AtDate field if non-nil, zero value otherwise.

### GetAtDateOk

`func (o *MutualFundHoldings) GetAtDateOk() (*string, bool)`

GetAtDateOk returns a tuple with the AtDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAtDate

`func (o *MutualFundHoldings) SetAtDate(v string)`

SetAtDate sets AtDate field to given value.

### HasAtDate

`func (o *MutualFundHoldings) HasAtDate() bool`

HasAtDate returns a boolean if a field has been set.

### GetNumberOfHoldings

`func (o *MutualFundHoldings) GetNumberOfHoldings() int64`

GetNumberOfHoldings returns the NumberOfHoldings field if non-nil, zero value otherwise.

### GetNumberOfHoldingsOk

`func (o *MutualFundHoldings) GetNumberOfHoldingsOk() (*int64, bool)`

GetNumberOfHoldingsOk returns a tuple with the NumberOfHoldings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumberOfHoldings

`func (o *MutualFundHoldings) SetNumberOfHoldings(v int64)`

SetNumberOfHoldings sets NumberOfHoldings field to given value.

### HasNumberOfHoldings

`func (o *MutualFundHoldings) HasNumberOfHoldings() bool`

HasNumberOfHoldings returns a boolean if a field has been set.

### GetHoldings

`func (o *MutualFundHoldings) GetHoldings() []MutualFundHoldingsData`

GetHoldings returns the Holdings field if non-nil, zero value otherwise.

### GetHoldingsOk

`func (o *MutualFundHoldings) GetHoldingsOk() (*[]MutualFundHoldingsData, bool)`

GetHoldingsOk returns a tuple with the Holdings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHoldings

`func (o *MutualFundHoldings) SetHoldings(v []MutualFundHoldingsData)`

SetHoldings sets Holdings field to given value.

### HasHoldings

`func (o *MutualFundHoldings) HasHoldings() bool`

HasHoldings returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


