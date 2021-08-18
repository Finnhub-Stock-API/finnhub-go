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

// SupplyChainRelationships struct for SupplyChainRelationships
type SupplyChainRelationships struct {
	// symbol
	Symbol *string `json:"symbol,omitempty"`
	// Key customers and suppliers.
	Data *[]KeyCustomersSuppliers `json:"data,omitempty"`
}

// NewSupplyChainRelationships instantiates a new SupplyChainRelationships object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSupplyChainRelationships() *SupplyChainRelationships {
	this := SupplyChainRelationships{}
	return &this
}

// NewSupplyChainRelationshipsWithDefaults instantiates a new SupplyChainRelationships object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSupplyChainRelationshipsWithDefaults() *SupplyChainRelationships {
	this := SupplyChainRelationships{}
	return &this
}

// GetSymbol returns the Symbol field value if set, zero value otherwise.
func (o *SupplyChainRelationships) GetSymbol() string {
	if o == nil || o.Symbol == nil {
		var ret string
		return ret
	}
	return *o.Symbol
}

// GetSymbolOk returns a tuple with the Symbol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SupplyChainRelationships) GetSymbolOk() (*string, bool) {
	if o == nil || o.Symbol == nil {
		return nil, false
	}
	return o.Symbol, true
}

// HasSymbol returns a boolean if a field has been set.
func (o *SupplyChainRelationships) HasSymbol() bool {
	if o != nil && o.Symbol != nil {
		return true
	}

	return false
}

// SetSymbol gets a reference to the given string and assigns it to the Symbol field.
func (o *SupplyChainRelationships) SetSymbol(v string) {
	o.Symbol = &v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *SupplyChainRelationships) GetData() []KeyCustomersSuppliers {
	if o == nil || o.Data == nil {
		var ret []KeyCustomersSuppliers
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SupplyChainRelationships) GetDataOk() (*[]KeyCustomersSuppliers, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *SupplyChainRelationships) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given []KeyCustomersSuppliers and assigns it to the Data field.
func (o *SupplyChainRelationships) SetData(v []KeyCustomersSuppliers) {
	o.Data = &v
}

func (o SupplyChainRelationships) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Symbol != nil {
		toSerialize["symbol"] = o.Symbol
	}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullableSupplyChainRelationships struct {
	value *SupplyChainRelationships
	isSet bool
}

func (v NullableSupplyChainRelationships) Get() *SupplyChainRelationships {
	return v.value
}

func (v *NullableSupplyChainRelationships) Set(val *SupplyChainRelationships) {
	v.value = val
	v.isSet = true
}

func (v NullableSupplyChainRelationships) IsSet() bool {
	return v.isSet
}

func (v *NullableSupplyChainRelationships) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSupplyChainRelationships(val *SupplyChainRelationships) *NullableSupplyChainRelationships {
	return &NullableSupplyChainRelationships{value: val, isSet: true}
}

func (v NullableSupplyChainRelationships) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSupplyChainRelationships) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


