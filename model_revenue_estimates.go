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

// RevenueEstimates struct for RevenueEstimates
type RevenueEstimates struct {
	// List of estimates
	Data *[]RevenueEstimatesInfo `json:"data,omitempty"`
	// Frequency: annual or quarterly.
	Freq *string `json:"freq,omitempty"`
	// Company symbol.
	Symbol *string `json:"symbol,omitempty"`
}

// NewRevenueEstimates instantiates a new RevenueEstimates object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRevenueEstimates() *RevenueEstimates {
	this := RevenueEstimates{}
	return &this
}

// NewRevenueEstimatesWithDefaults instantiates a new RevenueEstimates object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRevenueEstimatesWithDefaults() *RevenueEstimates {
	this := RevenueEstimates{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *RevenueEstimates) GetData() []RevenueEstimatesInfo {
	if o == nil || o.Data == nil {
		var ret []RevenueEstimatesInfo
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RevenueEstimates) GetDataOk() (*[]RevenueEstimatesInfo, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *RevenueEstimates) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given []RevenueEstimatesInfo and assigns it to the Data field.
func (o *RevenueEstimates) SetData(v []RevenueEstimatesInfo) {
	o.Data = &v
}

// GetFreq returns the Freq field value if set, zero value otherwise.
func (o *RevenueEstimates) GetFreq() string {
	if o == nil || o.Freq == nil {
		var ret string
		return ret
	}
	return *o.Freq
}

// GetFreqOk returns a tuple with the Freq field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RevenueEstimates) GetFreqOk() (*string, bool) {
	if o == nil || o.Freq == nil {
		return nil, false
	}
	return o.Freq, true
}

// HasFreq returns a boolean if a field has been set.
func (o *RevenueEstimates) HasFreq() bool {
	if o != nil && o.Freq != nil {
		return true
	}

	return false
}

// SetFreq gets a reference to the given string and assigns it to the Freq field.
func (o *RevenueEstimates) SetFreq(v string) {
	o.Freq = &v
}

// GetSymbol returns the Symbol field value if set, zero value otherwise.
func (o *RevenueEstimates) GetSymbol() string {
	if o == nil || o.Symbol == nil {
		var ret string
		return ret
	}
	return *o.Symbol
}

// GetSymbolOk returns a tuple with the Symbol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RevenueEstimates) GetSymbolOk() (*string, bool) {
	if o == nil || o.Symbol == nil {
		return nil, false
	}
	return o.Symbol, true
}

// HasSymbol returns a boolean if a field has been set.
func (o *RevenueEstimates) HasSymbol() bool {
	if o != nil && o.Symbol != nil {
		return true
	}

	return false
}

// SetSymbol gets a reference to the given string and assigns it to the Symbol field.
func (o *RevenueEstimates) SetSymbol(v string) {
	o.Symbol = &v
}

func (o RevenueEstimates) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	if o.Freq != nil {
		toSerialize["freq"] = o.Freq
	}
	if o.Symbol != nil {
		toSerialize["symbol"] = o.Symbol
	}
	return json.Marshal(toSerialize)
}

type NullableRevenueEstimates struct {
	value *RevenueEstimates
	isSet bool
}

func (v NullableRevenueEstimates) Get() *RevenueEstimates {
	return v.value
}

func (v *NullableRevenueEstimates) Set(val *RevenueEstimates) {
	v.value = val
	v.isSet = true
}

func (v NullableRevenueEstimates) IsSet() bool {
	return v.isSet
}

func (v *NullableRevenueEstimates) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRevenueEstimates(val *RevenueEstimates) *NullableRevenueEstimates {
	return &NullableRevenueEstimates{value: val, isSet: true}
}

func (v NullableRevenueEstimates) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRevenueEstimates) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


