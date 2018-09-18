package plantest

import (
	"github.com/influxdata/flux/plan"
	uuid "github.com/satori/go.uuid"
)

func RandomProcedureID() plan.ProcedureID {
	id, err := uuid.NewV4()
	if err != nil {
		panic(err)
	}
	return plan.ProcedureID(id)
}
