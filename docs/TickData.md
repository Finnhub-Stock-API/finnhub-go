# TickData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**S** | Pointer to **string** | Symbol. | [optional] 
**Skip** | Pointer to **int64** | Number of ticks skipped. | [optional] 
**Count** | Pointer to **int64** | Number of ticks returned. If &lt;code&gt;count&lt;/code&gt; &lt; &lt;code&gt;limit&lt;/code&gt;, all data for that date has been returned. | [optional] 
**Total** | Pointer to **int64** | Total number of ticks for that date. | [optional] 
**V** | Pointer to **[]float32** | List of volume data. | [optional] 
**P** | Pointer to **[]float32** | List of price data. | [optional] 
**T** | Pointer to **[]int64** | List of timestamp in UNIX ms. | [optional] 
**X** | Pointer to **[]string** | List of venues/exchanges. A list of exchange codes can be found &lt;a target&#x3D;\&quot;_blank\&quot; href&#x3D;\&quot;https://docs.google.com/spreadsheets/d/1Tj53M1svmr-hfEtbk6_NpVR1yAyGLMaH6ByYU6CG0ZY/edit?usp&#x3D;sharing\&quot;,&gt;here&lt;/a&gt; | [optional] 
**C** | Pointer to **[][]string** | List of trade conditions. A comprehensive list of trade conditions code can be found &lt;a target&#x3D;\&quot;_blank\&quot; href&#x3D;\&quot;https://docs.google.com/spreadsheets/d/1PUxiSWPHSODbaTaoL2Vef6DgU-yFtlRGZf19oBb9Hp0/edit?usp&#x3D;sharing\&quot;&gt;here&lt;/a&gt; | [optional] 

## Methods

### NewTickData

`func NewTickData() *TickData`

NewTickData instantiates a new TickData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTickDataWithDefaults

`func NewTickDataWithDefaults() *TickData`

NewTickDataWithDefaults instantiates a new TickData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetS

`func (o *TickData) GetS() string`

GetS returns the S field if non-nil, zero value otherwise.

### GetSOk

`func (o *TickData) GetSOk() (*string, bool)`

GetSOk returns a tuple with the S field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetS

`func (o *TickData) SetS(v string)`

SetS sets S field to given value.

### HasS

`func (o *TickData) HasS() bool`

HasS returns a boolean if a field has been set.

### GetSkip

`func (o *TickData) GetSkip() int64`

GetSkip returns the Skip field if non-nil, zero value otherwise.

### GetSkipOk

`func (o *TickData) GetSkipOk() (*int64, bool)`

GetSkipOk returns a tuple with the Skip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkip

`func (o *TickData) SetSkip(v int64)`

SetSkip sets Skip field to given value.

### HasSkip

`func (o *TickData) HasSkip() bool`

HasSkip returns a boolean if a field has been set.

### GetCount

`func (o *TickData) GetCount() int64`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *TickData) GetCountOk() (*int64, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *TickData) SetCount(v int64)`

SetCount sets Count field to given value.

### HasCount

`func (o *TickData) HasCount() bool`

HasCount returns a boolean if a field has been set.

### GetTotal

`func (o *TickData) GetTotal() int64`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *TickData) GetTotalOk() (*int64, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *TickData) SetTotal(v int64)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *TickData) HasTotal() bool`

HasTotal returns a boolean if a field has been set.

### GetV

`func (o *TickData) GetV() []float32`

GetV returns the V field if non-nil, zero value otherwise.

### GetVOk

`func (o *TickData) GetVOk() (*[]float32, bool)`

GetVOk returns a tuple with the V field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetV

`func (o *TickData) SetV(v []float32)`

SetV sets V field to given value.

### HasV

`func (o *TickData) HasV() bool`

HasV returns a boolean if a field has been set.

### GetP

`func (o *TickData) GetP() []float32`

GetP returns the P field if non-nil, zero value otherwise.

### GetPOk

`func (o *TickData) GetPOk() (*[]float32, bool)`

GetPOk returns a tuple with the P field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetP

`func (o *TickData) SetP(v []float32)`

SetP sets P field to given value.

### HasP

`func (o *TickData) HasP() bool`

HasP returns a boolean if a field has been set.

### GetT

`func (o *TickData) GetT() []int64`

GetT returns the T field if non-nil, zero value otherwise.

### GetTOk

`func (o *TickData) GetTOk() (*[]int64, bool)`

GetTOk returns a tuple with the T field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetT

`func (o *TickData) SetT(v []int64)`

SetT sets T field to given value.

### HasT

`func (o *TickData) HasT() bool`

HasT returns a boolean if a field has been set.

### GetX

`func (o *TickData) GetX() []string`

GetX returns the X field if non-nil, zero value otherwise.

### GetXOk

`func (o *TickData) GetXOk() (*[]string, bool)`

GetXOk returns a tuple with the X field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetX

`func (o *TickData) SetX(v []string)`

SetX sets X field to given value.

### HasX

`func (o *TickData) HasX() bool`

HasX returns a boolean if a field has been set.

### GetC

`func (o *TickData) GetC() [][]interface{}`

GetC returns the C field if non-nil, zero value otherwise.

### GetCOk

`func (o *TickData) GetCOk() (*[][]interface{}, bool)`

GetCOk returns a tuple with the C field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetC

`func (o *TickData) SetC(v [][]interface{})`

SetC sets C field to given value.

### HasC

`func (o *TickData) HasC() bool`

HasC returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


