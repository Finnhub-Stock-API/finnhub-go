# SocialSentiment

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Symbol** | Pointer to **string** | Company symbol. | [optional] 
**Data** | Pointer to [**[]SentimentContent**](SentimentContent.md) | Sentiment data. | [optional] 

## Methods

### NewSocialSentiment

`func NewSocialSentiment() *SocialSentiment`

NewSocialSentiment instantiates a new SocialSentiment object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSocialSentimentWithDefaults

`func NewSocialSentimentWithDefaults() *SocialSentiment`

NewSocialSentimentWithDefaults instantiates a new SocialSentiment object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSymbol

`func (o *SocialSentiment) GetSymbol() string`

GetSymbol returns the Symbol field if non-nil, zero value otherwise.

### GetSymbolOk

`func (o *SocialSentiment) GetSymbolOk() (*string, bool)`

GetSymbolOk returns a tuple with the Symbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSymbol

`func (o *SocialSentiment) SetSymbol(v string)`

SetSymbol sets Symbol field to given value.

### HasSymbol

`func (o *SocialSentiment) HasSymbol() bool`

HasSymbol returns a boolean if a field has been set.

### GetData

`func (o *SocialSentiment) GetData() []SentimentContent`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *SocialSentiment) GetDataOk() (*[]SentimentContent, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *SocialSentiment) SetData(v []SentimentContent)`

SetData sets Data field to given value.

### HasData

`func (o *SocialSentiment) HasData() bool`

HasData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


