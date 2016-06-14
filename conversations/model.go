package conversations

type Conversation struct {
	ID             string        `json:"id,omitempty"`
	Name           string        `json:"name,omitempty"`
	StartTime      string        `json:"startTime"`
	EndTime        string        `json:"endTime,omitempty"`
	Address        string        `json:"address,omitempty"`
	Participants   []Participant `json:"participants"`
	RecordingState string        `json:"recordingState,omitempty"`
	SelfURI        string        `json:"selfUri,omitempty"`
}

type Participant struct {
	ID                     string             `json:"id,omitempty"`
	StartTime              string             `json:"startTime,omitempty"`
	EndTime                string             `json:"endTime,omitempty"`
	ConnectedTime          string             `json:"connectedTime,omitempty"`
	Name                   string             `json:"name,omitempty"`
	UserURI                string             `json:"userUri,omitempty"`
	UserID                 string             `json:"userId,omitempty"`
	ExternalContactID      string             `json:"externalContactId,omitempty"`
	QueueID                string             `json:"queueId,omitempty"`
	GroupID                string             `json:"groupId,omitempty"`
	QueueName              string             `json:"queueName,omitempty"`
	Purpose                string             `json:"purpose,omitempty"`
	ParticipantType        string             `json:"participantType,omitempty"`
	ConsultParticipantID   string             `json:"consultParticipantId,omitempty"`
	Address                string             `json:"address,omitempty"`
	Ani                    string             `json:"ani,omitempty"`
	Dnis                   string             `json:"dnis,omitempty"`
	Locale                 string             `json:"locale,omitempty"`
	WrapupRequired         bool               `json:"wrapupRequired,omitempty"`
	WrapupPrompt           string             `json:"wrapupPrompt,omitempty"`
	WrapupTimeoutMS        float64            `json:"wrapupTimeoutMS,omitempty"`
	WrapupSkipped          bool               `json:"wrapupSkipped,omitempty"`
	Wrapup                 *Wrapup            `json:"wrapup,omitempty"` // need to be a pointer so that it is not marshalled into an empty JSON object if not required
	MonitoredParticipantID string             `json:"monitoredParticipantId,omitempty"`
	Attributes             map[string]string  `json:"attributes,omitempty"`
	Calls                  []Call             `json:"calls,omitempty"`
	Callbacks              []Callback         `json:"callbacks,omitempty"`
	Chats                  []ConversationChat `json:"chats,omitempty"`
	Emails                 []Email            `json:"emails,omitempty"`
	SocialExpressions      []SocialExpression `json:"socialExpressions,omitempty"`
	Videos                 []Video            `json:"videos,omitempty"`
	Evaluations            []Evaluation       `json:"evaluations,omitempty"`
}

type Wrapup struct {
	Code            string  `json:"code,omitempty"`
	Name            string  `json:"name,omitempty"`
	Notes           string  `json:"notes,omitempty"`
	DurationSeconds float64 `json:"durationSeconds,omitempty"`
	EndTime         string  `json:"endTime,omitempty"`
	Provisional     bool    `json:"provisional,omitempty"`
}

type Call struct {
	State            string `json:"state,omitempty"`
	ID               string `json:"id,omitempty"`
	Direction        string `json:"direction,omitempty"`
	Recording        bool   `json:"recording,omitempty"`
	RecordingState   string `json:"recordingState,omitempty"`
	Muted            bool   `json:"muted,omitempty"`
	Confined         bool   `json:"confined,omitempty"`
	Held             bool   `json:"held,omitempty"`
	RecordingID      string `json:"recordingId,omitempty"`
	DisconnectType   string `json:"disconnectType,omitempty"`
	StartHoldTime    string `json:"startHoldTime,omitempty"`
	DocumentID       string `json:"documentId,omitempty"`
	ConnectedTime    string `json:"connectedTime,omitempty"`
	DisconnectedTime string `json:"disconnectedTime,omitempty"`
}

