# Development

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Symbol** | Pointer to **string** | Company symbol. | [optional] 
**Datetime** | Pointer to **string** | Published time in &lt;code&gt;YYYY-MM-DD HH:MM:SS&lt;/code&gt; format. | [optional] 
**Headline** | Pointer to **string** | Development headline. | [optional] 
**Description** | Pointer to **string** | Development description. | [optional] 
**Url** | Pointer to **string** | URL. | [optional] 

## Methods

### NewDevelopment

`func NewDevelopment() *Development`

NewDevelopment instantiates a new Development object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDevelopmentWithDefaults

`func NewDevelopmentWithDefaults() *Development`

NewDevelopmentWithDefaults instantiates a new Development object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSymbol

`func (o *Development) GetSymbol() string`

GetSymbol returns the Symbol field if non-nil, zero value otherwise.

### GetSymbolOk

`func (o *Development) GetSymbolOk() (*string, bool)`

GetSymbolOk returns a tuple with the Symbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSymbol

`func (o *Development) SetSymbol(v string)`

SetSymbol sets Symbol field to given value.

### HasSymbol

`func (o *Development) HasSymbol() bool`

HasSymbol returns a boolean if a field has been set.

### GetDatetime

`func (o *Development) GetDatetime() string`

GetDatetime returns the Datetime field if non-nil, zero value otherwise.

### GetDatetimeOk

`func (o *Development) GetDatetimeOk() (*string, bool)`

GetDatetimeOk returns a tuple with the Datetime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatetime

`func (o *Development) SetDatetime(v string)`

SetDatetime sets Datetime field to given value.

### HasDatetime

`func (o *Development) HasDatetime() bool`

HasDatetime returns a boolean if a field has been set.

### GetHeadline

`func (o *Development) GetHeadline() string`

GetHeadline returns the Headline field if non-nil, zero value otherwise.

### GetHeadlineOk

`func (o *Development) GetHeadlineOk() (*string, bool)`

GetHeadlineOk returns a tuple with the Headline field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeadline

`func (o *Development) SetHeadline(v string)`

SetHeadline sets Headline field to given value.

### HasHeadline

`func (o *Development) HasHeadline() bool`

HasHeadline returns a boolean if a field has been set.

### GetDescription

`func (o *Development) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Development) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Development) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Development) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetUrl

`func (o *Development) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *Development) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *Development) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *Development) HasUrl() bool`

HasUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


