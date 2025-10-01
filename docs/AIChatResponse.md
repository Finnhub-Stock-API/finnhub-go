# AIChatResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ChatId** | Pointer to **string** | Chat ID. | [optional] 
**Content** | Pointer to **string** | Response text. | [optional] 
**QuerySummary** | Pointer to **string** | Query summary | [optional] 
**RelatedQueries** | Pointer to **[]map[string]interface{}** | Related queries. | [optional] 
**Tickers** | Pointer to **[]map[string]interface{}** | List of tickers mentioned. | [optional] 
**Sources** | Pointer to **[]map[string]interface{}** | Sources. | [optional] 
**Widgets** | Pointer to **[]map[string]interface{}** | Widgets. | [optional] 

## Methods

### NewAIChatResponse

`func NewAIChatResponse() *AIChatResponse`

NewAIChatResponse instantiates a new AIChatResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAIChatResponseWithDefaults

`func NewAIChatResponseWithDefaults() *AIChatResponse`

NewAIChatResponseWithDefaults instantiates a new AIChatResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChatId

`func (o *AIChatResponse) GetChatId() string`

GetChatId returns the ChatId field if non-nil, zero value otherwise.

### GetChatIdOk

`func (o *AIChatResponse) GetChatIdOk() (*string, bool)`

GetChatIdOk returns a tuple with the ChatId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChatId

`func (o *AIChatResponse) SetChatId(v string)`

SetChatId sets ChatId field to given value.

### HasChatId

`func (o *AIChatResponse) HasChatId() bool`

HasChatId returns a boolean if a field has been set.

### GetContent

`func (o *AIChatResponse) GetContent() string`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *AIChatResponse) GetContentOk() (*string, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *AIChatResponse) SetContent(v string)`

SetContent sets Content field to given value.

### HasContent

`func (o *AIChatResponse) HasContent() bool`

HasContent returns a boolean if a field has been set.

### GetQuerySummary

`func (o *AIChatResponse) GetQuerySummary() string`

GetQuerySummary returns the QuerySummary field if non-nil, zero value otherwise.

### GetQuerySummaryOk

`func (o *AIChatResponse) GetQuerySummaryOk() (*string, bool)`

GetQuerySummaryOk returns a tuple with the QuerySummary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuerySummary

`func (o *AIChatResponse) SetQuerySummary(v string)`

SetQuerySummary sets QuerySummary field to given value.

### HasQuerySummary

`func (o *AIChatResponse) HasQuerySummary() bool`

HasQuerySummary returns a boolean if a field has been set.

### GetRelatedQueries

`func (o *AIChatResponse) GetRelatedQueries() []map[string]interface{}`

GetRelatedQueries returns the RelatedQueries field if non-nil, zero value otherwise.

### GetRelatedQueriesOk

`func (o *AIChatResponse) GetRelatedQueriesOk() (*[]map[string]interface{}, bool)`

GetRelatedQueriesOk returns a tuple with the RelatedQueries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelatedQueries

`func (o *AIChatResponse) SetRelatedQueries(v []map[string]interface{})`

SetRelatedQueries sets RelatedQueries field to given value.

### HasRelatedQueries

`func (o *AIChatResponse) HasRelatedQueries() bool`

HasRelatedQueries returns a boolean if a field has been set.

### GetTickers

`func (o *AIChatResponse) GetTickers() []map[string]interface{}`

GetTickers returns the Tickers field if non-nil, zero value otherwise.

### GetTickersOk

`func (o *AIChatResponse) GetTickersOk() (*[]map[string]interface{}, bool)`

GetTickersOk returns a tuple with the Tickers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTickers

`func (o *AIChatResponse) SetTickers(v []map[string]interface{})`

SetTickers sets Tickers field to given value.

### HasTickers

`func (o *AIChatResponse) HasTickers() bool`

HasTickers returns a boolean if a field has been set.

### GetSources

`func (o *AIChatResponse) GetSources() []map[string]interface{}`

GetSources returns the Sources field if non-nil, zero value otherwise.

### GetSourcesOk

`func (o *AIChatResponse) GetSourcesOk() (*[]map[string]interface{}, bool)`

GetSourcesOk returns a tuple with the Sources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSources

`func (o *AIChatResponse) SetSources(v []map[string]interface{})`

SetSources sets Sources field to given value.

### HasSources

`func (o *AIChatResponse) HasSources() bool`

HasSources returns a boolean if a field has been set.

### GetWidgets

`func (o *AIChatResponse) GetWidgets() []map[string]interface{}`

GetWidgets returns the Widgets field if non-nil, zero value otherwise.

### GetWidgetsOk

`func (o *AIChatResponse) GetWidgetsOk() (*[]map[string]interface{}, bool)`

GetWidgetsOk returns a tuple with the Widgets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWidgets

`func (o *AIChatResponse) SetWidgets(v []map[string]interface{})`

SetWidgets sets Widgets field to given value.

### HasWidgets

`func (o *AIChatResponse) HasWidgets() bool`

HasWidgets returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


