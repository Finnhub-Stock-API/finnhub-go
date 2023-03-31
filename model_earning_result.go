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

// EarningResult struct for EarningResult
type EarningResult struct {
	// Actual earning result.
	Actual *float32 `json:"actual,omitempty"`
	// Estimated earning.
	Estimate *float32 `json:"estimate,omitempty"`
	// Surprise - The difference between actual and estimate.
	Surprise *float32 `json:"surprise,omitempty"`
	// Surprise percent.
	SurprisePercent *float32 `json:"surprisePercent,omitempty"`
	// Reported period.
	Period *string `json:"period,omitempty"`
	// Company symbol.
	Symbol *string `json:"symbol,omitempty"`
	// Earnings year.
	Year *int64 `json:"year,omitempty"`
	// Earnings quarter.
	Quarter *int64 `json:"quarter,omitempty"`
}

// NewEarningResult instantiates a new EarningResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEarningResult() *EarningResult {
	this := EarningResult{}
	return &this
}

// NewEarningResultWithDefaults instantiates a new EarningResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEarningResultWithDefaults() *EarningResult {
	this := EarningResult{}
	return &this
}

// GetActual returns the Actual field value if set, zero value otherwise.
func (o *EarningResult) GetActual() float32 {
	if o == nil || o.Actual == nil {
		var ret float32
		return ret
	}
	return *o.Actual
}

// GetActualOk returns a tuple with the Actual field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EarningResult) GetActualOk() (*float32, bool) {
	if o == nil || o.Actual == nil {
		return nil, false
	}
	return o.Actual, true
}

// HasActual returns a boolean if a field has been set.
func (o *EarningResult) HasActual() bool {
	if o != nil && o.Actual != nil {
		return true
	}

	return false
}

// SetActual gets a reference to the given float32 and assigns it to the Actual field.
func (o *EarningResult) SetActual(v float32) {
	o.Actual = &v
}

// GetEstimate returns the Estimate field value if set, zero value otherwise.
func (o *EarningResult) GetEstimate() float32 {
	if o == nil || o.Estimate == nil {
		var ret float32
		return ret
	}
	return *o.Estimate
}

// GetEstimateOk returns a tuple with the Estimate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EarningResult) GetEstimateOk() (*float32, bool) {
	if o == nil || o.Estimate == nil {
		return nil, false
	}
	return o.Estimate, true
}

// HasEstimate returns a boolean if a field has been set.
func (o *EarningResult) HasEstimate() bool {
	if o != nil && o.Estimate != nil {
		return true
	}

	return false
}

// SetEstimate gets a reference to the given float32 and assigns it to the Estimate field.
func (o *EarningResult) SetEstimate(v float32) {
	o.Estimate = &v
}

// GetSurprise returns the Surprise field value if set, zero value otherwise.
func (o *EarningResult) GetSurprise() float32 {
	if o == nil || o.Surprise == nil {
		var ret float32
		return ret
	}
	return *o.Surprise
}

// GetSurpriseOk returns a tuple with the Surprise field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EarningResult) GetSurpriseOk() (*float32, bool) {
	if o == nil || o.Surprise == nil {
		return nil, false
	}
	return o.Surprise, true
}

// HasSurprise returns a boolean if a field has been set.
func (o *EarningResult) HasSurprise() bool {
	if o != nil && o.Surprise != nil {
		return true
	}

	return false
}

// SetSurprise gets a reference to the given float32 and assigns it to the Surprise field.
func (o *EarningResult) SetSurprise(v float32) {
	o.Surprise = &v
}

// GetSurprisePercent returns the SurprisePercent field value if set, zero value otherwise.
func (o *EarningResult) GetSurprisePercent() float32 {
	if o == nil || o.SurprisePercent == nil {
		var ret float32
		return ret
	}
	return *o.SurprisePercent
}

// GetSurprisePercentOk returns a tuple with the SurprisePercent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EarningResult) GetSurprisePercentOk() (*float32, bool) {
	if o == nil || o.SurprisePercent == nil {
		return nil, false
	}
	return o.SurprisePercent, true
}

// HasSurprisePercent returns a boolean if a field has been set.
func (o *EarningResult) HasSurprisePercent() bool {
	if o != nil && o.SurprisePercent != nil {
		return true
	}

	return false
}

// SetSurprisePercent gets a reference to the given float32 and assigns it to the SurprisePercent field.
func (o *EarningResult) SetSurprisePercent(v float32) {
	o.SurprisePercent = &v
}

