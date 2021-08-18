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

// EconomicEvent struct for EconomicEvent
type EconomicEvent struct {
	// Actual release
	Actual *float32 `json:"actual,omitempty"`
	// Previous release
	Prev *float32 `json:"prev,omitempty"`
	// Country
	Country *string `json:"country,omitempty"`
	// Unit
	Unit *string `json:"unit,omitempty"`
	// Estimate
	Estimate *float32 `json:"estimate,omitempty"`
	// Event
	Event *string `json:"event,omitempty"`
	// Impact level
	Impact *string `json:"impact,omitempty"`
	// Release time
	Time *string `json:"time,omitempty"`
}

// NewEconomicEvent instantiates a new EconomicEvent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEconomicEvent() *EconomicEvent {
	this := EconomicEvent{}
	return &this
}

// NewEconomicEventWithDefaults instantiates a new EconomicEvent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEconomicEventWithDefaults() *EconomicEvent {
	this := EconomicEvent{}
	return &this
}

// GetActual returns the Actual field value if set, zero value otherwise.
func (o *EconomicEvent) GetActual() float32 {
	if o == nil || o.Actual == nil {
		var ret float32
		return ret
	}
	return *o.Actual
}

// GetActualOk returns a tuple with the Actual field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EconomicEvent) GetActualOk() (*float32, bool) {
	if o == nil || o.Actual == nil {
		return nil, false
	}
	return o.Actual, true
}

// HasActual returns a boolean if a field has been set.
func (o *EconomicEvent) HasActual() bool {
	if o != nil && o.Actual != nil {
		return true
	}

	return false
}

// SetActual gets a reference to the given float32 and assigns it to the Actual field.
func (o *EconomicEvent) SetActual(v float32) {
	o.Actual = &v
}

// GetPrev returns the Prev field value if set, zero value otherwise.
func (o *EconomicEvent) GetPrev() float32 {
	if o == nil || o.Prev == nil {
		var ret float32
		return ret
	}
	return *o.Prev
}

// GetPrevOk returns a tuple with the Prev field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EconomicEvent) GetPrevOk() (*float32, bool) {
	if o == nil || o.Prev == nil {
		return nil, false
	}
	return o.Prev, true
}

// HasPrev returns a boolean if a field has been set.
func (o *EconomicEvent) HasPrev() bool {
	if o != nil && o.Prev != nil {
		return true
	}

	return false
}

// SetPrev gets a reference to the given float32 and assigns it to the Prev field.
func (o *EconomicEvent) SetPrev(v float32) {
	o.Prev = &v
}

// GetCountry returns the Country field value if set, zero value otherwise.
func (o *EconomicEvent) GetCountry() string {
	if o == nil || o.Country == nil {
		var ret string
		return ret
	}
	return *o.Country
}

// GetCountryOk returns a tuple with the Country field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EconomicEvent) GetCountryOk() (*string, bool) {
	if o == nil || o.Country == nil {
		return nil, false
	}
	return o.Country, true
}

// HasCountry returns a boolean if a field has been set.
func (o *EconomicEvent) HasCountry() bool {
	if o != nil && o.Country != nil {
		return true
	}

	return false
}

// SetCountry gets a reference to the given string and assigns it to the Country field.
func (o *EconomicEvent) SetCountry(v string) {
	o.Country = &v
}

// GetUnit returns the Unit field value if set, zero value otherwise.
func (o *EconomicEvent) GetUnit() string {
	if o == nil || o.Unit == nil {
		var ret string
		return ret
	}
	return *o.Unit
}

// GetUnitOk returns a tuple with the Unit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EconomicEvent) GetUnitOk() (*string, bool) {
	if o == nil || o.Unit == nil {
		return nil, false
	}
	return o.Unit, true
}

// HasUnit returns a boolean if a field has been set.
func (o *EconomicEvent) HasUnit() bool {
	if o != nil && o.Unit != nil {
		return true
	}

	return false
}

// SetUnit gets a reference to the given string and assigns it to the Unit field.
func (o *EconomicEvent) SetUnit(v string) {
	o.Unit = &v
}

// GetEstimate returns the Estimate field value if set, zero value otherwise.
func (o *EconomicEvent) GetEstimate() float32 {
	if o == nil || o.Estimate == nil {
		var ret float32
		return ret
	}
	return *o.Estimate
}

