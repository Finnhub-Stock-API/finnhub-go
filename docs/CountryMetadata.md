# CountryMetadata

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Country** | Pointer to **string** | Country name | [optional] 
**Code2** | Pointer to **string** | Alpha 2 code | [optional] 
**Code3** | Pointer to **string** | Alpha 3 code | [optional] 
**CodeNo** | Pointer to **string** | UN code | [optional] 
**Currency** | Pointer to **string** | Currency name | [optional] 
**CurrencyCode** | Pointer to **string** | Currency code | [optional] 
**Region** | Pointer to **string** | Region | [optional] 
**SubRegion** | Pointer to **string** | Sub-Region | [optional] 
**Rating** | Pointer to **string** | Moody&#39;s credit risk rating. | [optional] 
**DefaultSpread** | Pointer to **float32** | Default spread | [optional] 
**CountryRiskPremium** | Pointer to **float32** | Country risk premium | [optional] 
**EquityRiskPremium** | Pointer to **float32** | Equity risk premium | [optional] 

## Methods

### NewCountryMetadata

`func NewCountryMetadata() *CountryMetadata`

NewCountryMetadata instantiates a new CountryMetadata object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCountryMetadataWithDefaults

`func NewCountryMetadataWithDefaults() *CountryMetadata`

NewCountryMetadataWithDefaults instantiates a new CountryMetadata object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCountry

`func (o *CountryMetadata) GetCountry() string`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *CountryMetadata) GetCountryOk() (*string, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *CountryMetadata) SetCountry(v string)`

SetCountry sets Country field to given value.

### HasCountry

`func (o *CountryMetadata) HasCountry() bool`

HasCountry returns a boolean if a field has been set.

### GetCode2

`func (o *CountryMetadata) GetCode2() string`

GetCode2 returns the Code2 field if non-nil, zero value otherwise.

### GetCode2Ok

`func (o *CountryMetadata) GetCode2Ok() (*string, bool)`

GetCode2Ok returns a tuple with the Code2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode2

`func (o *CountryMetadata) SetCode2(v string)`

SetCode2 sets Code2 field to given value.

### HasCode2

`func (o *CountryMetadata) HasCode2() bool`

HasCode2 returns a boolean if a field has been set.

### GetCode3

`func (o *CountryMetadata) GetCode3() string`

GetCode3 returns the Code3 field if non-nil, zero value otherwise.

### GetCode3Ok

`func (o *CountryMetadata) GetCode3Ok() (*string, bool)`

GetCode3Ok returns a tuple with the Code3 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode3

`func (o *CountryMetadata) SetCode3(v string)`

SetCode3 sets Code3 field to given value.

### HasCode3

`func (o *CountryMetadata) HasCode3() bool`

HasCode3 returns a boolean if a field has been set.

### GetCodeNo

`func (o *CountryMetadata) GetCodeNo() string`

GetCodeNo returns the CodeNo field if non-nil, zero value otherwise.

### GetCodeNoOk

`func (o *CountryMetadata) GetCodeNoOk() (*string, bool)`

GetCodeNoOk returns a tuple with the CodeNo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCodeNo

`func (o *CountryMetadata) SetCodeNo(v string)`

SetCodeNo sets CodeNo field to given value.

### HasCodeNo

`func (o *CountryMetadata) HasCodeNo() bool`

HasCodeNo returns a boolean if a field has been set.

### GetCurrency

`func (o *CountryMetadata) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *CountryMetadata) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *CountryMetadata) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *CountryMetadata) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetCurrencyCode

`func (o *CountryMetadata) GetCurrencyCode() string`

GetCurrencyCode returns the CurrencyCode field if non-nil, zero value otherwise.

### GetCurrencyCodeOk

`func (o *CountryMetadata) GetCurrencyCodeOk() (*string, bool)`

GetCurrencyCodeOk returns a tuple with the CurrencyCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencyCode

`func (o *CountryMetadata) SetCurrencyCode(v string)`

SetCurrencyCode sets CurrencyCode field to given value.

### HasCurrencyCode

`func (o *CountryMetadata) HasCurrencyCode() bool`

HasCurrencyCode returns a boolean if a field has been set.

### GetRegion

`func (o *CountryMetadata) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *CountryMetadata) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *CountryMetadata) SetRegion(v string)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *CountryMetadata) HasRegion() bool`

HasRegion returns a boolean if a field has been set.

### GetSubRegion

`func (o *CountryMetadata) GetSubRegion() string`

GetSubRegion returns the SubRegion field if non-nil, zero value otherwise.

### GetSubRegionOk

`func (o *CountryMetadata) GetSubRegionOk() (*string, bool)`

GetSubRegionOk returns a tuple with the SubRegion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubRegion

`func (o *CountryMetadata) SetSubRegion(v string)`

SetSubRegion sets SubRegion field to given value.

### HasSubRegion

`func (o *CountryMetadata) HasSubRegion() bool`

HasSubRegion returns a boolean if a field has been set.

### GetRating

`func (o *CountryMetadata) GetRating() string`

GetRating returns the Rating field if non-nil, zero value otherwise.

### GetRatingOk

`func (o *CountryMetadata) GetRatingOk() (*string, bool)`

GetRatingOk returns a tuple with the Rating field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRating

`func (o *CountryMetadata) SetRating(v string)`

SetRating sets Rating field to given value.

### HasRating

`func (o *CountryMetadata) HasRating() bool`

HasRating returns a boolean if a field has been set.

### GetDefaultSpread

`func (o *CountryMetadata) GetDefaultSpread() float32`

GetDefaultSpread returns the DefaultSpread field if non-nil, zero value otherwise.

### GetDefaultSpreadOk

`func (o *CountryMetadata) GetDefaultSpreadOk() (*float32, bool)`

GetDefaultSpreadOk returns a tuple with the DefaultSpread field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultSpread

`func (o *CountryMetadata) SetDefaultSpread(v float32)`

SetDefaultSpread sets DefaultSpread field to given value.

### HasDefaultSpread

`func (o *CountryMetadata) HasDefaultSpread() bool`

HasDefaultSpread returns a boolean if a field has been set.

### GetCountryRiskPremium

`func (o *CountryMetadata) GetCountryRiskPremium() float32`

GetCountryRiskPremium returns the CountryRiskPremium field if non-nil, zero value otherwise.

### GetCountryRiskPremiumOk

`func (o *CountryMetadata) GetCountryRiskPremiumOk() (*float32, bool)`

GetCountryRiskPremiumOk returns a tuple with the CountryRiskPremium field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryRiskPremium

`func (o *CountryMetadata) SetCountryRiskPremium(v float32)`

SetCountryRiskPremium sets CountryRiskPremium field to given value.

### HasCountryRiskPremium

`func (o *CountryMetadata) HasCountryRiskPremium() bool`

HasCountryRiskPremium returns a boolean if a field has been set.

### GetEquityRiskPremium

`func (o *CountryMetadata) GetEquityRiskPremium() float32`

GetEquityRiskPremium returns the EquityRiskPremium field if non-nil, zero value otherwise.

### GetEquityRiskPremiumOk

`func (o *CountryMetadata) GetEquityRiskPremiumOk() (*float32, bool)`

GetEquityRiskPremiumOk returns a tuple with the EquityRiskPremium field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEquityRiskPremium

`func (o *CountryMetadata) SetEquityRiskPremium(v float32)`

SetEquityRiskPremium sets EquityRiskPremium field to given value.

### HasEquityRiskPremium

`func (o *CountryMetadata) HasEquityRiskPremium() bool`

HasEquityRiskPremium returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


