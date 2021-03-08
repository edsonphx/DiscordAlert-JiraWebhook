package main

import "strings"

type WebhookResponse struct {
	Timestamp    uint64    `json:"timestamp"`
	WebhookEvent string    `json:"webhookEvent"`
	Comment      Comment   `json:"comment"`
	User         User      `json:"user"`
	Project      Project   `json:"project"`
	Issue        Issue     `json:"issue"`
	Changelog    ChangeLog `json:"changelog"`
}

func (w *WebhookResponse) getProjectName(toLower bool) string {

	var name string

	if w.Project.isEmpty() {
		name = w.Issue.Fields.Project.Name
	} else {
		name = w.Project.Name
	}

	if toLower {
		name = strings.ToLower(name)
	}

	return name
}

type Comment struct {
	ID     string `json:"id"`
	Author User   `json:"author"`
	Body   string `json:"body"`
}

type Project struct {
	ID          interface{} `json:"id"`
	Key         string      `json:"key"`
	Name        string      `json:"name"`
	ProjectLead User        `json:"projectLead"`
}

func (p *Project) isEmpty() bool {
	if p.Key == "" {
		return true
	}

	return false
}

func (p *Project) getProjectURL() string {
	return "https://" + JiraSiteName + ".atlassian.net/browse/" + p.Key
}

type Issue struct {
	Key    string `json:"key"`
	Fields Fields `json:"fields"`
}

func (i *Issue) getIssueURL() string {
	return "https://" + JiraSiteName + ".atlassian.net/browse/" + i.Key
}

type Fields struct {
	IssueType   BasicField `json:"issuetype"`
	Project     Project    `json:"project"`
	Priority    BasicField `json:"priority"`
	Status      BasicField `json:"status"`
	Description string     `json:"description"`
	Summary     string     `json:"summary"`
}

type BasicField struct {
	Name string `json:"name"`
}

type User struct {
	ID        string            `json:"accountId"`
	Name      string            `json:"displayName"`
	AvatarURL map[string]string `json:"avatarUrls"`
}

func (u User) getProfileURL() string {
	return "https://" + JiraSiteName + ".atlassian.net/jira/people/" + u.ID
}

type ChangeLog struct {
	ID    string  `json:"id"`
	Items []Items `json:"items"`
}

type Items struct {
	Field     string `json:"field"`
	LastValue string `json:"fromString"`
	NewValue  string `json:"toString"`
}
