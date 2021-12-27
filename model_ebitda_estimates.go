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

// EbitdaEstimates struct for EbitdaEstimates
type EbitdaEstimates struct {
	// List of estimates
	Data *[]EbitdaEstimatesInfo `json:"data,omitempty"`
	// Frequency: annual or quarterly.
	Freq *string `json:"freq,omitempty"`
	// Company symbol.
	Symbol *string `json:"symbol,omitempty"`
}

// NewEbitdaEstimates instantiates a new EbitdaEstimates object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEbitdaEstimates() *EbitdaEstimates {
	this := EbitdaEstimates{}
	return &this
}

// NewEbitdaEstimatesWithDefaults instantiates a new EbitdaEstimates object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEbitdaEstimatesWithDefaults() *EbitdaEstimates {
	this := EbitdaEstimates{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *EbitdaEstimates) GetData() []EbitdaEstimatesInfo {
	if o == nil || o.Data == nil {
		var ret []EbitdaEstimatesInfo
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EbitdaEstimates) GetDataOk() (*[]EbitdaEstimatesInfo, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *EbitdaEstimates) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given []EbitdaEstimatesInfo and assigns it to the Data field.
func (o *EbitdaEstimates) SetData(v []EbitdaEstimatesInfo) {
	o.Data = &v
}

// GetFreq returns the Freq field value if set, zero value otherwise.
func (o *EbitdaEstimates) GetFreq() string {
	if o == nil || o.Freq == nil {
		var ret string
		return ret
	}
	return *o.Freq
}

// GetFreqOk returns a tuple with the Freq field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EbitdaEstimates) GetFreqOk() (*string, bool) {
	if o == nil || o.Freq == nil {
		return nil, false
	}
	return o.Freq, true
}

// HasFreq returns a boolean if a field has been set.
func (o *EbitdaEstimates) HasFreq() bool {
	if o != nil && o.Freq != nil {
		return true
	}

	return false
}

// SetFreq gets a reference to the given string and assigns it to the Freq field.
func (o *EbitdaEstimates) SetFreq(v string) {
	o.Freq = &v
}

// GetSymbol returns the Symbol field value if set, zero value otherwise.
func (o *EbitdaEstimates) GetSymbol() string {
	if o == nil || o.Symbol == nil {
		var ret string
		return ret
	}
	return *o.Symbol
}

// GetSymbolOk returns a tuple with the Symbol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EbitdaEstimates) GetSymbolOk() (*string, bool) {
	if o == nil || o.Symbol == nil {
		return nil, false
	}
	return o.Symbol, true
}

// HasSymbol returns a boolean if a field has been set.
func (o *EbitdaEstimates) HasSymbol() bool {
	if o != nil && o.Symbol != nil {
		return true
	}

	return false
}

// SetSymbol gets a reference to the given string and assigns it to the Symbol field.
func (o *EbitdaEstimates) SetSymbol(v string) {
	o.Symbol = &v
}

func (o EbitdaEstimates) MarshalJSON() ([]byte, error) {
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

type NullableEbitdaEstimates struct {
	value *EbitdaEstimates
	isSet bool
}

func (v NullableEbitdaEstimates) Get() *EbitdaEstimates {
	return v.value
}

func (v *NullableEbitdaEstimates) Set(val *EbitdaEstimates) {
	v.value = val
	v.isSet = true
}

func (v NullableEbitdaEstimates) IsSet() bool {
	return v.isSet
}

func (v *NullableEbitdaEstimates) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEbitdaEstimates(val *EbitdaEstimates) *NullableEbitdaEstimates {
	return &NullableEbitdaEstimates{value: val, isSet: true}
}

func (v NullableEbitdaEstimates) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEbitdaEstimates) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

