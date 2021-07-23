# EconomicData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to **[]map[string]interface{}** | Array of economic data for requested code. | [optional] 
**Code** | Pointer to **string** | Finnhub economic code | [optional] 

## Methods

### NewEconomicData

`func NewEconomicData() *EconomicData`

NewEconomicData instantiates a new EconomicData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEconomicDataWithDefaults

`func NewEconomicDataWithDefaults() *EconomicData`

NewEconomicDataWithDefaults instantiates a new EconomicData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *EconomicData) GetData() []map[string]interface{}`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *EconomicData) GetDataOk() (*[]map[string]interface{}, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *EconomicData) SetData(v []map[string]interface{})`

SetData sets Data field to given value.

### HasData

`func (o *EconomicData) HasData() bool`

HasData returns a boolean if a field has been set.

### GetCode

`func (o *EconomicData) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *EconomicData) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *EconomicData) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *EconomicData) HasCode() bool`

HasCode returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


