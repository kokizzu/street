package zImport

import (
	"fmt"
	"sync"
	"time"

	"github.com/kpango/fastime"
)

type ImporterStat struct {
	Total int

	upserted  int
	skipped   int
	failed    int
	warn      int
	startTime *time.Time

	PrintEvery int

	mutex    *sync.Mutex
	warnings map[string]int
}

func (s *ImporterStat) Print(opt ...any) {
	progress := s.upserted + s.skipped + s.failed
	if s.startTime == nil {
		t := fastime.Now()
		s.startTime = &t
	}
	if s.PrintEvery == 0 {
		s.PrintEvery = 100
	}
	if len(opt) == 0 && progress%s.PrintEvery != 0 {
		return
	}
	fmt.Printf("\r    Upserted: %d, Skipped: %d, Warn: %d, Failed: %d | %.2f%% | %.1fs",
		s.upserted, s.skipped, s.warn, s.failed,
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

func (s *ImporterStat) Warn(str string) {
	if s.mutex == nil {
		s.mutex = &sync.Mutex{}
	}
	s.mutex.Lock()
	s.warnings[str] += 1
	s.mutex.Unlock()
}
