package bigcache

import (
	"encoding/json"
	"sync/atomic"
)

// Stats stores cache statistics
type Stats struct {
	hits       atomic.Int64
	misses     atomic.Int64
	delHits    atomic.Int64
	delMisses  atomic.Int64
	collisions atomic.Int64
}

// Hits returns the number of successfully found keys
func (s *Stats) Hits() int64 {
	return s.hits.Load()
}

func (s *Stats) AddHits(delta int64) int64 {
	return s.hits.Add(delta)
}

// Misses returns the number of not found keys
func (s *Stats) Misses() int64 {
	return s.misses.Load()
}

func (s *Stats) AddMisses(delta int64) int64 {
	return s.misses.Add(delta)
}

// DelHits returns the number of successfully deleted keys
func (s *Stats) DelHits() int64 {
	return s.delHits.Load()
}

func (s *Stats) AddDelHits(delta int64) int64 {
	return s.delHits.Add(delta)
}

// DelMisses returns the number of not deleted keys
func (s *Stats) DelMisses() int64 {
	return s.delMisses.Load()
}

func (s *Stats) AddDelMisses(delta int64) int64 {
	return s.delMisses.Add(delta)
}

// Collisions returns the number of happened key-collisions
func (s *Stats) Collisions() int64 {
	return s.collisions.Load()
}

func (s *Stats) AddCollisions(delta int64) int64 {
	return s.collisions.Add(delta)
}

func (s *Stats) Copy() *Stats {
	return NewStats(
		s.Hits(),
		s.Misses(),
		s.DelHits(),
		s.DelMisses(),
		s.Collisions(),
	)
}

func NewStats(hits, misses, delHits, delMisses, collisions int64) *Stats {
	stats := &Stats{
		hits:       atomic.Int64{},
		misses:     atomic.Int64{},
		delHits:    atomic.Int64{},
		delMisses:  atomic.Int64{},
		collisions: atomic.Int64{},
	}

	stats.hits.Store(hits)
	stats.misses.Store(misses)
	stats.delHits.Store(delHits)
	stats.delMisses.Store(delMisses)
	stats.collisions.Store(collisions)

	return stats
}

type jsonStats struct {
	Hits       int64 `json:"hits"`
	Misses     int64 `json:"misses"`
	DelHits    int64 `json:"delete_hits"`
	DelMisses  int64 `json:"delete_misses"`
	Collisions int64 `json:"collisions"`
}

func (s *Stats) MarshalJSON() ([]byte, error) {
	jsonStats := jsonStats{
		Hits:       s.hits.Load(),
		Misses:     s.misses.Load(),
		DelHits:    s.delHits.Load(),
		DelMisses:  s.delMisses.Load(),
		Collisions: s.collisions.Load(),
	}

	return json.Marshal(jsonStats)
}

func (s *Stats) UnmarshalJSON(data []byte) error {
	jsonStats := jsonStats{}
	if err := json.Unmarshal(data, &jsonStats); err != nil {
		return err
	}

	s.hits.Store(jsonStats.Hits)
	s.misses.Store(jsonStats.Misses)
	s.delHits.Store(jsonStats.DelHits)
	s.delMisses.Store(jsonStats.DelMisses)
	s.collisions.Store(jsonStats.Collisions)

	return nil
}
