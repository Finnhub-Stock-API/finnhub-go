# BasicFinancials

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Symbol** | Pointer to **string** | Symbol of the company. | [optional] 
**MetricType** | Pointer to **string** | Metric type. | [optional] 
**Series** | Pointer to **map[string]interface{}** |  | [optional] 
**Metric** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewBasicFinancials

`func NewBasicFinancials() *BasicFinancials`

NewBasicFinancials instantiates a new BasicFinancials object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBasicFinancialsWithDefaults

`func NewBasicFinancialsWithDefaults() *BasicFinancials`

NewBasicFinancialsWithDefaults instantiates a new BasicFinancials object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSymbol

`func (o *BasicFinancials) GetSymbol() string`

GetSymbol returns the Symbol field if non-nil, zero value otherwise.

### GetSymbolOk

`func (o *BasicFinancials) GetSymbolOk() (*string, bool)`

GetSymbolOk returns a tuple with the Symbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSymbol

`func (o *BasicFinancials) SetSymbol(v string)`

SetSymbol sets Symbol field to given value.

### HasSymbol

`func (o *BasicFinancials) HasSymbol() bool`

HasSymbol returns a boolean if a field has been set.

### GetMetricType

`func (o *BasicFinancials) GetMetricType() string`

GetMetricType returns the MetricType field if non-nil, zero value otherwise.

### GetMetricTypeOk

`func (o *BasicFinancials) GetMetricTypeOk() (*string, bool)`

GetMetricTypeOk returns a tuple with the MetricType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetricType

`func (o *BasicFinancials) SetMetricType(v string)`

SetMetricType sets MetricType field to given value.

### HasMetricType

`func (o *BasicFinancials) HasMetricType() bool`

HasMetricType returns a boolean if a field has been set.

### GetSeries

`func (o *BasicFinancials) GetSeries() map[string]interface{}`

GetSeries returns the Series field if non-nil, zero value otherwise.

### GetSeriesOk

`func (o *BasicFinancials) GetSeriesOk() (*map[string]interface{}, bool)`

GetSeriesOk returns a tuple with the Series field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeries

`func (o *BasicFinancials) SetSeries(v map[string]interface{})`

SetSeries sets Series field to given value.

### HasSeries

`func (o *BasicFinancials) HasSeries() bool`

HasSeries returns a boolean if a field has been set.

### GetMetric

`func (o *BasicFinancials) GetMetric() map[string]interface{}`

GetMetric returns the Metric field if non-nil, zero value otherwise.

### GetMetricOk

`func (o *BasicFinancials) GetMetricOk() (*map[string]interface{}, bool)`

GetMetricOk returns a tuple with the Metric field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetric

`func (o *BasicFinancials) SetMetric(v map[string]interface{})`

SetMetric sets Metric field to given value.

### HasMetric

`func (o *BasicFinancials) HasMetric() bool`

HasMetric returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


