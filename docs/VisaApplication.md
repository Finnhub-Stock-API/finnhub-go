# VisaApplication

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Year** | Pointer to **int64** | Year. | [optional] 
**Quarter** | Pointer to **int64** | Quarter. | [optional] 
**Symbol** | Pointer to **string** | Symbol. | [optional] 
**CaseNumber** | Pointer to **string** | Case number. | [optional] 
**CaseStatus** | Pointer to **string** | Case status. | [optional] 
**ReceivedDate** | Pointer to **string** | Received date. | [optional] 
**VisaClass** | Pointer to **string** | Visa class. | [optional] 
**JobTitle** | Pointer to **string** | Job Title. | [optional] 
**SocCode** | Pointer to **string** | SOC Code. A list of SOC code can be found &lt;a href&#x3D;\&quot;https://www.bls.gov/oes/current/oes_stru.htm\&quot; target&#x3D;\&quot;_blank\&quot;&gt;here&lt;/a&gt;. | [optional] 
**FullTimePosition** | Pointer to **string** | Full-time position flag. | [optional] 
**BeginDate** | Pointer to **string** | Job&#39;s start date. | [optional] 
**EndDate** | Pointer to **string** | Job&#39;s end date. | [optional] 
**EmployerName** | Pointer to **string** | Company&#39;s name. | [optional] 
**WorksiteAddress** | Pointer to **string** | Worksite address. | [optional] 
**WorksiteCity** | Pointer to **string** | Worksite city. | [optional] 
**WorksiteCounty** | Pointer to **string** | Worksite county. | [optional] 
**WorksiteState** | Pointer to **string** | Worksite state. | [optional] 
**WorksitePostalCode** | Pointer to **string** | Worksite postal code. | [optional] 
**WageRangeFrom** | Pointer to **float32** | Wage range from. | [optional] 
**WageRangeTo** | Pointer to **float32** | Wage range to. | [optional] 
**WaveUnitOfPay** | Pointer to **string** | Wage unit of pay. | [optional] 
**WageLevel** | Pointer to **string** | Wage level. | [optional] 
**H1bDependent** | Pointer to **string** | H1B dependent flag. | [optional] 

## Methods

### NewVisaApplication

`func NewVisaApplication() *VisaApplication`

NewVisaApplication instantiates a new VisaApplication object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVisaApplicationWithDefaults

`func NewVisaApplicationWithDefaults() *VisaApplication`

NewVisaApplicationWithDefaults instantiates a new VisaApplication object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetYear

`func (o *VisaApplication) GetYear() int64`

GetYear returns the Year field if non-nil, zero value otherwise.

### GetYearOk

`func (o *VisaApplication) GetYearOk() (*int64, bool)`

GetYearOk returns a tuple with the Year field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetYear

`func (o *VisaApplication) SetYear(v int64)`

SetYear sets Year field to given value.

### HasYear

`func (o *VisaApplication) HasYear() bool`

HasYear returns a boolean if a field has been set.

### GetQuarter

`func (o *VisaApplication) GetQuarter() int64`

GetQuarter returns the Quarter field if non-nil, zero value otherwise.

### GetQuarterOk

`func (o *VisaApplication) GetQuarterOk() (*int64, bool)`

GetQuarterOk returns a tuple with the Quarter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuarter

`func (o *VisaApplication) SetQuarter(v int64)`

SetQuarter sets Quarter field to given value.

### HasQuarter

`func (o *VisaApplication) HasQuarter() bool`

HasQuarter returns a boolean if a field has been set.

### GetSymbol

`func (o *VisaApplication) GetSymbol() string`

GetSymbol returns the Symbol field if non-nil, zero value otherwise.

### GetSymbolOk

`func (o *VisaApplication) GetSymbolOk() (*string, bool)`

GetSymbolOk returns a tuple with the Symbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSymbol

`func (o *VisaApplication) SetSymbol(v string)`

SetSymbol sets Symbol field to given value.

### HasSymbol

`func (o *VisaApplication) HasSymbol() bool`

HasSymbol returns a boolean if a field has been set.

### GetCaseNumber

`func (o *VisaApplication) GetCaseNumber() string`

