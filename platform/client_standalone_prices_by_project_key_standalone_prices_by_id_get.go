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

type ByProjectKeyStandalonePricesByIDRequestMethodGet struct {
	url     string
	client  *Client
	headers http.Header
	params  *ByProjectKeyStandalonePricesByIDRequestMethodGetInput
}

func (r *ByProjectKeyStandalonePricesByIDRequestMethodGet) Dump() map[string]interface{} {
	return map[string]interface{}{
		"url":    r.url,
		"params": r.params,
	}
}

type ByProjectKeyStandalonePricesByIDRequestMethodGetInput struct {
	Expand []string
}

func (input *ByProjectKeyStandalonePricesByIDRequestMethodGetInput) Values() url.Values {
	values := url.Values{}
	for _, v := range input.Expand {
		values.Add("expand", fmt.Sprintf("%v", v))
	}
	return values
}

func (rb *ByProjectKeyStandalonePricesByIDRequestMethodGet) Expand(v []string) *ByProjectKeyStandalonePricesByIDRequestMethodGet {
	if rb.params == nil {
		rb.params = &ByProjectKeyStandalonePricesByIDRequestMethodGetInput{}
	}
	rb.params.Expand = v
	return rb
}

func (rb *ByProjectKeyStandalonePricesByIDRequestMethodGet) WithQueryParams(input ByProjectKeyStandalonePricesByIDRequestMethodGetInput) *ByProjectKeyStandalonePricesByIDRequestMethodGet {
	rb.params = &input
	return rb
}
func (rb *ByProjectKeyStandalonePricesByIDRequestMethodGet) WithHeaders(headers http.Header) *ByProjectKeyStandalonePricesByIDRequestMethodGet {
	rb.headers = headers
	return rb
}
func (rb *ByProjectKeyStandalonePricesByIDRequestMethodGet) Execute(ctx context.Context) (result *StandalonePrice, err error) {
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
