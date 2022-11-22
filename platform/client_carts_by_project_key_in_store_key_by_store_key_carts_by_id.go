package platform

// Generated file, please do not change!!!

import (
	"fmt"
)

type ByProjectKeyInStoreKeyByStoreKeyCartsByIDRequestBuilder struct {
	projectKey string
	storeKey   string
	id         string
	client     *Client
}

/**
*	Returns a cart by its ID from a specific Store.
*	If the cart exists in the project but does not have the store field,
*	or the store field references a different store, this method returns a ResourceNotFound error.
*	The cart may not contain up-to-date prices, discounts etc.
*	If you want to ensure they're up-to-date, send an Update request with the Recalculate update action instead.
*
 */
func (rb *ByProjectKeyInStoreKeyByStoreKeyCartsByIDRequestBuilder) Get() *ByProjectKeyInStoreKeyByStoreKeyCartsByIDRequestMethodGet {
	return &ByProjectKeyInStoreKeyByStoreKeyCartsByIDRequestMethodGet{
		url:    fmt.Sprintf("/%s/in-store/key=%s/carts/%s", rb.projectKey, rb.storeKey, rb.id),
		client: rb.client,
	}
}

/**
*	Updates a [Cart](ctp:api:type:Cart) in the Store specified by `storeKey`.
*	If the Cart exists in the Project but does not have the `store` field, or the `store` field references a different Store, a [ResourceNotFound](ctp:api:type:ResourceNotFoundError) error is returned.
*
 */
func (rb *ByProjectKeyInStoreKeyByStoreKeyCartsByIDRequestBuilder) Post(body CartUpdate) *ByProjectKeyInStoreKeyByStoreKeyCartsByIDRequestMethodPost {
	return &ByProjectKeyInStoreKeyByStoreKeyCartsByIDRequestMethodPost{
		body:   body,
		url:    fmt.Sprintf("/%s/in-store/key=%s/carts/%s", rb.projectKey, rb.storeKey, rb.id),
		client: rb.client,
	}
}

func (rb *ByProjectKeyInStoreKeyByStoreKeyCartsByIDRequestBuilder) Delete() *ByProjectKeyInStoreKeyByStoreKeyCartsByIDRequestMethodDelete {
	return &ByProjectKeyInStoreKeyByStoreKeyCartsByIDRequestMethodDelete{
		url:    fmt.Sprintf("/%s/in-store/key=%s/carts/%s", rb.projectKey, rb.storeKey, rb.id),
		client: rb.client,
	}
}
