package native

import (
	"github.com/infinitete/ontology-tool/methods/smartcontract/native/governance"
)

func RegisterNative() {
	governance.RegisterGovernance()
}
