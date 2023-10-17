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

// EarningsEstimatesInfo struct for EarningsEstimatesInfo
type EarningsEstimatesInfo struct {
	// Average EPS estimates including Finnhub's proprietary estimates.
	EpsAvg *float32 `json:"epsAvg,omitempty"`
	// Highest estimate.
	EpsHigh *float32 `json:"epsHigh,omitempty"`
	// Lowest estimate.
	EpsLow *float32 `json:"epsLow,omitempty"`
	// Number of Analysts.
	NumberAnalysts *int64 `json:"numberAnalysts,omitempty"`
	// Period.
	Period *string `json:"period,omitempty"`
	// Fiscal year.
	Year *int64 `json:"year,omitempty"`
	// Fiscal quarter.
	Quarter *int64 `json:"quarter,omitempty"`
}

// NewEarningsEstimatesInfo instantiates a new EarningsEstimatesInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEarningsEstimatesInfo() *EarningsEstimatesInfo {
	this := EarningsEstimatesInfo{}
	return &this
}

// NewEarningsEstimatesInfoWithDefaults instantiates a new EarningsEstimatesInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEarningsEstimatesInfoWithDefaults() *EarningsEstimatesInfo {
	this := EarningsEstimatesInfo{}
	return &this
}

// GetEpsAvg returns the EpsAvg field value if set, zero value otherwise.
func (o *EarningsEstimatesInfo) GetEpsAvg() float32 {
	if o == nil || o.EpsAvg == nil {
		var ret float32
		return ret
	}
	return *o.EpsAvg
}

// GetEpsAvgOk returns a tuple with the EpsAvg field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EarningsEstimatesInfo) GetEpsAvgOk() (*float32, bool) {
	if o == nil || o.EpsAvg == nil {
		return nil, false
	}
	return o.EpsAvg, true
}

// HasEpsAvg returns a boolean if a field has been set.
func (o *EarningsEstimatesInfo) HasEpsAvg() bool {
	if o != nil && o.EpsAvg != nil {
		return true
	}

	return false
}

// SetEpsAvg gets a reference to the given float32 and assigns it to the EpsAvg field.
func (o *EarningsEstimatesInfo) SetEpsAvg(v float32) {
	o.EpsAvg = &v
}

// GetEpsHigh returns the EpsHigh field value if set, zero value otherwise.
func (o *EarningsEstimatesInfo) GetEpsHigh() float32 {
	if o == nil || o.EpsHigh == nil {
		var ret float32
		return ret
	}
	return *o.EpsHigh
}

// GetEpsHighOk returns a tuple with the EpsHigh field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EarningsEstimatesInfo) GetEpsHighOk() (*float32, bool) {
	if o == nil || o.EpsHigh == nil {
		return nil, false
	}
	return o.EpsHigh, true
}

// HasEpsHigh returns a boolean if a field has been set.
func (o *EarningsEstimatesInfo) HasEpsHigh() bool {
	if o != nil && o.EpsHigh != nil {
		return true
	}

	return false
}

// SetEpsHigh gets a reference to the given float32 and assigns it to the EpsHigh field.
func (o *EarningsEstimatesInfo) SetEpsHigh(v float32) {
	o.EpsHigh = &v
}

// GetEpsLow returns the EpsLow field value if set, zero value otherwise.
func (o *EarningsEstimatesInfo) GetEpsLow() float32 {
	if o == nil || o.EpsLow == nil {
		var ret float32
		return ret
	}
	return *o.EpsLow
}

// GetEpsLowOk returns a tuple with the EpsLow field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EarningsEstimatesInfo) GetEpsLowOk() (*float32, bool) {
	if o == nil || o.EpsLow == nil {
		return nil, false
	}
	return o.EpsLow, true
}

// HasEpsLow returns a boolean if a field has been set.
func (o *EarningsEstimatesInfo) HasEpsLow() bool {
	if o != nil && o.EpsLow != nil {
		return true
	}

	return false
}

// SetEpsLow gets a reference to the given float32 and assigns it to the EpsLow field.
func (o *EarningsEstimatesInfo) SetEpsLow(v float32) {
	o.EpsLow = &v
}

// GetNumberAnalysts returns the NumberAnalysts field value if set, zero value otherwise.
func (o *EarningsEstimatesInfo) GetNumberAnalysts() int64 {
	if o == nil || o.NumberAnalysts == nil {
		var ret int64
		return ret
	}
	return *o.NumberAnalysts
}

