# MutualFundProfile

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Symbol** | Pointer to **string** | Symbol. | [optional] 
**Profile** | Pointer to [**MutualFundProfileData**](MutualFundProfileData.md) |  | [optional] 

## Methods

### NewMutualFundProfile

`func NewMutualFundProfile() *MutualFundProfile`

NewMutualFundProfile instantiates a new MutualFundProfile object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMutualFundProfileWithDefaults

`func NewMutualFundProfileWithDefaults() *MutualFundProfile`

NewMutualFundProfileWithDefaults instantiates a new MutualFundProfile object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSymbol

`func (o *MutualFundProfile) GetSymbol() string`

GetSymbol returns the Symbol field if non-nil, zero value otherwise.

### GetSymbolOk

`func (o *MutualFundProfile) GetSymbolOk() (*string, bool)`

GetSymbolOk returns a tuple with the Symbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSymbol

`func (o *MutualFundProfile) SetSymbol(v string)`

SetSymbol sets Symbol field to given value.

### HasSymbol

`func (o *MutualFundProfile) HasSymbol() bool`

HasSymbol returns a boolean if a field has been set.

### GetProfile

`func (o *MutualFundProfile) GetProfile() MutualFundProfileData`

GetProfile returns the Profile field if non-nil, zero value otherwise.

### GetProfileOk

`func (o *MutualFundProfile) GetProfileOk() (*MutualFundProfileData, bool)`

GetProfileOk returns a tuple with the Profile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfile

`func (o *MutualFundProfile) SetProfile(v MutualFundProfileData)`

SetProfile sets Profile field to given value.

### HasProfile

`func (o *MutualFundProfile) HasProfile() bool`

HasProfile returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


