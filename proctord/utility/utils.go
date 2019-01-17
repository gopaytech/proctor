package utility

import (
	"bytes"
	"fmt"
	"strings"
)

const ClientError = "malformed request"
const NonExistentProcClientError = "proc name non existent"
const InvalidCronExpressionClientError = "Cron expression invalid"
const InvalidEmailIdClientError = "Provided invalid Email ID"
const InvalidTagError  = "Tag(s) are missing"
const DuplicateJobNameArgsClientError = "provided duplicate combination of job name and args for scheduling"
const ServerError = "Something went wrong"

const UnauthorizedErrorMissingConfig = "EMAIL_ID or ACCESS_TOKEN is not present in proctor config file."
const UnauthorizedErrorInvalidConfig = "Please check the EMAIL_ID and ACCESS_TOKEN validity in proctor config file."
const GenericListCmdError = "Error fetching list of procs. Please check configuration and network connectivity"
const GenericProcCmdError = "Error executing proc. Please check configuration and network connectivity"
const GenericDescribeCmdError = "Error fetching description of proc. Please check configuration and network connectivity"

const UnauthorizedErrorHeader = "Unauthorized Access!!!"
const GenericTimeoutErrorHeader = "Connection Timeout!!!"
const GenericNetworkErrorHeader = "Network Error!!!"
const GenericResponseErrorHeader = "Server Error!!!"

const ConfigProctorHostMissingError = "Config Error!!!\nMandatory config PROCTOR_HOST is missing in Proctor Config file."
const GenericTimeoutErrorBody = "Please check your Internet/VPN connection for connectivity to ProctorD."
const ClientOutdatedErrorMessage = "Your Proctor client is using an outdated version: %s . To continue using proctor, please upgrade to latest version."

const JobSubmissionSuccess = "success"
const JobSubmissionClientError = "client_error"
const JobSubmissionServerError = "server_error"

const JobSucceeded = "SUCCEEDED"
const JobFailed = "FAILED"
const JobWaiting = "WAITING"
const JobExecutionStatusFetchError = "JOB_EXECUTION_STATUS_FETCH_ERROR"
const NoDefinitiveJobExecutionStatusFound = "NO_DEFINITIVE_JOB_EXECUTION_STATUS_FOUND"

const JobNameContextKey = "job_name"
const UserEmailContextKey = "user_email"
const JobArgsContextKey = "job_args"
const ImageNameContextKey = "image_name"
const JobNameSubmittedForExecutionContextKey = "job_name_submitted_for_execution"
const JobSubmissionStatusContextKey = "job_sumission_status"
const JobSchedulingStatusContextKey = "job_scheduling_status"

const UserEmailHeaderKey = "Email-Id"
const AccessTokenHeaderKey = "Access-Token"
const ClientVersionHeaderKey = "Client-Version"

const WorkerEmail = "worker@proctor"

func MergeMaps(mapOne, mapTwo map[string]string) map[string]string {
	result := make(map[string]string)

	for k, v := range mapOne {
		result[k] = v
	}
	for k, v := range mapTwo {
		result[k] = v
	}
	return result
}

func MapToString(someMap map[string]string) string {
	b := new(bytes.Buffer)
	for key, value := range someMap {
		fmt.Fprintf(b, "%s=\"%s\",", key, value)
	}
	return strings.TrimRight(b.String(), ",")
}
