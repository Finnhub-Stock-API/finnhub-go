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

// ETFProfileData struct for ETFProfileData
type ETFProfileData struct {
	// Name
	Name *string `json:"name,omitempty"`
	// Asset Class.
	AssetClass *string `json:"assetClass,omitempty"`
	// Investment Segment.
	InvestmentSegment *string `json:"investmentSegment,omitempty"`
	// AUM.
	Aum *float32 `json:"aum,omitempty"`
	// NAV.
	Nav *float32 `json:"nav,omitempty"`
	// NAV currency.
	NavCurrency *string `json:"navCurrency,omitempty"`
	// Expense ratio. For non-US funds, this is the <a href=\"https://www.esma.europa.eu/sites/default/files/library/2015/11/09_1028_final_kid_ongoing_charges_methodology_for_publication_u_2_.pdf\" target=\"_blank\">KID ongoing charges<a/>.
	ExpenseRatio *float32 `json:"expenseRatio,omitempty"`
	// Tracking Index.
	TrackingIndex *string `json:"trackingIndex,omitempty"`
	// ETF issuer.
	EtfCompany *string `json:"etfCompany,omitempty"`
	// ETF domicile.
	Domicile *string `json:"domicile,omitempty"`
	// Inception date.
	InceptionDate *string `json:"inceptionDate,omitempty"`
	// ETF's website.
	Website *string `json:"website,omitempty"`
	// Logo.
	Logo *string `json:"logo,omitempty"`
	// ISIN.
	Isin *string `json:"isin,omitempty"`
	// CUSIP.
	Cusip *string `json:"cusip,omitempty"`
	// P/E.
	PriceToEarnings *float32 `json:"priceToEarnings,omitempty"`
	// P/B.
	PriceToBook *float32 `json:"priceToBook,omitempty"`
	// 30-day average volume.
	AvgVolume *float32 `json:"avgVolume,omitempty"`
	// ETF's description.
	Description *string `json:"description,omitempty"`
	// Whether the ETF is inverse
	IsInverse *bool `json:"isInverse,omitempty"`
	// Whether the ETF is leveraged
	IsLeveraged *bool `json:"isLeveraged,omitempty"`
	// Leverage factor.
	LeverageFactor *float32 `json:"leverageFactor,omitempty"`
}

// NewETFProfileData instantiates a new ETFProfileData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewETFProfileData() *ETFProfileData {
	this := ETFProfileData{}
	return &this
}

// NewETFProfileDataWithDefaults instantiates a new ETFProfileData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewETFProfileDataWithDefaults() *ETFProfileData {
	this := ETFProfileData{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ETFProfileData) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ETFProfileData) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ETFProfileData) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ETFProfileData) SetName(v string) {
	o.Name = &v
}

// GetAssetClass returns the AssetClass field value if set, zero value otherwise.
func (o *ETFProfileData) GetAssetClass() string {
	if o == nil || o.AssetClass == nil {
		var ret string
		return ret
	}
	return *o.AssetClass
}

// GetAssetClassOk returns a tuple with the AssetClass field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ETFProfileData) GetAssetClassOk() (*string, bool) {
	if o == nil || o.AssetClass == nil {
		return nil, false
	}
	return o.AssetClass, true
}

// HasAssetClass returns a boolean if a field has been set.
func (o *ETFProfileData) HasAssetClass() bool {
	if o != nil && o.AssetClass != nil {
		return true
	}

	return false
}

// SetAssetClass gets a reference to the given string and assigns it to the AssetClass field.
func (o *ETFProfileData) SetAssetClass(v string) {
	o.AssetClass = &v
}

// GetInvestmentSegment returns the InvestmentSegment field value if set, zero value otherwise.
func (o *ETFProfileData) GetInvestmentSegment() string {
	if o == nil || o.InvestmentSegment == nil {
		var ret string
		return ret
	}
	return *o.InvestmentSegment
}

