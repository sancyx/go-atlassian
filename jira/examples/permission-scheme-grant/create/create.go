package main

import (
	"context"
	"github.com/ctreminiom/go-atlassian/jira"
	"log"
	"os"
)

func main() {

	/*
		----------- Set an environment variable in git bash -----------
		export HOST="https://ctreminiom.atlassian.net/"
		export MAIL="MAIL_ADDRESS"
		export TOKEN="TOKEN_API"

		Docs: https://stackoverflow.com/questions/34169721/set-an-environment-variable-in-git-bash
	*/

	var (
		host  = os.Getenv("HOST")
		mail  = os.Getenv("MAIL")
		token = os.Getenv("TOKEN")
	)

	atlassian, err := jira.New(nil, host)
	if err != nil {
		log.Fatal(err)
	}

	atlassian.Auth.SetBasicAuth(mail, token)

	var permissionSchemeID = 10001

	grant := &jira.PermissionGrantPayloadScheme{
		Holder: &jira.PermissionGrantHolderScheme{
			Parameter: "jira-administrators-system",
			Type:      "group",
		},
		Permission: "EDIT_ISSUES",
	}

	permissionGrant, response, err := atlassian.Permission.Scheme.Grant.Create(context.Background(), permissionSchemeID, grant)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("HTTP Endpoint Used", response.Endpoint)
	log.Println(permissionGrant)
}
