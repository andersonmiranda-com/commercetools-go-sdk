package platform

// Generated file, please do not change!!!

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

type ByProjectKeyInStoreKeyByStoreKeyMeShoppingListsKeyByKeyRequestMethodGet struct {
	url     string
	client  *Client
	headers http.Header
	params  *ByProjectKeyInStoreKeyByStoreKeyMeShoppingListsKeyByKeyRequestMethodGetInput
}

func (r *ByProjectKeyInStoreKeyByStoreKeyMeShoppingListsKeyByKeyRequestMethodGet) Dump() map[string]interface{} {
	return map[string]interface{}{
		"url":    r.url,
		"params": r.params,
	}
}

type ByProjectKeyInStoreKeyByStoreKeyMeShoppingListsKeyByKeyRequestMethodGetInput struct {
	Expand []string
}

func (input *ByProjectKeyInStoreKeyByStoreKeyMeShoppingListsKeyByKeyRequestMethodGetInput) Values() url.Values {
	values := url.Values{}
	for _, v := range input.Expand {
		values.Add("expand", fmt.Sprintf("%v", v))
	}
	return values
}

func (rb *ByProjectKeyInStoreKeyByStoreKeyMeShoppingListsKeyByKeyRequestMethodGet) Expand(v []string) *ByProjectKeyInStoreKeyByStoreKeyMeShoppingListsKeyByKeyRequestMethodGet {
	if rb.params == nil {
		rb.params = &ByProjectKeyInStoreKeyByStoreKeyMeShoppingListsKeyByKeyRequestMethodGetInput{}
	}
	rb.params.Expand = v
	return rb
}

func (rb *ByProjectKeyInStoreKeyByStoreKeyMeShoppingListsKeyByKeyRequestMethodGet) WithQueryParams(input ByProjectKeyInStoreKeyByStoreKeyMeShoppingListsKeyByKeyRequestMethodGetInput) *ByProjectKeyInStoreKeyByStoreKeyMeShoppingListsKeyByKeyRequestMethodGet {
	rb.params = &input
	return rb
}
func (rb *ByProjectKeyInStoreKeyByStoreKeyMeShoppingListsKeyByKeyRequestMethodGet) WithHeaders(headers http.Header) *ByProjectKeyInStoreKeyByStoreKeyMeShoppingListsKeyByKeyRequestMethodGet {
	rb.headers = headers
	return rb
}

/**
*	If a ShoppingList exists in a Project but does _not_ have the `store` field, or the `store` field references a different Store, the [ResourceNotFound](/errors#404-not-found-1) error is returned.
*
 */
func (rb *ByProjectKeyInStoreKeyByStoreKeyMeShoppingListsKeyByKeyRequestMethodGet) Execute(ctx context.Context) (result *ShoppingList, err error) {
	var queryParams url.Values
	if rb.params != nil {
		queryParams = rb.params.Values()
	} else {
		queryParams = url.Values{}
	}
	resp, err := rb.client.get(
		ctx,
		rb.url,
		queryParams,
		rb.headers,
	)

	if err != nil {
		return nil, err
	}
	content, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	switch resp.StatusCode {
	case 200:
		err = json.Unmarshal(content, &result)
		return result, nil
	case 400, 401, 403, 500, 502, 503:
		errorObj := ErrorResponse{}
		err = json.Unmarshal(content, &errorObj)
		if err != nil {
			return nil, err
		}
		return nil, errorObj
	case 404:
		return nil, ErrNotFound
	default:
		result := GenericRequestError{
			StatusCode: resp.StatusCode,
			Content:    content,
			Response:   resp,
		}
		return nil, result
	}

}
