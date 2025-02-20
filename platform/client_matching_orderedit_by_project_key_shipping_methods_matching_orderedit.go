package platform

// Generated file, please do not change!!!

import (
	"fmt"
)

type ByProjectKeyShippingMethodsMatchingOrdereditRequestBuilder struct {
	projectKey string
	client     *Client
}

/**
*	Retrieves all the ShippingMethods that can ship to the given [Location](/../api/projects/zones#location) for an [OrderEdit](/../api/projects/order-edits).
*
*	If the OrderEdit preview cannot be generated, an [EditPreviewFailed](ctp:api:type:EditPreviewFailedError) error is returned.
*
 */
func (rb *ByProjectKeyShippingMethodsMatchingOrdereditRequestBuilder) Get() *ByProjectKeyShippingMethodsMatchingOrdereditRequestMethodGet {
	return &ByProjectKeyShippingMethodsMatchingOrdereditRequestMethodGet{
		url:    fmt.Sprintf("/%s/shipping-methods/matching-orderedit", rb.projectKey),
		client: rb.client,
	}
}
