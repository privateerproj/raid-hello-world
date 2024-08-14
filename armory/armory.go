package armory

import (
	hclog "github.com/hashicorp/go-hclog"
	"github.com/spf13/viper"

	"github.com/privateerproj/privateer-sdk/raidengine"
	"github.com/privateerproj/privateer-sdk/utils"
)

// Conforms to the Armory interface type
type SVC struct {
	Tactics map[string][]raidengine.Strike     // Required, allows you to sort which strikes are run for each control
	Log     hclog.Logger                       // Recommended, allows you to set the log level for each log message
	Results map[string]raidengine.StrikeResult // Optional, allows cross referencing between strikes
}

// Optionally, retrieve config variables using Viper.
var user string

func init() {
	user = viper.GetString("user")
}

func (a *SVC) SetLogger(loggerName string) hclog.Logger {
	a.Log = raidengine.GetLogger(loggerName, false)
	return a.Log
}

func (a *SVC) GetTactics() map[string][]raidengine.Strike {
	return a.Tactics
}

// -----
// Strike and Movements for CCC_OS_C1_TR01 
// -----

// CCC_OS_C1_TR01 conforms to the Strike function type
func (a *SVC) CCC_OS_C1_TR01() (strikeName string, result raidengine.StrikeResult) {
	// set default return values
	strikeName = "CCC_OS_C1_TR01"
	result = raidengine.StrikeResult{
		Passed:      false,
		Description: "All supported network data protocols must be running on secure channels.",
		Message:     "Strike has not yet started.", // This message will be overwritten by subsequent movements
		DocsURL:     "https://maintainer.com/docs/raids/SVC",
		ControlID:   "CCC-Taxonomy-1",
		Movements:   make(map[string]raidengine.MovementResult),
	}

	CCC_OS_C1_TR01_T01_Result := CCC_OS_C1_TR01_T01()
	result.Movements["CCC_OS_C1_TR01_T01"] = CCC_OS_C1_TR01_T01_Result
	if !CCC_OS_C1_TR01_T01_Result.Passed {
		result.Message = CCC_OS_C1_TR01_T01_Result.Message
		return
	}

	// TODO: Additional movement calls go here

	return
}

func CCC_OS_C1_TR01_T01() (result raidengine.MovementResult) {
	result = raidengine.MovementResult{
		Description: "JokerName must be found in the runtime configuration.",
		Function:    utils.CallerPath(0),
	}
	
	// TODO: Movement logic goes here
	return
}


// -----
// Strike and Movements for CCC_OS_C1_TR02 
// -----

// CCC_OS_C1_TR02 conforms to the Strike function type
func (a *SVC) CCC_OS_C1_TR02() (strikeName string, result raidengine.StrikeResult) {
	// set default return values
	strikeName = "CCC_OS_C1_TR02"
	result = raidengine.StrikeResult{
		Passed:      false,
		Description: "All clear text channels should be disabled.",
		Message:     "Strike has not yet started.", // This message will be overwritten by subsequent movements
		DocsURL:     "https://maintainer.com/docs/raids/SVC",
		ControlID:   "CCC-Taxonomy-1",
		Movements:   make(map[string]raidengine.MovementResult),
	}

	CCC_OS_C1_TR02_T01_Result := CCC_OS_C1_TR02_T01()
	result.Movements["CCC_OS_C1_TR02_T01"] = CCC_OS_C1_TR02_T01_Result
	if !CCC_OS_C1_TR02_T01_Result.Passed {
		result.Message = CCC_OS_C1_TR02_T01_Result.Message
		return
	}

	// TODO: Additional movement calls go here

	return
}

func CCC_OS_C1_TR02_T01() (result raidengine.MovementResult) {
	result = raidengine.MovementResult{
		Description: "JokerName must be found in the runtime configuration.",
		Function:    utils.CallerPath(0),
	}
	
	// TODO: Movement logic goes here
	return
}


// -----
// Strike and Movements for CCC_OS_C1_TR03 
// -----

