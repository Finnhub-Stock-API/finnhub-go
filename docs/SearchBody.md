# SearchBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Query** | **string** | Search query | 
**Isins** | Pointer to **string** | List of isin to search, comma separated (Max: 50). | [optional] 
**Cusips** | Pointer to **string** | List of cusip to search, comma separated (Max: 50). | [optional] 
**Ciks** | Pointer to **string** | List of SEC Center Index Key to search, comma separated (Max: 50). | [optional] 
**SedarIds** | Pointer to **string** | List of SEDAR issuer number to search, comma separated (Max: 50). | [optional] 
**ChIds** | Pointer to **string** | List of Companies House number to search, comma separated (Max: 50). | [optional] 
**Symbols** | Pointer to **string** | List of symbols to search, comma separated (Max: 50). | [optional] 
**Sedols** | Pointer to **string** | List of sedols to search, comma separated (Max: 50). | [optional] 
**Sources** | Pointer to **string** | List of sources to search, comma separated (Max: 50). Look at &lt;code&gt;/filter&lt;/code&gt; endpoint to see all available values. | [optional] 
**Forms** | Pointer to **string** | List of forms to search, comma separated (Max: 50). Look at &lt;code&gt;/filter&lt;/code&gt; endpoint to see all available values. | [optional] 
**Gics** | Pointer to **string** | List of gics to search, comma separated (Max: 50). Look at &lt;code&gt;/filter&lt;/code&gt; endpoint to see all available values. | [optional] 
**Naics** | Pointer to **string** | List of sources to search, comma separated (Max: 50). Look at &lt;code&gt;/filter&lt;/code&gt; endpoint to see all available values. | [optional] 
**Exhibits** | Pointer to **string** | List of exhibits to search, comma separated (Max: 50). Look at &lt;code&gt;/filter&lt;/code&gt; endpoint to see all available values. | [optional] 
**Exchanges** | Pointer to **string** | List of exchanges to search, comma separated (Max: 50). Look at &lt;code&gt;/filter&lt;/code&gt; endpoint to see all available values. | [optional] 
**Countries** | Pointer to **string** | List of sources to search, comma separated (Max: 50). Look at &lt;code&gt;/filter&lt;/code&gt; endpoint to see all available values. | [optional] 
**Acts** | Pointer to **string** | List of SEC&#39;s exchanges act to search, comma separated. Look at &lt;code&gt;/filter&lt;/code&gt; endpoint to see all available values. | [optional] 
**Caps** | Pointer to **string** | List of market capitalization to search, comma separated. Look at &lt;code&gt;/filter&lt;/code&gt; endpoint to see all available values. | [optional] 
**FromDate** | Pointer to **string** | Search from date in format: YYYY-MM-DD, default from the last 2 years | [optional] 
**ToDate** | Pointer to **string** | Search to date in format: YYYY-MM-DD, default to today | [optional] 
**Page** | Pointer to **string** | Use for pagination, default to page 1 | [optional] 
**Sort** | Pointer to **string** | Sort result by, default: sortMostRecent. Look at &lt;code&gt;/filter&lt;/code&gt; endpoint to see all available values. | [optional] 
**Highlighted** | Pointer to **bool** | Enable highlight in returned filings. If enabled, only return 10 results each time | [optional] 

## Methods

### NewSearchBody

`func NewSearchBody(query string, ) *SearchBody`

NewSearchBody instantiates a new SearchBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSearchBodyWithDefaults

`func NewSearchBodyWithDefaults() *SearchBody`

NewSearchBodyWithDefaults instantiates a new SearchBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetQuery

`func (o *SearchBody) GetQuery() string`

GetQuery returns the Query field if non-nil, zero value otherwise.

### GetQueryOk

`func (o *SearchBody) GetQueryOk() (*string, bool)`

GetQueryOk returns a tuple with the Query field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuery

`func (o *SearchBody) SetQuery(v string)`

SetQuery sets Query field to given value.


### GetIsins

`func (o *SearchBody) GetIsins() string`

GetIsins returns the Isins field if non-nil, zero value otherwise.

### GetIsinsOk

`func (o *SearchBody) GetIsinsOk() (*string, bool)`

GetIsinsOk returns a tuple with the Isins field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsins

`func (o *SearchBody) SetIsins(v string)`

SetIsins sets Isins field to given value.

### HasIsins

`func (o *SearchBody) HasIsins() bool`

HasIsins returns a boolean if a field has been set.

### GetCusips

`func (o *SearchBody) GetCusips() string`

GetCusips returns the Cusips field if non-nil, zero value otherwise.

### GetCusipsOk

`func (o *SearchBody) GetCusipsOk() (*string, bool)`

GetCusipsOk returns a tuple with the Cusips field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCusips

`func (o *SearchBody) SetCusips(v string)`

SetCusips sets Cusips field to given value.

### HasCusips

`func (o *SearchBody) HasCusips() bool`

