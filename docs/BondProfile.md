# BondProfile

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Isin** | Pointer to **string** | ISIN. | [optional] 
**Cusip** | Pointer to **string** | Cusip. | [optional] 
**Figi** | Pointer to **string** | FIGI. | [optional] 
**Coupon** | Pointer to **float32** | Coupon. | [optional] 
**MaturityDate** | Pointer to **string** | Period. | [optional] 
**OfferingPrice** | Pointer to **float32** | Offering price. | [optional] 
**IssueDate** | Pointer to **string** | Issue date. | [optional] 
**BondType** | Pointer to **string** | Bond type. | [optional] 
**DebtType** | Pointer to **string** | Bond type. | [optional] 
**IndustryGroup** | Pointer to **string** | Industry. | [optional] 
**IndustrySubGroup** | Pointer to **string** | Sub-Industry. | [optional] 
**Asset** | Pointer to **string** | Asset. | [optional] 
**AssetType** | Pointer to **string** | Asset. | [optional] 
**DatedDate** | Pointer to **string** | Dated date. | [optional] 
**FirstCouponDate** | Pointer to **string** | First coupon date. | [optional] 
**OriginalOffering** | Pointer to **float32** | Offering amount. | [optional] 
**AmountOutstanding** | Pointer to **float32** | Outstanding amount. | [optional] 
**PaymentFrequency** | Pointer to **string** | Payment frequency. | [optional] 
**SecurityLevel** | Pointer to **string** | Security level. | [optional] 
**Callable** | Pointer to **bool** | Callable. | [optional] 
**CouponType** | Pointer to **string** | Coupon type. | [optional] 

## Methods

### NewBondProfile

`func NewBondProfile() *BondProfile`

NewBondProfile instantiates a new BondProfile object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBondProfileWithDefaults

`func NewBondProfileWithDefaults() *BondProfile`

NewBondProfileWithDefaults instantiates a new BondProfile object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIsin

`func (o *BondProfile) GetIsin() string`

GetIsin returns the Isin field if non-nil, zero value otherwise.

### GetIsinOk

`func (o *BondProfile) GetIsinOk() (*string, bool)`

GetIsinOk returns a tuple with the Isin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsin

`func (o *BondProfile) SetIsin(v string)`

SetIsin sets Isin field to given value.

### HasIsin

`func (o *BondProfile) HasIsin() bool`

HasIsin returns a boolean if a field has been set.

### GetCusip

`func (o *BondProfile) GetCusip() string`

GetCusip returns the Cusip field if non-nil, zero value otherwise.

### GetCusipOk

`func (o *BondProfile) GetCusipOk() (*string, bool)`

GetCusipOk returns a tuple with the Cusip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCusip

`func (o *BondProfile) SetCusip(v string)`

SetCusip sets Cusip field to given value.

### HasCusip

`func (o *BondProfile) HasCusip() bool`

HasCusip returns a boolean if a field has been set.

### GetFigi

`func (o *BondProfile) GetFigi() string`

GetFigi returns the Figi field if non-nil, zero value otherwise.

### GetFigiOk

`func (o *BondProfile) GetFigiOk() (*string, bool)`

GetFigiOk returns a tuple with the Figi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFigi

`func (o *BondProfile) SetFigi(v string)`

SetFigi sets Figi field to given value.

### HasFigi

`func (o *BondProfile) HasFigi() bool`

HasFigi returns a boolean if a field has been set.

### GetCoupon

`func (o *BondProfile) GetCoupon() float32`

GetCoupon returns the Coupon field if non-nil, zero value otherwise.

### GetCouponOk

`func (o *BondProfile) GetCouponOk() (*float32, bool)`

GetCouponOk returns a tuple with the Coupon field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCoupon

`func (o *BondProfile) SetCoupon(v float32)`

SetCoupon sets Coupon field to given value.

### HasCoupon

`func (o *BondProfile) HasCoupon() bool`

HasCoupon returns a boolean if a field has been set.

### GetMaturityDate

`func (o *BondProfile) GetMaturityDate() string`

GetMaturityDate returns the MaturityDate field if non-nil, zero value otherwise.

### GetMaturityDateOk

`func (o *BondProfile) GetMaturityDateOk() (*string, bool)`

GetMaturityDateOk returns a tuple with the MaturityDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaturityDate

`func (o *BondProfile) SetMaturityDate(v string)`

SetMaturityDate sets MaturityDate field to given value.

### HasMaturityDate

`func (o *BondProfile) HasMaturityDate() bool`

HasMaturityDate returns a boolean if a field has been set.

### GetOfferingPrice

`func (o *BondProfile) GetOfferingPrice() float32`

GetOfferingPrice returns the OfferingPrice field if non-nil, zero value otherwise.

