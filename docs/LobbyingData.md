# LobbyingData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Symbol** | Pointer to **string** | Symbol. | [optional] 
**Name** | Pointer to **string** | Company&#39;s name. | [optional] 
**Description** | Pointer to **string** | Description. | [optional] 
**Country** | Pointer to **string** | Country. | [optional] 
**Year** | Pointer to **int64** | Year. | [optional] 
**Period** | Pointer to **string** | Period. | [optional] 
**Income** | Pointer to **float32** | Income reported by lobbying firms. | [optional] 
**Expenses** | Pointer to **float32** | Expenses reported by the company. | [optional] 
**DocumentUrl** | Pointer to **string** | Document&#39;s URL. | [optional] 
**PostedName** | Pointer to **string** | Posted name. | [optional] 
**Date** | Pointer to **string** | Date. | [optional] 
**ClientId** | Pointer to **string** | Client ID. | [optional] 
**RegistrantId** | Pointer to **string** | Registrant ID. | [optional] 
**SenateId** | Pointer to **string** | Senate ID. | [optional] 
**HouseregistrantId** | Pointer to **string** | House registrant ID. | [optional] 

## Methods

### NewLobbyingData

`func NewLobbyingData() *LobbyingData`

NewLobbyingData instantiates a new LobbyingData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLobbyingDataWithDefaults

`func NewLobbyingDataWithDefaults() *LobbyingData`

NewLobbyingDataWithDefaults instantiates a new LobbyingData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSymbol

`func (o *LobbyingData) GetSymbol() string`

GetSymbol returns the Symbol field if non-nil, zero value otherwise.

### GetSymbolOk

`func (o *LobbyingData) GetSymbolOk() (*string, bool)`

GetSymbolOk returns a tuple with the Symbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSymbol

`func (o *LobbyingData) SetSymbol(v string)`

SetSymbol sets Symbol field to given value.

### HasSymbol

`func (o *LobbyingData) HasSymbol() bool`

HasSymbol returns a boolean if a field has been set.

### GetName

`func (o *LobbyingData) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *LobbyingData) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *LobbyingData) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *LobbyingData) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *LobbyingData) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *LobbyingData) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *LobbyingData) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *LobbyingData) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetCountry

`func (o *LobbyingData) GetCountry() string`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *LobbyingData) GetCountryOk() (*string, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *LobbyingData) SetCountry(v string)`

SetCountry sets Country field to given value.

### HasCountry

`func (o *LobbyingData) HasCountry() bool`

HasCountry returns a boolean if a field has been set.

### GetYear

`func (o *LobbyingData) GetYear() int64`

GetYear returns the Year field if non-nil, zero value otherwise.

### GetYearOk

`func (o *LobbyingData) GetYearOk() (*int64, bool)`

GetYearOk returns a tuple with the Year field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetYear

`func (o *LobbyingData) SetYear(v int64)`

SetYear sets Year field to given value.

### HasYear

`func (o *LobbyingData) HasYear() bool`

HasYear returns a boolean if a field has been set.

### GetPeriod

`func (o *LobbyingData) GetPeriod() string`

GetPeriod returns the Period field if non-nil, zero value otherwise.

### GetPeriodOk

`func (o *LobbyingData) GetPeriodOk() (*string, bool)`

GetPeriodOk returns a tuple with the Period field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeriod

`func (o *LobbyingData) SetPeriod(v string)`

SetPeriod sets Period field to given value.

### HasPeriod

`func (o *LobbyingData) HasPeriod() bool`

HasPeriod returns a boolean if a field has been set.

### GetIncome

`func (o *LobbyingData) GetIncome() float32`

GetIncome returns the Income field if non-nil, zero value otherwise.

### GetIncomeOk

`func (o *LobbyingData) GetIncomeOk() (*float32, bool)`

GetIncomeOk returns a tuple with the Income field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncome

`func (o *LobbyingData) SetIncome(v float32)`

SetIncome sets Income field to given value.

### HasIncome

`func (o *LobbyingData) HasIncome() bool`

HasIncome returns a boolean if a field has been set.

### GetExpenses

`func (o *LobbyingData) GetExpenses() float32`

GetExpenses returns the Expenses field if non-nil, zero value otherwise.

### GetExpensesOk

