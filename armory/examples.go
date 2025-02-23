package armory

import (
	"github.com/privateerproj/privateer-sdk/pluginkit"
	"github.com/privateerproj/privateer-sdk/utils"
)

//
// !
// !!
// !!!
//
// This file is for reference purposes only
// These are not customized or generated for your use case
// Delete this as soon as you start adding your own changes
//
// !!!
// !!
// !
//

var globalObject interface{}

// Example of a testSet that calls an invasive and non-invasive test.
// Any number or combination of tests can be called
func ExampleTestSet01() (testSetName string, result pluginkit.TestSetResult) {
	// set default return values
	testSetName = "Example_TestSet_01"
	result = pluginkit.TestSetResult{
		Description: "The service enforces the use of secure transport protocols for all network communications (e.g., TLS 1.2 or higher).",
		Message:     "TestSet has not yet started.",            // This message will be overwritten by subsequent tests
		DocsURL:     "https://maintainer.com/docs/plugins/DEV", // This is an optional link to documentation that will help users better understand the testSet
		ControlID:   "CCC.C01",                                 // This is the control ID that the testSet is testing against
		Tests:       make(map[string]pluginkit.TestResult),     // This map will be populated with the results of each test
		Passed:      false,                                     // This will be updated to true if a test passes, and back to false if a test fails
	}

	result.ExecuteTest(ExampleTest0101)

	// if a test relies on another test to pass, add this type of condition
	if result.Tests["ExampleTest0101"].Passed {
		// if a test could potentially cause harm to the target env, flag it as invasive like this
		result.ExecuteInvasiveTest(ExampleTest0102)
	}

	return
}

// ExampleTest0101 does not apply a change to the system
func ExampleTest0101() (testResult pluginkit.TestResult) {
	// Pretend we're making some API call or other logic to determine if the test is applicable
	customLogicResults := true

	testResult = pluginkit.TestResult{
		Description: "Making an API call to see if HTTPS is enforced.",
		Function:    utils.CallerPath(0), // This allows interested users to jump directly to the code that is executing this test
		Passed:      customLogicResults,
	}
	return
}

// ExampleTest0102 applies an invasive change to the system. Not all changes are invasive, but this one is.
// Use ExecuteInvasiveTest() to ensure it is run only when the user has opted in to potentially destructive changes.
func ExampleTest0102() (testResult pluginkit.TestResult) {
	// The functions here can be defined whereever you like
	// If you have a lot of changes or plan to reuse them, you may want to put them in a separate file
	change1 := pluginkit.NewChange(
		"targetName",
		"This change should create a new storage object", // For logging purposes. This will be overwritten by the result of a successful apply function.
		applyChange,
		revertChange,
	)

	// Any intended changes should be applied before returning the test result
	change1.Apply()

	// A future release may have better object handling for objects returned by the change
	// For now, toss it onto a global variable if you need to access it later
	globalObject = change1.TargetObject

	// If the change is not needed for subsequent tests, revert it now
	// A future release will use this logic to multi-thread the revert process
	// Any changes that are not reverted within the test will be reverted together at the end of the testSet
	change1.Revert()

	// Note that we are not setting Passed to true or false. That will be determined by ExecuteTest() or ExecuteInvasiveTest()
	testResult = pluginkit.TestResult{
		Description: "Making an API call to see if HTTPS is enforced.",
		Function:    utils.CallerPath(0), // This allows interested users to jump directly to the code that is executing this test
		Changes: map[string]*pluginkit.Change{
			"TestChange1": change1,
		},
	}
	return
}

// Mock function to simulate applying a change
func applyChange() (modifiedObject interface{}, err error) {
	// Replace with actual logic
	return nil, nil
}

// Mock function to simulate undoing a change
func revertChange() error {
	// Replace with actual logic
	return nil // Return an error here to simulate failure
}
