# KeyCustomersSuppliers

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Symbol** | Pointer to **string** | Symbol | [optional] 
**Name** | Pointer to **string** | Name | [optional] 
**Country** | Pointer to **string** | Country | [optional] 
**Industry** | Pointer to **string** | Industry | [optional] 
**Customer** | Pointer to **bool** | Whether the company is a customer. | [optional] 
**Supplier** | Pointer to **bool** | Whether the company is a supplier | [optional] 
**OneMonthCorrelation** | Pointer to **float32** | 1-month price correlation | [optional] 
**OneYearCorrelation** | Pointer to **float32** | 1-year price correlation | [optional] 
**SixMonthCorrelation** | Pointer to **float32** | 6-month price correlation | [optional] 
**ThreeMonthCorrelation** | Pointer to **float32** | 3-month price correlation | [optional] 
**TwoWeekCorrelation** | Pointer to **float32** | 2-week price correlation | [optional] 
**TwoYearCorrelation** | Pointer to **float32** | 2-year price correlation | [optional] 

## Methods

### NewKeyCustomersSuppliers

`func NewKeyCustomersSuppliers() *KeyCustomersSuppliers`

NewKeyCustomersSuppliers instantiates a new KeyCustomersSuppliers object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKeyCustomersSuppliersWithDefaults

`func NewKeyCustomersSuppliersWithDefaults() *KeyCustomersSuppliers`

NewKeyCustomersSuppliersWithDefaults instantiates a new KeyCustomersSuppliers object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSymbol

`func (o *KeyCustomersSuppliers) GetSymbol() string`

GetSymbol returns the Symbol field if non-nil, zero value otherwise.

### GetSymbolOk

`func (o *KeyCustomersSuppliers) GetSymbolOk() (*string, bool)`

GetSymbolOk returns a tuple with the Symbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSymbol

`func (o *KeyCustomersSuppliers) SetSymbol(v string)`

SetSymbol sets Symbol field to given value.

### HasSymbol

`func (o *KeyCustomersSuppliers) HasSymbol() bool`

HasSymbol returns a boolean if a field has been set.

### GetName

`func (o *KeyCustomersSuppliers) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *KeyCustomersSuppliers) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *KeyCustomersSuppliers) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *KeyCustomersSuppliers) HasName() bool`

HasName returns a boolean if a field has been set.

### GetCountry

