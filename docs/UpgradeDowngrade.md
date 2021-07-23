# UpgradeDowngrade

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Symbol** | Pointer to **string** | Company symbol. | [optional] 
**GradeTime** | Pointer to **int64** | Upgrade/downgrade time in UNIX timestamp. | [optional] 
**FromGrade** | Pointer to **string** | From grade. | [optional] 
**ToGrade** | Pointer to **string** | To grade. | [optional] 
**Company** | Pointer to **string** | Company/analyst who did the upgrade/downgrade. | [optional] 
**Action** | Pointer to **string** | Action can take any of the following values: &lt;code&gt;up(upgrade), down(downgrade), main(maintains), init(initiate), reit(reiterate)&lt;/code&gt;. | [optional] 

## Methods

### NewUpgradeDowngrade

`func NewUpgradeDowngrade() *UpgradeDowngrade`

NewUpgradeDowngrade instantiates a new UpgradeDowngrade object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpgradeDowngradeWithDefaults

`func NewUpgradeDowngradeWithDefaults() *UpgradeDowngrade`

NewUpgradeDowngradeWithDefaults instantiates a new UpgradeDowngrade object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSymbol

`func (o *UpgradeDowngrade) GetSymbol() string`

GetSymbol returns the Symbol field if non-nil, zero value otherwise.

### GetSymbolOk

`func (o *UpgradeDowngrade) GetSymbolOk() (*string, bool)`

GetSymbolOk returns a tuple with the Symbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSymbol

`func (o *UpgradeDowngrade) SetSymbol(v string)`

SetSymbol sets Symbol field to given value.

### HasSymbol

`func (o *UpgradeDowngrade) HasSymbol() bool`

HasSymbol returns a boolean if a field has been set.

### GetGradeTime

`func (o *UpgradeDowngrade) GetGradeTime() int64`

GetGradeTime returns the GradeTime field if non-nil, zero value otherwise.

### GetGradeTimeOk

`func (o *UpgradeDowngrade) GetGradeTimeOk() (*int64, bool)`

GetGradeTimeOk returns a tuple with the GradeTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGradeTime

`func (o *UpgradeDowngrade) SetGradeTime(v int64)`

SetGradeTime sets GradeTime field to given value.

### HasGradeTime

`func (o *UpgradeDowngrade) HasGradeTime() bool`

HasGradeTime returns a boolean if a field has been set.

### GetFromGrade

`func (o *UpgradeDowngrade) GetFromGrade() string`

GetFromGrade returns the FromGrade field if non-nil, zero value otherwise.

### GetFromGradeOk

`func (o *UpgradeDowngrade) GetFromGradeOk() (*string, bool)`

GetFromGradeOk returns a tuple with the FromGrade field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFromGrade

`func (o *UpgradeDowngrade) SetFromGrade(v string)`

SetFromGrade sets FromGrade field to given value.

### HasFromGrade

`func (o *UpgradeDowngrade) HasFromGrade() bool`

HasFromGrade returns a boolean if a field has been set.

### GetToGrade

`func (o *UpgradeDowngrade) GetToGrade() string`

GetToGrade returns the ToGrade field if non-nil, zero value otherwise.

### GetToGradeOk

`func (o *UpgradeDowngrade) GetToGradeOk() (*string, bool)`

GetToGradeOk returns a tuple with the ToGrade field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToGrade

`func (o *UpgradeDowngrade) SetToGrade(v string)`

SetToGrade sets ToGrade field to given value.

### HasToGrade

`func (o *UpgradeDowngrade) HasToGrade() bool`

HasToGrade returns a boolean if a field has been set.

### GetCompany

`func (o *UpgradeDowngrade) GetCompany() string`

GetCompany returns the Company field if non-nil, zero value otherwise.

### GetCompanyOk

`func (o *UpgradeDowngrade) GetCompanyOk() (*string, bool)`

GetCompanyOk returns a tuple with the Company field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompany

`func (o *UpgradeDowngrade) SetCompany(v string)`

SetCompany sets Company field to given value.

### HasCompany

`func (o *UpgradeDowngrade) HasCompany() bool`

HasCompany returns a boolean if a field has been set.

### GetAction

`func (o *UpgradeDowngrade) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *UpgradeDowngrade) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *UpgradeDowngrade) SetAction(v string)`

SetAction sets Action field to given value.

### HasAction

`func (o *UpgradeDowngrade) HasAction() bool`

HasAction returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


