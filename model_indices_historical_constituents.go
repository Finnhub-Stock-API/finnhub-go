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

// IndicesHistoricalConstituents struct for IndicesHistoricalConstituents
type IndicesHistoricalConstituents struct {
	// Index's symbol.
	Symbol *string `json:"symbol,omitempty"`
	// Array of historical constituents.
	HistoricalConstituents *[]IndexHistoricalConstituent `json:"historicalConstituents,omitempty"`
}

// NewIndicesHistoricalConstituents instantiates a new IndicesHistoricalConstituents object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIndicesHistoricalConstituents() *IndicesHistoricalConstituents {
	this := IndicesHistoricalConstituents{}
	return &this
}

// NewIndicesHistoricalConstituentsWithDefaults instantiates a new IndicesHistoricalConstituents object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIndicesHistoricalConstituentsWithDefaults() *IndicesHistoricalConstituents {
	this := IndicesHistoricalConstituents{}
	return &this
}

// GetSymbol returns the Symbol field value if set, zero value otherwise.
func (o *IndicesHistoricalConstituents) GetSymbol() string {
	if o == nil || o.Symbol == nil {
		var ret string
		return ret
	}
	return *o.Symbol
}

// GetSymbolOk returns a tuple with the Symbol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IndicesHistoricalConstituents) GetSymbolOk() (*string, bool) {
	if o == nil || o.Symbol == nil {
		return nil, false
	}
	return o.Symbol, true
}

// HasSymbol returns a boolean if a field has been set.
func (o *IndicesHistoricalConstituents) HasSymbol() bool {
	if o != nil && o.Symbol != nil {
		return true
	}

	return false
}

// SetSymbol gets a reference to the given string and assigns it to the Symbol field.
func (o *IndicesHistoricalConstituents) SetSymbol(v string) {
	o.Symbol = &v
}

// GetHistoricalConstituents returns the HistoricalConstituents field value if set, zero value otherwise.
func (o *IndicesHistoricalConstituents) GetHistoricalConstituents() []IndexHistoricalConstituent {
	if o == nil || o.HistoricalConstituents == nil {
		var ret []IndexHistoricalConstituent
		return ret
	}
	return *o.HistoricalConstituents
}

// GetHistoricalConstituentsOk returns a tuple with the HistoricalConstituents field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IndicesHistoricalConstituents) GetHistoricalConstituentsOk() (*[]IndexHistoricalConstituent, bool) {
	if o == nil || o.HistoricalConstituents == nil {
		return nil, false
	}
	return o.HistoricalConstituents, true
}

// HasHistoricalConstituents returns a boolean if a field has been set.
func (o *IndicesHistoricalConstituents) HasHistoricalConstituents() bool {
	if o != nil && o.HistoricalConstituents != nil {
		return true
	}

	return false
}

// SetHistoricalConstituents gets a reference to the given []IndexHistoricalConstituent and assigns it to the HistoricalConstituents field.
func (o *IndicesHistoricalConstituents) SetHistoricalConstituents(v []IndexHistoricalConstituent) {
	o.HistoricalConstituents = &v
}

func (o IndicesHistoricalConstituents) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Symbol != nil {
		toSerialize["symbol"] = o.Symbol
	}
	if o.HistoricalConstituents != nil {
		toSerialize["historicalConstituents"] = o.HistoricalConstituents
	}
	return json.Marshal(toSerialize)
}

type NullableIndicesHistoricalConstituents struct {
	value *IndicesHistoricalConstituents
	isSet bool
}

func (v NullableIndicesHistoricalConstituents) Get() *IndicesHistoricalConstituents {
	return v.value
}

func (v *NullableIndicesHistoricalConstituents) Set(val *IndicesHistoricalConstituents) {
	v.value = val
	v.isSet = true
}

func (v NullableIndicesHistoricalConstituents) IsSet() bool {
	return v.isSet
}

func (v *NullableIndicesHistoricalConstituents) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIndicesHistoricalConstituents(val *IndicesHistoricalConstituents) *NullableIndicesHistoricalConstituents {
	return &NullableIndicesHistoricalConstituents{value: val, isSet: true}
}

func (v NullableIndicesHistoricalConstituents) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIndicesHistoricalConstituents) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


