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

// IndexHistoricalConstituent struct for IndexHistoricalConstituent
type IndexHistoricalConstituent struct {
	// Symbol
	Symbol *string `json:"symbol,omitempty"`
	// <code>add</code> or <code>remove</code>.
	Action *string `json:"action,omitempty"`
	// Date of joining or leaving the index.
	Date *string `json:"date,omitempty"`
}

// NewIndexHistoricalConstituent instantiates a new IndexHistoricalConstituent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIndexHistoricalConstituent() *IndexHistoricalConstituent {
	this := IndexHistoricalConstituent{}
	return &this
}

// NewIndexHistoricalConstituentWithDefaults instantiates a new IndexHistoricalConstituent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIndexHistoricalConstituentWithDefaults() *IndexHistoricalConstituent {
	this := IndexHistoricalConstituent{}
	return &this
}

// GetSymbol returns the Symbol field value if set, zero value otherwise.
func (o *IndexHistoricalConstituent) GetSymbol() string {
	if o == nil || o.Symbol == nil {
		var ret string
		return ret
	}
	return *o.Symbol
}

// GetSymbolOk returns a tuple with the Symbol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IndexHistoricalConstituent) GetSymbolOk() (*string, bool) {
	if o == nil || o.Symbol == nil {
		return nil, false
	}
	return o.Symbol, true
}

// HasSymbol returns a boolean if a field has been set.
func (o *IndexHistoricalConstituent) HasSymbol() bool {
	if o != nil && o.Symbol != nil {
		return true
	}

	return false
}

// SetSymbol gets a reference to the given string and assigns it to the Symbol field.
func (o *IndexHistoricalConstituent) SetSymbol(v string) {
	o.Symbol = &v
}

// GetAction returns the Action field value if set, zero value otherwise.
func (o *IndexHistoricalConstituent) GetAction() string {
	if o == nil || o.Action == nil {
		var ret string
		return ret
	}
	return *o.Action
}

// GetActionOk returns a tuple with the Action field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IndexHistoricalConstituent) GetActionOk() (*string, bool) {
	if o == nil || o.Action == nil {
		return nil, false
	}
	return o.Action, true
}

// HasAction returns a boolean if a field has been set.
func (o *IndexHistoricalConstituent) HasAction() bool {
	if o != nil && o.Action != nil {
		return true
	}

	return false
}

// SetAction gets a reference to the given string and assigns it to the Action field.
func (o *IndexHistoricalConstituent) SetAction(v string) {
	o.Action = &v
}

// GetDate returns the Date field value if set, zero value otherwise.
func (o *IndexHistoricalConstituent) GetDate() string {
	if o == nil || o.Date == nil {
		var ret string
		return ret
	}
	return *o.Date
}

// GetDateOk returns a tuple with the Date field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IndexHistoricalConstituent) GetDateOk() (*string, bool) {
	if o == nil || o.Date == nil {
		return nil, false
	}
	return o.Date, true
}

// HasDate returns a boolean if a field has been set.
func (o *IndexHistoricalConstituent) HasDate() bool {
	if o != nil && o.Date != nil {
		return true
	}

	return false
}

// SetDate gets a reference to the given string and assigns it to the Date field.
func (o *IndexHistoricalConstituent) SetDate(v string) {
	o.Date = &v
}

func (o IndexHistoricalConstituent) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Symbol != nil {
		toSerialize["symbol"] = o.Symbol
	}
	if o.Action != nil {
		toSerialize["action"] = o.Action
	}
	if o.Date != nil {
		toSerialize["date"] = o.Date
	}
	return json.Marshal(toSerialize)
}

type NullableIndexHistoricalConstituent struct {
	value *IndexHistoricalConstituent
	isSet bool
}

func (v NullableIndexHistoricalConstituent) Get() *IndexHistoricalConstituent {
	return v.value
}

func (v *NullableIndexHistoricalConstituent) Set(val *IndexHistoricalConstituent) {
	v.value = val
	v.isSet = true
}

func (v NullableIndexHistoricalConstituent) IsSet() bool {
	return v.isSet
}

func (v *NullableIndexHistoricalConstituent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIndexHistoricalConstituent(val *IndexHistoricalConstituent) *NullableIndexHistoricalConstituent {
	return &NullableIndexHistoricalConstituent{value: val, isSet: true}
}

func (v NullableIndexHistoricalConstituent) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIndexHistoricalConstituent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


