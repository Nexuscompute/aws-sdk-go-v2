// Code generated by smithy-go-codegen DO NOT EDIT.

package guardduty

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/guardduty/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Describes which data sources are enabled for the member account's detector.
//
// There might be regional differences because some data sources might not be
// available in all the Amazon Web Services Regions where GuardDuty is presently
// supported. For more information, see [Regions and endpoints].
//
// [Regions and endpoints]: https://docs.aws.amazon.com/guardduty/latest/ug/guardduty_regions.html
func (c *Client) GetMemberDetectors(ctx context.Context, params *GetMemberDetectorsInput, optFns ...func(*Options)) (*GetMemberDetectorsOutput, error) {
	if params == nil {
		params = &GetMemberDetectorsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetMemberDetectors", params, optFns, c.addOperationGetMemberDetectorsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetMemberDetectorsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetMemberDetectorsInput struct {

	// The account ID of the member account.
	//
	// This member is required.
	AccountIds []string

	// The detector ID for the administrator account.
	//
	// This member is required.
	DetectorId *string

	noSmithyDocumentSerde
}

type GetMemberDetectorsOutput struct {

	// An object that describes which data sources are enabled for a member account.
	//
	// This member is required.
	MemberDataSourceConfigurations []types.MemberDataSourceConfiguration

	// A list of member account IDs that were unable to be processed along with an
	// explanation for why they were not processed.
	//
	// This member is required.
	UnprocessedAccounts []types.UnprocessedAccount

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetMemberDetectorsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpGetMemberDetectors{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpGetMemberDetectors{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "GetMemberDetectors"); err != nil {
		return fmt.Errorf("add protocol finalizers: %v", err)
	}

	if err = addlegacyEndpointContextSetter(stack, options); err != nil {
		return err
	}
	if err = addSetLoggerMiddleware(stack, options); err != nil {
		return err
	}
	if err = addClientRequestID(stack); err != nil {
		return err
	}
	if err = addComputeContentLength(stack); err != nil {
		return err
	}
	if err = addResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = addComputePayloadSHA256(stack); err != nil {
		return err
	}
	if err = addRetry(stack, options); err != nil {
		return err
	}
	if err = addRawResponseToMetadata(stack); err != nil {
		return err
	}
	if err = addRecordResponseTiming(stack); err != nil {
		return err
	}
	if err = addClientUserAgent(stack, options); err != nil {
		return err
	}
	if err = smithyhttp.AddErrorCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = addSetLegacyContextSigningOptionsMiddleware(stack); err != nil {
		return err
	}
	if err = addTimeOffsetBuild(stack, c); err != nil {
		return err
	}
	if err = addOpGetMemberDetectorsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetMemberDetectors(options.Region), middleware.Before); err != nil {
		return err
	}
	if err = addRecursionDetection(stack); err != nil {
		return err
	}
	if err = addRequestIDRetrieverMiddleware(stack); err != nil {
		return err
	}
	if err = addResponseErrorMiddleware(stack); err != nil {
		return err
	}
	if err = addRequestResponseLogging(stack, options); err != nil {
		return err
	}
	if err = addDisableHTTPSMiddleware(stack, options); err != nil {
		return err
	}
	return nil
}

func newServiceMetadataMiddleware_opGetMemberDetectors(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "GetMemberDetectors",
	}
}
