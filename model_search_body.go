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

// SearchBody struct for SearchBody
type SearchBody struct {
	// Search query
	Query string `json:"query"`
	// List of isin to search, comma separated (Max: 50).
	Isins *string `json:"isins,omitempty"`
	// List of cusip to search, comma separated (Max: 50).
	Cusips *string `json:"cusips,omitempty"`
	// List of SEC Center Index Key to search, comma separated (Max: 50).
	Ciks *string `json:"ciks,omitempty"`
	// List of SEDAR issuer number to search, comma separated (Max: 50).
	SedarIds *string `json:"sedarIds,omitempty"`
	// List of Companies House number to search, comma separated (Max: 50).
	ChIds *string `json:"chIds,omitempty"`
	// List of symbols to search, comma separated (Max: 50).
	Symbols *string `json:"symbols,omitempty"`
	// List of sedols to search, comma separated (Max: 50).
	Sedols *string `json:"sedols,omitempty"`
	// List of sources to search, comma separated (Max: 50). Look at <code>/filter</code> endpoint to see all available values.
	Sources *string `json:"sources,omitempty"`
	// List of forms to search, comma separated (Max: 50). Look at <code>/filter</code> endpoint to see all available values.
	Forms *string `json:"forms,omitempty"`
	// List of gics to search, comma separated (Max: 50). Look at <code>/filter</code> endpoint to see all available values.
	Gics *string `json:"gics,omitempty"`
	// List of sources to search, comma separated (Max: 50). Look at <code>/filter</code> endpoint to see all available values.
	Naics *string `json:"naics,omitempty"`
	// List of exhibits to search, comma separated (Max: 50). Look at <code>/filter</code> endpoint to see all available values.
	Exhibits *string `json:"exhibits,omitempty"`
	// List of exchanges to search, comma separated (Max: 50). Look at <code>/filter</code> endpoint to see all available values.
	Exchanges *string `json:"exchanges,omitempty"`
	// List of sources to search, comma separated (Max: 50). Look at <code>/filter</code> endpoint to see all available values.
	Countries *string `json:"countries,omitempty"`
	// List of SEC's exchanges act to search, comma separated. Look at <code>/filter</code> endpoint to see all available values.
	Acts *string `json:"acts,omitempty"`
	// List of market capitalization to search, comma separated. Look at <code>/filter</code> endpoint to see all available values.
	Caps *string `json:"caps,omitempty"`
	// Search from date in format: YYYY-MM-DD, default from the last 2 years
	FromDate *string `json:"fromDate,omitempty"`
	// Search to date in format: YYYY-MM-DD, default to today
	ToDate *string `json:"toDate,omitempty"`
	// Use for pagination, default to page 1
	Page *string `json:"page,omitempty"`
	// Sort result by, default: sortMostRecent. Look at <code>/filter</code> endpoint to see all available values.
	Sort *string `json:"sort,omitempty"`
	// Enable highlight in returned filings. If enabled, only return 10 results each time
	Highlighted *bool `json:"highlighted,omitempty"`
}

// NewSearchBody instantiates a new SearchBody object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSearchBody(query string) *SearchBody {
	this := SearchBody{}
	this.Query = query
	return &this
}

// NewSearchBodyWithDefaults instantiates a new SearchBody object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSearchBodyWithDefaults() *SearchBody {
	this := SearchBody{}
	return &this
}

// GetQuery returns the Query field value
func (o *SearchBody) GetQuery() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Query
}

// GetQueryOk returns a tuple with the Query field value
// and a boolean to check if the value has been set.
func (o *SearchBody) GetQueryOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Query, true
}

// SetQuery sets field value
func (o *SearchBody) SetQuery(v string) {
	o.Query = v
}

// GetIsins returns the Isins field value if set, zero value otherwise.
func (o *SearchBody) GetIsins() string {
	if o == nil || o.Isins == nil {
		var ret string
		return ret
	}
	return *o.Isins
}

// GetIsinsOk returns a tuple with the Isins field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchBody) GetIsinsOk() (*string, bool) {
	if o == nil || o.Isins == nil {
		return nil, false
	}
	return o.Isins, true
}

// HasIsins returns a boolean if a field has been set.
func (o *SearchBody) HasIsins() bool {
	if o != nil && o.Isins != nil {
		return true
	}

	return false
}

// SetIsins gets a reference to the given string and assigns it to the Isins field.
func (o *SearchBody) SetIsins(v string) {
	o.Isins = &v
}

