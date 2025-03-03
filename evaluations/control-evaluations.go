package evaluations

import (
	"fmt"

	"github.com/revanite-io/sci/pkg/layer4"
)


// 
// Data Protection Control Family

func CCC_C01() (evaluation *layer4.ControlEvaluation) {
	evaluation = &layer4.ControlEvaluation{
		Control_Id:        "CCC.C01",
		Remediation_Guide: "",
	}
	
	evaluation.AddAssessment(
		"CCC.C01.TR01",
		"When a port is exposed for non-SSH network traffic, all traffic MUST include a TLS handshake AND be encrypted using TLS 1.2 or higher.",
		[]string{ 
			"tlp_clear",
			"tlp_green",
			"tlp_amber",
			"tlp_red",
		},
		[]layer4.AssessmentStep{
			reusable_step_example,
		},
	)
	
	evaluation.AddAssessment(
		"CCC.C01.TR02",
		"When a port is exposed for SSH network traffic, all traffic MUST include a SSH handshake AND be encrypted using SSHv2 or higher.",
		[]string{ 
			"tlp_clear",
			"tlp_green",
			"tlp_amber",
			"tlp_red",
		},
		[]layer4.AssessmentStep{
			reusable_step_example,
		},
	)
	
	return
}

func CCC_C06() (evaluation *layer4.ControlEvaluation) {
	evaluation = &layer4.ControlEvaluation{
		Control_Id:        "CCC.C06",
		Remediation_Guide: "",
	}
	
	evaluation.AddAssessment(
		"CCC.C06.TR01",
		"When a deployment request is made, the service MUST validate that the deployment region is not to a restricted or regions or availability zones.",
		[]string{ 
			"tlp_clear",
			"tlp_green",
			"tlp_amber",
			"tlp_red",
		},
		[]layer4.AssessmentStep{
			reusable_step_example,
		},
	)
	
	evaluation.AddAssessment(
		"CCC.C06.TR02",
		"When a deployment request is made, the service MUST validate that replication of data, backups, and disaster recovery operations will not occur in restricted regions or availability zones.",
		[]string{ 
			"tlp_clear",
			"tlp_green",
			"tlp_amber",
			"tlp_red",
		},
		[]layer4.AssessmentStep{
			reusable_step_example,
		},
	)
	
	return
}

func CCC_C08() (evaluation *layer4.ControlEvaluation) {
	evaluation = &layer4.ControlEvaluation{
		Control_Id:        "CCC.C08",
		Remediation_Guide: "",
	}
	
	evaluation.AddAssessment(
		"CCC.C08.TR01",
		"When data is stored, the service MUST ensure that data is replicated across multiple availability zones or regions.",
		[]string{ 
			"tlp_green",
			"tlp_amber",
			"tlp_red",
		},
		[]layer4.AssessmentStep{
			reusable_step_example,
		},
	)
	
	evaluation.AddAssessment(
		"CCC.C08.TR02",
		"When data is replicated across multiple zones or regions, the service MUST be able to verify the replication state, including the replication locations and data synchronization status.",
		[]string{ 
			"tlp_green",
			"tlp_amber",
			"tlp_red",
		},
		[]layer4.AssessmentStep{
			reusable_step_example,
		},
	)
	
	return
}

func CCC_C09() (evaluation *layer4.ControlEvaluation) {
	evaluation = &layer4.ControlEvaluation{
		Control_Id:        "CCC.C09",
		Remediation_Guide: "",
	}
	
	evaluation.AddAssessment(
		"CCC.C09.TR01",
		"When access logs are stored, the service MUST ensure that access logs cannot be accessed without proper authorization.",
		[]string{ 
			"tlp_amber",
			"tlp_red",
			"tlp_green",
			"tlp_clear",
		},
		[]layer4.AssessmentStep{
			reusable_step_example,
		},
	)
	
	evaluation.AddAssessment(
		"CCC.C09.TR02",
		"When access logs are stored, the service MUST ensure that access logs cannot be modified without proper authorization.",
		[]string{ 
			"tlp_amber",
			"tlp_red",
			"tlp_green",
			"tlp_clear",
		},
		[]layer4.AssessmentStep{
			reusable_step_example,
		},
	)
	
	evaluation.AddAssessment(
		"CCC.C09.TR03",
		"When access logs are stored, the service MUST ensure that access logs cannot be deleted without proper authorization.",
		[]string{ 
			"tlp_amber",
			"tlp_red",
			"tlp_green",
			"tlp_clear",
		},
		[]layer4.AssessmentStep{
			reusable_step_example,
		},
	)
	
	return
}

func CCC_C10() (evaluation *layer4.ControlEvaluation) {
	evaluation = &layer4.ControlEvaluation{
		Control_Id:        "CCC.C10",
		Remediation_Guide: "",
	}
	
	evaluation.AddAssessment(
		"CCC.C10.TR01",
		"When data is replicated, the service MUST ensure that replication is restricted to explicitly trusted destinations.",
		[]string{ 
			"tlp_green",
			"tlp_amber",
			"tlp_red",
		},
		[]layer4.AssessmentStep{
			reusable_step_example,
		},
	)
	
	return
}


// TODO: This is only for reference, and should be deleted
type PayloadTypeExample struct {
	Organization struct {
		RequiresTwoFactorAuthentication bool `json:"requiresTwoFactorAuthentication"`
	} `json:"organization"`
}

// TODO: This is only for reference, and should be deleted
func reusable_step_example(payloadData interface{}, changes map[string]*layer4.Change) (result layer4.Result, message string) {
	payload, ok := payloadData.(*CustomDataObject)
	if !ok {
		return layer4.Unknown, fmt.Sprintf("Malformed assessment: expected payload type %T, got %T (%v)", &CustomDataObject{}, payloadData, payloadData)
	}
	if payload.yourData != "" {
		return layer4.Passed, fmt.Sprint("Your data ", payload.yourData)
	}
	return layer4.Unknown, "Not implemented"
}
