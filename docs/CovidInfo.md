# CovidInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**State** | Pointer to **string** | State. | [optional] 
**Case** | Pointer to **float32** | Number of confirmed cases. | [optional] 
**Death** | Pointer to **float32** | Number of confirmed deaths. | [optional] 
**Updated** | Pointer to **string** | Updated time. | [optional] 

## Methods

### NewCovidInfo

`func NewCovidInfo() *CovidInfo`

NewCovidInfo instantiates a new CovidInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCovidInfoWithDefaults

`func NewCovidInfoWithDefaults() *CovidInfo`

NewCovidInfoWithDefaults instantiates a new CovidInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetState

`func (o *CovidInfo) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *CovidInfo) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *CovidInfo) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *CovidInfo) HasState() bool`

HasState returns a boolean if a field has been set.

### GetCase

`func (o *CovidInfo) GetCase() float32`

GetCase returns the Case field if non-nil, zero value otherwise.

### GetCaseOk

`func (o *CovidInfo) GetCaseOk() (*float32, bool)`

GetCaseOk returns a tuple with the Case field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCase

`func (o *CovidInfo) SetCase(v float32)`

SetCase sets Case field to given value.

### HasCase

`func (o *CovidInfo) HasCase() bool`

HasCase returns a boolean if a field has been set.

### GetDeath

`func (o *CovidInfo) GetDeath() float32`

GetDeath returns the Death field if non-nil, zero value otherwise.

### GetDeathOk

`func (o *CovidInfo) GetDeathOk() (*float32, bool)`

GetDeathOk returns a tuple with the Death field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeath

`func (o *CovidInfo) SetDeath(v float32)`

SetDeath sets Death field to given value.

### HasDeath

`func (o *CovidInfo) HasDeath() bool`

HasDeath returns a boolean if a field has been set.

### GetUpdated

`func (o *CovidInfo) GetUpdated() string`

GetUpdated returns the Updated field if non-nil, zero value otherwise.

### GetUpdatedOk

`func (o *CovidInfo) GetUpdatedOk() (*string, bool)`

GetUpdatedOk returns a tuple with the Updated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdated

`func (o *CovidInfo) SetUpdated(v string)`

SetUpdated sets Updated field to given value.

### HasUpdated

`func (o *CovidInfo) HasUpdated() bool`

HasUpdated returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