// GetCusips returns the Cusips field value if set, zero value otherwise.
func (o *SearchBody) GetCusips() string {
	if o == nil || o.Cusips == nil {
		var ret string
		return ret
	}
	return *o.Cusips
}

// GetCusipsOk returns a tuple with the Cusips field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchBody) GetCusipsOk() (*string, bool) {
	if o == nil || o.Cusips == nil {
		return nil, false
	}
	return o.Cusips, true
}

// HasCusips returns a boolean if a field has been set.
func (o *SearchBody) HasCusips() bool {
	if o != nil && o.Cusips != nil {
		return true
	}

	return false
}

// SetCusips gets a reference to the given string and assigns it to the Cusips field.
func (o *SearchBody) SetCusips(v string) {
	o.Cusips = &v
}

// GetCiks returns the Ciks field value if set, zero value otherwise.
func (o *SearchBody) GetCiks() string {
	if o == nil || o.Ciks == nil {
		var ret string
		return ret
	}
	return *o.Ciks
}

// GetCiksOk returns a tuple with the Ciks field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchBody) GetCiksOk() (*string, bool) {
	if o == nil || o.Ciks == nil {
		return nil, false
	}
	return o.Ciks, true
}

// HasCiks returns a boolean if a field has been set.
func (o *SearchBody) HasCiks() bool {
	if o != nil && o.Ciks != nil {
		return true
	}

	return false
}

// SetCiks gets a reference to the given string and assigns it to the Ciks field.
func (o *SearchBody) SetCiks(v string) {
	o.Ciks = &v
}

// GetSedarIds returns the SedarIds field value if set, zero value otherwise.
func (o *SearchBody) GetSedarIds() string {
	if o == nil || o.SedarIds == nil {
		var ret string
		return ret
	}
	return *o.SedarIds
}

// GetSedarIdsOk returns a tuple with the SedarIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchBody) GetSedarIdsOk() (*string, bool) {
	if o == nil || o.SedarIds == nil {
		return nil, false
	}
	return o.SedarIds, true
}

// HasSedarIds returns a boolean if a field has been set.
func (o *SearchBody) HasSedarIds() bool {
	if o != nil && o.SedarIds != nil {
		return true
	}

	return false
}

// SetSedarIds gets a reference to the given string and assigns it to the SedarIds field.
func (o *SearchBody) SetSedarIds(v string) {
	o.SedarIds = &v
}

// GetChIds returns the ChIds field value if set, zero value otherwise.
func (o *SearchBody) GetChIds() string {
	if o == nil || o.ChIds == nil {
		var ret string
		return ret
	}
	return *o.ChIds
}

// GetChIdsOk returns a tuple with the ChIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchBody) GetChIdsOk() (*string, bool) {
	if o == nil || o.ChIds == nil {
		return nil, false
	}
	return o.ChIds, true
}

// HasChIds returns a boolean if a field has been set.
func (o *SearchBody) HasChIds() bool {
	if o != nil && o.ChIds != nil {
		return true
	}

	return false
}

// SetChIds gets a reference to the given string and assigns it to the ChIds field.
func (o *SearchBody) SetChIds(v string) {
	o.ChIds = &v
}

// GetSymbols returns the Symbols field value if set, zero value otherwise.
func (o *SearchBody) GetSymbols() string {
	if o == nil || o.Symbols == nil {
		var ret string
		return ret
	}
	return *o.Symbols
}

// GetSymbolsOk returns a tuple with the Symbols field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchBody) GetSymbolsOk() (*string, bool) {
	if o == nil || o.Symbols == nil {
		return nil, false
	}
	return o.Symbols, true
}

// HasSymbols returns a boolean if a field has been set.
func (o *SearchBody) HasSymbols() bool {
	if o != nil && o.Symbols != nil {
		return true
	}

	return false
}

// SetSymbols gets a reference to the given string and assigns it to the Symbols field.
func (o *SearchBody) SetSymbols(v string) {
	o.Symbols = &v
}

// GetSedols returns the Sedols field value if set, zero value otherwise.
func (o *SearchBody) GetSedols() string {
	if o == nil || o.Sedols == nil {
		var ret string
		return ret
	}
	return *o.Sedols
}

// GetSedolsOk returns a tuple with the Sedols field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchBody) GetSedolsOk() (*string, bool) {
	if o == nil || o.Sedols == nil {
		return nil, false
	}
	return o.Sedols, true
}

// HasSedols returns a boolean if a field has been set.
func (o *SearchBody) HasSedols() bool {
	if o != nil && o.Sedols != nil {
		return true
	}

	return false
}

