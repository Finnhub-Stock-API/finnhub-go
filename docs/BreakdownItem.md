# BreakdownItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccessNumber** | Pointer to **string** | Access number of the report from which the data is sourced. | [optional] 
**Breakdown** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewBreakdownItem

`func NewBreakdownItem() *BreakdownItem`

NewBreakdownItem instantiates a new BreakdownItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBreakdownItemWithDefaults

`func NewBreakdownItemWithDefaults() *BreakdownItem`

NewBreakdownItemWithDefaults instantiates a new BreakdownItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccessNumber

`func (o *BreakdownItem) GetAccessNumber() string`

GetAccessNumber returns the AccessNumber field if non-nil, zero value otherwise.

### GetAccessNumberOk

`func (o *BreakdownItem) GetAccessNumberOk() (*string, bool)`

GetAccessNumberOk returns a tuple with the AccessNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessNumber

`func (o *BreakdownItem) SetAccessNumber(v string)`

SetAccessNumber sets AccessNumber field to given value.

### HasAccessNumber

`func (o *BreakdownItem) HasAccessNumber() bool`

HasAccessNumber returns a boolean if a field has been set.

### GetBreakdown

`func (o *BreakdownItem) GetBreakdown() map[string]interface{}`

GetBreakdown returns the Breakdown field if non-nil, zero value otherwise.

### GetBreakdownOk

`func (o *BreakdownItem) GetBreakdownOk() (*map[string]interface{}, bool)`

GetBreakdownOk returns a tuple with the Breakdown field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBreakdown

`func (o *BreakdownItem) SetBreakdown(v map[string]interface{})`

SetBreakdown sets Breakdown field to given value.

### HasBreakdown

`func (o *BreakdownItem) HasBreakdown() bool`

HasBreakdown returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


