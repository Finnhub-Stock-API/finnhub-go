# \DefaultApi

All URIs are relative to *https://finnhub.io/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AggregateIndicator**](DefaultApi.md#AggregateIndicator) | **Get** /scan/technical-indicator | Aggregate Indicators
[**CompanyBasicFinancials**](DefaultApi.md#CompanyBasicFinancials) | **Get** /stock/metric | Basic Financials
[**CompanyEarnings**](DefaultApi.md#CompanyEarnings) | **Get** /stock/earnings | Earnings Surprises
[**CompanyEpsEstimates**](DefaultApi.md#CompanyEpsEstimates) | **Get** /stock/eps-estimate | Earnings Estimates
[**CompanyExecutive**](DefaultApi.md#CompanyExecutive) | **Get** /stock/executive | Company Executive
[**CompanyNews**](DefaultApi.md#CompanyNews) | **Get** /company-news | Company News
[**CompanyPeers**](DefaultApi.md#CompanyPeers) | **Get** /stock/peers | Peers
[**CompanyProfile**](DefaultApi.md#CompanyProfile) | **Get** /stock/profile | Company Profile
[**CompanyProfile2**](DefaultApi.md#CompanyProfile2) | **Get** /stock/profile2 | Company Profile 2
[**CompanyRevenueEstimates**](DefaultApi.md#CompanyRevenueEstimates) | **Get** /stock/revenue-estimate | Revenue Estimates
[**Covid19**](DefaultApi.md#Covid19) | **Get** /covid19/us | COVID-19
[**CryptoCandles**](DefaultApi.md#CryptoCandles) | **Get** /crypto/candle | Crypto Candles
[**CryptoExchanges**](DefaultApi.md#CryptoExchanges) | **Get** /crypto/exchange | Crypto Exchanges
[**CryptoSymbols**](DefaultApi.md#CryptoSymbols) | **Get** /crypto/symbol | Crypto Symbol
[**EarningsCalendar**](DefaultApi.md#EarningsCalendar) | **Get** /calendar/earnings | Earnings Calendar
[**Filings**](DefaultApi.md#Filings) | **Get** /stock/filings | Filings
[**Financials**](DefaultApi.md#Financials) | **Get** /stock/financials | Financial Statements
[**FinancialsReported**](DefaultApi.md#FinancialsReported) | **Get** /stock/financials-reported | Financials As Reported
[**ForexCandles**](DefaultApi.md#ForexCandles) | **Get** /forex/candle | Forex Candles
[**ForexExchanges**](DefaultApi.md#ForexExchanges) | **Get** /forex/exchange | Forex Exchanges
[**ForexRates**](DefaultApi.md#ForexRates) | **Get** /forex/rates | Forex rates
[**ForexSymbols**](DefaultApi.md#ForexSymbols) | **Get** /forex/symbol | Forex Symbol
[**FundOwnership**](DefaultApi.md#FundOwnership) | **Get** /stock/fund-ownership | Fund Ownership
[**GeneralNews**](DefaultApi.md#GeneralNews) | **Get** /news | General News
[**InvestorsOwnership**](DefaultApi.md#InvestorsOwnership) | **Get** /stock/investor-ownership | Investors Ownership
[**IpoCalendar**](DefaultApi.md#IpoCalendar) | **Get** /calendar/ipo | IPO Calendar
[**MajorDevelopments**](DefaultApi.md#MajorDevelopments) | **Get** /major-development | Major Developments
[**NewsSentiment**](DefaultApi.md#NewsSentiment) | **Get** /news-sentiment | News Sentiment
[**PatternRecognition**](DefaultApi.md#PatternRecognition) | **Get** /scan/pattern | Pattern Recognition
[**PriceTarget**](DefaultApi.md#PriceTarget) | **Get** /stock/price-target | Price Target
[**Quote**](DefaultApi.md#Quote) | **Get** /quote | Quote
[**RecommendationTrends**](DefaultApi.md#RecommendationTrends) | **Get** /stock/recommendation | Recommendation Trends
[**StockCandles**](DefaultApi.md#StockCandles) | **Get** /stock/candle | Stock Candles
[**StockDividends**](DefaultApi.md#StockDividends) | **Get** /stock/dividend | Dividends
[**StockSplits**](DefaultApi.md#StockSplits) | **Get** /stock/split | Splits
[**StockSymbols**](DefaultApi.md#StockSymbols) | **Get** /stock/symbol | Stock Symbol
[**StockTick**](DefaultApi.md#StockTick) | **Get** /stock/tick | Tick Data
[**SupportResistance**](DefaultApi.md#SupportResistance) | **Get** /scan/support-resistance | Support/Resistance
[**TechnicalIndicator**](DefaultApi.md#TechnicalIndicator) | **Get** /indicator | Technical Indicators
[**Transcripts**](DefaultApi.md#Transcripts) | **Get** /stock/transcripts | Earnings Call Transcripts
[**TranscriptsList**](DefaultApi.md#TranscriptsList) | **Get** /stock/transcripts/list | Earnings Call Transcripts List
[**UpgradeDowngrade**](DefaultApi.md#UpgradeDowngrade) | **Get** /stock/upgrade-downgrade | Stock Upgrade/Downgrade



## AggregateIndicator

> AggregateIndicators AggregateIndicator(ctx, symbol, resolution)

Aggregate Indicators

Get aggregate signal of multiple technical indicators such as MACD, RSI, Moving Average v.v.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**symbol** | **string**| symbol | 
**resolution** | **string**| Supported resolution includes &lt;code&gt;1, 5, 15, 30, 60, D, W, M &lt;/code&gt;.Some timeframes might not be available depending on the exchange. | 

### Return type

[**AggregateIndicators**](AggregateIndicators.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CompanyBasicFinancials

> BasicFinancials CompanyBasicFinancials(ctx, symbol, metric)

Basic Financials

Get company basic financials such as margin, P/E ratio, 52-week high/low etc.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**symbol** | **string**| Symbol of the company: AAPL. | 
**metric** | **string**| Metric type. Can be 1 of the following values &lt;code&gt;all, price, valuation, margin&lt;/code&gt; | 

### Return type

[**BasicFinancials**](BasicFinancials.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CompanyEarnings

> []EarningResult CompanyEarnings(ctx, symbol, optional)

Earnings Surprises

Get company historical quarterly earnings surprise going back to 2000.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**symbol** | **string**| Symbol of the company: AAPL. | 
 **optional** | ***CompanyEarningsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a CompanyEarningsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **optional.Int64**| Limit number of period returned. Leave blank to get the full history. | 

### Return type

[**[]EarningResult**](EarningResult.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CompanyEpsEstimates

> EarningsEstimates CompanyEpsEstimates(ctx, symbol, optional)

Earnings Estimates

Get company's EPS estimates.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**symbol** | **string**| Symbol of the company: AAPL. | 
 **optional** | ***CompanyEpsEstimatesOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a CompanyEpsEstimatesOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **freq** | **optional.String**| Can take 1 of the following values: &lt;code&gt;annual, quarterly&lt;/code&gt;. Default to &lt;code&gt;quarterly&lt;/code&gt; | 

### Return type

[**EarningsEstimates**](EarningsEstimates.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CompanyExecutive

> CompanyExecutive CompanyExecutive(ctx, symbol)

Company Executive

Get a list of company's executives and members of the Board.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**symbol** | **string**| Symbol of the company: AAPL. | 

### Return type

[**CompanyExecutive**](CompanyExecutive.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CompanyNews

> []News CompanyNews(ctx, symbol, from, to)

Company News

List latest company news by symbol. This endpoint is only available for US companies.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**symbol** | **string**| Company symbol. | 
**from** | **string**| From date &lt;code&gt;YYYY-MM-DD&lt;/code&gt;. | 
**to** | **string**| To date &lt;code&gt;YYYY-MM-DD&lt;/code&gt;. | 

### Return type

[**[]News**](News.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CompanyPeers

> []string CompanyPeers(ctx, symbol)

Peers

Get company peers. Return a list of peers in the same country and GICS sub-industry

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**symbol** | **string**| Symbol of the company: AAPL. | 

### Return type

**[]string**

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CompanyProfile

> CompanyProfile CompanyProfile(ctx, optional)

Company Profile

Get general information of a company. You can query by symbol, ISIN or CUSIP

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***CompanyProfileOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a CompanyProfileOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **symbol** | **optional.String**| Symbol of the company: AAPL, SBIN.NS e.g. | 
 **isin** | **optional.String**| ISIN | 
 **cusip** | **optional.String**| CUSIP | 

### Return type

[**CompanyProfile**](CompanyProfile.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CompanyProfile2

> CompanyProfile2 CompanyProfile2(ctx, optional)

Company Profile 2

Get general information of a company. You can query by symbol, ISIN or CUSIP. This is the free version of <a href=\"#company-profile\">Company Profile</a>.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***CompanyProfile2Opts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a CompanyProfile2Opts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **symbol** | **optional.String**| Symbol of the company: AAPL, SBIN.NS e.g. | 
 **isin** | **optional.String**| ISIN | 
 **cusip** | **optional.String**| CUSIP | 

### Return type

[**CompanyProfile2**](CompanyProfile2.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CompanyRevenueEstimates

> RevenueEstimates CompanyRevenueEstimates(ctx, symbol, optional)

Revenue Estimates

Get company's revenue estimates.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**symbol** | **string**| Symbol of the company: AAPL. | 
 **optional** | ***CompanyRevenueEstimatesOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a CompanyRevenueEstimatesOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **freq** | **optional.String**| Can take 1 of the following values: &lt;code&gt;annual, quarterly&lt;/code&gt;. Default to &lt;code&gt;quarterly&lt;/code&gt; | 

### Return type

[**RevenueEstimates**](RevenueEstimates.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Covid19

> Covid19 Covid19(ctx, )

COVID-19

Get real-time updates on the number of COVID-19 (Corona virus) cases in the US with a state-by-state breakdown. Data is sourced from CDC and reputable sources. You can also access this API <a href=\"https://rapidapi.com/Finnhub/api/finnhub-real-time-covid-19\" target=\"_blank\" rel=\"nofollow\">here</a>

### Required Parameters

This endpoint does not need any parameter.

### Return type

[**Covid19**](COVID-19.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CryptoCandles

> CryptoCandles CryptoCandles(ctx, symbol, resolution, optional)

Crypto Candles

Get candlestick data for crypto symbols.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**symbol** | **string**| Use symbol returned in &lt;code&gt;/crypto/symbol&lt;/code&gt; endpoint for this field. | 
**resolution** | **string**| Supported resolution includes &lt;code&gt;1, 5, 15, 30, 60, D, W, M &lt;/code&gt;.Some timeframes might not be available depending on the exchange. | 
 **optional** | ***CryptoCandlesOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a CryptoCandlesOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **from** | **optional.Int64**| UNIX timestamp. Interval initial value. If count is not provided, this field is required | 
 **to** | **optional.Int64**| UNIX timestamp. Interval end value. If count is not provided, this field is required | 
 **format** | **optional.String**| By default, &lt;code&gt;format&#x3D;json&lt;/code&gt;. Strings &lt;code&gt;json&lt;/code&gt; and &lt;code&gt;csv&lt;/code&gt; are accepted. | 
 **count** | **optional.Int64**| Shortcut to set &lt;code&gt;to&#x3D;Unix.Now&lt;/code&gt; and &lt;code&gt;from&#x3D;Unix.Now - count * resolution_second&lt;/code&gt;. | 

### Return type

[**CryptoCandles**](CryptoCandles.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CryptoExchanges

> []string CryptoExchanges(ctx, )

Crypto Exchanges

List supported crypto exchanges

### Required Parameters

This endpoint does not need any parameter.

### Return type

**[]string**

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CryptoSymbols

> []CryptoSymbol CryptoSymbols(ctx, exchange)

Crypto Symbol

List supported crypto symbols by exchange

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**exchange** | **string**| Exchange you want to get the list of symbols from. | 

### Return type

[**[]CryptoSymbol**](CryptoSymbol.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EarningsCalendar

> []EarningRelease EarningsCalendar(ctx, optional)

Earnings Calendar

Get historical and coming earnings release dating back to 2003. You can setup <a href=\"#webhook\">webhook</a> to receive real-time earnings update.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***EarningsCalendarOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a EarningsCalendarOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **from** | **optional.String**| From date: 2020-03-15. | 
 **to** | **optional.String**| To date: 2020-03-16. | 
 **symbol** | **optional.String**| Filter by symbol: AAPL. | 
 **international** | [**optional.Interface of interface{}**](.md)| Set to &lt;code&gt;true&lt;/code&gt; to include international markets. Default value is &lt;code&gt;false&lt;/code&gt; | 

### Return type

[**[]EarningRelease**](EarningRelease.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Filings

> []Filing Filings(ctx, optional)

Filings

List company's filing. Limit to 250 documents at a time. This data is available for bulk download on <a href=\"https://www.kaggle.com/finnhub/sec-filings\" target=\"_blank\">Kaggle SEC Filings database</a>.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***FilingsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a FilingsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **symbol** | **optional.String**| Symbol. Leave &lt;code&gt;symbol&lt;/code&gt;,&lt;code&gt;cik&lt;/code&gt; and &lt;code&gt;accessNumber&lt;/code&gt; empty to list latest filings. | 
 **cik** | **optional.String**| CIK. | 
 **accessNumber** | **optional.String**| Access number of a specific report you want to retrieve data from. | 
 **form** | **optional.String**| Filter by form. You can use this value &lt;code&gt;NT 10-K&lt;/code&gt; to find non-timely filings for a company. | 
 **from** | **optional.String**| From date: 2020-03-15. | 
 **to** | **optional.String**| To date: 2020-03-16. | 

### Return type

[**[]Filing**](Filing.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Financials

> FinancialStatements Financials(ctx, symbol, statement, freq)

Financial Statements

Get standardized balance sheet, income statement and cash flow for global companies going back 30+ years.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**symbol** | **string**| Symbol of the company: AAPL. | 
**statement** | **string**| Statement can take 1 of these values &lt;code&gt;bs, ic, cf&lt;/code&gt; for Balance Sheet, Income Statement, Cash Flow respectively. | 
**freq** | **string**| Frequency can take 1 of these values &lt;code&gt;annual, quarterly, ttm, ytd&lt;/code&gt;.  TTM (Trailing Twelve Months) option is available for Income Statement and Cash Flow. YTD (Year To Date) option is only available for Cash Flow. | 

### Return type

[**FinancialStatements**](FinancialStatements.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FinancialsReported

> FinancialsAsReported FinancialsReported(ctx, optional)

Financials As Reported

Get financials as reported. This data is available for bulk download on <a href=\"https://www.kaggle.com/finnhub/reported-financials\" target=\"_blank\">Kaggle SEC Financials database</a>.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***FinancialsReportedOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a FinancialsReportedOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **symbol** | **optional.String**| Symbol. | 
 **cik** | **optional.String**| CIK. | 
 **accessNumber** | **optional.String**| Access number of a specific report you want to retrieve financials from. | 
 **freq** | **optional.String**| Frequency. Can be either &lt;code&gt;annual&lt;/code&gt; or &lt;code&gt;quarterly&lt;/code&gt;. Default to &lt;code&gt;annual&lt;/code&gt;. | 

### Return type

[**FinancialsAsReported**](FinancialsAsReported.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ForexCandles

> ForexCandles ForexCandles(ctx, symbol, resolution, optional)

Forex Candles

Get candlestick data for forex symbols.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**symbol** | **string**| Use symbol returned in &lt;code&gt;/forex/symbol&lt;/code&gt; endpoint for this field. | 
**resolution** | **string**| Supported resolution includes &lt;code&gt;1, 5, 15, 30, 60, D, W, M &lt;/code&gt;.Some timeframes might not be available depending on the exchange. | 
 **optional** | ***ForexCandlesOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ForexCandlesOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **from** | **optional.Int64**| UNIX timestamp. Interval initial value. If count is not provided, this field is required | 
 **to** | **optional.Int64**| UNIX timestamp. Interval end value. If count is not provided, this field is required | 
 **format** | **optional.String**| By default, &lt;code&gt;format&#x3D;json&lt;/code&gt;. Strings &lt;code&gt;json&lt;/code&gt; and &lt;code&gt;csv&lt;/code&gt; are accepted. | 

### Return type

[**ForexCandles**](ForexCandles.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ForexExchanges

> []string ForexExchanges(ctx, )

Forex Exchanges

List supported forex exchanges

### Required Parameters

This endpoint does not need any parameter.

### Return type

**[]string**

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ForexRates

> Forexrates ForexRates(ctx, optional)

Forex rates

Get rates for all forex pairs. Ideal for currency conversion

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ForexRatesOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ForexRatesOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **base** | **optional.String**| Base currency. Default to EUR. | 

### Return type

[**Forexrates**](Forexrates.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ForexSymbols

> []ForexSymbol ForexSymbols(ctx, exchange)

Forex Symbol

List supported forex symbols.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**exchange** | **string**| Exchange you want to get the list of symbols from. | 

### Return type

[**[]ForexSymbol**](ForexSymbol.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FundOwnership

> FundOwnership FundOwnership(ctx, symbol, optional)

Fund Ownership

Get a full list fund and institutional investors of a company in descending order of the number of shares held.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**symbol** | **string**| Symbol of the company: AAPL. | 
 **optional** | ***FundOwnershipOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a FundOwnershipOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **optional.Int64**| Limit number of results. Leave empty to get the full list. | 

### Return type

[**FundOwnership**](FundOwnership.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GeneralNews

> []News GeneralNews(ctx, category, optional)

General News

Get latest market news.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**category** | **string**| This parameter can be 1 of the following values &lt;code&gt;general, forex, crypto, merger&lt;/code&gt;. | 
 **optional** | ***GeneralNewsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GeneralNewsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **minId** | **optional.String**| Use this field to get only news after this ID. Default to 0 | 

### Return type

[**[]News**](News.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## InvestorsOwnership

> InvestorsOwnership InvestorsOwnership(ctx, symbol, optional)

Investors Ownership

Get a full list of shareholders/investors of a company in descending order of the number of shares held.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**symbol** | **string**| Symbol of the company: AAPL. | 
 **optional** | ***InvestorsOwnershipOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a InvestorsOwnershipOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **optional.Int64**| Limit number of results. Leave empty to get the full list. | 

### Return type

[**InvestorsOwnership**](InvestorsOwnership.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IpoCalendar

> []IpoEvent IpoCalendar(ctx, from, to)

IPO Calendar

Get recent and coming IPO.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**from** | **string**| From date: 2020-03-15. | 
**to** | **string**| To date: 2020-03-16. | 

### Return type

[**[]IpoEvent**](IPOEvent.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MajorDevelopments

> MajorDevelopments MajorDevelopments(ctx, symbol, optional)

Major Developments

List latest major developments of a company going back 20 years with 12M+ data points. This data can be used to highlight the most significant events.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**symbol** | **string**| Company symbol. | 
 **optional** | ***MajorDevelopmentsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a MajorDevelopmentsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **from** | **optional.String**| From time: 2020-01-01. | 
 **to** | **optional.String**| To time: 2020-01-05. | 

### Return type

[**MajorDevelopments**](MajorDevelopments.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## NewsSentiment

> NewsSentiment NewsSentiment(ctx, symbol)

News Sentiment

Get company's news sentiment and statistics. This endpoint is only available for US companies.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**symbol** | **string**| Company symbol. | 

### Return type

[**NewsSentiment**](NewsSentiment.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatternRecognition

> []map[string]interface{} PatternRecognition(ctx, symbol, resolution)

Pattern Recognition

Run pattern recognition algorithm on a symbol. Support double top/bottom, triple top/bottom, head and shoulders, triangle, wedge, channel, flag, and candlestick patterns.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**symbol** | **string**| Symbol | 
**resolution** | **string**| Supported resolution includes &lt;code&gt;1, 5, 15, 30, 60, D, W, M &lt;/code&gt;.Some timeframes might not be available depending on the exchange. | 

### Return type

**[]map[string]interface{}**

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PriceTarget

> PriceTarget PriceTarget(ctx, symbol)

Price Target

Get latest price target consensus.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**symbol** | **string**| Symbol of the company: AAPL. | 

### Return type

[**PriceTarget**](PriceTarget.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Quote

> Quote Quote(ctx, symbol)

Quote

<p>Get quote data for stocks. Constant polling is not recommended. Use websocket if you need real-time update.</p><p> Real-time stock prices for international markets are supported for Enterprise clients via our partner's feed. <a href=\"mailto:support@finnhub.io\">Contact Us</a> to learn more.</p>

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**symbol** | **string**| Symbol | 

### Return type

[**Quote**](Quote.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RecommendationTrends

> RecommendationTrends RecommendationTrends(ctx, symbol)

Recommendation Trends

Get latest analyst recommendation trends for a company.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**symbol** | **string**| Symbol of the company: AAPL. | 

### Return type

[**RecommendationTrends**](RecommendationTrends.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StockCandles

> StockCandles StockCandles(ctx, symbol, resolution, optional)

Stock Candles

<p>Get candlestick data for stocks going back 25 years.</p><p> Real-time stock prices for international markets are supported for Enterprise clients via our partner's feed. <a href=\"mailto:support@finnhub.io\">Contact Us</a> to learn more.</p>

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**symbol** | **string**| Symbol. | 
**resolution** | **string**| Supported resolution includes &lt;code&gt;1, 5, 15, 30, 60, D, W, M &lt;/code&gt;.Some timeframes might not be available depending on the exchange. | 
 **optional** | ***StockCandlesOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a StockCandlesOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **from** | **optional.Int64**| UNIX timestamp. Interval initial value. If count is not provided, this field is required | 
 **to** | **optional.Int64**| UNIX timestamp. Interval end value. If count is not provided, this field is required | 
 **format** | **optional.String**| By default, &lt;code&gt;format&#x3D;json&lt;/code&gt;. Strings &lt;code&gt;json&lt;/code&gt; and &lt;code&gt;csv&lt;/code&gt; are accepted. | 
 **adjusted** | **optional.String**| By default, &lt;code&gt;adjusted&#x3D;false&lt;/code&gt;. Use &lt;code&gt;true&lt;/code&gt; to get adjusted data. | 

### Return type

[**StockCandles**](StockCandles.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StockDividends

> Dividends StockDividends(ctx, symbol, from, to)

Dividends

Get dividends data for common stocks going back 30 years.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**symbol** | **string**| Symbol. | 
**from** | **string**| YYYY-MM-DD. | 
**to** | **string**| YYYY-MM-DD. | 

### Return type

[**Dividends**](Dividends.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StockSplits

> Splits StockSplits(ctx, symbol, from, to)

Splits

Get splits data for stocks.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**symbol** | **string**| Symbol. | 
**from** | **string**| YYYY-MM-DD. | 
**to** | **string**| YYYY-MM-DD. | 

### Return type

[**Splits**](Splits.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StockSymbols

> []Stock StockSymbols(ctx, exchange)

Stock Symbol

List supported stocks.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**exchange** | **string**| Exchange you want to get the list of symbols from. List of exchanges with fundamental data can be found &lt;a href&#x3D;\&quot;https://docs.google.com/spreadsheets/d/1I3pBxjfXB056-g_JYf_6o3Rns3BV2kMGG1nCatb91ls/edit?usp&#x3D;sharing\&quot; target&#x3D;\&quot;_blank\&quot;&gt;here&lt;/a&gt;. | 

### Return type

[**[]Stock**](Stock.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StockTick

> TickData StockTick(ctx, symbol, date)

Tick Data

<p>Get historical tick data for US stocks from all 13 exchanges. Return csv format. You can send the request directly to our tick server at <a href=\"https://tick.finnhub.io/\">https://tick.finnhub.io/</a> with the same path and parameters or get redirected there if you call our main server. Data is updated at the end of each trading day.</p><p>Tick data from 1985 is available for Enterprise clients via our partner's feed. <a href=\"mailto:support@finnhub.io\">Contact us</a> to learn more.</p>

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**symbol** | **string**| Symbol. | 
**date** | **string**| Date: 2020-04-02. | 

### Return type

[**TickData**](TickData.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SupportResistance

> []float32 SupportResistance(ctx, symbol, resolution)

Support/Resistance

Get support and resistance levels for a symbol.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**symbol** | **string**| Symbol | 
**resolution** | **string**| Supported resolution includes &lt;code&gt;1, 5, 15, 30, 60, D, W, M &lt;/code&gt;.Some timeframes might not be available depending on the exchange. | 

### Return type

**[]float32**

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TechnicalIndicator

> TechnicalIndicators TechnicalIndicator(ctx, symbol, resolution, from, to, indicator, optional)

Technical Indicators

Return technical indicator with price data. List of supported indicators can be found <a href=\"https://docs.google.com/spreadsheets/d/1ylUvKHVYN2E87WdwIza8ROaCpd48ggEl1k5i5SgA29k/edit?usp=sharing\" target=\"_blank\">here</a>.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**symbol** | **string**| symbol | 
**resolution** | **string**| Supported resolution includes &lt;code&gt;1, 5, 15, 30, 60, D, W, M &lt;/code&gt;.Some timeframes might not be available depending on the exchange. | 
**from** | **int64**| UNIX timestamp. Interval initial value. If count is not provided, this field is required | 
**to** | **int64**| UNIX timestamp. Interval end value. If count is not provided, this field is required | 
**indicator** | **string**| Indicator name. Full list can be found &lt;a href&#x3D;\&quot;https://docs.google.com/spreadsheets/d/1ylUvKHVYN2E87WdwIza8ROaCpd48ggEl1k5i5SgA29k/edit?usp&#x3D;sharing\&quot; target&#x3D;\&quot;_blank\&quot;&gt;here&lt;/a&gt;. | 
 **optional** | ***TechnicalIndicatorOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a TechnicalIndicatorOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------





 **indicatorSpecificFields** | [**optional.Interface of map[string]interface{}**](.md)| Check out &lt;a href&#x3D;\&quot;https://docs.google.com/spreadsheets/d/1ylUvKHVYN2E87WdwIza8ROaCpd48ggEl1k5i5SgA29k/edit?usp&#x3D;sharing\&quot; target&#x3D;\&quot;_blank\&quot;&gt;this page&lt;/a&gt; to see which indicators and params are supported. | 

### Return type

[**TechnicalIndicators**](TechnicalIndicators.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Transcripts

> EarningsCallTranscripts Transcripts(ctx, id)

Earnings Call Transcripts

<p>Get earnings call transcripts, audio and participants' list. This endpoint is only available for US companies. Earnings call transcripts for international markets are available for Enterprise clients via our partner's feed. <a href=\"mailto:support@finnhub.io\">Contact us</a> to learn more.</p><p>17+ years of data is available with 170,000+ audio which add up to 6TB in size.</p>

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| Transcript&#39;s id obtained with &lt;a href&#x3D;\&quot;#transcripts-list\&quot;&gt;Transcripts List endpoint&lt;/a&gt;. | 

### Return type

[**EarningsCallTranscripts**](EarningsCallTranscripts.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TranscriptsList

> EarningsCallTranscriptsList TranscriptsList(ctx, symbol)

Earnings Call Transcripts List

List earnings call transcripts' metadata. This endpoint is only available for US companies. Earnings call transcripts for international markets are available for Enterprise clients via our partner's feed. <a href=\"mailto:support@finnhub.io\">Contact us</a> to learn more.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**symbol** | **string**| Company symbol: AAPL. Leave empty to list the latest transcripts | 

### Return type

[**EarningsCallTranscriptsList**](EarningsCallTranscriptsList.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpgradeDowngrade

> []UpgradeDowngrade UpgradeDowngrade(ctx, optional)

Stock Upgrade/Downgrade

Get latest stock upgrade and downgrade.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***UpgradeDowngradeOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UpgradeDowngradeOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **symbol** | **optional.String**| Symbol of the company: AAPL. If left blank, the API will return latest stock upgrades/downgrades. | 
 **from** | **optional.String**| From date: 2000-03-15. | 
 **to** | **optional.String**| To date: 2020-03-16. | 

### Return type

[**[]UpgradeDowngrade**](UpgradeDowngrade.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

