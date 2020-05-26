/*
 * Finnhub API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package finnhub
import (
	"time"
)
// Filing struct for Filing
type Filing struct {
	// Access number.
	AccessNumber string `json:"accessNumber,omitempty"`
	// Symbol.
	Symbol string `json:"symbol,omitempty"`
	// CIK.
	Cik string `json:"cik,omitempty"`
	// Form type.
	Form string `json:"form,omitempty"`
	// Filed date <code>%Y-%m-%d %H:%M:%S</code>.
	FiledDate time.Time `json:"filedDate,omitempty"`
	// Accepted date <code>%Y-%m-%d %H:%M:%S</code>.
	AcceptedDate time.Time `json:"acceptedDate,omitempty"`
	// Report's URL.
	ReportUrl string `json:"reportUrl,omitempty"`
	// Filing's URL.
	FilingUrl string `json:"filingUrl,omitempty"`
}