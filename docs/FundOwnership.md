# FundOwnership

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Symbol** | Pointer to **string** | Symbol of the company. | [optional] 
**Ownership** | Pointer to [**[]FundOwnershipInfo**](FundOwnershipInfo.md) | Array of investors with detailed information about their holdings. | [optional] 

## Methods

### NewFundOwnership

`func NewFundOwnership() *FundOwnership`

NewFundOwnership instantiates a new FundOwnership object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFundOwnershipWithDefaults

`func NewFundOwnershipWithDefaults() *FundOwnership`

NewFundOwnershipWithDefaults instantiates a new FundOwnership object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSymbol

`func (o *FundOwnership) GetSymbol() string`

GetSymbol returns the Symbol field if non-nil, zero value otherwise.

### GetSymbolOk

`func (o *FundOwnership) GetSymbolOk() (*string, bool)`

GetSymbolOk returns a tuple with the Symbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSymbol

`func (o *FundOwnership) SetSymbol(v string)`

SetSymbol sets Symbol field to given value.

### HasSymbol

`func (o *FundOwnership) HasSymbol() bool`

HasSymbol returns a boolean if a field has been set.

### GetOwnership

`func (o *FundOwnership) GetOwnership() []FundOwnershipInfo`

GetOwnership returns the Ownership field if non-nil, zero value otherwise.

### GetOwnershipOk

`func (o *FundOwnership) GetOwnershipOk() (*[]FundOwnershipInfo, bool)`

GetOwnershipOk returns a tuple with the Ownership field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnership

`func (o *FundOwnership) SetOwnership(v []FundOwnershipInfo)`

SetOwnership sets Ownership field to given value.

### HasOwnership

`func (o *FundOwnership) HasOwnership() bool`

HasOwnership returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


