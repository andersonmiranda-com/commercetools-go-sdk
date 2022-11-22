package platform

// Generated file, please do not change!!!

import (
	"fmt"
)

type ByProjectKeyProductDiscountsMatchingRequestBuilder struct {
	projectKey string
	client     *Client
}

/**
*	This endpoint can be used to simulate which Product Discounts would be applied if a specified Product Variant had a specified Price.
*	Given Product and Product Variant IDs and a Price, this endpoint will return the [ProductDiscount](ctp:api:type:ProductDiscount) that would have been applied to that Price.
*
*	If a Product Discount could not be found that could be applied to the Price of a Product Variant, a [NoMatchingProductDiscountFound](ctp:api:type:NoMatchingProductDiscountFoundError) error is returned.
*
 */
func (rb *ByProjectKeyProductDiscountsMatchingRequestBuilder) Post(body ProductDiscountMatchQuery) *ByProjectKeyProductDiscountsMatchingRequestMethodPost {
	return &ByProjectKeyProductDiscountsMatchingRequestMethodPost{
		body:   body,
		url:    fmt.Sprintf("/%s/product-discounts/matching", rb.projectKey),
		client: rb.client,
	}
}
