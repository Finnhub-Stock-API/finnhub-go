# InsiderTransactions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Symbol** | Pointer to **string** | Symbol of the company. | [optional] 
**Data** | Pointer to [**[]Transactions**](Transactions.md) | Array of insider transactions. | [optional] 

## Methods

### NewInsiderTransactions

`func NewInsiderTransactions() *InsiderTransactions`

NewInsiderTransactions instantiates a new InsiderTransactions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInsiderTransactionsWithDefaults

`func NewInsiderTransactionsWithDefaults() *InsiderTransactions`

NewInsiderTransactionsWithDefaults instantiates a new InsiderTransactions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSymbol

`func (o *InsiderTransactions) GetSymbol() string`

GetSymbol returns the Symbol field if non-nil, zero value otherwise.

### GetSymbolOk

`func (o *InsiderTransactions) GetSymbolOk() (*string, bool)`

GetSymbolOk returns a tuple with the Symbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSymbol

`func (o *InsiderTransactions) SetSymbol(v string)`

SetSymbol sets Symbol field to given value.

### HasSymbol

`func (o *InsiderTransactions) HasSymbol() bool`

HasSymbol returns a boolean if a field has been set.

### GetData

`func (o *InsiderTransactions) GetData() []Transactions`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *InsiderTransactions) GetDataOk() (*[]Transactions, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *InsiderTransactions) SetData(v []Transactions)`

SetData sets Data field to given value.

### HasData

`func (o *InsiderTransactions) HasData() bool`

HasData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


