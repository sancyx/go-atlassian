package models

import (
	"errors"
)

var (
	ErrNoAdminOrganizationError = errors.New("admin: no organization id set")
	ErrNoAdminDomainIDError     = errors.New("admin: no domain id set")
	ErrNoEventIDError           = errors.New("admin: no event id set")
	ErrNoAdminPolicyError       = errors.New("admin: no organization policy id set")
	ErrNoAdminDirectoryIDError  = errors.New("admin: no directory id set")
	ErrNoAdminGroupIDError      = errors.New("admin: no group id set")
	ErrNoAdminGroupNameError    = errors.New("admin: no group name set")
	ErrNoAdminUserIDError       = errors.New("admin: no user id set")
	ErrNoAdminAccountIDError    = errors.New("admin: no account id set")
	ErrNoAdminUserTokenError    = errors.New("admin: no user token id set")

	ErrNoCustomerMailError        = errors.New("sm: no customer mail set")
	ErrNoCustomerDisplayNameError = errors.New("sm: no customer display name set")
	ErrNoKBQueryError             = errors.New("sm: no knowledge base query set")
	ErrNoOrganizationNameError    = errors.New("sm: no organization name set")
	ErrNoCommentBodyError         = errors.New("sm/jira: no comment body set")
	ErrNoServiceDeskIDError       = errors.New("sm: no service desk id set")
	ErrNoRequestTypeIDError       = errors.New("sm: no request type id set")
	ErrNoFileNameError            = errors.New("sm: no file name set")
	ErrNoFileReaderError          = errors.New("sm: no io.Reader set")

	ErrNoContentIDError        = errors.New("confluence: no content id set")
	ErrNoCQLError              = errors.New("confluence: no CQL query set")
	ErrNoContentTypeError      = errors.New("confluence: no content type set")
	ErrInvalidContentTypeError = errors.New("confluence: invalid content type: (page, comment, attachment)")
	ValidContentTypes          = []string{"page", "comment", "attachment"}
	ErrNoContentLabelError     = errors.New("confluence: no content label set")
	ErrNoContentPropertyError  = errors.New("confluence: no content property set")
	ErrNoSpaceNameError        = errors.New("confluence: no space name set")
	ErrNoSpaceKeyError         = errors.New("confluence: no space key set")

	ErrNoBoardIDError  = errors.New("agile: no board id set")
	ErrNoFilterIDError = errors.New("agile: no filter id set")
	ErrNoEpicIDError   = errors.New("agile: no epic id set")
	ErrNoSprintIDError = errors.New("agile: no sprint id set")

	ErrNoApplicationRoleError         = errors.New("jira: no application role key set")
	ErrNoDashboardIDError             = errors.New("jira: no dashboard id set")
	ErrNoGroupNameError               = errors.New("jira: no group name set")
	ErrNoGroupIDError                 = errors.New("jira: no group name set")
	ErrNoIssueKeyOrIDError            = errors.New("jira: no issue key/id set")
	ErrNoIssueSchemeError             = errors.New("jira: no jira.IssueScheme set")
	ErrNoTransitionIDError            = errors.New("jira: no transition id set")
	ErrNoAttachmentIDError            = errors.New("jira: no attachment id set")
	ErrNoAttachmentNameError          = errors.New("jira: no attachment filename set")
	ErrNoReaderError                  = errors.New("jira: no reader set")
	ErrNoCommentIDError               = errors.New("jira: no comment id set")
	ErrNoProjectIDError               = errors.New("jira: no project id set")
	ErrNoFieldIDError                 = errors.New("jira: no field id set")
	ErrNoFieldContextIDError          = errors.New("jira: no field context id set")
	ErrNoIssueTypesError              = errors.New("jira: no issue types id's set")
	ErrNoProjectsError                = errors.New("jira: no projects set")
	ErrNoContextOptionIDError         = errors.New("jira: no field context option id set")
	ErrNoTypeIDError                  = errors.New("jira: no link id set")
	ErrNoLinkTypeIDError              = errors.New("jira: no link type id set")
	ErrNoPriorityIDError              = errors.New("jira: no priority id set")
	ErrNoResolutionIDError            = errors.New("jira: no resolution id set")
	ErrNoJQLError                     = errors.New("jira: no sql set")
	ErrNoIssueTypeIDError             = errors.New("jira: no issue type id set")
	ErrNoIssueTypeScreenSchemeIDError = errors.New("jira: no issue type screen scheme id set")
	ErrNoScreenSchemeIDError          = errors.New("jira: no screen scheme id set")
	ErrNoAccountIDError               = errors.New("jira: no account id set")
	ErrNoWorklogIDError               = errors.New("jira: no worklog id set")
	ErrNpWorklogsError                = errors.New("jira: no worklog's id set")
	ErrNoPermissionSchemeIDError      = errors.New("jira: no permission scheme id set")
	ErrNoPermissionGrantIDError       = errors.New("jira: no permission grant id set")
	ErrNoComponentIDError             = errors.New("jira: no component id set")
	ErrProjectTypeKeyError            = errors.New("jira: no project type key set")
	ErrNoProjectNameError             = errors.New("jira: no project name set")
	ErrNoVersionIDError               = errors.New("jira: no version id set")
	ErrNoScreenNameError              = errors.New("jira: no screen name set")
	ErrNoScreenTabNameError           = errors.New("jira: no screen tab name set")
	ErrNoAccountSliceError            = errors.New("jira: no account id's set")
	ErrNoProjectKeySliceError         = errors.New("jira: no project key's set")
	ErrNoWorkflowIDError              = errors.New("jira: no workflow id set")
	ErrNoWorkflowSchemeIDError        = errors.New("jira: no workflow scheme id set")
	ErrNoScreenIDError                = errors.New("jira: no screen id set")
	ErrNoScreenTabIDError             = errors.New("jira: no screen tab id set")
	ErrNoFieldConfigurationNameError  = errors.New("jira: no field configuration name set")
)
