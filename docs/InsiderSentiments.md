# InsiderSentiments

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Symbol** | Pointer to **string** | Symbol of the company. | [optional] 
**Data** | Pointer to [**[]InsiderSentimentsData**](InsiderSentimentsData.md) | Array of sentiment data. | [optional] 

## Methods

### NewInsiderSentiments

`func NewInsiderSentiments() *InsiderSentiments`

NewInsiderSentiments instantiates a new InsiderSentiments object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInsiderSentimentsWithDefaults

`func NewInsiderSentimentsWithDefaults() *InsiderSentiments`

NewInsiderSentimentsWithDefaults instantiates a new InsiderSentiments object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSymbol

`func (o *InsiderSentiments) GetSymbol() string`

GetSymbol returns the Symbol field if non-nil, zero value otherwise.

### GetSymbolOk

`func (o *InsiderSentiments) GetSymbolOk() (*string, bool)`

GetSymbolOk returns a tuple with the Symbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSymbol

`func (o *InsiderSentiments) SetSymbol(v string)`

SetSymbol sets Symbol field to given value.

### HasSymbol

`func (o *InsiderSentiments) HasSymbol() bool`

HasSymbol returns a boolean if a field has been set.

### GetData

`func (o *InsiderSentiments) GetData() []InsiderSentimentsData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *InsiderSentiments) GetDataOk() (*[]InsiderSentimentsData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *InsiderSentiments) SetData(v []InsiderSentimentsData)`

SetData sets Data field to given value.

### HasData

`func (o *InsiderSentiments) HasData() bool`

HasData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


