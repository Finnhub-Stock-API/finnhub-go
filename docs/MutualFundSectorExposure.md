# MutualFundSectorExposure

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Symbol** | Pointer to **string** | Mutual symbol. | [optional] 
**SectorExposure** | Pointer to **[]map[string]interface{}** | Array of sector and exposure levels. | [optional] 

## Methods

### NewMutualFundSectorExposure

`func NewMutualFundSectorExposure() *MutualFundSectorExposure`

NewMutualFundSectorExposure instantiates a new MutualFundSectorExposure object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMutualFundSectorExposureWithDefaults

`func NewMutualFundSectorExposureWithDefaults() *MutualFundSectorExposure`

NewMutualFundSectorExposureWithDefaults instantiates a new MutualFundSectorExposure object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSymbol

`func (o *MutualFundSectorExposure) GetSymbol() string`

GetSymbol returns the Symbol field if non-nil, zero value otherwise.

### GetSymbolOk

`func (o *MutualFundSectorExposure) GetSymbolOk() (*string, bool)`

GetSymbolOk returns a tuple with the Symbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSymbol

`func (o *MutualFundSectorExposure) SetSymbol(v string)`

SetSymbol sets Symbol field to given value.

### HasSymbol

`func (o *MutualFundSectorExposure) HasSymbol() bool`

HasSymbol returns a boolean if a field has been set.

### GetSectorExposure

`func (o *MutualFundSectorExposure) GetSectorExposure() []map[string]interface{}`

GetSectorExposure returns the SectorExposure field if non-nil, zero value otherwise.

### GetSectorExposureOk

`func (o *MutualFundSectorExposure) GetSectorExposureOk() (*[]map[string]interface{}, bool)`

GetSectorExposureOk returns a tuple with the SectorExposure field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSectorExposure

`func (o *MutualFundSectorExposure) SetSectorExposure(v []map[string]interface{})`

SetSectorExposure sets SectorExposure field to given value.

### HasSectorExposure

`func (o *MutualFundSectorExposure) HasSectorExposure() bool`

HasSectorExposure returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