// GetNumberAnalystsOk returns a tuple with the NumberAnalysts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EarningsEstimatesInfo) GetNumberAnalystsOk() (*int64, bool) {
	if o == nil || o.NumberAnalysts == nil {
		return nil, false
	}
	return o.NumberAnalysts, true
}

// HasNumberAnalysts returns a boolean if a field has been set.
func (o *EarningsEstimatesInfo) HasNumberAnalysts() bool {
	if o != nil && o.NumberAnalysts != nil {
		return true
	}

	return false
}

// SetNumberAnalysts gets a reference to the given int64 and assigns it to the NumberAnalysts field.
func (o *EarningsEstimatesInfo) SetNumberAnalysts(v int64) {
	o.NumberAnalysts = &v
}

// GetPeriod returns the Period field value if set, zero value otherwise.
func (o *EarningsEstimatesInfo) GetPeriod() string {
	if o == nil || o.Period == nil {
		var ret string
		return ret
	}
	return *o.Period
}

// GetPeriodOk returns a tuple with the Period field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EarningsEstimatesInfo) GetPeriodOk() (*string, bool) {
	if o == nil || o.Period == nil {
		return nil, false
	}
	return o.Period, true
}

// HasPeriod returns a boolean if a field has been set.
func (o *EarningsEstimatesInfo) HasPeriod() bool {
	if o != nil && o.Period != nil {
		return true
	}

	return false
}

// SetPeriod gets a reference to the given string and assigns it to the Period field.
func (o *EarningsEstimatesInfo) SetPeriod(v string) {
	o.Period = &v
}

// GetYear returns the Year field value if set, zero value otherwise.
func (o *EarningsEstimatesInfo) GetYear() int64 {
	if o == nil || o.Year == nil {
		var ret int64
		return ret
	}
	return *o.Year
}

// GetYearOk returns a tuple with the Year field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EarningsEstimatesInfo) GetYearOk() (*int64, bool) {
	if o == nil || o.Year == nil {
		return nil, false
	}
	return o.Year, true
}

// HasYear returns a boolean if a field has been set.
func (o *EarningsEstimatesInfo) HasYear() bool {
	if o != nil && o.Year != nil {
		return true
	}

	return false
}

// SetYear gets a reference to the given int64 and assigns it to the Year field.
func (o *EarningsEstimatesInfo) SetYear(v int64) {
	o.Year = &v
}

// GetQuarter returns the Quarter field value if set, zero value otherwise.
func (o *EarningsEstimatesInfo) GetQuarter() int64 {
	if o == nil || o.Quarter == nil {
		var ret int64
		return ret
	}
	return *o.Quarter
}

// GetQuarterOk returns a tuple with the Quarter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EarningsEstimatesInfo) GetQuarterOk() (*int64, bool) {
	if o == nil || o.Quarter == nil {
		return nil, false
	}
	return o.Quarter, true
}

// HasQuarter returns a boolean if a field has been set.
func (o *EarningsEstimatesInfo) HasQuarter() bool {
	if o != nil && o.Quarter != nil {
		return true
	}

	return false
}

// SetQuarter gets a reference to the given int64 and assigns it to the Quarter field.
func (o *EarningsEstimatesInfo) SetQuarter(v int64) {
	o.Quarter = &v
}

func (o EarningsEstimatesInfo) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.EpsAvg != nil {
		toSerialize["epsAvg"] = o.EpsAvg
	}
	if o.EpsHigh != nil {
		toSerialize["epsHigh"] = o.EpsHigh
	}
	if o.EpsLow != nil {
		toSerialize["epsLow"] = o.EpsLow
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

type NullableEarningsEstimatesInfo struct {
	value *EarningsEstimatesInfo
	isSet bool
}

func (v NullableEarningsEstimatesInfo) Get() *EarningsEstimatesInfo {
	return v.value
}

func (v *NullableEarningsEstimatesInfo) Set(val *EarningsEstimatesInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableEarningsEstimatesInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableEarningsEstimatesInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEarningsEstimatesInfo(val *EarningsEstimatesInfo) *NullableEarningsEstimatesInfo {
	return &NullableEarningsEstimatesInfo{value: val, isSet: true}
}

func (v NullableEarningsEstimatesInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEarningsEstimatesInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


