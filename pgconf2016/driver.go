import (
	"database/sql"
	"fmt"
	"time"
)

type PgV2Driver struct {
	uri   string
	cache *cache.PgCache
	DB    *sql.DB
}

// START OMIT
func (pg *PgDriver) Initialize(uri string) error {
	if err := pg.Connect(uri); err != nil {
		return err
	}

	//Add Notify Listener
	reportProblem := func(ev pq.ListenerEventType, err error) {
		if err != nil {
			fmt.Println(err.Error())
		}
	}

	listener := pq.NewListener(pg.uri, 10*time.Second, time.Minute, reportProblem)
	err := listener.Listen("new_fields")
	if err != nil {
		panic(err)
	}
}

// END OMIT
