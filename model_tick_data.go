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

// TickData struct for TickData
type TickData struct {
	// Symbol.
	S *string `json:"s,omitempty"`
	// Number of ticks skipped.
	Skip *int64 `json:"skip,omitempty"`
	// Number of ticks returned. If <code>count</code> < <code>limit</code>, all data for that date has been returned.
	Count *int64 `json:"count,omitempty"`
	// Total number of ticks for that date.
	Total *int64 `json:"total,omitempty"`
	// List of volume data.
	V *[]float32 `json:"v,omitempty"`
	// List of price data.
	P *[]float32 `json:"p,omitempty"`
	// List of timestamp in UNIX ms.
	T *[]int64 `json:"t,omitempty"`
	// List of venues/exchanges. A list of exchange codes can be found <a target=\"_blank\" href=\"https://docs.google.com/spreadsheets/d/1Tj53M1svmr-hfEtbk6_NpVR1yAyGLMaH6ByYU6CG0ZY/edit?usp=sharing\",>here</a>
	X *[]string `json:"x,omitempty"`
	// List of trade conditions. A comprehensive list of trade conditions code can be found <a target=\"_blank\" href=\"https://docs.google.com/spreadsheets/d/1PUxiSWPHSODbaTaoL2Vef6DgU-yFtlRGZf19oBb9Hp0/edit?usp=sharing\">here</a>
	C *[][]string `json:"c,omitempty"`
}

// NewTickData instantiates a new TickData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTickData() *TickData {
	this := TickData{}
	return &this
}

// NewTickDataWithDefaults instantiates a new TickData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTickDataWithDefaults() *TickData {
	this := TickData{}
	return &this
}

// GetS returns the S field value if set, zero value otherwise.
func (o *TickData) GetS() string {
	if o == nil || o.S == nil {
		var ret string
		return ret
	}
	return *o.S
}

// GetSOk returns a tuple with the S field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TickData) GetSOk() (*string, bool) {
	if o == nil || o.S == nil {
		return nil, false
	}
	return o.S, true
}

// HasS returns a boolean if a field has been set.
func (o *TickData) HasS() bool {
	if o != nil && o.S != nil {
		return true
	}

	return false
}

// SetS gets a reference to the given string and assigns it to the S field.
func (o *TickData) SetS(v string) {
	o.S = &v
}

// GetSkip returns the Skip field value if set, zero value otherwise.
func (o *TickData) GetSkip() int64 {
	if o == nil || o.Skip == nil {
		var ret int64
		return ret
	}
	return *o.Skip
}

// GetSkipOk returns a tuple with the Skip field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TickData) GetSkipOk() (*int64, bool) {
	if o == nil || o.Skip == nil {
		return nil, false
	}
	return o.Skip, true
}

// HasSkip returns a boolean if a field has been set.
func (o *TickData) HasSkip() bool {
	if o != nil && o.Skip != nil {
		return true
	}

	return false
}

// SetSkip gets a reference to the given int64 and assigns it to the Skip field.
func (o *TickData) SetSkip(v int64) {
	o.Skip = &v
}

// GetCount returns the Count field value if set, zero value otherwise.
func (o *TickData) GetCount() int64 {
	if o == nil || o.Count == nil {
		var ret int64
		return ret
	}
	return *o.Count
}

// GetCountOk returns a tuple with the Count field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TickData) GetCountOk() (*int64, bool) {
	if o == nil || o.Count == nil {
		return nil, false
	}
	return o.Count, true
}

// HasCount returns a boolean if a field has been set.
func (o *TickData) HasCount() bool {
	if o != nil && o.Count != nil {
		return true
	}

	return false
}

// SetCount gets a reference to the given int64 and assigns it to the Count field.
func (o *TickData) SetCount(v int64) {
	o.Count = &v
}

// GetTotal returns the Total field value if set, zero value otherwise.
func (o *TickData) GetTotal() int64 {
	if o == nil || o.Total == nil {
		var ret int64
		return ret
	}
	return *o.Total
}

// GetTotalOk returns a tuple with the Total field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TickData) GetTotalOk() (*int64, bool) {
	if o == nil || o.Total == nil {
		return nil, false
	}
	return o.Total, true
}

// HasTotal returns a boolean if a field has been set.
func (o *TickData) HasTotal() bool {
	if o != nil && o.Total != nil {
		return true
	}

	return false
}

// SetTotal gets a reference to the given int64 and assigns it to the Total field.
func (o *TickData) SetTotal(v int64) {
	o.Total = &v
}

