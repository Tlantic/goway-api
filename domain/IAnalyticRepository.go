package domain

type QueryOptions struct {
	Client string
	Product string
}

type IAnalyticRepository interface {
	GetLastHourRequest(queryOptions *QueryOptions) (error,[]LogRecord)
}
