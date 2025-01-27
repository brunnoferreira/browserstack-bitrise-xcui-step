package main

const (
	POLLING_INTERVAL_IN_MS             = 30000 // 30 secs
	BROWSERSTACK_DOMAIN                = "https://api-cloud.browserstack.com"
	APP_UPLOAD_ENDPOINT                = "/app-automate/xcuitest/v2/app"
	TEST_SUITE_UPLOAD_ENDPOINT         = "/app-automate/xcuitest/v2/test-suite"
	APP_AUTOMATE_BUILD_ENDPOINT        = "/app-automate/xcuitest/v2/build"
	APP_AUTOMATE_BUILD_STATUS_ENDPOINT = "/app-automate/xcuitest/v2/builds/"
	APP_AUTOMATE_BUILD_DASHBOARD_URL   = "https://app-automate.browserstack.com/dashboard/v2/builds/"
	TEST_RUNNER_RELATIVE_PATH_BITRISE  = "/Debug-iphoneos/AirwallexUITests-Runner.app"
	TEST_APP_RELATIVE_PATH_BITRISE     = "/Debug-iphoneos/Airwallex.app"
	TEST_RUNNER_ZIP_FILE_NAME          = "test_suite.zip"
	TEST_APP_ZIP_FILE_NAME             = "awx.ipa"

	SAMPLE_APP        = "bs://b91841adbf33515fef7a1cca869a9526a86f9a0e"
	SAMPLE_TEST_SUITE = "bs://535a0932c8a785384b8470ec6166e093cd3b2c5f"
	SAMPLE_BUILD_ID   = "56fec97937b22c785a6c5e08c13f629d505f5cd9"

	UPLOAD_APP_ERROR         = "Failed to upload app on BrowserStack, error : %s"
	FILE_ERROR_1             = "File error 1"
	FILE_ERROR_2             = "File error 2"
	FILE_ERROR_3             = "File error 3"
	FILE_ERROR_4             = "File error 4"
	FILE_ERROR_5             = "File error 5"
	FILE_ERROR_6             = "File error 6"
	FILE_ERROR_7             = "File error 7"
	FILE_ERROR_8             = "File error 8"
	INVALID_FILE_TYPE_ERROR  = "Failed to upload test suite on BrowserStack, error: invalid file type"
	BUILD_FAILED_ERROR       = "Failed to execute build on BrowserStack, error: %s"
	FETCH_BUILD_STATUS_ERROR = "Failed to fetch test results, error: %s"
	HTTP_ERROR               = "Something went wrong while processing your request, error: %s"
	APP_APP_NOT_FOUND_1      = "App Failure 1"
	APP_APP_NOT_FOUND_2      = "App Failure 2"
	RUNNER_APP_NOT_FOUND_1   = "Failure 1"
	RUNNER_APP_NOT_FOUND_2   = "Failure 2"
	RUNNER_APP_NOT_FOUND_3   = "Failure 3"
	RUNNER_APP_NOT_FOUND_4   = "Failure 4"
	IPA_NOT_FOUND            = "app_ipa_path: couldn’t find the iOS app (.ipa file). Please add the $BITRISE_IPA_PATH from Xcode Archive & Export for iOS step or the absolute path of iOS app (.ipa file)"
	FILE_ZIP_ERROR           = "Something went wrong while processing the test-suite, error: %s"
	APP_COPY_ERROR           = "App copy, error: %s"
	APP_DIR_ERROR            = "App dir, error: %s"
	APP_ZIP_ERROR            = "App zip, error: %s"
)
