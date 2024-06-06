// Code generated by smithy-go-codegen DO NOT EDIT.

package glue

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/glue/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates a data quality ruleset with DQDL rules applied to a specified Glue
// table.
//
// You create the ruleset using the Data Quality Definition Language (DQDL). For
// more information, see the Glue developer guide.
func (c *Client) CreateDataQualityRuleset(ctx context.Context, params *CreateDataQualityRulesetInput, optFns ...func(*Options)) (*CreateDataQualityRulesetOutput, error) {
	if params == nil {
		params = &CreateDataQualityRulesetInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateDataQualityRuleset", params, optFns, c.addOperationCreateDataQualityRulesetMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateDataQualityRulesetOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateDataQualityRulesetInput struct {

	// A unique name for the data quality ruleset.
	//
	// This member is required.
	Name *string

	// A Data Quality Definition Language (DQDL) ruleset. For more information, see
	// the Glue developer guide.
	//
	// This member is required.
	Ruleset *string

	// Used for idempotency and is recommended to be set to a random ID (such as a
	// UUID) to avoid creating or starting multiple instances of the same resource.
	ClientToken *string

	// A description of the data quality ruleset.
	Description *string

	// A list of tags applied to the data quality ruleset.
	Tags map[string]string

	// A target table associated with the data quality ruleset.
	TargetTable *types.DataQualityTargetTable

	noSmithyDocumentSerde
}

type CreateDataQualityRulesetOutput struct {

	// A unique name for the data quality ruleset.
	Name *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateDataQualityRulesetMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpCreateDataQualityRuleset{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpCreateDataQualityRuleset{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "CreateDataQualityRuleset"); err != nil {
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
	if err = addOpCreateDataQualityRulesetValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateDataQualityRuleset(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateDataQualityRuleset(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "CreateDataQualityRuleset",
	}
}
