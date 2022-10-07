# SymbolChange

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FromDate** | Pointer to **string** | From date. | [optional] 
**ToDate** | Pointer to **string** | To date. | [optional] 
**Data** | Pointer to [**[]SymbolChangeInfo**](SymbolChangeInfo.md) | Array of symbol change events. | [optional] 

## Methods

### NewSymbolChange

`func NewSymbolChange() *SymbolChange`

NewSymbolChange instantiates a new SymbolChange object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSymbolChangeWithDefaults

`func NewSymbolChangeWithDefaults() *SymbolChange`

NewSymbolChangeWithDefaults instantiates a new SymbolChange object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFromDate

`func (o *SymbolChange) GetFromDate() string`

GetFromDate returns the FromDate field if non-nil, zero value otherwise.

### GetFromDateOk

`func (o *SymbolChange) GetFromDateOk() (*string, bool)`

GetFromDateOk returns a tuple with the FromDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFromDate

`func (o *SymbolChange) SetFromDate(v string)`

SetFromDate sets FromDate field to given value.

### HasFromDate

`func (o *SymbolChange) HasFromDate() bool`

HasFromDate returns a boolean if a field has been set.

### GetToDate

`func (o *SymbolChange) GetToDate() string`

GetToDate returns the ToDate field if non-nil, zero value otherwise.

### GetToDateOk

`func (o *SymbolChange) GetToDateOk() (*string, bool)`

GetToDateOk returns a tuple with the ToDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToDate

`func (o *SymbolChange) SetToDate(v string)`

SetToDate sets ToDate field to given value.

### HasToDate

`func (o *SymbolChange) HasToDate() bool`

HasToDate returns a boolean if a field has been set.

### GetData

`func (o *SymbolChange) GetData() []SymbolChangeInfo`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *SymbolChange) GetDataOk() (*[]SymbolChangeInfo, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *SymbolChange) SetData(v []SymbolChangeInfo)`

SetData sets Data field to given value.

### HasData

`func (o *SymbolChange) HasData() bool`

HasData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


