package platform

// Generated file, please do not change!!!

import (
	"encoding/json"
)

type GraphQLError struct {
	Message   string                 `json:"message"`
	Locations []GraphQLErrorLocation `json:"locations"`
	Path      []interface{}          `json:"path"`
}

func (obj GraphQLError) Error() string {
	if obj.Message != "" {
		return obj.Message
	}
	return "unknown GraphQLError: failed to parse error response"
}

type GraphQLErrorLocation struct {
	Line   int `json:"line"`
	Column int `json:"column"`
}

type GraphQLRequest struct {
	Query         string               `json:"query"`
	OperationName *string              `json:"operationName,omitempty"`
	Variables     *GraphQLVariablesMap `json:"variables,omitempty"`
}

type GraphQLResponse struct {
	Data   interface{}    `json:"data,omitempty"`
	Errors []GraphQLError `json:"errors"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj GraphQLResponse) MarshalJSON() ([]byte, error) {
	type Alias GraphQLResponse
	data, err := json.Marshal(struct {
		*Alias
	}{Alias: (*Alias)(&obj)})
	if err != nil {
		return nil, err
	}

	raw := make(map[string]interface{})
	if err := json.Unmarshal(data, &raw); err != nil {
		return nil, err
	}

	if raw["errors"] == nil {
		delete(raw, "errors")
	}

	return json.Marshal(raw)

}

type GraphQLVariablesMap map[string]interface{}
