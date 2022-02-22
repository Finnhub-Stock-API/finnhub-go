# VisaApplicationResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Symbol** | Pointer to **string** | Symbol. | [optional] 
**Data** | Pointer to [**[]VisaApplication**](VisaApplication.md) | Array of H1b and Permanent visa applications. | [optional] 

## Methods

### NewVisaApplicationResult

`func NewVisaApplicationResult() *VisaApplicationResult`

NewVisaApplicationResult instantiates a new VisaApplicationResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVisaApplicationResultWithDefaults

`func NewVisaApplicationResultWithDefaults() *VisaApplicationResult`

NewVisaApplicationResultWithDefaults instantiates a new VisaApplicationResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSymbol

`func (o *VisaApplicationResult) GetSymbol() string`

GetSymbol returns the Symbol field if non-nil, zero value otherwise.

### GetSymbolOk

`func (o *VisaApplicationResult) GetSymbolOk() (*string, bool)`

GetSymbolOk returns a tuple with the Symbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSymbol

`func (o *VisaApplicationResult) SetSymbol(v string)`

SetSymbol sets Symbol field to given value.

### HasSymbol

`func (o *VisaApplicationResult) HasSymbol() bool`

HasSymbol returns a boolean if a field has been set.

### GetData

`func (o *VisaApplicationResult) GetData() []VisaApplication`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *VisaApplicationResult) GetDataOk() (*[]VisaApplication, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *VisaApplicationResult) SetData(v []VisaApplication)`

SetData sets Data field to given value.

### HasData

`func (o *VisaApplicationResult) HasData() bool`

HasData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


