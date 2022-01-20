# UsptoPatentResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Symbol** | Pointer to **string** | Symbol. | [optional] 
**Data** | Pointer to [**[]UsptoPatent**](UsptoPatent.md) | Array of patents. | [optional] 

## Methods

### NewUsptoPatentResult

`func NewUsptoPatentResult() *UsptoPatentResult`

NewUsptoPatentResult instantiates a new UsptoPatentResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUsptoPatentResultWithDefaults

`func NewUsptoPatentResultWithDefaults() *UsptoPatentResult`

NewUsptoPatentResultWithDefaults instantiates a new UsptoPatentResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSymbol

`func (o *UsptoPatentResult) GetSymbol() string`

GetSymbol returns the Symbol field if non-nil, zero value otherwise.

### GetSymbolOk

`func (o *UsptoPatentResult) GetSymbolOk() (*string, bool)`

GetSymbolOk returns a tuple with the Symbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSymbol

`func (o *UsptoPatentResult) SetSymbol(v string)`

SetSymbol sets Symbol field to given value.

### HasSymbol

`func (o *UsptoPatentResult) HasSymbol() bool`

HasSymbol returns a boolean if a field has been set.

### GetData

`func (o *UsptoPatentResult) GetData() []UsptoPatent`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *UsptoPatentResult) GetDataOk() (*[]UsptoPatent, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *UsptoPatentResult) SetData(v []UsptoPatent)`

SetData sets Data field to given value.

### HasData

`func (o *UsptoPatentResult) HasData() bool`

HasData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


