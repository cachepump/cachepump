package pump

import (
	"sync"

	"github.com/robfig/cron/v3"
)

var storageCronIDs = make(StorageCronID)
var wrkMtx sync.RWMutex

// StorageCronID is a structure of storage for all cron job ID of ran jobs.
type StorageCronID map[string]cron.EntryID

// getWorkindSourceNames returns all names of source were define in configuration yaml file and are ran now.
func getWorkindSourceNames() (names []string) {
	wrkMtx.RLock()
	defer wrkMtx.RUnlock()
	for name := range storageCronIDs {
		names = append(names, name)
	}
	return names
}

// getWorkedID returns an ID of cron job for updating data from source with selected name if this job is ran now.
// If source with selected name is not run now argument 'ok' will be false.
func getWorkedID(name string) (id cron.EntryID, ok bool) {
	wrkMtx.Lock()
	defer wrkMtx.Unlock()
	id, ok = storageCronIDs[name]
	return id, ok
}

// setWorkedID saves a source with selected name and id to storage for all cron job ID of ran jobs.
func setWorkedID(name string, id cron.EntryID) {
	wrkMtx.Lock()
	defer wrkMtx.Unlock()
	storageCronIDs[name] = id
}

// delWorkindSource deletes a source with selected name from storage for all cron job ID of ran jobs.
func delWorkindSource(name string) {
	wrkMtx.Lock()
	defer wrkMtx.Unlock()
	delete(storageCronIDs, name)
}