// GetInvestmentSegmentOk returns a tuple with the InvestmentSegment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ETFProfileData) GetInvestmentSegmentOk() (*string, bool) {
	if o == nil || o.InvestmentSegment == nil {
		return nil, false
	}
	return o.InvestmentSegment, true
}

// HasInvestmentSegment returns a boolean if a field has been set.
func (o *ETFProfileData) HasInvestmentSegment() bool {
	if o != nil && o.InvestmentSegment != nil {
		return true
	}

	return false
}

// SetInvestmentSegment gets a reference to the given string and assigns it to the InvestmentSegment field.
func (o *ETFProfileData) SetInvestmentSegment(v string) {
	o.InvestmentSegment = &v
}

// GetAum returns the Aum field value if set, zero value otherwise.
func (o *ETFProfileData) GetAum() float32 {
	if o == nil || o.Aum == nil {
		var ret float32
		return ret
	}
	return *o.Aum
}

// GetAumOk returns a tuple with the Aum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ETFProfileData) GetAumOk() (*float32, bool) {
	if o == nil || o.Aum == nil {
		return nil, false
	}
	return o.Aum, true
}

// HasAum returns a boolean if a field has been set.
func (o *ETFProfileData) HasAum() bool {
	if o != nil && o.Aum != nil {
		return true
	}

	return false
}

// SetAum gets a reference to the given float32 and assigns it to the Aum field.
func (o *ETFProfileData) SetAum(v float32) {
	o.Aum = &v
}

// GetNav returns the Nav field value if set, zero value otherwise.
func (o *ETFProfileData) GetNav() float32 {
	if o == nil || o.Nav == nil {
		var ret float32
		return ret
	}
	return *o.Nav
}

// GetNavOk returns a tuple with the Nav field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ETFProfileData) GetNavOk() (*float32, bool) {
	if o == nil || o.Nav == nil {
		return nil, false
	}
	return o.Nav, true
}

// HasNav returns a boolean if a field has been set.
func (o *ETFProfileData) HasNav() bool {
	if o != nil && o.Nav != nil {
		return true
	}

	return false
}

// SetNav gets a reference to the given float32 and assigns it to the Nav field.
func (o *ETFProfileData) SetNav(v float32) {
	o.Nav = &v
}

// GetNavCurrency returns the NavCurrency field value if set, zero value otherwise.
func (o *ETFProfileData) GetNavCurrency() string {
	if o == nil || o.NavCurrency == nil {
		var ret string
		return ret
	}
	return *o.NavCurrency
}

// GetNavCurrencyOk returns a tuple with the NavCurrency field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ETFProfileData) GetNavCurrencyOk() (*string, bool) {
	if o == nil || o.NavCurrency == nil {
		return nil, false
	}
	return o.NavCurrency, true
}

// HasNavCurrency returns a boolean if a field has been set.
func (o *ETFProfileData) HasNavCurrency() bool {
	if o != nil && o.NavCurrency != nil {
		return true
	}

	return false
}

// SetNavCurrency gets a reference to the given string and assigns it to the NavCurrency field.
func (o *ETFProfileData) SetNavCurrency(v string) {
	o.NavCurrency = &v
}

// GetExpenseRatio returns the ExpenseRatio field value if set, zero value otherwise.
func (o *ETFProfileData) GetExpenseRatio() float32 {
	if o == nil || o.ExpenseRatio == nil {
		var ret float32
		return ret
	}
	return *o.ExpenseRatio
}

// GetExpenseRatioOk returns a tuple with the ExpenseRatio field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ETFProfileData) GetExpenseRatioOk() (*float32, bool) {
	if o == nil || o.ExpenseRatio == nil {
		return nil, false
	}
	return o.ExpenseRatio, true
}

// HasExpenseRatio returns a boolean if a field has been set.
func (o *ETFProfileData) HasExpenseRatio() bool {
	if o != nil && o.ExpenseRatio != nil {
		return true
	}

	return false
}

