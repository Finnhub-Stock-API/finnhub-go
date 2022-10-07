# BondTickData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Skip** | Pointer to **int64** | Number of ticks skipped. | [optional] 
**Count** | Pointer to **int64** | Number of ticks returned. If &lt;code&gt;count&lt;/code&gt; &lt; &lt;code&gt;limit&lt;/code&gt;, all data for that date has been returned. | [optional] 
**Total** | Pointer to **int64** | Total number of ticks for that date. | [optional] 
**V** | Pointer to **[]float32** | List of volume data. | [optional] 
**P** | Pointer to **[]float32** | List of price data. | [optional] 
**T** | Pointer to **[]int64** | List of timestamp in UNIX ms. | [optional] 
**Si** | Pointer to **[]string** | List of values showing the side (Buy/sell) of each trade. List of supported values: &lt;a target&#x3D;\&quot;_blank\&quot; href&#x3D;\&quot;https://docs.google.com/spreadsheets/d/1O3aueXSPOqo7Iuyz4PqDG6yZunHsX8BTefZ2kFk5pz4/edit?usp&#x3D;sharing\&quot;,&gt;here&lt;/a&gt; | [optional] 
**Cp** | Pointer to **[]string** | List of values showing the counterparty of each trade. List of supported values: &lt;a target&#x3D;\&quot;_blank\&quot; href&#x3D;\&quot;https://docs.google.com/spreadsheets/d/1O3aueXSPOqo7Iuyz4PqDG6yZunHsX8BTefZ2kFk5pz4/edit?usp&#x3D;sharing\&quot;,&gt;here&lt;/a&gt; | [optional] 
**C** | Pointer to **[][]string** | List of trade conditions. A comprehensive list of trade conditions code can be found &lt;a target&#x3D;\&quot;_blank\&quot; href&#x3D;\&quot;https://docs.google.com/spreadsheets/d/1O3aueXSPOqo7Iuyz4PqDG6yZunHsX8BTefZ2kFk5pz4/edit?usp&#x3D;sharing\&quot;&gt;here&lt;/a&gt; | [optional] 

## Methods

### NewBondTickData

`func NewBondTickData() *BondTickData`

NewBondTickData instantiates a new BondTickData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBondTickDataWithDefaults

`func NewBondTickDataWithDefaults() *BondTickData`

NewBondTickDataWithDefaults instantiates a new BondTickData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSkip

`func (o *BondTickData) GetSkip() int64`

GetSkip returns the Skip field if non-nil, zero value otherwise.

### GetSkipOk

`func (o *BondTickData) GetSkipOk() (*int64, bool)`

GetSkipOk returns a tuple with the Skip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkip

`func (o *BondTickData) SetSkip(v int64)`

SetSkip sets Skip field to given value.

### HasSkip

`func (o *BondTickData) HasSkip() bool`

HasSkip returns a boolean if a field has been set.

### GetCount

`func (o *BondTickData) GetCount() int64`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *BondTickData) GetCountOk() (*int64, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *BondTickData) SetCount(v int64)`

SetCount sets Count field to given value.

### HasCount

`func (o *BondTickData) HasCount() bool`

HasCount returns a boolean if a field has been set.

### GetTotal

`func (o *BondTickData) GetTotal() int64`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *BondTickData) GetTotalOk() (*int64, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *BondTickData) SetTotal(v int64)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *BondTickData) HasTotal() bool`

HasTotal returns a boolean if a field has been set.

### GetV

`func (o *BondTickData) GetV() []float32`

GetV returns the V field if non-nil, zero value otherwise.

### GetVOk

`func (o *BondTickData) GetVOk() (*[]float32, bool)`

GetVOk returns a tuple with the V field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetV

`func (o *BondTickData) SetV(v []float32)`

SetV sets V field to given value.

### HasV

`func (o *BondTickData) HasV() bool`

HasV returns a boolean if a field has been set.

### GetP

`func (o *BondTickData) GetP() []float32`

GetP returns the P field if non-nil, zero value otherwise.

### GetPOk

`func (o *BondTickData) GetPOk() (*[]float32, bool)`

GetPOk returns a tuple with the P field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetP

`func (o *BondTickData) SetP(v []float32)`

SetP sets P field to given value.

### HasP

`func (o *BondTickData) HasP() bool`

HasP returns a boolean if a field has been set.

### GetT

`func (o *BondTickData) GetT() []int64`

GetT returns the T field if non-nil, zero value otherwise.

### GetTOk

`func (o *BondTickData) GetTOk() (*[]int64, bool)`

GetTOk returns a tuple with the T field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetT

`func (o *BondTickData) SetT(v []int64)`

SetT sets T field to given value.

### HasT

`func (o *BondTickData) HasT() bool`

HasT returns a boolean if a field has been set.

### GetSi

`func (o *BondTickData) GetSi() []string`

GetSi returns the Si field if non-nil, zero value otherwise.

### GetSiOk

`func (o *BondTickData) GetSiOk() (*[]string, bool)`

GetSiOk returns a tuple with the Si field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSi

`func (o *BondTickData) SetSi(v []string)`

SetSi sets Si field to given value.

### HasSi

`func (o *BondTickData) HasSi() bool`

HasSi returns a boolean if a field has been set.

### GetCp

`func (o *BondTickData) GetCp() []string`

GetCp returns the Cp field if non-nil, zero value otherwise.

### GetCpOk

`func (o *BondTickData) GetCpOk() (*[]string, bool)`

GetCpOk returns a tuple with the Cp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCp

`func (o *BondTickData) SetCp(v []string)`

SetCp sets Cp field to given value.

### HasCp

`func (o *BondTickData) HasCp() bool`

HasCp returns a boolean if a field has been set.

### GetC

`func (o *BondTickData) GetC() [][]string`

GetC returns the C field if non-nil, zero value otherwise.

### GetCOk

`func (o *BondTickData) GetCOk() (*[][]string, bool)`

GetCOk returns a tuple with the C field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetC

`func (o *BondTickData) SetC(v [][]string)`

SetC sets C field to given value.

### HasC

`func (o *BondTickData) HasC() bool`

HasC returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