`func (o *KeyCustomersSuppliers) GetCountry() string`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *KeyCustomersSuppliers) GetCountryOk() (*string, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *KeyCustomersSuppliers) SetCountry(v string)`

SetCountry sets Country field to given value.

### HasCountry

`func (o *KeyCustomersSuppliers) HasCountry() bool`

HasCountry returns a boolean if a field has been set.

### GetIndustry

`func (o *KeyCustomersSuppliers) GetIndustry() string`

GetIndustry returns the Industry field if non-nil, zero value otherwise.

### GetIndustryOk

`func (o *KeyCustomersSuppliers) GetIndustryOk() (*string, bool)`

GetIndustryOk returns a tuple with the Industry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndustry

`func (o *KeyCustomersSuppliers) SetIndustry(v string)`

SetIndustry sets Industry field to given value.

### HasIndustry

`func (o *KeyCustomersSuppliers) HasIndustry() bool`

HasIndustry returns a boolean if a field has been set.

### GetCustomer

`func (o *KeyCustomersSuppliers) GetCustomer() bool`

GetCustomer returns the Customer field if non-nil, zero value otherwise.

### GetCustomerOk

`func (o *KeyCustomersSuppliers) GetCustomerOk() (*bool, bool)`

GetCustomerOk returns a tuple with the Customer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomer

`func (o *KeyCustomersSuppliers) SetCustomer(v bool)`

SetCustomer sets Customer field to given value.

### HasCustomer

`func (o *KeyCustomersSuppliers) HasCustomer() bool`

HasCustomer returns a boolean if a field has been set.

### GetSupplier

`func (o *KeyCustomersSuppliers) GetSupplier() bool`

GetSupplier returns the Supplier field if non-nil, zero value otherwise.

### GetSupplierOk

`func (o *KeyCustomersSuppliers) GetSupplierOk() (*bool, bool)`

GetSupplierOk returns a tuple with the Supplier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupplier

`func (o *KeyCustomersSuppliers) SetSupplier(v bool)`

SetSupplier sets Supplier field to given value.

### HasSupplier

`func (o *KeyCustomersSuppliers) HasSupplier() bool`

HasSupplier returns a boolean if a field has been set.

### GetOneMonthCorrelation

`func (o *KeyCustomersSuppliers) GetOneMonthCorrelation() float32`

GetOneMonthCorrelation returns the OneMonthCorrelation field if non-nil, zero value otherwise.

### GetOneMonthCorrelationOk

`func (o *KeyCustomersSuppliers) GetOneMonthCorrelationOk() (*float32, bool)`

GetOneMonthCorrelationOk returns a tuple with the OneMonthCorrelation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOneMonthCorrelation

`func (o *KeyCustomersSuppliers) SetOneMonthCorrelation(v float32)`

SetOneMonthCorrelation sets OneMonthCorrelation field to given value.

### HasOneMonthCorrelation

`func (o *KeyCustomersSuppliers) HasOneMonthCorrelation() bool`

HasOneMonthCorrelation returns a boolean if a field has been set.

### GetOneYearCorrelation

`func (o *KeyCustomersSuppliers) GetOneYearCorrelation() float32`

GetOneYearCorrelation returns the OneYearCorrelation field if non-nil, zero value otherwise.

### GetOneYearCorrelationOk

`func (o *KeyCustomersSuppliers) GetOneYearCorrelationOk() (*float32, bool)`

GetOneYearCorrelationOk returns a tuple with the OneYearCorrelation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOneYearCorrelation

`func (o *KeyCustomersSuppliers) SetOneYearCorrelation(v float32)`

SetOneYearCorrelation sets OneYearCorrelation field to given value.

### HasOneYearCorrelation

`func (o *KeyCustomersSuppliers) HasOneYearCorrelation() bool`

HasOneYearCorrelation returns a boolean if a field has been set.

### GetSixMonthCorrelation

`func (o *KeyCustomersSuppliers) GetSixMonthCorrelation() float32`

GetSixMonthCorrelation returns the SixMonthCorrelation field if non-nil, zero value otherwise.

### GetSixMonthCorrelationOk

`func (o *KeyCustomersSuppliers) GetSixMonthCorrelationOk() (*float32, bool)`

GetSixMonthCorrelationOk returns a tuple with the SixMonthCorrelation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSixMonthCorrelation

`func (o *KeyCustomersSuppliers) SetSixMonthCorrelation(v float32)`

SetSixMonthCorrelation sets SixMonthCorrelation field to given value.

### HasSixMonthCorrelation

`func (o *KeyCustomersSuppliers) HasSixMonthCorrelation() bool`

HasSixMonthCorrelation returns a boolean if a field has been set.

### GetThreeMonthCorrelation

`func (o *KeyCustomersSuppliers) GetThreeMonthCorrelation() float32`

GetThreeMonthCorrelation returns the ThreeMonthCorrelation field if non-nil, zero value otherwise.

### GetThreeMonthCorrelationOk

`func (o *KeyCustomersSuppliers) GetThreeMonthCorrelationOk() (*float32, bool)`

GetThreeMonthCorrelationOk returns a tuple with the ThreeMonthCorrelation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThreeMonthCorrelation

`func (o *KeyCustomersSuppliers) SetThreeMonthCorrelation(v float32)`

SetThreeMonthCorrelation sets ThreeMonthCorrelation field to given value.

### HasThreeMonthCorrelation

`func (o *KeyCustomersSuppliers) HasThreeMonthCorrelation() bool`

HasThreeMonthCorrelation returns a boolean if a field has been set.

### GetTwoWeekCorrelation

`func (o *KeyCustomersSuppliers) GetTwoWeekCorrelation() float32`

GetTwoWeekCorrelation returns the TwoWeekCorrelation field if non-nil, zero value otherwise.

### GetTwoWeekCorrelationOk

`func (o *KeyCustomersSuppliers) GetTwoWeekCorrelationOk() (*float32, bool)`

GetTwoWeekCorrelationOk returns a tuple with the TwoWeekCorrelation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTwoWeekCorrelation

`func (o *KeyCustomersSuppliers) SetTwoWeekCorrelation(v float32)`

SetTwoWeekCorrelation sets TwoWeekCorrelation field to given value.

### HasTwoWeekCorrelation

`func (o *KeyCustomersSuppliers) HasTwoWeekCorrelation() bool`

HasTwoWeekCorrelation returns a boolean if a field has been set.

### GetTwoYearCorrelation

`func (o *KeyCustomersSuppliers) GetTwoYearCorrelation() float32`

GetTwoYearCorrelation returns the TwoYearCorrelation field if non-nil, zero value otherwise.

### GetTwoYearCorrelationOk

`func (o *KeyCustomersSuppliers) GetTwoYearCorrelationOk() (*float32, bool)`

GetTwoYearCorrelationOk returns a tuple with the TwoYearCorrelation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTwoYearCorrelation

`func (o *KeyCustomersSuppliers) SetTwoYearCorrelation(v float32)`

SetTwoYearCorrelation sets TwoYearCorrelation field to given value.

### HasTwoYearCorrelation

`func (o *KeyCustomersSuppliers) HasTwoYearCorrelation() bool`

HasTwoYearCorrelation returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


