import (
	"strconv"
	"strings"
)

// START OMIT
func (pg *PgDriver) UpdateCache(l *pq.Listener) {
	for {
		select {
		case n := <-l.Notify:
			switch n.Channel {
			case "new_field":
				farr := strings.Split(n.Extra, ", ")
				vt, _ := strconv.Atoi(farr[3])
				f := &pb.Field{
					Id:           farr[0],
					ObjectTypeId: farr[1],
					Title:        farr[2],
					ValueType:    pb.Field_ValueType(vt),
				}
				go pg.cache.AddFieldtoCache(f)
			}
		}
	}
}

// END OMIT
