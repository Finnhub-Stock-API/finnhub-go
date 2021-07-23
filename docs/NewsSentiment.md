# NewsSentiment

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Buzz** | Pointer to [**CompanyNewsStatistics**](CompanyNewsStatistics.md) |  | [optional] 
**CompanyNewsScore** | Pointer to **float32** | News score. | [optional] 
**SectorAverageBullishPercent** | Pointer to **float32** | Sector average bullish percent. | [optional] 
**SectorAverageNewsScore** | Pointer to **float32** | Sectore average score. | [optional] 
**Sentiment** | Pointer to [**Sentiment**](Sentiment.md) |  | [optional] 
**Symbol** | Pointer to **string** | Requested symbol. | [optional] 

## Methods

### NewNewsSentiment

`func NewNewsSentiment() *NewsSentiment`

NewNewsSentiment instantiates a new NewsSentiment object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNewsSentimentWithDefaults

`func NewNewsSentimentWithDefaults() *NewsSentiment`

NewNewsSentimentWithDefaults instantiates a new NewsSentiment object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBuzz

`func (o *NewsSentiment) GetBuzz() CompanyNewsStatistics`

GetBuzz returns the Buzz field if non-nil, zero value otherwise.

### GetBuzzOk

`func (o *NewsSentiment) GetBuzzOk() (*CompanyNewsStatistics, bool)`

GetBuzzOk returns a tuple with the Buzz field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuzz

`func (o *NewsSentiment) SetBuzz(v CompanyNewsStatistics)`

SetBuzz sets Buzz field to given value.

### HasBuzz

`func (o *NewsSentiment) HasBuzz() bool`

HasBuzz returns a boolean if a field has been set.

### GetCompanyNewsScore

`func (o *NewsSentiment) GetCompanyNewsScore() float32`

GetCompanyNewsScore returns the CompanyNewsScore field if non-nil, zero value otherwise.

### GetCompanyNewsScoreOk

`func (o *NewsSentiment) GetCompanyNewsScoreOk() (*float32, bool)`

GetCompanyNewsScoreOk returns a tuple with the CompanyNewsScore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyNewsScore

`func (o *NewsSentiment) SetCompanyNewsScore(v float32)`

SetCompanyNewsScore sets CompanyNewsScore field to given value.

### HasCompanyNewsScore

`func (o *NewsSentiment) HasCompanyNewsScore() bool`

HasCompanyNewsScore returns a boolean if a field has been set.

### GetSectorAverageBullishPercent

`func (o *NewsSentiment) GetSectorAverageBullishPercent() float32`

GetSectorAverageBullishPercent returns the SectorAverageBullishPercent field if non-nil, zero value otherwise.

### GetSectorAverageBullishPercentOk

`func (o *NewsSentiment) GetSectorAverageBullishPercentOk() (*float32, bool)`

GetSectorAverageBullishPercentOk returns a tuple with the SectorAverageBullishPercent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSectorAverageBullishPercent

`func (o *NewsSentiment) SetSectorAverageBullishPercent(v float32)`

SetSectorAverageBullishPercent sets SectorAverageBullishPercent field to given value.

### HasSectorAverageBullishPercent

`func (o *NewsSentiment) HasSectorAverageBullishPercent() bool`

HasSectorAverageBullishPercent returns a boolean if a field has been set.

### GetSectorAverageNewsScore

`func (o *NewsSentiment) GetSectorAverageNewsScore() float32`

GetSectorAverageNewsScore returns the SectorAverageNewsScore field if non-nil, zero value otherwise.

### GetSectorAverageNewsScoreOk

`func (o *NewsSentiment) GetSectorAverageNewsScoreOk() (*float32, bool)`

GetSectorAverageNewsScoreOk returns a tuple with the SectorAverageNewsScore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSectorAverageNewsScore

`func (o *NewsSentiment) SetSectorAverageNewsScore(v float32)`

SetSectorAverageNewsScore sets SectorAverageNewsScore field to given value.

### HasSectorAverageNewsScore

`func (o *NewsSentiment) HasSectorAverageNewsScore() bool`

HasSectorAverageNewsScore returns a boolean if a field has been set.

### GetSentiment

`func (o *NewsSentiment) GetSentiment() Sentiment`

GetSentiment returns the Sentiment field if non-nil, zero value otherwise.

### GetSentimentOk

`func (o *NewsSentiment) GetSentimentOk() (*Sentiment, bool)`

GetSentimentOk returns a tuple with the Sentiment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSentiment

`func (o *NewsSentiment) SetSentiment(v Sentiment)`

SetSentiment sets Sentiment field to given value.

### HasSentiment

`func (o *NewsSentiment) HasSentiment() bool`

HasSentiment returns a boolean if a field has been set.

### GetSymbol

`func (o *NewsSentiment) GetSymbol() string`

GetSymbol returns the Symbol field if non-nil, zero value otherwise.

### GetSymbolOk

`func (o *NewsSentiment) GetSymbolOk() (*string, bool)`

GetSymbolOk returns a tuple with the Symbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSymbol

`func (o *NewsSentiment) SetSymbol(v string)`

SetSymbol sets Symbol field to given value.

### HasSymbol

`func (o *NewsSentiment) HasSymbol() bool`

HasSymbol returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


