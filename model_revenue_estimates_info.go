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

// RevenueEstimatesInfo struct for RevenueEstimatesInfo
type RevenueEstimatesInfo struct {
	// Average revenue estimates including Finnhub's proprietary estimates.
	RevenueAvg *float32 `json:"revenueAvg,omitempty"`
	// Highest estimate.
	RevenueHigh *float32 `json:"revenueHigh,omitempty"`
	// Lowest estimate.
	RevenueLow *float32 `json:"revenueLow,omitempty"`
	// Number of Analysts.
	NumberAnalysts *int64 `json:"numberAnalysts,omitempty"`
	// Period.
	Period *string `json:"period,omitempty"`
	// Fiscal year.
	Year *int64 `json:"year,omitempty"`
	// Fiscal quarter.
	Quarter *int64 `json:"quarter,omitempty"`
}

// NewRevenueEstimatesInfo instantiates a new RevenueEstimatesInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRevenueEstimatesInfo() *RevenueEstimatesInfo {
	this := RevenueEstimatesInfo{}
	return &this
}

// NewRevenueEstimatesInfoWithDefaults instantiates a new RevenueEstimatesInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRevenueEstimatesInfoWithDefaults() *RevenueEstimatesInfo {
	this := RevenueEstimatesInfo{}
	return &this
}

// GetRevenueAvg returns the RevenueAvg field value if set, zero value otherwise.
func (o *RevenueEstimatesInfo) GetRevenueAvg() float32 {
	if o == nil || o.RevenueAvg == nil {
		var ret float32
		return ret
	}
	return *o.RevenueAvg
}

// GetRevenueAvgOk returns a tuple with the RevenueAvg field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RevenueEstimatesInfo) GetRevenueAvgOk() (*float32, bool) {
	if o == nil || o.RevenueAvg == nil {
		return nil, false
	}
	return o.RevenueAvg, true
}

// HasRevenueAvg returns a boolean if a field has been set.
func (o *RevenueEstimatesInfo) HasRevenueAvg() bool {
	if o != nil && o.RevenueAvg != nil {
		return true
	}

	return false
}

// SetRevenueAvg gets a reference to the given float32 and assigns it to the RevenueAvg field.
func (o *RevenueEstimatesInfo) SetRevenueAvg(v float32) {
	o.RevenueAvg = &v
}

// GetRevenueHigh returns the RevenueHigh field value if set, zero value otherwise.
func (o *RevenueEstimatesInfo) GetRevenueHigh() float32 {
	if o == nil || o.RevenueHigh == nil {
		var ret float32
		return ret
	}
	return *o.RevenueHigh
}

// GetRevenueHighOk returns a tuple with the RevenueHigh field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RevenueEstimatesInfo) GetRevenueHighOk() (*float32, bool) {
	if o == nil || o.RevenueHigh == nil {
		return nil, false
	}
	return o.RevenueHigh, true
}

// HasRevenueHigh returns a boolean if a field has been set.
func (o *RevenueEstimatesInfo) HasRevenueHigh() bool {
	if o != nil && o.RevenueHigh != nil {
		return true
	}

	return false
}

// SetRevenueHigh gets a reference to the given float32 and assigns it to the RevenueHigh field.
func (o *RevenueEstimatesInfo) SetRevenueHigh(v float32) {
	o.RevenueHigh = &v
}

// GetRevenueLow returns the RevenueLow field value if set, zero value otherwise.
func (o *RevenueEstimatesInfo) GetRevenueLow() float32 {
	if o == nil || o.RevenueLow == nil {
		var ret float32
		return ret
	}
	return *o.RevenueLow
}

// GetRevenueLowOk returns a tuple with the RevenueLow field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RevenueEstimatesInfo) GetRevenueLowOk() (*float32, bool) {
	if o == nil || o.RevenueLow == nil {
		return nil, false
	}
	return o.RevenueLow, true
}

// HasRevenueLow returns a boolean if a field has been set.
func (o *RevenueEstimatesInfo) HasRevenueLow() bool {
	if o != nil && o.RevenueLow != nil {
		return true
	}

	return false
}

// SetRevenueLow gets a reference to the given float32 and assigns it to the RevenueLow field.
func (o *RevenueEstimatesInfo) SetRevenueLow(v float32) {
	o.RevenueLow = &v
}

// GetNumberAnalysts returns the NumberAnalysts field value if set, zero value otherwise.
func (o *RevenueEstimatesInfo) GetNumberAnalysts() int64 {
	if o == nil || o.NumberAnalysts == nil {
		var ret int64
		return ret
	}
	return *o.NumberAnalysts
}

