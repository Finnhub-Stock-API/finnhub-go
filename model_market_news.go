/*
 * Finnhub API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 1.0.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package finnhub

import (
	"encoding/json"
)

// MarketNews struct for MarketNews
type MarketNews struct {
	// News category.
	Category *string `json:"category,omitempty"`
	// Published time in UNIX timestamp.
	Datetime *int64 `json:"datetime,omitempty"`
	// News headline.
	Headline *string `json:"headline,omitempty"`
	// News ID. This value can be used for <code>minId</code> params to get the latest news only.
	Id *int64 `json:"id,omitempty"`
	// Thumbnail image URL.
	Image *string `json:"image,omitempty"`
	// Related stocks and companies mentioned in the article.
	Related *string `json:"related,omitempty"`
	// News source.
	Source *string `json:"source,omitempty"`
	// News summary.
	Summary *string `json:"summary,omitempty"`
	// URL of the original article.
	Url *string `json:"url,omitempty"`
}

// NewMarketNews instantiates a new MarketNews object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMarketNews() *MarketNews {
	this := MarketNews{}
	return &this
}

// NewMarketNewsWithDefaults instantiates a new MarketNews object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMarketNewsWithDefaults() *MarketNews {
	this := MarketNews{}
	return &this
}

// GetCategory returns the Category field value if set, zero value otherwise.
func (o *MarketNews) GetCategory() string {
	if o == nil || o.Category == nil {
		var ret string
		return ret
	}
	return *o.Category
}

// GetCategoryOk returns a tuple with the Category field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MarketNews) GetCategoryOk() (*string, bool) {
	if o == nil || o.Category == nil {
		return nil, false
	}
	return o.Category, true
}

// HasCategory returns a boolean if a field has been set.
func (o *MarketNews) HasCategory() bool {
	if o != nil && o.Category != nil {
		return true
	}

	return false
}

// SetCategory gets a reference to the given string and assigns it to the Category field.
func (o *MarketNews) SetCategory(v string) {
	o.Category = &v
}

// GetDatetime returns the Datetime field value if set, zero value otherwise.
func (o *MarketNews) GetDatetime() int64 {
	if o == nil || o.Datetime == nil {
		var ret int64
		return ret
	}
	return *o.Datetime
}

// GetDatetimeOk returns a tuple with the Datetime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MarketNews) GetDatetimeOk() (*int64, bool) {
	if o == nil || o.Datetime == nil {
		return nil, false
	}
	return o.Datetime, true
}

// HasDatetime returns a boolean if a field has been set.
func (o *MarketNews) HasDatetime() bool {
	if o != nil && o.Datetime != nil {
		return true
	}

	return false
}

// SetDatetime gets a reference to the given int64 and assigns it to the Datetime field.
func (o *MarketNews) SetDatetime(v int64) {
	o.Datetime = &v
}

// GetHeadline returns the Headline field value if set, zero value otherwise.
func (o *MarketNews) GetHeadline() string {
	if o == nil || o.Headline == nil {
		var ret string
		return ret
	}
	return *o.Headline
}

// GetHeadlineOk returns a tuple with the Headline field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MarketNews) GetHeadlineOk() (*string, bool) {
	if o == nil || o.Headline == nil {
		return nil, false
	}
	return o.Headline, true
}

// HasHeadline returns a boolean if a field has been set.
func (o *MarketNews) HasHeadline() bool {
	if o != nil && o.Headline != nil {
		return true
	}

	return false
}

// SetHeadline gets a reference to the given string and assigns it to the Headline field.
func (o *MarketNews) SetHeadline(v string) {
	o.Headline = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *MarketNews) GetId() int64 {
	if o == nil || o.Id == nil {
		var ret int64
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MarketNews) GetIdOk() (*int64, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *MarketNews) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given int64 and assigns it to the Id field.
func (o *MarketNews) SetId(v int64) {
	o.Id = &v
}

// GetImage returns the Image field value if set, zero value otherwise.
func (o *MarketNews) GetImage() string {
	if o == nil || o.Image == nil {
		var ret string
		return ret
	}
	return *o.Image
}

// GetImageOk returns a tuple with the Image field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MarketNews) GetImageOk() (*string, bool) {
	if o == nil || o.Image == nil {
		return nil, false
	}
	return o.Image, true
}

// HasImage returns a boolean if a field has been set.
func (o *MarketNews) HasImage() bool {
	if o != nil && o.Image != nil {
		return true
	}

	return false
}

// SetImage gets a reference to the given string and assigns it to the Image field.
func (o *MarketNews) SetImage(v string) {
	o.Image = &v
}

// GetRelated returns the Related field value if set, zero value otherwise.
func (o *MarketNews) GetRelated() string {
	if o == nil || o.Related == nil {
		var ret string
		return ret
	}
	return *o.Related
}

// GetRelatedOk returns a tuple with the Related field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MarketNews) GetRelatedOk() (*string, bool) {
	if o == nil || o.Related == nil {
		return nil, false
	}
	return o.Related, true
}

// HasRelated returns a boolean if a field has been set.
func (o *MarketNews) HasRelated() bool {
	if o != nil && o.Related != nil {
		return true
	}

	return false
}

// SetRelated gets a reference to the given string and assigns it to the Related field.
func (o *MarketNews) SetRelated(v string) {
	o.Related = &v
}

// GetSource returns the Source field value if set, zero value otherwise.
func (o *MarketNews) GetSource() string {
	if o == nil || o.Source == nil {
		var ret string
		return ret
	}
	return *o.Source
}

// GetSourceOk returns a tuple with the Source field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MarketNews) GetSourceOk() (*string, bool) {
	if o == nil || o.Source == nil {
		return nil, false
	}
	return o.Source, true
}

// HasSource returns a boolean if a field has been set.
func (o *MarketNews) HasSource() bool {
	if o != nil && o.Source != nil {
		return true
	}

	return false
}

// SetSource gets a reference to the given string and assigns it to the Source field.
func (o *MarketNews) SetSource(v string) {
	o.Source = &v
}

// GetSummary returns the Summary field value if set, zero value otherwise.
func (o *MarketNews) GetSummary() string {
	if o == nil || o.Summary == nil {
		var ret string
		return ret
	}
	return *o.Summary
}

// GetSummaryOk returns a tuple with the Summary field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MarketNews) GetSummaryOk() (*string, bool) {
	if o == nil || o.Summary == nil {
		return nil, false
	}
	return o.Summary, true
}

// HasSummary returns a boolean if a field has been set.
func (o *MarketNews) HasSummary() bool {
	if o != nil && o.Summary != nil {
		return true
	}

	return false
}

// SetSummary gets a reference to the given string and assigns it to the Summary field.
func (o *MarketNews) SetSummary(v string) {
	o.Summary = &v
}

// GetUrl returns the Url field value if set, zero value otherwise.
func (o *MarketNews) GetUrl() string {
	if o == nil || o.Url == nil {
		var ret string
		return ret
	}
	return *o.Url
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MarketNews) GetUrlOk() (*string, bool) {
	if o == nil || o.Url == nil {
		return nil, false
	}
	return o.Url, true
}

// HasUrl returns a boolean if a field has been set.
func (o *MarketNews) HasUrl() bool {
	if o != nil && o.Url != nil {
		return true
	}

	return false
}

// SetUrl gets a reference to the given string and assigns it to the Url field.
func (o *MarketNews) SetUrl(v string) {
	o.Url = &v
}

func (o MarketNews) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Category != nil {
		toSerialize["category"] = o.Category
	}
	if o.Datetime != nil {
		toSerialize["datetime"] = o.Datetime
	}
	if o.Headline != nil {
		toSerialize["headline"] = o.Headline
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Image != nil {
		toSerialize["image"] = o.Image
	}
	if o.Related != nil {
		toSerialize["related"] = o.Related
	}
	if o.Source != nil {
		toSerialize["source"] = o.Source
	}
	if o.Summary != nil {
		toSerialize["summary"] = o.Summary
	}
	if o.Url != nil {
		toSerialize["url"] = o.Url
	}
	return json.Marshal(toSerialize)
}

type NullableMarketNews struct {
	value *MarketNews
	isSet bool
}

func (v NullableMarketNews) Get() *MarketNews {
	return v.value
}

func (v *NullableMarketNews) Set(val *MarketNews) {
	v.value = val
	v.isSet = true
}

func (v NullableMarketNews) IsSet() bool {
	return v.isSet
}

func (v *NullableMarketNews) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMarketNews(val *MarketNews) *NullableMarketNews {
	return &NullableMarketNews{value: val, isSet: true}
}

func (v NullableMarketNews) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMarketNews) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

