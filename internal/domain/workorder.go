package domain

import (
	"errors"
	"fmt"
)

type Status string

const (
	Draft                   Status = "draft"
	Researching             Status = "researching"
	SolutionReview          Status = "solution_review"
	Planning                Status = "planning"
	AwaitingPlanApproval    Status = "awaiting_plan_approval"
	Implementing            Status = "implementing"
	Validating              Status = "validating"
	IndependentReview       Status = "independent_review"
	AwaitingPublishApproval Status = "awaiting_publish_approval"
	Completed               Status = "completed"
)

var ErrInvalidTransition = errors.New("invalid work-order lifecycle transition")
var ErrHumanApprovalRequired = errors.New("human approval required")

type Actor struct {
	ID   string `json:"id"`
	Type string `json:"type"`
}

func (a Actor) IsHuman() bool { return a.Type == "human" }

type WorkOrder struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Status      Status `json:"status"`
	Version     int    `json:"version"`
}

func NewWorkOrder(id, title, description string) (*WorkOrder, error) {
	if id == "" || title == "" { return nil, errors.New("id and title are required") }
	return &WorkOrder{ID:id, Title:title, Description:description, Status:Draft}, nil
}

func (w *WorkOrder) MoveTo(next Status, actor Actor) error {
	if requiresHumanApproval(w.Status, next) && !actor.IsHuman() { return ErrHumanApprovalRequired }
	if !allowed(w.Status, next) { return fmt.Errorf("%w: %s -> %s", ErrInvalidTransition, w.Status, next) }
	w.Status = next
	return nil
}

func requiresHumanApproval(from, to Status) bool {
	return (from == AwaitingPlanApproval && to == Implementing) || (from == AwaitingPublishApproval && to == Completed)
}

func allowed(from, to Status) bool {
	return map[Status]Status{
		Draft: Researching, Researching: SolutionReview, SolutionReview: Planning,
		Planning: AwaitingPlanApproval, AwaitingPlanApproval: Implementing,
		Implementing: Validating, Validating: IndependentReview,
		IndependentReview: AwaitingPublishApproval, AwaitingPublishApproval: Completed,
	}[from] == to
}
