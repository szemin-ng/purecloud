package analytics

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

type ObservationQuery struct {
	Filter  *AnalyticsQueryFilter `json:"filter,omitempty"`
	Metrics []string              `json:"metrics,omitempty"`
}

// ======================== Responses ========================

type AggregateQueryResponse struct {
	Results []AggregateDataContainer `json:"results,omitempty"`
}

type AnalyticsConversationQueryResponse struct {
	Conversations []AnalyticsConversation `json:"conversations,omitempty"`
	Aggregations  []AggregationResult     `json:"aggregations,omitempty"`
}

type ObservationQueryResponse struct {
	Results []ObservationDataContainer `json:"results,omitempty"`
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
	Max         float64 `json:"max,omitempty"`
	Min         float64 `json:"min,omitempty"`
	Count       float64 `json:"count,omitempty"`
	Sum         float64 `json:"sum,omitempty"`
	Current     float64 `json:"current,omitempty"`
	Ratio       float64 `json:"ratio,omitempty"`
	Numerator   float64 `json:"numerator,omitempty"`
	Denominator float64 `json:"denominator,omitempty"`
	Target      float64 `json:"target,omitempty"`
}

type ObservationDataContainer struct {
	Group GroupMap              `json:"group,omitempty"`
	Data  []AggregateMetricData `json:"data,omitempty"`
}

type StatisticalResponse struct {
	Interval string                `json:"interval,omitempty"`
	Metrics  []AggregateMetricData `json:"metrics,omitempty"`
}
