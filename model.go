package purecloud

// ======================== A Conversation ========================

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

// ======================== Queries ========================

type AggregationQuery struct {
	Interval                     string                `json:"interval,omitempty"`
	Granularity                  string                `json:"granularity,omitempty"`
	GroupBy                      []string              `json:"groupBy,omitempty"`
	Filter                       *AnalyticsQueryFilter `json:"filter,omitempty"` // need to be a pointer so that it is not marshalled into an empty JSON object if not required
	Metrics                      []string              `json:"metrics,omitempty"`
	FlattenMultiValuedDimensions bool                  `json:"flattenMultivaluedDimensions,omitempty"`
}

type ConversationQuery struct {
	Interval            string                      `json:"interval,omitempty"`
	ConversationFilters []AnalyticsQueryFilter      `json:"conversationFilters,omitempty"`
	EvaluationFilters   []AnalyticsQueryFilter      `json:"evaluationFilters,omitempty"`
	SegmentFilters      []AnalyticsQueryFilter      `json:"segmentFilters,omitempty"`
	Aggregations        []AnalyticsQueryAggregation `json:"aggregations,omitempty"`
	Order               string                      `json:"order,omitempty"`
	OrderBy             []string                    `json:"orderBy,omitempty"`
}

// ======================== Responses ========================

type AggregateQueryResponse struct {
	Results []AggregateDataContainer `json:"results,omitempty"`
}

type AnalyticsConversationQueryResponse struct {
	Conversations []AnalyticsConversation `json:"conversations,omitempty"`
	Aggregations  []AggregationResult     `json:"aggregations,omitempty"`
}

// ======================== Common Structs for Embeddeding  ========================

type AggregateDataContainer struct {
	Group GroupMap              `json:"group,omitempty"`
	Data  []StatisticalResponse `json:"data,omitempty"`
}

type AggregateMetricData struct {
	Metric string     `json:"metric,omitempty"`
	Stats  MetricStat `json:"stats,omitempty"`
}

type AggregationRange struct {
	Gte float64 `json:"gte,omitempty"`
	Lt  float64 `json:"lt,omitempty"`
}

type AggregationResult struct {
	Type      string                   `json:"type,omitempty"`
	Dimension string                   `json:"dimension,omitempty"`
	Metric    string                   `json:"metric,omitempty"`
	Count     float64                  `json:"count,omitempty"`
	Results   []AggregationResultEntry `json:"results,omitempty"`
}

type AggregationResultEntry struct {
	Count float64 `json:"count,omitempty"`
	Value string  `json:"value,omitempty"`
	Gte   float64 `json:"gte,omitempty"`
	Lt    float64 `json:"lt,omitempty"`
}

type AnalyticsConversation struct {
	ConversationID    string                 `json:"conversationId,omitempty"`
	ConversationStart string                 `json:"conversationStart,omitempty"`
	Participants      []AnalyticsParticipant `json:"participants,omitempty"`
	Evaluations       []AnalyticsEvaluation  `json:"evaluations,omitempty"`
}

type AnalyticsEvaluation struct {
	EvaluationID           string  `json:"evaluationId,omitempty"`
	EvaluatorID            string  `json:"evaluatorId,omitempty"`
	UserID                 string  `json:"userId,omitempty"`
	EventTime              string  `json:"eventTime,omitempty"`
	QueueID                string  `json:"queueId,omitempty"`
	FormID                 string  `json:"formId,omitempty"`
	ContextID              string  `json:"contextId,omitempty"`
	FormName               string  `json:"formName,omitempty"`
	GetoTotalScore         float64 `json:"getoTotalScore,omitempty"`
	GetoTotalCriticalScore float64 `json:"getoTotalCriticalScore,omitempty"`
}

type AnalyticsParticipant struct {
	ParticipantID   string             `json:"participantId,omitempty"`
	ParticipantName string             `json:"participantName,omitempty"`
	UserID          string             `json:"userId,omitempty"`
	Purpose         string             `json:"purpose,omitempty"`
	Sessions        []AnalyticsSession `json:"sessions,omitempty"`
}

type AnalyticsQueryAggregation struct {
	Type      string             `json:"type,omitempty"`
	Dimension string             `json:"dimension,omitempty"`
	Metric    string             `json:"metric,omitempty"`
	Size      float64            `json:"size,omitempty"`
	Ranges    []AggregationRange `json:"AggregationRange,omitempty"`
}

type AnalyticsQueryClause struct {
	Type       string                    `json:"type,omitempty"`
	Predicates []AnalyticsQueryPredicate `json:"predicates,omitempty"`
}

type AnalyticsQueryFilter struct {
	Type       string                    `json:"type,omitempty"`
	Clauses    []AnalyticsQueryClause    `json:"clauses,omitempty"`
	Predicates []AnalyticsQueryPredicate `json:"predicates,omitempty"`
}

type AnalyticsQueryPredicate struct {
	Type         string `json:"type,omitempty"`
	Dimension    string `json:"dimension,omitempty"`
	PropertyType string `json:"propertyType,omitempty"`
	Property     string `json:"property,omitempty"`
	Metric       string `json:"metric,omitempty"`
	Operator     string `json:"operator,omitempty"`
	Value        string `json:"value,omitempty"`
}

type AnalyticsSession struct {
	MediaType string `json:"mediaType,omitempty"`
	SessionID string `json:"sessionId,omitempty"`
	Ani       string `json:"ani,omitempty"`
	Direction string `json:"direction,omitempty"`
	Dnis      string `json:"dnis,omitempty"`
}

type GroupMap struct {
	MediaType string `json:"mediaType,omitempty"`
	QueueID   string `json:"queueId,omitempty"`
}

type MetricStat struct {
	Max   float64 `json:"max,omitempty"`
	Count float64 `json:"count,omitempty"`
	Sum   float64 `json:"sum,omitempty"`
}

type StatisticalResponse struct {
	Interval string                `json:"interval,omitempty"`
	Metrics  []AggregateMetricData `json:"metrics,omitempty"`
}
