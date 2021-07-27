# ETFHoldingsData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Symbol** | Pointer to **string** | Symbol description | [optional] 
**Name** | Pointer to **string** | Security name | [optional] 
**Isin** | Pointer to **string** | ISIN. | [optional] 
**Cusip** | Pointer to **string** | CUSIP. | [optional] 
**Share** | Pointer to **float32** | Number of shares owned by the ETF. | [optional] 
**Percent** | Pointer to **float32** | Portfolio&#39;s percent | [optional] 
**Value** | Pointer to **float32** | Market value | [optional] 

## Methods

### NewETFHoldingsData

`func NewETFHoldingsData() *ETFHoldingsData`

NewETFHoldingsData instantiates a new ETFHoldingsData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewETFHoldingsDataWithDefaults

`func NewETFHoldingsDataWithDefaults() *ETFHoldingsData`

NewETFHoldingsDataWithDefaults instantiates a new ETFHoldingsData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSymbol

`func (o *ETFHoldingsData) GetSymbol() string`

GetSymbol returns the Symbol field if non-nil, zero value otherwise.

### GetSymbolOk

`func (o *ETFHoldingsData) GetSymbolOk() (*string, bool)`

GetSymbolOk returns a tuple with the Symbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSymbol

`func (o *ETFHoldingsData) SetSymbol(v string)`

SetSymbol sets Symbol field to given value.

### HasSymbol

`func (o *ETFHoldingsData) HasSymbol() bool`

HasSymbol returns a boolean if a field has been set.

### GetName

`func (o *ETFHoldingsData) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ETFHoldingsData) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ETFHoldingsData) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ETFHoldingsData) HasName() bool`

HasName returns a boolean if a field has been set.

### GetIsin

`func (o *ETFHoldingsData) GetIsin() string`

GetIsin returns the Isin field if non-nil, zero value otherwise.

### GetIsinOk

`func (o *ETFHoldingsData) GetIsinOk() (*string, bool)`

GetIsinOk returns a tuple with the Isin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsin

`func (o *ETFHoldingsData) SetIsin(v string)`

SetIsin sets Isin field to given value.

### HasIsin

`func (o *ETFHoldingsData) HasIsin() bool`

HasIsin returns a boolean if a field has been set.

### GetCusip

`func (o *ETFHoldingsData) GetCusip() string`

GetCusip returns the Cusip field if non-nil, zero value otherwise.

### GetCusipOk

`func (o *ETFHoldingsData) GetCusipOk() (*string, bool)`

GetCusipOk returns a tuple with the Cusip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCusip

`func (o *ETFHoldingsData) SetCusip(v string)`

SetCusip sets Cusip field to given value.

### HasCusip

`func (o *ETFHoldingsData) HasCusip() bool`

HasCusip returns a boolean if a field has been set.

### GetShare

`func (o *ETFHoldingsData) GetShare() float32`

GetShare returns the Share field if non-nil, zero value otherwise.

### GetShareOk

`func (o *ETFHoldingsData) GetShareOk() (*float32, bool)`

GetShareOk returns a tuple with the Share field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShare

`func (o *ETFHoldingsData) SetShare(v float32)`

SetShare sets Share field to given value.

### HasShare

`func (o *ETFHoldingsData) HasShare() bool`

HasShare returns a boolean if a field has been set.

### GetPercent

`func (o *ETFHoldingsData) GetPercent() float32`

GetPercent returns the Percent field if non-nil, zero value otherwise.

### GetPercentOk

`func (o *ETFHoldingsData) GetPercentOk() (*float32, bool)`

GetPercentOk returns a tuple with the Percent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPercent

`func (o *ETFHoldingsData) SetPercent(v float32)`

SetPercent sets Percent field to given value.

### HasPercent

`func (o *ETFHoldingsData) HasPercent() bool`

HasPercent returns a boolean if a field has been set.

### GetValue

`func (o *ETFHoldingsData) GetValue() float32`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *ETFHoldingsData) GetValueOk() (*float32, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *ETFHoldingsData) SetValue(v float32)`

SetValue sets Value field to given value.

### HasValue

`func (o *ETFHoldingsData) HasValue() bool`

HasValue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


