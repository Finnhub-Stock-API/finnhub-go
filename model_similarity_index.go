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

// SimilarityIndex struct for SimilarityIndex
type SimilarityIndex struct {
	// Symbol.
	Symbol *string `json:"symbol,omitempty"`
	// CIK.
	Cik *string `json:"cik,omitempty"`
	// Array of filings with its cosine similarity compared to the same report of the previous year.
	Similarity *[]map[string]interface{} `json:"similarity,omitempty"`
}

// NewSimilarityIndex instantiates a new SimilarityIndex object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSimilarityIndex() *SimilarityIndex {
	this := SimilarityIndex{}
	return &this
}

// NewSimilarityIndexWithDefaults instantiates a new SimilarityIndex object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSimilarityIndexWithDefaults() *SimilarityIndex {
	this := SimilarityIndex{}
	return &this
}

// GetSymbol returns the Symbol field value if set, zero value otherwise.
func (o *SimilarityIndex) GetSymbol() string {
	if o == nil || o.Symbol == nil {
		var ret string
		return ret
	}
	return *o.Symbol
}

// GetSymbolOk returns a tuple with the Symbol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SimilarityIndex) GetSymbolOk() (*string, bool) {
	if o == nil || o.Symbol == nil {
		return nil, false
	}
	return o.Symbol, true
}

// HasSymbol returns a boolean if a field has been set.
func (o *SimilarityIndex) HasSymbol() bool {
	if o != nil && o.Symbol != nil {
		return true
	}

	return false
}

// SetSymbol gets a reference to the given string and assigns it to the Symbol field.
func (o *SimilarityIndex) SetSymbol(v string) {
	o.Symbol = &v
}

// GetCik returns the Cik field value if set, zero value otherwise.
func (o *SimilarityIndex) GetCik() string {
	if o == nil || o.Cik == nil {
		var ret string
		return ret
	}
	return *o.Cik
}

// GetCikOk returns a tuple with the Cik field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SimilarityIndex) GetCikOk() (*string, bool) {
	if o == nil || o.Cik == nil {
		return nil, false
	}
	return o.Cik, true
}

// HasCik returns a boolean if a field has been set.
func (o *SimilarityIndex) HasCik() bool {
	if o != nil && o.Cik != nil {
		return true
	}

	return false
}

// SetCik gets a reference to the given string and assigns it to the Cik field.
func (o *SimilarityIndex) SetCik(v string) {
	o.Cik = &v
}

// GetSimilarity returns the Similarity field value if set, zero value otherwise.
func (o *SimilarityIndex) GetSimilarity() []map[string]interface{} {
	if o == nil || o.Similarity == nil {
		var ret []map[string]interface{}
		return ret
	}
	return *o.Similarity
}

// GetSimilarityOk returns a tuple with the Similarity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SimilarityIndex) GetSimilarityOk() (*[]map[string]interface{}, bool) {
	if o == nil || o.Similarity == nil {
		return nil, false
	}
	return o.Similarity, true
}

// HasSimilarity returns a boolean if a field has been set.
func (o *SimilarityIndex) HasSimilarity() bool {
	if o != nil && o.Similarity != nil {
		return true
	}

	return false
}

// SetSimilarity gets a reference to the given []map[string]interface{} and assigns it to the Similarity field.
func (o *SimilarityIndex) SetSimilarity(v []map[string]interface{}) {
	o.Similarity = &v
}

func (o SimilarityIndex) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Symbol != nil {
		toSerialize["symbol"] = o.Symbol
	}
	if o.Cik != nil {
		toSerialize["cik"] = o.Cik
	}
	if o.Similarity != nil {
		toSerialize["similarity"] = o.Similarity
	}
	return json.Marshal(toSerialize)
}

type NullableSimilarityIndex struct {
	value *SimilarityIndex
	isSet bool
}

func (v NullableSimilarityIndex) Get() *SimilarityIndex {
	return v.value
}

func (v *NullableSimilarityIndex) Set(val *SimilarityIndex) {
	v.value = val
	v.isSet = true
}

func (v NullableSimilarityIndex) IsSet() bool {
	return v.isSet
}

func (v *NullableSimilarityIndex) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSimilarityIndex(val *SimilarityIndex) *NullableSimilarityIndex {
	return &NullableSimilarityIndex{value: val, isSet: true}
}

func (v NullableSimilarityIndex) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSimilarityIndex) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


