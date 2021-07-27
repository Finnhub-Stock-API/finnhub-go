# ETFsProfile

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Symbol** | Pointer to **string** | Symbol. | [optional] 
**Profile** | Pointer to [**ETFProfileData**](ETFProfileData.md) |  | [optional] 

## Methods

### NewETFsProfile

`func NewETFsProfile() *ETFsProfile`

NewETFsProfile instantiates a new ETFsProfile object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewETFsProfileWithDefaults

`func NewETFsProfileWithDefaults() *ETFsProfile`

NewETFsProfileWithDefaults instantiates a new ETFsProfile object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSymbol

`func (o *ETFsProfile) GetSymbol() string`

GetSymbol returns the Symbol field if non-nil, zero value otherwise.

### GetSymbolOk

`func (o *ETFsProfile) GetSymbolOk() (*string, bool)`

GetSymbolOk returns a tuple with the Symbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSymbol

`func (o *ETFsProfile) SetSymbol(v string)`

SetSymbol sets Symbol field to given value.

### HasSymbol

`func (o *ETFsProfile) HasSymbol() bool`

HasSymbol returns a boolean if a field has been set.

### GetProfile

`func (o *ETFsProfile) GetProfile() ETFProfileData`

GetProfile returns the Profile field if non-nil, zero value otherwise.

### GetProfileOk

`func (o *ETFsProfile) GetProfileOk() (*ETFProfileData, bool)`

GetProfileOk returns a tuple with the Profile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfile

`func (o *ETFsProfile) SetProfile(v ETFProfileData)`

SetProfile sets Profile field to given value.

### HasProfile

`func (o *ETFsProfile) HasProfile() bool`

HasProfile returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