// SetSedols gets a reference to the given string and assigns it to the Sedols field.
func (o *SearchBody) SetSedols(v string) {
	o.Sedols = &v
}

// GetSources returns the Sources field value if set, zero value otherwise.
func (o *SearchBody) GetSources() string {
	if o == nil || o.Sources == nil {
		var ret string
		return ret
	}
	return *o.Sources
}

// GetSourcesOk returns a tuple with the Sources field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchBody) GetSourcesOk() (*string, bool) {
	if o == nil || o.Sources == nil {
		return nil, false
	}
	return o.Sources, true
}

// HasSources returns a boolean if a field has been set.
func (o *SearchBody) HasSources() bool {
	if o != nil && o.Sources != nil {
		return true
	}

	return false
}

// SetSources gets a reference to the given string and assigns it to the Sources field.
func (o *SearchBody) SetSources(v string) {
	o.Sources = &v
}

// GetForms returns the Forms field value if set, zero value otherwise.
func (o *SearchBody) GetForms() string {
	if o == nil || o.Forms == nil {
		var ret string
		return ret
	}
	return *o.Forms
}

// GetFormsOk returns a tuple with the Forms field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchBody) GetFormsOk() (*string, bool) {
	if o == nil || o.Forms == nil {
		return nil, false
	}
	return o.Forms, true
}

// HasForms returns a boolean if a field has been set.
func (o *SearchBody) HasForms() bool {
	if o != nil && o.Forms != nil {
		return true
	}

	return false
}

// SetForms gets a reference to the given string and assigns it to the Forms field.
func (o *SearchBody) SetForms(v string) {
	o.Forms = &v
}

// GetGics returns the Gics field value if set, zero value otherwise.
func (o *SearchBody) GetGics() string {
	if o == nil || o.Gics == nil {
		var ret string
		return ret
	}
	return *o.Gics
}

// GetGicsOk returns a tuple with the Gics field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchBody) GetGicsOk() (*string, bool) {
	if o == nil || o.Gics == nil {
		return nil, false
	}
	return o.Gics, true
}

// HasGics returns a boolean if a field has been set.
func (o *SearchBody) HasGics() bool {
	if o != nil && o.Gics != nil {
		return true
	}

	return false
}

// SetGics gets a reference to the given string and assigns it to the Gics field.
func (o *SearchBody) SetGics(v string) {
	o.Gics = &v
}

// GetNaics returns the Naics field value if set, zero value otherwise.
func (o *SearchBody) GetNaics() string {
	if o == nil || o.Naics == nil {
		var ret string
		return ret
	}
	return *o.Naics
}

// GetNaicsOk returns a tuple with the Naics field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchBody) GetNaicsOk() (*string, bool) {
	if o == nil || o.Naics == nil {
		return nil, false
	}
	return o.Naics, true
}

// HasNaics returns a boolean if a field has been set.
func (o *SearchBody) HasNaics() bool {
	if o != nil && o.Naics != nil {
		return true
	}

	return false
}

// SetNaics gets a reference to the given string and assigns it to the Naics field.
func (o *SearchBody) SetNaics(v string) {
	o.Naics = &v
}

// GetExhibits returns the Exhibits field value if set, zero value otherwise.
func (o *SearchBody) GetExhibits() string {
	if o == nil || o.Exhibits == nil {
		var ret string
		return ret
	}
	return *o.Exhibits
}

// GetExhibitsOk returns a tuple with the Exhibits field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchBody) GetExhibitsOk() (*string, bool) {
	if o == nil || o.Exhibits == nil {
		return nil, false
	}
	return o.Exhibits, true
}

// HasExhibits returns a boolean if a field has been set.
func (o *SearchBody) HasExhibits() bool {
	if o != nil && o.Exhibits != nil {
		return true
	}

	return false
}

// SetExhibits gets a reference to the given string and assigns it to the Exhibits field.
func (o *SearchBody) SetExhibits(v string) {
	o.Exhibits = &v
}

// GetExchanges returns the Exchanges field value if set, zero value otherwise.
func (o *SearchBody) GetExchanges() string {
	if o == nil || o.Exchanges == nil {
		var ret string
		return ret
	}
	return *o.Exchanges
}

// GetExchangesOk returns a tuple with the Exchanges field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchBody) GetExchangesOk() (*string, bool) {
	if o == nil || o.Exchanges == nil {
		return nil, false
	}
	return o.Exchanges, true
}

