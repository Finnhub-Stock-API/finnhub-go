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

// BasicFinancials struct for BasicFinancials
type BasicFinancials struct {
	// Symbol of the company.
	Symbol *string `json:"symbol,omitempty"`
	// Metric type.
	MetricType *string `json:"metricType,omitempty"`
	Series *map[string]interface{} `json:"series,omitempty"`
	Metric *map[string]interface{} `json:"metric,omitempty"`
}

// NewBasicFinancials instantiates a new BasicFinancials object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBasicFinancials() *BasicFinancials {
	this := BasicFinancials{}
	return &this
}

// NewBasicFinancialsWithDefaults instantiates a new BasicFinancials object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBasicFinancialsWithDefaults() *BasicFinancials {
	this := BasicFinancials{}
	return &this
}

// GetSymbol returns the Symbol field value if set, zero value otherwise.
func (o *BasicFinancials) GetSymbol() string {
	if o == nil || o.Symbol == nil {
		var ret string
		return ret
	}
	return *o.Symbol
}

// GetSymbolOk returns a tuple with the Symbol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BasicFinancials) GetSymbolOk() (*string, bool) {
	if o == nil || o.Symbol == nil {
		return nil, false
	}
	return o.Symbol, true
}

// HasSymbol returns a boolean if a field has been set.
func (o *BasicFinancials) HasSymbol() bool {
	if o != nil && o.Symbol != nil {
		return true
	}

	return false
}

// SetSymbol gets a reference to the given string and assigns it to the Symbol field.
func (o *BasicFinancials) SetSymbol(v string) {
	o.Symbol = &v
}

// GetMetricType returns the MetricType field value if set, zero value otherwise.
func (o *BasicFinancials) GetMetricType() string {
	if o == nil || o.MetricType == nil {
		var ret string
		return ret
	}
	return *o.MetricType
}

// GetMetricTypeOk returns a tuple with the MetricType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BasicFinancials) GetMetricTypeOk() (*string, bool) {
	if o == nil || o.MetricType == nil {
		return nil, false
	}
	return o.MetricType, true
}

// HasMetricType returns a boolean if a field has been set.
func (o *BasicFinancials) HasMetricType() bool {
	if o != nil && o.MetricType != nil {
		return true
	}

	return false
}

// SetMetricType gets a reference to the given string and assigns it to the MetricType field.
func (o *BasicFinancials) SetMetricType(v string) {
	o.MetricType = &v
}

// GetSeries returns the Series field value if set, zero value otherwise.
func (o *BasicFinancials) GetSeries() map[string]interface{} {
	if o == nil || o.Series == nil {
		var ret map[string]interface{}
		return ret
	}
	return *o.Series
}

// GetSeriesOk returns a tuple with the Series field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BasicFinancials) GetSeriesOk() (*map[string]interface{}, bool) {
	if o == nil || o.Series == nil {
		return nil, false
	}
	return o.Series, true
}

// HasSeries returns a boolean if a field has been set.
func (o *BasicFinancials) HasSeries() bool {
	if o != nil && o.Series != nil {
		return true
	}

	return false
}

// SetSeries gets a reference to the given map[string]interface{} and assigns it to the Series field.
func (o *BasicFinancials) SetSeries(v map[string]interface{}) {
	o.Series = &v
}

// GetMetric returns the Metric field value if set, zero value otherwise.
func (o *BasicFinancials) GetMetric() map[string]interface{} {
	if o == nil || o.Metric == nil {
		var ret map[string]interface{}
		return ret
	}
	return *o.Metric
}

// GetMetricOk returns a tuple with the Metric field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BasicFinancials) GetMetricOk() (*map[string]interface{}, bool) {
	if o == nil || o.Metric == nil {
		return nil, false
	}
	return o.Metric, true
}

// HasMetric returns a boolean if a field has been set.
func (o *BasicFinancials) HasMetric() bool {
	if o != nil && o.Metric != nil {
		return true
	}

	return false
}

// SetMetric gets a reference to the given map[string]interface{} and assigns it to the Metric field.
func (o *BasicFinancials) SetMetric(v map[string]interface{}) {
	o.Metric = &v
}

func (o BasicFinancials) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Symbol != nil {
		toSerialize["symbol"] = o.Symbol
	}
	if o.MetricType != nil {
		toSerialize["metricType"] = o.MetricType
	}
	if o.Series != nil {
		toSerialize["series"] = o.Series
	}
	if o.Metric != nil {
		toSerialize["metric"] = o.Metric
	}
	return json.Marshal(toSerialize)
}

type NullableBasicFinancials struct {
	value *BasicFinancials
	isSet bool
}

func (v NullableBasicFinancials) Get() *BasicFinancials {
	return v.value
}

func (v *NullableBasicFinancials) Set(val *BasicFinancials) {
	v.value = val
	v.isSet = true
}

func (v NullableBasicFinancials) IsSet() bool {
	return v.isSet
}

func (v *NullableBasicFinancials) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBasicFinancials(val *BasicFinancials) *NullableBasicFinancials {
	return &NullableBasicFinancials{value: val, isSet: true}
}

func (v NullableBasicFinancials) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBasicFinancials) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


