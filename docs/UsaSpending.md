# UsaSpending

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Symbol** | Pointer to **string** | Symbol. | [optional] 
**RecipientName** | Pointer to **string** | Company&#39;s name. | [optional] 
**RecipientParentName** | Pointer to **string** | Company&#39;s name. | [optional] 
**AwardDescription** | Pointer to **string** | Description. | [optional] 
**Country** | Pointer to **string** | Recipient&#39;s country. | [optional] 
**ActionDate** | Pointer to **string** | Period. | [optional] 
**TotalValue** | Pointer to **float32** | Income reported by lobbying firms. | [optional] 
**PerformanceStartDate** | Pointer to **string** | Performance start date. | [optional] 
**PerformanceEndDate** | Pointer to **string** | Performance end date. | [optional] 
**AwardingAgencyName** | Pointer to **string** | Award agency. | [optional] 
**AwardingSubAgencyName** | Pointer to **string** | Award sub-agency. | [optional] 
**AwardingOfficeName** | Pointer to **string** | Award office name. | [optional] 
**PerformanceCountry** | Pointer to **string** | Performance country. | [optional] 
**PerformanceCity** | Pointer to **string** | Performance city. | [optional] 
**PerformanceCounty** | Pointer to **string** | Performance county. | [optional] 
**PerformanceState** | Pointer to **string** | Performance state. | [optional] 
**PerformanceZipCode** | Pointer to **string** | Performance zip code. | [optional] 
**PerformanceCongressionalDistrict** | Pointer to **string** | Performance congressional district. | [optional] 
**NaicsCode** | Pointer to **string** | NAICS code. | [optional] 
**Permalink** | Pointer to **string** | Permalink. | [optional] 

## Methods

### NewUsaSpending

`func NewUsaSpending() *UsaSpending`

NewUsaSpending instantiates a new UsaSpending object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUsaSpendingWithDefaults

`func NewUsaSpendingWithDefaults() *UsaSpending`

NewUsaSpendingWithDefaults instantiates a new UsaSpending object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSymbol

`func (o *UsaSpending) GetSymbol() string`

GetSymbol returns the Symbol field if non-nil, zero value otherwise.

### GetSymbolOk

`func (o *UsaSpending) GetSymbolOk() (*string, bool)`

GetSymbolOk returns a tuple with the Symbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSymbol

`func (o *UsaSpending) SetSymbol(v string)`

SetSymbol sets Symbol field to given value.

### HasSymbol

`func (o *UsaSpending) HasSymbol() bool`

HasSymbol returns a boolean if a field has been set.

### GetRecipientName

`func (o *UsaSpending) GetRecipientName() string`

GetRecipientName returns the RecipientName field if non-nil, zero value otherwise.

### GetRecipientNameOk

`func (o *UsaSpending) GetRecipientNameOk() (*string, bool)`

GetRecipientNameOk returns a tuple with the RecipientName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecipientName

`func (o *UsaSpending) SetRecipientName(v string)`

SetRecipientName sets RecipientName field to given value.

### HasRecipientName

`func (o *UsaSpending) HasRecipientName() bool`

HasRecipientName returns a boolean if a field has been set.

### GetRecipientParentName

`func (o *UsaSpending) GetRecipientParentName() string`

GetRecipientParentName returns the RecipientParentName field if non-nil, zero value otherwise.

### GetRecipientParentNameOk

`func (o *UsaSpending) GetRecipientParentNameOk() (*string, bool)`

GetRecipientParentNameOk returns a tuple with the RecipientParentName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecipientParentName

`func (o *UsaSpending) SetRecipientParentName(v string)`

SetRecipientParentName sets RecipientParentName field to given value.

### HasRecipientParentName

`func (o *UsaSpending) HasRecipientParentName() bool`

HasRecipientParentName returns a boolean if a field has been set.

### GetAwardDescription

`func (o *UsaSpending) GetAwardDescription() string`

GetAwardDescription returns the AwardDescription field if non-nil, zero value otherwise.

### GetAwardDescriptionOk

`func (o *UsaSpending) GetAwardDescriptionOk() (*string, bool)`

GetAwardDescriptionOk returns a tuple with the AwardDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAwardDescription

`func (o *UsaSpending) SetAwardDescription(v string)`

SetAwardDescription sets AwardDescription field to given value.

### HasAwardDescription

`func (o *UsaSpending) HasAwardDescription() bool`

HasAwardDescription returns a boolean if a field has been set.

### GetCountry

