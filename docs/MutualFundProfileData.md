# MutualFundProfileData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | Name | [optional] 
**Category** | Pointer to **string** | Fund&#39;s category. | [optional] 
**InvestmentSegment** | Pointer to **string** | Investment Segment. | [optional] 
**TotalNav** | Pointer to **float32** | NAV. | [optional] 
**ExpenseRatio** | Pointer to **float32** | Expense ratio. | [optional] 
**Benchmark** | Pointer to **string** | Index benchmark. | [optional] 
**InceptionDate** | Pointer to **string** | Inception date. | [optional] 
**Description** | Pointer to **string** | Fund&#39;s description. | [optional] 
**FundFamily** | Pointer to **string** | Fund Family. | [optional] 
**Manager** | Pointer to **string** | Fund&#39;s managers. | [optional] 
**Status** | Pointer to **string** | Status. | [optional] 
**Beta** | Pointer to **float32** | Beta. | [optional] 
**DeferredLoad** | Pointer to **float32** | Deferred load. | [optional] 
**Fee12b1** | Pointer to **float32** | 12B-1 fee. | [optional] 
**FrontLoad** | Pointer to **float32** | Front Load. | [optional] 
**IraMinInvestment** | Pointer to **float32** | IRA minimum investment. | [optional] 
**Isin** | Pointer to **string** | ISIN. | [optional] 
**Cusip** | Pointer to **string** | CUSIP. | [optional] 
**MaxRedemptionFee** | Pointer to **float32** | Max redemption fee. | [optional] 
**StandardMinInvestment** | Pointer to **float32** | Minimum investment for standard accounts. | [optional] 
**Turnover** | Pointer to **float32** | Turnover. | [optional] 

## Methods

### NewMutualFundProfileData

`func NewMutualFundProfileData() *MutualFundProfileData`

NewMutualFundProfileData instantiates a new MutualFundProfileData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMutualFundProfileDataWithDefaults

`func NewMutualFundProfileDataWithDefaults() *MutualFundProfileData`

NewMutualFundProfileDataWithDefaults instantiates a new MutualFundProfileData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *MutualFundProfileData) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *MutualFundProfileData) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *MutualFundProfileData) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *MutualFundProfileData) HasName() bool`

HasName returns a boolean if a field has been set.

### GetCategory

`func (o *MutualFundProfileData) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *MutualFundProfileData) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *MutualFundProfileData) SetCategory(v string)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *MutualFundProfileData) HasCategory() bool`

HasCategory returns a boolean if a field has been set.

### GetInvestmentSegment

`func (o *MutualFundProfileData) GetInvestmentSegment() string`

GetInvestmentSegment returns the InvestmentSegment field if non-nil, zero value otherwise.

### GetInvestmentSegmentOk

`func (o *MutualFundProfileData) GetInvestmentSegmentOk() (*string, bool)`

GetInvestmentSegmentOk returns a tuple with the InvestmentSegment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvestmentSegment

`func (o *MutualFundProfileData) SetInvestmentSegment(v string)`

SetInvestmentSegment sets InvestmentSegment field to given value.

### HasInvestmentSegment

`func (o *MutualFundProfileData) HasInvestmentSegment() bool`

HasInvestmentSegment returns a boolean if a field has been set.

### GetTotalNav

`func (o *MutualFundProfileData) GetTotalNav() float32`

GetTotalNav returns the TotalNav field if non-nil, zero value otherwise.

### GetTotalNavOk

`func (o *MutualFundProfileData) GetTotalNavOk() (*float32, bool)`

GetTotalNavOk returns a tuple with the TotalNav field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalNav

`func (o *MutualFundProfileData) SetTotalNav(v float32)`

SetTotalNav sets TotalNav field to given value.

### HasTotalNav

`func (o *MutualFundProfileData) HasTotalNav() bool`

HasTotalNav returns a boolean if a field has been set.

### GetExpenseRatio

`func (o *MutualFundProfileData) GetExpenseRatio() float32`

GetExpenseRatio returns the ExpenseRatio field if non-nil, zero value otherwise.

### GetExpenseRatioOk

`func (o *MutualFundProfileData) GetExpenseRatioOk() (*float32, bool)`

GetExpenseRatioOk returns a tuple with the ExpenseRatio field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpenseRatio

`func (o *MutualFundProfileData) SetExpenseRatio(v float32)`

SetExpenseRatio sets ExpenseRatio field to given value.

### HasExpenseRatio

`func (o *MutualFundProfileData) HasExpenseRatio() bool`

HasExpenseRatio returns a boolean if a field has been set.

### GetBenchmark

`func (o *MutualFundProfileData) GetBenchmark() string`

GetBenchmark returns the Benchmark field if non-nil, zero value otherwise.

### GetBenchmarkOk

`func (o *MutualFundProfileData) GetBenchmarkOk() (*string, bool)`

GetBenchmarkOk returns a tuple with the Benchmark field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBenchmark

`func (o *MutualFundProfileData) SetBenchmark(v string)`

SetBenchmark sets Benchmark field to given value.

### HasBenchmark

`func (o *MutualFundProfileData) HasBenchmark() bool`

HasBenchmark returns a boolean if a field has been set.

### GetInceptionDate

`func (o *MutualFundProfileData) GetInceptionDate() string`

GetInceptionDate returns the InceptionDate field if non-nil, zero value otherwise.

### GetInceptionDateOk

`func (o *MutualFundProfileData) GetInceptionDateOk() (*string, bool)`

GetInceptionDateOk returns a tuple with the InceptionDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInceptionDate

`func (o *MutualFundProfileData) SetInceptionDate(v string)`

SetInceptionDate sets InceptionDate field to given value.

### HasInceptionDate

`func (o *MutualFundProfileData) HasInceptionDate() bool`

HasInceptionDate returns a boolean if a field has been set.

### GetDescription

`func (o *MutualFundProfileData) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *MutualFundProfileData) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *MutualFundProfileData) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *MutualFundProfileData) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetFundFamily

