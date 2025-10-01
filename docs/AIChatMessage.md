# AIChatMessage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Role** | Pointer to **string** | Role system/user | [optional] 
**Content** | Pointer to **string** | Content | [optional] 

## Methods

### NewAIChatMessage

`func NewAIChatMessage() *AIChatMessage`

NewAIChatMessage instantiates a new AIChatMessage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAIChatMessageWithDefaults

`func NewAIChatMessageWithDefaults() *AIChatMessage`

NewAIChatMessageWithDefaults instantiates a new AIChatMessage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRole

`func (o *AIChatMessage) GetRole() string`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *AIChatMessage) GetRoleOk() (*string, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *AIChatMessage) SetRole(v string)`

SetRole sets Role field to given value.

### HasRole

`func (o *AIChatMessage) HasRole() bool`

HasRole returns a boolean if a field has been set.

### GetContent

`func (o *AIChatMessage) GetContent() string`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *AIChatMessage) GetContentOk() (*string, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *AIChatMessage) SetContent(v string)`

SetContent sets Content field to given value.

### HasContent

`func (o *AIChatMessage) HasContent() bool`

HasContent returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


