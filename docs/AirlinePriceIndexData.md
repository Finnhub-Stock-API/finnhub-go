# AirlinePriceIndexData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**[]AirlinePriceIndex**](AirlinePriceIndex.md) | Array of price index. | [optional] 
**Airline** | Pointer to **string** | Airline name | [optional] 
**From** | Pointer to **string** | From date | [optional] 
**To** | Pointer to **string** | To date | [optional] 

## Methods

### NewAirlinePriceIndexData

`func NewAirlinePriceIndexData() *AirlinePriceIndexData`

NewAirlinePriceIndexData instantiates a new AirlinePriceIndexData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAirlinePriceIndexDataWithDefaults

`func NewAirlinePriceIndexDataWithDefaults() *AirlinePriceIndexData`

NewAirlinePriceIndexDataWithDefaults instantiates a new AirlinePriceIndexData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *AirlinePriceIndexData) GetData() []AirlinePriceIndex`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *AirlinePriceIndexData) GetDataOk() (*[]AirlinePriceIndex, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *AirlinePriceIndexData) SetData(v []AirlinePriceIndex)`

SetData sets Data field to given value.

### HasData

`func (o *AirlinePriceIndexData) HasData() bool`

HasData returns a boolean if a field has been set.

### GetAirline

`func (o *AirlinePriceIndexData) GetAirline() string`

GetAirline returns the Airline field if non-nil, zero value otherwise.

### GetAirlineOk

`func (o *AirlinePriceIndexData) GetAirlineOk() (*string, bool)`

GetAirlineOk returns a tuple with the Airline field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAirline

`func (o *AirlinePriceIndexData) SetAirline(v string)`

SetAirline sets Airline field to given value.

### HasAirline

`func (o *AirlinePriceIndexData) HasAirline() bool`

HasAirline returns a boolean if a field has been set.

### GetFrom

`func (o *AirlinePriceIndexData) GetFrom() string`

GetFrom returns the From field if non-nil, zero value otherwise.

### GetFromOk

`func (o *AirlinePriceIndexData) GetFromOk() (*string, bool)`

GetFromOk returns a tuple with the From field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrom

`func (o *AirlinePriceIndexData) SetFrom(v string)`

SetFrom sets From field to given value.

### HasFrom

`func (o *AirlinePriceIndexData) HasFrom() bool`

HasFrom returns a boolean if a field has been set.

### GetTo

`func (o *AirlinePriceIndexData) GetTo() string`

GetTo returns the To field if non-nil, zero value otherwise.

### GetToOk

`func (o *AirlinePriceIndexData) GetToOk() (*string, bool)`

GetToOk returns a tuple with the To field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTo

`func (o *AirlinePriceIndexData) SetTo(v string)`

SetTo sets To field to given value.

### HasTo

`func (o *AirlinePriceIndexData) HasTo() bool`

HasTo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