// GetV returns the V field value if set, zero value otherwise.
func (o *TickData) GetV() []float32 {
	if o == nil || o.V == nil {
		var ret []float32
		return ret
	}
	return *o.V
}

// GetVOk returns a tuple with the V field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TickData) GetVOk() (*[]float32, bool) {
	if o == nil || o.V == nil {
		return nil, false
	}
	return o.V, true
}

// HasV returns a boolean if a field has been set.
func (o *TickData) HasV() bool {
	if o != nil && o.V != nil {
		return true
	}

	return false
}

// SetV gets a reference to the given []float32 and assigns it to the V field.
func (o *TickData) SetV(v []float32) {
	o.V = &v
}

// GetP returns the P field value if set, zero value otherwise.
func (o *TickData) GetP() []float32 {
	if o == nil || o.P == nil {
		var ret []float32
		return ret
	}
	return *o.P
}

// GetPOk returns a tuple with the P field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TickData) GetPOk() (*[]float32, bool) {
	if o == nil || o.P == nil {
		return nil, false
	}
	return o.P, true
}

// HasP returns a boolean if a field has been set.
func (o *TickData) HasP() bool {
	if o != nil && o.P != nil {
		return true
	}

	return false
}

// SetP gets a reference to the given []float32 and assigns it to the P field.
func (o *TickData) SetP(v []float32) {
	o.P = &v
}

// GetT returns the T field value if set, zero value otherwise.
func (o *TickData) GetT() []int64 {
	if o == nil || o.T == nil {
		var ret []int64
		return ret
	}
	return *o.T
}

// GetTOk returns a tuple with the T field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TickData) GetTOk() (*[]int64, bool) {
	if o == nil || o.T == nil {
		return nil, false
	}
	return o.T, true
}

// HasT returns a boolean if a field has been set.
func (o *TickData) HasT() bool {
	if o != nil && o.T != nil {
		return true
	}

	return false
}

// SetT gets a reference to the given []int64 and assigns it to the T field.
func (o *TickData) SetT(v []int64) {
	o.T = &v
}

// GetX returns the X field value if set, zero value otherwise.
func (o *TickData) GetX() []string {
	if o == nil || o.X == nil {
		var ret []string
		return ret
	}
	return *o.X
}

// GetXOk returns a tuple with the X field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TickData) GetXOk() (*[]string, bool) {
	if o == nil || o.X == nil {
		return nil, false
	}
	return o.X, true
}

// HasX returns a boolean if a field has been set.
func (o *TickData) HasX() bool {
	if o != nil && o.X != nil {
		return true
	}

	return false
}

// SetX gets a reference to the given []string and assigns it to the X field.
func (o *TickData) SetX(v []string) {
	o.X = &v
}

// GetC returns the C field value if set, zero value otherwise.
func (o *TickData) GetC() [][]string {
	if o == nil || o.C == nil {
		var ret [][]string
		return ret
	}
	return *o.C
}

// GetCOk returns a tuple with the C field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TickData) GetCOk() (*[][]string, bool) {
	if o == nil || o.C == nil {
		return nil, false
	}
	return o.C, true
}

// HasC returns a boolean if a field has been set.
func (o *TickData) HasC() bool {
	if o != nil && o.C != nil {
		return true
	}

	return false
}

// SetC gets a reference to the given [][]string and assigns it to the C field.
func (o *TickData) SetC(v [][]string) {
	o.C = &v
}

func (o TickData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.S != nil {
		toSerialize["s"] = o.S
	}
	if o.Skip != nil {
		toSerialize["skip"] = o.Skip
	}
	if o.Count != nil {
		toSerialize["count"] = o.Count
	}
	if o.Total != nil {
		toSerialize["total"] = o.Total
	}
	if o.V != nil {
		toSerialize["v"] = o.V
	}
	if o.P != nil {
		toSerialize["p"] = o.P
	}
	if o.T != nil {
		toSerialize["t"] = o.T
	}
	if o.X != nil {
		toSerialize["x"] = o.X
	}
	if o.C != nil {
		toSerialize["c"] = o.C
	}
	return json.Marshal(toSerialize)
}

type NullableTickData struct {
	value *TickData
	isSet bool
}

func (v NullableTickData) Get() *TickData {
	return v.value
}

func (v *NullableTickData) Set(val *TickData) {
	v.value = val
	v.isSet = true
}

func (v NullableTickData) IsSet() bool {
	return v.isSet
}

func (v *NullableTickData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTickData(val *TickData) *NullableTickData {
	return &NullableTickData{value: val, isSet: true}
}

func (v NullableTickData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTickData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


