# ETFsCountryExposure

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Symbol** | Pointer to **string** | ETF symbol. | [optional] 
**CountryExposure** | Pointer to **[]map[string]interface{}** | Array of countries and and exposure levels. | [optional] 

## Methods

### NewETFsCountryExposure

`func NewETFsCountryExposure() *ETFsCountryExposure`

NewETFsCountryExposure instantiates a new ETFsCountryExposure object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewETFsCountryExposureWithDefaults

`func NewETFsCountryExposureWithDefaults() *ETFsCountryExposure`

NewETFsCountryExposureWithDefaults instantiates a new ETFsCountryExposure object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSymbol

`func (o *ETFsCountryExposure) GetSymbol() string`

GetSymbol returns the Symbol field if non-nil, zero value otherwise.

### GetSymbolOk

`func (o *ETFsCountryExposure) GetSymbolOk() (*string, bool)`

GetSymbolOk returns a tuple with the Symbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSymbol

`func (o *ETFsCountryExposure) SetSymbol(v string)`

SetSymbol sets Symbol field to given value.

### HasSymbol

`func (o *ETFsCountryExposure) HasSymbol() bool`

HasSymbol returns a boolean if a field has been set.

### GetCountryExposure

`func (o *ETFsCountryExposure) GetCountryExposure() []map[string]interface{}`

GetCountryExposure returns the CountryExposure field if non-nil, zero value otherwise.

### GetCountryExposureOk

`func (o *ETFsCountryExposure) GetCountryExposureOk() (*[]map[string]interface{}, bool)`

GetCountryExposureOk returns a tuple with the CountryExposure field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryExposure

`func (o *ETFsCountryExposure) SetCountryExposure(v []map[string]interface{})`

SetCountryExposure sets CountryExposure field to given value.

### HasCountryExposure

`func (o *ETFsCountryExposure) HasCountryExposure() bool`

HasCountryExposure returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


