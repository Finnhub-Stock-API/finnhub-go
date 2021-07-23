# Forexrates

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Base** | Pointer to **string** | Base currency. | [optional] 
**Quote** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewForexrates

`func NewForexrates() *Forexrates`

NewForexrates instantiates a new Forexrates object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewForexratesWithDefaults

`func NewForexratesWithDefaults() *Forexrates`

NewForexratesWithDefaults instantiates a new Forexrates object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBase

`func (o *Forexrates) GetBase() string`

GetBase returns the Base field if non-nil, zero value otherwise.

### GetBaseOk

`func (o *Forexrates) GetBaseOk() (*string, bool)`

GetBaseOk returns a tuple with the Base field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBase

`func (o *Forexrates) SetBase(v string)`

SetBase sets Base field to given value.

### HasBase

`func (o *Forexrates) HasBase() bool`

HasBase returns a boolean if a field has been set.

### GetQuote

`func (o *Forexrates) GetQuote() map[string]interface{}`

GetQuote returns the Quote field if non-nil, zero value otherwise.

### GetQuoteOk

`func (o *Forexrates) GetQuoteOk() (*map[string]interface{}, bool)`

GetQuoteOk returns a tuple with the Quote field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuote

`func (o *Forexrates) SetQuote(v map[string]interface{})`

SetQuote sets Quote field to given value.

### HasQuote

`func (o *Forexrates) HasQuote() bool`

HasQuote returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