// CCC_OS_C1_TR03 conforms to the Strike function type
func (a *SVC) CCC_OS_C1_TR03() (strikeName string, result raidengine.StrikeResult) {
	// set default return values
	strikeName = "CCC_OS_C1_TR03"
	result = raidengine.StrikeResult{
		Passed:      false,
		Description: "The cipher suite implemented for ensuring the integrity and confidentiality of data should conform with the latest suggested cipher suites. [NIST proposed latest standard cipher suites](&lt;[#](https://csrc.nist.gov/pubs/sp/800/52/r2/final)&gt;).",
		Message:     "Strike has not yet started.", // This message will be overwritten by subsequent movements
		DocsURL:     "https://maintainer.com/docs/raids/SVC",
		ControlID:   "CCC-Taxonomy-1",
		Movements:   make(map[string]raidengine.MovementResult),
	}

	CCC_OS_C1_TR03_T01_Result := CCC_OS_C1_TR03_T01()
	result.Movements["CCC_OS_C1_TR03_T01"] = CCC_OS_C1_TR03_T01_Result
	if !CCC_OS_C1_TR03_T01_Result.Passed {
		result.Message = CCC_OS_C1_TR03_T01_Result.Message
		return
	}

	// TODO: Additional movement calls go here

	return
}

func CCC_OS_C1_TR03_T01() (result raidengine.MovementResult) {
	result = raidengine.MovementResult{
		Description: "JokerName must be found in the runtime configuration.",
		Function:    utils.CallerPath(0),
	}
	
	// TODO: Movement logic goes here
	return
}
 

// -----
// Strike and Movements for CCC_OS_C2_TR01 
// -----

// CCC_OS_C2_TR01 conforms to the Strike function type
func (a *SVC) CCC_OS_C2_TR01() (strikeName string, result raidengine.StrikeResult) {
	// set default return values
	strikeName = "CCC_OS_C2_TR01"
	result = raidengine.StrikeResult{
		Passed:      false,
		Description: "Verify that data stored in the object storage bucket is encrypted using industry-standard algorithms.",
		Message:     "Strike has not yet started.", // This message will be overwritten by subsequent movements
		DocsURL:     "https://maintainer.com/docs/raids/SVC",
		ControlID:   "CCC-Taxonomy-1",
		Movements:   make(map[string]raidengine.MovementResult),
	}

	CCC_OS_C2_TR01_T01_Result := CCC_OS_C2_TR01_T01()
	result.Movements["CCC_OS_C2_TR01_T01"] = CCC_OS_C2_TR01_T01_Result
	if !CCC_OS_C2_TR01_T01_Result.Passed {
		result.Message = CCC_OS_C2_TR01_T01_Result.Message
		return
	}

	// TODO: Additional movement calls go here

	return
}

func CCC_OS_C2_TR01_T01() (result raidengine.MovementResult) {
	result = raidengine.MovementResult{
		Description: "JokerName must be found in the runtime configuration.",
		Function:    utils.CallerPath(0),
	}
	
	// TODO: Movement logic goes here
	return
}


// -----
// Strike and Movements for CCC_OS_C2_TR02 
// -----

// CCC_OS_C2_TR02 conforms to the Strike function type
func (a *SVC) CCC_OS_C2_TR02() (strikeName string, result raidengine.StrikeResult) {
	// set default return values
	strikeName = "CCC_OS_C2_TR02"
	result = raidengine.StrikeResult{
		Passed:      false,
		Description: "Ensure that encryption keys are managed securely and rotated periodically.",
		Message:     "Strike has not yet started.", // This message will be overwritten by subsequent movements
		DocsURL:     "https://maintainer.com/docs/raids/SVC",
		ControlID:   "CCC-Taxonomy-1",
		Movements:   make(map[string]raidengine.MovementResult),
	}

	CCC_OS_C2_TR02_T01_Result := CCC_OS_C2_TR02_T01()
	result.Movements["CCC_OS_C2_TR02_T01"] = CCC_OS_C2_TR02_T01_Result
	if !CCC_OS_C2_TR02_T01_Result.Passed {
		result.Message = CCC_OS_C2_TR02_T01_Result.Message
		return
	}

	// TODO: Additional movement calls go here

	return
}

func CCC_OS_C2_TR02_T01() (result raidengine.MovementResult) {
	result = raidengine.MovementResult{
		Description: "JokerName must be found in the runtime configuration.",
		Function:    utils.CallerPath(0),
	}
	
	// TODO: Movement logic goes here
	return
}


// -----
// Strike and Movements for CCC_OS_C2_TR03 
// -----

// CCC_OS_C2_TR03 conforms to the Strike function type
func (a *SVC) CCC_OS_C2_TR03() (strikeName string, result raidengine.StrikeResult) {
	// set default return values
	strikeName = "CCC_OS_C2_TR03"
	result = raidengine.StrikeResult{
		Passed:      false,
		Description: "Confirm that decryption is only possible through authorized access mechanisms.",
		Message:     "Strike has not yet started.", // This message will be overwritten by subsequent movements
		DocsURL:     "https://maintainer.com/docs/raids/SVC",
		ControlID:   "CCC-Taxonomy-1",
		Movements:   make(map[string]raidengine.MovementResult),
	}

	CCC_OS_C2_TR03_T01_Result := CCC_OS_C2_TR03_T01()
	result.Movements["CCC_OS_C2_TR03_T01"] = CCC_OS_C2_TR03_T01_Result
	if !CCC_OS_C2_TR03_T01_Result.Passed {
		result.Message = CCC_OS_C2_TR03_T01_Result.Message
		return
	}

	// TODO: Additional movement calls go here

	return
}

