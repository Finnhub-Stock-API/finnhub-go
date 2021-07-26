# EarningsCallTranscripts

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Symbol** | Pointer to **string** | Company symbol. | [optional] 
**Transcript** | Pointer to [**[]TranscriptContent**](TranscriptContent.md) | Transcript content. | [optional] 
**Participant** | Pointer to [**[]TranscriptParticipant**](TranscriptParticipant.md) | Participant list | [optional] 
**Audio** | Pointer to **string** | Audio link. | [optional] 
**Id** | Pointer to **string** | Transcript&#39;s ID. | [optional] 
**Title** | Pointer to **string** | Title. | [optional] 
**Time** | Pointer to **string** | Time of the event. | [optional] 
**Year** | Pointer to **int64** | Year of earnings result in the case of earnings call transcript. | [optional] 
**Quarter** | Pointer to **int64** | Quarter of earnings result in the case of earnings call transcript. | [optional] 

## Methods

### NewEarningsCallTranscripts

`func NewEarningsCallTranscripts() *EarningsCallTranscripts`

NewEarningsCallTranscripts instantiates a new EarningsCallTranscripts object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEarningsCallTranscriptsWithDefaults

`func NewEarningsCallTranscriptsWithDefaults() *EarningsCallTranscripts`

NewEarningsCallTranscriptsWithDefaults instantiates a new EarningsCallTranscripts object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSymbol

`func (o *EarningsCallTranscripts) GetSymbol() string`

GetSymbol returns the Symbol field if non-nil, zero value otherwise.

### GetSymbolOk

`func (o *EarningsCallTranscripts) GetSymbolOk() (*string, bool)`

GetSymbolOk returns a tuple with the Symbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSymbol

`func (o *EarningsCallTranscripts) SetSymbol(v string)`

SetSymbol sets Symbol field to given value.

### HasSymbol

`func (o *EarningsCallTranscripts) HasSymbol() bool`

HasSymbol returns a boolean if a field has been set.

### GetTranscript

`func (o *EarningsCallTranscripts) GetTranscript() []TranscriptContent`

GetTranscript returns the Transcript field if non-nil, zero value otherwise.

### GetTranscriptOk

`func (o *EarningsCallTranscripts) GetTranscriptOk() (*[]TranscriptContent, bool)`

GetTranscriptOk returns a tuple with the Transcript field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTranscript

`func (o *EarningsCallTranscripts) SetTranscript(v []TranscriptContent)`

SetTranscript sets Transcript field to given value.

### HasTranscript

`func (o *EarningsCallTranscripts) HasTranscript() bool`

HasTranscript returns a boolean if a field has been set.

### GetParticipant

`func (o *EarningsCallTranscripts) GetParticipant() []TranscriptParticipant`

GetParticipant returns the Participant field if non-nil, zero value otherwise.

### GetParticipantOk

`func (o *EarningsCallTranscripts) GetParticipantOk() (*[]TranscriptParticipant, bool)`

GetParticipantOk returns a tuple with the Participant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParticipant

`func (o *EarningsCallTranscripts) SetParticipant(v []TranscriptParticipant)`

SetParticipant sets Participant field to given value.

### HasParticipant

`func (o *EarningsCallTranscripts) HasParticipant() bool`

HasParticipant returns a boolean if a field has been set.

### GetAudio

`func (o *EarningsCallTranscripts) GetAudio() string`

GetAudio returns the Audio field if non-nil, zero value otherwise.

### GetAudioOk

`func (o *EarningsCallTranscripts) GetAudioOk() (*string, bool)`

GetAudioOk returns a tuple with the Audio field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAudio

`func (o *EarningsCallTranscripts) SetAudio(v string)`

SetAudio sets Audio field to given value.

### HasAudio

`func (o *EarningsCallTranscripts) HasAudio() bool`

HasAudio returns a boolean if a field has been set.

### GetId

`func (o *EarningsCallTranscripts) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *EarningsCallTranscripts) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *EarningsCallTranscripts) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *EarningsCallTranscripts) HasId() bool`

HasId returns a boolean if a field has been set.

### GetTitle

`func (o *EarningsCallTranscripts) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *EarningsCallTranscripts) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *EarningsCallTranscripts) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *EarningsCallTranscripts) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetTime

`func (o *EarningsCallTranscripts) GetTime() string`

GetTime returns the Time field if non-nil, zero value otherwise.

### GetTimeOk

`func (o *EarningsCallTranscripts) GetTimeOk() (*string, bool)`

GetTimeOk returns a tuple with the Time field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTime

`func (o *EarningsCallTranscripts) SetTime(v string)`

SetTime sets Time field to given value.

### HasTime

`func (o *EarningsCallTranscripts) HasTime() bool`

HasTime returns a boolean if a field has been set.

### GetYear

`func (o *EarningsCallTranscripts) GetYear() int64`

GetYear returns the Year field if non-nil, zero value otherwise.

### GetYearOk

`func (o *EarningsCallTranscripts) GetYearOk() (*int64, bool)`

GetYearOk returns a tuple with the Year field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetYear

`func (o *EarningsCallTranscripts) SetYear(v int64)`

SetYear sets Year field to given value.

### HasYear

`func (o *EarningsCallTranscripts) HasYear() bool`

HasYear returns a boolean if a field has been set.

### GetQuarter

`func (o *EarningsCallTranscripts) GetQuarter() int64`

GetQuarter returns the Quarter field if non-nil, zero value otherwise.

### GetQuarterOk

`func (o *EarningsCallTranscripts) GetQuarterOk() (*int64, bool)`

GetQuarterOk returns a tuple with the Quarter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuarter

`func (o *EarningsCallTranscripts) SetQuarter(v int64)`

SetQuarter sets Quarter field to given value.

### HasQuarter

`func (o *EarningsCallTranscripts) HasQuarter() bool`

HasQuarter returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