### GetOfferingPriceOk

`func (o *BondProfile) GetOfferingPriceOk() (*float32, bool)`

GetOfferingPriceOk returns a tuple with the OfferingPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOfferingPrice

`func (o *BondProfile) SetOfferingPrice(v float32)`

SetOfferingPrice sets OfferingPrice field to given value.

### HasOfferingPrice

`func (o *BondProfile) HasOfferingPrice() bool`

HasOfferingPrice returns a boolean if a field has been set.

### GetIssueDate

`func (o *BondProfile) GetIssueDate() string`

GetIssueDate returns the IssueDate field if non-nil, zero value otherwise.

### GetIssueDateOk

`func (o *BondProfile) GetIssueDateOk() (*string, bool)`

GetIssueDateOk returns a tuple with the IssueDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssueDate

`func (o *BondProfile) SetIssueDate(v string)`

SetIssueDate sets IssueDate field to given value.

### HasIssueDate

`func (o *BondProfile) HasIssueDate() bool`

HasIssueDate returns a boolean if a field has been set.

### GetBondType

`func (o *BondProfile) GetBondType() string`

GetBondType returns the BondType field if non-nil, zero value otherwise.

### GetBondTypeOk

`func (o *BondProfile) GetBondTypeOk() (*string, bool)`

GetBondTypeOk returns a tuple with the BondType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBondType

`func (o *BondProfile) SetBondType(v string)`

SetBondType sets BondType field to given value.

### HasBondType

`func (o *BondProfile) HasBondType() bool`

HasBondType returns a boolean if a field has been set.

### GetDebtType

`func (o *BondProfile) GetDebtType() string`

GetDebtType returns the DebtType field if non-nil, zero value otherwise.

### GetDebtTypeOk

`func (o *BondProfile) GetDebtTypeOk() (*string, bool)`

GetDebtTypeOk returns a tuple with the DebtType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDebtType

`func (o *BondProfile) SetDebtType(v string)`

SetDebtType sets DebtType field to given value.

### HasDebtType

`func (o *BondProfile) HasDebtType() bool`

HasDebtType returns a boolean if a field has been set.

### GetIndustryGroup

`func (o *BondProfile) GetIndustryGroup() string`

GetIndustryGroup returns the IndustryGroup field if non-nil, zero value otherwise.

### GetIndustryGroupOk

`func (o *BondProfile) GetIndustryGroupOk() (*string, bool)`

GetIndustryGroupOk returns a tuple with the IndustryGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndustryGroup

`func (o *BondProfile) SetIndustryGroup(v string)`

SetIndustryGroup sets IndustryGroup field to given value.

### HasIndustryGroup

`func (o *BondProfile) HasIndustryGroup() bool`

HasIndustryGroup returns a boolean if a field has been set.

### GetIndustrySubGroup

`func (o *BondProfile) GetIndustrySubGroup() string`

GetIndustrySubGroup returns the IndustrySubGroup field if non-nil, zero value otherwise.

### GetIndustrySubGroupOk

`func (o *BondProfile) GetIndustrySubGroupOk() (*string, bool)`

GetIndustrySubGroupOk returns a tuple with the IndustrySubGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndustrySubGroup

`func (o *BondProfile) SetIndustrySubGroup(v string)`

SetIndustrySubGroup sets IndustrySubGroup field to given value.

### HasIndustrySubGroup

`func (o *BondProfile) HasIndustrySubGroup() bool`

HasIndustrySubGroup returns a boolean if a field has been set.

### GetAsset

`func (o *BondProfile) GetAsset() string`

GetAsset returns the Asset field if non-nil, zero value otherwise.

### GetAssetOk

`func (o *BondProfile) GetAssetOk() (*string, bool)`

GetAssetOk returns a tuple with the Asset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAsset

`func (o *BondProfile) SetAsset(v string)`

SetAsset sets Asset field to given value.

### HasAsset

`func (o *BondProfile) HasAsset() bool`

HasAsset returns a boolean if a field has been set.

### GetAssetType

`func (o *BondProfile) GetAssetType() string`

GetAssetType returns the AssetType field if non-nil, zero value otherwise.

### GetAssetTypeOk

`func (o *BondProfile) GetAssetTypeOk() (*string, bool)`

GetAssetTypeOk returns a tuple with the AssetType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssetType

`func (o *BondProfile) SetAssetType(v string)`

SetAssetType sets AssetType field to given value.

### HasAssetType

`func (o *BondProfile) HasAssetType() bool`

HasAssetType returns a boolean if a field has been set.

### GetDatedDate

`func (o *BondProfile) GetDatedDate() string`

GetDatedDate returns the DatedDate field if non-nil, zero value otherwise.

### GetDatedDateOk