func CCC_OS_C2_TR03_T01() (result raidengine.MovementResult) {
	result = raidengine.MovementResult{
		Description: "JokerName must be found in the runtime configuration.",
		Function:    utils.CallerPath(0),
	}
	
	// TODO: Movement logic goes here
	return
}
 

// -----
// Strike and Movements for CCC_OS_C3_TR01 
// -----

// CCC_OS_C3_TR01 conforms to the Strike function type
func (a *SVC) CCC_OS_C3_TR01() (strikeName string, result raidengine.StrikeResult) {
	// set default return values
	strikeName = "CCC_OS_C3_TR01"
	result = raidengine.StrikeResult{
		Passed:      false,
		Description: "Verify that MFA is enforced for all access attempts to the object storage bucket.",
		Message:     "Strike has not yet started.", // This message will be overwritten by subsequent movements
		DocsURL:     "https://maintainer.com/docs/raids/SVC",
		ControlID:   "CCC-Taxonomy-1",
		Movements:   make(map[string]raidengine.MovementResult),
	}

	CCC_OS_C3_TR01_T01_Result := CCC_OS_C3_TR01_T01()
	result.Movements["CCC_OS_C3_TR01_T01"] = CCC_OS_C3_TR01_T01_Result
	if !CCC_OS_C3_TR01_T01_Result.Passed {
		result.Message = CCC_OS_C3_TR01_T01_Result.Message
		return
	}

	// TODO: Additional movement calls go here

	return
}

func CCC_OS_C3_TR01_T01() (result raidengine.MovementResult) {
	result = raidengine.MovementResult{
		Description: "JokerName must be found in the runtime configuration.",
		Function:    utils.CallerPath(0),
	}
	
	// TODO: Movement logic goes here
	return
}


// -----
// Strike and Movements for CCC_OS_C3_TR02 
// -----

// CCC_OS_C3_TR02 conforms to the Strike function type
func (a *SVC) CCC_OS_C3_TR02() (strikeName string, result raidengine.StrikeResult) {
	// set default return values
	strikeName = "CCC_OS_C3_TR02"
	result = raidengine.StrikeResult{
		Passed:      false,
		Description: "Ensure that MFA is required for all administrative access to the storage management interface.",
		Message:     "Strike has not yet started.", // This message will be overwritten by subsequent movements
		DocsURL:     "https://maintainer.com/docs/raids/SVC",
		ControlID:   "CCC-Taxonomy-1",
		Movements:   make(map[string]raidengine.MovementResult),
	}

	CCC_OS_C3_TR02_T01_Result := CCC_OS_C3_TR02_T01()
	result.Movements["CCC_OS_C3_TR02_T01"] = CCC_OS_C3_TR02_T01_Result
	if !CCC_OS_C3_TR02_T01_Result.Passed {
		result.Message = CCC_OS_C3_TR02_T01_Result.Message
		return
	}

	// TODO: Additional movement calls go here

	return
}

func CCC_OS_C3_TR02_T01() (result raidengine.MovementResult) {
	result = raidengine.MovementResult{
		Description: "JokerName must be found in the runtime configuration.",
		Function:    utils.CallerPath(0),
	}
	
	// TODO: Movement logic goes here
	return
}


// -----
// Strike and Movements for CCC_OS_C3_TR03 
// -----

// CCC_OS_C3_TR03 conforms to the Strike function type
func (a *SVC) CCC_OS_C3_TR03() (strikeName string, result raidengine.StrikeResult) {
	// set default return values
	strikeName = "CCC_OS_C3_TR03"
	result = raidengine.StrikeResult{
		Passed:      false,
		Description: "Confirm that users are unable to access the object storage bucket without completing MFA.",
		Message:     "Strike has not yet started.", // This message will be overwritten by subsequent movements
		DocsURL:     "https://maintainer.com/docs/raids/SVC",
		ControlID:   "CCC-Taxonomy-1",
		Movements:   make(map[string]raidengine.MovementResult),
	}

	CCC_OS_C3_TR03_T01_Result := CCC_OS_C3_TR03_T01()
	result.Movements["CCC_OS_C3_TR03_T01"] = CCC_OS_C3_TR03_T01_Result
	if !CCC_OS_C3_TR03_T01_Result.Passed {
		result.Message = CCC_OS_C3_TR03_T01_Result.Message
		return
	}

	// TODO: Additional movement calls go here

	return
}

