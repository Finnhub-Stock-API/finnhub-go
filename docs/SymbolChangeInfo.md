# SymbolChangeInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AtDate** | Pointer to **string** | Event&#39;s date. | [optional] 
**OldSymbol** | Pointer to **string** | Old symbol. | [optional] 
**NewSymbol** | Pointer to **string** | New symbol. | [optional] 

## Methods

### NewSymbolChangeInfo

`func NewSymbolChangeInfo() *SymbolChangeInfo`

NewSymbolChangeInfo instantiates a new SymbolChangeInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSymbolChangeInfoWithDefaults

`func NewSymbolChangeInfoWithDefaults() *SymbolChangeInfo`

NewSymbolChangeInfoWithDefaults instantiates a new SymbolChangeInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAtDate

`func (o *SymbolChangeInfo) GetAtDate() string`

GetAtDate returns the AtDate field if non-nil, zero value otherwise.

### GetAtDateOk

`func (o *SymbolChangeInfo) GetAtDateOk() (*string, bool)`

GetAtDateOk returns a tuple with the AtDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAtDate

`func (o *SymbolChangeInfo) SetAtDate(v string)`

SetAtDate sets AtDate field to given value.

### HasAtDate

`func (o *SymbolChangeInfo) HasAtDate() bool`

HasAtDate returns a boolean if a field has been set.

### GetOldSymbol

`func (o *SymbolChangeInfo) GetOldSymbol() string`

GetOldSymbol returns the OldSymbol field if non-nil, zero value otherwise.

### GetOldSymbolOk

`func (o *SymbolChangeInfo) GetOldSymbolOk() (*string, bool)`

GetOldSymbolOk returns a tuple with the OldSymbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOldSymbol

`func (o *SymbolChangeInfo) SetOldSymbol(v string)`

SetOldSymbol sets OldSymbol field to given value.

### HasOldSymbol

`func (o *SymbolChangeInfo) HasOldSymbol() bool`

HasOldSymbol returns a boolean if a field has been set.

### GetNewSymbol

`func (o *SymbolChangeInfo) GetNewSymbol() string`

GetNewSymbol returns the NewSymbol field if non-nil, zero value otherwise.

### GetNewSymbolOk

`func (o *SymbolChangeInfo) GetNewSymbolOk() (*string, bool)`

GetNewSymbolOk returns a tuple with the NewSymbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewSymbol

`func (o *SymbolChangeInfo) SetNewSymbol(v string)`

SetNewSymbol sets NewSymbol field to given value.

### HasNewSymbol

`func (o *SymbolChangeInfo) HasNewSymbol() bool`

HasNewSymbol returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


