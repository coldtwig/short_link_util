package stat

import (
	"go/http-api/pkg/event"
	"log"
)

type StatServiceDeps struct {
	Eventbus       *event.EventBus
	StatRepository *StatRepository
}

type StatService struct {
	Eventbus       *event.EventBus
	StatRepository *StatRepository
}

func NewStatService(deps *StatServiceDeps) *StatService {
	return &StatService{
		Eventbus:       deps.Eventbus,
		StatRepository: deps.StatRepository,
	}
}

func (s *StatService) AddClick() {
	for msg := range s.Eventbus.Subscribe() {
		if msg.Type == event.EventLinkVisited {
			id, ok := msg.Data.(uint)
			if !ok {
				log.Fatalln("Bad EventLinkVisited Data:", msg.Data)
				continue
			}
			s.StatRepository.AddClick(id)
		}
	}
}
