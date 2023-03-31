# SearchResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Count** | Pointer to **int32** | Total filing matched your search criteria. | [optional] 
**Took** | Pointer to **int32** | Time took to execute your search query on our server, value in ms. | [optional] 
**Page** | Pointer to **int32** | Current search page | [optional] 
**Filings** | Pointer to [**[]FilingResponse**](FilingResponse.md) | Filing match your search criteria. | [optional] 

## Methods

### NewSearchResponse

`func NewSearchResponse() *SearchResponse`

NewSearchResponse instantiates a new SearchResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSearchResponseWithDefaults

`func NewSearchResponseWithDefaults() *SearchResponse`

NewSearchResponseWithDefaults instantiates a new SearchResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCount

`func (o *SearchResponse) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *SearchResponse) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *SearchResponse) SetCount(v int32)`

SetCount sets Count field to given value.

### HasCount

`func (o *SearchResponse) HasCount() bool`

HasCount returns a boolean if a field has been set.

### GetTook

`func (o *SearchResponse) GetTook() int32`

GetTook returns the Took field if non-nil, zero value otherwise.

### GetTookOk

`func (o *SearchResponse) GetTookOk() (*int32, bool)`

GetTookOk returns a tuple with the Took field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTook

`func (o *SearchResponse) SetTook(v int32)`

SetTook sets Took field to given value.

### HasTook

`func (o *SearchResponse) HasTook() bool`

HasTook returns a boolean if a field has been set.

### GetPage

`func (o *SearchResponse) GetPage() int32`

GetPage returns the Page field if non-nil, zero value otherwise.

### GetPageOk

`func (o *SearchResponse) GetPageOk() (*int32, bool)`

GetPageOk returns a tuple with the Page field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPage

`func (o *SearchResponse) SetPage(v int32)`

SetPage sets Page field to given value.

### HasPage

`func (o *SearchResponse) HasPage() bool`

HasPage returns a boolean if a field has been set.

### GetFilings

`func (o *SearchResponse) GetFilings() []FilingResponse`

GetFilings returns the Filings field if non-nil, zero value otherwise.

### GetFilingsOk

`func (o *SearchResponse) GetFilingsOk() (*[]FilingResponse, bool)`

GetFilingsOk returns a tuple with the Filings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilings

`func (o *SearchResponse) SetFilings(v []FilingResponse)`

SetFilings sets Filings field to given value.

### HasFilings

`func (o *SearchResponse) HasFilings() bool`

HasFilings returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


