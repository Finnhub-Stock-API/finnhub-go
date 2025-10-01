# AIChatBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Messages** | [**[]AIChatMessage**](AIChatMessage.md) | Messages | 
**Stream** | Pointer to **bool** | Stream responses | [optional] 

## Methods

### NewAIChatBody

`func NewAIChatBody(messages []AIChatMessage, ) *AIChatBody`

NewAIChatBody instantiates a new AIChatBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAIChatBodyWithDefaults

`func NewAIChatBodyWithDefaults() *AIChatBody`

NewAIChatBodyWithDefaults instantiates a new AIChatBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessages

`func (o *AIChatBody) GetMessages() []AIChatMessage`

GetMessages returns the Messages field if non-nil, zero value otherwise.

### GetMessagesOk

`func (o *AIChatBody) GetMessagesOk() (*[]AIChatMessage, bool)`

GetMessagesOk returns a tuple with the Messages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessages

`func (o *AIChatBody) SetMessages(v []AIChatMessage)`

SetMessages sets Messages field to given value.


### GetStream

`func (o *AIChatBody) GetStream() bool`

GetStream returns the Stream field if non-nil, zero value otherwise.

### GetStreamOk

`func (o *AIChatBody) GetStreamOk() (*bool, bool)`

GetStreamOk returns a tuple with the Stream field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStream

`func (o *AIChatBody) SetStream(v bool)`

SetStream sets Stream field to given value.

### HasStream

`func (o *AIChatBody) HasStream() bool`

HasStream returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