`func (o *UsaSpending) GetCountry() string`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *UsaSpending) GetCountryOk() (*string, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *UsaSpending) SetCountry(v string)`

SetCountry sets Country field to given value.

### HasCountry

`func (o *UsaSpending) HasCountry() bool`

HasCountry returns a boolean if a field has been set.

### GetActionDate

`func (o *UsaSpending) GetActionDate() string`

GetActionDate returns the ActionDate field if non-nil, zero value otherwise.

### GetActionDateOk

`func (o *UsaSpending) GetActionDateOk() (*string, bool)`

GetActionDateOk returns a tuple with the ActionDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActionDate

`func (o *UsaSpending) SetActionDate(v string)`

SetActionDate sets ActionDate field to given value.

### HasActionDate

`func (o *UsaSpending) HasActionDate() bool`

HasActionDate returns a boolean if a field has been set.

### GetTotalValue

`func (o *UsaSpending) GetTotalValue() float32`

GetTotalValue returns the TotalValue field if non-nil, zero value otherwise.

### GetTotalValueOk

`func (o *UsaSpending) GetTotalValueOk() (*float32, bool)`

GetTotalValueOk returns a tuple with the TotalValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalValue

`func (o *UsaSpending) SetTotalValue(v float32)`

SetTotalValue sets TotalValue field to given value.

### HasTotalValue

`func (o *UsaSpending) HasTotalValue() bool`

HasTotalValue returns a boolean if a field has been set.

### GetPerformanceStartDate

`func (o *UsaSpending) GetPerformanceStartDate() string`

GetPerformanceStartDate returns the PerformanceStartDate field if non-nil, zero value otherwise.

### GetPerformanceStartDateOk

`func (o *UsaSpending) GetPerformanceStartDateOk() (*string, bool)`

GetPerformanceStartDateOk returns a tuple with the PerformanceStartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPerformanceStartDate

`func (o *UsaSpending) SetPerformanceStartDate(v string)`

SetPerformanceStartDate sets PerformanceStartDate field to given value.

### HasPerformanceStartDate

`func (o *UsaSpending) HasPerformanceStartDate() bool`

HasPerformanceStartDate returns a boolean if a field has been set.

### GetPerformanceEndDate

`func (o *UsaSpending) GetPerformanceEndDate() string`

GetPerformanceEndDate returns the PerformanceEndDate field if non-nil, zero value otherwise.

### GetPerformanceEndDateOk

`func (o *UsaSpending) GetPerformanceEndDateOk() (*string, bool)`

GetPerformanceEndDateOk returns a tuple with the PerformanceEndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPerformanceEndDate

`func (o *UsaSpending) SetPerformanceEndDate(v string)`

SetPerformanceEndDate sets PerformanceEndDate field to given value.

### HasPerformanceEndDate

`func (o *UsaSpending) HasPerformanceEndDate() bool`

HasPerformanceEndDate returns a boolean if a field has been set.

### GetAwardingAgencyName

`func (o *UsaSpending) GetAwardingAgencyName() string`

GetAwardingAgencyName returns the AwardingAgencyName field if non-nil, zero value otherwise.

### GetAwardingAgencyNameOk

`func (o *UsaSpending) GetAwardingAgencyNameOk() (*string, bool)`

GetAwardingAgencyNameOk returns a tuple with the AwardingAgencyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAwardingAgencyName

`func (o *UsaSpending) SetAwardingAgencyName(v string)`

SetAwardingAgencyName sets AwardingAgencyName field to given value.

### HasAwardingAgencyName

`func (o *UsaSpending) HasAwardingAgencyName() bool`

HasAwardingAgencyName returns a boolean if a field has been set.

### GetAwardingSubAgencyName

`func (o *UsaSpending) GetAwardingSubAgencyName() string`

GetAwardingSubAgencyName returns the AwardingSubAgencyName field if non-nil, zero value otherwise.

### GetAwardingSubAgencyNameOk

`func (o *UsaSpending) GetAwardingSubAgencyNameOk() (*string, bool)`

GetAwardingSubAgencyNameOk returns a tuple with the AwardingSubAgencyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAwardingSubAgencyName

`func (o *UsaSpending) SetAwardingSubAgencyName(v string)`

SetAwardingSubAgencyName sets AwardingSubAgencyName field to given value.

### HasAwardingSubAgencyName

`func (o *UsaSpending) HasAwardingSubAgencyName() bool`

HasAwardingSubAgencyName returns a boolean if a field has been set.

### GetAwardingOfficeName

`func (o *UsaSpending) GetAwardingOfficeName() string`

GetAwardingOfficeName returns the AwardingOfficeName field if non-nil, zero value otherwise.

### GetAwardingOfficeNameOk

`func (o *UsaSpending) GetAwardingOfficeNameOk() (*string, bool)`

GetAwardingOfficeNameOk returns a tuple with the AwardingOfficeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAwardingOfficeName

`func (o *UsaSpending) SetAwardingOfficeName(v string)`

SetAwardingOfficeName sets AwardingOfficeName field to given value.

### HasAwardingOfficeName

`func (o *UsaSpending) HasAwardingOfficeName() bool`

HasAwardingOfficeName returns a boolean if a field has been set.

### GetPerformanceCountry

`func (o *UsaSpending) GetPerformanceCountry() string`

GetPerformanceCountry returns the PerformanceCountry field if non-nil, zero value otherwise.

### GetPerformanceCountryOk

`func (o *UsaSpending) GetPerformanceCountryOk() (*string, bool)`

GetPerformanceCountryOk returns a tuple with the PerformanceCountry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPerformanceCountry

`func (o *UsaSpending) SetPerformanceCountry(v string)`

SetPerformanceCountry sets PerformanceCountry field to given value.

### HasPerformanceCountry

`func (o *UsaSpending) HasPerformanceCountry() bool`

HasPerformanceCountry returns a boolean if a field has been set.

### GetPerformanceCity

`func (o *UsaSpending) GetPerformanceCity() string`

GetPerformanceCity returns the PerformanceCity field if non-nil, zero value otherwise.

### GetPerformanceCityOk

`func (o *UsaSpending) GetPerformanceCityOk() (*string, bool)`

GetPerformanceCityOk returns a tuple with the PerformanceCity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPerformanceCity

`func (o *UsaSpending) SetPerformanceCity(v string)`

SetPerformanceCity sets PerformanceCity field to given value.

### HasPerformanceCity

`func (o *UsaSpending) HasPerformanceCity() bool`

HasPerformanceCity returns a boolean if a field has been set.

### GetPerformanceCounty

`func (o *UsaSpending) GetPerformanceCounty() string`

GetPerformanceCounty returns the PerformanceCounty field if non-nil, zero value otherwise.

### GetPerformanceCountyOk

`func (o *UsaSpending) GetPerformanceCountyOk() (*string, bool)`

GetPerformanceCountyOk returns a tuple with the PerformanceCounty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPerformanceCounty

`func (o *UsaSpending) SetPerformanceCounty(v string)`

SetPerformanceCounty sets PerformanceCounty field to given value.

### HasPerformanceCounty

`func (o *UsaSpending) HasPerformanceCounty() bool`

HasPerformanceCounty returns a boolean if a field has been set.

### GetPerformanceState

`func (o *UsaSpending) GetPerformanceState() string`

GetPerformanceState returns the PerformanceState field if non-nil, zero value otherwise.

### GetPerformanceStateOk

`func (o *UsaSpending) GetPerformanceStateOk() (*string, bool)`

GetPerformanceStateOk returns a tuple with the PerformanceState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPerformanceState

`func (o *UsaSpending) SetPerformanceState(v string)`

SetPerformanceState sets PerformanceState field to given value.

### HasPerformanceState

`func (o *UsaSpending) HasPerformanceState() bool`

HasPerformanceState returns a boolean if a field has been set.

### GetPerformanceZipCode

`func (o *UsaSpending) GetPerformanceZipCode() string`

GetPerformanceZipCode returns the PerformanceZipCode field if non-nil, zero value otherwise.

### GetPerformanceZipCodeOk

`func (o *UsaSpending) GetPerformanceZipCodeOk() (*string, bool)`

GetPerformanceZipCodeOk returns a tuple with the PerformanceZipCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPerformanceZipCode

`func (o *UsaSpending) SetPerformanceZipCode(v string)`

SetPerformanceZipCode sets PerformanceZipCode field to given value.

### HasPerformanceZipCode

`func (o *UsaSpending) HasPerformanceZipCode() bool`

HasPerformanceZipCode returns a boolean if a field has been set.

### GetPerformanceCongressionalDistrict

`func (o *UsaSpending) GetPerformanceCongressionalDistrict() string`

GetPerformanceCongressionalDistrict returns the PerformanceCongressionalDistrict field if non-nil, zero value otherwise.

### GetPerformanceCongressionalDistrictOk

`func (o *UsaSpending) GetPerformanceCongressionalDistrictOk() (*string, bool)`

GetPerformanceCongressionalDistrictOk returns a tuple with the PerformanceCongressionalDistrict field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPerformanceCongressionalDistrict

`func (o *UsaSpending) SetPerformanceCongressionalDistrict(v string)`

SetPerformanceCongressionalDistrict sets PerformanceCongressionalDistrict field to given value.

### HasPerformanceCongressionalDistrict

`func (o *UsaSpending) HasPerformanceCongressionalDistrict() bool`

HasPerformanceCongressionalDistrict returns a boolean if a field has been set.

### GetNaicsCode

`func (o *UsaSpending) GetNaicsCode() string`

GetNaicsCode returns the NaicsCode field if non-nil, zero value otherwise.

### GetNaicsCodeOk

`func (o *UsaSpending) GetNaicsCodeOk() (*string, bool)`

GetNaicsCodeOk returns a tuple with the NaicsCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNaicsCode

`func (o *UsaSpending) SetNaicsCode(v string)`

SetNaicsCode sets NaicsCode field to given value.

### HasNaicsCode

`func (o *UsaSpending) HasNaicsCode() bool`

HasNaicsCode returns a boolean if a field has been set.

### GetPermalink

`func (o *UsaSpending) GetPermalink() string`

GetPermalink returns the Permalink field if non-nil, zero value otherwise.

### GetPermalinkOk

`func (o *UsaSpending) GetPermalinkOk() (*string, bool)`

GetPermalinkOk returns a tuple with the Permalink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermalink

`func (o *UsaSpending) SetPermalink(v string)`

SetPermalink sets Permalink field to given value.

### HasPermalink

`func (o *UsaSpending) HasPermalink() bool`

HasPermalink returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


