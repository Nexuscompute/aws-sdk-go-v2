// Code generated by smithy-go-codegen DO NOT EDIT.

package mturk

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/mturk/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

//	The CreateHITType operation creates a new HIT type. This operation allows you
//
// to define a standard set of HIT properties to use when creating HITs. If you
// register a HIT type with values that match an existing HIT type, the HIT type ID
// of the existing type will be returned.
func (c *Client) CreateHITType(ctx context.Context, params *CreateHITTypeInput, optFns ...func(*Options)) (*CreateHITTypeOutput, error) {
	if params == nil {
		params = &CreateHITTypeInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateHITType", params, optFns, c.addOperationCreateHITTypeMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateHITTypeOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateHITTypeInput struct {

	//  The amount of time, in seconds, that a Worker has to complete the HIT after
	// accepting it. If a Worker does not complete the assignment within the specified
	// duration, the assignment is considered abandoned. If the HIT is still active
	// (that is, its lifetime has not elapsed), the assignment becomes available for
	// other users to find and accept.
	//
	// This member is required.
	AssignmentDurationInSeconds *int64

	//  A general description of the HIT. A description includes detailed information
	// about the kind of task the HIT contains. On the Amazon Mechanical Turk web site,
	// the HIT description appears in the expanded view of search results, and in the
	// HIT and assignment screens. A good description gives the user enough information
	// to evaluate the HIT before accepting it.
	//
	// This member is required.
	Description *string

	//  The amount of money the Requester will pay a Worker for successfully
	// completing the HIT.
	//
	// This member is required.
	Reward *string

	//  The title of the HIT. A title should be short and descriptive about the kind
	// of task the HIT contains. On the Amazon Mechanical Turk web site, the HIT title
	// appears in search results, and everywhere the HIT is mentioned.
	//
	// This member is required.
	Title *string

	//  The number of seconds after an assignment for the HIT has been submitted,
	// after which the assignment is considered Approved automatically unless the
	// Requester explicitly rejects it.
	AutoApprovalDelayInSeconds *int64

	//  One or more words or phrases that describe the HIT, separated by commas. These
	// words are used in searches to find HITs.
	Keywords *string

	//  Conditions that a Worker's Qualifications must meet in order to accept the
	// HIT. A HIT can have between zero and ten Qualification requirements. All
	// requirements must be met in order for a Worker to accept the HIT. Additionally,
	// other actions can be restricted using the ActionsGuarded field on each
	// QualificationRequirement structure.
	QualificationRequirements []types.QualificationRequirement

	noSmithyDocumentSerde
}

type CreateHITTypeOutput struct {

	//  The ID of the newly registered HIT type.
	HITTypeId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateHITTypeMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpCreateHITType{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpCreateHITType{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "CreateHITType"); err != nil {
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
	if err = addSpanRetryLoop(stack, options); err != nil {
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
	if err = addUserAgentRetryMode(stack, options); err != nil {
		return err
	}
	if err = addOpCreateHITTypeValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateHITType(options.Region), middleware.Before); err != nil {
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
	if err = addSpanInitializeStart(stack); err != nil {
		return err
	}
	if err = addSpanInitializeEnd(stack); err != nil {
		return err
	}
	if err = addSpanBuildRequestStart(stack); err != nil {
		return err
	}
	if err = addSpanBuildRequestEnd(stack); err != nil {
		return err
	}
	return nil
}

func newServiceMetadataMiddleware_opCreateHITType(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "CreateHITType",
	}
}
