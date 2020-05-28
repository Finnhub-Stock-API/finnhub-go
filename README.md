# Go API client for finnhub.io

## Overview
- API documentation: https://finnhub.io/docs/api
- API version: 1.0
- Package version: 0.0.3

## Installation

Install package:

```shell
$ go get -u github.com/Finnhub-Stock-API/finnhub-go
```

Example:

```golang
package main

import (
	"context"
	"fmt"

	"github.com/antihax/optional"

	finnhub "github.com/Finnhub-Stock-API/finnhub-go"
)

func main() {
	client := finnhub.NewAPIClient(finnhub.NewConfiguration()).DefaultApi
	auth := context.WithValue(context.Background(), finnhub.ContextAPIKey, finnhub.APIKey{
		Key: "bqlsn1frh5rfdbi8u17g",
	})

	// Example with required parameters
	news, _, err := client.CompanyNews(auth, "AAPL", "2020-01-01", "2020-05-01")
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v", news)

	// Example with required and optional parameters
	investorsOwnershipOpts := &finnhub.InvestorsOwnershipOpts{Limit: optional.NewInt64(10)}
	ownerships, _, err := client.InvestorsOwnership(auth, "AAPL", investorsOwnershipOpts)
	if err != nil {
		panic(err)
	}
	fmt.Printf("\n\n%+v", ownerships)

	// Example with mutiple optional params
	stockCandlesOpts := &finnhub.StockCandlesOpts{
		From: optional.NewInt64(1572651390),
		To:   optional.NewInt64(1575243390),
	}
	stockCandles, _, err := client.StockCandles(auth, "AAPL", "D", stockCandlesOpts)
	if err != nil {
		panic(err)
	}
	fmt.Printf("\n\n%+v", stockCandles)
}
```


## API Endpoints

