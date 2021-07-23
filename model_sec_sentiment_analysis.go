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

// SECSentimentAnalysis struct for SECSentimentAnalysis
type SECSentimentAnalysis struct {
	// Access number.
	AccessNumber *string `json:"accessNumber,omitempty"`
	// Symbol.
	Symbol *string `json:"symbol,omitempty"`
	// CIK.
	Cik *string `json:"cik,omitempty"`
	Sentiment *FilingSentiment `json:"sentiment,omitempty"`
}

// NewSECSentimentAnalysis instantiates a new SECSentimentAnalysis object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSECSentimentAnalysis() *SECSentimentAnalysis {
	this := SECSentimentAnalysis{}
	return &this
}

// NewSECSentimentAnalysisWithDefaults instantiates a new SECSentimentAnalysis object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSECSentimentAnalysisWithDefaults() *SECSentimentAnalysis {
	this := SECSentimentAnalysis{}
	return &this
}

// GetAccessNumber returns the AccessNumber field value if set, zero value otherwise.
func (o *SECSentimentAnalysis) GetAccessNumber() string {
	if o == nil || o.AccessNumber == nil {
		var ret string
		return ret
	}
	return *o.AccessNumber
}

// GetAccessNumberOk returns a tuple with the AccessNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SECSentimentAnalysis) GetAccessNumberOk() (*string, bool) {
	if o == nil || o.AccessNumber == nil {
		return nil, false
	}
	return o.AccessNumber, true
}

// HasAccessNumber returns a boolean if a field has been set.
func (o *SECSentimentAnalysis) HasAccessNumber() bool {
	if o != nil && o.AccessNumber != nil {
		return true
	}

	return false
}

// SetAccessNumber gets a reference to the given string and assigns it to the AccessNumber field.
func (o *SECSentimentAnalysis) SetAccessNumber(v string) {
	o.AccessNumber = &v
}

// GetSymbol returns the Symbol field value if set, zero value otherwise.
func (o *SECSentimentAnalysis) GetSymbol() string {
	if o == nil || o.Symbol == nil {
		var ret string
		return ret
	}
	return *o.Symbol
}

// GetSymbolOk returns a tuple with the Symbol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SECSentimentAnalysis) GetSymbolOk() (*string, bool) {
	if o == nil || o.Symbol == nil {
		return nil, false
	}
	return o.Symbol, true
}

// HasSymbol returns a boolean if a field has been set.
func (o *SECSentimentAnalysis) HasSymbol() bool {
	if o != nil && o.Symbol != nil {
		return true
	}

	return false
}

// SetSymbol gets a reference to the given string and assigns it to the Symbol field.
func (o *SECSentimentAnalysis) SetSymbol(v string) {
	o.Symbol = &v
}

// GetCik returns the Cik field value if set, zero value otherwise.
func (o *SECSentimentAnalysis) GetCik() string {
	if o == nil || o.Cik == nil {
		var ret string
		return ret
	}
	return *o.Cik
}

// GetCikOk returns a tuple with the Cik field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SECSentimentAnalysis) GetCikOk() (*string, bool) {
	if o == nil || o.Cik == nil {
		return nil, false
	}
	return o.Cik, true
}

// HasCik returns a boolean if a field has been set.
func (o *SECSentimentAnalysis) HasCik() bool {
	if o != nil && o.Cik != nil {
		return true
	}

	return false
}

// SetCik gets a reference to the given string and assigns it to the Cik field.
func (o *SECSentimentAnalysis) SetCik(v string) {
	o.Cik = &v
}

// GetSentiment returns the Sentiment field value if set, zero value otherwise.
func (o *SECSentimentAnalysis) GetSentiment() FilingSentiment {
	if o == nil || o.Sentiment == nil {
		var ret FilingSentiment
		return ret
	}
	return *o.Sentiment
}

// GetSentimentOk returns a tuple with the Sentiment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SECSentimentAnalysis) GetSentimentOk() (*FilingSentiment, bool) {
	if o == nil || o.Sentiment == nil {
		return nil, false
	}
	return o.Sentiment, true
}

// HasSentiment returns a boolean if a field has been set.
func (o *SECSentimentAnalysis) HasSentiment() bool {
	if o != nil && o.Sentiment != nil {
		return true
	}

	return false
}

// SetSentiment gets a reference to the given FilingSentiment and assigns it to the Sentiment field.
func (o *SECSentimentAnalysis) SetSentiment(v FilingSentiment) {
	o.Sentiment = &v
}

func (o SECSentimentAnalysis) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AccessNumber != nil {
		toSerialize["accessNumber"] = o.AccessNumber
	}
	if o.Symbol != nil {
		toSerialize["symbol"] = o.Symbol
	}
	if o.Cik != nil {
		toSerialize["cik"] = o.Cik
	}
	if o.Sentiment != nil {
		toSerialize["sentiment"] = o.Sentiment
	}
	return json.Marshal(toSerialize)
}

type NullableSECSentimentAnalysis struct {
	value *SECSentimentAnalysis
	isSet bool
}

func (v NullableSECSentimentAnalysis) Get() *SECSentimentAnalysis {
	return v.value
}

func (v *NullableSECSentimentAnalysis) Set(val *SECSentimentAnalysis) {
	v.value = val
	v.isSet = true
}

func (v NullableSECSentimentAnalysis) IsSet() bool {
	return v.isSet
}

func (v *NullableSECSentimentAnalysis) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSECSentimentAnalysis(val *SECSentimentAnalysis) *NullableSECSentimentAnalysis {
	return &NullableSECSentimentAnalysis{value: val, isSet: true}
}

func (v NullableSECSentimentAnalysis) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSECSentimentAnalysis) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


