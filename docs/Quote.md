# Quote

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**O** | Pointer to **float32** | Open price of the day | [optional] 
**H** | Pointer to **float32** | High price of the day | [optional] 
**L** | Pointer to **float32** | Low price of the day | [optional] 
**C** | Pointer to **float32** | Current price | [optional] 
**Pc** | Pointer to **float32** | Previous close price | [optional] 
**D** | Pointer to **float32** | Change | [optional] 
**Dp** | Pointer to **float32** | Percent change | [optional] 

## Methods

### NewQuote

`func NewQuote() *Quote`

NewQuote instantiates a new Quote object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQuoteWithDefaults

`func NewQuoteWithDefaults() *Quote`

NewQuoteWithDefaults instantiates a new Quote object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetO

`func (o *Quote) GetO() float32`

GetO returns the O field if non-nil, zero value otherwise.

### GetOOk

`func (o *Quote) GetOOk() (*float32, bool)`

GetOOk returns a tuple with the O field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetO

`func (o *Quote) SetO(v float32)`

SetO sets O field to given value.

### HasO

`func (o *Quote) HasO() bool`

HasO returns a boolean if a field has been set.

### GetH

`func (o *Quote) GetH() float32`

GetH returns the H field if non-nil, zero value otherwise.

### GetHOk

`func (o *Quote) GetHOk() (*float32, bool)`

GetHOk returns a tuple with the H field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetH

`func (o *Quote) SetH(v float32)`

SetH sets H field to given value.

### HasH

`func (o *Quote) HasH() bool`

HasH returns a boolean if a field has been set.

### GetL

`func (o *Quote) GetL() float32`

GetL returns the L field if non-nil, zero value otherwise.

### GetLOk

`func (o *Quote) GetLOk() (*float32, bool)`

GetLOk returns a tuple with the L field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetL

`func (o *Quote) SetL(v float32)`

SetL sets L field to given value.

### HasL

`func (o *Quote) HasL() bool`

HasL returns a boolean if a field has been set.

### GetC

`func (o *Quote) GetC() float32`

GetC returns the C field if non-nil, zero value otherwise.

### GetCOk

`func (o *Quote) GetCOk() (*float32, bool)`

GetCOk returns a tuple with the C field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetC

`func (o *Quote) SetC(v float32)`

SetC sets C field to given value.

### HasC

`func (o *Quote) HasC() bool`

HasC returns a boolean if a field has been set.

### GetPc

`func (o *Quote) GetPc() float32`

GetPc returns the Pc field if non-nil, zero value otherwise.

### GetPcOk

`func (o *Quote) GetPcOk() (*float32, bool)`

GetPcOk returns a tuple with the Pc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPc

`func (o *Quote) SetPc(v float32)`

SetPc sets Pc field to given value.

### HasPc

`func (o *Quote) HasPc() bool`

HasPc returns a boolean if a field has been set.

### GetD

`func (o *Quote) GetD() float32`

GetD returns the D field if non-nil, zero value otherwise.

### GetDOk

`func (o *Quote) GetDOk() (*float32, bool)`

GetDOk returns a tuple with the D field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetD

`func (o *Quote) SetD(v float32)`

SetD sets D field to given value.

### HasD

`func (o *Quote) HasD() bool`

HasD returns a boolean if a field has been set.

### GetDp

`func (o *Quote) GetDp() float32`

GetDp returns the Dp field if non-nil, zero value otherwise.

### GetDpOk

`func (o *Quote) GetDpOk() (*float32, bool)`

GetDpOk returns a tuple with the Dp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDp

`func (o *Quote) SetDp(v float32)`

SetDp sets Dp field to given value.

### HasDp

`func (o *Quote) HasDp() bool`

HasDp returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