All URIs are relative to *https://finnhub.io/api/v1*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*DefaultApi* | [**AggregateIndicator**](docs/DefaultApi.md#aggregateindicator) | **Get** /scan/technical-indicator | Aggregate Indicators
*DefaultApi* | [**CompanyEarnings**](docs/DefaultApi.md#companyearnings) | **Get** /stock/earnings | Earnings Surprises
*DefaultApi* | [**CompanyEpsEstimates**](docs/DefaultApi.md#companyepsestimates) | **Get** /stock/eps-estimate | Earnings Estimates
*DefaultApi* | [**CompanyExecutive**](docs/DefaultApi.md#companyexecutive) | **Get** /stock/executive | Company Executive
*DefaultApi* | [**CompanyMetrics**](docs/DefaultApi.md#companymetrics) | **Get** /stock/metric | Metrics
*DefaultApi* | [**CompanyNews**](docs/DefaultApi.md#companynews) | **Get** /company-news | Company News
*DefaultApi* | [**CompanyPeers**](docs/DefaultApi.md#companypeers) | **Get** /stock/peers | Peers
*DefaultApi* | [**CompanyProfile**](docs/DefaultApi.md#companyprofile) | **Get** /stock/profile | Company Profile
*DefaultApi* | [**CompanyProfile2**](docs/DefaultApi.md#companyprofile2) | **Get** /stock/profile2 | Company Profile 2
*DefaultApi* | [**CompanyRevenueEstimates**](docs/DefaultApi.md#companyrevenueestimates) | **Get** /stock/revenue-estimate | Revenue Estimates
*DefaultApi* | [**Covid19**](docs/DefaultApi.md#covid19) | **Get** /covid19/us | COVID-19
*DefaultApi* | [**CryptoCandles**](docs/DefaultApi.md#cryptocandles) | **Get** /crypto/candle | Crypto Candles
*DefaultApi* | [**CryptoExchanges**](docs/DefaultApi.md#cryptoexchanges) | **Get** /crypto/exchange | Crypto Exchanges
*DefaultApi* | [**CryptoSymbols**](docs/DefaultApi.md#cryptosymbols) | **Get** /crypto/symbol | Crypto Symbol
*DefaultApi* | [**EarningsCalendar**](docs/DefaultApi.md#earningscalendar) | **Get** /calendar/earnings | Earnings Calendar
*DefaultApi* | [**Filings**](docs/DefaultApi.md#filings) | **Get** /stock/filings | Filings
*DefaultApi* | [**Financials**](docs/DefaultApi.md#financials) | **Get** /stock/financials | Financial Statements
*DefaultApi* | [**FinancialsReported**](docs/DefaultApi.md#financialsreported) | **Get** /stock/financials-reported | Financials As Reported
*DefaultApi* | [**ForexCandles**](docs/DefaultApi.md#forexcandles) | **Get** /forex/candle | Forex Candles
*DefaultApi* | [**ForexExchanges**](docs/DefaultApi.md#forexexchanges) | **Get** /forex/exchange | Forex Exchanges
*DefaultApi* | [**ForexRates**](docs/DefaultApi.md#forexrates) | **Get** /forex/rates | Forex rates
*DefaultApi* | [**ForexSymbols**](docs/DefaultApi.md#forexsymbols) | **Get** /forex/symbol | Forex Symbol
*DefaultApi* | [**FundOwnership**](docs/DefaultApi.md#fundownership) | **Get** /stock/fund-ownership | Fund Ownership
*DefaultApi* | [**GeneralNews**](docs/DefaultApi.md#generalnews) | **Get** /news | General News
*DefaultApi* | [**InvestorsOwnership**](docs/DefaultApi.md#investorsownership) | **Get** /stock/investor-ownership | Investors Ownership
*DefaultApi* | [**IpoCalendar**](docs/DefaultApi.md#ipocalendar) | **Get** /calendar/ipo | IPO Calendar
*DefaultApi* | [**MajorDevelopments**](docs/DefaultApi.md#majordevelopments) | **Get** /major-development | Major Developments
*DefaultApi* | [**NewsSentiment**](docs/DefaultApi.md#newssentiment) | **Get** /news-sentiment | News Sentiment
*DefaultApi* | [**PatternRecognition**](docs/DefaultApi.md#patternrecognition) | **Get** /scan/pattern | Pattern Recognition
*DefaultApi* | [**PriceTarget**](docs/DefaultApi.md#pricetarget) | **Get** /stock/price-target | Price Target
*DefaultApi* | [**Quote**](docs/DefaultApi.md#quote) | **Get** /quote | Quote
*DefaultApi* | [**RecommendationTrends**](docs/DefaultApi.md#recommendationtrends) | **Get** /stock/recommendation | Recommendation Trends
*DefaultApi* | [**StockCandles**](docs/DefaultApi.md#stockcandles) | **Get** /stock/candle | Stock Candles
*DefaultApi* | [**StockDividends**](docs/DefaultApi.md#stockdividends) | **Get** /stock/dividend | Dividends
*DefaultApi* | [**StockSplits**](docs/DefaultApi.md#stocksplits) | **Get** /stock/split | Splits
*DefaultApi* | [**StockSymbols**](docs/DefaultApi.md#stocksymbols) | **Get** /stock/symbol | Stock Symbol
*DefaultApi* | [**StockTick**](docs/DefaultApi.md#stocktick) | **Get** /stock/tick | Tick Data
*DefaultApi* | [**SupportResistance**](docs/DefaultApi.md#supportresistance) | **Get** /scan/support-resistance | Support/Resistance
*DefaultApi* | [**TechnicalIndicator**](docs/DefaultApi.md#technicalindicator) | **Get** /indicator | Technical Indicators
*DefaultApi* | [**Transcripts**](docs/DefaultApi.md#transcripts) | **Get** /stock/transcripts | Earnings Call Transcripts
*DefaultApi* | [**TranscriptsList**](docs/DefaultApi.md#transcriptslist) | **Get** /stock/transcripts/list | Earnings Call Transcripts List
*DefaultApi* | [**UpgradeDowngrade**](docs/DefaultApi.md#upgradedowngrade) | **Get** /stock/upgrade-downgrade | Stock Upgrade/Downgrade


## Models

 - [AggregateIndicators](docs/AggregateIndicators.md)
 - [Company](docs/Company.md)
 - [CompanyExecutive](docs/CompanyExecutive.md)
 - [CompanyNewsStatistics](docs/CompanyNewsStatistics.md)
 - [CompanyProfile](docs/CompanyProfile.md)
 - [CompanyProfile2](docs/CompanyProfile2.md)
 - [Covid19](docs/Covid19.md)
 - [CryptoCandles](docs/CryptoCandles.md)
 - [CryptoSymbol](docs/CryptoSymbol.md)
 - [Development](docs/Development.md)
 - [Dividends](docs/Dividends.md)
 - [EarningEstimate](docs/EarningEstimate.md)
 - [EarningRelease](docs/EarningRelease.md)
 - [EarningResult](docs/EarningResult.md)
 - [EarningsCallTranscripts](docs/EarningsCallTranscripts.md)
 - [EarningsCallTranscriptsList](docs/EarningsCallTranscriptsList.md)
 - [EarningsEstimates](docs/EarningsEstimates.md)
 - [Estimate](docs/Estimate.md)
 - [Filing](docs/Filing.md)
 - [FinancialStatements](docs/FinancialStatements.md)
 - [FinancialsAsReported](docs/FinancialsAsReported.md)
 - [ForexCandles](docs/ForexCandles.md)
 - [ForexSymbol](docs/ForexSymbol.md)
 - [Forexrates](docs/Forexrates.md)
 - [FundOwnership](docs/FundOwnership.md)
 - [Indicator](docs/Indicator.md)
 - [Investor](docs/Investor.md)
 - [InvestorsOwnership](docs/InvestorsOwnership.md)
 - [IpoEvent](docs/IpoEvent.md)
 - [MajorDevelopments](docs/MajorDevelopments.md)
 - [Metrics](docs/Metrics.md)
 - [News](docs/News.md)
 - [NewsSentiment](docs/NewsSentiment.md)
 - [PriceTarget](docs/PriceTarget.md)
 - [Quote](docs/Quote.md)
 - [RecommendationTrends](docs/RecommendationTrends.md)
 - [Report](docs/Report.md)
 - [RevenueEstimates](docs/RevenueEstimates.md)
 - [Sentiment](docs/Sentiment.md)
 - [Splits](docs/Splits.md)
 - [Stock](docs/Stock.md)
 - [StockCandles](docs/StockCandles.md)
 - [StockTranscripts](docs/StockTranscripts.md)
 - [TechnicalAnalysis](docs/TechnicalAnalysis.md)
 - [TechnicalIndicators](docs/TechnicalIndicators.md)
 - [TickData](docs/TickData.md)
 - [TranscriptContent](docs/TranscriptContent.md)
 - [TranscriptParticipant](docs/TranscriptParticipant.md)
 - [Trend](docs/Trend.md)
 - [UpgradeDowngrade](docs/UpgradeDowngrade.md)


## License

MIT License

Copyright (c) 2020 Finnhub

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
