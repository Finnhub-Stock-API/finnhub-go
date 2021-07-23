# Transactions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | Insider&#39;s name. | [optional] 
**Share** | Pointer to **int64** | Number of shares held after the transaction. | [optional] 
**Change** | Pointer to **int64** | Number of share changed from the last period. A positive value suggests a &lt;code&gt;BUY&lt;/code&gt; transaction. A negative value suggests a &lt;code&gt;SELL&lt;/code&gt; transaction. | [optional] 
**FilingDate** | Pointer to **string** | Filing date. | [optional] 
**TransactionDate** | Pointer to **string** | Transaction date. | [optional] 
**TransactionPrice** | Pointer to **float32** | Average transaction price. | [optional] 
**TransactionCode** | Pointer to **string** | Transaction code. A list of codes and their meanings can be found &lt;a href&#x3D;\&quot;https://www.sec.gov/about/forms/form4data.pdf\&quot; target&#x3D;\&quot;_blank\&quot; rel&#x3D;\&quot;noopener\&quot;&gt;here&lt;/a&gt;. | [optional] 

## Methods

### NewTransactions

`func NewTransactions() *Transactions`

NewTransactions instantiates a new Transactions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTransactionsWithDefaults

`func NewTransactionsWithDefaults() *Transactions`

NewTransactionsWithDefaults instantiates a new Transactions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *Transactions) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Transactions) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Transactions) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Transactions) HasName() bool`

HasName returns a boolean if a field has been set.

### GetShare

`func (o *Transactions) GetShare() int64`

GetShare returns the Share field if non-nil, zero value otherwise.

### GetShareOk

`func (o *Transactions) GetShareOk() (*int64, bool)`

GetShareOk returns a tuple with the Share field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShare

`func (o *Transactions) SetShare(v int64)`

SetShare sets Share field to given value.

### HasShare

`func (o *Transactions) HasShare() bool`

HasShare returns a boolean if a field has been set.

### GetChange

`func (o *Transactions) GetChange() int64`

GetChange returns the Change field if non-nil, zero value otherwise.

### GetChangeOk

`func (o *Transactions) GetChangeOk() (*int64, bool)`

GetChangeOk returns a tuple with the Change field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChange

`func (o *Transactions) SetChange(v int64)`

SetChange sets Change field to given value.

### HasChange

`func (o *Transactions) HasChange() bool`

HasChange returns a boolean if a field has been set.

### GetFilingDate

`func (o *Transactions) GetFilingDate() string`

GetFilingDate returns the FilingDate field if non-nil, zero value otherwise.

### GetFilingDateOk

`func (o *Transactions) GetFilingDateOk() (*string, bool)`

GetFilingDateOk returns a tuple with the FilingDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilingDate

`func (o *Transactions) SetFilingDate(v string)`

SetFilingDate sets FilingDate field to given value.

### HasFilingDate

`func (o *Transactions) HasFilingDate() bool`

HasFilingDate returns a boolean if a field has been set.

### GetTransactionDate

`func (o *Transactions) GetTransactionDate() string`

GetTransactionDate returns the TransactionDate field if non-nil, zero value otherwise.

### GetTransactionDateOk

`func (o *Transactions) GetTransactionDateOk() (*string, bool)`

GetTransactionDateOk returns a tuple with the TransactionDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionDate

`func (o *Transactions) SetTransactionDate(v string)`

SetTransactionDate sets TransactionDate field to given value.

### HasTransactionDate

`func (o *Transactions) HasTransactionDate() bool`

HasTransactionDate returns a boolean if a field has been set.

### GetTransactionPrice

`func (o *Transactions) GetTransactionPrice() float32`

GetTransactionPrice returns the TransactionPrice field if non-nil, zero value otherwise.

### GetTransactionPriceOk

`func (o *Transactions) GetTransactionPriceOk() (*float32, bool)`

GetTransactionPriceOk returns a tuple with the TransactionPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionPrice

`func (o *Transactions) SetTransactionPrice(v float32)`

SetTransactionPrice sets TransactionPrice field to given value.

### HasTransactionPrice

`func (o *Transactions) HasTransactionPrice() bool`

HasTransactionPrice returns a boolean if a field has been set.

### GetTransactionCode

`func (o *Transactions) GetTransactionCode() string`

GetTransactionCode returns the TransactionCode field if non-nil, zero value otherwise.

### GetTransactionCodeOk

`func (o *Transactions) GetTransactionCodeOk() (*string, bool)`

GetTransactionCodeOk returns a tuple with the TransactionCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionCode

`func (o *Transactions) SetTransactionCode(v string)`

SetTransactionCode sets TransactionCode field to given value.

### HasTransactionCode

`func (o *Transactions) HasTransactionCode() bool`

HasTransactionCode returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