GetCaseNumber returns the CaseNumber field if non-nil, zero value otherwise.

### GetCaseNumberOk

`func (o *VisaApplication) GetCaseNumberOk() (*string, bool)`

GetCaseNumberOk returns a tuple with the CaseNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaseNumber

`func (o *VisaApplication) SetCaseNumber(v string)`

SetCaseNumber sets CaseNumber field to given value.

### HasCaseNumber

`func (o *VisaApplication) HasCaseNumber() bool`

HasCaseNumber returns a boolean if a field has been set.

### GetCaseStatus

`func (o *VisaApplication) GetCaseStatus() string`

GetCaseStatus returns the CaseStatus field if non-nil, zero value otherwise.

### GetCaseStatusOk

`func (o *VisaApplication) GetCaseStatusOk() (*string, bool)`

GetCaseStatusOk returns a tuple with the CaseStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaseStatus

`func (o *VisaApplication) SetCaseStatus(v string)`

SetCaseStatus sets CaseStatus field to given value.

### HasCaseStatus

`func (o *VisaApplication) HasCaseStatus() bool`

HasCaseStatus returns a boolean if a field has been set.

### GetReceivedDate

`func (o *VisaApplication) GetReceivedDate() string`

GetReceivedDate returns the ReceivedDate field if non-nil, zero value otherwise.

### GetReceivedDateOk

`func (o *VisaApplication) GetReceivedDateOk() (*string, bool)`

GetReceivedDateOk returns a tuple with the ReceivedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReceivedDate

`func (o *VisaApplication) SetReceivedDate(v string)`

SetReceivedDate sets ReceivedDate field to given value.

### HasReceivedDate

`func (o *VisaApplication) HasReceivedDate() bool`

HasReceivedDate returns a boolean if a field has been set.

### GetVisaClass

`func (o *VisaApplication) GetVisaClass() string`

GetVisaClass returns the VisaClass field if non-nil, zero value otherwise.

### GetVisaClassOk

`func (o *VisaApplication) GetVisaClassOk() (*string, bool)`

GetVisaClassOk returns a tuple with the VisaClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisaClass

`func (o *VisaApplication) SetVisaClass(v string)`

SetVisaClass sets VisaClass field to given value.

### HasVisaClass

`func (o *VisaApplication) HasVisaClass() bool`

HasVisaClass returns a boolean if a field has been set.

### GetJobTitle

`func (o *VisaApplication) GetJobTitle() string`

GetJobTitle returns the JobTitle field if non-nil, zero value otherwise.

### GetJobTitleOk

`func (o *VisaApplication) GetJobTitleOk() (*string, bool)`

GetJobTitleOk returns a tuple with the JobTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobTitle

`func (o *VisaApplication) SetJobTitle(v string)`

SetJobTitle sets JobTitle field to given value.

### HasJobTitle

`func (o *VisaApplication) HasJobTitle() bool`

HasJobTitle returns a boolean if a field has been set.

### GetSocCode

`func (o *VisaApplication) GetSocCode() string`

GetSocCode returns the SocCode field if non-nil, zero value otherwise.

### GetSocCodeOk

`func (o *VisaApplication) GetSocCodeOk() (*string, bool)`

GetSocCodeOk returns a tuple with the SocCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSocCode

`func (o *VisaApplication) SetSocCode(v string)`

SetSocCode sets SocCode field to given value.

### HasSocCode

`func (o *VisaApplication) HasSocCode() bool`

HasSocCode returns a boolean if a field has been set.

### GetFullTimePosition

`func (o *VisaApplication) GetFullTimePosition() string`

GetFullTimePosition returns the FullTimePosition field if non-nil, zero value otherwise.

### GetFullTimePositionOk

`func (o *VisaApplication) GetFullTimePositionOk() (*string, bool)`

GetFullTimePositionOk returns a tuple with the FullTimePosition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFullTimePosition

`func (o *VisaApplication) SetFullTimePosition(v string)`

SetFullTimePosition sets FullTimePosition field to given value.

### HasFullTimePosition

`func (o *VisaApplication) HasFullTimePosition() bool`

HasFullTimePosition returns a boolean if a field has been set.

### GetBeginDate

`func (o *VisaApplication) GetBeginDate() string`

