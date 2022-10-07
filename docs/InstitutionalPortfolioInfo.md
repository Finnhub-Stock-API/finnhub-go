# InstitutionalPortfolioInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Symbol** | Pointer to **string** | Symbol. | [optional] 
**Cusip** | Pointer to **string** | CUSIP. | [optional] 
**Name** | Pointer to **string** | Position&#39;s name. | [optional] 
**PutCall** | Pointer to **string** | &lt;code&gt;put&lt;/code&gt; or &lt;code&gt;call&lt;/code&gt; for options. | [optional] 
**Change** | Pointer to **float32** | Number of shares change. | [optional] 
**NoVoting** | Pointer to **float32** | Number of shares with no voting rights. | [optional] 
**Percentage** | Pointer to **float32** | Percentage of portfolio. | [optional] 
**Share** | Pointer to **float32** | Number of shares. | [optional] 
**SharedVoting** | Pointer to **float32** | Number of shares with shared voting rights. | [optional] 
**SoleVoting** | Pointer to **float32** | Number of shares with sole voting rights. | [optional] 
**Value** | Pointer to **float32** | Position value. | [optional] 

## Methods

### NewInstitutionalPortfolioInfo

`func NewInstitutionalPortfolioInfo() *InstitutionalPortfolioInfo`

NewInstitutionalPortfolioInfo instantiates a new InstitutionalPortfolioInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInstitutionalPortfolioInfoWithDefaults

`func NewInstitutionalPortfolioInfoWithDefaults() *InstitutionalPortfolioInfo`

NewInstitutionalPortfolioInfoWithDefaults instantiates a new InstitutionalPortfolioInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSymbol

`func (o *InstitutionalPortfolioInfo) GetSymbol() string`

GetSymbol returns the Symbol field if non-nil, zero value otherwise.

### GetSymbolOk

`func (o *InstitutionalPortfolioInfo) GetSymbolOk() (*string, bool)`

GetSymbolOk returns a tuple with the Symbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSymbol

`func (o *InstitutionalPortfolioInfo) SetSymbol(v string)`

SetSymbol sets Symbol field to given value.

### HasSymbol

`func (o *InstitutionalPortfolioInfo) HasSymbol() bool`

HasSymbol returns a boolean if a field has been set.

### GetCusip

`func (o *InstitutionalPortfolioInfo) GetCusip() string`

GetCusip returns the Cusip field if non-nil, zero value otherwise.

### GetCusipOk

`func (o *InstitutionalPortfolioInfo) GetCusipOk() (*string, bool)`

GetCusipOk returns a tuple with the Cusip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCusip

`func (o *InstitutionalPortfolioInfo) SetCusip(v string)`

SetCusip sets Cusip field to given value.

### HasCusip

`func (o *InstitutionalPortfolioInfo) HasCusip() bool`

HasCusip returns a boolean if a field has been set.

### GetName

