package main

import (
	"log"
	"time"
)

type pushEvent struct {
	After   string      `json:"after"`
	BaseRef interface{} `json:"base_ref"`
	Before  string      `json:"before"`
	Commits []struct {
		Added  []interface{} `json:"added"`
		Author struct {
			Email    string `json:"email"`
			Name     string `json:"name"`
			Username string `json:"username"`
		} `json:"author"`
		Committer struct {
			Email    string `json:"email"`
			Name     string `json:"name"`
			Username string `json:"username"`
		} `json:"committer"`
		Distinct  bool          `json:"distinct"`
		ID        string        `json:"id"`
		Message   string        `json:"message"`
		Modified  []string      `json:"modified"`
		Removed   []interface{} `json:"removed"`
		Timestamp string        `json:"timestamp"`
		TreeID    string        `json:"tree_id"`
		URL       string        `json:"url"`
	} `json:"commits"`
	Compare    string `json:"compare"`
	Created    bool   `json:"created"`
	Deleted    bool   `json:"deleted"`
	Forced     bool   `json:"forced"`
	HeadCommit struct {
		Added  []interface{} `json:"added"`
		Author struct {
			Email    string `json:"email"`
			Name     string `json:"name"`
			Username string `json:"username"`
		} `json:"author"`
		Committer struct {
			Email    string `json:"email"`
			Name     string `json:"name"`
			Username string `json:"username"`
		} `json:"committer"`
		Distinct  bool          `json:"distinct"`
		ID        string        `json:"id"`
		Message   string        `json:"message"`
		Modified  []string      `json:"modified"`
		Removed   []interface{} `json:"removed"`
		Timestamp string        `json:"timestamp"`
		TreeID    string        `json:"tree_id"`
		URL       string        `json:"url"`
	} `json:"head_commit"`
	Pusher struct {
		Email string `json:"email"`
		Name  string `json:"name"`
	} `json:"pusher"`
	Ref        string `json:"ref"`
	Repository struct {
		ArchiveURL       string      `json:"archive_url"`
		AssigneesURL     string      `json:"assignees_url"`
		BlobsURL         string      `json:"blobs_url"`
		BranchesURL      string      `json:"branches_url"`
		CloneURL         string      `json:"clone_url"`
		CollaboratorsURL string      `json:"collaborators_url"`
		CommentsURL      string      `json:"comments_url"`
		CommitsURL       string      `json:"commits_url"`
		CompareURL       string      `json:"compare_url"`
		ContentsURL      string      `json:"contents_url"`
		ContributorsURL  string      `json:"contributors_url"`
		CreatedAt        int         `json:"created_at"`
		DefaultBranch    string      `json:"default_branch"`
		Description      string      `json:"description"`
		DownloadsURL     string      `json:"downloads_url"`
		EventsURL        string      `json:"events_url"`
		Fork             bool        `json:"fork"`
		Forks            int         `json:"forks"`
		ForksCount       int         `json:"forks_count"`
		ForksURL         string      `json:"forks_url"`
		FullName         string      `json:"full_name"`
		GitCommitsURL    string      `json:"git_commits_url"`
		GitRefsURL       string      `json:"git_refs_url"`
		GitTagsURL       string      `json:"git_tags_url"`
		GitURL           string      `json:"git_url"`
		HasDownloads     bool        `json:"has_downloads"`
		HasIssues        bool        `json:"has_issues"`
		HasPages         bool        `json:"has_pages"`
		HasWiki          bool        `json:"has_wiki"`
		Homepage         interface{} `json:"homepage"`
		HooksURL         string      `json:"hooks_url"`
		HTMLURL          string      `json:"html_url"`
		ID               int         `json:"id"`
		IssueCommentURL  string      `json:"issue_comment_url"`
		IssueEventsURL   string      `json:"issue_events_url"`
		IssuesURL        string      `json:"issues_url"`
		KeysURL          string      `json:"keys_url"`
		LabelsURL        string      `json:"labels_url"`
		Language         interface{} `json:"language"`
		LanguagesURL     string      `json:"languages_url"`
		MasterBranch     string      `json:"master_branch"`
		MergesURL        string      `json:"merges_url"`
		MilestonesURL    string      `json:"milestones_url"`
		MirrorURL        interface{} `json:"mirror_url"`
		Name             string      `json:"name"`
		NotificationsURL string      `json:"notifications_url"`
		OpenIssues       int         `json:"open_issues"`
		OpenIssuesCount  int         `json:"open_issues_count"`
		Owner            struct {
			Email string `json:"email"`
			Name  string `json:"name"`
		} `json:"owner"`
		Private         bool   `json:"private"`
		PullsURL        string `json:"pulls_url"`
		PushedAt        int    `json:"pushed_at"`
		ReleasesURL     string `json:"releases_url"`
		Size            int    `json:"size"`
		SSHURL          string `json:"ssh_url"`
		Stargazers      int    `json:"stargazers"`
		StargazersCount int    `json:"stargazers_count"`
		StargazersURL   string `json:"stargazers_url"`
		StatusesURL     string `json:"statuses_url"`
		SubscribersURL  string `json:"subscribers_url"`
		SubscriptionURL string `json:"subscription_url"`
		SvnURL          string `json:"svn_url"`
		TagsURL         string `json:"tags_url"`
		TeamsURL        string `json:"teams_url"`
		TreesURL        string `json:"trees_url"`
		UpdatedAt       string `json:"updated_at"`
		URL             string `json:"url"`
		Watchers        int    `json:"watchers"`
		WatchersCount   int    `json:"watchers_count"`
	} `json:"repository"`
	Sender struct {
		AvatarURL         string `json:"avatar_url"`
		EventsURL         string `json:"events_url"`
		FollowersURL      string `json:"followers_url"`
		FollowingURL      string `json:"following_url"`
		GistsURL          string `json:"gists_url"`
		GravatarID        string `json:"gravatar_id"`
		HTMLURL           string `json:"html_url"`
		ID                int    `json:"id"`
		Login             string `json:"login"`
		OrganizationsURL  string `json:"organizations_url"`
		ReceivedEventsURL string `json:"received_events_url"`
		ReposURL          string `json:"repos_url"`
		SiteAdmin         bool   `json:"site_admin"`
		StarredURL        string `json:"starred_url"`
		SubscriptionsURL  string `json:"subscriptions_url"`
		Type              string `json:"type"`
		URL               string `json:"url"`
	} `json:"sender"`
}

