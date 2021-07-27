# ETFSectorExposureData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Industry** | Pointer to **string** | Industry | [optional] 
**Exposure** | Pointer to **float32** | Percent of exposure. | [optional] 

## Methods

### NewETFSectorExposureData

`func NewETFSectorExposureData() *ETFSectorExposureData`

NewETFSectorExposureData instantiates a new ETFSectorExposureData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewETFSectorExposureDataWithDefaults

`func NewETFSectorExposureDataWithDefaults() *ETFSectorExposureData`

NewETFSectorExposureDataWithDefaults instantiates a new ETFSectorExposureData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIndustry

`func (o *ETFSectorExposureData) GetIndustry() string`

GetIndustry returns the Industry field if non-nil, zero value otherwise.

### GetIndustryOk

`func (o *ETFSectorExposureData) GetIndustryOk() (*string, bool)`

GetIndustryOk returns a tuple with the Industry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndustry

`func (o *ETFSectorExposureData) SetIndustry(v string)`

SetIndustry sets Industry field to given value.

### HasIndustry

`func (o *ETFSectorExposureData) HasIndustry() bool`

HasIndustry returns a boolean if a field has been set.

### GetExposure

`func (o *ETFSectorExposureData) GetExposure() float32`

GetExposure returns the Exposure field if non-nil, zero value otherwise.

### GetExposureOk

`func (o *ETFSectorExposureData) GetExposureOk() (*float32, bool)`

GetExposureOk returns a tuple with the Exposure field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExposure

`func (o *ETFSectorExposureData) SetExposure(v float32)`

SetExposure sets Exposure field to given value.

### HasExposure

`func (o *ETFSectorExposureData) HasExposure() bool`

HasExposure returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