// SetExpenseRatio gets a reference to the given float32 and assigns it to the ExpenseRatio field.
func (o *ETFProfileData) SetExpenseRatio(v float32) {
	o.ExpenseRatio = &v
}

// GetTrackingIndex returns the TrackingIndex field value if set, zero value otherwise.
func (o *ETFProfileData) GetTrackingIndex() string {
	if o == nil || o.TrackingIndex == nil {
		var ret string
		return ret
	}
	return *o.TrackingIndex
}

// GetTrackingIndexOk returns a tuple with the TrackingIndex field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ETFProfileData) GetTrackingIndexOk() (*string, bool) {
	if o == nil || o.TrackingIndex == nil {
		return nil, false
	}
	return o.TrackingIndex, true
}

// HasTrackingIndex returns a boolean if a field has been set.
func (o *ETFProfileData) HasTrackingIndex() bool {
	if o != nil && o.TrackingIndex != nil {
		return true
	}

	return false
}

// SetTrackingIndex gets a reference to the given string and assigns it to the TrackingIndex field.
func (o *ETFProfileData) SetTrackingIndex(v string) {
	o.TrackingIndex = &v
}

// GetEtfCompany returns the EtfCompany field value if set, zero value otherwise.
func (o *ETFProfileData) GetEtfCompany() string {
	if o == nil || o.EtfCompany == nil {
		var ret string
		return ret
	}
	return *o.EtfCompany
}

// GetEtfCompanyOk returns a tuple with the EtfCompany field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ETFProfileData) GetEtfCompanyOk() (*string, bool) {
	if o == nil || o.EtfCompany == nil {
		return nil, false
	}
	return o.EtfCompany, true
}

// HasEtfCompany returns a boolean if a field has been set.
func (o *ETFProfileData) HasEtfCompany() bool {
	if o != nil && o.EtfCompany != nil {
		return true
	}

	return false
}

// SetEtfCompany gets a reference to the given string and assigns it to the EtfCompany field.
func (o *ETFProfileData) SetEtfCompany(v string) {
	o.EtfCompany = &v
}

// GetDomicile returns the Domicile field value if set, zero value otherwise.
func (o *ETFProfileData) GetDomicile() string {
	if o == nil || o.Domicile == nil {
		var ret string
		return ret
	}
	return *o.Domicile
}

// GetDomicileOk returns a tuple with the Domicile field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ETFProfileData) GetDomicileOk() (*string, bool) {
	if o == nil || o.Domicile == nil {
		return nil, false
	}
	return o.Domicile, true
}

// HasDomicile returns a boolean if a field has been set.
func (o *ETFProfileData) HasDomicile() bool {
	if o != nil && o.Domicile != nil {
		return true
	}

	return false
}

// SetDomicile gets a reference to the given string and assigns it to the Domicile field.
func (o *ETFProfileData) SetDomicile(v string) {
	o.Domicile = &v
}

// GetInceptionDate returns the InceptionDate field value if set, zero value otherwise.
func (o *ETFProfileData) GetInceptionDate() string {
	if o == nil || o.InceptionDate == nil {
		var ret string
		return ret
	}
	return *o.InceptionDate
}

// GetInceptionDateOk returns a tuple with the InceptionDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ETFProfileData) GetInceptionDateOk() (*string, bool) {
	if o == nil || o.InceptionDate == nil {
		return nil, false
	}
	return o.InceptionDate, true
}

// HasInceptionDate returns a boolean if a field has been set.
func (o *ETFProfileData) HasInceptionDate() bool {
	if o != nil && o.InceptionDate != nil {
		return true
	}

	return false
}

// SetInceptionDate gets a reference to the given string and assigns it to the InceptionDate field.
func (o *ETFProfileData) SetInceptionDate(v string) {
	o.InceptionDate = &v
}

