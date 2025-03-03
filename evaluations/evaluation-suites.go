package evaluations

import "github.com/revanite-io/sci/pkg/layer4"

var (
	FINOS_CCC = []*layer4.ControlEvaluation{
		CCC_C01(),
		CCC_C06(),
		CCC_C08(),
		CCC_C09(),
		CCC_C10(),
		
	}
)
