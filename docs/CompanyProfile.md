# CompanyProfile

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Address** | Pointer to **string** | Address of company&#39;s headquarter. | [optional] 
**City** | Pointer to **string** | City of company&#39;s headquarter. | [optional] 
**Country** | Pointer to **string** | Country of company&#39;s headquarter. | [optional] 
**Currency** | Pointer to **string** | Currency used in company filings. | [optional] 
**Cusip** | Pointer to **string** | CUSIP number. | [optional] 
**Sedol** | Pointer to **string** | Sedol number. | [optional] 
**Description** | Pointer to **string** | Company business summary. | [optional] 
**Exchange** | Pointer to **string** | Listed exchange. | [optional] 
**Ggroup** | Pointer to **string** | GICS industry group. | [optional] 
**Gind** | Pointer to **string** | GICS industry. | [optional] 
**Gsector** | Pointer to **string** | GICS sector. | [optional] 
**Gsubind** | Pointer to **string** | GICS sub-industry. | [optional] 
**Isin** | Pointer to **string** | ISIN number. | [optional] 
**NaicsNationalIndustry** | Pointer to **string** | NAICS national industry. | [optional] 
**Naics** | Pointer to **string** | NAICS industry. | [optional] 
**NaicsSector** | Pointer to **string** | NAICS sector. | [optional] 
**NaicsSubsector** | Pointer to **string** | NAICS subsector. | [optional] 
**Name** | Pointer to **string** | Company name. | [optional] 
**Phone** | Pointer to **string** | Company phone number. | [optional] 
**State** | Pointer to **string** | State of company&#39;s headquarter. | [optional] 
**Ticker** | Pointer to **string** | Company symbol/ticker as used on the listed exchange. | [optional] 
**Weburl** | Pointer to **string** | Company website. | [optional] 
**Ipo** | Pointer to **string** | IPO date. | [optional] 
**MarketCapitalization** | Pointer to **float32** | Market Capitalization. | [optional] 
**ShareOutstanding** | Pointer to **float32** | Number of oustanding shares. | [optional] 
**EmployeeTotal** | Pointer to **int64** | Number of employee. | [optional] 
**Logo** | Pointer to **string** | Logo image. | [optional] 
**FinnhubIndustry** | Pointer to **string** | Finnhub industry classification. | [optional] 

## Methods

### NewCompanyProfile

`func NewCompanyProfile() *CompanyProfile`

NewCompanyProfile instantiates a new CompanyProfile object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCompanyProfileWithDefaults

`func NewCompanyProfileWithDefaults() *CompanyProfile`

NewCompanyProfileWithDefaults instantiates a new CompanyProfile object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddress

`func (o *CompanyProfile) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *CompanyProfile) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *CompanyProfile) SetAddress(v string)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *CompanyProfile) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### GetCity

`func (o *CompanyProfile) GetCity() string`

GetCity returns the City field if non-nil, zero value otherwise.

### GetCityOk

`func (o *CompanyProfile) GetCityOk() (*string, bool)`

GetCityOk returns a tuple with the City field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCity

`func (o *CompanyProfile) SetCity(v string)`

SetCity sets City field to given value.

### HasCity

`func (o *CompanyProfile) HasCity() bool`

HasCity returns a boolean if a field has been set.

### GetCountry

`func (o *CompanyProfile) GetCountry() string`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *CompanyProfile) GetCountryOk() (*string, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *CompanyProfile) SetCountry(v string)`

SetCountry sets Country field to given value.

### HasCountry

`func (o *CompanyProfile) HasCountry() bool`

HasCountry returns a boolean if a field has been set.

### GetCurrency

