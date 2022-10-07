# IsinChange

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FromDate** | Pointer to **string** | From date. | [optional] 
**ToDate** | Pointer to **string** | To date. | [optional] 
**Data** | Pointer to [**[]IsinChangeInfo**](IsinChangeInfo.md) | Array of ISIN change events. | [optional] 

## Methods

### NewIsinChange

`func NewIsinChange() *IsinChange`

NewIsinChange instantiates a new IsinChange object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIsinChangeWithDefaults

`func NewIsinChangeWithDefaults() *IsinChange`

NewIsinChangeWithDefaults instantiates a new IsinChange object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFromDate

`func (o *IsinChange) GetFromDate() string`

GetFromDate returns the FromDate field if non-nil, zero value otherwise.

### GetFromDateOk

`func (o *IsinChange) GetFromDateOk() (*string, bool)`

GetFromDateOk returns a tuple with the FromDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFromDate

`func (o *IsinChange) SetFromDate(v string)`

SetFromDate sets FromDate field to given value.

### HasFromDate

`func (o *IsinChange) HasFromDate() bool`

HasFromDate returns a boolean if a field has been set.

### GetToDate

`func (o *IsinChange) GetToDate() string`

GetToDate returns the ToDate field if non-nil, zero value otherwise.

### GetToDateOk

`func (o *IsinChange) GetToDateOk() (*string, bool)`

GetToDateOk returns a tuple with the ToDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToDate

`func (o *IsinChange) SetToDate(v string)`

SetToDate sets ToDate field to given value.

### HasToDate

`func (o *IsinChange) HasToDate() bool`

HasToDate returns a boolean if a field has been set.

### GetData

`func (o *IsinChange) GetData() []IsinChangeInfo`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *IsinChange) GetDataOk() (*[]IsinChangeInfo, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *IsinChange) SetData(v []IsinChangeInfo)`

SetData sets Data field to given value.

### HasData

`func (o *IsinChange) HasData() bool`

HasData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


