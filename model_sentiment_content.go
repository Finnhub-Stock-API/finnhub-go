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

// SentimentContent struct for SentimentContent
type SentimentContent struct {
	// Number of mentions
	Mention *int64 `json:"mention,omitempty"`
	// Number of positive mentions
	PositiveMention *int64 `json:"positiveMention,omitempty"`
	// Number of negative mentions
	NegativeMention *int64 `json:"negativeMention,omitempty"`
	// Positive score. Range 0-1
	PositiveScore *float32 `json:"positiveScore,omitempty"`
	// Negative score. Range 0-1
	NegativeScore *float32 `json:"negativeScore,omitempty"`
	// Final score. Range: -1 to 1 with 1 is very positive and -1 is very negative
	Score *float32 `json:"score,omitempty"`
	// Period.
	AtTime *string `json:"atTime,omitempty"`
}

// NewSentimentContent instantiates a new SentimentContent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSentimentContent() *SentimentContent {
	this := SentimentContent{}
	return &this
}

// NewSentimentContentWithDefaults instantiates a new SentimentContent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSentimentContentWithDefaults() *SentimentContent {
	this := SentimentContent{}
	return &this
}

// GetMention returns the Mention field value if set, zero value otherwise.
func (o *SentimentContent) GetMention() int64 {
	if o == nil || o.Mention == nil {
		var ret int64
		return ret
	}
	return *o.Mention
}

// GetMentionOk returns a tuple with the Mention field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SentimentContent) GetMentionOk() (*int64, bool) {
	if o == nil || o.Mention == nil {
		return nil, false
	}
	return o.Mention, true
}

// HasMention returns a boolean if a field has been set.
func (o *SentimentContent) HasMention() bool {
	if o != nil && o.Mention != nil {
		return true
	}

	return false
}

// SetMention gets a reference to the given int64 and assigns it to the Mention field.
func (o *SentimentContent) SetMention(v int64) {
	o.Mention = &v
}

// GetPositiveMention returns the PositiveMention field value if set, zero value otherwise.
func (o *SentimentContent) GetPositiveMention() int64 {
	if o == nil || o.PositiveMention == nil {
		var ret int64
		return ret
	}
	return *o.PositiveMention
}

// GetPositiveMentionOk returns a tuple with the PositiveMention field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SentimentContent) GetPositiveMentionOk() (*int64, bool) {
	if o == nil || o.PositiveMention == nil {
		return nil, false
	}
	return o.PositiveMention, true
}

// HasPositiveMention returns a boolean if a field has been set.
func (o *SentimentContent) HasPositiveMention() bool {
	if o != nil && o.PositiveMention != nil {
		return true
	}

	return false
}

// SetPositiveMention gets a reference to the given int64 and assigns it to the PositiveMention field.
func (o *SentimentContent) SetPositiveMention(v int64) {
	o.PositiveMention = &v
}

// GetNegativeMention returns the NegativeMention field value if set, zero value otherwise.
func (o *SentimentContent) GetNegativeMention() int64 {
	if o == nil || o.NegativeMention == nil {
		var ret int64
		return ret
	}
	return *o.NegativeMention
}

// GetNegativeMentionOk returns a tuple with the NegativeMention field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SentimentContent) GetNegativeMentionOk() (*int64, bool) {
	if o == nil || o.NegativeMention == nil {
		return nil, false
	}
	return o.NegativeMention, true
}

// HasNegativeMention returns a boolean if a field has been set.
func (o *SentimentContent) HasNegativeMention() bool {
	if o != nil && o.NegativeMention != nil {
		return true
	}

	return false
}

// SetNegativeMention gets a reference to the given int64 and assigns it to the NegativeMention field.
func (o *SentimentContent) SetNegativeMention(v int64) {
	o.NegativeMention = &v
}

// GetPositiveScore returns the PositiveScore field value if set, zero value otherwise.
func (o *SentimentContent) GetPositiveScore() float32 {
	if o == nil || o.PositiveScore == nil {
		var ret float32
		return ret
	}
	return *o.PositiveScore
}

// GetPositiveScoreOk returns a tuple with the PositiveScore field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SentimentContent) GetPositiveScoreOk() (*float32, bool) {
	if o == nil || o.PositiveScore == nil {
		return nil, false
	}
	return o.PositiveScore, true
}

