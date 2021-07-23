# FilingSentiment

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Negative** | Pointer to **float32** | % of negative words in the filing. | [optional] 
**Positive** | Pointer to **float32** | % of positive words in the filing. | [optional] 
**Polarity** | Pointer to **float32** | % of polarity words in the filing. | [optional] 
**Litigious** | Pointer to **float32** | % of litigious words in the filing. | [optional] 
**Uncertainty** | Pointer to **float32** | % of uncertainty words in the filing. | [optional] 
**Constraining** | Pointer to **float32** | % of constraining words in the filing. | [optional] 
**ModalWeak** | Pointer to **float32** | % of modal-weak words in the filing. | [optional] 
**ModalStrong** | Pointer to **float32** | % of modal-strong words in the filing. | [optional] 
**ModalModerate** | Pointer to **float32** | % of modal-moderate words in the filing. | [optional] 

## Methods

### NewFilingSentiment

`func NewFilingSentiment() *FilingSentiment`

NewFilingSentiment instantiates a new FilingSentiment object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFilingSentimentWithDefaults

`func NewFilingSentimentWithDefaults() *FilingSentiment`

NewFilingSentimentWithDefaults instantiates a new FilingSentiment object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNegative

`func (o *FilingSentiment) GetNegative() float32`

GetNegative returns the Negative field if non-nil, zero value otherwise.

### GetNegativeOk

`func (o *FilingSentiment) GetNegativeOk() (*float32, bool)`

GetNegativeOk returns a tuple with the Negative field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNegative

`func (o *FilingSentiment) SetNegative(v float32)`

SetNegative sets Negative field to given value.

### HasNegative

`func (o *FilingSentiment) HasNegative() bool`

HasNegative returns a boolean if a field has been set.

### GetPositive

`func (o *FilingSentiment) GetPositive() float32`

GetPositive returns the Positive field if non-nil, zero value otherwise.

### GetPositiveOk

`func (o *FilingSentiment) GetPositiveOk() (*float32, bool)`

GetPositiveOk returns a tuple with the Positive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPositive

`func (o *FilingSentiment) SetPositive(v float32)`

SetPositive sets Positive field to given value.

### HasPositive

`func (o *FilingSentiment) HasPositive() bool`

HasPositive returns a boolean if a field has been set.

### GetPolarity

`func (o *FilingSentiment) GetPolarity() float32`

GetPolarity returns the Polarity field if non-nil, zero value otherwise.

### GetPolarityOk

`func (o *FilingSentiment) GetPolarityOk() (*float32, bool)`

GetPolarityOk returns a tuple with the Polarity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolarity

`func (o *FilingSentiment) SetPolarity(v float32)`

SetPolarity sets Polarity field to given value.

### HasPolarity

`func (o *FilingSentiment) HasPolarity() bool`

HasPolarity returns a boolean if a field has been set.

### GetLitigious

`func (o *FilingSentiment) GetLitigious() float32`

GetLitigious returns the Litigious field if non-nil, zero value otherwise.

### GetLitigiousOk

`func (o *FilingSentiment) GetLitigiousOk() (*float32, bool)`

GetLitigiousOk returns a tuple with the Litigious field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLitigious

`func (o *FilingSentiment) SetLitigious(v float32)`

SetLitigious sets Litigious field to given value.

### HasLitigious

`func (o *FilingSentiment) HasLitigious() bool`

HasLitigious returns a boolean if a field has been set.

### GetUncertainty

`func (o *FilingSentiment) GetUncertainty() float32`

GetUncertainty returns the Uncertainty field if non-nil, zero value otherwise.

### GetUncertaintyOk

`func (o *FilingSentiment) GetUncertaintyOk() (*float32, bool)`

GetUncertaintyOk returns a tuple with the Uncertainty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUncertainty

`func (o *FilingSentiment) SetUncertainty(v float32)`

SetUncertainty sets Uncertainty field to given value.

### HasUncertainty

`func (o *FilingSentiment) HasUncertainty() bool`

HasUncertainty returns a boolean if a field has been set.

### GetConstraining

`func (o *FilingSentiment) GetConstraining() float32`

GetConstraining returns the Constraining field if non-nil, zero value otherwise.

### GetConstrainingOk

`func (o *FilingSentiment) GetConstrainingOk() (*float32, bool)`

GetConstrainingOk returns a tuple with the Constraining field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConstraining

`func (o *FilingSentiment) SetConstraining(v float32)`

SetConstraining sets Constraining field to given value.

### HasConstraining

`func (o *FilingSentiment) HasConstraining() bool`

HasConstraining returns a boolean if a field has been set.

### GetModalWeak

`func (o *FilingSentiment) GetModalWeak() float32`

GetModalWeak returns the ModalWeak field if non-nil, zero value otherwise.

### GetModalWeakOk

`func (o *FilingSentiment) GetModalWeakOk() (*float32, bool)`

GetModalWeakOk returns a tuple with the ModalWeak field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModalWeak

`func (o *FilingSentiment) SetModalWeak(v float32)`

SetModalWeak sets ModalWeak field to given value.

### HasModalWeak

`func (o *FilingSentiment) HasModalWeak() bool`

HasModalWeak returns a boolean if a field has been set.

### GetModalStrong

`func (o *FilingSentiment) GetModalStrong() float32`

GetModalStrong returns the ModalStrong field if non-nil, zero value otherwise.

### GetModalStrongOk

`func (o *FilingSentiment) GetModalStrongOk() (*float32, bool)`

GetModalStrongOk returns a tuple with the ModalStrong field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModalStrong

`func (o *FilingSentiment) SetModalStrong(v float32)`

SetModalStrong sets ModalStrong field to given value.

### HasModalStrong

`func (o *FilingSentiment) HasModalStrong() bool`

HasModalStrong returns a boolean if a field has been set.

### GetModalModerate

`func (o *FilingSentiment) GetModalModerate() float32`

GetModalModerate returns the ModalModerate field if non-nil, zero value otherwise.

### GetModalModerateOk

`func (o *FilingSentiment) GetModalModerateOk() (*float32, bool)`

GetModalModerateOk returns a tuple with the ModalModerate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModalModerate

`func (o *FilingSentiment) SetModalModerate(v float32)`

SetModalModerate sets ModalModerate field to given value.

### HasModalModerate

`func (o *FilingSentiment) HasModalModerate() bool`

HasModalModerate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