// GetWebsite returns the Website field value if set, zero value otherwise.
func (o *ETFProfileData) GetWebsite() string {
	if o == nil || o.Website == nil {
		var ret string
		return ret
	}
	return *o.Website
}

// GetWebsiteOk returns a tuple with the Website field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ETFProfileData) GetWebsiteOk() (*string, bool) {
	if o == nil || o.Website == nil {
		return nil, false
	}
	return o.Website, true
}

// HasWebsite returns a boolean if a field has been set.
func (o *ETFProfileData) HasWebsite() bool {
	if o != nil && o.Website != nil {
		return true
	}

	return false
}

// SetWebsite gets a reference to the given string and assigns it to the Website field.
func (o *ETFProfileData) SetWebsite(v string) {
	o.Website = &v
}

// GetLogo returns the Logo field value if set, zero value otherwise.
func (o *ETFProfileData) GetLogo() string {
	if o == nil || o.Logo == nil {
		var ret string
		return ret
	}
	return *o.Logo
}

// GetLogoOk returns a tuple with the Logo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ETFProfileData) GetLogoOk() (*string, bool) {
	if o == nil || o.Logo == nil {
		return nil, false
	}
	return o.Logo, true
}

// HasLogo returns a boolean if a field has been set.
func (o *ETFProfileData) HasLogo() bool {
	if o != nil && o.Logo != nil {
		return true
	}

	return false
}

// SetLogo gets a reference to the given string and assigns it to the Logo field.
func (o *ETFProfileData) SetLogo(v string) {
	o.Logo = &v
}

// GetIsin returns the Isin field value if set, zero value otherwise.
func (o *ETFProfileData) GetIsin() string {
	if o == nil || o.Isin == nil {
		var ret string
		return ret
	}
	return *o.Isin
}

// GetIsinOk returns a tuple with the Isin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ETFProfileData) GetIsinOk() (*string, bool) {
	if o == nil || o.Isin == nil {
		return nil, false
	}
	return o.Isin, true
}

// HasIsin returns a boolean if a field has been set.
func (o *ETFProfileData) HasIsin() bool {
	if o != nil && o.Isin != nil {
		return true
	}

	return false
}

// SetIsin gets a reference to the given string and assigns it to the Isin field.
func (o *ETFProfileData) SetIsin(v string) {
	o.Isin = &v
}

// GetCusip returns the Cusip field value if set, zero value otherwise.
func (o *ETFProfileData) GetCusip() string {
	if o == nil || o.Cusip == nil {
		var ret string
		return ret
	}
	return *o.Cusip
}

// GetCusipOk returns a tuple with the Cusip field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ETFProfileData) GetCusipOk() (*string, bool) {
	if o == nil || o.Cusip == nil {
		return nil, false
	}
	return o.Cusip, true
}

// HasCusip returns a boolean if a field has been set.
func (o *ETFProfileData) HasCusip() bool {
	if o != nil && o.Cusip != nil {
		return true
	}

	return false
}

// SetCusip gets a reference to the given string and assigns it to the Cusip field.
func (o *ETFProfileData) SetCusip(v string) {
	o.Cusip = &v
}

// GetPriceToEarnings returns the PriceToEarnings field value if set, zero value otherwise.
func (o *ETFProfileData) GetPriceToEarnings() float32 {
	if o == nil || o.PriceToEarnings == nil {
		var ret float32
		return ret
	}
	return *o.PriceToEarnings
}

// GetPriceToEarningsOk returns a tuple with the PriceToEarnings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ETFProfileData) GetPriceToEarningsOk() (*float32, bool) {
	if o == nil || o.PriceToEarnings == nil {
		return nil, false
	}
	return o.PriceToEarnings, true
}

// HasPriceToEarnings returns a boolean if a field has been set.
func (o *ETFProfileData) HasPriceToEarnings() bool {
	if o != nil && o.PriceToEarnings != nil {
		return true
	}

	return false
}

