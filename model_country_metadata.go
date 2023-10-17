/*
Finnhub API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package finnhub

import (
	"encoding/json"
)

// CountryMetadata struct for CountryMetadata
type CountryMetadata struct {
	// Country name
	Country *string `json:"country,omitempty"`
	// Alpha 2 code
	Code2 *string `json:"code2,omitempty"`
	// Alpha 3 code
	Code3 *string `json:"code3,omitempty"`
	// UN code
	CodeNo *string `json:"codeNo,omitempty"`
	// Currency name
	Currency *string `json:"currency,omitempty"`
	// Currency code
	CurrencyCode *string `json:"currencyCode,omitempty"`
	// Region
	Region *string `json:"region,omitempty"`
	// Sub-Region
	SubRegion *string `json:"subRegion,omitempty"`
	// Moody's credit risk rating.
	Rating *string `json:"rating,omitempty"`
	// Default spread
	DefaultSpread *float32 `json:"defaultSpread,omitempty"`
	// Country risk premium
	CountryRiskPremium *float32 `json:"countryRiskPremium,omitempty"`
	// Equity risk premium
	EquityRiskPremium *float32 `json:"equityRiskPremium,omitempty"`
}

// NewCountryMetadata instantiates a new CountryMetadata object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCountryMetadata() *CountryMetadata {
	this := CountryMetadata{}
	return &this
}

// NewCountryMetadataWithDefaults instantiates a new CountryMetadata object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCountryMetadataWithDefaults() *CountryMetadata {
	this := CountryMetadata{}
	return &this
}

// GetCountry returns the Country field value if set, zero value otherwise.
func (o *CountryMetadata) GetCountry() string {
	if o == nil || o.Country == nil {
		var ret string
		return ret
	}
	return *o.Country
}

// GetCountryOk returns a tuple with the Country field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CountryMetadata) GetCountryOk() (*string, bool) {
	if o == nil || o.Country == nil {
		return nil, false
	}
	return o.Country, true
}

// HasCountry returns a boolean if a field has been set.
func (o *CountryMetadata) HasCountry() bool {
	if o != nil && o.Country != nil {
		return true
	}

	return false
}

// SetCountry gets a reference to the given string and assigns it to the Country field.
func (o *CountryMetadata) SetCountry(v string) {
	o.Country = &v
}

// GetCode2 returns the Code2 field value if set, zero value otherwise.
func (o *CountryMetadata) GetCode2() string {
	if o == nil || o.Code2 == nil {
		var ret string
		return ret
	}
	return *o.Code2
}

// GetCode2Ok returns a tuple with the Code2 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CountryMetadata) GetCode2Ok() (*string, bool) {
	if o == nil || o.Code2 == nil {
		return nil, false
	}
	return o.Code2, true
}

// HasCode2 returns a boolean if a field has been set.
func (o *CountryMetadata) HasCode2() bool {
	if o != nil && o.Code2 != nil {
		return true
	}

	return false
}

// SetCode2 gets a reference to the given string and assigns it to the Code2 field.
func (o *CountryMetadata) SetCode2(v string) {
	o.Code2 = &v
}

// GetCode3 returns the Code3 field value if set, zero value otherwise.
func (o *CountryMetadata) GetCode3() string {
	if o == nil || o.Code3 == nil {
		var ret string
		return ret
	}
	return *o.Code3
}

// GetCode3Ok returns a tuple with the Code3 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CountryMetadata) GetCode3Ok() (*string, bool) {
	if o == nil || o.Code3 == nil {
		return nil, false
	}
	return o.Code3, true
}

// HasCode3 returns a boolean if a field has been set.
func (o *CountryMetadata) HasCode3() bool {
	if o != nil && o.Code3 != nil {
		return true
	}

	return false
}

// SetCode3 gets a reference to the given string and assigns it to the Code3 field.
func (o *CountryMetadata) SetCode3(v string) {
	o.Code3 = &v
}

// GetCodeNo returns the CodeNo field value if set, zero value otherwise.
func (o *CountryMetadata) GetCodeNo() string {
	if o == nil || o.CodeNo == nil {
		var ret string
		return ret
	}
	return *o.CodeNo
}

// GetCodeNoOk returns a tuple with the CodeNo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CountryMetadata) GetCodeNoOk() (*string, bool) {
	if o == nil || o.CodeNo == nil {
		return nil, false
	}
	return o.CodeNo, true
}

// HasCodeNo returns a boolean if a field has been set.
func (o *CountryMetadata) HasCodeNo() bool {
	if o != nil && o.CodeNo != nil {
		return true
	}

	return false
}

// SetCodeNo gets a reference to the given string and assigns it to the CodeNo field.
func (o *CountryMetadata) SetCodeNo(v string) {
	o.CodeNo = &v
}

// GetCurrency returns the Currency field value if set, zero value otherwise.
func (o *CountryMetadata) GetCurrency() string {
	if o == nil || o.Currency == nil {
		var ret string
		return ret
	}
	return *o.Currency
}

// GetCurrencyOk returns a tuple with the Currency field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CountryMetadata) GetCurrencyOk() (*string, bool) {
	if o == nil || o.Currency == nil {
		return nil, false
	}
	return o.Currency, true
}

// HasCurrency returns a boolean if a field has been set.
func (o *CountryMetadata) HasCurrency() bool {
	if o != nil && o.Currency != nil {
		return true
	}

	return false
}

// SetCurrency gets a reference to the given string and assigns it to the Currency field.
func (o *CountryMetadata) SetCurrency(v string) {
	o.Currency = &v
}

// GetCurrencyCode returns the CurrencyCode field value if set, zero value otherwise.
func (o *CountryMetadata) GetCurrencyCode() string {
	if o == nil || o.CurrencyCode == nil {
		var ret string
		return ret
	}
	return *o.CurrencyCode
}

// GetCurrencyCodeOk returns a tuple with the CurrencyCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CountryMetadata) GetCurrencyCodeOk() (*string, bool) {
	if o == nil || o.CurrencyCode == nil {
		return nil, false
	}
	return o.CurrencyCode, true
}

// HasCurrencyCode returns a boolean if a field has been set.
func (o *CountryMetadata) HasCurrencyCode() bool {
	if o != nil && o.CurrencyCode != nil {
		return true
	}

	return false
}

// SetCurrencyCode gets a reference to the given string and assigns it to the CurrencyCode field.
func (o *CountryMetadata) SetCurrencyCode(v string) {
	o.CurrencyCode = &v
}

// GetRegion returns the Region field value if set, zero value otherwise.
func (o *CountryMetadata) GetRegion() string {
	if o == nil || o.Region == nil {
		var ret string
		return ret
	}
	return *o.Region
}

// GetRegionOk returns a tuple with the Region field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CountryMetadata) GetRegionOk() (*string, bool) {
	if o == nil || o.Region == nil {
		return nil, false
	}
	return o.Region, true
}

// HasRegion returns a boolean if a field has been set.
func (o *CountryMetadata) HasRegion() bool {
	if o != nil && o.Region != nil {
		return true
	}

	return false
}

// SetRegion gets a reference to the given string and assigns it to the Region field.
func (o *CountryMetadata) SetRegion(v string) {
	o.Region = &v
}

// GetSubRegion returns the SubRegion field value if set, zero value otherwise.
func (o *CountryMetadata) GetSubRegion() string {
	if o == nil || o.SubRegion == nil {
		var ret string
		return ret
	}
	return *o.SubRegion
}

// GetSubRegionOk returns a tuple with the SubRegion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CountryMetadata) GetSubRegionOk() (*string, bool) {
	if o == nil || o.SubRegion == nil {
		return nil, false
	}
	return o.SubRegion, true
}

// HasSubRegion returns a boolean if a field has been set.
func (o *CountryMetadata) HasSubRegion() bool {
	if o != nil && o.SubRegion != nil {
		return true
	}

	return false
}

// SetSubRegion gets a reference to the given string and assigns it to the SubRegion field.
func (o *CountryMetadata) SetSubRegion(v string) {
	o.SubRegion = &v
}

// GetRating returns the Rating field value if set, zero value otherwise.
func (o *CountryMetadata) GetRating() string {
	if o == nil || o.Rating == nil {
		var ret string
		return ret
	}
	return *o.Rating
}

// GetRatingOk returns a tuple with the Rating field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CountryMetadata) GetRatingOk() (*string, bool) {
	if o == nil || o.Rating == nil {
		return nil, false
	}
	return o.Rating, true
}

// HasRating returns a boolean if a field has been set.
func (o *CountryMetadata) HasRating() bool {
	if o != nil && o.Rating != nil {
		return true
	}

	return false
}

// SetRating gets a reference to the given string and assigns it to the Rating field.
func (o *CountryMetadata) SetRating(v string) {
	o.Rating = &v
}

// GetDefaultSpread returns the DefaultSpread field value if set, zero value otherwise.
func (o *CountryMetadata) GetDefaultSpread() float32 {
	if o == nil || o.DefaultSpread == nil {
		var ret float32
		return ret
	}
	return *o.DefaultSpread
}

// GetDefaultSpreadOk returns a tuple with the DefaultSpread field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CountryMetadata) GetDefaultSpreadOk() (*float32, bool) {
	if o == nil || o.DefaultSpread == nil {
		return nil, false
	}
	return o.DefaultSpread, true
}

// HasDefaultSpread returns a boolean if a field has been set.
func (o *CountryMetadata) HasDefaultSpread() bool {
	if o != nil && o.DefaultSpread != nil {
		return true
	}

	return false
}

// SetDefaultSpread gets a reference to the given float32 and assigns it to the DefaultSpread field.
func (o *CountryMetadata) SetDefaultSpread(v float32) {
	o.DefaultSpread = &v
}

// GetCountryRiskPremium returns the CountryRiskPremium field value if set, zero value otherwise.
func (o *CountryMetadata) GetCountryRiskPremium() float32 {
	if o == nil || o.CountryRiskPremium == nil {
		var ret float32
		return ret
	}
	return *o.CountryRiskPremium
}

// GetCountryRiskPremiumOk returns a tuple with the CountryRiskPremium field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CountryMetadata) GetCountryRiskPremiumOk() (*float32, bool) {
	if o == nil || o.CountryRiskPremium == nil {
		return nil, false
	}
	return o.CountryRiskPremium, true
}

// HasCountryRiskPremium returns a boolean if a field has been set.
func (o *CountryMetadata) HasCountryRiskPremium() bool {
	if o != nil && o.CountryRiskPremium != nil {
		return true
	}

	return false
}

// SetCountryRiskPremium gets a reference to the given float32 and assigns it to the CountryRiskPremium field.
func (o *CountryMetadata) SetCountryRiskPremium(v float32) {
	o.CountryRiskPremium = &v
}

// GetEquityRiskPremium returns the EquityRiskPremium field value if set, zero value otherwise.
func (o *CountryMetadata) GetEquityRiskPremium() float32 {
	if o == nil || o.EquityRiskPremium == nil {
		var ret float32
		return ret
	}
	return *o.EquityRiskPremium
}

// GetEquityRiskPremiumOk returns a tuple with the EquityRiskPremium field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CountryMetadata) GetEquityRiskPremiumOk() (*float32, bool) {
	if o == nil || o.EquityRiskPremium == nil {
		return nil, false
	}
	return o.EquityRiskPremium, true
}

// HasEquityRiskPremium returns a boolean if a field has been set.
func (o *CountryMetadata) HasEquityRiskPremium() bool {
	if o != nil && o.EquityRiskPremium != nil {
		return true
	}

	return false
}

// SetEquityRiskPremium gets a reference to the given float32 and assigns it to the EquityRiskPremium field.
func (o *CountryMetadata) SetEquityRiskPremium(v float32) {
	o.EquityRiskPremium = &v
}

func (o CountryMetadata) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Country != nil {
		toSerialize["country"] = o.Country
	}
	if o.Code2 != nil {
		toSerialize["code2"] = o.Code2
	}
	if o.Code3 != nil {
		toSerialize["code3"] = o.Code3
	}
	if o.CodeNo != nil {
		toSerialize["codeNo"] = o.CodeNo
	}
	if o.Currency != nil {
		toSerialize["currency"] = o.Currency
	}
	if o.CurrencyCode != nil {
		toSerialize["currencyCode"] = o.CurrencyCode
	}
	if o.Region != nil {
		toSerialize["region"] = o.Region
	}
	if o.SubRegion != nil {
		toSerialize["subRegion"] = o.SubRegion
	}
	if o.Rating != nil {
		toSerialize["rating"] = o.Rating
	}
	if o.DefaultSpread != nil {
		toSerialize["defaultSpread"] = o.DefaultSpread
	}
	if o.CountryRiskPremium != nil {
		toSerialize["countryRiskPremium"] = o.CountryRiskPremium
	}
	if o.EquityRiskPremium != nil {
		toSerialize["equityRiskPremium"] = o.EquityRiskPremium
	}
	return json.Marshal(toSerialize)
}

type NullableCountryMetadata struct {
	value *CountryMetadata
	isSet bool
}

func (v NullableCountryMetadata) Get() *CountryMetadata {
	return v.value
}

func (v *NullableCountryMetadata) Set(val *CountryMetadata) {
	v.value = val
	v.isSet = true
}

func (v NullableCountryMetadata) IsSet() bool {
	return v.isSet
}

func (v *NullableCountryMetadata) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCountryMetadata(val *CountryMetadata) *NullableCountryMetadata {
	return &NullableCountryMetadata{value: val, isSet: true}
}

func (v NullableCountryMetadata) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCountryMetadata) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