// HasPositiveScore returns a boolean if a field has been set.
func (o *SentimentContent) HasPositiveScore() bool {
	if o != nil && o.PositiveScore != nil {
		return true
	}

	return false
}

// SetPositiveScore gets a reference to the given float32 and assigns it to the PositiveScore field.
func (o *SentimentContent) SetPositiveScore(v float32) {
	o.PositiveScore = &v
}

// GetNegativeScore returns the NegativeScore field value if set, zero value otherwise.
func (o *SentimentContent) GetNegativeScore() float32 {
	if o == nil || o.NegativeScore == nil {
		var ret float32
		return ret
	}
	return *o.NegativeScore
}

// GetNegativeScoreOk returns a tuple with the NegativeScore field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SentimentContent) GetNegativeScoreOk() (*float32, bool) {
	if o == nil || o.NegativeScore == nil {
		return nil, false
	}
	return o.NegativeScore, true
}

// HasNegativeScore returns a boolean if a field has been set.
func (o *SentimentContent) HasNegativeScore() bool {
	if o != nil && o.NegativeScore != nil {
		return true
	}

	return false
}

// SetNegativeScore gets a reference to the given float32 and assigns it to the NegativeScore field.
func (o *SentimentContent) SetNegativeScore(v float32) {
	o.NegativeScore = &v
}

// GetScore returns the Score field value if set, zero value otherwise.
func (o *SentimentContent) GetScore() float32 {
	if o == nil || o.Score == nil {
		var ret float32
		return ret
	}
	return *o.Score
}

// GetScoreOk returns a tuple with the Score field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SentimentContent) GetScoreOk() (*float32, bool) {
	if o == nil || o.Score == nil {
		return nil, false
	}
	return o.Score, true
}

// HasScore returns a boolean if a field has been set.
func (o *SentimentContent) HasScore() bool {
	if o != nil && o.Score != nil {
		return true
	}

	return false
}

// SetScore gets a reference to the given float32 and assigns it to the Score field.
func (o *SentimentContent) SetScore(v float32) {
	o.Score = &v
}

// GetAtTime returns the AtTime field value if set, zero value otherwise.
func (o *SentimentContent) GetAtTime() string {
	if o == nil || o.AtTime == nil {
		var ret string
		return ret
	}
	return *o.AtTime
}

// GetAtTimeOk returns a tuple with the AtTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SentimentContent) GetAtTimeOk() (*string, bool) {
	if o == nil || o.AtTime == nil {
		return nil, false
	}
	return o.AtTime, true
}

// HasAtTime returns a boolean if a field has been set.
func (o *SentimentContent) HasAtTime() bool {
	if o != nil && o.AtTime != nil {
		return true
	}

	return false
}

// SetAtTime gets a reference to the given string and assigns it to the AtTime field.
func (o *SentimentContent) SetAtTime(v string) {
	o.AtTime = &v
}

func (o SentimentContent) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Mention != nil {
		toSerialize["mention"] = o.Mention
	}
	if o.PositiveMention != nil {
		toSerialize["positiveMention"] = o.PositiveMention
	}
	if o.NegativeMention != nil {
		toSerialize["negativeMention"] = o.NegativeMention
	}
	if o.PositiveScore != nil {
		toSerialize["positiveScore"] = o.PositiveScore
	}
	if o.NegativeScore != nil {
		toSerialize["negativeScore"] = o.NegativeScore
	}
	if o.Score != nil {
		toSerialize["score"] = o.Score
	}
	if o.AtTime != nil {
		toSerialize["atTime"] = o.AtTime
	}
	return json.Marshal(toSerialize)
}

type NullableSentimentContent struct {
	value *SentimentContent
	isSet bool
}

func (v NullableSentimentContent) Get() *SentimentContent {
	return v.value
}

func (v *NullableSentimentContent) Set(val *SentimentContent) {
	v.value = val
	v.isSet = true
}

func (v NullableSentimentContent) IsSet() bool {
	return v.isSet
}

func (v *NullableSentimentContent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSentimentContent(val *SentimentContent) *NullableSentimentContent {
	return &NullableSentimentContent{value: val, isSet: true}
}

func (v NullableSentimentContent) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSentimentContent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


