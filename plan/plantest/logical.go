package plantest

import (
	"github.com/influxdata/flux/plan"
	uuid "github.com/satori/go.uuid"
)

func RandomProcedureID() plan.ProcedureID {
	return plan.ProcedureID(uuid.Must(uuid.NewV4()))
}
