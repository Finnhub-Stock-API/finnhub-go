# IndicesConstituents

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Symbol** | Pointer to **string** | Index&#39;s symbol. | [optional] 
**Constituents** | Pointer to **[]string** | Array of constituents. | [optional] 
**ConstituentsBreakdown** | Pointer to [**[]IndicesConstituentsBreakdown**](IndicesConstituentsBreakdown.md) | Array of constituents&#39; details. | [optional] 

## Methods

### NewIndicesConstituents

`func NewIndicesConstituents() *IndicesConstituents`

NewIndicesConstituents instantiates a new IndicesConstituents object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIndicesConstituentsWithDefaults

`func NewIndicesConstituentsWithDefaults() *IndicesConstituents`

NewIndicesConstituentsWithDefaults instantiates a new IndicesConstituents object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSymbol

`func (o *IndicesConstituents) GetSymbol() string`

GetSymbol returns the Symbol field if non-nil, zero value otherwise.

### GetSymbolOk

`func (o *IndicesConstituents) GetSymbolOk() (*string, bool)`

GetSymbolOk returns a tuple with the Symbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSymbol

`func (o *IndicesConstituents) SetSymbol(v string)`

SetSymbol sets Symbol field to given value.

### HasSymbol

`func (o *IndicesConstituents) HasSymbol() bool`

HasSymbol returns a boolean if a field has been set.

### GetConstituents

`func (o *IndicesConstituents) GetConstituents() []string`

GetConstituents returns the Constituents field if non-nil, zero value otherwise.

### GetConstituentsOk

`func (o *IndicesConstituents) GetConstituentsOk() (*[]string, bool)`

GetConstituentsOk returns a tuple with the Constituents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConstituents

`func (o *IndicesConstituents) SetConstituents(v []string)`

SetConstituents sets Constituents field to given value.

### HasConstituents

`func (o *IndicesConstituents) HasConstituents() bool`

HasConstituents returns a boolean if a field has been set.

### GetConstituentsBreakdown

`func (o *IndicesConstituents) GetConstituentsBreakdown() []IndicesConstituentsBreakdown`

GetConstituentsBreakdown returns the ConstituentsBreakdown field if non-nil, zero value otherwise.

### GetConstituentsBreakdownOk

`func (o *IndicesConstituents) GetConstituentsBreakdownOk() (*[]IndicesConstituentsBreakdown, bool)`

GetConstituentsBreakdownOk returns a tuple with the ConstituentsBreakdown field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConstituentsBreakdown

`func (o *IndicesConstituents) SetConstituentsBreakdown(v []IndicesConstituentsBreakdown)`

SetConstituentsBreakdown sets ConstituentsBreakdown field to given value.

### HasConstituentsBreakdown

`func (o *IndicesConstituents) HasConstituentsBreakdown() bool`

HasConstituentsBreakdown returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


