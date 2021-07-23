# LastBidAsk

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**B** | Pointer to **float32** | Bid price. | [optional] 
**A** | Pointer to **float32** | Ask price. | [optional] 
**Bv** | Pointer to **float32** | Bid volume. | [optional] 
**Av** | Pointer to **float32** | Ask volume. | [optional] 
**T** | Pointer to **int64** | Reference UNIX timestamp in ms. | [optional] 

## Methods

### NewLastBidAsk

`func NewLastBidAsk() *LastBidAsk`

NewLastBidAsk instantiates a new LastBidAsk object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLastBidAskWithDefaults

`func NewLastBidAskWithDefaults() *LastBidAsk`

NewLastBidAskWithDefaults instantiates a new LastBidAsk object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetB

`func (o *LastBidAsk) GetB() float32`

GetB returns the B field if non-nil, zero value otherwise.

### GetBOk

`func (o *LastBidAsk) GetBOk() (*float32, bool)`

GetBOk returns a tuple with the B field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetB

`func (o *LastBidAsk) SetB(v float32)`

SetB sets B field to given value.

### HasB

`func (o *LastBidAsk) HasB() bool`

HasB returns a boolean if a field has been set.

### GetA

`func (o *LastBidAsk) GetA() float32`

GetA returns the A field if non-nil, zero value otherwise.

### GetAOk

`func (o *LastBidAsk) GetAOk() (*float32, bool)`

GetAOk returns a tuple with the A field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetA

`func (o *LastBidAsk) SetA(v float32)`

SetA sets A field to given value.

### HasA

`func (o *LastBidAsk) HasA() bool`

HasA returns a boolean if a field has been set.

### GetBv

`func (o *LastBidAsk) GetBv() float32`

GetBv returns the Bv field if non-nil, zero value otherwise.

### GetBvOk

`func (o *LastBidAsk) GetBvOk() (*float32, bool)`

GetBvOk returns a tuple with the Bv field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBv

`func (o *LastBidAsk) SetBv(v float32)`

SetBv sets Bv field to given value.

### HasBv

`func (o *LastBidAsk) HasBv() bool`

HasBv returns a boolean if a field has been set.

### GetAv

`func (o *LastBidAsk) GetAv() float32`

GetAv returns the Av field if non-nil, zero value otherwise.

### GetAvOk

`func (o *LastBidAsk) GetAvOk() (*float32, bool)`

GetAvOk returns a tuple with the Av field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAv

`func (o *LastBidAsk) SetAv(v float32)`

SetAv sets Av field to given value.

### HasAv

`func (o *LastBidAsk) HasAv() bool`

HasAv returns a boolean if a field has been set.

### GetT

`func (o *LastBidAsk) GetT() int64`

GetT returns the T field if non-nil, zero value otherwise.

### GetTOk

`func (o *LastBidAsk) GetTOk() (*int64, bool)`

GetTOk returns a tuple with the T field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetT

`func (o *LastBidAsk) SetT(v int64)`

SetT sets T field to given value.

### HasT

`func (o *LastBidAsk) HasT() bool`

HasT returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


