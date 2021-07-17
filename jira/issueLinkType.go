package jira

import (
	"context"
	"fmt"
	"net/http"
)

type IssueLinkTypeService struct{ client *Client }

type IssueLinkTypeSearchScheme struct {
	IssueLinkTypes []*LinkTypeScheme `json:"issueLinkTypes,omitempty"`
}

type IssueLinkTypeScheme struct {
	ID      string `json:"id"`
	Name    string `json:"name"`
	Inward  string `json:"inward"`
	Outward string `json:"outward"`
	Self    string `json:"self"`
}

// Gets returns a list of all issue link types.
// Docs: https://docs.go-atlassian.io/jira-software-cloud/issues/link/types#get-issue-link-types
// Official Docs: https://developer.atlassian.com/cloud/jira/platform/rest/v3/api-group-issue-link-types/#api-rest-api-3-issuelinktype-get
func (i *IssueLinkTypeService) Gets(ctx context.Context) (result *IssueLinkTypeSearchScheme, response *ResponseScheme,
	err error) {

	var endpoint = "rest/api/3/issueLinkType"

	request, err := i.client.newRequest(ctx, http.MethodGet, endpoint, nil)
	if err != nil {
		return
	}

	request.Header.Set("Accept", "application/json")

	response, err = i.client.call(request, nil)
	if err != nil {
		return
	}

	return
}

// Get returns an issue link type.
// Docs: https://docs.go-atlassian.io/jira-software-cloud/issues/link/types#get-issue-link-type
// Official Docs: https://developer.atlassian.com/cloud/jira/platform/rest/v3/api-group-issue-link-types/#api-rest-api-3-issuelinktype-issuelinktypeid-get
func (i *IssueLinkTypeService) Get(ctx context.Context, issueLinkTypeID string) (result *LinkTypeScheme,
	response *ResponseScheme, err error) {

	if len(issueLinkTypeID) == 0 {
		return nil, nil, notIssueLinkTypeIDError
	}

	var endpoint = fmt.Sprintf("rest/api/3/issueLinkType/%v", issueLinkTypeID)

	request, err := i.client.newRequest(ctx, http.MethodGet, endpoint, nil)
	if err != nil {
		return
	}

	request.Header.Set("Accept", "application/json")

	response, err = i.client.call(request, &result)
	if err != nil {
		return
	}

	return
}

// Create creates an issue link type.
// Use this operation to create descriptions of the reasons why issues are linked.
// The issue link type consists of a name and descriptions for a link's inward and outward relationships.
// Docs: https://docs.go-atlassian.io/jira-software-cloud/issues/link/types#create-issue-link-type
// Official Docs: https://developer.atlassian.com/cloud/jira/platform/rest/v3/api-group-issue-link-types/#api-rest-api-3-issuelinktype-post
func (i *IssueLinkTypeService) Create(ctx context.Context, payload *LinkTypeScheme) (result *LinkTypeScheme,
	response *ResponseScheme, err error) {

	var endpoint = "rest/api/3/issueLinkType"

	payloadAsReader, err := transformStructToReader(payload)
	if err != nil {
		return nil, nil, err
	}

	request, err := i.client.newRequest(ctx, http.MethodPost, endpoint, payloadAsReader)
	if err != nil {
		return
	}

	request.Header.Set("Accept", "application/json")
	request.Header.Set("Content-Type", "application/json")

	response, err = i.client.call(request, &result)
	if err != nil {
		return
	}

	return
}

// Update updates an issue link type.
// Docs: https://docs.go-atlassian.io/jira-software-cloud/issues/link/types#update-issue-link-type
// Official Docs: https://developer.atlassian.com/cloud/jira/platform/rest/v3/api-group-issue-link-types/#api-rest-api-3-issuelinktype-issuelinktypeid-put
func (i *IssueLinkTypeService) Update(ctx context.Context, issueLinkTypeID string, payload *LinkTypeScheme) (
	result *LinkTypeScheme, response *ResponseScheme, err error) {

	if len(issueLinkTypeID) == 0 {
		return nil, nil, notIssueLinkTypeIDError
	}

	var endpoint = fmt.Sprintf("rest/api/3/issueLinkType/%v", issueLinkTypeID)

	payloadAsReader, err := transformStructToReader(payload)
	if err != nil {
		return nil, nil, err
	}

	request, err := i.client.newRequest(ctx, http.MethodPut, endpoint, payloadAsReader)
	if err != nil {
		return
	}

	request.Header.Set("Accept", "application/json")
	request.Header.Set("Content-Type", "application/json")

	response, err = i.client.call(request, &result)
	if err != nil {
		return
	}

	return
}

// Delete deletes an issue link type.
// Docs: https://docs.go-atlassian.io/jira-software-cloud/issues/link/types#delete-issue-link-type
// Official Docs: https://developer.atlassian.com/cloud/jira/platform/rest/v3/api-group-issue-link-types/#api-rest-api-3-issuelinktype-issuelinktypeid-delete
func (i *IssueLinkTypeService) Delete(ctx context.Context, issueLinkTypeID string) (response *ResponseScheme, err error) {

	if len(issueLinkTypeID) == 0 {
		return nil, notIssueLinkTypeIDError
	}

	var endpoint = fmt.Sprintf("rest/api/3/issueLinkType/%v", issueLinkTypeID)

	request, err := i.client.newRequest(ctx, http.MethodDelete, endpoint, nil)
	if err != nil {
		return
	}

	response, err = i.client.call(request, nil)
	if err != nil {
		return
	}

	return
}

var (
	notIssueLinkTypeIDError = fmt.Errorf("error!, please provide a valid issueLinkTypeID value")
)
