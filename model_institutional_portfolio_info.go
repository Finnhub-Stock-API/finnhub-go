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

// InstitutionalPortfolioInfo struct for InstitutionalPortfolioInfo
type InstitutionalPortfolioInfo struct {
	// Symbol.
	Symbol *string `json:"symbol,omitempty"`
	// CUSIP.
	Cusip *string `json:"cusip,omitempty"`
	// Position's name.
	Name *string `json:"name,omitempty"`
	// <code>put</code> or <code>call</code> for options.
	PutCall *string `json:"putCall,omitempty"`
	// Number of shares change.
	Change *float32 `json:"change,omitempty"`
	// Number of shares with no voting rights.
	NoVoting *float32 `json:"noVoting,omitempty"`
	// Percentage of portfolio.
	Percentage *float32 `json:"percentage,omitempty"`
	// Number of shares.
	Share *float32 `json:"share,omitempty"`
	// Number of shares with shared voting rights.
	SharedVoting *float32 `json:"sharedVoting,omitempty"`
	// Number of shares with sole voting rights.
	SoleVoting *float32 `json:"soleVoting,omitempty"`
	// Position value.
	Value *float32 `json:"value,omitempty"`
}

// NewInstitutionalPortfolioInfo instantiates a new InstitutionalPortfolioInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInstitutionalPortfolioInfo() *InstitutionalPortfolioInfo {
	this := InstitutionalPortfolioInfo{}
	return &this
}

// NewInstitutionalPortfolioInfoWithDefaults instantiates a new InstitutionalPortfolioInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInstitutionalPortfolioInfoWithDefaults() *InstitutionalPortfolioInfo {
	this := InstitutionalPortfolioInfo{}
	return &this
}

// GetSymbol returns the Symbol field value if set, zero value otherwise.
func (o *InstitutionalPortfolioInfo) GetSymbol() string {
	if o == nil || o.Symbol == nil {
		var ret string
		return ret
	}
	return *o.Symbol
}

// GetSymbolOk returns a tuple with the Symbol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InstitutionalPortfolioInfo) GetSymbolOk() (*string, bool) {
	if o == nil || o.Symbol == nil {
		return nil, false
	}
	return o.Symbol, true
}

// HasSymbol returns a boolean if a field has been set.
func (o *InstitutionalPortfolioInfo) HasSymbol() bool {
	if o != nil && o.Symbol != nil {
		return true
	}

	return false
}

// SetSymbol gets a reference to the given string and assigns it to the Symbol field.
func (o *InstitutionalPortfolioInfo) SetSymbol(v string) {
	o.Symbol = &v
}

// GetCusip returns the Cusip field value if set, zero value otherwise.
func (o *InstitutionalPortfolioInfo) GetCusip() string {
	if o == nil || o.Cusip == nil {
		var ret string
		return ret
	}
	return *o.Cusip
}

// GetCusipOk returns a tuple with the Cusip field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InstitutionalPortfolioInfo) GetCusipOk() (*string, bool) {
	if o == nil || o.Cusip == nil {
		return nil, false
	}
	return o.Cusip, true
}

// HasCusip returns a boolean if a field has been set.
func (o *InstitutionalPortfolioInfo) HasCusip() bool {
	if o != nil && o.Cusip != nil {
		return true
	}

	return false
}

// SetCusip gets a reference to the given string and assigns it to the Cusip field.
func (o *InstitutionalPortfolioInfo) SetCusip(v string) {
	o.Cusip = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *InstitutionalPortfolioInfo) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InstitutionalPortfolioInfo) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *InstitutionalPortfolioInfo) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *InstitutionalPortfolioInfo) SetName(v string) {
	o.Name = &v
}

// GetPutCall returns the PutCall field value if set, zero value otherwise.
func (o *InstitutionalPortfolioInfo) GetPutCall() string {
	if o == nil || o.PutCall == nil {
		var ret string
		return ret
	}
	return *o.PutCall
}

// GetPutCallOk returns a tuple with the PutCall field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InstitutionalPortfolioInfo) GetPutCallOk() (*string, bool) {
	if o == nil || o.PutCall == nil {
		return nil, false
	}
	return o.PutCall, true
}