HasCusips returns a boolean if a field has been set.

### GetCiks

`func (o *SearchBody) GetCiks() string`

GetCiks returns the Ciks field if non-nil, zero value otherwise.

### GetCiksOk

`func (o *SearchBody) GetCiksOk() (*string, bool)`

GetCiksOk returns a tuple with the Ciks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCiks

`func (o *SearchBody) SetCiks(v string)`

SetCiks sets Ciks field to given value.

### HasCiks

`func (o *SearchBody) HasCiks() bool`

HasCiks returns a boolean if a field has been set.

### GetSedarIds

`func (o *SearchBody) GetSedarIds() string`

GetSedarIds returns the SedarIds field if non-nil, zero value otherwise.

### GetSedarIdsOk

`func (o *SearchBody) GetSedarIdsOk() (*string, bool)`

GetSedarIdsOk returns a tuple with the SedarIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSedarIds

`func (o *SearchBody) SetSedarIds(v string)`

SetSedarIds sets SedarIds field to given value.

### HasSedarIds

`func (o *SearchBody) HasSedarIds() bool`

HasSedarIds returns a boolean if a field has been set.

### GetChIds

`func (o *SearchBody) GetChIds() string`

GetChIds returns the ChIds field if non-nil, zero value otherwise.

### GetChIdsOk

`func (o *SearchBody) GetChIdsOk() (*string, bool)`

GetChIdsOk returns a tuple with the ChIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChIds

`func (o *SearchBody) SetChIds(v string)`

SetChIds sets ChIds field to given value.

### HasChIds

`func (o *SearchBody) HasChIds() bool`

HasChIds returns a boolean if a field has been set.

### GetSymbols

`func (o *SearchBody) GetSymbols() string`

GetSymbols returns the Symbols field if non-nil, zero value otherwise.

### GetSymbolsOk

`func (o *SearchBody) GetSymbolsOk() (*string, bool)`

GetSymbolsOk returns a tuple with the Symbols field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSymbols

`func (o *SearchBody) SetSymbols(v string)`

SetSymbols sets Symbols field to given value.

### HasSymbols

`func (o *SearchBody) HasSymbols() bool`

HasSymbols returns a boolean if a field has been set.

### GetSedols

`func (o *SearchBody) GetSedols() string`

GetSedols returns the Sedols field if non-nil, zero value otherwise.

### GetSedolsOk

`func (o *SearchBody) GetSedolsOk() (*string, bool)`

GetSedolsOk returns a tuple with the Sedols field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSedols

`func (o *SearchBody) SetSedols(v string)`

SetSedols sets Sedols field to given value.

### HasSedols

`func (o *SearchBody) HasSedols() bool`

HasSedols returns a boolean if a field has been set.

### GetSources

`func (o *SearchBody) GetSources() string`

GetSources returns the Sources field if non-nil, zero value otherwise.

### GetSourcesOk

`func (o *SearchBody) GetSourcesOk() (*string, bool)`

GetSourcesOk returns a tuple with the Sources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSources

`func (o *SearchBody) SetSources(v string)`

SetSources sets Sources field to given value.

### HasSources

`func (o *SearchBody) HasSources() bool`

HasSources returns a boolean if a field has been set.

### GetForms

`func (o *SearchBody) GetForms() string`

GetForms returns the Forms field if non-nil, zero value otherwise.

### GetFormsOk

`func (o *SearchBody) GetFormsOk() (*string, bool)`

GetFormsOk returns a tuple with the Forms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForms

`func (o *SearchBody) SetForms(v string)`

SetForms sets Forms field to given value.

### HasForms

`func (o *SearchBody) HasForms() bool`

HasForms returns a boolean if a field has been set.

### GetGics

`func (o *SearchBody) GetGics() string`

GetGics returns the Gics field if non-nil, zero value otherwise.

### GetGicsOk

`func (o *SearchBody) GetGicsOk() (*string, bool)`

GetGicsOk returns a tuple with the Gics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGics

`func (o *SearchBody) SetGics(v string)`

SetGics sets Gics field to given value.

### HasGics

`func (o *SearchBody) HasGics() bool`

HasGics returns a boolean if a field has been set.

### GetNaics

`func (o *SearchBody) GetNaics() string`

GetNaics returns the Naics field if non-nil, zero value otherwise.

### GetNaicsOk

`func (o *SearchBody) GetNaicsOk() (*string, bool)`

GetNaicsOk returns a tuple with the Naics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNaics

`func (o *SearchBody) SetNaics(v string)`

SetNaics sets Naics field to given value.

### HasNaics

`func (o *SearchBody) HasNaics() bool`

HasNaics returns a boolean if a field has been set.

### GetExhibits

`func (o *SearchBody) GetExhibits() string`

GetExhibits returns the Exhibits field if non-nil, zero value otherwise.

### GetExhibitsOk

`func (o *SearchBody) GetExhibitsOk() (*string, bool)`