type Callback struct {
	State            string  `json:"state,omitempty"`
	ID               string  `json:"id,omitempty"`
	Direction        string  `json:"direction,omitempty"`
	Held             bool    `json:"held,omitempty"`
	DisconnectType   string  `json:"disconnectType,omitempty"`
	StartHoldTime    string  `json:"startHoldTime,omitempty"`
	ScriptID         string  `json:"scriptId,omitempty"`
	SkipEnabled      bool    `json:"skipEnabled,omitempty"`
	TimeoutSeconds   float64 `json:"timeoutSeconds,omitempty"`
	ConnectedTime    string  `json:"connectedTime,omitempty"`
	DisconnectedTime string  `json:"disconnectedTime,omitempty"`
}

type ConversationChat struct {
	State            string `json:"state,omitempty"`
	ID               string `json:"id,omitempty"`
	RoomID           string `json:"roomId,omitempty"`
	RecordingID      string `json:"recordingId,omitempty"`
	Held             bool   `json:"held,omitempty"`
	Direction        string `json:"direction,omitempty"`
	DisconnectType   string `json:"disconnectType,omitempty"`
	StartHoldTime    string `json:"startHoldTime,omitempty"`
	ConnectedTime    string `json:"connectedTime,omitempty"`
	DisconnectedTime string `json:"disconnectedTime,omitempty"`
}

type Email struct {
	State            string  `json:"state,omitempty"`
	ID               string  `json:"id,omitempty"`
	Held             bool    `json:"held,omitempty"`
	Subject          string  `json:"subject,omitempty"`
	MessagesSent     float64 `json:"messagesSent,omitempty"`
	Direction        string  `json:"direction,omitempty"`
	RecordingID      string  `json:"recordingId,omitempty"`
	DisconnectType   string  `json:"disconnectType,omitempty"`
	StartHoldTime    string  `json:"startHoldTime,omitempty"`
	ConnectedTime    string  `json:"connectedTime,omitempty"`
	DisconnectedTime string  `json:"disconnectedTime,omitempty"`
}

type SocialExpression struct {
	State            string `json:"state,omitempty"`
	ID               string `json:"id,omitempty"`
	SocialMediaID    string `json:"socialMediaId,omitempty"`
	SocialMediaHub   string `json:"socialMediaHub,omitempty"`
	SocialUserName   string `json:"socialUserName,omitempty"`
	PreviewText      string `json:"previewText,omitempty"`
	RecordingID      string `json:"recordingId,omitempty"`
	Held             bool   `json:"held,omitempty"`
	DisconnectType   string `json:"disconnectType,omitempty"`
	StartHoldTime    string `json:"startHoldTime,omitempty"`
	ConnectedTime    string `json:"connectedTime,omitempty"`
	DisconnectedTime string `json:"disconnectedTime,omitempty"`
}

type Video struct {
	State            string  `json:"state,omitempty"`
	ID               string  `json:"id,omitempty"`
	Context          string  `json:"context,omitempty"`
	AudioMuted       bool    `json:"audioMuted,omitempty"`
	VideoMuted       bool    `json:"videoMuted,omitempty"`
	SharingScreen    bool    `json:"sharingScreen,omitempty"`
	PeerCount        float64 `json:"peerCount,omitempty"`
	DisconnectType   string  `json:"disconnectType,omitempty"`
	ConnectedTime    string  `json:"connectedTime,omitempty"`
	DisconnectedTime string  `json:"disconnectedTime,omitempty"`
}

type Evaluation struct {
	ID             string `json:"id,omitempty"`
	Name           string `json:"name,omitempty"`
	Status         string `json:"status,omitempty"`
	AgentHasRead   bool   `json:"agentHasRead,omitempty"`
	ReleaseDate    string `json:"releaseDate,omitempty"`
	AssignedDate   string `json:"assignedDate,omitempty"`
	ChangedDate    string `json:"changedDate,omitempty"`
	NeverRelease   bool   `json:"neverRelease,omitempty"`
	ResourceID     string `json:"resourceId,omitempty"`
	ResourceType   string `json:"resourceType,omitempty"`
	Redacted       bool   `json:"redacted,omitempty"`
	IsScoringIndex bool   `json:"isScoringIndex,omitempty"`
	SelfURI        string `json:"selfUri,omitempty"`
}
