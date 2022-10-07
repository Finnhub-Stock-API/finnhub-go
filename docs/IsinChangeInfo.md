# IsinChangeInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AtDate** | Pointer to **string** | Event&#39;s date. | [optional] 
**OldIsin** | Pointer to **string** | Old ISIN. | [optional] 
**NewIsin** | Pointer to **string** | New ISIN. | [optional] 

## Methods

### NewIsinChangeInfo

`func NewIsinChangeInfo() *IsinChangeInfo`

NewIsinChangeInfo instantiates a new IsinChangeInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIsinChangeInfoWithDefaults

`func NewIsinChangeInfoWithDefaults() *IsinChangeInfo`

NewIsinChangeInfoWithDefaults instantiates a new IsinChangeInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAtDate

`func (o *IsinChangeInfo) GetAtDate() string`

GetAtDate returns the AtDate field if non-nil, zero value otherwise.

### GetAtDateOk

`func (o *IsinChangeInfo) GetAtDateOk() (*string, bool)`

GetAtDateOk returns a tuple with the AtDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAtDate

`func (o *IsinChangeInfo) SetAtDate(v string)`

SetAtDate sets AtDate field to given value.

### HasAtDate

`func (o *IsinChangeInfo) HasAtDate() bool`

HasAtDate returns a boolean if a field has been set.

### GetOldIsin

`func (o *IsinChangeInfo) GetOldIsin() string`

GetOldIsin returns the OldIsin field if non-nil, zero value otherwise.

### GetOldIsinOk

`func (o *IsinChangeInfo) GetOldIsinOk() (*string, bool)`

GetOldIsinOk returns a tuple with the OldIsin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOldIsin

`func (o *IsinChangeInfo) SetOldIsin(v string)`

SetOldIsin sets OldIsin field to given value.

### HasOldIsin

`func (o *IsinChangeInfo) HasOldIsin() bool`

HasOldIsin returns a boolean if a field has been set.

### GetNewIsin

`func (o *IsinChangeInfo) GetNewIsin() string`

GetNewIsin returns the NewIsin field if non-nil, zero value otherwise.

### GetNewIsinOk

`func (o *IsinChangeInfo) GetNewIsinOk() (*string, bool)`

GetNewIsinOk returns a tuple with the NewIsin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewIsin

`func (o *IsinChangeInfo) SetNewIsin(v string)`

SetNewIsin sets NewIsin field to given value.

### HasNewIsin

`func (o *IsinChangeInfo) HasNewIsin() bool`

HasNewIsin returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


