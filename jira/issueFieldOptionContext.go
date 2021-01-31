package jira

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
)

type FieldOptionContextService struct{ client *Client }

type FieldOptionContextParams struct {
	FieldID     string
	ContextID   int
	OptionID    int
	OnlyOptions bool
}

type FieldContextOptionScheme struct {
	MaxResults int                             `json:"maxResults,omitempty"`
	StartAt    int                             `json:"startAt,omitempty"`
	Total      int                             `json:"total,omitempty"`
	IsLast     bool                            `json:"isLast,omitempty"`
	Values     []FieldContextOptionValueScheme `json:"values,omitempty"`
}

type FieldContextOptionValueScheme struct {
	ID       string `json:"id,omitempty"`
	Value    string `json:"value,omitempty"`
	Disabled bool   `json:"disabled,omitempty"`
	OptionID string `json:"optionId,omitempty"`
}

type FieldContextOptionListScheme struct {
	Options []FieldContextOptionValueScheme `json:"options"`
}

// Returns a paginated list of all custom field option for a context.
// Options are returned first then cascading options, in the order they display in Jira.
// Docs: https://developer.atlassian.com/cloud/jira/platform/rest/v3/api-group-issue-custom-field-options/#api-rest-api-3-field-fieldid-context-contextid-option-get
func (f *FieldOptionContextService) Gets(ctx context.Context, opts *FieldOptionContextParams, startAt, maxResults int) (result *FieldContextOptionScheme, response *Response, err error) {

	params := url.Values{}
	params.Add("startAt", strconv.Itoa(startAt))
	params.Add("maxResults", strconv.Itoa(maxResults))

	if len(opts.FieldID) != 0 {
		params.Add("fieldId", opts.FieldID)
	}

	if opts.ContextID != 0 {
		params.Add("contextId", strconv.Itoa(opts.ContextID))
	}

	if opts.OptionID != 0 {
		params.Add("optionId", strconv.Itoa(opts.OptionID))
	}

	if opts.OnlyOptions {
		params.Add("onlyOptions", "true")
	}

	var endpoint = fmt.Sprintf("rest/api/3/field/%v/context/%v/option?%v", opts.FieldID, opts.ContextID, params.Encode())
	request, err := f.client.newRequest(ctx, http.MethodGet, endpoint, nil)
	if err != nil {
		return
	}
	request.Header.Set("Accept", "application/json")

	response, err = f.client.Do(request)
	if err != nil {
		return
	}

	result = new(FieldContextOptionScheme)
	if err = json.Unmarshal(response.BodyAsBytes, &result); err != nil {
		return
	}

	return
}

type CreateCustomFieldOptionPayloadScheme struct {
	Options []FieldContextOptionValueScheme `json:"options"`
}

// Creates options and, where the custom select field is of the type Select List (cascading),
// cascading options for a custom select field. The options are added to a context of the field.
// Docs: https://developer.atlassian.com/cloud/jira/platform/rest/v3/api-group-issue-custom-field-options/#api-rest-api-3-field-fieldid-context-contextid-option-post
func (f *FieldOptionContextService) Create(ctx context.Context, fieldID string, contextID int, payload *CreateCustomFieldOptionPayloadScheme) (result *FieldContextOptionListScheme, response *Response, err error) {

	var endpoint = fmt.Sprintf("rest/api/3/field/%v/context/%v/option", fieldID, contextID)

	request, err := f.client.newRequest(ctx, http.MethodPost, endpoint, payload)
	if err != nil {
		return
	}
	request.Header.Set("Accept", "application/json")
	request.Header.Set("Content-Type", "application/json")

	response, err = f.client.Do(request)
	if err != nil {
		return
	}

	result = new(FieldContextOptionListScheme)
	if err = json.Unmarshal(response.BodyAsBytes, &result); err != nil {
		return
	}

	return
}

// Updates the options of a custom field.
// If any of the options are not found, no options are updated.
// Options where the values in the request match the current values aren't updated and aren't reported in the response.
// Docs: https://developer.atlassian.com/cloud/jira/platform/rest/v3/api-group-issue-custom-field-options/#api-rest-api-3-field-fieldid-context-contextid-option-put
func (f *FieldOptionContextService) Update(ctx context.Context, fieldID string, contextID int, payload *CreateCustomFieldOptionPayloadScheme) (result *FieldContextOptionListScheme, response *Response, err error) {

	var endpoint = fmt.Sprintf("rest/api/3/field/%v/context/%v/option", fieldID, contextID)

	request, err := f.client.newRequest(ctx, http.MethodPut, endpoint, payload)
	if err != nil {
		return
	}
	request.Header.Set("Accept", "application/json")
	request.Header.Set("Content-Type", "application/json")

	response, err = f.client.Do(request)
	if err != nil {
		return
	}

	result = new(FieldContextOptionListScheme)
	if err = json.Unmarshal(response.BodyAsBytes, &result); err != nil {
		return
	}

	return
}

// Deletes a custom field option.
// Options with cascading options cannot be deleted without deleting the cascading options first.
func (f *FieldOptionContextService) Delete(ctx context.Context, fieldID string, contextID, optionID int) (response *Response, err error) {

	var endpoint = fmt.Sprintf("rest/api/3/field/%v/context/%v/option/%v", fieldID, contextID, optionID)

	request, err := f.client.newRequest(ctx, http.MethodDelete, endpoint, nil)
	if err != nil {
		return
	}

	response, err = f.client.Do(request)
	if err != nil {
		return
	}

	return
}
