package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"

	"github.com/gongbell/contractfuzzer/fuzz"
	"github.com/gongbell/contractfuzzer/server"
	"go.uber.org/zap"
)

var (
	abi_dir       = flag.String("abi_dir", "/verified_contract_abis", "input abi-dir")
	out_dir       = flag.String("out_dir", "/verified_contract_abis_fuzz", "input out-dir")
	contract_list = flag.String("contract_list", "/list/config/contracts.list", "specify contract list for fuzzing input")
	addr_seeds    = flag.String("addr_seeds", "/list/config/addressSeed.json", "specify address seedfile")
	int_seeds     = flag.String("int_seeds", "/list/config/intSeed.json", "specify int seedfile")
	uint_seeds    = flag.String("uint_seeds", "/list/config/uintSeed.json", "specify uint seedfile")
	string_seeds  = flag.String("string_seeds", "/list/config/stringSeed.json", "specify string seedfile")
	byte_seeds    = flag.String("byte_seeds", "/list/config/byteSeed.json", "specify bytes seedfile")
	bytes_seeds   = flag.String("bytes_seeds", "/list/config/bytesSeed.json", "specify bytes seedfile")
	fuzz_scale    = flag.Int("fuzz_scale", 5, "specify fuzz scale for each input param")
	input_scale   = flag.Int("input_scale", 8, "specify scale for fun")
	fstart        = flag.Int("fstart", 2, "specify fuzz scale for each input param")
	fend          = flag.Int("fend", 2, "specify fuzz scale for each input param")
	addr_map      = flag.String("addr_map", "/list/config/addrmap.csv", "set addr_map")
	abi_sigs_dir  = flag.String("abi_sigs_dir", "", "set abi_sigs_dir")
	bin_sigs_dir  = flag.String("bin_sigs_dir", "", "set bin_sigs_dir")
	listen_port   = flag.String("listen_port", "8888", "set listen_port")
	tester_port   = flag.String("tester_port", "http://localhost:6666", "set tester_port")
	reporter      = flag.String("reporter", "/reporter", "specifiy results records direcotry")
)

func main() {
	flag.Parse()

	logger, err := initLogger()
	if err != nil {
		log.Printf("Error while loading zap logger: %s", err)
		panic(err)
	}
	defer logger.Sync()

	if err := fuzz.Init(*contract_list, *addr_seeds, *int_seeds, *uint_seeds, *string_seeds, *byte_seeds, *bytes_seeds, *fuzz_scale, *input_scale, *fstart, *fend, *addr_map, *abi_sigs_dir, *bin_sigs_dir, *listen_port, *tester_port); err != nil {
		logger.Panic(fmt.Sprintf("Error while initializing fuzzer: %s\n", err))
		panic(err)
	}

	s := new(server.DefaultServer).Init(logger, *abi_dir, *out_dir, *addr_map, *reporter)
	if err = s.Run(); err != nil {
		logger.Panic(fmt.Sprintf("Error while starting the HTTP server: %s\n", err))
		panic(err)
	}
	<-fuzz.G_finish
}

func initLogger() (*zap.Logger, error) {
	rawJSON := []byte(`{
		"level": "debug",
		"encoding": "json",
		"outputPaths": ["stdout", "/tmp/logs"],
		"errorOutputPaths": ["stderr"],
		"encoderConfig": {
			"messageKey": "message",
			"levelKey": "level",
			"levelEncoder": "lowercase"
		}
	}`)

	var cfg zap.Config
	if err := json.Unmarshal(rawJSON, &cfg); err != nil {
		return nil, err
	}
	l, err := cfg.Build()
	if err != nil {
		return nil, err
	}
	return l, nil
}
