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

type ByProjectKeyShoppingListsByIDRequestMethodPost struct {
	body    ShoppingListUpdate
	url     string
	client  *Client
	headers http.Header
	params  *ByProjectKeyShoppingListsByIDRequestMethodPostInput
}

func (r *ByProjectKeyShoppingListsByIDRequestMethodPost) Dump() map[string]interface{} {
	return map[string]interface{}{
		"url":    r.url,
		"params": r.params,
	}
}

type ByProjectKeyShoppingListsByIDRequestMethodPostInput struct {
	Expand []string
}

func (input *ByProjectKeyShoppingListsByIDRequestMethodPostInput) Values() url.Values {
	values := url.Values{}
	for _, v := range input.Expand {
		values.Add("expand", fmt.Sprintf("%v", v))
	}
	return values
}

func (rb *ByProjectKeyShoppingListsByIDRequestMethodPost) Expand(v []string) *ByProjectKeyShoppingListsByIDRequestMethodPost {
	if rb.params == nil {
		rb.params = &ByProjectKeyShoppingListsByIDRequestMethodPostInput{}
	}
	rb.params.Expand = v
	return rb
}

func (rb *ByProjectKeyShoppingListsByIDRequestMethodPost) WithQueryParams(input ByProjectKeyShoppingListsByIDRequestMethodPostInput) *ByProjectKeyShoppingListsByIDRequestMethodPost {
	rb.params = &input
	return rb
}
func (rb *ByProjectKeyShoppingListsByIDRequestMethodPost) WithHeaders(headers http.Header) *ByProjectKeyShoppingListsByIDRequestMethodPost {
	rb.headers = headers
	return rb
}
func (rb *ByProjectKeyShoppingListsByIDRequestMethodPost) Execute(ctx context.Context) (result *ShoppingList, err error) {
	data, err := serializeInput(rb.body)
	if err != nil {
		return nil, err
	}
	var queryParams url.Values
	if rb.params != nil {
		queryParams = rb.params.Values()
	} else {
		queryParams = url.Values{}
	}
	resp, err := rb.client.post(
		ctx,
		rb.url,
		queryParams,
		rb.headers,
		data,
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
	case 409, 400, 401, 403, 500, 502, 503:
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
