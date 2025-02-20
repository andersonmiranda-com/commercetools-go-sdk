package platform

// Generated file, please do not change!!!

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
)

type ByProjectKeyCustomObjectsByContainerRequestMethodGet struct {
	url     string
	client  *Client
	headers http.Header
	params  *ByProjectKeyCustomObjectsByContainerRequestMethodGetInput
}

func (r *ByProjectKeyCustomObjectsByContainerRequestMethodGet) Dump() map[string]interface{} {
	return map[string]interface{}{
		"url":    r.url,
		"params": r.params,
	}
}

type ByProjectKeyCustomObjectsByContainerRequestMethodGetInput struct {
	Sort         []string
	Where        []string
	Expand       []string
	PredicateVar map[string][]string
	Limit        *int
	Offset       *int
	WithTotal    *bool
}

func (input *ByProjectKeyCustomObjectsByContainerRequestMethodGetInput) Values() url.Values {
	values := url.Values{}
	for _, v := range input.Sort {
		values.Add("sort", fmt.Sprintf("%v", v))
	}
	for _, v := range input.Where {
		values.Add("where", fmt.Sprintf("%v", v))
	}
	for _, v := range input.Expand {
		values.Add("expand", fmt.Sprintf("%v", v))
	}
	for k, v := range input.PredicateVar {
		for _, x := range v {
			values.Set(k, x)
		}
	}
	if input.Limit != nil {
		values.Add("limit", strconv.Itoa(*input.Limit))
	}
	if input.Offset != nil {
		values.Add("offset", strconv.Itoa(*input.Offset))
	}
	if input.WithTotal != nil {
		if *input.WithTotal {
			values.Add("withTotal", "true")
		} else {
			values.Add("withTotal", "false")
		}
	}
	return values
}

func (rb *ByProjectKeyCustomObjectsByContainerRequestMethodGet) Sort(v []string) *ByProjectKeyCustomObjectsByContainerRequestMethodGet {
	if rb.params == nil {
		rb.params = &ByProjectKeyCustomObjectsByContainerRequestMethodGetInput{}
	}
	rb.params.Sort = v
	return rb
}

func (rb *ByProjectKeyCustomObjectsByContainerRequestMethodGet) Where(v []string) *ByProjectKeyCustomObjectsByContainerRequestMethodGet {
	if rb.params == nil {
		rb.params = &ByProjectKeyCustomObjectsByContainerRequestMethodGetInput{}
	}
	rb.params.Where = v
	return rb
}

func (rb *ByProjectKeyCustomObjectsByContainerRequestMethodGet) Expand(v []string) *ByProjectKeyCustomObjectsByContainerRequestMethodGet {
	if rb.params == nil {
		rb.params = &ByProjectKeyCustomObjectsByContainerRequestMethodGetInput{}
	}
	rb.params.Expand = v
	return rb
}

func (rb *ByProjectKeyCustomObjectsByContainerRequestMethodGet) PredicateVar(v map[string][]string) *ByProjectKeyCustomObjectsByContainerRequestMethodGet {
	if rb.params == nil {
		rb.params = &ByProjectKeyCustomObjectsByContainerRequestMethodGetInput{}
	}
	rb.params.PredicateVar = v
	return rb
}

func (rb *ByProjectKeyCustomObjectsByContainerRequestMethodGet) Limit(v int) *ByProjectKeyCustomObjectsByContainerRequestMethodGet {
	if rb.params == nil {
		rb.params = &ByProjectKeyCustomObjectsByContainerRequestMethodGetInput{}
	}
	rb.params.Limit = &v
	return rb
}

func (rb *ByProjectKeyCustomObjectsByContainerRequestMethodGet) Offset(v int) *ByProjectKeyCustomObjectsByContainerRequestMethodGet {
	if rb.params == nil {
		rb.params = &ByProjectKeyCustomObjectsByContainerRequestMethodGetInput{}
	}
	rb.params.Offset = &v
	return rb
}

func (rb *ByProjectKeyCustomObjectsByContainerRequestMethodGet) WithTotal(v bool) *ByProjectKeyCustomObjectsByContainerRequestMethodGet {
	if rb.params == nil {
		rb.params = &ByProjectKeyCustomObjectsByContainerRequestMethodGetInput{}
	}
	rb.params.WithTotal = &v
	return rb
}

func (rb *ByProjectKeyCustomObjectsByContainerRequestMethodGet) WithQueryParams(input ByProjectKeyCustomObjectsByContainerRequestMethodGetInput) *ByProjectKeyCustomObjectsByContainerRequestMethodGet {
	rb.params = &input
	return rb
}
func (rb *ByProjectKeyCustomObjectsByContainerRequestMethodGet) WithHeaders(headers http.Header) *ByProjectKeyCustomObjectsByContainerRequestMethodGet {
	rb.headers = headers
	return rb
}
func (rb *ByProjectKeyCustomObjectsByContainerRequestMethodGet) Execute(ctx context.Context) (result *CustomObjectPagedQueryResponse, err error) {
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