`func (o *MutualFundProfileData) GetFundFamily() string`

GetFundFamily returns the FundFamily field if non-nil, zero value otherwise.

### GetFundFamilyOk

`func (o *MutualFundProfileData) GetFundFamilyOk() (*string, bool)`

GetFundFamilyOk returns a tuple with the FundFamily field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFundFamily

`func (o *MutualFundProfileData) SetFundFamily(v string)`

SetFundFamily sets FundFamily field to given value.

### HasFundFamily

`func (o *MutualFundProfileData) HasFundFamily() bool`

HasFundFamily returns a boolean if a field has been set.

### GetManager

`func (o *MutualFundProfileData) GetManager() string`

GetManager returns the Manager field if non-nil, zero value otherwise.

### GetManagerOk

`func (o *MutualFundProfileData) GetManagerOk() (*string, bool)`

GetManagerOk returns a tuple with the Manager field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManager

`func (o *MutualFundProfileData) SetManager(v string)`

SetManager sets Manager field to given value.

### HasManager

`func (o *MutualFundProfileData) HasManager() bool`

HasManager returns a boolean if a field has been set.

### GetStatus

`func (o *MutualFundProfileData) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *MutualFundProfileData) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *MutualFundProfileData) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *MutualFundProfileData) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetBeta

`func (o *MutualFundProfileData) GetBeta() float32`

GetBeta returns the Beta field if non-nil, zero value otherwise.

### GetBetaOk

`func (o *MutualFundProfileData) GetBetaOk() (*float32, bool)`

GetBetaOk returns a tuple with the Beta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBeta

`func (o *MutualFundProfileData) SetBeta(v float32)`

SetBeta sets Beta field to given value.

### HasBeta

`func (o *MutualFundProfileData) HasBeta() bool`

HasBeta returns a boolean if a field has been set.

### GetDeferredLoad

`func (o *MutualFundProfileData) GetDeferredLoad() float32`

GetDeferredLoad returns the DeferredLoad field if non-nil, zero value otherwise.

### GetDeferredLoadOk

`func (o *MutualFundProfileData) GetDeferredLoadOk() (*float32, bool)`

GetDeferredLoadOk returns a tuple with the DeferredLoad field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeferredLoad

`func (o *MutualFundProfileData) SetDeferredLoad(v float32)`

SetDeferredLoad sets DeferredLoad field to given value.

### HasDeferredLoad

`func (o *MutualFundProfileData) HasDeferredLoad() bool`

HasDeferredLoad returns a boolean if a field has been set.

### GetFee12b1

`func (o *MutualFundProfileData) GetFee12b1() float32`

GetFee12b1 returns the Fee12b1 field if non-nil, zero value otherwise.

### GetFee12b1Ok

`func (o *MutualFundProfileData) GetFee12b1Ok() (*float32, bool)`

GetFee12b1Ok returns a tuple with the Fee12b1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFee12b1

`func (o *MutualFundProfileData) SetFee12b1(v float32)`

SetFee12b1 sets Fee12b1 field to given value.

### HasFee12b1

`func (o *MutualFundProfileData) HasFee12b1() bool`

HasFee12b1 returns a boolean if a field has been set.

### GetFrontLoad

`func (o *MutualFundProfileData) GetFrontLoad() float32`

GetFrontLoad returns the FrontLoad field if non-nil, zero value otherwise.

### GetFrontLoadOk

`func (o *MutualFundProfileData) GetFrontLoadOk() (*float32, bool)`

GetFrontLoadOk returns a tuple with the FrontLoad field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrontLoad

`func (o *MutualFundProfileData) SetFrontLoad(v float32)`

SetFrontLoad sets FrontLoad field to given value.

### HasFrontLoad

`func (o *MutualFundProfileData) HasFrontLoad() bool`

HasFrontLoad returns a boolean if a field has been set.

### GetIraMinInvestment

`func (o *MutualFundProfileData) GetIraMinInvestment() float32`

GetIraMinInvestment returns the IraMinInvestment field if non-nil, zero value otherwise.

### GetIraMinInvestmentOk

`func (o *MutualFundProfileData) GetIraMinInvestmentOk() (*float32, bool)`

GetIraMinInvestmentOk returns a tuple with the IraMinInvestment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIraMinInvestment

`func (o *MutualFundProfileData) SetIraMinInvestment(v float32)`

SetIraMinInvestment sets IraMinInvestment field to given value.

### HasIraMinInvestment

`func (o *MutualFundProfileData) HasIraMinInvestment() bool`

HasIraMinInvestment returns a boolean if a field has been set.

### GetIsin

`func (o *MutualFundProfileData) GetIsin() string`

GetIsin returns the Isin field if non-nil, zero value otherwise.

### GetIsinOk

`func (o *MutualFundProfileData) GetIsinOk() (*string, bool)`

GetIsinOk returns a tuple with the Isin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsin

`func (o *MutualFundProfileData) SetIsin(v string)`

SetIsin sets Isin field to given value.

### HasIsin

`func (o *MutualFundProfileData) HasIsin() bool`

HasIsin returns a boolean if a field has been set.

### GetCusip

`func (o *MutualFundProfileData) GetCusip() string`

GetCusip returns the Cusip field if non-nil, zero value otherwise.

### GetCusipOk

`func (o *MutualFundProfileData) GetCusipOk() (*string, bool)`

GetCusipOk returns a tuple with the Cusip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCusip

`func (o *MutualFundProfileData) SetCusip(v string)`

SetCusip sets Cusip field to given value.

### HasCusip

`func (o *MutualFundProfileData) HasCusip() bool`

HasCusip returns a boolean if a field has been set.

### GetMaxRedemptionFee

`func (o *MutualFundProfileData) GetMaxRedemptionFee() float32`

GetMaxRedemptionFee returns the MaxRedemptionFee field if non-nil, zero value otherwise.

### GetMaxRedemptionFeeOk

`func (o *MutualFundProfileData) GetMaxRedemptionFeeOk() (*float32, bool)`

GetMaxRedemptionFeeOk returns a tuple with the MaxRedemptionFee field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxRedemptionFee

`func (o *MutualFundProfileData) SetMaxRedemptionFee(v float32)`

SetMaxRedemptionFee sets MaxRedemptionFee field to given value.

### HasMaxRedemptionFee

`func (o *MutualFundProfileData) HasMaxRedemptionFee() bool`

HasMaxRedemptionFee returns a boolean if a field has been set.

### GetStandardMinInvestment

`func (o *MutualFundProfileData) GetStandardMinInvestment() float32`

GetStandardMinInvestment returns the StandardMinInvestment field if non-nil, zero value otherwise.

### GetStandardMinInvestmentOk

`func (o *MutualFundProfileData) GetStandardMinInvestmentOk() (*float32, bool)`

GetStandardMinInvestmentOk returns a tuple with the StandardMinInvestment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStandardMinInvestment

`func (o *MutualFundProfileData) SetStandardMinInvestment(v float32)`

SetStandardMinInvestment sets StandardMinInvestment field to given value.

### HasStandardMinInvestment

`func (o *MutualFundProfileData) HasStandardMinInvestment() bool`

HasStandardMinInvestment returns a boolean if a field has been set.

### GetTurnover

`func (o *MutualFundProfileData) GetTurnover() float32`

GetTurnover returns the Turnover field if non-nil, zero value otherwise.

### GetTurnoverOk

`func (o *MutualFundProfileData) GetTurnoverOk() (*float32, bool)`

GetTurnoverOk returns a tuple with the Turnover field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTurnover

`func (o *MutualFundProfileData) SetTurnover(v float32)`

SetTurnover sets Turnover field to given value.

### HasTurnover

`func (o *MutualFundProfileData) HasTurnover() bool`

HasTurnover returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


