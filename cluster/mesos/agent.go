package mesos

import (
	"sync"

	"github.com/docker/swarm/cluster"
	"github.com/mesos/mesos-go/mesosproto"
)

type agent struct {
	sync.RWMutex

	id     string
	offers map[string]*mesosproto.Offer
	cpus   int64
	mem    int64
	tasks  map[string]*task
	engine *cluster.Engine
}

func newAgent(sid string, e *cluster.Engine) *agent {
	return &agent{
		id:     sid,
		offers: make(map[string]*mesosproto.Offer),
		tasks:  make(map[string]*task),
		engine: e,
	}
}

func (s *agent) addOffer(offer *mesosproto.Offer) {
	s.Lock()
	s.offers[offer.Id.GetValue()] = offer
	offers := []*mesosproto.Offer{offer}
	s.cpus += int64(sumScalarResourceValue(offers, "cpus"))
	s.mem += int64(sumScalarResourceValue(offers, "mem")) * 1024 * 1024
	s.Unlock()
}

func (s *agent) addTask(task *task) {
	s.Lock()
	s.tasks[task.TaskInfo.TaskId.GetValue()] = task
	s.Unlock()
}

func (s *slave) removeOffer(offerID string, used bool) bool {
	s.Lock()
	defer s.Unlock()
	offer, found := s.offers[offerID]
	if found {
		if !used {
			offers := []*mesosproto.Offer{offer}
			s.cpus -= int64(sumScalarResourceValue(offers, "cpus"))
			s.mem -= int64(sumScalarResourceValue(offers, "mem")) * 1024 * 1024
		}
		delete(s.offers, offerID)
	}
	return found
}

func (s *agent) removeTask(taskID string) bool {
	s.Lock()
	defer s.Unlock()
	found := false
	_, found = s.tasks[taskID]
	if found {
		delete(s.tasks, taskID)
	}
	return found
}

func (s *agent) empty() bool {
	s.RLock()
	defer s.RUnlock()
	return len(s.offers) == 0 && len(s.tasks) == 0
}

func (s *agent) getTasks() map[string]*task {
	s.RLock()
	defer s.RUnlock()
	return s.tasks
}