// GetPeriod returns the Period field value if set, zero value otherwise.
func (o *EarningResult) GetPeriod() string {
	if o == nil || o.Period == nil {
		var ret string
		return ret
	}
	return *o.Period
}

// GetPeriodOk returns a tuple with the Period field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EarningResult) GetPeriodOk() (*string, bool) {
	if o == nil || o.Period == nil {
		return nil, false
	}
	return o.Period, true
}

// HasPeriod returns a boolean if a field has been set.
func (o *EarningResult) HasPeriod() bool {
	if o != nil && o.Period != nil {
		return true
	}

	return false
}

// SetPeriod gets a reference to the given string and assigns it to the Period field.
func (o *EarningResult) SetPeriod(v string) {
	o.Period = &v
}

// GetSymbol returns the Symbol field value if set, zero value otherwise.
func (o *EarningResult) GetSymbol() string {
	if o == nil || o.Symbol == nil {
		var ret string
		return ret
	}
	return *o.Symbol
}

// GetSymbolOk returns a tuple with the Symbol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EarningResult) GetSymbolOk() (*string, bool) {
	if o == nil || o.Symbol == nil {
		return nil, false
	}
	return o.Symbol, true
}

// HasSymbol returns a boolean if a field has been set.
func (o *EarningResult) HasSymbol() bool {
	if o != nil && o.Symbol != nil {
		return true
	}

	return false
}

// SetSymbol gets a reference to the given string and assigns it to the Symbol field.
func (o *EarningResult) SetSymbol(v string) {
	o.Symbol = &v
}

// GetYear returns the Year field value if set, zero value otherwise.
func (o *EarningResult) GetYear() int64 {
	if o == nil || o.Year == nil {
		var ret int64
		return ret
	}
	return *o.Year
}

// GetYearOk returns a tuple with the Year field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EarningResult) GetYearOk() (*int64, bool) {
	if o == nil || o.Year == nil {
		return nil, false
	}
	return o.Year, true
}

// HasYear returns a boolean if a field has been set.
func (o *EarningResult) HasYear() bool {
	if o != nil && o.Year != nil {
		return true
	}

	return false
}

// SetYear gets a reference to the given int64 and assigns it to the Year field.
func (o *EarningResult) SetYear(v int64) {
	o.Year = &v
}

// GetQuarter returns the Quarter field value if set, zero value otherwise.
func (o *EarningResult) GetQuarter() int64 {
	if o == nil || o.Quarter == nil {
		var ret int64
		return ret
	}
	return *o.Quarter
}

// GetQuarterOk returns a tuple with the Quarter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EarningResult) GetQuarterOk() (*int64, bool) {
	if o == nil || o.Quarter == nil {
		return nil, false
	}
	return o.Quarter, true
}

// HasQuarter returns a boolean if a field has been set.
func (o *EarningResult) HasQuarter() bool {
	if o != nil && o.Quarter != nil {
		return true
	}

	return false
}

// SetQuarter gets a reference to the given int64 and assigns it to the Quarter field.
func (o *EarningResult) SetQuarter(v int64) {
	o.Quarter = &v
}

func (o EarningResult) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Actual != nil {
		toSerialize["actual"] = o.Actual
	}
	if o.Estimate != nil {
		toSerialize["estimate"] = o.Estimate
	}
	if o.Surprise != nil {
		toSerialize["surprise"] = o.Surprise
	}
	if o.SurprisePercent != nil {
		toSerialize["surprisePercent"] = o.SurprisePercent
	}
	if o.Period != nil {
		toSerialize["period"] = o.Period
	}
	if o.Symbol != nil {
		toSerialize["symbol"] = o.Symbol
	}
	if o.Year != nil {
		toSerialize["year"] = o.Year
	}
	if o.Quarter != nil {
		toSerialize["quarter"] = o.Quarter
	}
	return json.Marshal(toSerialize)
}

type NullableEarningResult struct {
	value *EarningResult
	isSet bool
}

func (v NullableEarningResult) Get() *EarningResult {
	return v.value
}

func (v *NullableEarningResult) Set(val *EarningResult) {
	v.value = val
	v.isSet = true
}

func (v NullableEarningResult) IsSet() bool {
	return v.isSet
}

func (v *NullableEarningResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEarningResult(val *EarningResult) *NullableEarningResult {
	return &NullableEarningResult{value: val, isSet: true}
}

func (v NullableEarningResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEarningResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


