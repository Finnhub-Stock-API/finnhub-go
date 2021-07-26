# FundOwnershipInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | Investor&#39;s name. | [optional] 
**Share** | Pointer to **int64** | Number of shares held by the investor. | [optional] 
**Change** | Pointer to **int64** | Number of share changed (net buy or sell) from the last period. | [optional] 
**FilingDate** | Pointer to **string** | Filing date. | [optional] 
**PortfolioPercent** | Pointer to **float32** | Percent of the fund&#39;s portfolio comprised of the company&#39;s share. | [optional] 

## Methods

### NewFundOwnershipInfo

`func NewFundOwnershipInfo() *FundOwnershipInfo`

NewFundOwnershipInfo instantiates a new FundOwnershipInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFundOwnershipInfoWithDefaults

`func NewFundOwnershipInfoWithDefaults() *FundOwnershipInfo`

NewFundOwnershipInfoWithDefaults instantiates a new FundOwnershipInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *FundOwnershipInfo) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *FundOwnershipInfo) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *FundOwnershipInfo) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *FundOwnershipInfo) HasName() bool`

HasName returns a boolean if a field has been set.

### GetShare

`func (o *FundOwnershipInfo) GetShare() int64`

GetShare returns the Share field if non-nil, zero value otherwise.

### GetShareOk

`func (o *FundOwnershipInfo) GetShareOk() (*int64, bool)`

GetShareOk returns a tuple with the Share field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShare

`func (o *FundOwnershipInfo) SetShare(v int64)`

SetShare sets Share field to given value.

### HasShare

`func (o *FundOwnershipInfo) HasShare() bool`

HasShare returns a boolean if a field has been set.

### GetChange

`func (o *FundOwnershipInfo) GetChange() int64`

GetChange returns the Change field if non-nil, zero value otherwise.

### GetChangeOk

`func (o *FundOwnershipInfo) GetChangeOk() (*int64, bool)`

GetChangeOk returns a tuple with the Change field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChange

`func (o *FundOwnershipInfo) SetChange(v int64)`

SetChange sets Change field to given value.

### HasChange

`func (o *FundOwnershipInfo) HasChange() bool`

HasChange returns a boolean if a field has been set.

### GetFilingDate

`func (o *FundOwnershipInfo) GetFilingDate() string`

GetFilingDate returns the FilingDate field if non-nil, zero value otherwise.

### GetFilingDateOk

`func (o *FundOwnershipInfo) GetFilingDateOk() (*string, bool)`

GetFilingDateOk returns a tuple with the FilingDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilingDate

`func (o *FundOwnershipInfo) SetFilingDate(v string)`

SetFilingDate sets FilingDate field to given value.

### HasFilingDate

`func (o *FundOwnershipInfo) HasFilingDate() bool`

HasFilingDate returns a boolean if a field has been set.

### GetPortfolioPercent

`func (o *FundOwnershipInfo) GetPortfolioPercent() float32`

GetPortfolioPercent returns the PortfolioPercent field if non-nil, zero value otherwise.

### GetPortfolioPercentOk

`func (o *FundOwnershipInfo) GetPortfolioPercentOk() (*float32, bool)`

GetPortfolioPercentOk returns a tuple with the PortfolioPercent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortfolioPercent

`func (o *FundOwnershipInfo) SetPortfolioPercent(v float32)`

SetPortfolioPercent sets PortfolioPercent field to given value.

### HasPortfolioPercent

`func (o *FundOwnershipInfo) HasPortfolioPercent() bool`

HasPortfolioPercent returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


