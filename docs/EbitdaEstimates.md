# EbitdaEstimates

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**[]EbitdaEstimatesInfo**](EbitdaEstimatesInfo.md) | List of estimates | [optional] 
**Freq** | Pointer to **string** | Frequency: annual or quarterly. | [optional] 
**Symbol** | Pointer to **string** | Company symbol. | [optional] 

## Methods

### NewEbitdaEstimates

`func NewEbitdaEstimates() *EbitdaEstimates`

NewEbitdaEstimates instantiates a new EbitdaEstimates object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEbitdaEstimatesWithDefaults

`func NewEbitdaEstimatesWithDefaults() *EbitdaEstimates`

NewEbitdaEstimatesWithDefaults instantiates a new EbitdaEstimates object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *EbitdaEstimates) GetData() []EbitdaEstimatesInfo`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *EbitdaEstimates) GetDataOk() (*[]EbitdaEstimatesInfo, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *EbitdaEstimates) SetData(v []EbitdaEstimatesInfo)`

SetData sets Data field to given value.

### HasData

`func (o *EbitdaEstimates) HasData() bool`

HasData returns a boolean if a field has been set.

### GetFreq

`func (o *EbitdaEstimates) GetFreq() string`

GetFreq returns the Freq field if non-nil, zero value otherwise.

### GetFreqOk

`func (o *EbitdaEstimates) GetFreqOk() (*string, bool)`

GetFreqOk returns a tuple with the Freq field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFreq

`func (o *EbitdaEstimates) SetFreq(v string)`

SetFreq sets Freq field to given value.

### HasFreq

`func (o *EbitdaEstimates) HasFreq() bool`

HasFreq returns a boolean if a field has been set.

### GetSymbol

`func (o *EbitdaEstimates) GetSymbol() string`

GetSymbol returns the Symbol field if non-nil, zero value otherwise.

### GetSymbolOk

`func (o *EbitdaEstimates) GetSymbolOk() (*string, bool)`

GetSymbolOk returns a tuple with the Symbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSymbol

`func (o *EbitdaEstimates) SetSymbol(v string)`

SetSymbol sets Symbol field to given value.

### HasSymbol

`func (o *EbitdaEstimates) HasSymbol() bool`

HasSymbol returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


