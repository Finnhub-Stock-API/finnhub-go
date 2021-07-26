# MarketNews

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Category** | Pointer to **string** | News category. | [optional] 
**Datetime** | Pointer to **int64** | Published time in UNIX timestamp. | [optional] 
**Headline** | Pointer to **string** | News headline. | [optional] 
**Id** | Pointer to **int64** | News ID. This value can be used for &lt;code&gt;minId&lt;/code&gt; params to get the latest news only. | [optional] 
**Image** | Pointer to **string** | Thumbnail image URL. | [optional] 
**Related** | Pointer to **string** | Related stocks and companies mentioned in the article. | [optional] 
**Source** | Pointer to **string** | News source. | [optional] 
**Summary** | Pointer to **string** | News summary. | [optional] 
**Url** | Pointer to **string** | URL of the original article. | [optional] 

## Methods

### NewMarketNews

`func NewMarketNews() *MarketNews`

NewMarketNews instantiates a new MarketNews object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMarketNewsWithDefaults

`func NewMarketNewsWithDefaults() *MarketNews`

NewMarketNewsWithDefaults instantiates a new MarketNews object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCategory

`func (o *MarketNews) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *MarketNews) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *MarketNews) SetCategory(v string)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *MarketNews) HasCategory() bool`

HasCategory returns a boolean if a field has been set.

### GetDatetime

`func (o *MarketNews) GetDatetime() int64`

GetDatetime returns the Datetime field if non-nil, zero value otherwise.

### GetDatetimeOk

`func (o *MarketNews) GetDatetimeOk() (*int64, bool)`

GetDatetimeOk returns a tuple with the Datetime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatetime

`func (o *MarketNews) SetDatetime(v int64)`

SetDatetime sets Datetime field to given value.

### HasDatetime

`func (o *MarketNews) HasDatetime() bool`

HasDatetime returns a boolean if a field has been set.

### GetHeadline

`func (o *MarketNews) GetHeadline() string`

GetHeadline returns the Headline field if non-nil, zero value otherwise.

### GetHeadlineOk

`func (o *MarketNews) GetHeadlineOk() (*string, bool)`

GetHeadlineOk returns a tuple with the Headline field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeadline

`func (o *MarketNews) SetHeadline(v string)`

SetHeadline sets Headline field to given value.

### HasHeadline

`func (o *MarketNews) HasHeadline() bool`

HasHeadline returns a boolean if a field has been set.

### GetId

`func (o *MarketNews) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MarketNews) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MarketNews) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *MarketNews) HasId() bool`

HasId returns a boolean if a field has been set.

### GetImage

`func (o *MarketNews) GetImage() string`

GetImage returns the Image field if non-nil, zero value otherwise.

### GetImageOk

`func (o *MarketNews) GetImageOk() (*string, bool)`

GetImageOk returns a tuple with the Image field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImage

`func (o *MarketNews) SetImage(v string)`

SetImage sets Image field to given value.

### HasImage

`func (o *MarketNews) HasImage() bool`

HasImage returns a boolean if a field has been set.

### GetRelated

`func (o *MarketNews) GetRelated() string`

GetRelated returns the Related field if non-nil, zero value otherwise.

### GetRelatedOk

`func (o *MarketNews) GetRelatedOk() (*string, bool)`

GetRelatedOk returns a tuple with the Related field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelated

`func (o *MarketNews) SetRelated(v string)`

SetRelated sets Related field to given value.

### HasRelated

`func (o *MarketNews) HasRelated() bool`

HasRelated returns a boolean if a field has been set.

### GetSource

`func (o *MarketNews) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *MarketNews) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *MarketNews) SetSource(v string)`

SetSource sets Source field to given value.

### HasSource

`func (o *MarketNews) HasSource() bool`

HasSource returns a boolean if a field has been set.

### GetSummary

`func (o *MarketNews) GetSummary() string`

GetSummary returns the Summary field if non-nil, zero value otherwise.

### GetSummaryOk

`func (o *MarketNews) GetSummaryOk() (*string, bool)`

GetSummaryOk returns a tuple with the Summary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSummary

`func (o *MarketNews) SetSummary(v string)`

SetSummary sets Summary field to given value.

### HasSummary

`func (o *MarketNews) HasSummary() bool`

HasSummary returns a boolean if a field has been set.

### GetUrl

`func (o *MarketNews) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *MarketNews) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *MarketNews) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *MarketNews) HasUrl() bool`

HasUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


