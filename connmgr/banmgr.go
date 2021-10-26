package connmgr

import (
	"sync"
	"time"

	"github.com/pkt-cash/pktd/pktlog/log"
)

type dbs struct {
	bs          *DynamicBanScore
	lastUsedSec int64
}

type BanMgr struct {
	m  sync.Mutex
	bs map[string]dbs
}

func now() int64 {
	return time.Now().Unix()
}

func (b *BanMgr) GetScore(host string) *DynamicBanScore {
	b.m.Lock()
	if _, ok := b.bs[host]; !ok {
		log.Debugf("Create new banScore for [%s]", host)
		if b.bs == nil {
			b.bs = make(map[string]dbs)
		}
		b.bs[host] = dbs{
			bs:          &DynamicBanScore{},
			lastUsedSec: now(),
		}
	}
	bs := b.bs[host].bs
	b.m.Unlock()
	return bs
}
