# DocumentResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DocumentId** | Pointer to **string** | AlphaResearch internal document id. | [optional] 
**Title** | Pointer to **string** | Title for this document. | [optional] 
**Hits** | Pointer to **string** | Number of hit in this document | [optional] 
**Url** | Pointer to **string** | Link to render this document | [optional] 
**Format** | Pointer to **string** | Format of this document (can be html or pdf) | [optional] 
**Excerpts** | Pointer to [**[]ExcerptResponse**](ExcerptResponse.md) | Highlighted excerpts for this document | [optional] 

## Methods

### NewDocumentResponse

`func NewDocumentResponse() *DocumentResponse`

NewDocumentResponse instantiates a new DocumentResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDocumentResponseWithDefaults

`func NewDocumentResponseWithDefaults() *DocumentResponse`

NewDocumentResponseWithDefaults instantiates a new DocumentResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDocumentId

`func (o *DocumentResponse) GetDocumentId() string`

GetDocumentId returns the DocumentId field if non-nil, zero value otherwise.

### GetDocumentIdOk

`func (o *DocumentResponse) GetDocumentIdOk() (*string, bool)`

GetDocumentIdOk returns a tuple with the DocumentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentId

`func (o *DocumentResponse) SetDocumentId(v string)`

SetDocumentId sets DocumentId field to given value.

### HasDocumentId

`func (o *DocumentResponse) HasDocumentId() bool`

HasDocumentId returns a boolean if a field has been set.

### GetTitle

`func (o *DocumentResponse) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *DocumentResponse) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *DocumentResponse) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *DocumentResponse) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetHits

`func (o *DocumentResponse) GetHits() string`

GetHits returns the Hits field if non-nil, zero value otherwise.

### GetHitsOk

`func (o *DocumentResponse) GetHitsOk() (*string, bool)`

GetHitsOk returns a tuple with the Hits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHits

`func (o *DocumentResponse) SetHits(v string)`

SetHits sets Hits field to given value.

### HasHits

`func (o *DocumentResponse) HasHits() bool`

HasHits returns a boolean if a field has been set.

### GetUrl

`func (o *DocumentResponse) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *DocumentResponse) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *DocumentResponse) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *DocumentResponse) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetFormat

`func (o *DocumentResponse) GetFormat() string`

GetFormat returns the Format field if non-nil, zero value otherwise.

### GetFormatOk

`func (o *DocumentResponse) GetFormatOk() (*string, bool)`

GetFormatOk returns a tuple with the Format field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormat

`func (o *DocumentResponse) SetFormat(v string)`

SetFormat sets Format field to given value.

### HasFormat

`func (o *DocumentResponse) HasFormat() bool`

HasFormat returns a boolean if a field has been set.

### GetExcerpts

`func (o *DocumentResponse) GetExcerpts() []ExcerptResponse`

GetExcerpts returns the Excerpts field if non-nil, zero value otherwise.

### GetExcerptsOk

`func (o *DocumentResponse) GetExcerptsOk() (*[]ExcerptResponse, bool)`

GetExcerptsOk returns a tuple with the Excerpts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcerpts

`func (o *DocumentResponse) SetExcerpts(v []ExcerptResponse)`

SetExcerpts sets Excerpts field to given value.

### HasExcerpts

`func (o *DocumentResponse) HasExcerpts() bool`

HasExcerpts returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