// GetNumberAnalystsOk returns a tuple with the NumberAnalysts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RevenueEstimatesInfo) GetNumberAnalystsOk() (*int64, bool) {
	if o == nil || o.NumberAnalysts == nil {
		return nil, false
	}
	return o.NumberAnalysts, true
}

// HasNumberAnalysts returns a boolean if a field has been set.
func (o *RevenueEstimatesInfo) HasNumberAnalysts() bool {
	if o != nil && o.NumberAnalysts != nil {
		return true
	}

	return false
}

// SetNumberAnalysts gets a reference to the given int64 and assigns it to the NumberAnalysts field.
func (o *RevenueEstimatesInfo) SetNumberAnalysts(v int64) {
	o.NumberAnalysts = &v
}

// GetPeriod returns the Period field value if set, zero value otherwise.
func (o *RevenueEstimatesInfo) GetPeriod() string {
	if o == nil || o.Period == nil {
		var ret string
		return ret
	}
	return *o.Period
}

// GetPeriodOk returns a tuple with the Period field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RevenueEstimatesInfo) GetPeriodOk() (*string, bool) {
	if o == nil || o.Period == nil {
		return nil, false
	}
	return o.Period, true
}

// HasPeriod returns a boolean if a field has been set.
func (o *RevenueEstimatesInfo) HasPeriod() bool {
	if o != nil && o.Period != nil {
		return true
	}

	return false
}

// SetPeriod gets a reference to the given string and assigns it to the Period field.
func (o *RevenueEstimatesInfo) SetPeriod(v string) {
	o.Period = &v
}

// GetYear returns the Year field value if set, zero value otherwise.
func (o *RevenueEstimatesInfo) GetYear() int64 {
	if o == nil || o.Year == nil {
		var ret int64
		return ret
	}
	return *o.Year
}

// GetYearOk returns a tuple with the Year field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RevenueEstimatesInfo) GetYearOk() (*int64, bool) {
	if o == nil || o.Year == nil {
		return nil, false
	}
	return o.Year, true
}

// HasYear returns a boolean if a field has been set.
func (o *RevenueEstimatesInfo) HasYear() bool {
	if o != nil && o.Year != nil {
		return true
	}

	return false
}

// SetYear gets a reference to the given int64 and assigns it to the Year field.
func (o *RevenueEstimatesInfo) SetYear(v int64) {
	o.Year = &v
}

// GetQuarter returns the Quarter field value if set, zero value otherwise.
func (o *RevenueEstimatesInfo) GetQuarter() int64 {
	if o == nil || o.Quarter == nil {
		var ret int64
		return ret
	}
	return *o.Quarter
}

// GetQuarterOk returns a tuple with the Quarter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RevenueEstimatesInfo) GetQuarterOk() (*int64, bool) {
	if o == nil || o.Quarter == nil {
		return nil, false
	}
	return o.Quarter, true
}

// HasQuarter returns a boolean if a field has been set.
func (o *RevenueEstimatesInfo) HasQuarter() bool {
	if o != nil && o.Quarter != nil {
		return true
	}

	return false
}

// SetQuarter gets a reference to the given int64 and assigns it to the Quarter field.
func (o *RevenueEstimatesInfo) SetQuarter(v int64) {
	o.Quarter = &v
}

func (o RevenueEstimatesInfo) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.RevenueAvg != nil {
		toSerialize["revenueAvg"] = o.RevenueAvg
	}
	if o.RevenueHigh != nil {
		toSerialize["revenueHigh"] = o.RevenueHigh
	}
	if o.RevenueLow != nil {
		toSerialize["revenueLow"] = o.RevenueLow
	}
	if o.NumberAnalysts != nil {
		toSerialize["numberAnalysts"] = o.NumberAnalysts
	}
	if o.Period != nil {
		toSerialize["period"] = o.Period
	}
	if o.Year != nil {
		toSerialize["year"] = o.Year
	}
	if o.Quarter != nil {
		toSerialize["quarter"] = o.Quarter
	}
	return json.Marshal(toSerialize)
}

type NullableRevenueEstimatesInfo struct {
	value *RevenueEstimatesInfo
	isSet bool
}

func (v NullableRevenueEstimatesInfo) Get() *RevenueEstimatesInfo {
	return v.value
}

func (v *NullableRevenueEstimatesInfo) Set(val *RevenueEstimatesInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableRevenueEstimatesInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableRevenueEstimatesInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRevenueEstimatesInfo(val *RevenueEstimatesInfo) *NullableRevenueEstimatesInfo {
	return &NullableRevenueEstimatesInfo{value: val, isSet: true}
}

func (v NullableRevenueEstimatesInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRevenueEstimatesInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


