package oracle

import "errors"

type Oracle interface {
	Detect(snapshot EventsSnapshot) bool
}

var ErrOracleDoesntExist = errors.New("oracle doesn't exist")

func GetOracles(oracleNames []string) []Oracle {
	oracles := make([]Oracle, len(oracleNames))

	for _, oracleName := range oracleNames {
		oracle, err := GetOracleFromName(oracleName)
		if err != nil {
			continue
		}
		oracles = append(oracles, oracle)
	}
	return oracles
}

func GetOracleFromName(name string) (Oracle, error) {
	switch name {
	case "delegate":
		return DelegateOracle{}, nil
	case "exception-disorder":
		return ExceptionDisorderOracle{}, nil
	case "gasless-send":
		return GaslessSendOracle{}, nil
	case "number-dependency":
		return NumberDependencyOracle{}, nil
	case "reentrancy":
		return ReentrancyOracle{}, nil
	case "timestamp-dependency":
		return TimestampDependencyOracle{}, nil
	}
	return nil, ErrOracleDoesntExist
}
