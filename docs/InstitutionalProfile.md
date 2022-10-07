# InstitutionalProfile

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cik** | Pointer to **string** | CIK. | [optional] 
**Data** | Pointer to [**[]InstitutionalProfileInfo**](InstitutionalProfileInfo.md) | Array of investors. | [optional] 

## Methods

### NewInstitutionalProfile

`func NewInstitutionalProfile() *InstitutionalProfile`

NewInstitutionalProfile instantiates a new InstitutionalProfile object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInstitutionalProfileWithDefaults

`func NewInstitutionalProfileWithDefaults() *InstitutionalProfile`

NewInstitutionalProfileWithDefaults instantiates a new InstitutionalProfile object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCik

`func (o *InstitutionalProfile) GetCik() string`

GetCik returns the Cik field if non-nil, zero value otherwise.

### GetCikOk

`func (o *InstitutionalProfile) GetCikOk() (*string, bool)`

GetCikOk returns a tuple with the Cik field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCik

`func (o *InstitutionalProfile) SetCik(v string)`

SetCik sets Cik field to given value.

### HasCik

`func (o *InstitutionalProfile) HasCik() bool`

HasCik returns a boolean if a field has been set.

### GetData

`func (o *InstitutionalProfile) GetData() []InstitutionalProfileInfo`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *InstitutionalProfile) GetDataOk() (*[]InstitutionalProfileInfo, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *InstitutionalProfile) SetData(v []InstitutionalProfileInfo)`

SetData sets Data field to given value.

### HasData

`func (o *InstitutionalProfile) HasData() bool`

HasData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


