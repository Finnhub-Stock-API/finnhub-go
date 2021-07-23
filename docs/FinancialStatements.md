# FinancialStatements

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Symbol** | Pointer to **string** | Symbol of the company. | [optional] 
**Financials** | Pointer to **[]map[string]interface{}** | An array of map of key, value pairs containing the data for each period. | [optional] 

## Methods

### NewFinancialStatements

`func NewFinancialStatements() *FinancialStatements`

NewFinancialStatements instantiates a new FinancialStatements object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFinancialStatementsWithDefaults

`func NewFinancialStatementsWithDefaults() *FinancialStatements`

NewFinancialStatementsWithDefaults instantiates a new FinancialStatements object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSymbol

`func (o *FinancialStatements) GetSymbol() string`

GetSymbol returns the Symbol field if non-nil, zero value otherwise.

### GetSymbolOk

`func (o *FinancialStatements) GetSymbolOk() (*string, bool)`

GetSymbolOk returns a tuple with the Symbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSymbol

`func (o *FinancialStatements) SetSymbol(v string)`

SetSymbol sets Symbol field to given value.

### HasSymbol

`func (o *FinancialStatements) HasSymbol() bool`

HasSymbol returns a boolean if a field has been set.

### GetFinancials

`func (o *FinancialStatements) GetFinancials() []map[string]interface{}`

GetFinancials returns the Financials field if non-nil, zero value otherwise.

### GetFinancialsOk

`func (o *FinancialStatements) GetFinancialsOk() (*[]map[string]interface{}, bool)`

GetFinancialsOk returns a tuple with the Financials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFinancials

`func (o *FinancialStatements) SetFinancials(v []map[string]interface{})`

SetFinancials sets Financials field to given value.

### HasFinancials

`func (o *FinancialStatements) HasFinancials() bool`

HasFinancials returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


