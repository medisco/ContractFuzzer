package oracle

type NumberDependencyOracle struct{}

func (o NumberDependencyOracle) Detect(snapshot EventsSnapshot) bool {
	return snapshot.BlockNumber && (snapshot.StorageChanged || snapshot.EtherTransfer || snapshot.SendOp)
}
