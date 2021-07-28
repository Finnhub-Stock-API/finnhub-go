# SimilarityIndex

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Symbol** | Pointer to **string** | Symbol. | [optional] 
**Cik** | Pointer to **string** | CIK. | [optional] 
**Similarity** | Pointer to [**[]SimilarityIndexInfo**](SimilarityIndexInfo.md) | Array of filings with its cosine similarity compared to the same report of the previous year. | [optional] 

## Methods

### NewSimilarityIndex

`func NewSimilarityIndex() *SimilarityIndex`

NewSimilarityIndex instantiates a new SimilarityIndex object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSimilarityIndexWithDefaults

`func NewSimilarityIndexWithDefaults() *SimilarityIndex`

NewSimilarityIndexWithDefaults instantiates a new SimilarityIndex object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSymbol

`func (o *SimilarityIndex) GetSymbol() string`

GetSymbol returns the Symbol field if non-nil, zero value otherwise.

### GetSymbolOk

`func (o *SimilarityIndex) GetSymbolOk() (*string, bool)`

GetSymbolOk returns a tuple with the Symbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSymbol

`func (o *SimilarityIndex) SetSymbol(v string)`

SetSymbol sets Symbol field to given value.

### HasSymbol

`func (o *SimilarityIndex) HasSymbol() bool`

HasSymbol returns a boolean if a field has been set.

### GetCik

`func (o *SimilarityIndex) GetCik() string`

GetCik returns the Cik field if non-nil, zero value otherwise.

### GetCikOk

`func (o *SimilarityIndex) GetCikOk() (*string, bool)`

GetCikOk returns a tuple with the Cik field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCik

`func (o *SimilarityIndex) SetCik(v string)`

SetCik sets Cik field to given value.

### HasCik

`func (o *SimilarityIndex) HasCik() bool`

HasCik returns a boolean if a field has been set.

### GetSimilarity

`func (o *SimilarityIndex) GetSimilarity() []SimilarityIndexInfo`

GetSimilarity returns the Similarity field if non-nil, zero value otherwise.

### GetSimilarityOk

`func (o *SimilarityIndex) GetSimilarityOk() (*[]SimilarityIndexInfo, bool)`

GetSimilarityOk returns a tuple with the Similarity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSimilarity

`func (o *SimilarityIndex) SetSimilarity(v []SimilarityIndexInfo)`

SetSimilarity sets Similarity field to given value.

### HasSimilarity

`func (o *SimilarityIndex) HasSimilarity() bool`

HasSimilarity returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


