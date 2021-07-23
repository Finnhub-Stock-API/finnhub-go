# MutualFundsProfile

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Symbol** | Pointer to **string** | Symbol. | [optional] 
**Profile** | Pointer to [**ETFProfileData**](ETFProfileData.md) |  | [optional] 

## Methods

### NewMutualFundsProfile

`func NewMutualFundsProfile() *MutualFundsProfile`

NewMutualFundsProfile instantiates a new MutualFundsProfile object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMutualFundsProfileWithDefaults

`func NewMutualFundsProfileWithDefaults() *MutualFundsProfile`

NewMutualFundsProfileWithDefaults instantiates a new MutualFundsProfile object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSymbol

`func (o *MutualFundsProfile) GetSymbol() string`

GetSymbol returns the Symbol field if non-nil, zero value otherwise.

### GetSymbolOk

`func (o *MutualFundsProfile) GetSymbolOk() (*string, bool)`

GetSymbolOk returns a tuple with the Symbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSymbol

`func (o *MutualFundsProfile) SetSymbol(v string)`

SetSymbol sets Symbol field to given value.

### HasSymbol

`func (o *MutualFundsProfile) HasSymbol() bool`

HasSymbol returns a boolean if a field has been set.

### GetProfile

`func (o *MutualFundsProfile) GetProfile() ETFProfileData`

GetProfile returns the Profile field if non-nil, zero value otherwise.

### GetProfileOk

`func (o *MutualFundsProfile) GetProfileOk() (*ETFProfileData, bool)`

GetProfileOk returns a tuple with the Profile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfile

`func (o *MutualFundsProfile) SetProfile(v ETFProfileData)`

SetProfile sets Profile field to given value.

### HasProfile

`func (o *MutualFundsProfile) HasProfile() bool`

HasProfile returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