// HasPutCall returns a boolean if a field has been set.
func (o *InstitutionalPortfolioInfo) HasPutCall() bool {
	if o != nil && o.PutCall != nil {
		return true
	}

	return false
}

// SetPutCall gets a reference to the given string and assigns it to the PutCall field.
func (o *InstitutionalPortfolioInfo) SetPutCall(v string) {
	o.PutCall = &v
}

// GetChange returns the Change field value if set, zero value otherwise.
func (o *InstitutionalPortfolioInfo) GetChange() float32 {
	if o == nil || o.Change == nil {
		var ret float32
		return ret
	}
	return *o.Change
}

// GetChangeOk returns a tuple with the Change field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InstitutionalPortfolioInfo) GetChangeOk() (*float32, bool) {
	if o == nil || o.Change == nil {
		return nil, false
	}
	return o.Change, true
}

// HasChange returns a boolean if a field has been set.
func (o *InstitutionalPortfolioInfo) HasChange() bool {
	if o != nil && o.Change != nil {
		return true
	}

	return false
}

// SetChange gets a reference to the given float32 and assigns it to the Change field.
func (o *InstitutionalPortfolioInfo) SetChange(v float32) {
	o.Change = &v
}

// GetNoVoting returns the NoVoting field value if set, zero value otherwise.
func (o *InstitutionalPortfolioInfo) GetNoVoting() float32 {
	if o == nil || o.NoVoting == nil {
		var ret float32
		return ret
	}
	return *o.NoVoting
}

// GetNoVotingOk returns a tuple with the NoVoting field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InstitutionalPortfolioInfo) GetNoVotingOk() (*float32, bool) {
	if o == nil || o.NoVoting == nil {
		return nil, false
	}
	return o.NoVoting, true
}

// HasNoVoting returns a boolean if a field has been set.
func (o *InstitutionalPortfolioInfo) HasNoVoting() bool {
	if o != nil && o.NoVoting != nil {
		return true
	}

	return false
}

// SetNoVoting gets a reference to the given float32 and assigns it to the NoVoting field.
func (o *InstitutionalPortfolioInfo) SetNoVoting(v float32) {
	o.NoVoting = &v
}

// GetPercentage returns the Percentage field value if set, zero value otherwise.
func (o *InstitutionalPortfolioInfo) GetPercentage() float32 {
	if o == nil || o.Percentage == nil {
		var ret float32
		return ret
	}
	return *o.Percentage
}

// GetPercentageOk returns a tuple with the Percentage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InstitutionalPortfolioInfo) GetPercentageOk() (*float32, bool) {
	if o == nil || o.Percentage == nil {
		return nil, false
	}
	return o.Percentage, true
}

// HasPercentage returns a boolean if a field has been set.
func (o *InstitutionalPortfolioInfo) HasPercentage() bool {
	if o != nil && o.Percentage != nil {
		return true
	}

	return false
}

// SetPercentage gets a reference to the given float32 and assigns it to the Percentage field.
func (o *InstitutionalPortfolioInfo) SetPercentage(v float32) {
	o.Percentage = &v
}

// GetShare returns the Share field value if set, zero value otherwise.
func (o *InstitutionalPortfolioInfo) GetShare() float32 {
	if o == nil || o.Share == nil {
		var ret float32
		return ret
	}
	return *o.Share
}

// GetShareOk returns a tuple with the Share field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InstitutionalPortfolioInfo) GetShareOk() (*float32, bool) {
	if o == nil || o.Share == nil {
		return nil, false
	}
	return o.Share, true
}

// HasShare returns a boolean if a field has been set.
func (o *InstitutionalPortfolioInfo) HasShare() bool {
	if o != nil && o.Share != nil {
		return true
	}

	return false
}

// SetShare gets a reference to the given float32 and assigns it to the Share field.
func (o *InstitutionalPortfolioInfo) SetShare(v float32) {
	o.Share = &v
}

