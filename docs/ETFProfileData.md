# ETFProfileData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | Name | [optional] 
**AssetClass** | Pointer to **string** | Asset Class. | [optional] 
**InvestmentSegment** | Pointer to **string** | Investment Segment. | [optional] 
**Aum** | Pointer to **float32** | AUM. | [optional] 
**Nav** | Pointer to **float32** | NAV. | [optional] 
**NavCurrency** | Pointer to **string** | NAV currency. | [optional] 
**ExpenseRatio** | Pointer to **float32** | Expense ratio. For non-US funds, this is the &lt;a href&#x3D;\&quot;https://www.esma.europa.eu/sites/default/files/library/2015/11/09_1028_final_kid_ongoing_charges_methodology_for_publication_u_2_.pdf\&quot; target&#x3D;\&quot;_blank\&quot;&gt;KID ongoing charges&lt;a/&gt;. | [optional] 
**TrackingIndex** | Pointer to **string** | Tracking Index. | [optional] 
**EtfCompany** | Pointer to **string** | ETF issuer. | [optional] 
**Domicile** | Pointer to **string** | ETF domicile. | [optional] 
**InceptionDate** | Pointer to **string** | Inception date. | [optional] 
**Website** | Pointer to **string** | ETF&#39;s website. | [optional] 
**Logo** | Pointer to **string** | Logo. | [optional] 
**Isin** | Pointer to **string** | ISIN. | [optional] 
**Cusip** | Pointer to **string** | CUSIP. | [optional] 
**PriceToEarnings** | Pointer to **float32** | P/E. | [optional] 
**PriceToBook** | Pointer to **float32** | P/B. | [optional] 
**AvgVolume** | Pointer to **float32** | 30-day average volume. | [optional] 
**Description** | Pointer to **string** | ETF&#39;s description. | [optional] 
**IsInverse** | Pointer to **bool** | Whether the ETF is inverse | [optional] 
**IsLeveraged** | Pointer to **bool** | Whether the ETF is leveraged | [optional] 
**LeverageFactor** | Pointer to **float32** | Leverage factor. | [optional] 

## Methods

### NewETFProfileData

`func NewETFProfileData() *ETFProfileData`

NewETFProfileData instantiates a new ETFProfileData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewETFProfileDataWithDefaults

`func NewETFProfileDataWithDefaults() *ETFProfileData`