// SetPriceToEarnings gets a reference to the given float32 and assigns it to the PriceToEarnings field.
func (o *ETFProfileData) SetPriceToEarnings(v float32) {
	o.PriceToEarnings = &v
}

// GetPriceToBook returns the PriceToBook field value if set, zero value otherwise.
func (o *ETFProfileData) GetPriceToBook() float32 {
	if o == nil || o.PriceToBook == nil {
		var ret float32
		return ret
	}
	return *o.PriceToBook
}

// GetPriceToBookOk returns a tuple with the PriceToBook field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ETFProfileData) GetPriceToBookOk() (*float32, bool) {
	if o == nil || o.PriceToBook == nil {
		return nil, false
	}
	return o.PriceToBook, true
}

// HasPriceToBook returns a boolean if a field has been set.
func (o *ETFProfileData) HasPriceToBook() bool {
	if o != nil && o.PriceToBook != nil {
		return true
	}

	return false
}

// SetPriceToBook gets a reference to the given float32 and assigns it to the PriceToBook field.
func (o *ETFProfileData) SetPriceToBook(v float32) {
	o.PriceToBook = &v
}

// GetAvgVolume returns the AvgVolume field value if set, zero value otherwise.
func (o *ETFProfileData) GetAvgVolume() float32 {
	if o == nil || o.AvgVolume == nil {
		var ret float32
		return ret
	}
	return *o.AvgVolume
}

// GetAvgVolumeOk returns a tuple with the AvgVolume field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ETFProfileData) GetAvgVolumeOk() (*float32, bool) {
	if o == nil || o.AvgVolume == nil {
		return nil, false
	}
	return o.AvgVolume, true
}

// HasAvgVolume returns a boolean if a field has been set.
func (o *ETFProfileData) HasAvgVolume() bool {
	if o != nil && o.AvgVolume != nil {
		return true
	}

	return false
}

// SetAvgVolume gets a reference to the given float32 and assigns it to the AvgVolume field.
func (o *ETFProfileData) SetAvgVolume(v float32) {
	o.AvgVolume = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *ETFProfileData) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ETFProfileData) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *ETFProfileData) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *ETFProfileData) SetDescription(v string) {
	o.Description = &v
}

// GetIsInverse returns the IsInverse field value if set, zero value otherwise.
func (o *ETFProfileData) GetIsInverse() bool {
	if o == nil || o.IsInverse == nil {
		var ret bool
		return ret
	}
	return *o.IsInverse
}

// GetIsInverseOk returns a tuple with the IsInverse field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ETFProfileData) GetIsInverseOk() (*bool, bool) {
	if o == nil || o.IsInverse == nil {
		return nil, false
	}
	return o.IsInverse, true
}

// HasIsInverse returns a boolean if a field has been set.
func (o *ETFProfileData) HasIsInverse() bool {
	if o != nil && o.IsInverse != nil {
		return true
	}

	return false
}

// SetIsInverse gets a reference to the given bool and assigns it to the IsInverse field.
func (o *ETFProfileData) SetIsInverse(v bool) {
	o.IsInverse = &v
}

// GetIsLeveraged returns the IsLeveraged field value if set, zero value otherwise.
func (o *ETFProfileData) GetIsLeveraged() bool {
	if o == nil || o.IsLeveraged == nil {
		var ret bool
		return ret
	}
	return *o.IsLeveraged
}

// GetIsLeveragedOk returns a tuple with the IsLeveraged field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ETFProfileData) GetIsLeveragedOk() (*bool, bool) {
	if o == nil || o.IsLeveraged == nil {
		return nil, false
	}
	return o.IsLeveraged, true
}

// HasIsLeveraged returns a boolean if a field has been set.
func (o *ETFProfileData) HasIsLeveraged() bool {
	if o != nil && o.IsLeveraged != nil {
		return true
	}

	return false
}

