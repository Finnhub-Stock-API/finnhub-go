# PressRelease

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Symbol** | Pointer to **string** | Company symbol. | [optional] 
**MajorDevelopment** | Pointer to [**[]Development**](Development.md) | Array of major developments. | [optional] 

## Methods

### NewPressRelease

`func NewPressRelease() *PressRelease`

NewPressRelease instantiates a new PressRelease object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPressReleaseWithDefaults

`func NewPressReleaseWithDefaults() *PressRelease`

NewPressReleaseWithDefaults instantiates a new PressRelease object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSymbol

`func (o *PressRelease) GetSymbol() string`

GetSymbol returns the Symbol field if non-nil, zero value otherwise.

### GetSymbolOk

`func (o *PressRelease) GetSymbolOk() (*string, bool)`

GetSymbolOk returns a tuple with the Symbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSymbol

`func (o *PressRelease) SetSymbol(v string)`

SetSymbol sets Symbol field to given value.

### HasSymbol

`func (o *PressRelease) HasSymbol() bool`

HasSymbol returns a boolean if a field has been set.

### GetMajorDevelopment

`func (o *PressRelease) GetMajorDevelopment() []Development`

GetMajorDevelopment returns the MajorDevelopment field if non-nil, zero value otherwise.

### GetMajorDevelopmentOk

`func (o *PressRelease) GetMajorDevelopmentOk() (*[]Development, bool)`

GetMajorDevelopmentOk returns a tuple with the MajorDevelopment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMajorDevelopment

`func (o *PressRelease) SetMajorDevelopment(v []Development)`

SetMajorDevelopment sets MajorDevelopment field to given value.

### HasMajorDevelopment

`func (o *PressRelease) HasMajorDevelopment() bool`

HasMajorDevelopment returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


