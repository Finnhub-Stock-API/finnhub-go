# CompanyExecutive

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Symbol** | Pointer to **string** | Company symbol. | [optional] 
**Executive** | Pointer to [**[]Company**](Company.md) | Array of company&#39;s executives and members of the Board. | [optional] 

## Methods

### NewCompanyExecutive

`func NewCompanyExecutive() *CompanyExecutive`

NewCompanyExecutive instantiates a new CompanyExecutive object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCompanyExecutiveWithDefaults

`func NewCompanyExecutiveWithDefaults() *CompanyExecutive`

NewCompanyExecutiveWithDefaults instantiates a new CompanyExecutive object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSymbol

`func (o *CompanyExecutive) GetSymbol() string`

GetSymbol returns the Symbol field if non-nil, zero value otherwise.

### GetSymbolOk

`func (o *CompanyExecutive) GetSymbolOk() (*string, bool)`

GetSymbolOk returns a tuple with the Symbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSymbol

`func (o *CompanyExecutive) SetSymbol(v string)`

SetSymbol sets Symbol field to given value.

### HasSymbol

`func (o *CompanyExecutive) HasSymbol() bool`

HasSymbol returns a boolean if a field has been set.

### GetExecutive

`func (o *CompanyExecutive) GetExecutive() []Company`

GetExecutive returns the Executive field if non-nil, zero value otherwise.

### GetExecutiveOk

`func (o *CompanyExecutive) GetExecutiveOk() (*[]Company, bool)`

GetExecutiveOk returns a tuple with the Executive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecutive

`func (o *CompanyExecutive) SetExecutive(v []Company)`

SetExecutive sets Executive field to given value.

### HasExecutive

`func (o *CompanyExecutive) HasExecutive() bool`

HasExecutive returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


