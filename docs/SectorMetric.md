# SectorMetric

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Region** | Pointer to **string** | Region. | [optional] 
**Data** | Pointer to [**[]SectorMetricData**](SectorMetricData.md) | Metrics for each sector. | [optional] 

## Methods

### NewSectorMetric

`func NewSectorMetric() *SectorMetric`

NewSectorMetric instantiates a new SectorMetric object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSectorMetricWithDefaults

`func NewSectorMetricWithDefaults() *SectorMetric`

NewSectorMetricWithDefaults instantiates a new SectorMetric object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRegion

`func (o *SectorMetric) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *SectorMetric) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *SectorMetric) SetRegion(v string)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *SectorMetric) HasRegion() bool`

HasRegion returns a boolean if a field has been set.

### GetData

`func (o *SectorMetric) GetData() []SectorMetricData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *SectorMetric) GetDataOk() (*[]SectorMetricData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *SectorMetric) SetData(v []SectorMetricData)`

SetData sets Data field to given value.

### HasData

`func (o *SectorMetric) HasData() bool`

HasData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


