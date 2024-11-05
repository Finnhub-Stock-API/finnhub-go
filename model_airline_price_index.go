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

// AirlinePriceIndex struct for AirlinePriceIndex
type AirlinePriceIndex struct {
	// Date
	Date *string `json:"date,omitempty"`
	// Price Index
	PriceIndex *float32 `json:"priceIndex,omitempty"`
	// Daily average ticket price.
	DailyAvgPrice *float32 `json:"dailyAvgPrice,omitempty"`
}

// NewAirlinePriceIndex instantiates a new AirlinePriceIndex object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAirlinePriceIndex() *AirlinePriceIndex {
	this := AirlinePriceIndex{}
	return &this
}

// NewAirlinePriceIndexWithDefaults instantiates a new AirlinePriceIndex object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAirlinePriceIndexWithDefaults() *AirlinePriceIndex {
	this := AirlinePriceIndex{}
	return &this
}

// GetDate returns the Date field value if set, zero value otherwise.
func (o *AirlinePriceIndex) GetDate() string {
	if o == nil || o.Date == nil {
		var ret string
		return ret
	}
	return *o.Date
}

// GetDateOk returns a tuple with the Date field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AirlinePriceIndex) GetDateOk() (*string, bool) {
	if o == nil || o.Date == nil {
		return nil, false
	}
	return o.Date, true
}

// HasDate returns a boolean if a field has been set.
func (o *AirlinePriceIndex) HasDate() bool {
	if o != nil && o.Date != nil {
		return true
	}

	return false
}

// SetDate gets a reference to the given string and assigns it to the Date field.
func (o *AirlinePriceIndex) SetDate(v string) {
	o.Date = &v
}

// GetPriceIndex returns the PriceIndex field value if set, zero value otherwise.
func (o *AirlinePriceIndex) GetPriceIndex() float32 {
	if o == nil || o.PriceIndex == nil {
		var ret float32
		return ret
	}
	return *o.PriceIndex
}

// GetPriceIndexOk returns a tuple with the PriceIndex field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AirlinePriceIndex) GetPriceIndexOk() (*float32, bool) {
	if o == nil || o.PriceIndex == nil {
		return nil, false
	}
	return o.PriceIndex, true
}

// HasPriceIndex returns a boolean if a field has been set.
func (o *AirlinePriceIndex) HasPriceIndex() bool {
	if o != nil && o.PriceIndex != nil {
		return true
	}

	return false
}

// SetPriceIndex gets a reference to the given float32 and assigns it to the PriceIndex field.
func (o *AirlinePriceIndex) SetPriceIndex(v float32) {
	o.PriceIndex = &v
}

// GetDailyAvgPrice returns the DailyAvgPrice field value if set, zero value otherwise.
func (o *AirlinePriceIndex) GetDailyAvgPrice() float32 {
	if o == nil || o.DailyAvgPrice == nil {
		var ret float32
		return ret
	}
	return *o.DailyAvgPrice
}

// GetDailyAvgPriceOk returns a tuple with the DailyAvgPrice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AirlinePriceIndex) GetDailyAvgPriceOk() (*float32, bool) {
	if o == nil || o.DailyAvgPrice == nil {
		return nil, false
	}
	return o.DailyAvgPrice, true
}

// HasDailyAvgPrice returns a boolean if a field has been set.
func (o *AirlinePriceIndex) HasDailyAvgPrice() bool {
	if o != nil && o.DailyAvgPrice != nil {
		return true
	}

	return false
}

// SetDailyAvgPrice gets a reference to the given float32 and assigns it to the DailyAvgPrice field.
func (o *AirlinePriceIndex) SetDailyAvgPrice(v float32) {
	o.DailyAvgPrice = &v
}

func (o AirlinePriceIndex) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Date != nil {
		toSerialize["date"] = o.Date
	}
	if o.PriceIndex != nil {
		toSerialize["priceIndex"] = o.PriceIndex
	}
	if o.DailyAvgPrice != nil {
		toSerialize["dailyAvgPrice"] = o.DailyAvgPrice
	}
	return json.Marshal(toSerialize)
}

type NullableAirlinePriceIndex struct {
	value *AirlinePriceIndex
	isSet bool
}

func (v NullableAirlinePriceIndex) Get() *AirlinePriceIndex {
	return v.value
}

func (v *NullableAirlinePriceIndex) Set(val *AirlinePriceIndex) {
	v.value = val
	v.isSet = true
}

func (v NullableAirlinePriceIndex) IsSet() bool {
	return v.isSet
}

func (v *NullableAirlinePriceIndex) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAirlinePriceIndex(val *AirlinePriceIndex) *NullableAirlinePriceIndex {
	return &NullableAirlinePriceIndex{value: val, isSet: true}
}

func (v NullableAirlinePriceIndex) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAirlinePriceIndex) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