GetBeginDate returns the BeginDate field if non-nil, zero value otherwise.

### GetBeginDateOk

`func (o *VisaApplication) GetBeginDateOk() (*string, bool)`

GetBeginDateOk returns a tuple with the BeginDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBeginDate

`func (o *VisaApplication) SetBeginDate(v string)`

SetBeginDate sets BeginDate field to given value.

### HasBeginDate

`func (o *VisaApplication) HasBeginDate() bool`

HasBeginDate returns a boolean if a field has been set.

### GetEndDate

`func (o *VisaApplication) GetEndDate() string`

GetEndDate returns the EndDate field if non-nil, zero value otherwise.

### GetEndDateOk

`func (o *VisaApplication) GetEndDateOk() (*string, bool)`

GetEndDateOk returns a tuple with the EndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDate

`func (o *VisaApplication) SetEndDate(v string)`

SetEndDate sets EndDate field to given value.

### HasEndDate

`func (o *VisaApplication) HasEndDate() bool`

HasEndDate returns a boolean if a field has been set.

### GetEmployerName

`func (o *VisaApplication) GetEmployerName() string`

GetEmployerName returns the EmployerName field if non-nil, zero value otherwise.

### GetEmployerNameOk

`func (o *VisaApplication) GetEmployerNameOk() (*string, bool)`

GetEmployerNameOk returns a tuple with the EmployerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmployerName

`func (o *VisaApplication) SetEmployerName(v string)`

SetEmployerName sets EmployerName field to given value.

### HasEmployerName

`func (o *VisaApplication) HasEmployerName() bool`

HasEmployerName returns a boolean if a field has been set.

### GetWorksiteAddress

`func (o *VisaApplication) GetWorksiteAddress() string`

GetWorksiteAddress returns the WorksiteAddress field if non-nil, zero value otherwise.

### GetWorksiteAddressOk

`func (o *VisaApplication) GetWorksiteAddressOk() (*string, bool)`

GetWorksiteAddressOk returns a tuple with the WorksiteAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorksiteAddress

`func (o *VisaApplication) SetWorksiteAddress(v string)`

SetWorksiteAddress sets WorksiteAddress field to given value.

### HasWorksiteAddress

`func (o *VisaApplication) HasWorksiteAddress() bool`

HasWorksiteAddress returns a boolean if a field has been set.

### GetWorksiteCity

`func (o *VisaApplication) GetWorksiteCity() string`

GetWorksiteCity returns the WorksiteCity field if non-nil, zero value otherwise.

### GetWorksiteCityOk

`func (o *VisaApplication) GetWorksiteCityOk() (*string, bool)`

GetWorksiteCityOk returns a tuple with the WorksiteCity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorksiteCity

`func (o *VisaApplication) SetWorksiteCity(v string)`

SetWorksiteCity sets WorksiteCity field to given value.

### HasWorksiteCity

`func (o *VisaApplication) HasWorksiteCity() bool`

HasWorksiteCity returns a boolean if a field has been set.

### GetWorksiteCounty

`func (o *VisaApplication) GetWorksiteCounty() string`

GetWorksiteCounty returns the WorksiteCounty field if non-nil, zero value otherwise.

### GetWorksiteCountyOk

`func (o *VisaApplication) GetWorksiteCountyOk() (*string, bool)`

GetWorksiteCountyOk returns a tuple with the WorksiteCounty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorksiteCounty

`func (o *VisaApplication) SetWorksiteCounty(v string)`

SetWorksiteCounty sets WorksiteCounty field to given value.

### HasWorksiteCounty

`func (o *VisaApplication) HasWorksiteCounty() bool`

HasWorksiteCounty returns a boolean if a field has been set.

### GetWorksiteState

`func (o *VisaApplication) GetWorksiteState() string`

GetWorksiteState returns the WorksiteState field if non-nil, zero value otherwise.

### GetWorksiteStateOk

`func (o *VisaApplication) GetWorksiteStateOk() (*string, bool)`

GetWorksiteStateOk returns a tuple with the WorksiteState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorksiteState

`func (o *VisaApplication) SetWorksiteState(v string)`

SetWorksiteState sets WorksiteState field to given value.

### HasWorksiteState

