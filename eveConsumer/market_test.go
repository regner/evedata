package eveConsumer

import (
	"testing"
	"time"

	"github.com/garyburd/redigo/redis"
)

func TestMarketAddRegion(t *testing.T) {
	r := ctx.Cache.Get()
	defer r.Close()
	eC.marketRegionAddRegion(1, time.Now().UTC().Unix(), r)
}

func TestMarketRegionConsumer(t *testing.T) {
	r := ctx.Cache.Get()
	defer r.Close()
	_, err := marketRegionConsumer(eC, r)
	if err != nil {
		t.Error(err)
		return
	}
}

func TestMarketHistoryTrigger(t *testing.T) {
	_, err := marketHistoryTrigger(eC)
	if err != nil {
		t.Error(err)
		return
	}
}

func TestMarketOrderConsumer(t *testing.T) {
	r := ctx.Cache.Get()
	defer r.Close()
	for {
		_, err := marketOrderConsumer(eC, r)
		if err != nil {
			t.Error(err)
			return
		}
		if i, _ := redis.Int(r.Do("SCARD", "EVEDATA_marketOrders")); i == 0 {
			break
		}
	}
}

// This is bugged due to the ESI Spec.
func TestMarketHistoryConsumer(t *testing.T) {
	r := ctx.Cache.Get()
	defer r.Close()
	j := 0
	for {
		j++
		_, err := marketHistoryConsumer(eC, r)
		if err != nil {
			t.Error(err)
			return
		}
		if i, _ := redis.Int(r.Do("SCARD", "EVEDATA_marketHistory")); i == 0 || j > 5 {
			break
		}
	}
}

func TestMarketMaintTrigger(t *testing.T) {
	_, err := marketMaintTrigger(eC)
	if err != nil {
		t.Error(err)
		return
	}
}