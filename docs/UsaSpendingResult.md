# UsaSpendingResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Symbol** | Pointer to **string** | Symbol. | [optional] 
**Data** | Pointer to [**[]UsaSpending**](UsaSpending.md) | Array of government&#39;s spending data points. | [optional] 

## Methods

### NewUsaSpendingResult

`func NewUsaSpendingResult() *UsaSpendingResult`

NewUsaSpendingResult instantiates a new UsaSpendingResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUsaSpendingResultWithDefaults

`func NewUsaSpendingResultWithDefaults() *UsaSpendingResult`

NewUsaSpendingResultWithDefaults instantiates a new UsaSpendingResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSymbol

`func (o *UsaSpendingResult) GetSymbol() string`

GetSymbol returns the Symbol field if non-nil, zero value otherwise.

### GetSymbolOk

`func (o *UsaSpendingResult) GetSymbolOk() (*string, bool)`

GetSymbolOk returns a tuple with the Symbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSymbol

`func (o *UsaSpendingResult) SetSymbol(v string)`

SetSymbol sets Symbol field to given value.

### HasSymbol

`func (o *UsaSpendingResult) HasSymbol() bool`

HasSymbol returns a boolean if a field has been set.

### GetData

`func (o *UsaSpendingResult) GetData() []UsaSpending`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *UsaSpendingResult) GetDataOk() (*[]UsaSpending, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *UsaSpendingResult) SetData(v []UsaSpending)`

SetData sets Data field to given value.

### HasData

`func (o *UsaSpendingResult) HasData() bool`

HasData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


