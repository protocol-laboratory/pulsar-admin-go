package padmin

type Option func(*options)

type options struct {
	backlogQuotaType BacklogQuotaType
}

func WithBacklogQuotaType(backlogQuotaType BacklogQuotaType) Option {
	return func(o *options) {
		o.backlogQuotaType = backlogQuotaType
	}
}
