package jira
import (
    "github.com/sensu/uchiwa/uchiwa/config"
    "github.com/sensu/uchiwa/uchiwa/structs"
    "github.com/andygrunwald/go-jira"
    "errors"
    "io/ioutil"
)

func CreateJiraTicket(jiraconfig config.Jira, silencedata structs.Silence) (string, error) {
jiraClient, err := jira.NewClient(nil, jiraconfig.URL)
    if err != nil {
        return "", errors.New("couldn't create Jira client. Check the server address")
    }
    jiraClient.Authentication.SetBasicAuth(jiraconfig.User, jiraconfig.Pass)
    i := jira.Issue{
        Fields: &jira.IssueFields{
            Assignee: &jira.User{
                Name: "-1",
            },
            Reporter: &jira.User{
                Name: jiraconfig.User,
            },
            Description: silencedata.Check + " alerted on " + silencedata.Subscription,
            Type: jira.IssueType{
                Name: "Incident",
            },
            Project: jira.Project{
                Key: jiraconfig.Project,
            },
            Summary: "Sensu alert: " + silencedata.Check + " on " + silencedata.Subscription + " in " + silencedata.Dc,
        },
    }
    issuenew, resp, err := jiraClient.Issue.Create(&i)
    if err != nil {
        body, _ := ioutil.ReadAll(resp.Body)
        return "", errors.New(string(body))
    }
	return jiraconfig.URL + "browse/" + issuenew.Key, nil
}
