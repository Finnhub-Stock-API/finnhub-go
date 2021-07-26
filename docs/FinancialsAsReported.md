# FinancialsAsReported

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Symbol** | Pointer to **string** | Symbol | [optional] 
**Cik** | Pointer to **string** | CIK | [optional] 
**Data** | Pointer to [**[]Report**](Report.md) | Array of filings. | [optional] 

## Methods

### NewFinancialsAsReported

`func NewFinancialsAsReported() *FinancialsAsReported`

NewFinancialsAsReported instantiates a new FinancialsAsReported object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFinancialsAsReportedWithDefaults

`func NewFinancialsAsReportedWithDefaults() *FinancialsAsReported`

NewFinancialsAsReportedWithDefaults instantiates a new FinancialsAsReported object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSymbol

`func (o *FinancialsAsReported) GetSymbol() string`

GetSymbol returns the Symbol field if non-nil, zero value otherwise.

### GetSymbolOk

`func (o *FinancialsAsReported) GetSymbolOk() (*string, bool)`

GetSymbolOk returns a tuple with the Symbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSymbol

`func (o *FinancialsAsReported) SetSymbol(v string)`

SetSymbol sets Symbol field to given value.

### HasSymbol

`func (o *FinancialsAsReported) HasSymbol() bool`

HasSymbol returns a boolean if a field has been set.

### GetCik

`func (o *FinancialsAsReported) GetCik() string`

GetCik returns the Cik field if non-nil, zero value otherwise.

### GetCikOk

`func (o *FinancialsAsReported) GetCikOk() (*string, bool)`

GetCikOk returns a tuple with the Cik field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCik

`func (o *FinancialsAsReported) SetCik(v string)`

SetCik sets Cik field to given value.

### HasCik

`func (o *FinancialsAsReported) HasCik() bool`

HasCik returns a boolean if a field has been set.

### GetData

`func (o *FinancialsAsReported) GetData() []Report`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *FinancialsAsReported) GetDataOk() (*[]Report, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *FinancialsAsReported) SetData(v []Report)`

SetData sets Data field to given value.

### HasData

`func (o *FinancialsAsReported) HasData() bool`

HasData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


