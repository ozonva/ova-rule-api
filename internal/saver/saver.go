package saver

import (
	"errors"
	"sync"
	"time"

	"github.com/onsi/ginkgo"

	"github.com/ozonva/ova-rule-api/internal/flusher"
	"github.com/ozonva/ova-rule-api/internal/models"
)

type Saver interface {
	Init()
	Save(rule models.Rule) error
	Close()
}

// NewSaver возвращает Saver с поддержкой периодического сохранения.
func NewSaver(
	capacity uint,
	flusher flusher.Flusher,
	timeout time.Duration,
) Saver {
	return &saver{
		flusher: flusher,
		buffer:  make([]models.Rule, 0, capacity),
		timeout: timeout,
	}
}

type saver struct {
	sync.Mutex
	flusher  flusher.Flusher
	buffer   []models.Rule
	timeout  time.Duration
	notifyCh chan struct{}
}

func (s *saver) Init() {
	s.notifyCh = make(chan struct{}, 1)

	go func(ch <-chan struct{}) {
		defer ginkgo.GinkgoRecover()

		ticker := time.NewTicker(s.timeout)
		defer ticker.Stop()

		for {
			select {
			case <-ticker.C:
				s.flush()
			case <-s.notifyCh:
				s.flush()
				return
			}
		}
	}(s.notifyCh)
}

func (s *saver) Save(rule models.Rule) error {
	s.Lock()
	defer s.Unlock()

	if len(s.buffer) == cap(s.buffer) {
		return errors.New("saver buffer is full, wait")
	}

	s.buffer = append(s.buffer, rule)

	return nil
}

func (s *saver) Close() {
	close(s.notifyCh)
}

func (s *saver) flush() {
	s.Lock()
	defer s.Unlock()

	unsaved := s.flusher.Flush(s.buffer)

	s.buffer = s.buffer[:0]

	if len(unsaved) > 0 {
		s.buffer = append(s.buffer, unsaved...)
	}
}
