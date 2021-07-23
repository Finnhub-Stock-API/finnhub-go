# TranscriptContent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | Speaker&#39;s name | [optional] 
**Speech** | Pointer to **[]string** | Speaker&#39;s speech | [optional] 
**Session** | Pointer to **string** | Earnings calls section (management discussion or Q&amp;A) | [optional] 

## Methods

### NewTranscriptContent

`func NewTranscriptContent() *TranscriptContent`

NewTranscriptContent instantiates a new TranscriptContent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTranscriptContentWithDefaults

`func NewTranscriptContentWithDefaults() *TranscriptContent`

NewTranscriptContentWithDefaults instantiates a new TranscriptContent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *TranscriptContent) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TranscriptContent) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TranscriptContent) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *TranscriptContent) HasName() bool`

HasName returns a boolean if a field has been set.

### GetSpeech

`func (o *TranscriptContent) GetSpeech() []string`

GetSpeech returns the Speech field if non-nil, zero value otherwise.

### GetSpeechOk

`func (o *TranscriptContent) GetSpeechOk() (*[]string, bool)`

GetSpeechOk returns a tuple with the Speech field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpeech

`func (o *TranscriptContent) SetSpeech(v []string)`

SetSpeech sets Speech field to given value.

### HasSpeech

`func (o *TranscriptContent) HasSpeech() bool`

HasSpeech returns a boolean if a field has been set.

### GetSession

`func (o *TranscriptContent) GetSession() string`

GetSession returns the Session field if non-nil, zero value otherwise.

### GetSessionOk

`func (o *TranscriptContent) GetSessionOk() (*string, bool)`

GetSessionOk returns a tuple with the Session field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSession

`func (o *TranscriptContent) SetSession(v string)`

SetSession sets Session field to given value.

### HasSession

`func (o *TranscriptContent) HasSession() bool`

HasSession returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


