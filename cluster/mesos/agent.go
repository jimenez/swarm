package mesos

import (
	"github.com/docker/swarm/cluster"
	"github.com/mesos/mesos-go/mesosproto"
)

type agent struct {
	id     string
	offers map[string]*mesosproto.Offer
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
	s.offers[offer.Id.GetValue()] = offer
}

func (s *agent) addTask(task *task) {
	s.tasks[task.TaskInfo.TaskId.GetValue()] = task
}

func (s *agent) removeOffer(offerID string) bool {
	found := false
	_, found = s.offers[offerID]
	if found {
		delete(s.offers, offerID)
	}
	return found
}

func (s *agent) removeTask(taskID string) bool {
	found := false
	_, found = s.tasks[taskID]
	if found {
		delete(s.tasks, taskID)
	}
	return found
}

func (s *agent) empty() bool {
	return len(s.offers) == 0 && len(s.tasks) == 0
}

func (s *agent) getOffers() map[string]*mesosproto.Offer {
	return s.offers
}

func (s *agent) getTasks() map[string]*task {
	return s.tasks
}
