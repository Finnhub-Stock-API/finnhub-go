# SimilarityIndexInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cik** | Pointer to **string** | CIK. | [optional] 
**Item1** | Pointer to **float32** | Cosine similarity of Item 1 (Business). This number is only available for Annual reports. | [optional] 
**Item1a** | Pointer to **float32** | Cosine similarity of Item 1A (Risk Factors). This number is available for both Annual and Quarterly reports. | [optional] 
**Item2** | Pointer to **float32** | Cosine similarity of Item 2 (Management’s Discussion and Analysis of Financial Condition and Results of Operations). This number is only available for Quarterly reports. | [optional] 
**Item7** | Pointer to **float32** | Cosine similarity of Item 7 (Management’s Discussion and Analysis of Financial Condition and Results of Operations). This number is only available for Annual reports. | [optional] 
**Item7a** | Pointer to **float32** | Cosine similarity of Item 7A (Quantitative and Qualitative Disclosures About Market Risk). This number is only available for Annual reports. | [optional] 
**AccessNumber** | Pointer to **string** | Access number. | [optional] 
**Form** | Pointer to **string** | Form type. | [optional] 
**FiledDate** | Pointer to **string** | Filed date &lt;code&gt;%Y-%m-%d %H:%M:%S&lt;/code&gt;. | [optional] 
**AcceptedDate** | Pointer to **string** | Accepted date &lt;code&gt;%Y-%m-%d %H:%M:%S&lt;/code&gt;. | [optional] 
**ReportUrl** | Pointer to **string** | Report&#39;s URL. | [optional] 
**FilingUrl** | Pointer to **string** | Filing&#39;s URL. | [optional] 

## Methods

### NewSimilarityIndexInfo

`func NewSimilarityIndexInfo() *SimilarityIndexInfo`

NewSimilarityIndexInfo instantiates a new SimilarityIndexInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSimilarityIndexInfoWithDefaults

`func NewSimilarityIndexInfoWithDefaults() *SimilarityIndexInfo`

NewSimilarityIndexInfoWithDefaults instantiates a new SimilarityIndexInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCik

`func (o *SimilarityIndexInfo) GetCik() string`

GetCik returns the Cik field if non-nil, zero value otherwise.

### GetCikOk

`func (o *SimilarityIndexInfo) GetCikOk() (*string, bool)`

GetCikOk returns a tuple with the Cik field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCik

`func (o *SimilarityIndexInfo) SetCik(v string)`

SetCik sets Cik field to given value.

### HasCik

`func (o *SimilarityIndexInfo) HasCik() bool`

HasCik returns a boolean if a field has been set.

### GetItem1

`func (o *SimilarityIndexInfo) GetItem1() float32`

GetItem1 returns the Item1 field if non-nil, zero value otherwise.

### GetItem1Ok

`func (o *SimilarityIndexInfo) GetItem1Ok() (*float32, bool)`

GetItem1Ok returns a tuple with the Item1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItem1

`func (o *SimilarityIndexInfo) SetItem1(v float32)`

SetItem1 sets Item1 field to given value.

### HasItem1

`func (o *SimilarityIndexInfo) HasItem1() bool`

HasItem1 returns a boolean if a field has been set.

### GetItem1a

`func (o *SimilarityIndexInfo) GetItem1a() float32`

GetItem1a returns the Item1a field if non-nil, zero value otherwise.

### GetItem1aOk

`func (o *SimilarityIndexInfo) GetItem1aOk() (*float32, bool)`

GetItem1aOk returns a tuple with the Item1a field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItem1a

`func (o *SimilarityIndexInfo) SetItem1a(v float32)`

SetItem1a sets Item1a field to given value.

### HasItem1a

`func (o *SimilarityIndexInfo) HasItem1a() bool`

HasItem1a returns a boolean if a field has been set.

### GetItem2

`func (o *SimilarityIndexInfo) GetItem2() float32`

GetItem2 returns the Item2 field if non-nil, zero value otherwise.

### GetItem2Ok

`func (o *SimilarityIndexInfo) GetItem2Ok() (*float32, bool)`

GetItem2Ok returns a tuple with the Item2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItem2

`func (o *SimilarityIndexInfo) SetItem2(v float32)`

SetItem2 sets Item2 field to given value.

### HasItem2

`func (o *SimilarityIndexInfo) HasItem2() bool`

HasItem2 returns a boolean if a field has been set.

### GetItem7

`func (o *SimilarityIndexInfo) GetItem7() float32`

GetItem7 returns the Item7 field if non-nil, zero value otherwise.

### GetItem7Ok

`func (o *SimilarityIndexInfo) GetItem7Ok() (*float32, bool)`