type createEvent struct {
	Description  string `json:"description"`
	MasterBranch string `json:"master_branch"`
	PusherType   string `json:"pusher_type"`
	Ref          string `json:"ref"`
	RefType      string `json:"ref_type"`
	Repository   struct {
		ArchiveURL       string      `json:"archive_url"`
		AssigneesURL     string      `json:"assignees_url"`
		BlobsURL         string      `json:"blobs_url"`
		BranchesURL      string      `json:"branches_url"`
		CloneURL         string      `json:"clone_url"`
		CollaboratorsURL string      `json:"collaborators_url"`
		CommentsURL      string      `json:"comments_url"`
		CommitsURL       string      `json:"commits_url"`
		CompareURL       string      `json:"compare_url"`
		ContentsURL      string      `json:"contents_url"`
		ContributorsURL  string      `json:"contributors_url"`
		CreatedAt        string      `json:"created_at"`
		DefaultBranch    string      `json:"default_branch"`
		Description      string      `json:"description"`
		DownloadsURL     string      `json:"downloads_url"`
		EventsURL        string      `json:"events_url"`
		Fork             bool        `json:"fork"`
		Forks            int         `json:"forks"`
		ForksCount       int         `json:"forks_count"`
		ForksURL         string      `json:"forks_url"`
		FullName         string      `json:"full_name"`
		GitCommitsURL    string      `json:"git_commits_url"`
		GitRefsURL       string      `json:"git_refs_url"`
		GitTagsURL       string      `json:"git_tags_url"`
		GitURL           string      `json:"git_url"`
		HasDownloads     bool        `json:"has_downloads"`
		HasIssues        bool        `json:"has_issues"`
		HasPages         bool        `json:"has_pages"`
		HasWiki          bool        `json:"has_wiki"`
		Homepage         interface{} `json:"homepage"`
		HooksURL         string      `json:"hooks_url"`
		HTMLURL          string      `json:"html_url"`
		ID               int         `json:"id"`
		IssueCommentURL  string      `json:"issue_comment_url"`
		IssueEventsURL   string      `json:"issue_events_url"`
		IssuesURL        string      `json:"issues_url"`
		KeysURL          string      `json:"keys_url"`
		LabelsURL        string      `json:"labels_url"`
		Language         interface{} `json:"language"`
		LanguagesURL     string      `json:"languages_url"`
		MergesURL        string      `json:"merges_url"`
		MilestonesURL    string      `json:"milestones_url"`
		MirrorURL        interface{} `json:"mirror_url"`
		Name             string      `json:"name"`
		NotificationsURL string      `json:"notifications_url"`
		OpenIssues       int         `json:"open_issues"`
		OpenIssuesCount  int         `json:"open_issues_count"`
		Owner            struct {
			AvatarURL         string `json:"avatar_url"`
			EventsURL         string `json:"events_url"`
			FollowersURL      string `json:"followers_url"`
			FollowingURL      string `json:"following_url"`
			GistsURL          string `json:"gists_url"`
			GravatarID        string `json:"gravatar_id"`
			HTMLURL           string `json:"html_url"`
			ID                int    `json:"id"`
			Login             string `json:"login"`
			OrganizationsURL  string `json:"organizations_url"`
			ReceivedEventsURL string `json:"received_events_url"`
			ReposURL          string `json:"repos_url"`
			SiteAdmin         bool   `json:"site_admin"`
			StarredURL        string `json:"starred_url"`
			SubscriptionsURL  string `json:"subscriptions_url"`
			Type              string `json:"type"`
			URL               string `json:"url"`
		} `json:"owner"`
		Private         bool   `json:"private"`
		PullsURL        string `json:"pulls_url"`
		PushedAt        string `json:"pushed_at"`
		ReleasesURL     string `json:"releases_url"`
		Size            int    `json:"size"`
		SSHURL          string `json:"ssh_url"`
		StargazersCount int    `json:"stargazers_count"`
		StargazersURL   string `json:"stargazers_url"`
		StatusesURL     string `json:"statuses_url"`
		SubscribersURL  string `json:"subscribers_url"`
		SubscriptionURL string `json:"subscription_url"`
		SvnURL          string `json:"svn_url"`
		TagsURL         string `json:"tags_url"`
		TeamsURL        string `json:"teams_url"`
		TreesURL        string `json:"trees_url"`
		UpdatedAt       string `json:"updated_at"`
		URL             string `json:"url"`
		Watchers        int    `json:"watchers"`
		WatchersCount   int    `json:"watchers_count"`
	} `json:"repository"`
	Sender struct {
		AvatarURL         string `json:"avatar_url"`
		EventsURL         string `json:"events_url"`
		FollowersURL      string `json:"followers_url"`
		FollowingURL      string `json:"following_url"`
		GistsURL          string `json:"gists_url"`
		GravatarID        string `json:"gravatar_id"`
		HTMLURL           string `json:"html_url"`
		ID                int    `json:"id"`
		Login             string `json:"login"`
		OrganizationsURL  string `json:"organizations_url"`
		ReceivedEventsURL string `json:"received_events_url"`
		ReposURL          string `json:"repos_url"`
		SiteAdmin         bool   `json:"site_admin"`
		StarredURL        string `json:"starred_url"`
		SubscriptionsURL  string `json:"subscriptions_url"`
		Type              string `json:"type"`
		URL               string `json:"url"`
	} `json:"sender"`
}

type Version struct {
	App         string    `json:"app",storm:"index"`
	Tag         string    `json:"tag",storm:"id"`
	CurrentSHA  string    `json:"current_sha"`
	PreviousSHA string    `json:"previous_sha"`
	CreatedAt   time.Time `json:"created_at"`
}

var Info *log.Logger
var Error *log.Logger
