# Dividends

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Symbol** | Pointer to **string** | Symbol. | [optional] 
**Date** | Pointer to **string** | Ex-Dividend date. | [optional] 
**Amount** | Pointer to **float32** | Amount in local currency. | [optional] 
**AdjustedAmount** | Pointer to **float32** | Adjusted dividend. | [optional] 
**PayDate** | Pointer to **string** | Pay date. | [optional] 
**RecordDate** | Pointer to **string** | Record date. | [optional] 
**DeclarationDate** | Pointer to **string** | Declaration date. | [optional] 
**Currency** | Pointer to **string** | Currency. | [optional] 

## Methods

### NewDividends

`func NewDividends() *Dividends`

NewDividends instantiates a new Dividends object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDividendsWithDefaults

`func NewDividendsWithDefaults() *Dividends`

NewDividendsWithDefaults instantiates a new Dividends object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSymbol

`func (o *Dividends) GetSymbol() string`

GetSymbol returns the Symbol field if non-nil, zero value otherwise.

### GetSymbolOk

`func (o *Dividends) GetSymbolOk() (*string, bool)`

GetSymbolOk returns a tuple with the Symbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSymbol

`func (o *Dividends) SetSymbol(v string)`

SetSymbol sets Symbol field to given value.

### HasSymbol

`func (o *Dividends) HasSymbol() bool`

HasSymbol returns a boolean if a field has been set.

### GetDate

`func (o *Dividends) GetDate() string`

GetDate returns the Date field if non-nil, zero value otherwise.

### GetDateOk

`func (o *Dividends) GetDateOk() (*string, bool)`

GetDateOk returns a tuple with the Date field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDate

`func (o *Dividends) SetDate(v string)`

SetDate sets Date field to given value.

### HasDate

`func (o *Dividends) HasDate() bool`

HasDate returns a boolean if a field has been set.

### GetAmount

`func (o *Dividends) GetAmount() float32`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *Dividends) GetAmountOk() (*float32, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *Dividends) SetAmount(v float32)`

SetAmount sets Amount field to given value.

### HasAmount

`func (o *Dividends) HasAmount() bool`

HasAmount returns a boolean if a field has been set.

### GetAdjustedAmount

`func (o *Dividends) GetAdjustedAmount() float32`

GetAdjustedAmount returns the AdjustedAmount field if non-nil, zero value otherwise.

### GetAdjustedAmountOk

`func (o *Dividends) GetAdjustedAmountOk() (*float32, bool)`

GetAdjustedAmountOk returns a tuple with the AdjustedAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdjustedAmount

`func (o *Dividends) SetAdjustedAmount(v float32)`

SetAdjustedAmount sets AdjustedAmount field to given value.

### HasAdjustedAmount

`func (o *Dividends) HasAdjustedAmount() bool`

HasAdjustedAmount returns a boolean if a field has been set.

### GetPayDate

`func (o *Dividends) GetPayDate() string`

GetPayDate returns the PayDate field if non-nil, zero value otherwise.

### GetPayDateOk

`func (o *Dividends) GetPayDateOk() (*string, bool)`

GetPayDateOk returns a tuple with the PayDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayDate

`func (o *Dividends) SetPayDate(v string)`

SetPayDate sets PayDate field to given value.

### HasPayDate

`func (o *Dividends) HasPayDate() bool`

HasPayDate returns a boolean if a field has been set.

### GetRecordDate

`func (o *Dividends) GetRecordDate() string`

GetRecordDate returns the RecordDate field if non-nil, zero value otherwise.

### GetRecordDateOk

`func (o *Dividends) GetRecordDateOk() (*string, bool)`

GetRecordDateOk returns a tuple with the RecordDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordDate

`func (o *Dividends) SetRecordDate(v string)`

SetRecordDate sets RecordDate field to given value.

### HasRecordDate

`func (o *Dividends) HasRecordDate() bool`

HasRecordDate returns a boolean if a field has been set.

### GetDeclarationDate

`func (o *Dividends) GetDeclarationDate() string`

GetDeclarationDate returns the DeclarationDate field if non-nil, zero value otherwise.

### GetDeclarationDateOk

`func (o *Dividends) GetDeclarationDateOk() (*string, bool)`

GetDeclarationDateOk returns a tuple with the DeclarationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeclarationDate

`func (o *Dividends) SetDeclarationDate(v string)`

SetDeclarationDate sets DeclarationDate field to given value.

### HasDeclarationDate

`func (o *Dividends) HasDeclarationDate() bool`

HasDeclarationDate returns a boolean if a field has been set.

### GetCurrency

`func (o *Dividends) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *Dividends) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *Dividends) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *Dividends) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


