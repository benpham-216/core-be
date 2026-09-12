package application_test

import (
 "context"
 "errors"
 "testing"
 "github.com/benpham-216/core-be/internal/application"
 "github.com/benpham-216/core-be/internal/domain"
 "github.com/benpham-216/core-be/internal/eventstore"
)
func TestLifecycleRejectsShortcutAndAgentApproval(t *testing.T){
 ctx:=context.Background(); s:=application.NewService(eventstore.NewMemoryStore()); human:=domain.Actor{ID:"h",Type:"human"}; agent:=domain.Actor{ID:"a",Type:"agent"}
 w,err:=s.Create(ctx,"wo-1","Test","", "c-1",agent);if err!=nil{t.Fatal(err)}
 if _,err=s.Transition(ctx,w.ID,w.Version,domain.Completed,"c-2",agent);!errors.Is(err,domain.ErrInvalidTransition){t.Fatalf("want invalid transition, got %v",err)}
 for _,next:=range []domain.Status{domain.Researching,domain.SolutionReview,domain.Planning,domain.AwaitingPlanApproval}{ w,err=s.Transition(ctx,w.ID,w.Version,next,"c",agent);if err!=nil{t.Fatal(err)} }
 if _,err=s.Transition(ctx,w.ID,w.Version,domain.Implementing,"c",agent);!errors.Is(err,domain.ErrHumanApprovalRequired){t.Fatalf("want human approval error, got %v",err)}
 if _,err=s.Transition(ctx,w.ID,w.Version,domain.Implementing,"c",human);err!=nil{t.Fatal(err)}
}
func TestTransitionRejectsStaleVersion(t *testing.T){ctx:=context.Background();s:=application.NewService(eventstore.NewMemoryStore());a:=domain.Actor{ID:"a",Type:"agent"};w,_:=s.Create(ctx,"wo-2","Test","","c",a);_,err:=s.Transition(ctx,w.ID,0,domain.Researching,"c",a);if !errors.Is(err,eventstore.ErrConcurrency){t.Fatalf("want conflict, got %v",err)}}