`func (o *BondProfile) GetDatedDateOk() (*string, bool)`

GetDatedDateOk returns a tuple with the DatedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatedDate

`func (o *BondProfile) SetDatedDate(v string)`

SetDatedDate sets DatedDate field to given value.

### HasDatedDate

`func (o *BondProfile) HasDatedDate() bool`

HasDatedDate returns a boolean if a field has been set.

### GetFirstCouponDate

`func (o *BondProfile) GetFirstCouponDate() string`

GetFirstCouponDate returns the FirstCouponDate field if non-nil, zero value otherwise.

### GetFirstCouponDateOk

`func (o *BondProfile) GetFirstCouponDateOk() (*string, bool)`

GetFirstCouponDateOk returns a tuple with the FirstCouponDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstCouponDate

`func (o *BondProfile) SetFirstCouponDate(v string)`

SetFirstCouponDate sets FirstCouponDate field to given value.

### HasFirstCouponDate

`func (o *BondProfile) HasFirstCouponDate() bool`

HasFirstCouponDate returns a boolean if a field has been set.

### GetOriginalOffering

`func (o *BondProfile) GetOriginalOffering() float32`

GetOriginalOffering returns the OriginalOffering field if non-nil, zero value otherwise.

### GetOriginalOfferingOk

`func (o *BondProfile) GetOriginalOfferingOk() (*float32, bool)`

GetOriginalOfferingOk returns a tuple with the OriginalOffering field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginalOffering

`func (o *BondProfile) SetOriginalOffering(v float32)`

SetOriginalOffering sets OriginalOffering field to given value.

### HasOriginalOffering

`func (o *BondProfile) HasOriginalOffering() bool`

HasOriginalOffering returns a boolean if a field has been set.

### GetAmountOutstanding

`func (o *BondProfile) GetAmountOutstanding() float32`

GetAmountOutstanding returns the AmountOutstanding field if non-nil, zero value otherwise.

### GetAmountOutstandingOk

`func (o *BondProfile) GetAmountOutstandingOk() (*float32, bool)`

GetAmountOutstandingOk returns a tuple with the AmountOutstanding field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmountOutstanding

`func (o *BondProfile) SetAmountOutstanding(v float32)`

SetAmountOutstanding sets AmountOutstanding field to given value.

### HasAmountOutstanding

`func (o *BondProfile) HasAmountOutstanding() bool`

HasAmountOutstanding returns a boolean if a field has been set.

### GetPaymentFrequency

`func (o *BondProfile) GetPaymentFrequency() string`

GetPaymentFrequency returns the PaymentFrequency field if non-nil, zero value otherwise.

### GetPaymentFrequencyOk

`func (o *BondProfile) GetPaymentFrequencyOk() (*string, bool)`

GetPaymentFrequencyOk returns a tuple with the PaymentFrequency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentFrequency

`func (o *BondProfile) SetPaymentFrequency(v string)`

SetPaymentFrequency sets PaymentFrequency field to given value.

### HasPaymentFrequency

`func (o *BondProfile) HasPaymentFrequency() bool`

HasPaymentFrequency returns a boolean if a field has been set.

### GetSecurityLevel

`func (o *BondProfile) GetSecurityLevel() string`

GetSecurityLevel returns the SecurityLevel field if non-nil, zero value otherwise.

### GetSecurityLevelOk

`func (o *BondProfile) GetSecurityLevelOk() (*string, bool)`

GetSecurityLevelOk returns a tuple with the SecurityLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurityLevel

`func (o *BondProfile) SetSecurityLevel(v string)`

SetSecurityLevel sets SecurityLevel field to given value.

### HasSecurityLevel

`func (o *BondProfile) HasSecurityLevel() bool`

HasSecurityLevel returns a boolean if a field has been set.

### GetCallable

`func (o *BondProfile) GetCallable() bool`

GetCallable returns the Callable field if non-nil, zero value otherwise.

### GetCallableOk

`func (o *BondProfile) GetCallableOk() (*bool, bool)`

GetCallableOk returns a tuple with the Callable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallable

`func (o *BondProfile) SetCallable(v bool)`

SetCallable sets Callable field to given value.

### HasCallable

`func (o *BondProfile) HasCallable() bool`

HasCallable returns a boolean if a field has been set.

### GetCouponType

`func (o *BondProfile) GetCouponType() string`

GetCouponType returns the CouponType field if non-nil, zero value otherwise.

### GetCouponTypeOk

`func (o *BondProfile) GetCouponTypeOk() (*string, bool)`

GetCouponTypeOk returns a tuple with the CouponType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCouponType

`func (o *BondProfile) SetCouponType(v string)`

SetCouponType sets CouponType field to given value.

### HasCouponType

`func (o *BondProfile) HasCouponType() bool`

HasCouponType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


