// Code generated by smithy-go-codegen DO NOT EDIT.

package kendra

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Deletes a group so that all users and sub groups that belong to the group can
// no longer access documents only available to that group.
//
// For example, after deleting the group "Summer Interns", all interns who
// belonged to that group no longer see intern-only documents in their search
// results.
//
// If you want to delete or replace users or sub groups of a group, you need to
// use the PutPrincipalMapping operation. For example, if a user in the group
// "Engineering" leaves the engineering team and another user takes their place,
// you provide an updated list of users or sub groups that belong to the
// "Engineering" group when calling PutPrincipalMapping . You can update your
// internal list of users or sub groups and input this list when calling
// PutPrincipalMapping .
//
// DeletePrincipalMapping is currently not supported in the Amazon Web Services
// GovCloud (US-West) region.
func (c *Client) DeletePrincipalMapping(ctx context.Context, params *DeletePrincipalMappingInput, optFns ...func(*Options)) (*DeletePrincipalMappingOutput, error) {
	if params == nil {
		params = &DeletePrincipalMappingInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DeletePrincipalMapping", params, optFns, c.addOperationDeletePrincipalMappingMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DeletePrincipalMappingOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DeletePrincipalMappingInput struct {

	// The identifier of the group you want to delete.
	//
	// This member is required.
	GroupId *string

	// The identifier of the index you want to delete a group from.
	//
	// This member is required.
	IndexId *string

	// The identifier of the data source you want to delete a group from.
	//
	// A group can be tied to multiple data sources. You can delete a group from
	// accessing documents in a certain data source. For example, the groups
	// "Research", "Engineering", and "Sales and Marketing" are all tied to the
	// company's documents stored in the data sources Confluence and Salesforce. You
	// want to delete "Research" and "Engineering" groups from Salesforce, so that
	// these groups cannot access customer-related documents stored in Salesforce. Only
	// "Sales and Marketing" should access documents in the Salesforce data source.
	DataSourceId *string

	// The timestamp identifier you specify to ensure Amazon Kendra does not override
	// the latest DELETE action with previous actions. The highest number ID, which is
	// the ordering ID, is the latest action you want to process and apply on top of
	// other actions with lower number IDs. This prevents previous actions with lower
	// number IDs from possibly overriding the latest action.
	//
	// The ordering ID can be the Unix time of the last update you made to a group
	// members list. You would then provide this list when calling PutPrincipalMapping
	// . This ensures your DELETE action for that updated group with the latest
	// members list doesn't get overwritten by earlier DELETE actions for the same
	// group which are yet to be processed.
	//
	// The default ordering ID is the current Unix time in milliseconds that the
	// action was received by Amazon Kendra.
	OrderingId *int64

	noSmithyDocumentSerde
}

type DeletePrincipalMappingOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDeletePrincipalMappingMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDeletePrincipalMapping{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDeletePrincipalMapping{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "DeletePrincipalMapping"); err != nil {
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
	if err = addOpDeletePrincipalMappingValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDeletePrincipalMapping(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDeletePrincipalMapping(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "DeletePrincipalMapping",
	}
}
