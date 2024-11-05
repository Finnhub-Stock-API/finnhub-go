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

// AirlinePriceIndexData struct for AirlinePriceIndexData
type AirlinePriceIndexData struct {
	// Array of price index.
	Data *[]AirlinePriceIndex `json:"data,omitempty"`
	// Airline name
	Airline *string `json:"airline,omitempty"`
	// From date
	From *string `json:"from,omitempty"`
	// To date
	To *string `json:"to,omitempty"`
}

// NewAirlinePriceIndexData instantiates a new AirlinePriceIndexData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAirlinePriceIndexData() *AirlinePriceIndexData {
	this := AirlinePriceIndexData{}
	return &this
}

// NewAirlinePriceIndexDataWithDefaults instantiates a new AirlinePriceIndexData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAirlinePriceIndexDataWithDefaults() *AirlinePriceIndexData {
	this := AirlinePriceIndexData{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *AirlinePriceIndexData) GetData() []AirlinePriceIndex {
	if o == nil || o.Data == nil {
		var ret []AirlinePriceIndex
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AirlinePriceIndexData) GetDataOk() (*[]AirlinePriceIndex, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *AirlinePriceIndexData) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given []AirlinePriceIndex and assigns it to the Data field.
func (o *AirlinePriceIndexData) SetData(v []AirlinePriceIndex) {
	o.Data = &v
}

// GetAirline returns the Airline field value if set, zero value otherwise.
func (o *AirlinePriceIndexData) GetAirline() string {
	if o == nil || o.Airline == nil {
		var ret string
		return ret
	}
	return *o.Airline
}

// GetAirlineOk returns a tuple with the Airline field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AirlinePriceIndexData) GetAirlineOk() (*string, bool) {
	if o == nil || o.Airline == nil {
		return nil, false
	}
	return o.Airline, true
}

// HasAirline returns a boolean if a field has been set.
func (o *AirlinePriceIndexData) HasAirline() bool {
	if o != nil && o.Airline != nil {
		return true
	}

	return false
}

// SetAirline gets a reference to the given string and assigns it to the Airline field.
func (o *AirlinePriceIndexData) SetAirline(v string) {
	o.Airline = &v
}

// GetFrom returns the From field value if set, zero value otherwise.
func (o *AirlinePriceIndexData) GetFrom() string {
	if o == nil || o.From == nil {
		var ret string
		return ret
	}
	return *o.From
}

// GetFromOk returns a tuple with the From field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AirlinePriceIndexData) GetFromOk() (*string, bool) {
	if o == nil || o.From == nil {
		return nil, false
	}
	return o.From, true
}

// HasFrom returns a boolean if a field has been set.
func (o *AirlinePriceIndexData) HasFrom() bool {
	if o != nil && o.From != nil {
		return true
	}

	return false
}

// SetFrom gets a reference to the given string and assigns it to the From field.
func (o *AirlinePriceIndexData) SetFrom(v string) {
	o.From = &v
}

// GetTo returns the To field value if set, zero value otherwise.
func (o *AirlinePriceIndexData) GetTo() string {
	if o == nil || o.To == nil {
		var ret string
		return ret
	}
	return *o.To
}

// GetToOk returns a tuple with the To field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AirlinePriceIndexData) GetToOk() (*string, bool) {
	if o == nil || o.To == nil {
		return nil, false
	}
	return o.To, true
}

// HasTo returns a boolean if a field has been set.
func (o *AirlinePriceIndexData) HasTo() bool {
	if o != nil && o.To != nil {
		return true
	}

	return false
}

// SetTo gets a reference to the given string and assigns it to the To field.
func (o *AirlinePriceIndexData) SetTo(v string) {
	o.To = &v
}

func (o AirlinePriceIndexData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	if o.Airline != nil {
		toSerialize["airline"] = o.Airline
	}
	if o.From != nil {
		toSerialize["from"] = o.From
	}
	if o.To != nil {
		toSerialize["to"] = o.To
	}
	return json.Marshal(toSerialize)
}

type NullableAirlinePriceIndexData struct {
	value *AirlinePriceIndexData
	isSet bool
}

func (v NullableAirlinePriceIndexData) Get() *AirlinePriceIndexData {
	return v.value
}

func (v *NullableAirlinePriceIndexData) Set(val *AirlinePriceIndexData) {
	v.value = val
	v.isSet = true
}

func (v NullableAirlinePriceIndexData) IsSet() bool {
	return v.isSet
}

func (v *NullableAirlinePriceIndexData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAirlinePriceIndexData(val *AirlinePriceIndexData) *NullableAirlinePriceIndexData {
	return &NullableAirlinePriceIndexData{value: val, isSet: true}
}

func (v NullableAirlinePriceIndexData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAirlinePriceIndexData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

