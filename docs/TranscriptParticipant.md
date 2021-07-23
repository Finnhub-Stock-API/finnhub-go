# TranscriptParticipant

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | Participant&#39;s name | [optional] 
**Description** | Pointer to **string** | Participant&#39;s description | [optional] 
**Role** | Pointer to **string** | Whether the speak is a company&#39;s executive or an analyst | [optional] 

## Methods

### NewTranscriptParticipant

`func NewTranscriptParticipant() *TranscriptParticipant`

NewTranscriptParticipant instantiates a new TranscriptParticipant object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTranscriptParticipantWithDefaults

`func NewTranscriptParticipantWithDefaults() *TranscriptParticipant`

NewTranscriptParticipantWithDefaults instantiates a new TranscriptParticipant object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *TranscriptParticipant) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TranscriptParticipant) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TranscriptParticipant) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *TranscriptParticipant) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *TranscriptParticipant) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *TranscriptParticipant) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *TranscriptParticipant) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *TranscriptParticipant) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetRole

`func (o *TranscriptParticipant) GetRole() string`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *TranscriptParticipant) GetRoleOk() (*string, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *TranscriptParticipant) SetRole(v string)`

SetRole sets Role field to given value.

### HasRole

`func (o *TranscriptParticipant) HasRole() bool`

HasRole returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


