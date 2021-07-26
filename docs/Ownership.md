# Ownership

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Symbol** | Pointer to **string** | Symbol of the company. | [optional] 
**Ownership** | Pointer to [**[]OwnershipInfo**](OwnershipInfo.md) | Array of investors with detailed information about their holdings. | [optional] 

## Methods

### NewOwnership

`func NewOwnership() *Ownership`

NewOwnership instantiates a new Ownership object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOwnershipWithDefaults

`func NewOwnershipWithDefaults() *Ownership`

NewOwnershipWithDefaults instantiates a new Ownership object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSymbol

`func (o *Ownership) GetSymbol() string`

GetSymbol returns the Symbol field if non-nil, zero value otherwise.

### GetSymbolOk

`func (o *Ownership) GetSymbolOk() (*string, bool)`

GetSymbolOk returns a tuple with the Symbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSymbol

`func (o *Ownership) SetSymbol(v string)`

SetSymbol sets Symbol field to given value.

### HasSymbol

`func (o *Ownership) HasSymbol() bool`

HasSymbol returns a boolean if a field has been set.

### GetOwnership

`func (o *Ownership) GetOwnership() []OwnershipInfo`

GetOwnership returns the Ownership field if non-nil, zero value otherwise.

### GetOwnershipOk

`func (o *Ownership) GetOwnershipOk() (*[]OwnershipInfo, bool)`

GetOwnershipOk returns a tuple with the Ownership field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnership

`func (o *Ownership) SetOwnership(v []OwnershipInfo)`

SetOwnership sets Ownership field to given value.

### HasOwnership

`func (o *Ownership) HasOwnership() bool`

HasOwnership returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