NewETFProfileDataWithDefaults instantiates a new ETFProfileData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ETFProfileData) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ETFProfileData) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ETFProfileData) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ETFProfileData) HasName() bool`

HasName returns a boolean if a field has been set.

### GetAssetClass

`func (o *ETFProfileData) GetAssetClass() string`

GetAssetClass returns the AssetClass field if non-nil, zero value otherwise.

### GetAssetClassOk

`func (o *ETFProfileData) GetAssetClassOk() (*string, bool)`

GetAssetClassOk returns a tuple with the AssetClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssetClass

`func (o *ETFProfileData) SetAssetClass(v string)`

SetAssetClass sets AssetClass field to given value.

### HasAssetClass

`func (o *ETFProfileData) HasAssetClass() bool`

HasAssetClass returns a boolean if a field has been set.

### GetInvestmentSegment

`func (o *ETFProfileData) GetInvestmentSegment() string`

GetInvestmentSegment returns the InvestmentSegment field if non-nil, zero value otherwise.

### GetInvestmentSegmentOk

`func (o *ETFProfileData) GetInvestmentSegmentOk() (*string, bool)`

GetInvestmentSegmentOk returns a tuple with the InvestmentSegment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvestmentSegment

`func (o *ETFProfileData) SetInvestmentSegment(v string)`

SetInvestmentSegment sets InvestmentSegment field to given value.

### HasInvestmentSegment

`func (o *ETFProfileData) HasInvestmentSegment() bool`

HasInvestmentSegment returns a boolean if a field has been set.

### GetAum

`func (o *ETFProfileData) GetAum() float32`

GetAum returns the Aum field if non-nil, zero value otherwise.

### GetAumOk

`func (o *ETFProfileData) GetAumOk() (*float32, bool)`

GetAumOk returns a tuple with the Aum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAum

`func (o *ETFProfileData) SetAum(v float32)`

SetAum sets Aum field to given value.

### HasAum

`func (o *ETFProfileData) HasAum() bool`

HasAum returns a boolean if a field has been set.

### GetNav

`func (o *ETFProfileData) GetNav() float32`

GetNav returns the Nav field if non-nil, zero value otherwise.

### GetNavOk

`func (o *ETFProfileData) GetNavOk() (*float32, bool)`

GetNavOk returns a tuple with the Nav field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNav

`func (o *ETFProfileData) SetNav(v float32)`

SetNav sets Nav field to given value.

### HasNav

`func (o *ETFProfileData) HasNav() bool`

HasNav returns a boolean if a field has been set.

### GetNavCurrency

`func (o *ETFProfileData) GetNavCurrency() string`

GetNavCurrency returns the NavCurrency field if non-nil, zero value otherwise.

### GetNavCurrencyOk

`func (o *ETFProfileData) GetNavCurrencyOk() (*string, bool)`

GetNavCurrencyOk returns a tuple with the NavCurrency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNavCurrency

`func (o *ETFProfileData) SetNavCurrency(v string)`

SetNavCurrency sets NavCurrency field to given value.

### HasNavCurrency

`func (o *ETFProfileData) HasNavCurrency() bool`

HasNavCurrency returns a boolean if a field has been set.

### GetExpenseRatio

`func (o *ETFProfileData) GetExpenseRatio() float32`

GetExpenseRatio returns the ExpenseRatio field if non-nil, zero value otherwise.

### GetExpenseRatioOk

`func (o *ETFProfileData) GetExpenseRatioOk() (*float32, bool)`

GetExpenseRatioOk returns a tuple with the ExpenseRatio field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpenseRatio

`func (o *ETFProfileData) SetExpenseRatio(v float32)`

SetExpenseRatio sets ExpenseRatio field to given value.

### HasExpenseRatio

`func (o *ETFProfileData) HasExpenseRatio() bool`

HasExpenseRatio returns a boolean if a field has been set.

### GetTrackingIndex

`func (o *ETFProfileData) GetTrackingIndex() string`

GetTrackingIndex returns the TrackingIndex field if non-nil, zero value otherwise.

### GetTrackingIndexOk

`func (o *ETFProfileData) GetTrackingIndexOk() (*string, bool)`

GetTrackingIndexOk returns a tuple with the TrackingIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrackingIndex

`func (o *ETFProfileData) SetTrackingIndex(v string)`

SetTrackingIndex sets TrackingIndex field to given value.

### HasTrackingIndex

`func (o *ETFProfileData) HasTrackingIndex() bool`

HasTrackingIndex returns a boolean if a field has been set.

### GetEtfCompany

`func (o *ETFProfileData) GetEtfCompany() string`

GetEtfCompany returns the EtfCompany field if non-nil, zero value otherwise.

### GetEtfCompanyOk

`func (o *ETFProfileData) GetEtfCompanyOk() (*string, bool)`

GetEtfCompanyOk returns a tuple with the EtfCompany field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEtfCompany

`func (o *ETFProfileData) SetEtfCompany(v string)`

SetEtfCompany sets EtfCompany field to given value.

### HasEtfCompany

`func (o *ETFProfileData) HasEtfCompany() bool`

HasEtfCompany returns a boolean if a field has been set.

### GetDomicile

`func (o *ETFProfileData) GetDomicile() string`

GetDomicile returns the Domicile field if non-nil, zero value otherwise.

### GetDomicileOk

`func (o *ETFProfileData) GetDomicileOk() (*string, bool)`

GetDomicileOk returns a tuple with the Domicile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomicile

`func (o *ETFProfileData) SetDomicile(v string)`

SetDomicile sets Domicile field to given value.

### HasDomicile

`func (o *ETFProfileData) HasDomicile() bool`

HasDomicile returns a boolean if a field has been set.

### GetInceptionDate

`func (o *ETFProfileData) GetInceptionDate() string`

GetInceptionDate returns the InceptionDate field if non-nil, zero value otherwise.

### GetInceptionDateOk

`func (o *ETFProfileData) GetInceptionDateOk() (*string, bool)`

GetInceptionDateOk returns a tuple with the InceptionDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInceptionDate

`func (o *ETFProfileData) SetInceptionDate(v string)`

SetInceptionDate sets InceptionDate field to given value.

### HasInceptionDate

`func (o *ETFProfileData) HasInceptionDate() bool`

HasInceptionDate returns a boolean if a field has been set.

### GetWebsite

`func (o *ETFProfileData) GetWebsite() string`

GetWebsite returns the Website field if non-nil, zero value otherwise.

### GetWebsiteOk

`func (o *ETFProfileData) GetWebsiteOk() (*string, bool)`

GetWebsiteOk returns a tuple with the Website field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebsite

`func (o *ETFProfileData) SetWebsite(v string)`

SetWebsite sets Website field to given value.

### HasWebsite

`func (o *ETFProfileData) HasWebsite() bool`

HasWebsite returns a boolean if a field has been set.

### GetLogo

`func (o *ETFProfileData) GetLogo() string`

GetLogo returns the Logo field if non-nil, zero value otherwise.

### GetLogoOk

`func (o *ETFProfileData) GetLogoOk() (*string, bool)`

GetLogoOk returns a tuple with the Logo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogo

`func (o *ETFProfileData) SetLogo(v string)`

SetLogo sets Logo field to given value.

### HasLogo

`func (o *ETFProfileData) HasLogo() bool`

HasLogo returns a boolean if a field has been set.

### GetIsin

`func (o *ETFProfileData) GetIsin() string`

GetIsin returns the Isin field if non-nil, zero value otherwise.

### GetIsinOk

`func (o *ETFProfileData) GetIsinOk() (*string, bool)`

GetIsinOk returns a tuple with the Isin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsin

`func (o *ETFProfileData) SetIsin(v string)`

SetIsin sets Isin field to given value.

### HasIsin

`func (o *ETFProfileData) HasIsin() bool`

HasIsin returns a boolean if a field has been set.

### GetCusip

`func (o *ETFProfileData) GetCusip() string`

GetCusip returns the Cusip field if non-nil, zero value otherwise.

### GetCusipOk

`func (o *ETFProfileData) GetCusipOk() (*string, bool)`

GetCusipOk returns a tuple with the Cusip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCusip

`func (o *ETFProfileData) SetCusip(v string)`

SetCusip sets Cusip field to given value.

### HasCusip

`func (o *ETFProfileData) HasCusip() bool`

HasCusip returns a boolean if a field has been set.

### GetPriceToEarnings

`func (o *ETFProfileData) GetPriceToEarnings() float32`

GetPriceToEarnings returns the PriceToEarnings field if non-nil, zero value otherwise.

### GetPriceToEarningsOk

`func (o *ETFProfileData) GetPriceToEarningsOk() (*float32, bool)`

GetPriceToEarningsOk returns a tuple with the PriceToEarnings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceToEarnings

`func (o *ETFProfileData) SetPriceToEarnings(v float32)`

SetPriceToEarnings sets PriceToEarnings field to given value.

### HasPriceToEarnings

`func (o *ETFProfileData) HasPriceToEarnings() bool`

HasPriceToEarnings returns a boolean if a field has been set.

### GetPriceToBook

`func (o *ETFProfileData) GetPriceToBook() float32`

GetPriceToBook returns the PriceToBook field if non-nil, zero value otherwise.

### GetPriceToBookOk

`func (o *ETFProfileData) GetPriceToBookOk() (*float32, bool)`

GetPriceToBookOk returns a tuple with the PriceToBook field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceToBook

`func (o *ETFProfileData) SetPriceToBook(v float32)`

SetPriceToBook sets PriceToBook field to given value.

### HasPriceToBook

`func (o *ETFProfileData) HasPriceToBook() bool`

HasPriceToBook returns a boolean if a field has been set.

### GetAvgVolume

`func (o *ETFProfileData) GetAvgVolume() float32`

GetAvgVolume returns the AvgVolume field if non-nil, zero value otherwise.

### GetAvgVolumeOk

`func (o *ETFProfileData) GetAvgVolumeOk() (*float32, bool)`

GetAvgVolumeOk returns a tuple with the AvgVolume field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvgVolume

`func (o *ETFProfileData) SetAvgVolume(v float32)`

SetAvgVolume sets AvgVolume field to given value.

### HasAvgVolume

`func (o *ETFProfileData) HasAvgVolume() bool`

HasAvgVolume returns a boolean if a field has been set.

### GetDescription

`func (o *ETFProfileData) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ETFProfileData) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ETFProfileData) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ETFProfileData) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetIsInverse

`func (o *ETFProfileData) GetIsInverse() bool`

GetIsInverse returns the IsInverse field if non-nil, zero value otherwise.

### GetIsInverseOk

`func (o *ETFProfileData) GetIsInverseOk() (*bool, bool)`

GetIsInverseOk returns a tuple with the IsInverse field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsInverse

`func (o *ETFProfileData) SetIsInverse(v bool)`

SetIsInverse sets IsInverse field to given value.

### HasIsInverse

`func (o *ETFProfileData) HasIsInverse() bool`

HasIsInverse returns a boolean if a field has been set.

### GetIsLeveraged

`func (o *ETFProfileData) GetIsLeveraged() bool`

GetIsLeveraged returns the IsLeveraged field if non-nil, zero value otherwise.

### GetIsLeveragedOk

`func (o *ETFProfileData) GetIsLeveragedOk() (*bool, bool)`

GetIsLeveragedOk returns a tuple with the IsLeveraged field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsLeveraged

`func (o *ETFProfileData) SetIsLeveraged(v bool)`

SetIsLeveraged sets IsLeveraged field to given value.

### HasIsLeveraged

`func (o *ETFProfileData) HasIsLeveraged() bool`

HasIsLeveraged returns a boolean if a field has been set.

### GetLeverageFactor

`func (o *ETFProfileData) GetLeverageFactor() float32`

GetLeverageFactor returns the LeverageFactor field if non-nil, zero value otherwise.

### GetLeverageFactorOk

`func (o *ETFProfileData) GetLeverageFactorOk() (*float32, bool)`

GetLeverageFactorOk returns a tuple with the LeverageFactor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLeverageFactor

`func (o *ETFProfileData) SetLeverageFactor(v float32)`

SetLeverageFactor sets LeverageFactor field to given value.

### HasLeverageFactor

`func (o *ETFProfileData) HasLeverageFactor() bool`

HasLeverageFactor returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