`func (o *VisaApplication) HasWorksiteState() bool`

HasWorksiteState returns a boolean if a field has been set.

### GetWorksitePostalCode

`func (o *VisaApplication) GetWorksitePostalCode() string`

GetWorksitePostalCode returns the WorksitePostalCode field if non-nil, zero value otherwise.

### GetWorksitePostalCodeOk

`func (o *VisaApplication) GetWorksitePostalCodeOk() (*string, bool)`

GetWorksitePostalCodeOk returns a tuple with the WorksitePostalCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorksitePostalCode

`func (o *VisaApplication) SetWorksitePostalCode(v string)`

SetWorksitePostalCode sets WorksitePostalCode field to given value.

### HasWorksitePostalCode

`func (o *VisaApplication) HasWorksitePostalCode() bool`

HasWorksitePostalCode returns a boolean if a field has been set.

### GetWageRangeFrom

`func (o *VisaApplication) GetWageRangeFrom() float32`

GetWageRangeFrom returns the WageRangeFrom field if non-nil, zero value otherwise.

### GetWageRangeFromOk

`func (o *VisaApplication) GetWageRangeFromOk() (*float32, bool)`

GetWageRangeFromOk returns a tuple with the WageRangeFrom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWageRangeFrom

`func (o *VisaApplication) SetWageRangeFrom(v float32)`

SetWageRangeFrom sets WageRangeFrom field to given value.

### HasWageRangeFrom

`func (o *VisaApplication) HasWageRangeFrom() bool`

HasWageRangeFrom returns a boolean if a field has been set.

### GetWageRangeTo

`func (o *VisaApplication) GetWageRangeTo() float32`

GetWageRangeTo returns the WageRangeTo field if non-nil, zero value otherwise.

### GetWageRangeToOk

`func (o *VisaApplication) GetWageRangeToOk() (*float32, bool)`

GetWageRangeToOk returns a tuple with the WageRangeTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWageRangeTo

`func (o *VisaApplication) SetWageRangeTo(v float32)`

SetWageRangeTo sets WageRangeTo field to given value.

### HasWageRangeTo

`func (o *VisaApplication) HasWageRangeTo() bool`

HasWageRangeTo returns a boolean if a field has been set.

### GetWaveUnitOfPay

`func (o *VisaApplication) GetWaveUnitOfPay() string`

GetWaveUnitOfPay returns the WaveUnitOfPay field if non-nil, zero value otherwise.

### GetWaveUnitOfPayOk

`func (o *VisaApplication) GetWaveUnitOfPayOk() (*string, bool)`

GetWaveUnitOfPayOk returns a tuple with the WaveUnitOfPay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWaveUnitOfPay

`func (o *VisaApplication) SetWaveUnitOfPay(v string)`

SetWaveUnitOfPay sets WaveUnitOfPay field to given value.

### HasWaveUnitOfPay

`func (o *VisaApplication) HasWaveUnitOfPay() bool`

HasWaveUnitOfPay returns a boolean if a field has been set.

### GetWageLevel

`func (o *VisaApplication) GetWageLevel() string`

GetWageLevel returns the WageLevel field if non-nil, zero value otherwise.

### GetWageLevelOk

`func (o *VisaApplication) GetWageLevelOk() (*string, bool)`

GetWageLevelOk returns a tuple with the WageLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWageLevel

`func (o *VisaApplication) SetWageLevel(v string)`

SetWageLevel sets WageLevel field to given value.

### HasWageLevel

`func (o *VisaApplication) HasWageLevel() bool`

HasWageLevel returns a boolean if a field has been set.

### GetH1bDependent

`func (o *VisaApplication) GetH1bDependent() string`

GetH1bDependent returns the H1bDependent field if non-nil, zero value otherwise.

### GetH1bDependentOk

`func (o *VisaApplication) GetH1bDependentOk() (*string, bool)`

GetH1bDependentOk returns a tuple with the H1bDependent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetH1bDependent

`func (o *VisaApplication) SetH1bDependent(v string)`

SetH1bDependent sets H1bDependent field to given value.

### HasH1bDependent

`func (o *VisaApplication) HasH1bDependent() bool`

HasH1bDependent returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


