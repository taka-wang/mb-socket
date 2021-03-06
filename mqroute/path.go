package mqroute

// mqtt route path
const (
	mbtcpOnceRead       = "once.read"
	mbtcpOnceWrite      = "once.write"
	mbtcpGetTimeout     = "timeout.read"
	mbtcpSetTimeout     = "timeout.update"
	mbtcpCreatePoll     = "poll.create"
	mbtcpUpdatePoll     = "poll.update"
	mbtcpGetPoll        = "poll.read"
	mbtcpDeletePoll     = "poll.delete"
	mbtcpTogglePoll     = "poll.toggle"
	mbtcpGetPollHistory = "poll.history"
	mbtcpGetPolls       = "polls.read"
	mbtcpDeletePolls    = "polls.delete"
	mbtcpTogglePolls    = "polls.toggle"
	mbtcpImportPolls    = "polls.import"
	mbtcpExportPolls    = "polls.export"
	mbtcpData           = "data"
	mbtcpCreateFilter   = "filter.create"
	mbtcpUpdateFilter   = "filter.update"
	mbtcpGetFilter      = "filter.read"
	mbtcpDeleteFilter   = "filter.delete"
	mbtcpToggleFilter   = "filter.toggle"
	mbtcpGetFilters     = "filters.read"
	mbtcpDeleteFilters  = "filters.delete"
	mbtcpToggleFilters  = "filters.toggle"
	mbtcpImportFilters  = "filters.import"
	mbtcpExportFilters  = "filters.export"
)