`func (o *InstitutionalPortfolioInfo) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *InstitutionalPortfolioInfo) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *InstitutionalPortfolioInfo) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *InstitutionalPortfolioInfo) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPutCall

`func (o *InstitutionalPortfolioInfo) GetPutCall() string`

GetPutCall returns the PutCall field if non-nil, zero value otherwise.

### GetPutCallOk

`func (o *InstitutionalPortfolioInfo) GetPutCallOk() (*string, bool)`

GetPutCallOk returns a tuple with the PutCall field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPutCall

`func (o *InstitutionalPortfolioInfo) SetPutCall(v string)`

SetPutCall sets PutCall field to given value.

### HasPutCall

`func (o *InstitutionalPortfolioInfo) HasPutCall() bool`

HasPutCall returns a boolean if a field has been set.

### GetChange

`func (o *InstitutionalPortfolioInfo) GetChange() float32`

GetChange returns the Change field if non-nil, zero value otherwise.

### GetChangeOk

`func (o *InstitutionalPortfolioInfo) GetChangeOk() (*float32, bool)`

GetChangeOk returns a tuple with the Change field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChange

`func (o *InstitutionalPortfolioInfo) SetChange(v float32)`

SetChange sets Change field to given value.

### HasChange

`func (o *InstitutionalPortfolioInfo) HasChange() bool`

HasChange returns a boolean if a field has been set.

### GetNoVoting

`func (o *InstitutionalPortfolioInfo) GetNoVoting() float32`

GetNoVoting returns the NoVoting field if non-nil, zero value otherwise.

### GetNoVotingOk

`func (o *InstitutionalPortfolioInfo) GetNoVotingOk() (*float32, bool)`

GetNoVotingOk returns a tuple with the NoVoting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNoVoting

`func (o *InstitutionalPortfolioInfo) SetNoVoting(v float32)`

SetNoVoting sets NoVoting field to given value.

### HasNoVoting

`func (o *InstitutionalPortfolioInfo) HasNoVoting() bool`

HasNoVoting returns a boolean if a field has been set.

### GetPercentage

`func (o *InstitutionalPortfolioInfo) GetPercentage() float32`

GetPercentage returns the Percentage field if non-nil, zero value otherwise.

### GetPercentageOk

`func (o *InstitutionalPortfolioInfo) GetPercentageOk() (*float32, bool)`

GetPercentageOk returns a tuple with the Percentage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPercentage

`func (o *InstitutionalPortfolioInfo) SetPercentage(v float32)`

SetPercentage sets Percentage field to given value.

### HasPercentage

`func (o *InstitutionalPortfolioInfo) HasPercentage() bool`

HasPercentage returns a boolean if a field has been set.

### GetShare

`func (o *InstitutionalPortfolioInfo) GetShare() float32`

GetShare returns the Share field if non-nil, zero value otherwise.

### GetShareOk

`func (o *InstitutionalPortfolioInfo) GetShareOk() (*float32, bool)`

GetShareOk returns a tuple with the Share field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShare

`func (o *InstitutionalPortfolioInfo) SetShare(v float32)`

SetShare sets Share field to given value.

### HasShare

`func (o *InstitutionalPortfolioInfo) HasShare() bool`

HasShare returns a boolean if a field has been set.

### GetSharedVoting

`func (o *InstitutionalPortfolioInfo) GetSharedVoting() float32`

GetSharedVoting returns the SharedVoting field if non-nil, zero value otherwise.

### GetSharedVotingOk

`func (o *InstitutionalPortfolioInfo) GetSharedVotingOk() (*float32, bool)`

GetSharedVotingOk returns a tuple with the SharedVoting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharedVoting

`func (o *InstitutionalPortfolioInfo) SetSharedVoting(v float32)`

SetSharedVoting sets SharedVoting field to given value.

### HasSharedVoting

`func (o *InstitutionalPortfolioInfo) HasSharedVoting() bool`

HasSharedVoting returns a boolean if a field has been set.

### GetSoleVoting

`func (o *InstitutionalPortfolioInfo) GetSoleVoting() float32`

GetSoleVoting returns the SoleVoting field if non-nil, zero value otherwise.

### GetSoleVotingOk

`func (o *InstitutionalPortfolioInfo) GetSoleVotingOk() (*float32, bool)`

GetSoleVotingOk returns a tuple with the SoleVoting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSoleVoting

`func (o *InstitutionalPortfolioInfo) SetSoleVoting(v float32)`

SetSoleVoting sets SoleVoting field to given value.

### HasSoleVoting

`func (o *InstitutionalPortfolioInfo) HasSoleVoting() bool`

HasSoleVoting returns a boolean if a field has been set.

### GetValue

`func (o *InstitutionalPortfolioInfo) GetValue() float32`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *InstitutionalPortfolioInfo) GetValueOk() (*float32, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *InstitutionalPortfolioInfo) SetValue(v float32)`

SetValue sets Value field to given value.

### HasValue

`func (o *InstitutionalPortfolioInfo) HasValue() bool`

HasValue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


