# SocialSentiment

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Symbol** | Pointer to **string** | Company symbol. | [optional] 
**Reddit** | Pointer to **[]map[string]interface{}** | Reddit sentiment. | [optional] 
**Twitter** | Pointer to **[]map[string]interface{}** | Twitter sentiment. | [optional] 

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

### GetReddit

`func (o *SocialSentiment) GetReddit() []map[string]interface{}`

GetReddit returns the Reddit field if non-nil, zero value otherwise.

### GetRedditOk

`func (o *SocialSentiment) GetRedditOk() (*[]map[string]interface{}, bool)`

GetRedditOk returns a tuple with the Reddit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReddit

`func (o *SocialSentiment) SetReddit(v []map[string]interface{})`

SetReddit sets Reddit field to given value.

### HasReddit

`func (o *SocialSentiment) HasReddit() bool`

HasReddit returns a boolean if a field has been set.

### GetTwitter

`func (o *SocialSentiment) GetTwitter() []map[string]interface{}`

GetTwitter returns the Twitter field if non-nil, zero value otherwise.

### GetTwitterOk

`func (o *SocialSentiment) GetTwitterOk() (*[]map[string]interface{}, bool)`

GetTwitterOk returns a tuple with the Twitter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTwitter

`func (o *SocialSentiment) SetTwitter(v []map[string]interface{})`

SetTwitter sets Twitter field to given value.

### HasTwitter

`func (o *SocialSentiment) HasTwitter() bool`

HasTwitter returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