GetItem7Ok returns a tuple with the Item7 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItem7

`func (o *SimilarityIndexInfo) SetItem7(v float32)`

SetItem7 sets Item7 field to given value.

### HasItem7

`func (o *SimilarityIndexInfo) HasItem7() bool`

HasItem7 returns a boolean if a field has been set.

### GetItem7a

`func (o *SimilarityIndexInfo) GetItem7a() float32`

GetItem7a returns the Item7a field if non-nil, zero value otherwise.

### GetItem7aOk

`func (o *SimilarityIndexInfo) GetItem7aOk() (*float32, bool)`

GetItem7aOk returns a tuple with the Item7a field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItem7a

`func (o *SimilarityIndexInfo) SetItem7a(v float32)`

SetItem7a sets Item7a field to given value.

### HasItem7a

`func (o *SimilarityIndexInfo) HasItem7a() bool`

HasItem7a returns a boolean if a field has been set.

### GetAccessNumber

`func (o *SimilarityIndexInfo) GetAccessNumber() string`

GetAccessNumber returns the AccessNumber field if non-nil, zero value otherwise.

### GetAccessNumberOk

`func (o *SimilarityIndexInfo) GetAccessNumberOk() (*string, bool)`

GetAccessNumberOk returns a tuple with the AccessNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessNumber

`func (o *SimilarityIndexInfo) SetAccessNumber(v string)`

SetAccessNumber sets AccessNumber field to given value.

### HasAccessNumber

`func (o *SimilarityIndexInfo) HasAccessNumber() bool`

HasAccessNumber returns a boolean if a field has been set.

### GetForm

`func (o *SimilarityIndexInfo) GetForm() string`

GetForm returns the Form field if non-nil, zero value otherwise.

### GetFormOk

`func (o *SimilarityIndexInfo) GetFormOk() (*string, bool)`

GetFormOk returns a tuple with the Form field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForm

`func (o *SimilarityIndexInfo) SetForm(v string)`

SetForm sets Form field to given value.

### HasForm

`func (o *SimilarityIndexInfo) HasForm() bool`

HasForm returns a boolean if a field has been set.

### GetFiledDate

`func (o *SimilarityIndexInfo) GetFiledDate() string`

GetFiledDate returns the FiledDate field if non-nil, zero value otherwise.

### GetFiledDateOk

`func (o *SimilarityIndexInfo) GetFiledDateOk() (*string, bool)`

GetFiledDateOk returns a tuple with the FiledDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFiledDate

`func (o *SimilarityIndexInfo) SetFiledDate(v string)`

SetFiledDate sets FiledDate field to given value.

### HasFiledDate

`func (o *SimilarityIndexInfo) HasFiledDate() bool`

HasFiledDate returns a boolean if a field has been set.

### GetAcceptedDate

`func (o *SimilarityIndexInfo) GetAcceptedDate() string`

GetAcceptedDate returns the AcceptedDate field if non-nil, zero value otherwise.

### GetAcceptedDateOk

`func (o *SimilarityIndexInfo) GetAcceptedDateOk() (*string, bool)`

GetAcceptedDateOk returns a tuple with the AcceptedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcceptedDate

`func (o *SimilarityIndexInfo) SetAcceptedDate(v string)`

SetAcceptedDate sets AcceptedDate field to given value.

### HasAcceptedDate

`func (o *SimilarityIndexInfo) HasAcceptedDate() bool`

HasAcceptedDate returns a boolean if a field has been set.

### GetReportUrl

`func (o *SimilarityIndexInfo) GetReportUrl() string`

GetReportUrl returns the ReportUrl field if non-nil, zero value otherwise.

### GetReportUrlOk

`func (o *SimilarityIndexInfo) GetReportUrlOk() (*string, bool)`

GetReportUrlOk returns a tuple with the ReportUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReportUrl

`func (o *SimilarityIndexInfo) SetReportUrl(v string)`

SetReportUrl sets ReportUrl field to given value.

### HasReportUrl

`func (o *SimilarityIndexInfo) HasReportUrl() bool`

HasReportUrl returns a boolean if a field has been set.

### GetFilingUrl

`func (o *SimilarityIndexInfo) GetFilingUrl() string`

GetFilingUrl returns the FilingUrl field if non-nil, zero value otherwise.

### GetFilingUrlOk

`func (o *SimilarityIndexInfo) GetFilingUrlOk() (*string, bool)`

GetFilingUrlOk returns a tuple with the FilingUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilingUrl

`func (o *SimilarityIndexInfo) SetFilingUrl(v string)`

SetFilingUrl sets FilingUrl field to given value.

### HasFilingUrl

`func (o *SimilarityIndexInfo) HasFilingUrl() bool`

HasFilingUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


