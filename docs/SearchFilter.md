# SearchFilter

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Filter id, use with respective field in search query body. | [optional] 
**Name** | Pointer to **string** | Display name. | [optional] 

## Methods

### NewSearchFilter

`func NewSearchFilter() *SearchFilter`

NewSearchFilter instantiates a new SearchFilter object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSearchFilterWithDefaults

`func NewSearchFilterWithDefaults() *SearchFilter`

NewSearchFilterWithDefaults instantiates a new SearchFilter object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *SearchFilter) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SearchFilter) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SearchFilter) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *SearchFilter) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *SearchFilter) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SearchFilter) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SearchFilter) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *SearchFilter) HasName() bool`

HasName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


