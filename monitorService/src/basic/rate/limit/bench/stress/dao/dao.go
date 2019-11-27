package dao

import (
	"context"

	"github.com/DazzlingSun/monitorService/src/basic/cache/memcache"
	"github.com/DazzlingSun/monitorService/src/basic/cache/redis"
	xsql "github.com/DazzlingSun/monitorService/src/basic/database/sql"
	"github.com/DazzlingSun/monitorService/src/basic/rate/limit/bench/stress/conf"
)

// Dao dao
type Dao struct {
	c     *conf.Config
	mc    *memcache.Pool
	redis *redis.Pool
	db    *xsql.DB
}

// New init mysql db
func New(c *conf.Config) (dao *Dao) {
	dao = &Dao{
		c:     c,
		mc:    memcache.NewPool(c.Memcache),
		redis: redis.NewPool(c.Redis),
		db:    xsql.NewMySQL(c.MySQL),
	}
	return
}

// Close close the resource.
func (d *Dao) Close() {
	d.mc.Close()
	d.redis.Close()
	d.db.Close()
}

// Ping dao ping
func (d *Dao) Ping(c context.Context) error {
	return d.pingMC(c)
}

// pingMc ping
func (d *Dao) pingMC(c context.Context) (err error) {
	conn := d.mc.Get(c)
	defer conn.Close()
	return
}
