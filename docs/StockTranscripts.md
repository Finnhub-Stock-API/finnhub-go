# StockTranscripts

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Transcript&#39;s ID used to get the &lt;a href&#x3D;\&quot;#transcripts\&quot;&gt;full transcript&lt;/a&gt;. | [optional] 
**Title** | Pointer to **string** | Title. | [optional] 
**Time** | Pointer to **time.Time** | Time of the event. | [optional] 
**Year** | Pointer to **int64** | Year of earnings result in the case of earnings call transcript. | [optional] 
**Quarter** | Pointer to **int64** | Quarter of earnings result in the case of earnings call transcript. | [optional] 

## Methods

### NewStockTranscripts

`func NewStockTranscripts() *StockTranscripts`

NewStockTranscripts instantiates a new StockTranscripts object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStockTranscriptsWithDefaults

`func NewStockTranscriptsWithDefaults() *StockTranscripts`

NewStockTranscriptsWithDefaults instantiates a new StockTranscripts object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *StockTranscripts) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *StockTranscripts) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *StockTranscripts) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *StockTranscripts) HasId() bool`

HasId returns a boolean if a field has been set.

### GetTitle

`func (o *StockTranscripts) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *StockTranscripts) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *StockTranscripts) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *StockTranscripts) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetTime

`func (o *StockTranscripts) GetTime() time.Time`

GetTime returns the Time field if non-nil, zero value otherwise.

### GetTimeOk

`func (o *StockTranscripts) GetTimeOk() (*time.Time, bool)`

GetTimeOk returns a tuple with the Time field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTime

`func (o *StockTranscripts) SetTime(v time.Time)`

SetTime sets Time field to given value.

### HasTime

`func (o *StockTranscripts) HasTime() bool`

HasTime returns a boolean if a field has been set.

### GetYear

`func (o *StockTranscripts) GetYear() int64`

GetYear returns the Year field if non-nil, zero value otherwise.

### GetYearOk

`func (o *StockTranscripts) GetYearOk() (*int64, bool)`

GetYearOk returns a tuple with the Year field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetYear

`func (o *StockTranscripts) SetYear(v int64)`

SetYear sets Year field to given value.

### HasYear

`func (o *StockTranscripts) HasYear() bool`

HasYear returns a boolean if a field has been set.

### GetQuarter

`func (o *StockTranscripts) GetQuarter() int64`

GetQuarter returns the Quarter field if non-nil, zero value otherwise.

### GetQuarterOk

`func (o *StockTranscripts) GetQuarterOk() (*int64, bool)`

GetQuarterOk returns a tuple with the Quarter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuarter

`func (o *StockTranscripts) SetQuarter(v int64)`

SetQuarter sets Quarter field to given value.

### HasQuarter

`func (o *StockTranscripts) HasQuarter() bool`

HasQuarter returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


