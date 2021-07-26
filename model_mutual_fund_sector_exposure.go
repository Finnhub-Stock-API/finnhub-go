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

// MutualFundSectorExposure struct for MutualFundSectorExposure
type MutualFundSectorExposure struct {
	// Mutual symbol.
	Symbol *string `json:"symbol,omitempty"`
	// Array of sector and exposure levels.
	SectorExposure *[]MutualFundSectorExposureData `json:"sectorExposure,omitempty"`
}

// NewMutualFundSectorExposure instantiates a new MutualFundSectorExposure object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMutualFundSectorExposure() *MutualFundSectorExposure {
	this := MutualFundSectorExposure{}
	return &this
}

// NewMutualFundSectorExposureWithDefaults instantiates a new MutualFundSectorExposure object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMutualFundSectorExposureWithDefaults() *MutualFundSectorExposure {
	this := MutualFundSectorExposure{}
	return &this
}

// GetSymbol returns the Symbol field value if set, zero value otherwise.
func (o *MutualFundSectorExposure) GetSymbol() string {
	if o == nil || o.Symbol == nil {
		var ret string
		return ret
	}
	return *o.Symbol
}

// GetSymbolOk returns a tuple with the Symbol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MutualFundSectorExposure) GetSymbolOk() (*string, bool) {
	if o == nil || o.Symbol == nil {
		return nil, false
	}
	return o.Symbol, true
}

// HasSymbol returns a boolean if a field has been set.
func (o *MutualFundSectorExposure) HasSymbol() bool {
	if o != nil && o.Symbol != nil {
		return true
	}

	return false
}

// SetSymbol gets a reference to the given string and assigns it to the Symbol field.
func (o *MutualFundSectorExposure) SetSymbol(v string) {
	o.Symbol = &v
}

// GetSectorExposure returns the SectorExposure field value if set, zero value otherwise.
func (o *MutualFundSectorExposure) GetSectorExposure() []MutualFundSectorExposureData {
	if o == nil || o.SectorExposure == nil {
		var ret []MutualFundSectorExposureData
		return ret
	}
	return *o.SectorExposure
}

// GetSectorExposureOk returns a tuple with the SectorExposure field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MutualFundSectorExposure) GetSectorExposureOk() (*[]MutualFundSectorExposureData, bool) {
	if o == nil || o.SectorExposure == nil {
		return nil, false
	}
	return o.SectorExposure, true
}

// HasSectorExposure returns a boolean if a field has been set.
func (o *MutualFundSectorExposure) HasSectorExposure() bool {
	if o != nil && o.SectorExposure != nil {
		return true
	}

	return false
}

// SetSectorExposure gets a reference to the given []MutualFundSectorExposureData and assigns it to the SectorExposure field.
func (o *MutualFundSectorExposure) SetSectorExposure(v []MutualFundSectorExposureData) {
	o.SectorExposure = &v
}

func (o MutualFundSectorExposure) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Symbol != nil {
		toSerialize["symbol"] = o.Symbol
	}
	if o.SectorExposure != nil {
		toSerialize["sectorExposure"] = o.SectorExposure
	}
	return json.Marshal(toSerialize)
}

type NullableMutualFundSectorExposure struct {
	value *MutualFundSectorExposure
	isSet bool
}

func (v NullableMutualFundSectorExposure) Get() *MutualFundSectorExposure {
	return v.value
}

func (v *NullableMutualFundSectorExposure) Set(val *MutualFundSectorExposure) {
	v.value = val
	v.isSet = true
}

func (v NullableMutualFundSectorExposure) IsSet() bool {
	return v.isSet
}

func (v *NullableMutualFundSectorExposure) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMutualFundSectorExposure(val *MutualFundSectorExposure) *NullableMutualFundSectorExposure {
	return &NullableMutualFundSectorExposure{value: val, isSet: true}
}

func (v NullableMutualFundSectorExposure) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMutualFundSectorExposure) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


