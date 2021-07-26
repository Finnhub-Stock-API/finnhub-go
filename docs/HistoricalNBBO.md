# HistoricalNBBO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**S** | Pointer to **string** | Symbol. | [optional] 
**Skip** | Pointer to **int64** | Number of ticks skipped. | [optional] 
**Count** | Pointer to **int64** | Number of ticks returned. If &lt;code&gt;count&lt;/code&gt; &lt; &lt;code&gt;limit&lt;/code&gt;, all data for that date has been returned. | [optional] 
**Total** | Pointer to **int64** | Total number of ticks for that date. | [optional] 
**Av** | Pointer to **[]float32** | List of Ask volume data. | [optional] 
**A** | Pointer to **[]float32** | List of Ask price data. | [optional] 
**Ax** | Pointer to **[]string** | List of venues/exchanges - Ask price. A list of exchange codes can be found &lt;a target&#x3D;\&quot;_blank\&quot; href&#x3D;\&quot;https://docs.google.com/spreadsheets/d/1Tj53M1svmr-hfEtbk6_NpVR1yAyGLMaH6ByYU6CG0ZY/edit?usp&#x3D;sharing\&quot;,&gt;here&lt;/a&gt; | [optional] 
**Bv** | Pointer to **[]float32** | List of Bid volume data. | [optional] 
**B** | Pointer to **[]float32** | List of Bid price data. | [optional] 
**Bx** | Pointer to **[]string** | List of venues/exchanges - Bid price. A list of exchange codes can be found &lt;a target&#x3D;\&quot;_blank\&quot; href&#x3D;\&quot;https://docs.google.com/spreadsheets/d/1Tj53M1svmr-hfEtbk6_NpVR1yAyGLMaH6ByYU6CG0ZY/edit?usp&#x3D;sharing\&quot;,&gt;here&lt;/a&gt; | [optional] 
**T** | Pointer to **[]int64** | List of timestamp in UNIX ms. | [optional] 
**C** | Pointer to **[][]string** | List of quote conditions. A comprehensive list of quote conditions code can be found &lt;a target&#x3D;\&quot;_blank\&quot; href&#x3D;\&quot;https://docs.google.com/spreadsheets/d/1iiA6e7Osdtai0oPMOUzgAIKXCsay89dFDmsegz6OpEg/edit?usp&#x3D;sharing\&quot;&gt;here&lt;/a&gt; | [optional] 

## Methods

### NewHistoricalNBBO

`func NewHistoricalNBBO() *HistoricalNBBO`

NewHistoricalNBBO instantiates a new HistoricalNBBO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHistoricalNBBOWithDefaults

`func NewHistoricalNBBOWithDefaults() *HistoricalNBBO`

NewHistoricalNBBOWithDefaults instantiates a new HistoricalNBBO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetS

`func (o *HistoricalNBBO) GetS() string`

GetS returns the S field if non-nil, zero value otherwise.

### GetSOk

`func (o *HistoricalNBBO) GetSOk() (*string, bool)`

GetSOk returns a tuple with the S field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetS

`func (o *HistoricalNBBO) SetS(v string)`

SetS sets S field to given value.

### HasS

`func (o *HistoricalNBBO) HasS() bool`

HasS returns a boolean if a field has been set.

### GetSkip

`func (o *HistoricalNBBO) GetSkip() int64`

GetSkip returns the Skip field if non-nil, zero value otherwise.

### GetSkipOk

`func (o *HistoricalNBBO) GetSkipOk() (*int64, bool)`

GetSkipOk returns a tuple with the Skip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkip

`func (o *HistoricalNBBO) SetSkip(v int64)`

SetSkip sets Skip field to given value.

### HasSkip

`func (o *HistoricalNBBO) HasSkip() bool`

HasSkip returns a boolean if a field has been set.

### GetCount

`func (o *HistoricalNBBO) GetCount() int64`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *HistoricalNBBO) GetCountOk() (*int64, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *HistoricalNBBO) SetCount(v int64)`

SetCount sets Count field to given value.

### HasCount

`func (o *HistoricalNBBO) HasCount() bool`

HasCount returns a boolean if a field has been set.

### GetTotal

`func (o *HistoricalNBBO) GetTotal() int64`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *HistoricalNBBO) GetTotalOk() (*int64, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *HistoricalNBBO) SetTotal(v int64)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *HistoricalNBBO) HasTotal() bool`

