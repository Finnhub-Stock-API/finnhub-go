# Dividends2Info

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExDate** | Pointer to **string** | Ex-Dividend date. | [optional] 
**Amount** | Pointer to **float32** | Amount in local currency. | [optional] 

## Methods

### NewDividends2Info

`func NewDividends2Info() *Dividends2Info`

NewDividends2Info instantiates a new Dividends2Info object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDividends2InfoWithDefaults

`func NewDividends2InfoWithDefaults() *Dividends2Info`

NewDividends2InfoWithDefaults instantiates a new Dividends2Info object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExDate

`func (o *Dividends2Info) GetExDate() string`

GetExDate returns the ExDate field if non-nil, zero value otherwise.

### GetExDateOk

`func (o *Dividends2Info) GetExDateOk() (*string, bool)`

GetExDateOk returns a tuple with the ExDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExDate

`func (o *Dividends2Info) SetExDate(v string)`

SetExDate sets ExDate field to given value.

### HasExDate

`func (o *Dividends2Info) HasExDate() bool`

HasExDate returns a boolean if a field has been set.

### GetAmount

`func (o *Dividends2Info) GetAmount() float32`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *Dividends2Info) GetAmountOk() (*float32, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *Dividends2Info) SetAmount(v float32)`

SetAmount sets Amount field to given value.

### HasAmount

`func (o *Dividends2Info) HasAmount() bool`

HasAmount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


