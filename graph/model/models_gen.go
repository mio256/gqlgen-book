// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type Node interface {
	IsNode()
	GetID() string
}

type ProjectV2ItemContent interface {
	IsProjectV2ItemContent()
}

type AddProjectV2ItemByIDInput struct {
	ContentID string `json:"contentId"`
	ProjectID string `json:"projectId"`
}

type AddProjectV2ItemByIDPayload struct {
	Item *ProjectV2Item `json:"item"`
}

type Issue struct {
	ID           string                   `json:"id"`
	URL          string                   `json:"url"`
	Title        string                   `json:"title"`
	Closed       bool                     `json:"closed"`
	Number       int                      `json:"number"`
	Author       *User                    `json:"author"`
	Repository   *Repository              `json:"repository"`
	ProjectItems *ProjectV2ItemConnection `json:"projectItems"`
}

func (Issue) IsNode()            {}
func (this Issue) GetID() string { return this.ID }

func (Issue) IsProjectV2ItemContent() {}

type IssueConnection struct {
	Edges      []*IssueEdge `json:"edges"`
	Nodes      []*Issue     `json:"nodes"`
	PageInfo   *PageInfo    `json:"pageInfo"`
	TotalCount int          `json:"totalCount"`
}

type IssueEdge struct {
	Cursor string `json:"cursor"`
	Node   *Issue `json:"node"`
}

type PageInfo struct {
	EndCursor       *string `json:"endCursor"`
	HasNextPage     bool    `json:"hasNextPage"`
	HasPreviousPage bool    `json:"hasPreviousPage"`
	StartCursor     *string `json:"startCursor"`
}

type ProjectV2 struct {
	ID     string                   `json:"id"`
	Title  string                   `json:"title"`
	URL    string                   `json:"url"`
	Number int                      `json:"number"`
	Items  *ProjectV2ItemConnection `json:"items"`
	Owner  *User                    `json:"owner"`
}

func (ProjectV2) IsNode()            {}
func (this ProjectV2) GetID() string { return this.ID }

type ProjectV2Connection struct {
	Edges      []*ProjectV2Edge `json:"edges"`
	Nodes      []*ProjectV2     `json:"nodes"`
	PageInfo   *PageInfo        `json:"pageInfo"`
	TotalCount int              `json:"totalCount"`
}

type ProjectV2Edge struct {
	Cursor string     `json:"cursor"`
	Node   *ProjectV2 `json:"node"`
}

type ProjectV2Item struct {
	ID      string               `json:"id"`
	Project *ProjectV2           `json:"project"`
	Content ProjectV2ItemContent `json:"content"`
}

func (ProjectV2Item) IsNode()            {}
func (this ProjectV2Item) GetID() string { return this.ID }

type ProjectV2ItemConnection struct {
	Edges      []*ProjectV2ItemEdge `json:"edges"`
	Nodes      []*ProjectV2Item     `json:"nodes"`
	PageInfo   *PageInfo            `json:"pageInfo"`
	TotalCount int                  `json:"totalCount"`
}

type ProjectV2ItemEdge struct {
	Cursor string         `json:"cursor"`
	Node   *ProjectV2Item `json:"node"`
}

type PullRequest struct {
	ID           string                   `json:"id"`
	BaseRefName  string                   `json:"baseRefName"`
	Closed       bool                     `json:"closed"`
	HeadRefName  string                   `json:"headRefName"`
	URL          string                   `json:"url"`
	Number       int                      `json:"number"`
	Repository   *Repository              `json:"repository"`
	ProjectItems *ProjectV2ItemConnection `json:"projectItems"`
}

func (PullRequest) IsNode()            {}
func (this PullRequest) GetID() string { return this.ID }

func (PullRequest) IsProjectV2ItemContent() {}

type PullRequestConnection struct {
	Edges      []*PullRequestEdge `json:"edges"`
	Nodes      []*PullRequest     `json:"nodes"`
	PageInfo   *PageInfo          `json:"pageInfo"`
	TotalCount int                `json:"totalCount"`
}

type PullRequestEdge struct {
	Cursor string       `json:"cursor"`
	Node   *PullRequest `json:"node"`
}

type Repository struct {
	ID           string                 `json:"id"`
	Owner        *User                  `json:"owner"`
	Name         string                 `json:"name"`
	CreatedAt    string                 `json:"createdAt"`
	Issue        *Issue                 `json:"issue"`
	Issues       *IssueConnection       `json:"issues"`
	PullRequest  *PullRequest           `json:"pullRequest"`
	PullRequests *PullRequestConnection `json:"pullRequests"`
}

func (Repository) IsNode()            {}
func (this Repository) GetID() string { return this.ID }

type User struct {
	ID         string               `json:"id"`
	Name       string               `json:"name"`
	ProjectV2  *ProjectV2           `json:"projectV2"`
	ProjectV2s *ProjectV2Connection `json:"projectV2s"`
}

func (User) IsNode()            {}
func (this User) GetID() string { return this.ID }
