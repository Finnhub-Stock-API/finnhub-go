# InsiderSentimentsData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Symbol** | Pointer to **string** | Symbol. | [optional] 
**Year** | Pointer to **int64** | Year. | [optional] 
**Month** | Pointer to **int64** | Month. | [optional] 
**Change** | Pointer to **int64** | Net buying/selling from all insiders&#39; transactions. | [optional] 
**Mspr** | Pointer to **float32** | Monthly share purchase ratio. | [optional] 

## Methods

### NewInsiderSentimentsData

`func NewInsiderSentimentsData() *InsiderSentimentsData`

NewInsiderSentimentsData instantiates a new InsiderSentimentsData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInsiderSentimentsDataWithDefaults

`func NewInsiderSentimentsDataWithDefaults() *InsiderSentimentsData`

NewInsiderSentimentsDataWithDefaults instantiates a new InsiderSentimentsData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSymbol

`func (o *InsiderSentimentsData) GetSymbol() string`

GetSymbol returns the Symbol field if non-nil, zero value otherwise.

### GetSymbolOk

`func (o *InsiderSentimentsData) GetSymbolOk() (*string, bool)`

GetSymbolOk returns a tuple with the Symbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSymbol

`func (o *InsiderSentimentsData) SetSymbol(v string)`

SetSymbol sets Symbol field to given value.

### HasSymbol

`func (o *InsiderSentimentsData) HasSymbol() bool`

HasSymbol returns a boolean if a field has been set.

### GetYear

`func (o *InsiderSentimentsData) GetYear() int64`

GetYear returns the Year field if non-nil, zero value otherwise.

### GetYearOk

`func (o *InsiderSentimentsData) GetYearOk() (*int64, bool)`

GetYearOk returns a tuple with the Year field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetYear

`func (o *InsiderSentimentsData) SetYear(v int64)`

SetYear sets Year field to given value.

### HasYear

`func (o *InsiderSentimentsData) HasYear() bool`

HasYear returns a boolean if a field has been set.

### GetMonth

`func (o *InsiderSentimentsData) GetMonth() int64`

GetMonth returns the Month field if non-nil, zero value otherwise.

### GetMonthOk

`func (o *InsiderSentimentsData) GetMonthOk() (*int64, bool)`

GetMonthOk returns a tuple with the Month field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonth

`func (o *InsiderSentimentsData) SetMonth(v int64)`

SetMonth sets Month field to given value.

### HasMonth

`func (o *InsiderSentimentsData) HasMonth() bool`

HasMonth returns a boolean if a field has been set.

### GetChange

`func (o *InsiderSentimentsData) GetChange() int64`

GetChange returns the Change field if non-nil, zero value otherwise.

### GetChangeOk

`func (o *InsiderSentimentsData) GetChangeOk() (*int64, bool)`

GetChangeOk returns a tuple with the Change field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChange

`func (o *InsiderSentimentsData) SetChange(v int64)`

SetChange sets Change field to given value.

### HasChange

`func (o *InsiderSentimentsData) HasChange() bool`

HasChange returns a boolean if a field has been set.

### GetMspr

`func (o *InsiderSentimentsData) GetMspr() float32`

GetMspr returns the Mspr field if non-nil, zero value otherwise.

### GetMsprOk

`func (o *InsiderSentimentsData) GetMsprOk() (*float32, bool)`

GetMsprOk returns a tuple with the Mspr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMspr

`func (o *InsiderSentimentsData) SetMspr(v float32)`

SetMspr sets Mspr field to given value.

### HasMspr

`func (o *InsiderSentimentsData) HasMspr() bool`

HasMspr returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