HasTotal returns a boolean if a field has been set.

### GetAv

`func (o *HistoricalNBBO) GetAv() []float32`

GetAv returns the Av field if non-nil, zero value otherwise.

### GetAvOk

`func (o *HistoricalNBBO) GetAvOk() (*[]float32, bool)`

GetAvOk returns a tuple with the Av field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAv

`func (o *HistoricalNBBO) SetAv(v []float32)`

SetAv sets Av field to given value.

### HasAv

`func (o *HistoricalNBBO) HasAv() bool`

HasAv returns a boolean if a field has been set.

### GetA

`func (o *HistoricalNBBO) GetA() []float32`

GetA returns the A field if non-nil, zero value otherwise.

### GetAOk

`func (o *HistoricalNBBO) GetAOk() (*[]float32, bool)`

GetAOk returns a tuple with the A field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetA

`func (o *HistoricalNBBO) SetA(v []float32)`

SetA sets A field to given value.

### HasA

`func (o *HistoricalNBBO) HasA() bool`

HasA returns a boolean if a field has been set.

### GetAx

`func (o *HistoricalNBBO) GetAx() []string`

GetAx returns the Ax field if non-nil, zero value otherwise.

### GetAxOk

`func (o *HistoricalNBBO) GetAxOk() (*[]string, bool)`

GetAxOk returns a tuple with the Ax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAx

`func (o *HistoricalNBBO) SetAx(v []string)`

SetAx sets Ax field to given value.

### HasAx

`func (o *HistoricalNBBO) HasAx() bool`

HasAx returns a boolean if a field has been set.

### GetBv

`func (o *HistoricalNBBO) GetBv() []float32`

GetBv returns the Bv field if non-nil, zero value otherwise.

### GetBvOk

`func (o *HistoricalNBBO) GetBvOk() (*[]float32, bool)`

GetBvOk returns a tuple with the Bv field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBv

`func (o *HistoricalNBBO) SetBv(v []float32)`

SetBv sets Bv field to given value.

### HasBv

`func (o *HistoricalNBBO) HasBv() bool`

HasBv returns a boolean if a field has been set.

### GetB

`func (o *HistoricalNBBO) GetB() []float32`

GetB returns the B field if non-nil, zero value otherwise.

### GetBOk

`func (o *HistoricalNBBO) GetBOk() (*[]float32, bool)`

GetBOk returns a tuple with the B field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetB

`func (o *HistoricalNBBO) SetB(v []float32)`

SetB sets B field to given value.

### HasB

`func (o *HistoricalNBBO) HasB() bool`

HasB returns a boolean if a field has been set.

### GetBx

`func (o *HistoricalNBBO) GetBx() []string`

GetBx returns the Bx field if non-nil, zero value otherwise.

### GetBxOk

`func (o *HistoricalNBBO) GetBxOk() (*[]string, bool)`

GetBxOk returns a tuple with the Bx field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBx

`func (o *HistoricalNBBO) SetBx(v []string)`

SetBx sets Bx field to given value.

### HasBx

`func (o *HistoricalNBBO) HasBx() bool`

HasBx returns a boolean if a field has been set.

### GetT

`func (o *HistoricalNBBO) GetT() []int64`

GetT returns the T field if non-nil, zero value otherwise.

### GetTOk

`func (o *HistoricalNBBO) GetTOk() (*[]int64, bool)`

GetTOk returns a tuple with the T field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetT

`func (o *HistoricalNBBO) SetT(v []int64)`

SetT sets T field to given value.

### HasT

`func (o *HistoricalNBBO) HasT() bool`

HasT returns a boolean if a field has been set.

### GetC

`func (o *HistoricalNBBO) GetC() [][]string`

GetC returns the C field if non-nil, zero value otherwise.

### GetCOk

`func (o *HistoricalNBBO) GetCOk() (*[][]string, bool)`

GetCOk returns a tuple with the C field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetC

`func (o *HistoricalNBBO) SetC(v [][]string)`

SetC sets C field to given value.

### HasC

`func (o *HistoricalNBBO) HasC() bool`

HasC returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


