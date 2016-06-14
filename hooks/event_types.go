package hooks

type EventType string

var (
	CommitCommentEvent            EventType = "commit_comment"
	CreateEvent                   EventType = "create"
	DeleteEvent                   EventType = "delete"
	DeploymentEvent               EventType = "deployment"
	DeploymentStatusEvent         EventType = "deployment_status"
	ForkEvent                     EventType = "fork"
	GollumEvent                   EventType = "gollum"
	IssueCommentEvent             EventType = "issue_comment"
	IssuesEvent                   EventType = "issues"
	MemberEvent                   EventType = "member"
	MembershipEvent               EventType = "membership"
	PageBuildEvent                EventType = "page_build"
	PublicEvent                   EventType = "public"
	PullRequestEvent              EventType = "pull_request"
	PullRequestReviewCommentEvent EventType = "pull_request_review_comment"
	PushEvent                     EventType = "push"
	ReleaseEvent                  EventType = "release"
	RepositoryEvent               EventType = "repository"
	StatusEvent                   EventType = "status"
	TeamAddEvent                  EventType = "team_add"
	WatchEvent                    EventType = "watch"
)

func (e EventType) String() string {
	return string(e)
}
