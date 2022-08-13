# SectorMetricData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Sector** | Pointer to **string** | Sector | [optional] 
**Metrics** | Pointer to **map[string]interface{}** | Metrics data in key-value format. &lt;code&gt;a&lt;/code&gt; and &lt;code&gt;m&lt;/code&gt; fields are for average and median respectively. | [optional] 

## Methods

### NewSectorMetricData

`func NewSectorMetricData() *SectorMetricData`

NewSectorMetricData instantiates a new SectorMetricData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSectorMetricDataWithDefaults

`func NewSectorMetricDataWithDefaults() *SectorMetricData`

NewSectorMetricDataWithDefaults instantiates a new SectorMetricData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSector

`func (o *SectorMetricData) GetSector() string`

GetSector returns the Sector field if non-nil, zero value otherwise.

### GetSectorOk

`func (o *SectorMetricData) GetSectorOk() (*string, bool)`

GetSectorOk returns a tuple with the Sector field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSector

`func (o *SectorMetricData) SetSector(v string)`

SetSector sets Sector field to given value.

### HasSector

`func (o *SectorMetricData) HasSector() bool`

HasSector returns a boolean if a field has been set.

### GetMetrics

`func (o *SectorMetricData) GetMetrics() map[string]interface{}`

GetMetrics returns the Metrics field if non-nil, zero value otherwise.

### GetMetricsOk

`func (o *SectorMetricData) GetMetricsOk() (*map[string]interface{}, bool)`

GetMetricsOk returns a tuple with the Metrics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetrics

`func (o *SectorMetricData) SetMetrics(v map[string]interface{})`

SetMetrics sets Metrics field to given value.

### HasMetrics

`func (o *SectorMetricData) HasMetrics() bool`

HasMetrics returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


