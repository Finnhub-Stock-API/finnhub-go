# PriceTarget

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Symbol** | Pointer to **string** | Company symbol. | [optional] 
**TargetHigh** | Pointer to **float32** | Highes analysts&#39; target. | [optional] 
**TargetLow** | Pointer to **float32** | Lowest analysts&#39; target. | [optional] 
**TargetMean** | Pointer to **float32** | Mean of all analysts&#39; targets. | [optional] 
**TargetMedian** | Pointer to **float32** | Median of all analysts&#39; targets. | [optional] 
**NumberAnalysts** | Pointer to **int64** | Number of Analysts. | [optional] 
**LastUpdated** | Pointer to **string** | Updated time of the data | [optional] 

## Methods

### NewPriceTarget

`func NewPriceTarget() *PriceTarget`

NewPriceTarget instantiates a new PriceTarget object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPriceTargetWithDefaults

`func NewPriceTargetWithDefaults() *PriceTarget`

NewPriceTargetWithDefaults instantiates a new PriceTarget object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSymbol

`func (o *PriceTarget) GetSymbol() string`

GetSymbol returns the Symbol field if non-nil, zero value otherwise.

### GetSymbolOk

`func (o *PriceTarget) GetSymbolOk() (*string, bool)`

GetSymbolOk returns a tuple with the Symbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSymbol

`func (o *PriceTarget) SetSymbol(v string)`

SetSymbol sets Symbol field to given value.

### HasSymbol

`func (o *PriceTarget) HasSymbol() bool`

HasSymbol returns a boolean if a field has been set.

### GetTargetHigh

`func (o *PriceTarget) GetTargetHigh() float32`

GetTargetHigh returns the TargetHigh field if non-nil, zero value otherwise.

### GetTargetHighOk

`func (o *PriceTarget) GetTargetHighOk() (*float32, bool)`

GetTargetHighOk returns a tuple with the TargetHigh field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetHigh

`func (o *PriceTarget) SetTargetHigh(v float32)`

SetTargetHigh sets TargetHigh field to given value.

### HasTargetHigh

`func (o *PriceTarget) HasTargetHigh() bool`

HasTargetHigh returns a boolean if a field has been set.

### GetTargetLow

`func (o *PriceTarget) GetTargetLow() float32`

GetTargetLow returns the TargetLow field if non-nil, zero value otherwise.

### GetTargetLowOk

`func (o *PriceTarget) GetTargetLowOk() (*float32, bool)`

GetTargetLowOk returns a tuple with the TargetLow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetLow

`func (o *PriceTarget) SetTargetLow(v float32)`

SetTargetLow sets TargetLow field to given value.

### HasTargetLow

`func (o *PriceTarget) HasTargetLow() bool`

HasTargetLow returns a boolean if a field has been set.

### GetTargetMean

`func (o *PriceTarget) GetTargetMean() float32`

GetTargetMean returns the TargetMean field if non-nil, zero value otherwise.

### GetTargetMeanOk

`func (o *PriceTarget) GetTargetMeanOk() (*float32, bool)`

GetTargetMeanOk returns a tuple with the TargetMean field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetMean

`func (o *PriceTarget) SetTargetMean(v float32)`

SetTargetMean sets TargetMean field to given value.

### HasTargetMean

`func (o *PriceTarget) HasTargetMean() bool`

HasTargetMean returns a boolean if a field has been set.

### GetTargetMedian

`func (o *PriceTarget) GetTargetMedian() float32`

GetTargetMedian returns the TargetMedian field if non-nil, zero value otherwise.

### GetTargetMedianOk

`func (o *PriceTarget) GetTargetMedianOk() (*float32, bool)`

GetTargetMedianOk returns a tuple with the TargetMedian field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetMedian

`func (o *PriceTarget) SetTargetMedian(v float32)`

SetTargetMedian sets TargetMedian field to given value.

### HasTargetMedian

`func (o *PriceTarget) HasTargetMedian() bool`

HasTargetMedian returns a boolean if a field has been set.

### GetNumberAnalysts

`func (o *PriceTarget) GetNumberAnalysts() int64`

GetNumberAnalysts returns the NumberAnalysts field if non-nil, zero value otherwise.

### GetNumberAnalystsOk

`func (o *PriceTarget) GetNumberAnalystsOk() (*int64, bool)`

GetNumberAnalystsOk returns a tuple with the NumberAnalysts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumberAnalysts

`func (o *PriceTarget) SetNumberAnalysts(v int64)`

SetNumberAnalysts sets NumberAnalysts field to given value.

### HasNumberAnalysts

`func (o *PriceTarget) HasNumberAnalysts() bool`

HasNumberAnalysts returns a boolean if a field has been set.

### GetLastUpdated

`func (o *PriceTarget) GetLastUpdated() string`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *PriceTarget) GetLastUpdatedOk() (*string, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *PriceTarget) SetLastUpdated(v string)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *PriceTarget) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