`func (o *CompanyProfile) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *CompanyProfile) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *CompanyProfile) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *CompanyProfile) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetCusip

`func (o *CompanyProfile) GetCusip() string`

GetCusip returns the Cusip field if non-nil, zero value otherwise.

### GetCusipOk

`func (o *CompanyProfile) GetCusipOk() (*string, bool)`

GetCusipOk returns a tuple with the Cusip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCusip

`func (o *CompanyProfile) SetCusip(v string)`

SetCusip sets Cusip field to given value.

### HasCusip

`func (o *CompanyProfile) HasCusip() bool`

HasCusip returns a boolean if a field has been set.

### GetSedol

`func (o *CompanyProfile) GetSedol() string`

GetSedol returns the Sedol field if non-nil, zero value otherwise.

### GetSedolOk

`func (o *CompanyProfile) GetSedolOk() (*string, bool)`

GetSedolOk returns a tuple with the Sedol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSedol

`func (o *CompanyProfile) SetSedol(v string)`

SetSedol sets Sedol field to given value.

### HasSedol

`func (o *CompanyProfile) HasSedol() bool`

HasSedol returns a boolean if a field has been set.

### GetDescription

`func (o *CompanyProfile) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CompanyProfile) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CompanyProfile) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CompanyProfile) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetExchange

`func (o *CompanyProfile) GetExchange() string`

GetExchange returns the Exchange field if non-nil, zero value otherwise.

### GetExchangeOk

`func (o *CompanyProfile) GetExchangeOk() (*string, bool)`

GetExchangeOk returns a tuple with the Exchange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExchange

`func (o *CompanyProfile) SetExchange(v string)`

SetExchange sets Exchange field to given value.

### HasExchange

`func (o *CompanyProfile) HasExchange() bool`

HasExchange returns a boolean if a field has been set.

### GetGgroup

`func (o *CompanyProfile) GetGgroup() string`

GetGgroup returns the Ggroup field if non-nil, zero value otherwise.

### GetGgroupOk

`func (o *CompanyProfile) GetGgroupOk() (*string, bool)`

GetGgroupOk returns a tuple with the Ggroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGgroup

`func (o *CompanyProfile) SetGgroup(v string)`

SetGgroup sets Ggroup field to given value.

### HasGgroup

`func (o *CompanyProfile) HasGgroup() bool`

HasGgroup returns a boolean if a field has been set.

### GetGind

`func (o *CompanyProfile) GetGind() string`

GetGind returns the Gind field if non-nil, zero value otherwise.

### GetGindOk

`func (o *CompanyProfile) GetGindOk() (*string, bool)`

GetGindOk returns a tuple with the Gind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGind

`func (o *CompanyProfile) SetGind(v string)`

SetGind sets Gind field to given value.

### HasGind

`func (o *CompanyProfile) HasGind() bool`

HasGind returns a boolean if a field has been set.

### GetGsector

`func (o *CompanyProfile) GetGsector() string`

GetGsector returns the Gsector field if non-nil, zero value otherwise.

### GetGsectorOk

`func (o *CompanyProfile) GetGsectorOk() (*string, bool)`

GetGsectorOk returns a tuple with the Gsector field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGsector

`func (o *CompanyProfile) SetGsector(v string)`

SetGsector sets Gsector field to given value.

### HasGsector

`func (o *CompanyProfile) HasGsector() bool`

HasGsector returns a boolean if a field has been set.

### GetGsubind

`func (o *CompanyProfile) GetGsubind() string`

GetGsubind returns the Gsubind field if non-nil, zero value otherwise.

### GetGsubindOk

`func (o *CompanyProfile) GetGsubindOk() (*string, bool)`

GetGsubindOk returns a tuple with the Gsubind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGsubind

`func (o *CompanyProfile) SetGsubind(v string)`

SetGsubind sets Gsubind field to given value.

### HasGsubind

`func (o *CompanyProfile) HasGsubind() bool`

HasGsubind returns a boolean if a field has been set.

### GetIsin

`func (o *CompanyProfile) GetIsin() string`

GetIsin returns the Isin field if non-nil, zero value otherwise.

### GetIsinOk

`func (o *CompanyProfile) GetIsinOk() (*string, bool)`

GetIsinOk returns a tuple with the Isin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsin

`func (o *CompanyProfile) SetIsin(v string)`

SetIsin sets Isin field to given value.

### HasIsin

`func (o *CompanyProfile) HasIsin() bool`

HasIsin returns a boolean if a field has been set.

### GetNaicsNationalIndustry

`func (o *CompanyProfile) GetNaicsNationalIndustry() string`

GetNaicsNationalIndustry returns the NaicsNationalIndustry field if non-nil, zero value otherwise.

