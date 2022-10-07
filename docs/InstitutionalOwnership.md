# InstitutionalOwnership

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Symbol** | Pointer to **string** | Symbol. | [optional] 
**Cusip** | Pointer to **string** | Cusip. | [optional] 
**Data** | Pointer to [**[]InstitutionalOwnershipGroup**](InstitutionalOwnershipGroup.md) | Array of institutional investors. | [optional] 

## Methods

### NewInstitutionalOwnership

`func NewInstitutionalOwnership() *InstitutionalOwnership`

NewInstitutionalOwnership instantiates a new InstitutionalOwnership object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInstitutionalOwnershipWithDefaults

`func NewInstitutionalOwnershipWithDefaults() *InstitutionalOwnership`

NewInstitutionalOwnershipWithDefaults instantiates a new InstitutionalOwnership object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSymbol

`func (o *InstitutionalOwnership) GetSymbol() string`

GetSymbol returns the Symbol field if non-nil, zero value otherwise.

### GetSymbolOk

`func (o *InstitutionalOwnership) GetSymbolOk() (*string, bool)`

GetSymbolOk returns a tuple with the Symbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSymbol

`func (o *InstitutionalOwnership) SetSymbol(v string)`

SetSymbol sets Symbol field to given value.

### HasSymbol

`func (o *InstitutionalOwnership) HasSymbol() bool`

HasSymbol returns a boolean if a field has been set.

### GetCusip

`func (o *InstitutionalOwnership) GetCusip() string`

GetCusip returns the Cusip field if non-nil, zero value otherwise.

### GetCusipOk

`func (o *InstitutionalOwnership) GetCusipOk() (*string, bool)`

GetCusipOk returns a tuple with the Cusip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCusip

`func (o *InstitutionalOwnership) SetCusip(v string)`

SetCusip sets Cusip field to given value.

### HasCusip

`func (o *InstitutionalOwnership) HasCusip() bool`

HasCusip returns a boolean if a field has been set.

### GetData

`func (o *InstitutionalOwnership) GetData() []InstitutionalOwnershipGroup`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *InstitutionalOwnership) GetDataOk() (*[]InstitutionalOwnershipGroup, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *InstitutionalOwnership) SetData(v []InstitutionalOwnershipGroup)`

SetData sets Data field to given value.

### HasData

`func (o *InstitutionalOwnership) HasData() bool`

HasData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


