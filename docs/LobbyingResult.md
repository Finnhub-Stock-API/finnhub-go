# LobbyingResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Symbol** | Pointer to **string** | Symbol. | [optional] 
**Data** | Pointer to [**[]LobbyingData**](LobbyingData.md) | Array of lobbying activities. | [optional] 

## Methods

### NewLobbyingResult

`func NewLobbyingResult() *LobbyingResult`

NewLobbyingResult instantiates a new LobbyingResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLobbyingResultWithDefaults

`func NewLobbyingResultWithDefaults() *LobbyingResult`

NewLobbyingResultWithDefaults instantiates a new LobbyingResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSymbol

`func (o *LobbyingResult) GetSymbol() string`

GetSymbol returns the Symbol field if non-nil, zero value otherwise.

### GetSymbolOk

`func (o *LobbyingResult) GetSymbolOk() (*string, bool)`

GetSymbolOk returns a tuple with the Symbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSymbol

`func (o *LobbyingResult) SetSymbol(v string)`

SetSymbol sets Symbol field to given value.

### HasSymbol

`func (o *LobbyingResult) HasSymbol() bool`

HasSymbol returns a boolean if a field has been set.

### GetData

`func (o *LobbyingResult) GetData() []LobbyingData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *LobbyingResult) GetDataOk() (*[]LobbyingData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *LobbyingResult) SetData(v []LobbyingData)`

SetData sets Data field to given value.

### HasData

`func (o *LobbyingResult) HasData() bool`

HasData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


