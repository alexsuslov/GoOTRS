package model

import "github.com/sirupsen/logrus"

//go:generate go run gena/gena.go
//go:generate gofmt -w generated.go

type Ticket struct {
	Age                       int            `json:"Age"`
	ArchiveFlag               string         `json:"ArchiveFlag"`
	Articles                  []Article      `json:"Article"`
	ChangeBy                  ID             `json:"ChangeBy"`
	CreateBy                  ID             `json:"CreateBy"`
	Changed                   string         `json:"Changed"`
	Created                   string         `json:"Created"`
	CustomerID                *string         `json:"CustomerID"`
	CustomerUser              string         `json:"CustomerUser"`
	CustomerUserID            *string         `json:"CustomerUserID"`
	DynamicFields             []DynamicField `json:"DynamicField"`
	EscalationDestinationDate string         `json:"EscalationDestinationDate"`
	EscalationDestinationIn   string         `json:"EscalationDestinationIn"`
	EscalationDestinationTime string         `json:"EscalationDestinationTime"`
	EscalationResponseTime    string         `json:"EscalationResponseTime"`
	EscalationSolutionTime    string         `json:"EscalationSolutionTime"`
	EscalationTime            int            `json:"EscalationTime"`
	EscalationTimeWorkingTime int            `json:"EscalationTimeWorkingTime"`
	EscalationUpdateTime      string         `json:"EscalationUpdateTime"`
	GroupID                   ID             `json:"GroupID"`
	Lock                      string         `json:"Lock"`
	LockID                    ID             `json:"LockID"`
	Owner                     string         `json:"Owner"`
	OwnerID                   ID             `json:"OwnerID"`
	Priority                  string         `json:"Priority"`
	PriorityID                ID             `json:"PriorityID"`

	Queue   string `json:"Queue"`
	QueueID ID     `json:"QueueID"`

	RealTillTimeNotUsed string `json:"RealTillTimeNotUsed"`

	Responsible   string `json:"Responsible"`
	ResponsibleID ID     `json:"ResponsibleID"`

	SLA   string `json:"SLA"`
	SLAID ID     `json:"SLAID"`

	Service   string `json:"Service"`
	ServiceID ID     `json:"ServiceID"`

	SolutionTime int `json:"SolutionTime"`

	SolutionTimeDestinationDate *string `json:"SolutionTimeDestinationDate"`
	SolutionTimeDestinationTime *Time `json:"SolutionTimeDestinationTime"`
	SolutionTimeWorkingTime     int    `json:"SolutionTimeWorkingTime"`

	State   string `json:"State"`
	StateID ID     `json:"StateID"`

	TicketID     ID     `json:"TicketID"`
	TicketNumber string `json:"TicketNumber"`

	TimeUnit int    `json:"TimeUnit"`
	Title    string `json:"Title"`

	Type   string `json:"Type"`
	TypeID ID     `json:"TypeID"`

	UnlockTimeout string `json:"UnlockTimeout"`
	UntilTime     int    `json:"UntilTime"`
}

type DynamicField struct {
	Name  string      `json:"Name"`
	Value interface{} `json:"Value"`
}

type TicketUpdate struct {
	Title         string      `json:"Title"`

	QueueID       *ID          `json:"QueueID,omitempty"`
	Queue         *string      `json:"Queue,omitempty"`
	Lock          *string      `json:"Lock,omitempty"`
	StateID       *ID          `json:"StateID,omitempty"`
	Priority      *string      `json:"Priority,omitempty"`
	SLAID         *ID          `json:"SLAID,omitempty"`
	State         *string      `json:"State,omitempty"`
	Responsible   *string      `json:"Responsible,omitempty"`
	TypeID        *ID          `json:"TypeID,omitempty"`
	ServiceID     *ID          `json:"ServiceID,omitempty"`
	Owner         *string      `json:"Owner,omitempty"`
	ResponsibleID *ID          `json:"ResponsibleID,omitempty"`
	OwnerID       *ID          `json:"OwnerID,omitempty"`
	CustomerUser  *string      `json:"CustomerUser,omitempty"`
	LockID        *ID          `json:"LockID,omitempty"`
	Type          *string      `json:"Type,omitempty"`
	Service       *string      `json:"Service,omitempty"`
	SLA           *string      `json:"SLA,omitempty"`
	PriorityID    *ID          `json:"PriorityID,omitempty"`
	PendingTime   *PendingTime `json:"PendingTime,omitempty"`

	logger bool
	prev   *Ticket
}

/**
 *	System
 */
// log
func (TicketUpdate *TicketUpdate) log(FieldName string, from, to interface{}) *TicketUpdate {
	if TicketUpdate.logger {
		logrus.
			WithField("from", from).
			WithField("to", to).
			Infof("set %s", FieldName)
	}
	return TicketUpdate
}

// SetPendingTime set PendingTime
func (TicketUpdate *TicketUpdate) SetPendingTime(to *PendingTime) *TicketUpdate {
	TicketUpdate.PendingTime = to
	return TicketUpdate.log("PendingTime", "", to)

}
