# CongressionalTransaction

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AmountFrom** | Pointer to **float32** | Transaction amount from. | [optional] 
**AmountTo** | Pointer to **float32** | Transaction amount to. | [optional] 
**AssetName** | Pointer to **string** | Asset name. | [optional] 
**FilingDate** | Pointer to **string** | Filing date. | [optional] 
**Name** | Pointer to **string** | Name of the representative. | [optional] 
**OwnerType** | Pointer to **string** | Owner Type. | [optional] 
**Position** | Pointer to **string** | Position. | [optional] 
**Symbol** | Pointer to **string** | Symbol. | [optional] 
**TransactionDate** | Pointer to **string** | Transaction date. | [optional] 
**TransactionType** | Pointer to **string** | Transaction type &lt;code&gt;Sale&lt;/code&gt; or &lt;code&gt;Purchase&lt;/code&gt;. | [optional] 

## Methods

### NewCongressionalTransaction

`func NewCongressionalTransaction() *CongressionalTransaction`

NewCongressionalTransaction instantiates a new CongressionalTransaction object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCongressionalTransactionWithDefaults

`func NewCongressionalTransactionWithDefaults() *CongressionalTransaction`

NewCongressionalTransactionWithDefaults instantiates a new CongressionalTransaction object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmountFrom

`func (o *CongressionalTransaction) GetAmountFrom() float32`

GetAmountFrom returns the AmountFrom field if non-nil, zero value otherwise.

### GetAmountFromOk

`func (o *CongressionalTransaction) GetAmountFromOk() (*float32, bool)`

GetAmountFromOk returns a tuple with the AmountFrom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmountFrom

`func (o *CongressionalTransaction) SetAmountFrom(v float32)`

SetAmountFrom sets AmountFrom field to given value.

### HasAmountFrom

`func (o *CongressionalTransaction) HasAmountFrom() bool`

HasAmountFrom returns a boolean if a field has been set.

### GetAmountTo

`func (o *CongressionalTransaction) GetAmountTo() float32`

GetAmountTo returns the AmountTo field if non-nil, zero value otherwise.

### GetAmountToOk

`func (o *CongressionalTransaction) GetAmountToOk() (*float32, bool)`

GetAmountToOk returns a tuple with the AmountTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmountTo

`func (o *CongressionalTransaction) SetAmountTo(v float32)`

SetAmountTo sets AmountTo field to given value.

### HasAmountTo

`func (o *CongressionalTransaction) HasAmountTo() bool`

HasAmountTo returns a boolean if a field has been set.

### GetAssetName

`func (o *CongressionalTransaction) GetAssetName() string`

GetAssetName returns the AssetName field if non-nil, zero value otherwise.

### GetAssetNameOk

`func (o *CongressionalTransaction) GetAssetNameOk() (*string, bool)`

GetAssetNameOk returns a tuple with the AssetName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssetName

`func (o *CongressionalTransaction) SetAssetName(v string)`

SetAssetName sets AssetName field to given value.

### HasAssetName

`func (o *CongressionalTransaction) HasAssetName() bool`

HasAssetName returns a boolean if a field has been set.

### GetFilingDate

`func (o *CongressionalTransaction) GetFilingDate() string`

GetFilingDate returns the FilingDate field if non-nil, zero value otherwise.

### GetFilingDateOk

`func (o *CongressionalTransaction) GetFilingDateOk() (*string, bool)`

GetFilingDateOk returns a tuple with the FilingDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilingDate

`func (o *CongressionalTransaction) SetFilingDate(v string)`

SetFilingDate sets FilingDate field to given value.

### HasFilingDate

`func (o *CongressionalTransaction) HasFilingDate() bool`

HasFilingDate returns a boolean if a field has been set.

### GetName

`func (o *CongressionalTransaction) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CongressionalTransaction) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CongressionalTransaction) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CongressionalTransaction) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOwnerType

`func (o *CongressionalTransaction) GetOwnerType() string`

GetOwnerType returns the OwnerType field if non-nil, zero value otherwise.

### GetOwnerTypeOk

`func (o *CongressionalTransaction) GetOwnerTypeOk() (*string, bool)`

GetOwnerTypeOk returns a tuple with the OwnerType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerType

`func (o *CongressionalTransaction) SetOwnerType(v string)`

SetOwnerType sets OwnerType field to given value.

### HasOwnerType

`func (o *CongressionalTransaction) HasOwnerType() bool`

HasOwnerType returns a boolean if a field has been set.

### GetPosition

`func (o *CongressionalTransaction) GetPosition() string`

GetPosition returns the Position field if non-nil, zero value otherwise.

### GetPositionOk

`func (o *CongressionalTransaction) GetPositionOk() (*string, bool)`

GetPositionOk returns a tuple with the Position field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPosition

`func (o *CongressionalTransaction) SetPosition(v string)`

SetPosition sets Position field to given value.

### HasPosition

`func (o *CongressionalTransaction) HasPosition() bool`

HasPosition returns a boolean if a field has been set.

### GetSymbol

`func (o *CongressionalTransaction) GetSymbol() string`

GetSymbol returns the Symbol field if non-nil, zero value otherwise.

### GetSymbolOk

`func (o *CongressionalTransaction) GetSymbolOk() (*string, bool)`

GetSymbolOk returns a tuple with the Symbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSymbol

`func (o *CongressionalTransaction) SetSymbol(v string)`

SetSymbol sets Symbol field to given value.

### HasSymbol

`func (o *CongressionalTransaction) HasSymbol() bool`

HasSymbol returns a boolean if a field has been set.

### GetTransactionDate

`func (o *CongressionalTransaction) GetTransactionDate() string`

GetTransactionDate returns the TransactionDate field if non-nil, zero value otherwise.

### GetTransactionDateOk

`func (o *CongressionalTransaction) GetTransactionDateOk() (*string, bool)`

GetTransactionDateOk returns a tuple with the TransactionDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionDate

`func (o *CongressionalTransaction) SetTransactionDate(v string)`

SetTransactionDate sets TransactionDate field to given value.

### HasTransactionDate

`func (o *CongressionalTransaction) HasTransactionDate() bool`

HasTransactionDate returns a boolean if a field has been set.

### GetTransactionType

`func (o *CongressionalTransaction) GetTransactionType() string`

GetTransactionType returns the TransactionType field if non-nil, zero value otherwise.

### GetTransactionTypeOk

`func (o *CongressionalTransaction) GetTransactionTypeOk() (*string, bool)`

GetTransactionTypeOk returns a tuple with the TransactionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionType

`func (o *CongressionalTransaction) SetTransactionType(v string)`

SetTransactionType sets TransactionType field to given value.

### HasTransactionType

`func (o *CongressionalTransaction) HasTransactionType() bool`

HasTransactionType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


