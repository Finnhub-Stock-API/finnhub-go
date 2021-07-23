# SECSentimentAnalysis

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccessNumber** | Pointer to **string** | Access number. | [optional] 
**Symbol** | Pointer to **string** | Symbol. | [optional] 
**Cik** | Pointer to **string** | CIK. | [optional] 
**Sentiment** | Pointer to [**FilingSentiment**](FilingSentiment.md) |  | [optional] 

## Methods

### NewSECSentimentAnalysis

`func NewSECSentimentAnalysis() *SECSentimentAnalysis`

NewSECSentimentAnalysis instantiates a new SECSentimentAnalysis object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSECSentimentAnalysisWithDefaults

`func NewSECSentimentAnalysisWithDefaults() *SECSentimentAnalysis`

NewSECSentimentAnalysisWithDefaults instantiates a new SECSentimentAnalysis object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccessNumber

`func (o *SECSentimentAnalysis) GetAccessNumber() string`

GetAccessNumber returns the AccessNumber field if non-nil, zero value otherwise.

### GetAccessNumberOk

`func (o *SECSentimentAnalysis) GetAccessNumberOk() (*string, bool)`

GetAccessNumberOk returns a tuple with the AccessNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessNumber

`func (o *SECSentimentAnalysis) SetAccessNumber(v string)`

SetAccessNumber sets AccessNumber field to given value.

### HasAccessNumber

`func (o *SECSentimentAnalysis) HasAccessNumber() bool`

HasAccessNumber returns a boolean if a field has been set.

### GetSymbol

`func (o *SECSentimentAnalysis) GetSymbol() string`

GetSymbol returns the Symbol field if non-nil, zero value otherwise.

### GetSymbolOk

`func (o *SECSentimentAnalysis) GetSymbolOk() (*string, bool)`

GetSymbolOk returns a tuple with the Symbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSymbol

`func (o *SECSentimentAnalysis) SetSymbol(v string)`

SetSymbol sets Symbol field to given value.

### HasSymbol

`func (o *SECSentimentAnalysis) HasSymbol() bool`

HasSymbol returns a boolean if a field has been set.

### GetCik

`func (o *SECSentimentAnalysis) GetCik() string`

GetCik returns the Cik field if non-nil, zero value otherwise.

### GetCikOk

`func (o *SECSentimentAnalysis) GetCikOk() (*string, bool)`

GetCikOk returns a tuple with the Cik field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCik

`func (o *SECSentimentAnalysis) SetCik(v string)`

SetCik sets Cik field to given value.

### HasCik

`func (o *SECSentimentAnalysis) HasCik() bool`

HasCik returns a boolean if a field has been set.

### GetSentiment

`func (o *SECSentimentAnalysis) GetSentiment() FilingSentiment`

GetSentiment returns the Sentiment field if non-nil, zero value otherwise.

### GetSentimentOk

`func (o *SECSentimentAnalysis) GetSentimentOk() (*FilingSentiment, bool)`

GetSentimentOk returns a tuple with the Sentiment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSentiment

`func (o *SECSentimentAnalysis) SetSentiment(v FilingSentiment)`

SetSentiment sets Sentiment field to given value.

### HasSentiment

`func (o *SECSentimentAnalysis) HasSentiment() bool`

HasSentiment returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