`func (o *LobbyingData) GetExpensesOk() (*float32, bool)`

GetExpensesOk returns a tuple with the Expenses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpenses

`func (o *LobbyingData) SetExpenses(v float32)`

SetExpenses sets Expenses field to given value.

### HasExpenses

`func (o *LobbyingData) HasExpenses() bool`

HasExpenses returns a boolean if a field has been set.

### GetDocumentUrl

`func (o *LobbyingData) GetDocumentUrl() string`

GetDocumentUrl returns the DocumentUrl field if non-nil, zero value otherwise.

### GetDocumentUrlOk

`func (o *LobbyingData) GetDocumentUrlOk() (*string, bool)`

GetDocumentUrlOk returns a tuple with the DocumentUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentUrl

`func (o *LobbyingData) SetDocumentUrl(v string)`

SetDocumentUrl sets DocumentUrl field to given value.

### HasDocumentUrl

`func (o *LobbyingData) HasDocumentUrl() bool`

HasDocumentUrl returns a boolean if a field has been set.

### GetPostedName

`func (o *LobbyingData) GetPostedName() string`

GetPostedName returns the PostedName field if non-nil, zero value otherwise.

### GetPostedNameOk

`func (o *LobbyingData) GetPostedNameOk() (*string, bool)`

GetPostedNameOk returns a tuple with the PostedName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostedName

`func (o *LobbyingData) SetPostedName(v string)`

SetPostedName sets PostedName field to given value.

### HasPostedName

`func (o *LobbyingData) HasPostedName() bool`

HasPostedName returns a boolean if a field has been set.

### GetDate

`func (o *LobbyingData) GetDate() string`

GetDate returns the Date field if non-nil, zero value otherwise.

### GetDateOk

`func (o *LobbyingData) GetDateOk() (*string, bool)`

GetDateOk returns a tuple with the Date field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDate

`func (o *LobbyingData) SetDate(v string)`

SetDate sets Date field to given value.

### HasDate

`func (o *LobbyingData) HasDate() bool`

HasDate returns a boolean if a field has been set.

### GetClientId

`func (o *LobbyingData) GetClientId() string`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *LobbyingData) GetClientIdOk() (*string, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *LobbyingData) SetClientId(v string)`

SetClientId sets ClientId field to given value.

### HasClientId

`func (o *LobbyingData) HasClientId() bool`

HasClientId returns a boolean if a field has been set.

### GetRegistrantId

`func (o *LobbyingData) GetRegistrantId() string`

GetRegistrantId returns the RegistrantId field if non-nil, zero value otherwise.

### GetRegistrantIdOk

`func (o *LobbyingData) GetRegistrantIdOk() (*string, bool)`

GetRegistrantIdOk returns a tuple with the RegistrantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegistrantId

`func (o *LobbyingData) SetRegistrantId(v string)`

SetRegistrantId sets RegistrantId field to given value.

### HasRegistrantId

`func (o *LobbyingData) HasRegistrantId() bool`

HasRegistrantId returns a boolean if a field has been set.

### GetSenateId

`func (o *LobbyingData) GetSenateId() string`

GetSenateId returns the SenateId field if non-nil, zero value otherwise.

### GetSenateIdOk

`func (o *LobbyingData) GetSenateIdOk() (*string, bool)`

GetSenateIdOk returns a tuple with the SenateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSenateId

`func (o *LobbyingData) SetSenateId(v string)`

SetSenateId sets SenateId field to given value.

### HasSenateId

`func (o *LobbyingData) HasSenateId() bool`

HasSenateId returns a boolean if a field has been set.

### GetHouseregistrantId

`func (o *LobbyingData) GetHouseregistrantId() string`

GetHouseregistrantId returns the HouseregistrantId field if non-nil, zero value otherwise.

### GetHouseregistrantIdOk

`func (o *LobbyingData) GetHouseregistrantIdOk() (*string, bool)`

GetHouseregistrantIdOk returns a tuple with the HouseregistrantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHouseregistrantId

`func (o *LobbyingData) SetHouseregistrantId(v string)`

SetHouseregistrantId sets HouseregistrantId field to given value.

### HasHouseregistrantId

`func (o *LobbyingData) HasHouseregistrantId() bool`

HasHouseregistrantId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


