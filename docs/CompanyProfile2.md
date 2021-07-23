# CompanyProfile2

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Country** | Pointer to **string** | Country of company&#39;s headquarter. | [optional] 
**Currency** | Pointer to **string** | Currency used in company filings. | [optional] 
**Exchange** | Pointer to **string** | Listed exchange. | [optional] 
**Name** | Pointer to **string** | Company name. | [optional] 
**Ticker** | Pointer to **string** | Company symbol/ticker as used on the listed exchange. | [optional] 
**Ipo** | Pointer to **string** | IPO date. | [optional] 
**MarketCapitalization** | Pointer to **float32** | Market Capitalization. | [optional] 
**ShareOutstanding** | Pointer to **float32** | Number of oustanding shares. | [optional] 
**Logo** | Pointer to **string** | Logo image. | [optional] 
**Phone** | Pointer to **string** | Company phone number. | [optional] 
**Weburl** | Pointer to **string** | Company website. | [optional] 
**FinnhubIndustry** | Pointer to **string** | Finnhub industry classification. | [optional] 

## Methods

### NewCompanyProfile2

`func NewCompanyProfile2() *CompanyProfile2`

NewCompanyProfile2 instantiates a new CompanyProfile2 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCompanyProfile2WithDefaults

`func NewCompanyProfile2WithDefaults() *CompanyProfile2`

NewCompanyProfile2WithDefaults instantiates a new CompanyProfile2 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCountry

`func (o *CompanyProfile2) GetCountry() string`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *CompanyProfile2) GetCountryOk() (*string, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *CompanyProfile2) SetCountry(v string)`

SetCountry sets Country field to given value.

### HasCountry

`func (o *CompanyProfile2) HasCountry() bool`

HasCountry returns a boolean if a field has been set.

### GetCurrency

`func (o *CompanyProfile2) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *CompanyProfile2) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *CompanyProfile2) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *CompanyProfile2) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetExchange

`func (o *CompanyProfile2) GetExchange() string`

GetExchange returns the Exchange field if non-nil, zero value otherwise.

### GetExchangeOk

`func (o *CompanyProfile2) GetExchangeOk() (*string, bool)`

GetExchangeOk returns a tuple with the Exchange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExchange

`func (o *CompanyProfile2) SetExchange(v string)`

SetExchange sets Exchange field to given value.

### HasExchange

`func (o *CompanyProfile2) HasExchange() bool`

HasExchange returns a boolean if a field has been set.

### GetName

`func (o *CompanyProfile2) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CompanyProfile2) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CompanyProfile2) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CompanyProfile2) HasName() bool`

HasName returns a boolean if a field has been set.

### GetTicker

`func (o *CompanyProfile2) GetTicker() string`

GetTicker returns the Ticker field if non-nil, zero value otherwise.

### GetTickerOk

`func (o *CompanyProfile2) GetTickerOk() (*string, bool)`

GetTickerOk returns a tuple with the Ticker field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTicker

`func (o *CompanyProfile2) SetTicker(v string)`

SetTicker sets Ticker field to given value.

### HasTicker

`func (o *CompanyProfile2) HasTicker() bool`

HasTicker returns a boolean if a field has been set.

### GetIpo

`func (o *CompanyProfile2) GetIpo() string`

GetIpo returns the Ipo field if non-nil, zero value otherwise.

### GetIpoOk

`func (o *CompanyProfile2) GetIpoOk() (*string, bool)`

GetIpoOk returns a tuple with the Ipo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpo

`func (o *CompanyProfile2) SetIpo(v string)`

SetIpo sets Ipo field to given value.

### HasIpo

`func (o *CompanyProfile2) HasIpo() bool`

HasIpo returns a boolean if a field has been set.

### GetMarketCapitalization

`func (o *CompanyProfile2) GetMarketCapitalization() float32`

GetMarketCapitalization returns the MarketCapitalization field if non-nil, zero value otherwise.

### GetMarketCapitalizationOk

`func (o *CompanyProfile2) GetMarketCapitalizationOk() (*float32, bool)`

GetMarketCapitalizationOk returns a tuple with the MarketCapitalization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarketCapitalization

`func (o *CompanyProfile2) SetMarketCapitalization(v float32)`

SetMarketCapitalization sets MarketCapitalization field to given value.

### HasMarketCapitalization

`func (o *CompanyProfile2) HasMarketCapitalization() bool`

HasMarketCapitalization returns a boolean if a field has been set.

### GetShareOutstanding

`func (o *CompanyProfile2) GetShareOutstanding() float32`

GetShareOutstanding returns the ShareOutstanding field if non-nil, zero value otherwise.

### GetShareOutstandingOk

`func (o *CompanyProfile2) GetShareOutstandingOk() (*float32, bool)`

GetShareOutstandingOk returns a tuple with the ShareOutstanding field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShareOutstanding

`func (o *CompanyProfile2) SetShareOutstanding(v float32)`

SetShareOutstanding sets ShareOutstanding field to given value.

### HasShareOutstanding

`func (o *CompanyProfile2) HasShareOutstanding() bool`

HasShareOutstanding returns a boolean if a field has been set.

### GetLogo

`func (o *CompanyProfile2) GetLogo() string`

GetLogo returns the Logo field if non-nil, zero value otherwise.

### GetLogoOk

`func (o *CompanyProfile2) GetLogoOk() (*string, bool)`

GetLogoOk returns a tuple with the Logo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogo

`func (o *CompanyProfile2) SetLogo(v string)`

SetLogo sets Logo field to given value.

### HasLogo

`func (o *CompanyProfile2) HasLogo() bool`

HasLogo returns a boolean if a field has been set.

### GetPhone

`func (o *CompanyProfile2) GetPhone() string`

GetPhone returns the Phone field if non-nil, zero value otherwise.

### GetPhoneOk

`func (o *CompanyProfile2) GetPhoneOk() (*string, bool)`

GetPhoneOk returns a tuple with the Phone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhone

`func (o *CompanyProfile2) SetPhone(v string)`

SetPhone sets Phone field to given value.

### HasPhone

`func (o *CompanyProfile2) HasPhone() bool`

HasPhone returns a boolean if a field has been set.

### GetWeburl

`func (o *CompanyProfile2) GetWeburl() string`

GetWeburl returns the Weburl field if non-nil, zero value otherwise.

### GetWeburlOk

`func (o *CompanyProfile2) GetWeburlOk() (*string, bool)`

GetWeburlOk returns a tuple with the Weburl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeburl

`func (o *CompanyProfile2) SetWeburl(v string)`

SetWeburl sets Weburl field to given value.

### HasWeburl

`func (o *CompanyProfile2) HasWeburl() bool`

HasWeburl returns a boolean if a field has been set.

### GetFinnhubIndustry

`func (o *CompanyProfile2) GetFinnhubIndustry() string`

GetFinnhubIndustry returns the FinnhubIndustry field if non-nil, zero value otherwise.

### GetFinnhubIndustryOk

`func (o *CompanyProfile2) GetFinnhubIndustryOk() (*string, bool)`

GetFinnhubIndustryOk returns a tuple with the FinnhubIndustry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFinnhubIndustry

`func (o *CompanyProfile2) SetFinnhubIndustry(v string)`

SetFinnhubIndustry sets FinnhubIndustry field to given value.

### HasFinnhubIndustry

`func (o *CompanyProfile2) HasFinnhubIndustry() bool`

HasFinnhubIndustry returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


