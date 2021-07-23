/*
 * Finnhub API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 1.0.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package finnhub

import (
	"encoding/json"
)

// ETFHoldingsData struct for ETFHoldingsData
type ETFHoldingsData struct {
	// Symbol description
	Symbol *string `json:"symbol,omitempty"`
	// Security name
	Name *string `json:"name,omitempty"`
	// ISIN.
	Isin *string `json:"isin,omitempty"`
	// CUSIP.
	Cusip *string `json:"cusip,omitempty"`
	// Number of shares owned by the ETF.
	Share *float32 `json:"share,omitempty"`
	// Portfolio's percent
	Percent *float32 `json:"percent,omitempty"`
	// Market value
	Value *float32 `json:"value,omitempty"`
}

// NewETFHoldingsData instantiates a new ETFHoldingsData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewETFHoldingsData() *ETFHoldingsData {
	this := ETFHoldingsData{}
	return &this
}

// NewETFHoldingsDataWithDefaults instantiates a new ETFHoldingsData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewETFHoldingsDataWithDefaults() *ETFHoldingsData {
	this := ETFHoldingsData{}
	return &this
}

// GetSymbol returns the Symbol field value if set, zero value otherwise.
func (o *ETFHoldingsData) GetSymbol() string {
	if o == nil || o.Symbol == nil {
		var ret string
		return ret
	}
	return *o.Symbol
}

// GetSymbolOk returns a tuple with the Symbol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ETFHoldingsData) GetSymbolOk() (*string, bool) {
	if o == nil || o.Symbol == nil {
		return nil, false
	}
	return o.Symbol, true
}

// HasSymbol returns a boolean if a field has been set.
func (o *ETFHoldingsData) HasSymbol() bool {
	if o != nil && o.Symbol != nil {
		return true
	}

	return false
}

// SetSymbol gets a reference to the given string and assigns it to the Symbol field.
func (o *ETFHoldingsData) SetSymbol(v string) {
	o.Symbol = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ETFHoldingsData) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ETFHoldingsData) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ETFHoldingsData) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ETFHoldingsData) SetName(v string) {
	o.Name = &v
}

// GetIsin returns the Isin field value if set, zero value otherwise.
func (o *ETFHoldingsData) GetIsin() string {
	if o == nil || o.Isin == nil {
		var ret string
		return ret
	}
	return *o.Isin
}

// GetIsinOk returns a tuple with the Isin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ETFHoldingsData) GetIsinOk() (*string, bool) {
	if o == nil || o.Isin == nil {
		return nil, false
	}
	return o.Isin, true
}

// HasIsin returns a boolean if a field has been set.
func (o *ETFHoldingsData) HasIsin() bool {
	if o != nil && o.Isin != nil {
		return true
	}

	return false
}

// SetIsin gets a reference to the given string and assigns it to the Isin field.
func (o *ETFHoldingsData) SetIsin(v string) {
	o.Isin = &v
}

// GetCusip returns the Cusip field value if set, zero value otherwise.
func (o *ETFHoldingsData) GetCusip() string {
	if o == nil || o.Cusip == nil {
		var ret string
		return ret
	}
	return *o.Cusip
}

// GetCusipOk returns a tuple with the Cusip field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ETFHoldingsData) GetCusipOk() (*string, bool) {
	if o == nil || o.Cusip == nil {
		return nil, false
	}
	return o.Cusip, true
}

// HasCusip returns a boolean if a field has been set.
func (o *ETFHoldingsData) HasCusip() bool {
	if o != nil && o.Cusip != nil {
		return true
	}

	return false
}

// SetCusip gets a reference to the given string and assigns it to the Cusip field.
func (o *ETFHoldingsData) SetCusip(v string) {
	o.Cusip = &v
}

// GetShare returns the Share field value if set, zero value otherwise.
func (o *ETFHoldingsData) GetShare() float32 {
	if o == nil || o.Share == nil {
		var ret float32
		return ret
	}
	return *o.Share
}

// GetShareOk returns a tuple with the Share field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ETFHoldingsData) GetShareOk() (*float32, bool) {
	if o == nil || o.Share == nil {
		return nil, false
	}
	return o.Share, true
}

// HasShare returns a boolean if a field has been set.
func (o *ETFHoldingsData) HasShare() bool {
	if o != nil && o.Share != nil {
		return true
	}

	return false
}

// SetShare gets a reference to the given float32 and assigns it to the Share field.
func (o *ETFHoldingsData) SetShare(v float32) {
	o.Share = &v
}

// GetPercent returns the Percent field value if set, zero value otherwise.
func (o *ETFHoldingsData) GetPercent() float32 {
	if o == nil || o.Percent == nil {
		var ret float32
		return ret
	}
	return *o.Percent
}

// GetPercentOk returns a tuple with the Percent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ETFHoldingsData) GetPercentOk() (*float32, bool) {
	if o == nil || o.Percent == nil {
		return nil, false
	}
	return o.Percent, true
}

// HasPercent returns a boolean if a field has been set.
func (o *ETFHoldingsData) HasPercent() bool {
	if o != nil && o.Percent != nil {
		return true
	}

	return false
}

// SetPercent gets a reference to the given float32 and assigns it to the Percent field.
func (o *ETFHoldingsData) SetPercent(v float32) {
	o.Percent = &v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *ETFHoldingsData) GetValue() float32 {
	if o == nil || o.Value == nil {
		var ret float32
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ETFHoldingsData) GetValueOk() (*float32, bool) {
	if o == nil || o.Value == nil {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *ETFHoldingsData) HasValue() bool {
	if o != nil && o.Value != nil {
		return true
	}

	return false
}

// SetValue gets a reference to the given float32 and assigns it to the Value field.
func (o *ETFHoldingsData) SetValue(v float32) {
	o.Value = &v
}

func (o ETFHoldingsData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Symbol != nil {
		toSerialize["symbol"] = o.Symbol
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Isin != nil {
		toSerialize["isin"] = o.Isin
	}
	if o.Cusip != nil {
		toSerialize["cusip"] = o.Cusip
	}
	if o.Share != nil {
		toSerialize["share"] = o.Share
	}
	if o.Percent != nil {
		toSerialize["percent"] = o.Percent
	}
	if o.Value != nil {
		toSerialize["value"] = o.Value
	}
	return json.Marshal(toSerialize)
}

type NullableETFHoldingsData struct {
	value *ETFHoldingsData
	isSet bool
}

func (v NullableETFHoldingsData) Get() *ETFHoldingsData {
	return v.value
}

func (v *NullableETFHoldingsData) Set(val *ETFHoldingsData) {
	v.value = val
	v.isSet = true
}

func (v NullableETFHoldingsData) IsSet() bool {
	return v.isSet
}

func (v *NullableETFHoldingsData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableETFHoldingsData(val *ETFHoldingsData) *NullableETFHoldingsData {
	return &NullableETFHoldingsData{value: val, isSet: true}
}

func (v NullableETFHoldingsData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableETFHoldingsData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


