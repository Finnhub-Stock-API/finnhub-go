# SentimentContent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Mention** | Pointer to **int64** | Number of mentions | [optional] 
**PositiveMention** | Pointer to **int64** | Number of positive mentions | [optional] 
**NegativeMention** | Pointer to **int64** | Number of negative mentions | [optional] 
**PositiveScore** | Pointer to **float32** | Positive score. Range 0-1 | [optional] 
**NegativeScore** | Pointer to **float32** | Negative score. Range 0-1 | [optional] 
**Score** | Pointer to **float32** | Final score. Range: -1 to 1 with 1 is very positive and -1 is very negative | [optional] 
**AtTime** | Pointer to **string** | Period. | [optional] 

## Methods

### NewSentimentContent

`func NewSentimentContent() *SentimentContent`

NewSentimentContent instantiates a new SentimentContent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSentimentContentWithDefaults

`func NewSentimentContentWithDefaults() *SentimentContent`

NewSentimentContentWithDefaults instantiates a new SentimentContent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMention

`func (o *SentimentContent) GetMention() int64`

GetMention returns the Mention field if non-nil, zero value otherwise.

### GetMentionOk

`func (o *SentimentContent) GetMentionOk() (*int64, bool)`

GetMentionOk returns a tuple with the Mention field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMention

`func (o *SentimentContent) SetMention(v int64)`

SetMention sets Mention field to given value.

### HasMention

`func (o *SentimentContent) HasMention() bool`

HasMention returns a boolean if a field has been set.

### GetPositiveMention

`func (o *SentimentContent) GetPositiveMention() int64`

GetPositiveMention returns the PositiveMention field if non-nil, zero value otherwise.

### GetPositiveMentionOk

`func (o *SentimentContent) GetPositiveMentionOk() (*int64, bool)`

GetPositiveMentionOk returns a tuple with the PositiveMention field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPositiveMention

`func (o *SentimentContent) SetPositiveMention(v int64)`

SetPositiveMention sets PositiveMention field to given value.

### HasPositiveMention

`func (o *SentimentContent) HasPositiveMention() bool`

HasPositiveMention returns a boolean if a field has been set.

### GetNegativeMention

`func (o *SentimentContent) GetNegativeMention() int64`

GetNegativeMention returns the NegativeMention field if non-nil, zero value otherwise.

### GetNegativeMentionOk

`func (o *SentimentContent) GetNegativeMentionOk() (*int64, bool)`

GetNegativeMentionOk returns a tuple with the NegativeMention field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNegativeMention

`func (o *SentimentContent) SetNegativeMention(v int64)`

SetNegativeMention sets NegativeMention field to given value.

### HasNegativeMention

`func (o *SentimentContent) HasNegativeMention() bool`

HasNegativeMention returns a boolean if a field has been set.

### GetPositiveScore

`func (o *SentimentContent) GetPositiveScore() float32`

GetPositiveScore returns the PositiveScore field if non-nil, zero value otherwise.

### GetPositiveScoreOk

`func (o *SentimentContent) GetPositiveScoreOk() (*float32, bool)`

GetPositiveScoreOk returns a tuple with the PositiveScore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPositiveScore

`func (o *SentimentContent) SetPositiveScore(v float32)`

SetPositiveScore sets PositiveScore field to given value.

### HasPositiveScore

`func (o *SentimentContent) HasPositiveScore() bool`

HasPositiveScore returns a boolean if a field has been set.

### GetNegativeScore

`func (o *SentimentContent) GetNegativeScore() float32`

GetNegativeScore returns the NegativeScore field if non-nil, zero value otherwise.

### GetNegativeScoreOk

`func (o *SentimentContent) GetNegativeScoreOk() (*float32, bool)`

GetNegativeScoreOk returns a tuple with the NegativeScore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNegativeScore

`func (o *SentimentContent) SetNegativeScore(v float32)`

SetNegativeScore sets NegativeScore field to given value.

### HasNegativeScore

`func (o *SentimentContent) HasNegativeScore() bool`

HasNegativeScore returns a boolean if a field has been set.

### GetScore

`func (o *SentimentContent) GetScore() float32`

GetScore returns the Score field if non-nil, zero value otherwise.

### GetScoreOk

`func (o *SentimentContent) GetScoreOk() (*float32, bool)`

GetScoreOk returns a tuple with the Score field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScore

`func (o *SentimentContent) SetScore(v float32)`

SetScore sets Score field to given value.

### HasScore

`func (o *SentimentContent) HasScore() bool`

HasScore returns a boolean if a field has been set.

### GetAtTime

`func (o *SentimentContent) GetAtTime() string`

GetAtTime returns the AtTime field if non-nil, zero value otherwise.

### GetAtTimeOk

`func (o *SentimentContent) GetAtTimeOk() (*string, bool)`

GetAtTimeOk returns a tuple with the AtTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAtTime

`func (o *SentimentContent) SetAtTime(v string)`

SetAtTime sets AtTime field to given value.

### HasAtTime

`func (o *SentimentContent) HasAtTime() bool`

HasAtTime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


