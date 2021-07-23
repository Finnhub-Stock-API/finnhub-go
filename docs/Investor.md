# Investor

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | Investor&#39;s name. | [optional] 
**Share** | Pointer to **int64** | Number of shares held by the investor. | [optional] 
**Change** | Pointer to **int64** | Number of share changed (net buy or sell) from the last period. | [optional] 
**FilingDate** | Pointer to **string** | Filing date. | [optional] 

## Methods

### NewInvestor

`func NewInvestor() *Investor`

NewInvestor instantiates a new Investor object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInvestorWithDefaults

`func NewInvestorWithDefaults() *Investor`

NewInvestorWithDefaults instantiates a new Investor object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *Investor) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Investor) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Investor) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Investor) HasName() bool`

HasName returns a boolean if a field has been set.

### GetShare

`func (o *Investor) GetShare() int64`

GetShare returns the Share field if non-nil, zero value otherwise.

### GetShareOk

`func (o *Investor) GetShareOk() (*int64, bool)`

GetShareOk returns a tuple with the Share field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShare

`func (o *Investor) SetShare(v int64)`

SetShare sets Share field to given value.

### HasShare

`func (o *Investor) HasShare() bool`

HasShare returns a boolean if a field has been set.

### GetChange

`func (o *Investor) GetChange() int64`

GetChange returns the Change field if non-nil, zero value otherwise.

### GetChangeOk

`func (o *Investor) GetChangeOk() (*int64, bool)`

GetChangeOk returns a tuple with the Change field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChange

`func (o *Investor) SetChange(v int64)`

SetChange sets Change field to given value.

### HasChange

`func (o *Investor) HasChange() bool`

HasChange returns a boolean if a field has been set.

### GetFilingDate

`func (o *Investor) GetFilingDate() string`

GetFilingDate returns the FilingDate field if non-nil, zero value otherwise.

### GetFilingDateOk

`func (o *Investor) GetFilingDateOk() (*string, bool)`

GetFilingDateOk returns a tuple with the FilingDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilingDate

`func (o *Investor) SetFilingDate(v string)`

SetFilingDate sets FilingDate field to given value.

### HasFilingDate

`func (o *Investor) HasFilingDate() bool`

HasFilingDate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


