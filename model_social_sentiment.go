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

// SocialSentiment struct for SocialSentiment
type SocialSentiment struct {
	// Company symbol.
	Symbol *string `json:"symbol,omitempty"`
	// Sentiment data.
	Data *[]SentimentContent `json:"data,omitempty"`
}

// NewSocialSentiment instantiates a new SocialSentiment object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSocialSentiment() *SocialSentiment {
	this := SocialSentiment{}
	return &this
}

// NewSocialSentimentWithDefaults instantiates a new SocialSentiment object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSocialSentimentWithDefaults() *SocialSentiment {
	this := SocialSentiment{}
	return &this
}

// GetSymbol returns the Symbol field value if set, zero value otherwise.
func (o *SocialSentiment) GetSymbol() string {
	if o == nil || o.Symbol == nil {
		var ret string
		return ret
	}
	return *o.Symbol
}

// GetSymbolOk returns a tuple with the Symbol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SocialSentiment) GetSymbolOk() (*string, bool) {
	if o == nil || o.Symbol == nil {
		return nil, false
	}
	return o.Symbol, true
}

// HasSymbol returns a boolean if a field has been set.
func (o *SocialSentiment) HasSymbol() bool {
	if o != nil && o.Symbol != nil {
		return true
	}

	return false
}

// SetSymbol gets a reference to the given string and assigns it to the Symbol field.
func (o *SocialSentiment) SetSymbol(v string) {
	o.Symbol = &v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *SocialSentiment) GetData() []SentimentContent {
	if o == nil || o.Data == nil {
		var ret []SentimentContent
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SocialSentiment) GetDataOk() (*[]SentimentContent, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *SocialSentiment) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given []SentimentContent and assigns it to the Data field.
func (o *SocialSentiment) SetData(v []SentimentContent) {
	o.Data = &v
}

func (o SocialSentiment) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Symbol != nil {
		toSerialize["symbol"] = o.Symbol
	}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullableSocialSentiment struct {
	value *SocialSentiment
	isSet bool
}

func (v NullableSocialSentiment) Get() *SocialSentiment {
	return v.value
}

func (v *NullableSocialSentiment) Set(val *SocialSentiment) {
	v.value = val
	v.isSet = true
}

func (v NullableSocialSentiment) IsSet() bool {
	return v.isSet
}

func (v *NullableSocialSentiment) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSocialSentiment(val *SocialSentiment) *NullableSocialSentiment {
	return &NullableSocialSentiment{value: val, isSet: true}
}

func (v NullableSocialSentiment) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSocialSentiment) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


