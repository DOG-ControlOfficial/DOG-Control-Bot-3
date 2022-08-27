package serverchat

import "time"

type serverchat struct {
	Token       string    `json:"token"`
	Event       time.Time `json:"Event"`
	Type     string    `json:"type"`
	Message  struct {
		Name   string `json:"name"`
		Sender struct {
			Name        string `json:"name"`
			DisplayName string `json:"displayName"`
			AvatarURL   string `json:"avatarUrl"`
			IsVerified  string `json:"isverified"`
			Type        string `json:"type"`
		} `json:"sender"`
		CreateTime time.Time `json:"createTime"`
		Text       string    `json:"text"`
		Thread     struct {
			Name              string `json:"name"`
			RetentionSettings struct {
				State string `json:"state"`
			} `json:"retentionSettings"`
		} `json:"thread"`
		Space struct {
			Name string `json:"name"`
			Type string `json:"type"`
		} `json:"space"`
		ArgumentText string `json:"argumentText"`
	} `json:"message"`
	User struct {
		Name        string `json:"name"`
		DisplayName string `json:"displayName"`
		AvatarURL   string `json:"avatarUrl"`
		IsVerified  string `json:"isverified"`
		Type        string `json:"type"`
	} `json:"user"`
	Space struct {
		Name        string `json:"name"`
		Type        string `json:"type"`
		DisplayName string `json:"displayName"`
	} `json:"space"`
	ConfigCompleteRedirectURL string `json:"configCompleteRedirectUrl"`
}

type Reply struct {
	Name string `json:"name,omitempty"`
}

type Reply struct {
	Text   string       `json:"text"`
	Thread *ReplyThread `json:"thread,omitempty"`
}
