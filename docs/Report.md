# Report

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccessNumber** | Pointer to **string** | Access number. | [optional] 
**Symbol** | Pointer to **string** | Symbol. | [optional] 
**Cik** | Pointer to **string** | CIK. | [optional] 
**Year** | Pointer to **int64** | Year. | [optional] 
**Quarter** | Pointer to **int64** | Quarter. | [optional] 
**Form** | Pointer to **string** | Form type. | [optional] 
**StartDate** | Pointer to **string** | Period start date &lt;code&gt;%Y-%m-%d %H:%M:%S&lt;/code&gt;. | [optional] 
**EndDate** | Pointer to **string** | Period end date &lt;code&gt;%Y-%m-%d %H:%M:%S&lt;/code&gt;. | [optional] 
**FiledDate** | Pointer to **string** | Filed date &lt;code&gt;%Y-%m-%d %H:%M:%S&lt;/code&gt;. | [optional] 
**AcceptedDate** | Pointer to **string** | Accepted date &lt;code&gt;%Y-%m-%d %H:%M:%S&lt;/code&gt;. | [optional] 
**Report** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewReport

`func NewReport() *Report`

NewReport instantiates a new Report object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReportWithDefaults

`func NewReportWithDefaults() *Report`

NewReportWithDefaults instantiates a new Report object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccessNumber

`func (o *Report) GetAccessNumber() string`

GetAccessNumber returns the AccessNumber field if non-nil, zero value otherwise.

### GetAccessNumberOk

`func (o *Report) GetAccessNumberOk() (*string, bool)`

GetAccessNumberOk returns a tuple with the AccessNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessNumber

`func (o *Report) SetAccessNumber(v string)`

SetAccessNumber sets AccessNumber field to given value.

### HasAccessNumber

`func (o *Report) HasAccessNumber() bool`

HasAccessNumber returns a boolean if a field has been set.

### GetSymbol

`func (o *Report) GetSymbol() string`

GetSymbol returns the Symbol field if non-nil, zero value otherwise.

### GetSymbolOk

`func (o *Report) GetSymbolOk() (*string, bool)`

GetSymbolOk returns a tuple with the Symbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSymbol

`func (o *Report) SetSymbol(v string)`

SetSymbol sets Symbol field to given value.

### HasSymbol

`func (o *Report) HasSymbol() bool`

HasSymbol returns a boolean if a field has been set.

### GetCik

`func (o *Report) GetCik() string`

GetCik returns the Cik field if non-nil, zero value otherwise.

### GetCikOk

`func (o *Report) GetCikOk() (*string, bool)`

GetCikOk returns a tuple with the Cik field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCik

`func (o *Report) SetCik(v string)`

SetCik sets Cik field to given value.

### HasCik

`func (o *Report) HasCik() bool`

HasCik returns a boolean if a field has been set.

### GetYear

`func (o *Report) GetYear() int64`

GetYear returns the Year field if non-nil, zero value otherwise.

### GetYearOk

`func (o *Report) GetYearOk() (*int64, bool)`

GetYearOk returns a tuple with the Year field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetYear

`func (o *Report) SetYear(v int64)`

SetYear sets Year field to given value.

### HasYear

`func (o *Report) HasYear() bool`

HasYear returns a boolean if a field has been set.

### GetQuarter

`func (o *Report) GetQuarter() int64`

GetQuarter returns the Quarter field if non-nil, zero value otherwise.

### GetQuarterOk

`func (o *Report) GetQuarterOk() (*int64, bool)`

GetQuarterOk returns a tuple with the Quarter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuarter

`func (o *Report) SetQuarter(v int64)`

SetQuarter sets Quarter field to given value.

### HasQuarter

`func (o *Report) HasQuarter() bool`

HasQuarter returns a boolean if a field has been set.

### GetForm

`func (o *Report) GetForm() string`

GetForm returns the Form field if non-nil, zero value otherwise.

### GetFormOk

`func (o *Report) GetFormOk() (*string, bool)`

GetFormOk returns a tuple with the Form field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForm

`func (o *Report) SetForm(v string)`

SetForm sets Form field to given value.

### HasForm

`func (o *Report) HasForm() bool`

HasForm returns a boolean if a field has been set.

### GetStartDate

`func (o *Report) GetStartDate() string`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *Report) GetStartDateOk() (*string, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *Report) SetStartDate(v string)`

SetStartDate sets StartDate field to given value.

### HasStartDate

`func (o *Report) HasStartDate() bool`

HasStartDate returns a boolean if a field has been set.

### GetEndDate

`func (o *Report) GetEndDate() string`

GetEndDate returns the EndDate field if non-nil, zero value otherwise.

### GetEndDateOk

`func (o *Report) GetEndDateOk() (*string, bool)`

GetEndDateOk returns a tuple with the EndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDate

`func (o *Report) SetEndDate(v string)`

SetEndDate sets EndDate field to given value.

### HasEndDate

`func (o *Report) HasEndDate() bool`

HasEndDate returns a boolean if a field has been set.

### GetFiledDate

`func (o *Report) GetFiledDate() string`

GetFiledDate returns the FiledDate field if non-nil, zero value otherwise.

### GetFiledDateOk

`func (o *Report) GetFiledDateOk() (*string, bool)`

GetFiledDateOk returns a tuple with the FiledDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFiledDate

`func (o *Report) SetFiledDate(v string)`

SetFiledDate sets FiledDate field to given value.

### HasFiledDate

`func (o *Report) HasFiledDate() bool`

HasFiledDate returns a boolean if a field has been set.

### GetAcceptedDate

`func (o *Report) GetAcceptedDate() string`

GetAcceptedDate returns the AcceptedDate field if non-nil, zero value otherwise.

### GetAcceptedDateOk

`func (o *Report) GetAcceptedDateOk() (*string, bool)`

GetAcceptedDateOk returns a tuple with the AcceptedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcceptedDate

`func (o *Report) SetAcceptedDate(v string)`

SetAcceptedDate sets AcceptedDate field to given value.

### HasAcceptedDate

`func (o *Report) HasAcceptedDate() bool`

HasAcceptedDate returns a boolean if a field has been set.

### GetReport

`func (o *Report) GetReport() map[string]interface{}`

GetReport returns the Report field if non-nil, zero value otherwise.

### GetReportOk

`func (o *Report) GetReportOk() (*map[string]interface{}, bool)`

GetReportOk returns a tuple with the Report field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReport

`func (o *Report) SetReport(v map[string]interface{})`

SetReport sets Report field to given value.

### HasReport

`func (o *Report) HasReport() bool`

HasReport returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


