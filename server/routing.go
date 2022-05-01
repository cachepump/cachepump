package server

// routing define a mapping of URL to them handlers.
var routing = Routing{
	"/":       getCache,
	"/health": health,
}