### GetNaicsNationalIndustryOk

`func (o *CompanyProfile) GetNaicsNationalIndustryOk() (*string, bool)`

GetNaicsNationalIndustryOk returns a tuple with the NaicsNationalIndustry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNaicsNationalIndustry

`func (o *CompanyProfile) SetNaicsNationalIndustry(v string)`

SetNaicsNationalIndustry sets NaicsNationalIndustry field to given value.

### HasNaicsNationalIndustry

`func (o *CompanyProfile) HasNaicsNationalIndustry() bool`

HasNaicsNationalIndustry returns a boolean if a field has been set.

### GetNaics

`func (o *CompanyProfile) GetNaics() string`

GetNaics returns the Naics field if non-nil, zero value otherwise.

### GetNaicsOk

`func (o *CompanyProfile) GetNaicsOk() (*string, bool)`

GetNaicsOk returns a tuple with the Naics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNaics

`func (o *CompanyProfile) SetNaics(v string)`

SetNaics sets Naics field to given value.

### HasNaics

`func (o *CompanyProfile) HasNaics() bool`

HasNaics returns a boolean if a field has been set.

### GetNaicsSector

`func (o *CompanyProfile) GetNaicsSector() string`

GetNaicsSector returns the NaicsSector field if non-nil, zero value otherwise.

### GetNaicsSectorOk

`func (o *CompanyProfile) GetNaicsSectorOk() (*string, bool)`

GetNaicsSectorOk returns a tuple with the NaicsSector field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNaicsSector

`func (o *CompanyProfile) SetNaicsSector(v string)`

SetNaicsSector sets NaicsSector field to given value.

### HasNaicsSector

`func (o *CompanyProfile) HasNaicsSector() bool`

HasNaicsSector returns a boolean if a field has been set.

### GetNaicsSubsector

`func (o *CompanyProfile) GetNaicsSubsector() string`

GetNaicsSubsector returns the NaicsSubsector field if non-nil, zero value otherwise.

### GetNaicsSubsectorOk

`func (o *CompanyProfile) GetNaicsSubsectorOk() (*string, bool)`

GetNaicsSubsectorOk returns a tuple with the NaicsSubsector field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNaicsSubsector

`func (o *CompanyProfile) SetNaicsSubsector(v string)`

SetNaicsSubsector sets NaicsSubsector field to given value.

### HasNaicsSubsector

`func (o *CompanyProfile) HasNaicsSubsector() bool`

HasNaicsSubsector returns a boolean if a field has been set.

### GetName

