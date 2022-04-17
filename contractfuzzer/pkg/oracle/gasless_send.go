package oracle

type GaslessSendOracle struct{}

func (o GaslessSendOracle) Detect(snapshot EventsSnapshot) bool {
	return snapshot.GaslessSend
}
