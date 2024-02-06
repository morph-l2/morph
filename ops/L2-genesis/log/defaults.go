package log

import (
	"os"

	"github.com/scroll-tech/go-ethereum/log"
)

func SetupDefaults() {
	log.Root().SetHandler(
		log.LvlFilterHandler(
			log.LvlInfo,
			log.StreamHandler(os.Stdout, log.LogfmtFormat()),
		),
	)
}