func CCC_OS_C3_TR03_T01() (result raidengine.MovementResult) {
	result = raidengine.MovementResult{
		Description: "JokerName must be found in the runtime configuration.",
		Function:    utils.CallerPath(0),
	}
	
	// TODO: Movement logic goes here
	return
}
 

// -----
// Strike and Movements for CCC_OS_C4_TR01 
// -----

// CCC_OS_C4_TR01 conforms to the Strike function type
func (a *SVC) CCC_OS_C4_TR01() (strikeName string, result raidengine.StrikeResult) {
	// set default return values
	strikeName = "CCC_OS_C4_TR01"
	result = raidengine.StrikeResult{
		Passed:      false,
		Description: "Verify that data in the object storage bucket is protected by immutability settings.",
		Message:     "Strike has not yet started.", // This message will be overwritten by subsequent movements
		DocsURL:     "https://maintainer.com/docs/raids/SVC",
		ControlID:   "CCC-Taxonomy-1",
		Movements:   make(map[string]raidengine.MovementResult),
	}

	CCC_OS_C4_TR01_T01_Result := CCC_OS_C4_TR01_T01()
	result.Movements["CCC_OS_C4_TR01_T01"] = CCC_OS_C4_TR01_T01_Result
	if !CCC_OS_C4_TR01_T01_Result.Passed {
		result.Message = CCC_OS_C4_TR01_T01_Result.Message
		return
	}

	// TODO: Additional movement calls go here

	return
}

func CCC_OS_C4_TR01_T01() (result raidengine.MovementResult) {
	result = raidengine.MovementResult{
		Description: "JokerName must be found in the runtime configuration.",
		Function:    utils.CallerPath(0),
	}
	
	// TODO: Movement logic goes here
	return
}


// -----
// Strike and Movements for CCC_OS_C4_TR02 
// -----

// CCC_OS_C4_TR02 conforms to the Strike function type
func (a *SVC) CCC_OS_C4_TR02() (strikeName string, result raidengine.StrikeResult) {
	// set default return values
	strikeName = "CCC_OS_C4_TR02"
	result = raidengine.StrikeResult{
		Passed:      false,
		Description: "Ensure that attempts to modify or delete data within the immutability period are denied.",
		Message:     "Strike has not yet started.", // This message will be overwritten by subsequent movements
		DocsURL:     "https://maintainer.com/docs/raids/SVC",
		ControlID:   "CCC-Taxonomy-1",
		Movements:   make(map[string]raidengine.MovementResult),
	}

	CCC_OS_C4_TR02_T01_Result := CCC_OS_C4_TR02_T01()
	result.Movements["CCC_OS_C4_TR02_T01"] = CCC_OS_C4_TR02_T01_Result
	if !CCC_OS_C4_TR02_T01_Result.Passed {
		result.Message = CCC_OS_C4_TR02_T01_Result.Message
		return
	}

	// TODO: Additional movement calls go here

	return
}

func CCC_OS_C4_TR02_T01() (result raidengine.MovementResult) {
	result = raidengine.MovementResult{
		Description: "JokerName must be found in the runtime configuration.",
		Function:    utils.CallerPath(0),
	}
	
	// TODO: Movement logic goes here
	return
}


// -----
// Strike and Movements for CCC_OS_C4_TR03 
// -----

// CCC_OS_C4_TR03 conforms to the Strike function type
func (a *SVC) CCC_OS_C4_TR03() (strikeName string, result raidengine.StrikeResult) {
	// set default return values
	strikeName = "CCC_OS_C4_TR03"
	result = raidengine.StrikeResult{
		Passed:      false,
		Description: "Confirm that immutable data remains unchanged throughout the defined retention period.",
		Message:     "Strike has not yet started.", // This message will be overwritten by subsequent movements
		DocsURL:     "https://maintainer.com/docs/raids/SVC",
		ControlID:   "CCC-Taxonomy-1",
		Movements:   make(map[string]raidengine.MovementResult),
	}

	CCC_OS_C4_TR03_T01_Result := CCC_OS_C4_TR03_T01()
	result.Movements["CCC_OS_C4_TR03_T01"] = CCC_OS_C4_TR03_T01_Result
	if !CCC_OS_C4_TR03_T01_Result.Passed {
		result.Message = CCC_OS_C4_TR03_T01_Result.Message
		return
	}

	// TODO: Additional movement calls go here

	return
}

