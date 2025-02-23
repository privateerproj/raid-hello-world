package armory

import (
	"github.com/privateerproj/privateer-sdk/pluginkit"
)

var (
	Armory   = pluginkit.Armory{
		TestSuites: map[string][]pluginkit.TestSet{
		
			"tlp_amber": {
				CCC_C01_TR01,
				CCC_C01_TR02,
				CCC_C06_TR01,
				CCC_C06_TR02,
				CCC_C08_TR01,
				CCC_C08_TR02,
				CCC_C09_TR01,
				CCC_C09_TR02,
				CCC_C09_TR03,
				CCC_C10_TR01,
			},
			"tlp_clear": {
				CCC_C01_TR01,
				CCC_C01_TR02,
				CCC_C06_TR01,
				CCC_C06_TR02,
				CCC_C09_TR01,
				CCC_C09_TR02,
				CCC_C09_TR03,
			},
			"tlp_green": {
				CCC_C01_TR01,
				CCC_C01_TR02,
				CCC_C06_TR01,
				CCC_C06_TR02,
				CCC_C08_TR01,
				CCC_C08_TR02,
				CCC_C09_TR01,
				CCC_C09_TR02,
				CCC_C09_TR03,
				CCC_C10_TR01,
			},
			"tlp_red": {
				CCC_C01_TR01,
				CCC_C01_TR02,
				CCC_C06_TR01,
				CCC_C06_TR02,
				CCC_C08_TR01,
				CCC_C08_TR02,
				CCC_C09_TR01,
				CCC_C09_TR02,
				CCC_C09_TR03,
				CCC_C10_TR01,
			},
		},
	}
)
