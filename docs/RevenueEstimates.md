# RevenueEstimates

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to **[]map[string]interface{}** | List of estimates | [optional] 
**Freq** | Pointer to **string** | Frequency: annual or quarterly. | [optional] 
**Symbol** | Pointer to **string** | Company symbol. | [optional] 

## Methods

### NewRevenueEstimates

`func NewRevenueEstimates() *RevenueEstimates`

NewRevenueEstimates instantiates a new RevenueEstimates object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRevenueEstimatesWithDefaults

`func NewRevenueEstimatesWithDefaults() *RevenueEstimates`

NewRevenueEstimatesWithDefaults instantiates a new RevenueEstimates object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *RevenueEstimates) GetData() []map[string]interface{}`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *RevenueEstimates) GetDataOk() (*[]map[string]interface{}, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *RevenueEstimates) SetData(v []map[string]interface{})`

SetData sets Data field to given value.

### HasData

`func (o *RevenueEstimates) HasData() bool`

HasData returns a boolean if a field has been set.

### GetFreq

`func (o *RevenueEstimates) GetFreq() string`

GetFreq returns the Freq field if non-nil, zero value otherwise.

### GetFreqOk

`func (o *RevenueEstimates) GetFreqOk() (*string, bool)`

GetFreqOk returns a tuple with the Freq field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFreq

`func (o *RevenueEstimates) SetFreq(v string)`

SetFreq sets Freq field to given value.

### HasFreq

`func (o *RevenueEstimates) HasFreq() bool`

HasFreq returns a boolean if a field has been set.

### GetSymbol

`func (o *RevenueEstimates) GetSymbol() string`

GetSymbol returns the Symbol field if non-nil, zero value otherwise.

### GetSymbolOk

`func (o *RevenueEstimates) GetSymbolOk() (*string, bool)`

GetSymbolOk returns a tuple with the Symbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSymbol

`func (o *RevenueEstimates) SetSymbol(v string)`

SetSymbol sets Symbol field to given value.

### HasSymbol

`func (o *RevenueEstimates) HasSymbol() bool`

HasSymbol returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


