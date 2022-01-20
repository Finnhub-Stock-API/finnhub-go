# Go API client for finnhub.io

## Overview
- API documentation: https://finnhub.io/docs/api
- API version: 1.0.0
- Package version: 2.0.9

## Installation

### Using Go Modules
Make sure your project is using Go Modules (it will have a `go.mod` file in its
root if it already is):

``` sh
go mod init
```

Then, reference finnhub-go in a Go program with `import`:

``` go
import (
    finnhub "github.com/Finnhub-Stock-API/finnhub-go/v2"
)
```

Run any of the normal `go` commands (`build`/`install`/`test`). The Go
toolchain will resolve and fetch the finnhub-go module automatically.

Alternatively, you can also explicitly `go get` the package into a project:

```shell
$ go get -u github.com/Finnhub-Stock-API/finnhub-go/v2
```

### Using `go get`
If you don't want to use Go Modules, you can choose to get the library directly:
```shell
$ go get -u github.com/Finnhub-Stock-API/finnhub-go
```
Then, reference finnhub-go in a Go program with `import` (Note that no /v2 at the end):
``` go
import (
    finnhub "github.com/Finnhub-Stock-API/finnhub-go"
)
```


## Examples
Example (check out other methods documentation [here](https://pkg.go.dev/github.com/Finnhub-Stock-API/finnhub-go/v2)):

```golang
package main

import (
	"context"
	"fmt"
	finnhub "github.com/Finnhub-Stock-API/finnhub-go/v2"
)

func main() {
    cfg := finnhub.NewConfiguration()
    cfg.AddDefaultHeader("X-Finnhub-Token", "<API_key>")
    finnhubClient := finnhub.NewAPIClient(cfg).DefaultApi
	
    //Earnings calendar
    earningsCalendar, _, err := finnhubClient.EarningsCalendar(context.Background()).From("2021-07-01").To("2021-07-25").Execute()
    fmt.Printf("%+v\n", earningsCalendar)

    // NBBO
    bboData, _, err := finnhubClient.StockNbbo(context.Background()).Symbol("AAPL").Date("2021-07-23").Limit(50).Skip(0).Execute()
    fmt.Printf("%+v\n", bboData)

    // Bid ask
    lastBidAsk, _, err := finnhubClient.StockBidask(context.Background()).Symbol("AAPL").Execute()
    fmt.Printf("%+v\n", lastBidAsk)

    // Stock dividends 2
    dividends2, _, err := finnhubClient.StockBasicDividends(context.Background()).Symbol("KO").Execute()
    fmt.Printf("%+v\n", dividends2)

    //Stock candles
    stockCandles, _, err := finnhubClient.StockCandles(context.Background()).Symbol("AAPL").Resolution("D").From(1590988249).To(1591852249).Execute()
    fmt.Printf("%+v\n", stockCandles)

    // Example with required parameters
    news, _, err := finnhubClient.CompanyNews(context.Background()).Symbol("AAPL").From("2020-05-01").To("2020-05-01").Execute()
    if err != nil {
        fmt.Println(err)
    }
    fmt.Printf("%+v\n", news)

    // Example with required and optional parameters
    ownerships, _, err := finnhubClient.Ownership(context.Background()).Symbol("AAPL").Execute()
    fmt.Printf("%+v\n", ownerships)

    // Aggregate Indicator
    aggregateIndicator, _, err := finnhubClient.AggregateIndicator(context.Background()).Symbol("AAPL").Resolution("D").Execute()
    fmt.Printf("%+v\n", aggregateIndicator)

    // Basic financials
    basicFinancials, _, err := finnhubClient.CompanyBasicFinancials(context.Background()).Symbol("MSFT").Metric("all").Execute()
    fmt.Printf("%+v\n", basicFinancials)

    // Company earnings
    earningsSurprises, _, err := finnhubClient.CompanyEarnings(context.Background()).Symbol("AAPL").Execute()
    fmt.Printf("%+v\n", earningsSurprises)

    // Company EPS estimates
    epsEstimate, _, err := finnhubClient.CompanyEpsEstimates(context.Background()).Symbol("AAPL").Execute()
    fmt.Printf("%+v\n", epsEstimate)

    // Company executive
    executive, _, err := finnhubClient.CompanyExecutive(context.Background()).Symbol("AAPL").Execute()
    fmt.Printf("%+v\n", executive)

    // Company peers
    peers, _, err := finnhubClient.CompanyPeers(context.Background()).Symbol("AAPL").Execute()
    fmt.Printf("%+v\n", peers)

    // Company profile
    profile, _, err := finnhubClient.CompanyProfile(context.Background()).Symbol("AAPL").Execute()
    fmt.Printf("%+v\n", profile)

    profileISIN, _, err := finnhubClient.CompanyProfile(context.Background()).Isin("US0378331005").Execute()
    fmt.Printf("%+v\n", profileISIN)

    profileCusip, _, err := finnhubClient.CompanyProfile(context.Background()).Cusip("037833100").Execute()
    fmt.Printf("%+v\n", profileCusip)

    // Company profile2
    profile2, _, err := finnhubClient.CompanyProfile2(context.Background()).Symbol("AAPL").Execute()
    fmt.Printf("%+v\n", profile2)

    // Revenue Estimates
    revenueEstimates, _, err := finnhubClient.CompanyRevenueEstimates(context.Background()).Symbol("AAPL").Execute()
    fmt.Printf("%+v\n", revenueEstimates)

    // List country
    countries, _, err := finnhubClient.Country(context.Background()).Execute()
    fmt.Printf("%+v\n", countries)

    // Covid-19
    covid19, _, err := finnhubClient.Covid19(context.Background()).Execute()
    fmt.Printf("%+v\n", covid19)

    // FDA Calendar
    fdaCalendar, _, err := finnhubClient.FdaCommitteeMeetingCalendar(context.Background()).Execute()
    fmt.Printf("%+v\n", fdaCalendar)

    // Crypto candles
    cryptoCandles, _, err := finnhubClient.CryptoCandles(context.Background()).Symbol("BINANCE:BTCUSDT").Resolution("D").From(1590988249).To(1591852249).Execute()
    fmt.Printf("%+v\n", cryptoCandles)

    // Crypto exchanges
    cryptoExchange, _, err := finnhubClient.CryptoExchanges(context.Background()).Execute()
    fmt.Printf("%+v\n", cryptoExchange)

    // Crypto symbols
    cryptoSymbol, _, err := finnhubClient.CryptoSymbols(context.Background()).Exchange("BINANCE").Execute()
    fmt.Printf("%+v\n", cryptoSymbol[0:5])

    // Economic Calendar
    economicCalendar, _, err := finnhubClient.EconomicCalendar(context.Background()).Execute()
    fmt.Printf("%+v\n", economicCalendar)

    // Economic code
    economicCode, _, err := finnhubClient.EconomicCode(context.Background()).Execute()
    fmt.Printf("%+v\n", economicCode)

    // Economic data
    economicData, _, err := finnhubClient.EconomicData(context.Background()).Code("MA-USA-656880").Execute()
    fmt.Printf("%+v\n", economicData)

    // Filings
    filings, _, err := finnhubClient.Filings(context.Background()).Symbol("AAPL").Execute()
    fmt.Printf("%+v\n", filings)

    // International filings
    internationalFilings, _, err := finnhubClient.InternationalFilings(context.Background()).Symbol("RY.TO").Execute()
    fmt.Printf("%+v\n", internationalFilings)

    // Filings Sentiment
    filingsSentiment, _, err := finnhubClient.FilingsSentiment(context.Background()).AccessNumber("0000320193-20-000052").Execute()
    fmt.Printf("%+v\n", filingsSentiment)

    // Similarity Index
    similarityIndex, _, err := finnhubClient.SimilarityIndex(context.Background()).Symbol("AAPL").Execute()
    fmt.Printf("%+v\n", similarityIndex)

    // Financials
    financials, _, err := finnhubClient.Financials(context.Background()).Symbol("AAPL").Statement("bs").Freq("annual").Execute()
    fmt.Printf("%+v\n", financials)

    // Financials Reported
    financialsReported, _, err := finnhubClient.FinancialsReported(context.Background()).Symbol("AAPL").Execute()
    fmt.Printf("%+v\n", financialsReported)

    // Forex candles
    forexCandles, _, err := finnhubClient.ForexCandles(context.Background()).Symbol("OANDA:EUR_USD").Resolution("D").From(1590988249).To(1591852249).Execute()
    fmt.Printf("%+v\n", forexCandles)

    // Forex exchanges
    forexExchanges, _, err := finnhubClient.ForexExchanges(context.Background()).Execute()
    fmt.Printf("%+v\n", forexExchanges)

    // Forex rates
    forexRates, _, err := finnhubClient.ForexRates(context.Background()).Base("USD").Execute()
    fmt.Printf("%+v\n", forexRates)

    // Forex symbols
    forexSymbols, _, err := finnhubClient.ForexSymbols(context.Background()).Exchange("OANDA").Execute()
    fmt.Printf("%+v\n", forexSymbols)

    // Fund ownership
    fundOwnership, _, err := finnhubClient.FundOwnership(context.Background()).Symbol("AAPL").Execute()
    fmt.Printf("%+v\n", fundOwnership)

    // General news
    generalNews, _, err := finnhubClient.MarketNews(context.Background()).Category("general").Execute()
    fmt.Printf("%+v\n", generalNews)

    // Ipo calendar
    ipoCalendar, _, err := finnhubClient.IpoCalendar(context.Background()).From("2021-01-01").To("2021-06-30").Execute()
    fmt.Printf("%+v\n", ipoCalendar)

    // Press Releases
    majorDevelopment, _, err := finnhubClient.PressReleases(context.Background()).Symbol("AAPL").Execute()
    fmt.Printf("%+v\n", majorDevelopment)

    // News sentiment
    newsSentiment, _, err := finnhubClient.NewsSentiment(context.Background()).Symbol("AAPL").Execute()
    fmt.Printf("%+v\n", newsSentiment)

    // Pattern recognition
    patterns, _, err := finnhubClient.PatternRecognition(context.Background()).Symbol("AAPL").Resolution("D").Execute()
    fmt.Printf("%+v\n", patterns)

    // Price target
    priceTarget, _, err := finnhubClient.PriceTarget(context.Background()).Symbol("AAPL").Execute()
    fmt.Printf("%+v\n", priceTarget)

    // Quote
    quote, _, err := finnhubClient.Quote(context.Background()).Symbol("AAPL").Execute()
    fmt.Printf("%+v\n", quote)

    // Recommendation trends
    recommendationTrend, _, err := finnhubClient.RecommendationTrends(context.Background()).Symbol("AAPL").Execute()
    fmt.Printf("%+v\n", recommendationTrend)

    // Stock dividends
    dividends, _, err := finnhubClient.StockDividends(context.Background()).Symbol("KO").From("2019-01-01").To("2021-01-01").Execute()
    fmt.Printf("%+v\n", dividends)

    // Splits
    splits, _, err := finnhubClient.StockSplits(context.Background()).Symbol("AAPL").From("2000-01-01").To("2020-06-15").Execute()
    fmt.Printf("%+v\n", splits)

    // Stock symbols
    stockSymbols, _, err := finnhubClient.StockSymbols(context.Background()).Exchange("US").Execute()
    fmt.Printf("%+v\n", stockSymbols[0:5])

    // Support resistance
    supportResitance, _, err := finnhubClient.SupportResistance(context.Background()).Symbol("AAPL").Resolution("D").Execute()
    fmt.Printf("%+v\n", supportResitance)

    // Technical indicator
    technicalIndicator, _, err := finnhubClient.TechnicalIndicator(context.Background()).Symbol("AAPL").Resolution("D").From(1583098857).To(1584308457).Indicator("sma").IndicatorFields(map[string]interface{}{"timeperiod": 3}).Execute()
    fmt.Printf("%+v\n", technicalIndicator)

    // Transcripts
    transcripts, _, err := finnhubClient.Transcripts(context.Background()).Id("AAPL_162777").Execute()
    fmt.Printf("%+v\n", transcripts)

    // Transcripts list
    transcriptsList, _, err := finnhubClient.TranscriptsList(context.Background()).Symbol("AAPL").Execute()
    fmt.Printf("%+v\n", transcriptsList)

    // Upgrade/downgrade
    upgradeDowngrade, _, err := finnhubClient.UpgradeDowngrade(context.Background()).Symbol("BYND").Execute()
    fmt.Printf("%+v\n", upgradeDowngrade)

    // Tick Data
    tickData, _, err := finnhubClient.StockTick(context.Background()).Symbol("AAPL").Date("2021-07-23").Limit(50).Skip(0).Execute()
    fmt.Printf("%+v\n", tickData)

    // Indices Constituents
    indicesConstData, _, err := finnhubClient.IndicesConstituents(context.Background()).Symbol("^GSPC").Execute()
    fmt.Printf("%+v\n", indicesConstData)

    // Indices Historical Constituents
    indicesHistoricalConstData, _, err := finnhubClient.IndicesHistoricalConstituents(context.Background()).Symbol("^GSPC").Execute()
    fmt.Printf("%+v\n", indicesHistoricalConstData)

    // ETFs Profile
    etfsProfileData, _, err := finnhubClient.EtfsProfile(context.Background()).Symbol("SPY").Execute()
    fmt.Printf("%+v\n", etfsProfileData)

    // ETFs Holdings
    etfsHoldingsData, _, err := finnhubClient.EtfsHoldings(context.Background()).Symbol("SPY").Execute()
    fmt.Printf("%+v\n", etfsHoldingsData)

    // ETFs Industry Exposure
    etfsIndustryExposureData, _, err := finnhubClient.EtfsSectorExposure(context.Background()).Symbol("SPY").Execute()
    fmt.Printf("%+v\n", etfsIndustryExposureData)

    // ETFs Country Exposure
    etfsCountryExposureData, _, err := finnhubClient.EtfsCountryExposure(context.Background()).Symbol("SPY").Execute()
    fmt.Printf("%+v\n", etfsCountryExposureData)

    // Mutual Funds Profile
    mfProfileData, _, err := finnhubClient.MutualFundProfile(context.Background()).Symbol("VTSAX").Execute()
    fmt.Printf("%+v\n", mfProfileData)

    // Mutual Funds Holdings
    mfHoldingsData, _, err := finnhubClient.MutualFundHoldings(context.Background()).Symbol("VTSAX").Execute()
    fmt.Printf("%+v\n", mfHoldingsData)

    // Mutual Funds Industry Exposure
    mfIndustryExposureData, _, err := finnhubClient.MutualFundSectorExposure(context.Background()).Symbol("VTSAX").Execute()
    fmt.Printf("%+v\n", mfIndustryExposureData)

    // Mutual Funds Country Exposure
    mfCountryExposureData, _, err := finnhubClient.MutualFundCountryExposure(context.Background()).Symbol("VTSAX").Execute()
    fmt.Printf("%+v\n", mfCountryExposureData)

    // Insider Transactions
    insiderTransactions, _, err := finnhubClient.InsiderTransactions(context.Background()).Symbol("AAPL").From("2021-01-01").To("2021-07-30").Execute()
    fmt.Printf("%+v\n", insiderTransactions)

    // Revenue breakdown
    revenueBreakdown, _, err := finnhubClient.RevenueBreakdown(context.Background()).Symbol("AAPL").Execute()
    fmt.Printf("%+v\n", revenueBreakdown)

    // Social Sentiment
    socialSentiment, _, err := finnhubClient.SocialSentiment(context.Background()).Symbol("GME").Execute()
    fmt.Printf("%+v\n", socialSentiment)

    // Investment theme
    investmentTheme, _, err := finnhubClient.InvestmentThemes(context.Background()).Theme("financialExchangesData").Execute()
    fmt.Printf("%+v\n", investmentTheme)

    // Supply chain
    supplyChain, _, err := finnhubClient.SupplyChainRelationships(context.Background()).Symbol("AAPL").Execute()
    fmt.Printf("%+v\n", supplyChain)

    //Symbol lookup
    searchResult, _, err := finnhubClient.SymbolSearch(context.Background()).Q("AAPL").Execute()
    fmt.Printf("%+v\n", searchResult)
    
    // Company ESG
    companyESGScore, _, err := finnhubClient.CompanyEsgScore(context.Background()).Symbol("AAPL").Execute()
	fmt.Printf("%+v\n", companyESGScore)
    
    // Company Earnings Quality Score
    earningsQualityScore, _, err := finnhubClient.CompanyEarningsQualityScore(context.Background()).Symbol("AAPL").Freq("quarterly").Execute()
    if err != nil {
		panic(err)
	}
    fmt.Printf("%+v\n", earningsQualityScore)   
    
    // Crypto Profile
    cryptoProfile, _, err := finnhubClient.CryptoProfile(context.Background()).Symbol("BTC").Execute()
	if err != nil {
		panic(err)
	}
	fmt.Println(objectString(cryptoProfile))

    // EBITDA Estimates
    ebitdaEstimates, _, err := finnhubClient.CompanyEbitdaEstimates(context.Background()).Symbol("AAPL").Freq("annual").Execute()
    if err != nil {
        panic(err)
    }
    fmt.Printf("%+v\n", ebitdaEstimates)
    
    // EBIT Estimates
    ebitEstimates, _, err := finnhubClient.CompanyEbitEstimates(context.Background()).Symbol("AAPL").Freq("annual").Execute()
    if err != nil {
        panic(err)
    }
    fmt.Printf("%+v\n", ebitEstimates)

    // USPTO Patent
    uspto, _, err := finnhubClient.StockUsptoPatent(context.Background()).Symbol("NVDA").From("2021-01-01").To("2021-12-31").Execute()
    if err != nil {
        panic(err)
    }
    fmt.Printf("%+v\n", uspto)
}

```

## License

Apache License
