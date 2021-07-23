# ETFsSectorExposure

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Symbol** | Pointer to **string** | ETF symbol. | [optional] 
**SectorExposure** | Pointer to **[]map[string]interface{}** | Array of industries and exposure levels. | [optional] 

## Methods

### NewETFsSectorExposure

`func NewETFsSectorExposure() *ETFsSectorExposure`

NewETFsSectorExposure instantiates a new ETFsSectorExposure object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewETFsSectorExposureWithDefaults

`func NewETFsSectorExposureWithDefaults() *ETFsSectorExposure`

NewETFsSectorExposureWithDefaults instantiates a new ETFsSectorExposure object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSymbol

`func (o *ETFsSectorExposure) GetSymbol() string`

GetSymbol returns the Symbol field if non-nil, zero value otherwise.

### GetSymbolOk

`func (o *ETFsSectorExposure) GetSymbolOk() (*string, bool)`

GetSymbolOk returns a tuple with the Symbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSymbol

`func (o *ETFsSectorExposure) SetSymbol(v string)`

SetSymbol sets Symbol field to given value.

### HasSymbol

`func (o *ETFsSectorExposure) HasSymbol() bool`

HasSymbol returns a boolean if a field has been set.

### GetSectorExposure

`func (o *ETFsSectorExposure) GetSectorExposure() []map[string]interface{}`

GetSectorExposure returns the SectorExposure field if non-nil, zero value otherwise.

### GetSectorExposureOk

`func (o *ETFsSectorExposure) GetSectorExposureOk() (*[]map[string]interface{}, bool)`

GetSectorExposureOk returns a tuple with the SectorExposure field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSectorExposure

`func (o *ETFsSectorExposure) SetSectorExposure(v []map[string]interface{})`

SetSectorExposure sets SectorExposure field to given value.

### HasSectorExposure

`func (o *ETFsSectorExposure) HasSectorExposure() bool`

HasSectorExposure returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


