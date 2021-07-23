# News

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

### NewNews

`func NewNews() *News`

NewNews instantiates a new News object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNewsWithDefaults

`func NewNewsWithDefaults() *News`

NewNewsWithDefaults instantiates a new News object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCategory

`func (o *News) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *News) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *News) SetCategory(v string)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *News) HasCategory() bool`

HasCategory returns a boolean if a field has been set.

### GetDatetime

`func (o *News) GetDatetime() int64`

GetDatetime returns the Datetime field if non-nil, zero value otherwise.

### GetDatetimeOk

`func (o *News) GetDatetimeOk() (*int64, bool)`

GetDatetimeOk returns a tuple with the Datetime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatetime

`func (o *News) SetDatetime(v int64)`

SetDatetime sets Datetime field to given value.

### HasDatetime

`func (o *News) HasDatetime() bool`

HasDatetime returns a boolean if a field has been set.

### GetHeadline

`func (o *News) GetHeadline() string`

GetHeadline returns the Headline field if non-nil, zero value otherwise.

### GetHeadlineOk

`func (o *News) GetHeadlineOk() (*string, bool)`

GetHeadlineOk returns a tuple with the Headline field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeadline

`func (o *News) SetHeadline(v string)`

SetHeadline sets Headline field to given value.

### HasHeadline

`func (o *News) HasHeadline() bool`

HasHeadline returns a boolean if a field has been set.

### GetId

`func (o *News) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *News) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *News) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *News) HasId() bool`

HasId returns a boolean if a field has been set.

### GetImage

`func (o *News) GetImage() string`

GetImage returns the Image field if non-nil, zero value otherwise.

### GetImageOk

`func (o *News) GetImageOk() (*string, bool)`

GetImageOk returns a tuple with the Image field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImage

`func (o *News) SetImage(v string)`

SetImage sets Image field to given value.

### HasImage

`func (o *News) HasImage() bool`

HasImage returns a boolean if a field has been set.

### GetRelated

`func (o *News) GetRelated() string`

GetRelated returns the Related field if non-nil, zero value otherwise.

### GetRelatedOk

`func (o *News) GetRelatedOk() (*string, bool)`

GetRelatedOk returns a tuple with the Related field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelated

`func (o *News) SetRelated(v string)`

SetRelated sets Related field to given value.

### HasRelated

`func (o *News) HasRelated() bool`

HasRelated returns a boolean if a field has been set.

### GetSource

`func (o *News) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *News) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *News) SetSource(v string)`

SetSource sets Source field to given value.

### HasSource

`func (o *News) HasSource() bool`

HasSource returns a boolean if a field has been set.

### GetSummary

`func (o *News) GetSummary() string`

GetSummary returns the Summary field if non-nil, zero value otherwise.

### GetSummaryOk

`func (o *News) GetSummaryOk() (*string, bool)`

GetSummaryOk returns a tuple with the Summary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSummary

`func (o *News) SetSummary(v string)`

SetSummary sets Summary field to given value.

### HasSummary

`func (o *News) HasSummary() bool`

HasSummary returns a boolean if a field has been set.

### GetUrl

`func (o *News) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *News) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *News) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *News) HasUrl() bool`

HasUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


