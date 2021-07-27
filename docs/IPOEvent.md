# IPOEvent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Symbol** | Pointer to **string** | Symbol. | [optional] 
**Date** | Pointer to **string** | IPO date. | [optional] 
**Exchange** | Pointer to **string** | Exchange. | [optional] 
**Name** | Pointer to **string** | Company&#39;s name. | [optional] 
**Status** | Pointer to **string** | IPO status. Can take 1 of the following values: &lt;code&gt;expected&lt;/code&gt;,&lt;code&gt;priced&lt;/code&gt;,&lt;code&gt;withdrawn&lt;/code&gt;,&lt;code&gt;filed&lt;/code&gt; | [optional] 
**Price** | Pointer to **string** | Projected price or price range. | [optional] 
**NumberOfShares** | Pointer to **float32** | Number of shares offered during the IPO. | [optional] 
**TotalSharesValue** | Pointer to **float32** | Total shares value. | [optional] 

## Methods

### NewIPOEvent

`func NewIPOEvent() *IPOEvent`

NewIPOEvent instantiates a new IPOEvent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIPOEventWithDefaults

`func NewIPOEventWithDefaults() *IPOEvent`

NewIPOEventWithDefaults instantiates a new IPOEvent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSymbol

`func (o *IPOEvent) GetSymbol() string`

GetSymbol returns the Symbol field if non-nil, zero value otherwise.

### GetSymbolOk

`func (o *IPOEvent) GetSymbolOk() (*string, bool)`

GetSymbolOk returns a tuple with the Symbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSymbol

`func (o *IPOEvent) SetSymbol(v string)`

SetSymbol sets Symbol field to given value.

### HasSymbol

`func (o *IPOEvent) HasSymbol() bool`

HasSymbol returns a boolean if a field has been set.

### GetDate

`func (o *IPOEvent) GetDate() string`

GetDate returns the Date field if non-nil, zero value otherwise.

### GetDateOk

`func (o *IPOEvent) GetDateOk() (*string, bool)`

GetDateOk returns a tuple with the Date field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDate

`func (o *IPOEvent) SetDate(v string)`

SetDate sets Date field to given value.

### HasDate

`func (o *IPOEvent) HasDate() bool`

HasDate returns a boolean if a field has been set.

### GetExchange

`func (o *IPOEvent) GetExchange() string`

GetExchange returns the Exchange field if non-nil, zero value otherwise.

### GetExchangeOk

`func (o *IPOEvent) GetExchangeOk() (*string, bool)`

GetExchangeOk returns a tuple with the Exchange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExchange

`func (o *IPOEvent) SetExchange(v string)`

SetExchange sets Exchange field to given value.

### HasExchange

`func (o *IPOEvent) HasExchange() bool`

HasExchange returns a boolean if a field has been set.

### GetName

`func (o *IPOEvent) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *IPOEvent) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *IPOEvent) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *IPOEvent) HasName() bool`

HasName returns a boolean if a field has been set.

### GetStatus

`func (o *IPOEvent) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *IPOEvent) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *IPOEvent) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *IPOEvent) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetPrice

`func (o *IPOEvent) GetPrice() string`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *IPOEvent) GetPriceOk() (*string, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *IPOEvent) SetPrice(v string)`

SetPrice sets Price field to given value.

### HasPrice

`func (o *IPOEvent) HasPrice() bool`

HasPrice returns a boolean if a field has been set.

### GetNumberOfShares

`func (o *IPOEvent) GetNumberOfShares() float32`

GetNumberOfShares returns the NumberOfShares field if non-nil, zero value otherwise.

### GetNumberOfSharesOk

`func (o *IPOEvent) GetNumberOfSharesOk() (*float32, bool)`

GetNumberOfSharesOk returns a tuple with the NumberOfShares field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumberOfShares

`func (o *IPOEvent) SetNumberOfShares(v float32)`

SetNumberOfShares sets NumberOfShares field to given value.

### HasNumberOfShares

`func (o *IPOEvent) HasNumberOfShares() bool`

HasNumberOfShares returns a boolean if a field has been set.

### GetTotalSharesValue

`func (o *IPOEvent) GetTotalSharesValue() float32`

GetTotalSharesValue returns the TotalSharesValue field if non-nil, zero value otherwise.

### GetTotalSharesValueOk

`func (o *IPOEvent) GetTotalSharesValueOk() (*float32, bool)`

GetTotalSharesValueOk returns a tuple with the TotalSharesValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalSharesValue

`func (o *IPOEvent) SetTotalSharesValue(v float32)`

SetTotalSharesValue sets TotalSharesValue field to given value.

### HasTotalSharesValue

`func (o *IPOEvent) HasTotalSharesValue() bool`

HasTotalSharesValue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