GetExhibitsOk returns a tuple with the Exhibits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExhibits

`func (o *SearchBody) SetExhibits(v string)`

SetExhibits sets Exhibits field to given value.

### HasExhibits

`func (o *SearchBody) HasExhibits() bool`

HasExhibits returns a boolean if a field has been set.

### GetExchanges

`func (o *SearchBody) GetExchanges() string`

GetExchanges returns the Exchanges field if non-nil, zero value otherwise.

### GetExchangesOk

`func (o *SearchBody) GetExchangesOk() (*string, bool)`

GetExchangesOk returns a tuple with the Exchanges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExchanges

`func (o *SearchBody) SetExchanges(v string)`

SetExchanges sets Exchanges field to given value.

### HasExchanges

`func (o *SearchBody) HasExchanges() bool`

HasExchanges returns a boolean if a field has been set.

### GetCountries

`func (o *SearchBody) GetCountries() string`

GetCountries returns the Countries field if non-nil, zero value otherwise.

### GetCountriesOk

`func (o *SearchBody) GetCountriesOk() (*string, bool)`

GetCountriesOk returns a tuple with the Countries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountries

`func (o *SearchBody) SetCountries(v string)`

SetCountries sets Countries field to given value.

### HasCountries

`func (o *SearchBody) HasCountries() bool`

HasCountries returns a boolean if a field has been set.

### GetActs

`func (o *SearchBody) GetActs() string`

GetActs returns the Acts field if non-nil, zero value otherwise.

### GetActsOk

`func (o *SearchBody) GetActsOk() (*string, bool)`

GetActsOk returns a tuple with the Acts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActs

`func (o *SearchBody) SetActs(v string)`

SetActs sets Acts field to given value.

### HasActs

`func (o *SearchBody) HasActs() bool`

HasActs returns a boolean if a field has been set.

### GetCaps

`func (o *SearchBody) GetCaps() string`

GetCaps returns the Caps field if non-nil, zero value otherwise.

### GetCapsOk

`func (o *SearchBody) GetCapsOk() (*string, bool)`

GetCapsOk returns a tuple with the Caps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaps

`func (o *SearchBody) SetCaps(v string)`

SetCaps sets Caps field to given value.

### HasCaps

`func (o *SearchBody) HasCaps() bool`

HasCaps returns a boolean if a field has been set.

### GetFromDate

`func (o *SearchBody) GetFromDate() string`

GetFromDate returns the FromDate field if non-nil, zero value otherwise.

### GetFromDateOk

`func (o *SearchBody) GetFromDateOk() (*string, bool)`

GetFromDateOk returns a tuple with the FromDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFromDate

`func (o *SearchBody) SetFromDate(v string)`

SetFromDate sets FromDate field to given value.

### HasFromDate

`func (o *SearchBody) HasFromDate() bool`

HasFromDate returns a boolean if a field has been set.

### GetToDate

`func (o *SearchBody) GetToDate() string`

GetToDate returns the ToDate field if non-nil, zero value otherwise.

### GetToDateOk

`func (o *SearchBody) GetToDateOk() (*string, bool)`

GetToDateOk returns a tuple with the ToDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToDate

`func (o *SearchBody) SetToDate(v string)`

SetToDate sets ToDate field to given value.

### HasToDate

`func (o *SearchBody) HasToDate() bool`

HasToDate returns a boolean if a field has been set.

### GetPage

`func (o *SearchBody) GetPage() string`

GetPage returns the Page field if non-nil, zero value otherwise.

### GetPageOk

`func (o *SearchBody) GetPageOk() (*string, bool)`

GetPageOk returns a tuple with the Page field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPage

`func (o *SearchBody) SetPage(v string)`

SetPage sets Page field to given value.

### HasPage

`func (o *SearchBody) HasPage() bool`

HasPage returns a boolean if a field has been set.

### GetSort

`func (o *SearchBody) GetSort() string`

GetSort returns the Sort field if non-nil, zero value otherwise.

### GetSortOk

`func (o *SearchBody) GetSortOk() (*string, bool)`

GetSortOk returns a tuple with the Sort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSort

`func (o *SearchBody) SetSort(v string)`

SetSort sets Sort field to given value.

### HasSort

`func (o *SearchBody) HasSort() bool`

HasSort returns a boolean if a field has been set.

### GetHighlighted

`func (o *SearchBody) GetHighlighted() bool`

GetHighlighted returns the Highlighted field if non-nil, zero value otherwise.

### GetHighlightedOk

`func (o *SearchBody) GetHighlightedOk() (*bool, bool)`

GetHighlightedOk returns a tuple with the Highlighted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHighlighted

`func (o *SearchBody) SetHighlighted(v bool)`

SetHighlighted sets Highlighted field to given value.

### HasHighlighted

`func (o *SearchBody) HasHighlighted() bool`

HasHighlighted returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