// HasExchanges returns a boolean if a field has been set.
func (o *SearchBody) HasExchanges() bool {
	if o != nil && o.Exchanges != nil {
		return true
	}

	return false
}

// SetExchanges gets a reference to the given string and assigns it to the Exchanges field.
func (o *SearchBody) SetExchanges(v string) {
	o.Exchanges = &v
}

// GetCountries returns the Countries field value if set, zero value otherwise.
func (o *SearchBody) GetCountries() string {
	if o == nil || o.Countries == nil {
		var ret string
		return ret
	}
	return *o.Countries
}

// GetCountriesOk returns a tuple with the Countries field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchBody) GetCountriesOk() (*string, bool) {
	if o == nil || o.Countries == nil {
		return nil, false
	}
	return o.Countries, true
}

// HasCountries returns a boolean if a field has been set.
func (o *SearchBody) HasCountries() bool {
	if o != nil && o.Countries != nil {
		return true
	}

	return false
}

// SetCountries gets a reference to the given string and assigns it to the Countries field.
func (o *SearchBody) SetCountries(v string) {
	o.Countries = &v
}

// GetActs returns the Acts field value if set, zero value otherwise.
func (o *SearchBody) GetActs() string {
	if o == nil || o.Acts == nil {
		var ret string
		return ret
	}
	return *o.Acts
}

// GetActsOk returns a tuple with the Acts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchBody) GetActsOk() (*string, bool) {
	if o == nil || o.Acts == nil {
		return nil, false
	}
	return o.Acts, true
}

// HasActs returns a boolean if a field has been set.
func (o *SearchBody) HasActs() bool {
	if o != nil && o.Acts != nil {
		return true
	}

	return false
}

// SetActs gets a reference to the given string and assigns it to the Acts field.
func (o *SearchBody) SetActs(v string) {
	o.Acts = &v
}

// GetCaps returns the Caps field value if set, zero value otherwise.
func (o *SearchBody) GetCaps() string {
	if o == nil || o.Caps == nil {
		var ret string
		return ret
	}
	return *o.Caps
}

// GetCapsOk returns a tuple with the Caps field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchBody) GetCapsOk() (*string, bool) {
	if o == nil || o.Caps == nil {
		return nil, false
	}
	return o.Caps, true
}

// HasCaps returns a boolean if a field has been set.
func (o *SearchBody) HasCaps() bool {
	if o != nil && o.Caps != nil {
		return true
	}

	return false
}

// SetCaps gets a reference to the given string and assigns it to the Caps field.
func (o *SearchBody) SetCaps(v string) {
	o.Caps = &v
}

// GetFromDate returns the FromDate field value if set, zero value otherwise.
func (o *SearchBody) GetFromDate() string {
	if o == nil || o.FromDate == nil {
		var ret string
		return ret
	}
	return *o.FromDate
}

// GetFromDateOk returns a tuple with the FromDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchBody) GetFromDateOk() (*string, bool) {
	if o == nil || o.FromDate == nil {
		return nil, false
	}
	return o.FromDate, true
}

// HasFromDate returns a boolean if a field has been set.
func (o *SearchBody) HasFromDate() bool {
	if o != nil && o.FromDate != nil {
		return true
	}

	return false
}

// SetFromDate gets a reference to the given string and assigns it to the FromDate field.
func (o *SearchBody) SetFromDate(v string) {
	o.FromDate = &v
}

// GetToDate returns the ToDate field value if set, zero value otherwise.
func (o *SearchBody) GetToDate() string {
	if o == nil || o.ToDate == nil {
		var ret string
		return ret
	}
	return *o.ToDate
}

// GetToDateOk returns a tuple with the ToDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchBody) GetToDateOk() (*string, bool) {
	if o == nil || o.ToDate == nil {
		return nil, false
	}
	return o.ToDate, true
}

// HasToDate returns a boolean if a field has been set.
func (o *SearchBody) HasToDate() bool {
	if o != nil && o.ToDate != nil {
		return true
	}

	return false
}

// SetToDate gets a reference to the given string and assigns it to the ToDate field.
func (o *SearchBody) SetToDate(v string) {
	o.ToDate = &v
}

// GetPage returns the Page field value if set, zero value otherwise.
func (o *SearchBody) GetPage() string {
	if o == nil || o.Page == nil {
		var ret string
		return ret
	}
	return *o.Page
}

// GetPageOk returns a tuple with the Page field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchBody) GetPageOk() (*string, bool) {
	if o == nil || o.Page == nil {
		return nil, false
	}
	return o.Page, true
}

