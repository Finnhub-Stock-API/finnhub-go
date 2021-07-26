# SymbolLookup

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Result** | Pointer to [**[]SymbolLookupInfo**](SymbolLookupInfo.md) | Array of search results. | [optional] 
**Count** | Pointer to **int64** | Number of results. | [optional] 

## Methods

### NewSymbolLookup

`func NewSymbolLookup() *SymbolLookup`

NewSymbolLookup instantiates a new SymbolLookup object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSymbolLookupWithDefaults

`func NewSymbolLookupWithDefaults() *SymbolLookup`

NewSymbolLookupWithDefaults instantiates a new SymbolLookup object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResult

`func (o *SymbolLookup) GetResult() []SymbolLookupInfo`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *SymbolLookup) GetResultOk() (*[]SymbolLookupInfo, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *SymbolLookup) SetResult(v []SymbolLookupInfo)`

SetResult sets Result field to given value.

### HasResult

`func (o *SymbolLookup) HasResult() bool`

HasResult returns a boolean if a field has been set.

### GetCount

`func (o *SymbolLookup) GetCount() int64`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *SymbolLookup) GetCountOk() (*int64, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *SymbolLookup) SetCount(v int64)`

SetCount sets Count field to given value.

### HasCount

`func (o *SymbolLookup) HasCount() bool`

HasCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


