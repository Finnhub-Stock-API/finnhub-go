# EmployeeCount

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AtDate** | Pointer to **string** | Date of the reading | [optional] 
**Employee** | Pointer to **float32** | Value | [optional] 

## Methods

### NewEmployeeCount

`func NewEmployeeCount() *EmployeeCount`

NewEmployeeCount instantiates a new EmployeeCount object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEmployeeCountWithDefaults

`func NewEmployeeCountWithDefaults() *EmployeeCount`

NewEmployeeCountWithDefaults instantiates a new EmployeeCount object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAtDate

`func (o *EmployeeCount) GetAtDate() string`

GetAtDate returns the AtDate field if non-nil, zero value otherwise.

### GetAtDateOk

`func (o *EmployeeCount) GetAtDateOk() (*string, bool)`

GetAtDateOk returns a tuple with the AtDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAtDate

`func (o *EmployeeCount) SetAtDate(v string)`

SetAtDate sets AtDate field to given value.

### HasAtDate

`func (o *EmployeeCount) HasAtDate() bool`

HasAtDate returns a boolean if a field has been set.

### GetEmployee

`func (o *EmployeeCount) GetEmployee() float32`

GetEmployee returns the Employee field if non-nil, zero value otherwise.

### GetEmployeeOk

`func (o *EmployeeCount) GetEmployeeOk() (*float32, bool)`

GetEmployeeOk returns a tuple with the Employee field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmployee

`func (o *EmployeeCount) SetEmployee(v float32)`

SetEmployee sets Employee field to given value.

### HasEmployee

`func (o *EmployeeCount) HasEmployee() bool`

HasEmployee returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