// HasPage returns a boolean if a field has been set.
func (o *SearchBody) HasPage() bool {
	if o != nil && o.Page != nil {
		return true
	}

	return false
}

// SetPage gets a reference to the given string and assigns it to the Page field.
func (o *SearchBody) SetPage(v string) {
	o.Page = &v
}

// GetSort returns the Sort field value if set, zero value otherwise.
func (o *SearchBody) GetSort() string {
	if o == nil || o.Sort == nil {
		var ret string
		return ret
	}
	return *o.Sort
}

// GetSortOk returns a tuple with the Sort field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchBody) GetSortOk() (*string, bool) {
	if o == nil || o.Sort == nil {
		return nil, false
	}
	return o.Sort, true
}

// HasSort returns a boolean if a field has been set.
func (o *SearchBody) HasSort() bool {
	if o != nil && o.Sort != nil {
		return true
	}

	return false
}

// SetSort gets a reference to the given string and assigns it to the Sort field.
func (o *SearchBody) SetSort(v string) {
	o.Sort = &v
}

// GetHighlighted returns the Highlighted field value if set, zero value otherwise.
func (o *SearchBody) GetHighlighted() bool {
	if o == nil || o.Highlighted == nil {
		var ret bool
		return ret
	}
	return *o.Highlighted
}

// GetHighlightedOk returns a tuple with the Highlighted field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchBody) GetHighlightedOk() (*bool, bool) {
	if o == nil || o.Highlighted == nil {
		return nil, false
	}
	return o.Highlighted, true
}

// HasHighlighted returns a boolean if a field has been set.
func (o *SearchBody) HasHighlighted() bool {
	if o != nil && o.Highlighted != nil {
		return true
	}

	return false
}

// SetHighlighted gets a reference to the given bool and assigns it to the Highlighted field.
func (o *SearchBody) SetHighlighted(v bool) {
	o.Highlighted = &v
}

func (o SearchBody) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["query"] = o.Query
	}
	if o.Isins != nil {
		toSerialize["isins"] = o.Isins
	}
	if o.Cusips != nil {
		toSerialize["cusips"] = o.Cusips
	}
	if o.Ciks != nil {
		toSerialize["ciks"] = o.Ciks
	}
	if o.SedarIds != nil {
		toSerialize["sedarIds"] = o.SedarIds
	}
	if o.ChIds != nil {
		toSerialize["chIds"] = o.ChIds
	}
	if o.Symbols != nil {
		toSerialize["symbols"] = o.Symbols
	}
	if o.Sedols != nil {
		toSerialize["sedols"] = o.Sedols
	}
	if o.Sources != nil {
		toSerialize["sources"] = o.Sources
	}
	if o.Forms != nil {
		toSerialize["forms"] = o.Forms
	}
	if o.Gics != nil {
		toSerialize["gics"] = o.Gics
	}
	if o.Naics != nil {
		toSerialize["naics"] = o.Naics
	}
	if o.Exhibits != nil {
		toSerialize["exhibits"] = o.Exhibits
	}
	if o.Exchanges != nil {
		toSerialize["exchanges"] = o.Exchanges
	}
	if o.Countries != nil {
		toSerialize["countries"] = o.Countries
	}
	if o.Acts != nil {
		toSerialize["acts"] = o.Acts
	}
	if o.Caps != nil {
		toSerialize["caps"] = o.Caps
	}
	if o.FromDate != nil {
		toSerialize["fromDate"] = o.FromDate
	}
	if o.ToDate != nil {
		toSerialize["toDate"] = o.ToDate
	}
	if o.Page != nil {
		toSerialize["page"] = o.Page
	}
	if o.Sort != nil {
		toSerialize["sort"] = o.Sort
	}
	if o.Highlighted != nil {
		toSerialize["highlighted"] = o.Highlighted
	}
	return json.Marshal(toSerialize)
}

type NullableSearchBody struct {
	value *SearchBody
	isSet bool
}

func (v NullableSearchBody) Get() *SearchBody {
	return v.value
}

func (v *NullableSearchBody) Set(val *SearchBody) {
	v.value = val
	v.isSet = true
}

func (v NullableSearchBody) IsSet() bool {
	return v.isSet
}

func (v *NullableSearchBody) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSearchBody(val *SearchBody) *NullableSearchBody {
	return &NullableSearchBody{value: val, isSet: true}
}

func (v NullableSearchBody) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSearchBody) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


