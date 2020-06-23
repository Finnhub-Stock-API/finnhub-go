# TickData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**S** | **string** | Symbol. | [optional] 
**Skip** | **int64** | Number of ticks skipped. | [optional] 
**Count** | **int64** | Number of ticks returned. If &lt;code&gt;count&lt;/code&gt; &lt; &lt;code&gt;limit&lt;/code&gt;, all data for that date has been returned. | [optional] 
**Total** | **int64** | Total number of ticks for that date. | [optional] 
**V** | **[]float32** | List of volume data. | [optional] 
**P** | **[]float32** | List of price data. | [optional] 
**T** | **[]int64** | List of timestamp in UNIX ms. | [optional] 
**X** | **[]string** | List of venues/exchanges. | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