// SetIsLeveraged gets a reference to the given bool and assigns it to the IsLeveraged field.
func (o *ETFProfileData) SetIsLeveraged(v bool) {
	o.IsLeveraged = &v
}

// GetLeverageFactor returns the LeverageFactor field value if set, zero value otherwise.
func (o *ETFProfileData) GetLeverageFactor() float32 {
	if o == nil || o.LeverageFactor == nil {
		var ret float32
		return ret
	}
	return *o.LeverageFactor
}

// GetLeverageFactorOk returns a tuple with the LeverageFactor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ETFProfileData) GetLeverageFactorOk() (*float32, bool) {
	if o == nil || o.LeverageFactor == nil {
		return nil, false
	}
	return o.LeverageFactor, true
}

// HasLeverageFactor returns a boolean if a field has been set.
func (o *ETFProfileData) HasLeverageFactor() bool {
	if o != nil && o.LeverageFactor != nil {
		return true
	}

	return false
}

// SetLeverageFactor gets a reference to the given float32 and assigns it to the LeverageFactor field.
func (o *ETFProfileData) SetLeverageFactor(v float32) {
	o.LeverageFactor = &v
}

func (o ETFProfileData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.AssetClass != nil {
		toSerialize["assetClass"] = o.AssetClass
	}
	if o.InvestmentSegment != nil {
		toSerialize["investmentSegment"] = o.InvestmentSegment
	}
	if o.Aum != nil {
		toSerialize["aum"] = o.Aum
	}
	if o.Nav != nil {
		toSerialize["nav"] = o.Nav
	}
	if o.NavCurrency != nil {
		toSerialize["navCurrency"] = o.NavCurrency
	}
	if o.ExpenseRatio != nil {
		toSerialize["expenseRatio"] = o.ExpenseRatio
	}
	if o.TrackingIndex != nil {
		toSerialize["trackingIndex"] = o.TrackingIndex
	}
	if o.EtfCompany != nil {
		toSerialize["etfCompany"] = o.EtfCompany
	}
	if o.Domicile != nil {
		toSerialize["domicile"] = o.Domicile
	}
	if o.InceptionDate != nil {
		toSerialize["inceptionDate"] = o.InceptionDate
	}
	if o.Website != nil {
		toSerialize["website"] = o.Website
	}
	if o.Logo != nil {
		toSerialize["logo"] = o.Logo
	}
	if o.Isin != nil {
		toSerialize["isin"] = o.Isin
	}
	if o.Cusip != nil {
		toSerialize["cusip"] = o.Cusip
	}
	if o.PriceToEarnings != nil {
		toSerialize["priceToEarnings"] = o.PriceToEarnings
	}
	if o.PriceToBook != nil {
		toSerialize["priceToBook"] = o.PriceToBook
	}
	if o.AvgVolume != nil {
		toSerialize["avgVolume"] = o.AvgVolume
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.IsInverse != nil {
		toSerialize["isInverse"] = o.IsInverse
	}
	if o.IsLeveraged != nil {
		toSerialize["isLeveraged"] = o.IsLeveraged
	}
	if o.LeverageFactor != nil {
		toSerialize["leverageFactor"] = o.LeverageFactor
	}
	return json.Marshal(toSerialize)
}

type NullableETFProfileData struct {
	value *ETFProfileData
	isSet bool
}

func (v NullableETFProfileData) Get() *ETFProfileData {
	return v.value
}

func (v *NullableETFProfileData) Set(val *ETFProfileData) {
	v.value = val
	v.isSet = true
}

func (v NullableETFProfileData) IsSet() bool {
	return v.isSet
}

func (v *NullableETFProfileData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableETFProfileData(val *ETFProfileData) *NullableETFProfileData {
	return &NullableETFProfileData{value: val, isSet: true}
}

func (v NullableETFProfileData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableETFProfileData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


