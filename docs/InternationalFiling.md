# InternationalFiling

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Symbol** | Pointer to **string** | Symbol. | [optional] 
**CompanyName** | Pointer to **string** | Company name. | [optional] 
**FiledDate** | Pointer to **time.Time** | Filed date &lt;code&gt;%Y-%m-%d %H:%M:%S&lt;/code&gt;. | [optional] 
**Category** | Pointer to **string** | Category. | [optional] 
**Title** | Pointer to **string** | Document&#39;s title. | [optional] 
**Description** | Pointer to **string** | Document&#39;s description. | [optional] 
**Url** | Pointer to **string** | Url. | [optional] 
**Language** | Pointer to **string** | Language. | [optional] 
**Country** | Pointer to **string** | Country. | [optional] 

## Methods

### NewInternationalFiling

`func NewInternationalFiling() *InternationalFiling`

NewInternationalFiling instantiates a new InternationalFiling object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInternationalFilingWithDefaults

`func NewInternationalFilingWithDefaults() *InternationalFiling`

NewInternationalFilingWithDefaults instantiates a new InternationalFiling object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSymbol

`func (o *InternationalFiling) GetSymbol() string`

GetSymbol returns the Symbol field if non-nil, zero value otherwise.

### GetSymbolOk

`func (o *InternationalFiling) GetSymbolOk() (*string, bool)`

GetSymbolOk returns a tuple with the Symbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSymbol

`func (o *InternationalFiling) SetSymbol(v string)`

SetSymbol sets Symbol field to given value.

### HasSymbol

`func (o *InternationalFiling) HasSymbol() bool`

HasSymbol returns a boolean if a field has been set.

### GetCompanyName

`func (o *InternationalFiling) GetCompanyName() string`

GetCompanyName returns the CompanyName field if non-nil, zero value otherwise.

### GetCompanyNameOk

`func (o *InternationalFiling) GetCompanyNameOk() (*string, bool)`

GetCompanyNameOk returns a tuple with the CompanyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyName

`func (o *InternationalFiling) SetCompanyName(v string)`

SetCompanyName sets CompanyName field to given value.

### HasCompanyName

`func (o *InternationalFiling) HasCompanyName() bool`

HasCompanyName returns a boolean if a field has been set.

### GetFiledDate

`func (o *InternationalFiling) GetFiledDate() time.Time`

GetFiledDate returns the FiledDate field if non-nil, zero value otherwise.

### GetFiledDateOk

`func (o *InternationalFiling) GetFiledDateOk() (*time.Time, bool)`

GetFiledDateOk returns a tuple with the FiledDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFiledDate

`func (o *InternationalFiling) SetFiledDate(v time.Time)`

SetFiledDate sets FiledDate field to given value.

### HasFiledDate

`func (o *InternationalFiling) HasFiledDate() bool`

HasFiledDate returns a boolean if a field has been set.

### GetCategory

`func (o *InternationalFiling) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *InternationalFiling) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *InternationalFiling) SetCategory(v string)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *InternationalFiling) HasCategory() bool`

HasCategory returns a boolean if a field has been set.

### GetTitle

`func (o *InternationalFiling) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *InternationalFiling) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *InternationalFiling) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *InternationalFiling) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetDescription

`func (o *InternationalFiling) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *InternationalFiling) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *InternationalFiling) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *InternationalFiling) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetUrl

`func (o *InternationalFiling) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *InternationalFiling) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *InternationalFiling) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *InternationalFiling) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetLanguage

`func (o *InternationalFiling) GetLanguage() string`

GetLanguage returns the Language field if non-nil, zero value otherwise.

### GetLanguageOk

`func (o *InternationalFiling) GetLanguageOk() (*string, bool)`

GetLanguageOk returns a tuple with the Language field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanguage

`func (o *InternationalFiling) SetLanguage(v string)`

SetLanguage sets Language field to given value.

### HasLanguage

`func (o *InternationalFiling) HasLanguage() bool`

HasLanguage returns a boolean if a field has been set.

### GetCountry

`func (o *InternationalFiling) GetCountry() string`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *InternationalFiling) GetCountryOk() (*string, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *InternationalFiling) SetCountry(v string)`

SetCountry sets Country field to given value.

### HasCountry

`func (o *InternationalFiling) HasCountry() bool`

HasCountry returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


