# InstitutionalOwnershipInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cik** | Pointer to **string** | Investor&#39;s company CIK. | [optional] 
**Name** | Pointer to **string** | Firm&#39;s name. | [optional] 
**PutCall** | Pointer to **string** | &lt;code&gt;put&lt;/code&gt; or &lt;code&gt;call&lt;/code&gt; for options. | [optional] 
**Change** | Pointer to **float32** | Number of shares change. | [optional] 
**NoVoting** | Pointer to **float32** | Number of shares with no voting rights. | [optional] 
**Percentage** | Pointer to **float32** | Percentage of portfolio. | [optional] 
**Share** | Pointer to **float32** | News score. | [optional] 
**SharedVoting** | Pointer to **float32** | Number of shares with shared voting rights. | [optional] 
**SoleVoting** | Pointer to **float32** | Number of shares with sole voting rights. | [optional] 
**Value** | Pointer to **float32** | Position value. | [optional] 

## Methods

### NewInstitutionalOwnershipInfo

`func NewInstitutionalOwnershipInfo() *InstitutionalOwnershipInfo`

NewInstitutionalOwnershipInfo instantiates a new InstitutionalOwnershipInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInstitutionalOwnershipInfoWithDefaults

`func NewInstitutionalOwnershipInfoWithDefaults() *InstitutionalOwnershipInfo`

NewInstitutionalOwnershipInfoWithDefaults instantiates a new InstitutionalOwnershipInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCik

`func (o *InstitutionalOwnershipInfo) GetCik() string`

GetCik returns the Cik field if non-nil, zero value otherwise.

### GetCikOk

`func (o *InstitutionalOwnershipInfo) GetCikOk() (*string, bool)`

GetCikOk returns a tuple with the Cik field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCik

`func (o *InstitutionalOwnershipInfo) SetCik(v string)`

SetCik sets Cik field to given value.

### HasCik

`func (o *InstitutionalOwnershipInfo) HasCik() bool`

HasCik returns a boolean if a field has been set.

### GetName

`func (o *InstitutionalOwnershipInfo) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *InstitutionalOwnershipInfo) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *InstitutionalOwnershipInfo) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *InstitutionalOwnershipInfo) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPutCall

`func (o *InstitutionalOwnershipInfo) GetPutCall() string`

GetPutCall returns the PutCall field if non-nil, zero value otherwise.

### GetPutCallOk

`func (o *InstitutionalOwnershipInfo) GetPutCallOk() (*string, bool)`

GetPutCallOk returns a tuple with the PutCall field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPutCall

`func (o *InstitutionalOwnershipInfo) SetPutCall(v string)`

SetPutCall sets PutCall field to given value.

### HasPutCall

`func (o *InstitutionalOwnershipInfo) HasPutCall() bool`

HasPutCall returns a boolean if a field has been set.

### GetChange

`func (o *InstitutionalOwnershipInfo) GetChange() float32`

GetChange returns the Change field if non-nil, zero value otherwise.

### GetChangeOk

`func (o *InstitutionalOwnershipInfo) GetChangeOk() (*float32, bool)`

GetChangeOk returns a tuple with the Change field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChange

`func (o *InstitutionalOwnershipInfo) SetChange(v float32)`

SetChange sets Change field to given value.

### HasChange

`func (o *InstitutionalOwnershipInfo) HasChange() bool`

HasChange returns a boolean if a field has been set.

### GetNoVoting

`func (o *InstitutionalOwnershipInfo) GetNoVoting() float32`

GetNoVoting returns the NoVoting field if non-nil, zero value otherwise.

### GetNoVotingOk

`func (o *InstitutionalOwnershipInfo) GetNoVotingOk() (*float32, bool)`

GetNoVotingOk returns a tuple with the NoVoting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNoVoting

`func (o *InstitutionalOwnershipInfo) SetNoVoting(v float32)`

SetNoVoting sets NoVoting field to given value.

### HasNoVoting

`func (o *InstitutionalOwnershipInfo) HasNoVoting() bool`

HasNoVoting returns a boolean if a field has been set.

### GetPercentage

`func (o *InstitutionalOwnershipInfo) GetPercentage() float32`

GetPercentage returns the Percentage field if non-nil, zero value otherwise.

### GetPercentageOk

`func (o *InstitutionalOwnershipInfo) GetPercentageOk() (*float32, bool)`

GetPercentageOk returns a tuple with the Percentage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPercentage

`func (o *InstitutionalOwnershipInfo) SetPercentage(v float32)`

SetPercentage sets Percentage field to given value.

### HasPercentage

`func (o *InstitutionalOwnershipInfo) HasPercentage() bool`

HasPercentage returns a boolean if a field has been set.

### GetShare

`func (o *InstitutionalOwnershipInfo) GetShare() float32`

GetShare returns the Share field if non-nil, zero value otherwise.

### GetShareOk

`func (o *InstitutionalOwnershipInfo) GetShareOk() (*float32, bool)`

GetShareOk returns a tuple with the Share field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShare

`func (o *InstitutionalOwnershipInfo) SetShare(v float32)`

SetShare sets Share field to given value.

### HasShare

`func (o *InstitutionalOwnershipInfo) HasShare() bool`

HasShare returns a boolean if a field has been set.

### GetSharedVoting

`func (o *InstitutionalOwnershipInfo) GetSharedVoting() float32`

GetSharedVoting returns the SharedVoting field if non-nil, zero value otherwise.

### GetSharedVotingOk

`func (o *InstitutionalOwnershipInfo) GetSharedVotingOk() (*float32, bool)`

GetSharedVotingOk returns a tuple with the SharedVoting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharedVoting

`func (o *InstitutionalOwnershipInfo) SetSharedVoting(v float32)`

SetSharedVoting sets SharedVoting field to given value.

### HasSharedVoting

`func (o *InstitutionalOwnershipInfo) HasSharedVoting() bool`

HasSharedVoting returns a boolean if a field has been set.

### GetSoleVoting

`func (o *InstitutionalOwnershipInfo) GetSoleVoting() float32`

GetSoleVoting returns the SoleVoting field if non-nil, zero value otherwise.

### GetSoleVotingOk

`func (o *InstitutionalOwnershipInfo) GetSoleVotingOk() (*float32, bool)`

GetSoleVotingOk returns a tuple with the SoleVoting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSoleVoting

`func (o *InstitutionalOwnershipInfo) SetSoleVoting(v float32)`

SetSoleVoting sets SoleVoting field to given value.

### HasSoleVoting

`func (o *InstitutionalOwnershipInfo) HasSoleVoting() bool`

HasSoleVoting returns a boolean if a field has been set.

### GetValue

`func (o *InstitutionalOwnershipInfo) GetValue() float32`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *InstitutionalOwnershipInfo) GetValueOk() (*float32, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *InstitutionalOwnershipInfo) SetValue(v float32)`

SetValue sets Value field to given value.

### HasValue

`func (o *InstitutionalOwnershipInfo) HasValue() bool`

HasValue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


