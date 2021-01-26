package main

import (
	"container/list"
	"context"
	"log"
	"math/rand"
	"sync"
	"time"
)

type Request int
type Service struct {
	mu         sync.Mutex
	queue      *list.List
	sema       chan int
	loopSignal chan struct{}
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	service := NewService(ctx, 3)
	for i := 0; i < 10; i++ {
		if err := service.EnqueueRequest(Request(i)); err != nil {
			log.Fatalf("error sending request: %v", err)
		}
		<-time.After(time.Duration(rand.Intn(10)) * time.Millisecond)
	}
	for {
		time.Sleep(time.Second)
	}
}

func (s *Service) tickleLoop() {
	select {
	case s.loopSignal <- struct{}{}:
	default:

	}
}

func (s *Service) EnqueueRequest(req Request) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.queue.PushBack(req)
	log.Printf("Added request to queue with length %d\n", s.queue.Len())
	s.tickleLoop()
	return nil
}

func (s *Service) dequeue() Request {
	element := s.queue.Front()
	s.queue.Remove(element)
	return element.Value.(Request)
}

func (s *Service) replenish() {
	<-s.sema
	log.Printf("Replenishing semaphore, now %d/%d slots in use\n", len(s.sema), cap(s.sema))
	s.tickleLoop()
}

func (s *Service) process(req Request) {
	defer s.replenish()
	log.Printf("Processing request %v\n", req)
	// Simulate work
	<-time.After(time.Duration(rand.Intn(500)) * time.Millisecond)
}

func (s *Service) tryDequeue() {
	s.mu.Lock()
	defer s.mu.Unlock()
	if s.queue.Len() == 0 {
		return
	}
	select {
	case s.sema <- 1:
		req := s.dequeue()
		log.Printf("Dequeued request %v\n", req)
		go s.process(req)
	default:
		log.Printf("Recieved loop signal, but request limit is reached")
	}
}

func (s *Service) loop(ctx context.Context) {
	log.Println("Starting service loop")
	for {
		select {
		case <-s.loopSignal:
			s.tryDequeue()
		case <-ctx.Done():
			log.Printf("Lopp context cancelled")
			return
		}
	}
}

func NewService(ctx context.Context, requestLimit int) *Service {
	service := &Service{
		queue:      list.New(),
		sema:       make(chan int, requestLimit),
		loopSignal: make(chan struct{}, 1),
	}

	go service.loop(ctx)
	return service
}