`func (o *CompanyProfile) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CompanyProfile) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CompanyProfile) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CompanyProfile) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPhone

`func (o *CompanyProfile) GetPhone() string`

GetPhone returns the Phone field if non-nil, zero value otherwise.

### GetPhoneOk

`func (o *CompanyProfile) GetPhoneOk() (*string, bool)`

GetPhoneOk returns a tuple with the Phone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhone

`func (o *CompanyProfile) SetPhone(v string)`

SetPhone sets Phone field to given value.

### HasPhone

`func (o *CompanyProfile) HasPhone() bool`

HasPhone returns a boolean if a field has been set.

### GetState

`func (o *CompanyProfile) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *CompanyProfile) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *CompanyProfile) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *CompanyProfile) HasState() bool`

HasState returns a boolean if a field has been set.

### GetTicker

`func (o *CompanyProfile) GetTicker() string`

GetTicker returns the Ticker field if non-nil, zero value otherwise.

### GetTickerOk

`func (o *CompanyProfile) GetTickerOk() (*string, bool)`

GetTickerOk returns a tuple with the Ticker field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTicker

`func (o *CompanyProfile) SetTicker(v string)`

SetTicker sets Ticker field to given value.

### HasTicker

`func (o *CompanyProfile) HasTicker() bool`

HasTicker returns a boolean if a field has been set.

### GetWeburl

`func (o *CompanyProfile) GetWeburl() string`

GetWeburl returns the Weburl field if non-nil, zero value otherwise.

### GetWeburlOk

`func (o *CompanyProfile) GetWeburlOk() (*string, bool)`

GetWeburlOk returns a tuple with the Weburl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeburl

`func (o *CompanyProfile) SetWeburl(v string)`

SetWeburl sets Weburl field to given value.

### HasWeburl

`func (o *CompanyProfile) HasWeburl() bool`

HasWeburl returns a boolean if a field has been set.

### GetIpo

`func (o *CompanyProfile) GetIpo() string`

GetIpo returns the Ipo field if non-nil, zero value otherwise.

### GetIpoOk

`func (o *CompanyProfile) GetIpoOk() (*string, bool)`

GetIpoOk returns a tuple with the Ipo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpo

`func (o *CompanyProfile) SetIpo(v string)`

SetIpo sets Ipo field to given value.

### HasIpo

`func (o *CompanyProfile) HasIpo() bool`

HasIpo returns a boolean if a field has been set.

### GetMarketCapitalization

`func (o *CompanyProfile) GetMarketCapitalization() float32`

GetMarketCapitalization returns the MarketCapitalization field if non-nil, zero value otherwise.

### GetMarketCapitalizationOk

`func (o *CompanyProfile) GetMarketCapitalizationOk() (*float32, bool)`

GetMarketCapitalizationOk returns a tuple with the MarketCapitalization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarketCapitalization

`func (o *CompanyProfile) SetMarketCapitalization(v float32)`

SetMarketCapitalization sets MarketCapitalization field to given value.

### HasMarketCapitalization

`func (o *CompanyProfile) HasMarketCapitalization() bool`

HasMarketCapitalization returns a boolean if a field has been set.

### GetShareOutstanding

`func (o *CompanyProfile) GetShareOutstanding() float32`

GetShareOutstanding returns the ShareOutstanding field if non-nil, zero value otherwise.

### GetShareOutstandingOk

`func (o *CompanyProfile) GetShareOutstandingOk() (*float32, bool)`

GetShareOutstandingOk returns a tuple with the ShareOutstanding field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShareOutstanding

`func (o *CompanyProfile) SetShareOutstanding(v float32)`

SetShareOutstanding sets ShareOutstanding field to given value.

### HasShareOutstanding

`func (o *CompanyProfile) HasShareOutstanding() bool`

HasShareOutstanding returns a boolean if a field has been set.

### GetEmployeeTotal

`func (o *CompanyProfile) GetEmployeeTotal() int64`

GetEmployeeTotal returns the EmployeeTotal field if non-nil, zero value otherwise.

### GetEmployeeTotalOk

`func (o *CompanyProfile) GetEmployeeTotalOk() (*int64, bool)`

GetEmployeeTotalOk returns a tuple with the EmployeeTotal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmployeeTotal

`func (o *CompanyProfile) SetEmployeeTotal(v int64)`

SetEmployeeTotal sets EmployeeTotal field to given value.

### HasEmployeeTotal

`func (o *CompanyProfile) HasEmployeeTotal() bool`

HasEmployeeTotal returns a boolean if a field has been set.

### GetLogo

`func (o *CompanyProfile) GetLogo() string`

GetLogo returns the Logo field if non-nil, zero value otherwise.

### GetLogoOk

`func (o *CompanyProfile) GetLogoOk() (*string, bool)`

GetLogoOk returns a tuple with the Logo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogo

`func (o *CompanyProfile) SetLogo(v string)`

SetLogo sets Logo field to given value.

### HasLogo

`func (o *CompanyProfile) HasLogo() bool`

HasLogo returns a boolean if a field has been set.

### GetFinnhubIndustry

`func (o *CompanyProfile) GetFinnhubIndustry() string`

GetFinnhubIndustry returns the FinnhubIndustry field if non-nil, zero value otherwise.

### GetFinnhubIndustryOk

`func (o *CompanyProfile) GetFinnhubIndustryOk() (*string, bool)`

GetFinnhubIndustryOk returns a tuple with the FinnhubIndustry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFinnhubIndustry

`func (o *CompanyProfile) SetFinnhubIndustry(v string)`

SetFinnhubIndustry sets FinnhubIndustry field to given value.

### HasFinnhubIndustry

`func (o *CompanyProfile) HasFinnhubIndustry() bool`

HasFinnhubIndustry returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


