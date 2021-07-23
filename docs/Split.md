# Split

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Symbol** | Pointer to **string** | Symbol. | [optional] 
**Date** | Pointer to **string** | Split date. | [optional] 
**FromFactor** | Pointer to **float32** | From factor. | [optional] 
**ToFactor** | Pointer to **float32** | To factor. | [optional] 

## Methods

### NewSplit

`func NewSplit() *Split`

NewSplit instantiates a new Split object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSplitWithDefaults

`func NewSplitWithDefaults() *Split`

NewSplitWithDefaults instantiates a new Split object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSymbol

`func (o *Split) GetSymbol() string`

GetSymbol returns the Symbol field if non-nil, zero value otherwise.

### GetSymbolOk

`func (o *Split) GetSymbolOk() (*string, bool)`

GetSymbolOk returns a tuple with the Symbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSymbol

`func (o *Split) SetSymbol(v string)`

SetSymbol sets Symbol field to given value.

### HasSymbol

`func (o *Split) HasSymbol() bool`

HasSymbol returns a boolean if a field has been set.

### GetDate

`func (o *Split) GetDate() string`

GetDate returns the Date field if non-nil, zero value otherwise.

### GetDateOk

`func (o *Split) GetDateOk() (*string, bool)`

GetDateOk returns a tuple with the Date field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDate

`func (o *Split) SetDate(v string)`

SetDate sets Date field to given value.

### HasDate

`func (o *Split) HasDate() bool`

HasDate returns a boolean if a field has been set.

### GetFromFactor

`func (o *Split) GetFromFactor() float32`

GetFromFactor returns the FromFactor field if non-nil, zero value otherwise.

### GetFromFactorOk

`func (o *Split) GetFromFactorOk() (*float32, bool)`

GetFromFactorOk returns a tuple with the FromFactor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFromFactor

`func (o *Split) SetFromFactor(v float32)`

SetFromFactor sets FromFactor field to given value.

### HasFromFactor

`func (o *Split) HasFromFactor() bool`

HasFromFactor returns a boolean if a field has been set.

### GetToFactor

`func (o *Split) GetToFactor() float32`

GetToFactor returns the ToFactor field if non-nil, zero value otherwise.

### GetToFactorOk

`func (o *Split) GetToFactorOk() (*float32, bool)`

GetToFactorOk returns a tuple with the ToFactor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToFactor

`func (o *Split) SetToFactor(v float32)`

SetToFactor sets ToFactor field to given value.

### HasToFactor

`func (o *Split) HasToFactor() bool`

HasToFactor returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


