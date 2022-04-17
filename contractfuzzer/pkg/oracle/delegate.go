package oracle

type DelegateOracle struct{}

func (o DelegateOracle) Detect(snapshot EventsSnapshot) bool {
	return snapshot.Delegate
}