// GetSharedVoting returns the SharedVoting field value if set, zero value otherwise.
func (o *InstitutionalPortfolioInfo) GetSharedVoting() float32 {
	if o == nil || o.SharedVoting == nil {
		var ret float32
		return ret
	}
	return *o.SharedVoting
}

// GetSharedVotingOk returns a tuple with the SharedVoting field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InstitutionalPortfolioInfo) GetSharedVotingOk() (*float32, bool) {
	if o == nil || o.SharedVoting == nil {
		return nil, false
	}
	return o.SharedVoting, true
}

// HasSharedVoting returns a boolean if a field has been set.
func (o *InstitutionalPortfolioInfo) HasSharedVoting() bool {
	if o != nil && o.SharedVoting != nil {
		return true
	}

	return false
}

// SetSharedVoting gets a reference to the given float32 and assigns it to the SharedVoting field.
func (o *InstitutionalPortfolioInfo) SetSharedVoting(v float32) {
	o.SharedVoting = &v
}

// GetSoleVoting returns the SoleVoting field value if set, zero value otherwise.
func (o *InstitutionalPortfolioInfo) GetSoleVoting() float32 {
	if o == nil || o.SoleVoting == nil {
		var ret float32
		return ret
	}
	return *o.SoleVoting
}

// GetSoleVotingOk returns a tuple with the SoleVoting field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InstitutionalPortfolioInfo) GetSoleVotingOk() (*float32, bool) {
	if o == nil || o.SoleVoting == nil {
		return nil, false
	}
	return o.SoleVoting, true
}

// HasSoleVoting returns a boolean if a field has been set.
func (o *InstitutionalPortfolioInfo) HasSoleVoting() bool {
	if o != nil && o.SoleVoting != nil {
		return true
	}

	return false
}

// SetSoleVoting gets a reference to the given float32 and assigns it to the SoleVoting field.
func (o *InstitutionalPortfolioInfo) SetSoleVoting(v float32) {
	o.SoleVoting = &v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *InstitutionalPortfolioInfo) GetValue() float32 {
	if o == nil || o.Value == nil {
		var ret float32
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InstitutionalPortfolioInfo) GetValueOk() (*float32, bool) {
	if o == nil || o.Value == nil {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *InstitutionalPortfolioInfo) HasValue() bool {
	if o != nil && o.Value != nil {
		return true
	}

	return false
}

// SetValue gets a reference to the given float32 and assigns it to the Value field.
func (o *InstitutionalPortfolioInfo) SetValue(v float32) {
	o.Value = &v
}

func (o InstitutionalPortfolioInfo) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Symbol != nil {
		toSerialize["symbol"] = o.Symbol
	}
	if o.Cusip != nil {
		toSerialize["cusip"] = o.Cusip
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.PutCall != nil {
		toSerialize["putCall"] = o.PutCall
	}
	if o.Change != nil {
		toSerialize["change"] = o.Change
	}
	if o.NoVoting != nil {
		toSerialize["noVoting"] = o.NoVoting
	}
	if o.Percentage != nil {
		toSerialize["percentage"] = o.Percentage
	}
	if o.Share != nil {
		toSerialize["share"] = o.Share
	}
	if o.SharedVoting != nil {
		toSerialize["sharedVoting"] = o.SharedVoting
	}
	if o.SoleVoting != nil {
		toSerialize["soleVoting"] = o.SoleVoting
	}
	if o.Value != nil {
		toSerialize["value"] = o.Value
	}
	return json.Marshal(toSerialize)
}

type NullableInstitutionalPortfolioInfo struct {
	value *InstitutionalPortfolioInfo
	isSet bool
}

func (v NullableInstitutionalPortfolioInfo) Get() *InstitutionalPortfolioInfo {
	return v.value
}

func (v *NullableInstitutionalPortfolioInfo) Set(val *InstitutionalPortfolioInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableInstitutionalPortfolioInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableInstitutionalPortfolioInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInstitutionalPortfolioInfo(val *InstitutionalPortfolioInfo) *NullableInstitutionalPortfolioInfo {
	return &NullableInstitutionalPortfolioInfo{value: val, isSet: true}
}

func (v NullableInstitutionalPortfolioInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInstitutionalPortfolioInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

