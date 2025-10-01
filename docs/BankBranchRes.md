# BankBranchRes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**[]BankBranchData**](BankBranchData.md) | Array of branches. | [optional] 
**Symbol** | Pointer to **string** | Symbol | [optional] 

## Methods

### NewBankBranchRes

`func NewBankBranchRes() *BankBranchRes`

NewBankBranchRes instantiates a new BankBranchRes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBankBranchResWithDefaults

`func NewBankBranchResWithDefaults() *BankBranchRes`

NewBankBranchResWithDefaults instantiates a new BankBranchRes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *BankBranchRes) GetData() []BankBranchData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *BankBranchRes) GetDataOk() (*[]BankBranchData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *BankBranchRes) SetData(v []BankBranchData)`

SetData sets Data field to given value.

### HasData

`func (o *BankBranchRes) HasData() bool`

HasData returns a boolean if a field has been set.

### GetSymbol

`func (o *BankBranchRes) GetSymbol() string`

GetSymbol returns the Symbol field if non-nil, zero value otherwise.

### GetSymbolOk

`func (o *BankBranchRes) GetSymbolOk() (*string, bool)`

GetSymbolOk returns a tuple with the Symbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSymbol

`func (o *BankBranchRes) SetSymbol(v string)`

SetSymbol sets Symbol field to given value.

### HasSymbol

`func (o *BankBranchRes) HasSymbol() bool`

HasSymbol returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


