// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

// Package cloudwatchiface provides an interface to enable mocking the Amazon CloudWatch service client
// for testing your code.
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters.
package cloudwatchiface

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/cloudwatch"
)

// ClientAPI provides an interface to enable mocking the
// cloudwatch.Client methods. This make unit testing your code that
// calls out to the SDK's service client's calls easier.
//
// The best way to use this interface is so the SDK's service client's calls
// can be stubbed out for unit testing your code with the SDK without needing
// to inject custom request handlers into the SDK's request pipeline.
//
//    // myFunc uses an SDK service client to make a request to
//    // CloudWatch.
//    func myFunc(svc cloudwatchiface.ClientAPI) bool {
//        // Make svc.DeleteAlarms request
//    }
//
//    func main() {
//        cfg, err := external.LoadDefaultAWSConfig()
//        if err != nil {
//            panic("failed to load config, " + err.Error())
//        }
//
//        svc := cloudwatch.New(cfg)
//
//        myFunc(svc)
//    }
//
// In your _test.go file:
//
//    // Define a mock struct to be used in your unit tests of myFunc.
//    type mockClientClient struct {
//        cloudwatchiface.ClientPI
//    }
//    func (m *mockClientClient) DeleteAlarms(input *cloudwatch.DeleteAlarmsInput) (*cloudwatch.DeleteAlarmsOutput, error) {
//        // mock response/functionality
//    }
//
//    func TestMyFunc(t *testing.T) {
//        // Setup Test
//        mockSvc := &mockClientClient{}
//
//        myfunc(mockSvc)
//
//        // Verify myFunc's functionality
//    }
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters. Its suggested to use the pattern above for testing, or using
// tooling to generate mocks to satisfy the interfaces.
type ClientAPI interface {
	DeleteAlarmsRequest(*cloudwatch.DeleteAlarmsInput) cloudwatch.DeleteAlarmsRequest

	DeleteAnomalyDetectorRequest(*cloudwatch.DeleteAnomalyDetectorInput) cloudwatch.DeleteAnomalyDetectorRequest

	DeleteDashboardsRequest(*cloudwatch.DeleteDashboardsInput) cloudwatch.DeleteDashboardsRequest

	DescribeAlarmHistoryRequest(*cloudwatch.DescribeAlarmHistoryInput) cloudwatch.DescribeAlarmHistoryRequest

	DescribeAlarmsRequest(*cloudwatch.DescribeAlarmsInput) cloudwatch.DescribeAlarmsRequest

	DescribeAlarmsForMetricRequest(*cloudwatch.DescribeAlarmsForMetricInput) cloudwatch.DescribeAlarmsForMetricRequest

	DescribeAnomalyDetectorsRequest(*cloudwatch.DescribeAnomalyDetectorsInput) cloudwatch.DescribeAnomalyDetectorsRequest

	DisableAlarmActionsRequest(*cloudwatch.DisableAlarmActionsInput) cloudwatch.DisableAlarmActionsRequest

	EnableAlarmActionsRequest(*cloudwatch.EnableAlarmActionsInput) cloudwatch.EnableAlarmActionsRequest

	GetDashboardRequest(*cloudwatch.GetDashboardInput) cloudwatch.GetDashboardRequest

	GetMetricDataRequest(*cloudwatch.GetMetricDataInput) cloudwatch.GetMetricDataRequest

	GetMetricStatisticsRequest(*cloudwatch.GetMetricStatisticsInput) cloudwatch.GetMetricStatisticsRequest

	GetMetricWidgetImageRequest(*cloudwatch.GetMetricWidgetImageInput) cloudwatch.GetMetricWidgetImageRequest

	ListDashboardsRequest(*cloudwatch.ListDashboardsInput) cloudwatch.ListDashboardsRequest

	ListMetricsRequest(*cloudwatch.ListMetricsInput) cloudwatch.ListMetricsRequest

	ListTagsForResourceRequest(*cloudwatch.ListTagsForResourceInput) cloudwatch.ListTagsForResourceRequest

	PutAnomalyDetectorRequest(*cloudwatch.PutAnomalyDetectorInput) cloudwatch.PutAnomalyDetectorRequest

	PutDashboardRequest(*cloudwatch.PutDashboardInput) cloudwatch.PutDashboardRequest

	PutMetricAlarmRequest(*cloudwatch.PutMetricAlarmInput) cloudwatch.PutMetricAlarmRequest

	PutMetricDataRequest(*cloudwatch.PutMetricDataInput) cloudwatch.PutMetricDataRequest

	SetAlarmStateRequest(*cloudwatch.SetAlarmStateInput) cloudwatch.SetAlarmStateRequest

	TagResourceRequest(*cloudwatch.TagResourceInput) cloudwatch.TagResourceRequest

	UntagResourceRequest(*cloudwatch.UntagResourceInput) cloudwatch.UntagResourceRequest

	WaitUntilAlarmExists(context.Context, *cloudwatch.DescribeAlarmsInput, ...aws.WaiterOption) error
}

var _ ClientAPI = (*cloudwatch.Client)(nil)