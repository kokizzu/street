package zImport

import (
	"fmt"
	"time"

	"github.com/kpango/fastime"
)

type ImporterStat struct {
	Total int

	upserted  int
	skipped   int
	failed    int
	startTime *time.Time
}

func (s *ImporterStat) Print() {
	progress := s.upserted + s.skipped + s.failed
	if progress%100 != 0 {
		return
	}
	if s.startTime == nil {
		t := fastime.Now()
		s.startTime = &t
	}
	fmt.Printf("\rUpserted: %d, Skipped: %d, Failed: %d | %.2f%% | %.1fs",
		s.upserted, s.skipped, s.failed,
		float64(progress*100)/float64(s.Total),
		fastime.Since(*s.startTime).Seconds())
}

func (s *ImporterStat) Skip() {
	s.skipped++
}

func (s *ImporterStat) Ok(ok bool) {
	if ok {
		s.upserted++
	} else {
		s.failed++
	}
}
