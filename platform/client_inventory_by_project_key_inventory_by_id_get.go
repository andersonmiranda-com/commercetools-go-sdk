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

type ByProjectKeyInventoryByIDRequestMethodGet struct {
	url     string
	client  *Client
	headers http.Header
	params  *ByProjectKeyInventoryByIDRequestMethodGetInput
}

func (r *ByProjectKeyInventoryByIDRequestMethodGet) Dump() map[string]interface{} {
	return map[string]interface{}{
		"url":    r.url,
		"params": r.params,
	}
}

type ByProjectKeyInventoryByIDRequestMethodGetInput struct {
	Expand []string
}

func (input *ByProjectKeyInventoryByIDRequestMethodGetInput) Values() url.Values {
	values := url.Values{}
	for _, v := range input.Expand {
		values.Add("expand", fmt.Sprintf("%v", v))
	}
	return values
}

func (rb *ByProjectKeyInventoryByIDRequestMethodGet) Expand(v []string) *ByProjectKeyInventoryByIDRequestMethodGet {
	if rb.params == nil {
		rb.params = &ByProjectKeyInventoryByIDRequestMethodGetInput{}
	}
	rb.params.Expand = v
	return rb
}

func (rb *ByProjectKeyInventoryByIDRequestMethodGet) WithQueryParams(input ByProjectKeyInventoryByIDRequestMethodGetInput) *ByProjectKeyInventoryByIDRequestMethodGet {
	rb.params = &input
	return rb
}
func (rb *ByProjectKeyInventoryByIDRequestMethodGet) WithHeaders(headers http.Header) *ByProjectKeyInventoryByIDRequestMethodGet {
	rb.headers = headers
	return rb
}
func (rb *ByProjectKeyInventoryByIDRequestMethodGet) Execute(ctx context.Context) (result *InventoryEntry, err error) {
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