// GetEstimateOk returns a tuple with the Estimate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EconomicEvent) GetEstimateOk() (*float32, bool) {
	if o == nil || o.Estimate == nil {
		return nil, false
	}
	return o.Estimate, true
}

// HasEstimate returns a boolean if a field has been set.
func (o *EconomicEvent) HasEstimate() bool {
	if o != nil && o.Estimate != nil {
		return true
	}

	return false
}

// SetEstimate gets a reference to the given float32 and assigns it to the Estimate field.
func (o *EconomicEvent) SetEstimate(v float32) {
	o.Estimate = &v
}

// GetEvent returns the Event field value if set, zero value otherwise.
func (o *EconomicEvent) GetEvent() string {
	if o == nil || o.Event == nil {
		var ret string
		return ret
	}
	return *o.Event
}

// GetEventOk returns a tuple with the Event field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EconomicEvent) GetEventOk() (*string, bool) {
	if o == nil || o.Event == nil {
		return nil, false
	}
	return o.Event, true
}

// HasEvent returns a boolean if a field has been set.
func (o *EconomicEvent) HasEvent() bool {
	if o != nil && o.Event != nil {
		return true
	}

	return false
}

// SetEvent gets a reference to the given string and assigns it to the Event field.
func (o *EconomicEvent) SetEvent(v string) {
	o.Event = &v
}

// GetImpact returns the Impact field value if set, zero value otherwise.
func (o *EconomicEvent) GetImpact() string {
	if o == nil || o.Impact == nil {
		var ret string
		return ret
	}
	return *o.Impact
}

// GetImpactOk returns a tuple with the Impact field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EconomicEvent) GetImpactOk() (*string, bool) {
	if o == nil || o.Impact == nil {
		return nil, false
	}
	return o.Impact, true
}

// HasImpact returns a boolean if a field has been set.
func (o *EconomicEvent) HasImpact() bool {
	if o != nil && o.Impact != nil {
		return true
	}

	return false
}

// SetImpact gets a reference to the given string and assigns it to the Impact field.
func (o *EconomicEvent) SetImpact(v string) {
	o.Impact = &v
}

// GetTime returns the Time field value if set, zero value otherwise.
func (o *EconomicEvent) GetTime() string {
	if o == nil || o.Time == nil {
		var ret string
		return ret
	}
	return *o.Time
}

// GetTimeOk returns a tuple with the Time field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EconomicEvent) GetTimeOk() (*string, bool) {
	if o == nil || o.Time == nil {
		return nil, false
	}
	return o.Time, true
}

// HasTime returns a boolean if a field has been set.
func (o *EconomicEvent) HasTime() bool {
	if o != nil && o.Time != nil {
		return true
	}

	return false
}

// SetTime gets a reference to the given string and assigns it to the Time field.
func (o *EconomicEvent) SetTime(v string) {
	o.Time = &v
}

func (o EconomicEvent) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Actual != nil {
		toSerialize["actual"] = o.Actual
	}
	if o.Prev != nil {
		toSerialize["prev"] = o.Prev
	}
	if o.Country != nil {
		toSerialize["country"] = o.Country
	}
	if o.Unit != nil {
		toSerialize["unit"] = o.Unit
	}
	if o.Estimate != nil {
		toSerialize["estimate"] = o.Estimate
	}
	if o.Event != nil {
		toSerialize["event"] = o.Event
	}
	if o.Impact != nil {
		toSerialize["impact"] = o.Impact
	}
	if o.Time != nil {
		toSerialize["time"] = o.Time
	}
	return json.Marshal(toSerialize)
}

type NullableEconomicEvent struct {
	value *EconomicEvent
	isSet bool
}

func (v NullableEconomicEvent) Get() *EconomicEvent {
	return v.value
}

func (v *NullableEconomicEvent) Set(val *EconomicEvent) {
	v.value = val
	v.isSet = true
}

func (v NullableEconomicEvent) IsSet() bool {
	return v.isSet
}

func (v *NullableEconomicEvent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEconomicEvent(val *EconomicEvent) *NullableEconomicEvent {
	return &NullableEconomicEvent{value: val, isSet: true}
}

func (v NullableEconomicEvent) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEconomicEvent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


