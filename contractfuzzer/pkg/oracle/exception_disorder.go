package oracle

type ExceptionDisorderOracle struct {
}

func (o ExceptionDisorderOracle) Detect(snapshot EventsSnapshot) bool {
	return snapshot.ExceptionDisorder
}
