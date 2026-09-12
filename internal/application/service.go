package application

import (
	"context"
	"fmt"
	"time"

	"github.com/benpham-216/core-be/internal/domain"
	"github.com/benpham-216/core-be/internal/eventstore"
)

type Service struct { store eventstore.Store; now func() time.Time }
func NewService(store eventstore.Store) *Service { return &Service{store:store, now:time.Now().UTC} }

func (s *Service) Create(ctx context.Context, id, title, description, correlationID string, actor domain.Actor) (*domain.WorkOrder, error) {
	w, err := domain.NewWorkOrder(id,title,description); if err != nil { return nil, err }
	e := s.event(id, "work_order.created", actor, correlationID, map[string]any{"title":title,"description":description,"status":domain.Draft})
	if err := s.store.Append(ctx,id,0,e); err != nil { return nil, err }
	w.Version=1; return w,nil
}

func (s *Service) Transition(ctx context.Context, id string, expectedVersion int, next domain.Status, correlationID string, actor domain.Actor) (*domain.WorkOrder,error) {
	w, err := s.Get(ctx,id); if err != nil { return nil,err }
	if w.Version != expectedVersion { return nil,eventstore.ErrConcurrency }
	from:=w.Status; if err:=w.MoveTo(next,actor); err != nil{return nil,err}
	e:=s.event(id,"work_order.transitioned",actor,correlationID,map[string]any{"from":from,"to":next})
	if err:=s.store.Append(ctx,id,expectedVersion,e);err!=nil{return nil,err}
	w.Version++; return w,nil
}

func (s *Service) Get(ctx context.Context, id string) (*domain.WorkOrder,error) {
	events,err:=s.store.Load(ctx,id);if err!=nil{return nil,err};if len(events)==0{return nil,fmt.Errorf("work order %q not found",id)}
	w:=&domain.WorkOrder{ID:id}; for _,e:=range events { switch e.Type {case "work_order.created": w.Title,_=e.Payload["title"].(string);w.Description,_=e.Payload["description"].(string);w.Status=domain.Draft;case "work_order.transitioned": w.Status=domain.Status(e.Payload["to"].(string))};w.Version=e.Version };return w,nil
}
func (s *Service) Events(ctx context.Context,id string)([]domain.Event,error){return s.store.Load(ctx,id)}
func (s *Service) event(id,typ string,actor domain.Actor,correlationID string,payload map[string]any)domain.Event{return domain.Event{ID:fmt.Sprintf("evt-%d",s.now().UnixNano()),AggregateID:id,AggregateType:"work_order",Type:typ,Actor:actor,CorrelationID:correlationID,OccurredAt:s.now(),Payload:payload}}
