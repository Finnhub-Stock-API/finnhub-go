# MutualFundCountryExposure

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Symbol** | Pointer to **string** | Symbol. | [optional] 
**CountryExposure** | Pointer to [**[]MutualFundCountryExposureData**](MutualFundCountryExposureData.md) | Array of countries and and exposure levels. | [optional] 

## Methods

### NewMutualFundCountryExposure

`func NewMutualFundCountryExposure() *MutualFundCountryExposure`

NewMutualFundCountryExposure instantiates a new MutualFundCountryExposure object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMutualFundCountryExposureWithDefaults

`func NewMutualFundCountryExposureWithDefaults() *MutualFundCountryExposure`

NewMutualFundCountryExposureWithDefaults instantiates a new MutualFundCountryExposure object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSymbol

`func (o *MutualFundCountryExposure) GetSymbol() string`

GetSymbol returns the Symbol field if non-nil, zero value otherwise.

### GetSymbolOk

`func (o *MutualFundCountryExposure) GetSymbolOk() (*string, bool)`

GetSymbolOk returns a tuple with the Symbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSymbol

`func (o *MutualFundCountryExposure) SetSymbol(v string)`

SetSymbol sets Symbol field to given value.

### HasSymbol

`func (o *MutualFundCountryExposure) HasSymbol() bool`

HasSymbol returns a boolean if a field has been set.

### GetCountryExposure

`func (o *MutualFundCountryExposure) GetCountryExposure() []MutualFundCountryExposureData`

GetCountryExposure returns the CountryExposure field if non-nil, zero value otherwise.

### GetCountryExposureOk

`func (o *MutualFundCountryExposure) GetCountryExposureOk() (*[]MutualFundCountryExposureData, bool)`

GetCountryExposureOk returns a tuple with the CountryExposure field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryExposure

`func (o *MutualFundCountryExposure) SetCountryExposure(v []MutualFundCountryExposureData)`

SetCountryExposure sets CountryExposure field to given value.

### HasCountryExposure

`func (o *MutualFundCountryExposure) HasCountryExposure() bool`

HasCountryExposure returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


