# FilingResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FilingId** | Pointer to **string** | Filing Id in Alpharesearch platform | [optional] 
**Title** | Pointer to **string** | Filing title | [optional] 
**FilerId** | Pointer to **string** | Id of the entity submitted the filing | [optional] 
**Symbol** | Pointer to **map[string]interface{}** | List of symbol associate with this filing | [optional] 
**Name** | Pointer to **string** | Filer name | [optional] 
**AcceptanceDate** | Pointer to **string** | Date the filing is submitted. | [optional] 
**FiledDate** | Pointer to **string** | Date the filing is made available to the public | [optional] 
**ReportDate** | Pointer to **string** | Date as which the filing is reported | [optional] 
**Form** | Pointer to **string** | Filing Form | [optional] 
**Amend** | Pointer to **bool** | Amendment | [optional] 
**Source** | Pointer to **string** | Filing Source | [optional] 
**PageCount** | Pointer to **int32** | Estimate number of page when printing | [optional] 
**DocumentCount** | Pointer to **int32** | Number of document in this filing | [optional] 

## Methods

### NewFilingResponse

`func NewFilingResponse() *FilingResponse`

NewFilingResponse instantiates a new FilingResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFilingResponseWithDefaults

`func NewFilingResponseWithDefaults() *FilingResponse`

NewFilingResponseWithDefaults instantiates a new FilingResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFilingId

`func (o *FilingResponse) GetFilingId() string`

GetFilingId returns the FilingId field if non-nil, zero value otherwise.

### GetFilingIdOk

`func (o *FilingResponse) GetFilingIdOk() (*string, bool)`

GetFilingIdOk returns a tuple with the FilingId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilingId

`func (o *FilingResponse) SetFilingId(v string)`

SetFilingId sets FilingId field to given value.

### HasFilingId

`func (o *FilingResponse) HasFilingId() bool`

HasFilingId returns a boolean if a field has been set.

### GetTitle

`func (o *FilingResponse) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *FilingResponse) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *FilingResponse) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *FilingResponse) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetFilerId

`func (o *FilingResponse) GetFilerId() string`

GetFilerId returns the FilerId field if non-nil, zero value otherwise.

### GetFilerIdOk

`func (o *FilingResponse) GetFilerIdOk() (*string, bool)`

GetFilerIdOk returns a tuple with the FilerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilerId

`func (o *FilingResponse) SetFilerId(v string)`

SetFilerId sets FilerId field to given value.

### HasFilerId

`func (o *FilingResponse) HasFilerId() bool`

HasFilerId returns a boolean if a field has been set.

### GetSymbol

`func (o *FilingResponse) GetSymbol() map[string]interface{}`

GetSymbol returns the Symbol field if non-nil, zero value otherwise.

### GetSymbolOk

`func (o *FilingResponse) GetSymbolOk() (*map[string]interface{}, bool)`

GetSymbolOk returns a tuple with the Symbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSymbol

`func (o *FilingResponse) SetSymbol(v map[string]interface{})`

SetSymbol sets Symbol field to given value.

### HasSymbol

`func (o *FilingResponse) HasSymbol() bool`

HasSymbol returns a boolean if a field has been set.

### GetName

`func (o *FilingResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *FilingResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *FilingResponse) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *FilingResponse) HasName() bool`

HasName returns a boolean if a field has been set.

### GetAcceptanceDate

`func (o *FilingResponse) GetAcceptanceDate() string`

GetAcceptanceDate returns the AcceptanceDate field if non-nil, zero value otherwise.

### GetAcceptanceDateOk

`func (o *FilingResponse) GetAcceptanceDateOk() (*string, bool)`

GetAcceptanceDateOk returns a tuple with the AcceptanceDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcceptanceDate

`func (o *FilingResponse) SetAcceptanceDate(v string)`

SetAcceptanceDate sets AcceptanceDate field to given value.

### HasAcceptanceDate

`func (o *FilingResponse) HasAcceptanceDate() bool`

HasAcceptanceDate returns a boolean if a field has been set.

### GetFiledDate

`func (o *FilingResponse) GetFiledDate() string`

GetFiledDate returns the FiledDate field if non-nil, zero value otherwise.

### GetFiledDateOk

`func (o *FilingResponse) GetFiledDateOk() (*string, bool)`

GetFiledDateOk returns a tuple with the FiledDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFiledDate

`func (o *FilingResponse) SetFiledDate(v string)`

SetFiledDate sets FiledDate field to given value.

### HasFiledDate

`func (o *FilingResponse) HasFiledDate() bool`

HasFiledDate returns a boolean if a field has been set.

### GetReportDate

`func (o *FilingResponse) GetReportDate() string`

GetReportDate returns the ReportDate field if non-nil, zero value otherwise.

### GetReportDateOk

`func (o *FilingResponse) GetReportDateOk() (*string, bool)`

GetReportDateOk returns a tuple with the ReportDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReportDate

`func (o *FilingResponse) SetReportDate(v string)`

SetReportDate sets ReportDate field to given value.

### HasReportDate

`func (o *FilingResponse) HasReportDate() bool`

HasReportDate returns a boolean if a field has been set.

### GetForm

`func (o *FilingResponse) GetForm() string`

GetForm returns the Form field if non-nil, zero value otherwise.

### GetFormOk

`func (o *FilingResponse) GetFormOk() (*string, bool)`

GetFormOk returns a tuple with the Form field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForm

`func (o *FilingResponse) SetForm(v string)`

SetForm sets Form field to given value.

### HasForm

`func (o *FilingResponse) HasForm() bool`

HasForm returns a boolean if a field has been set.

### GetAmend

`func (o *FilingResponse) GetAmend() bool`

GetAmend returns the Amend field if non-nil, zero value otherwise.

### GetAmendOk

`func (o *FilingResponse) GetAmendOk() (*bool, bool)`

GetAmendOk returns a tuple with the Amend field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmend

`func (o *FilingResponse) SetAmend(v bool)`

SetAmend sets Amend field to given value.

### HasAmend

`func (o *FilingResponse) HasAmend() bool`

HasAmend returns a boolean if a field has been set.

### GetSource

`func (o *FilingResponse) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *FilingResponse) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *FilingResponse) SetSource(v string)`

SetSource sets Source field to given value.

### HasSource

`func (o *FilingResponse) HasSource() bool`

HasSource returns a boolean if a field has been set.

### GetPageCount

`func (o *FilingResponse) GetPageCount() int32`

GetPageCount returns the PageCount field if non-nil, zero value otherwise.

### GetPageCountOk

`func (o *FilingResponse) GetPageCountOk() (*int32, bool)`

GetPageCountOk returns a tuple with the PageCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageCount

`func (o *FilingResponse) SetPageCount(v int32)`

SetPageCount sets PageCount field to given value.

### HasPageCount

`func (o *FilingResponse) HasPageCount() bool`

HasPageCount returns a boolean if a field has been set.

### GetDocumentCount

`func (o *FilingResponse) GetDocumentCount() int32`

GetDocumentCount returns the DocumentCount field if non-nil, zero value otherwise.

### GetDocumentCountOk

`func (o *FilingResponse) GetDocumentCountOk() (*int32, bool)`

GetDocumentCountOk returns a tuple with the DocumentCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentCount

`func (o *FilingResponse) SetDocumentCount(v int32)`

SetDocumentCount sets DocumentCount field to given value.

### HasDocumentCount

`func (o *FilingResponse) HasDocumentCount() bool`

HasDocumentCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