func CCC_OS_C4_TR03_T01() (result raidengine.MovementResult) {
	result = raidengine.MovementResult{
		Description: "JokerName must be found in the runtime configuration.",
		Function:    utils.CallerPath(0),
	}
	
	// TODO: Movement logic goes here
	return
}
 

// -----
// Strike and Movements for CCC_OS_C5_TR01 
// -----

// CCC_OS_C5_TR01 conforms to the Strike function type
func (a *SVC) CCC_OS_C5_TR01() (strikeName string, result raidengine.StrikeResult) {
	// set default return values
	strikeName = "CCC_OS_C5_TR01"
	result = raidengine.StrikeResult{
		Passed:      false,
		Description: "Verify that all access attempts to the object storage bucket are logged.",
		Message:     "Strike has not yet started.", // This message will be overwritten by subsequent movements
		DocsURL:     "https://maintainer.com/docs/raids/SVC",
		ControlID:   "CCC-Taxonomy-1",
		Movements:   make(map[string]raidengine.MovementResult),
	}

	CCC_OS_C5_TR01_T01_Result := CCC_OS_C5_TR01_T01()
	result.Movements["CCC_OS_C5_TR01_T01"] = CCC_OS_C5_TR01_T01_Result
	if !CCC_OS_C5_TR01_T01_Result.Passed {
		result.Message = CCC_OS_C5_TR01_T01_Result.Message
		return
	}

	// TODO: Additional movement calls go here

	return
}

func CCC_OS_C5_TR01_T01() (result raidengine.MovementResult) {
	result = raidengine.MovementResult{
		Description: "JokerName must be found in the runtime configuration.",
		Function:    utils.CallerPath(0),
	}
	
	// TODO: Movement logic goes here
	return
}


// -----
// Strike and Movements for CCC_OS_C5_TR02 
// -----

// CCC_OS_C5_TR02 conforms to the Strike function type
func (a *SVC) CCC_OS_C5_TR02() (strikeName string, result raidengine.StrikeResult) {
	// set default return values
	strikeName = "CCC_OS_C5_TR02"
	result = raidengine.StrikeResult{
		Passed:      false,
		Description: "Ensure that all changes to the object storage bucket configurations are logged.",
		Message:     "Strike has not yet started.", // This message will be overwritten by subsequent movements
		DocsURL:     "https://maintainer.com/docs/raids/SVC",
		ControlID:   "CCC-Taxonomy-1",
		Movements:   make(map[string]raidengine.MovementResult),
	}

	CCC_OS_C5_TR02_T01_Result := CCC_OS_C5_TR02_T01()
	result.Movements["CCC_OS_C5_TR02_T01"] = CCC_OS_C5_TR02_T01_Result
	if !CCC_OS_C5_TR02_T01_Result.Passed {
		result.Message = CCC_OS_C5_TR02_T01_Result.Message
		return
	}

	// TODO: Additional movement calls go here

	return
}

func CCC_OS_C5_TR02_T01() (result raidengine.MovementResult) {
	result = raidengine.MovementResult{
		Description: "JokerName must be found in the runtime configuration.",
		Function:    utils.CallerPath(0),
	}
	
	// TODO: Movement logic goes here
	return
}


// -----
// Strike and Movements for CCC_OS_C5_TR03 
// -----

// CCC_OS_C5_TR03 conforms to the Strike function type
func (a *SVC) CCC_OS_C5_TR03() (strikeName string, result raidengine.StrikeResult) {
	// set default return values
	strikeName = "CCC_OS_C5_TR03"
	result = raidengine.StrikeResult{
		Passed:      false,
		Description: "Confirm that logs are protected against unauthorized access and tampering.",
		Message:     "Strike has not yet started.", // This message will be overwritten by subsequent movements
		DocsURL:     "https://maintainer.com/docs/raids/SVC",
		ControlID:   "CCC-Taxonomy-1",
		Movements:   make(map[string]raidengine.MovementResult),
	}

	CCC_OS_C5_TR03_T01_Result := CCC_OS_C5_TR03_T01()
	result.Movements["CCC_OS_C5_TR03_T01"] = CCC_OS_C5_TR03_T01_Result
	if !CCC_OS_C5_TR03_T01_Result.Passed {
		result.Message = CCC_OS_C5_TR03_T01_Result.Message
		return
	}

	// TODO: Additional movement calls go here

	return
}

func CCC_OS_C5_TR03_T01() (result raidengine.MovementResult) {
	result = raidengine.MovementResult{
		Description: "JokerName must be found in the runtime configuration.",
		Function:    utils.CallerPath(0),
	}
	
	// TODO: Movement logic goes here
	return
}
 