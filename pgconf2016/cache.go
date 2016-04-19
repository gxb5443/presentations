import "sync"

type PgCache struct {
	fields fCache
}

type fCache struct {
	m map[string]*pb.Field
	sync.RWMutex
}

func (c *PgCache) AddFieldtoCache(f *pb.Field) {
	c.fields.Lock()
	defer c.fields.Unlock()
	c.fields.m[f.Id] = f
}
