// Code generated by smithy-go-codegen DO NOT EDIT.

package partnercentralselling

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/partnercentralselling/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Fetches the Opportunity record from Partner Central by a given Identifier .
//
// Use the ListOpportunities action or the event notification (from Amazon
// EventBridge) to obtain this identifier.
func (c *Client) GetOpportunity(ctx context.Context, params *GetOpportunityInput, optFns ...func(*Options)) (*GetOpportunityOutput, error) {
	if params == nil {
		params = &GetOpportunityInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetOpportunity", params, optFns, c.addOperationGetOpportunityMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetOpportunityOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetOpportunityInput struct {

	// Specifies the catalog associated with the request. This field takes a string
	// value from a predefined list: AWS or Sandbox . The catalog determines which
	// environment the opportunity is fetched from. Use AWS to retrieve opportunities
	// in the Amazon Web Services catalog, and Sandbox to retrieve opportunities in a
	// secure, isolated testing environment.
	//
	// This member is required.
	Catalog *string

	// Read-only, system generated Opportunity unique identifier.
	//
	// This member is required.
	Identifier *string

	noSmithyDocumentSerde
}

type GetOpportunityOutput struct {

	// Specifies the catalog associated with the request. This field takes a string
	// value from a predefined list: AWS or Sandbox . The catalog determines which
	// environment the opportunity information is retrieved from. Use AWS to retrieve
	// opportunities in the Amazon Web Services catalog, and Sandbox to retrieve
	// opportunities in a secure and isolated testing environment.
	//
	// This member is required.
	Catalog *string

	// DateTime when the Opportunity was last created.
	//
	// This member is required.
	CreatedDate *time.Time

	// Read-only, system generated Opportunity unique identifier.
	//
	// This member is required.
	Id *string

	// DateTime when the opportunity was last modified.
	//
	// This member is required.
	LastModifiedDate *time.Time

	// Provides information about the associations of other entities with the
	// opportunity. These entities include identifiers for AWSProducts , Partner
	// Solutions , and AWSMarketplaceOffers .
	//
	// This member is required.
	RelatedEntityIdentifiers *types.RelatedEntityIdentifiers

	// The Amazon Resource Name (ARN) that uniquely identifies the opportunity.
	Arn *string

	// Specifies details of the customer associated with the Opportunity .
	Customer *types.Customer

	// An object that contains lifecycle details for the Opportunity .
	LifeCycle *types.LifeCycle

	// An object that contains marketing details for the Opportunity .
	Marketing *types.Marketing

	// Indicates whether the Opportunity pertains to a national security project. This
	// field must be set to true only when the customer's industry is Government.
	// Additional privacy and security measures apply during the review and management
	// process for opportunities marked as NationalSecurity .
	NationalSecurity types.NationalSecurity

	// Represents the internal team handling the opportunity. Specify the members
	// involved in collaborating on this opportunity within the partner's organization.
	OpportunityTeam []types.Contact

	// Specifies the opportunity type as renewal, new, or expansion.
	//
	// Opportunity types:
	//
	//   - New opportunity: Represents a new business opportunity with a potential
	//   customer that's not previously engaged with your solutions or services.
	//
	//   - Renewal opportunity: Represents an opportunity to renew an existing
	//   contract or subscription with a current customer, which helps to ensure service
	//   continuity.
	//
	//   - Expansion opportunity: Represents an opportunity to expand the scope of a
	//   customer's contract or subscription, either by adding new services or increasing
	//   the volume of existing services.
	OpportunityType types.OpportunityType

	// Specifies the opportunity's unique identifier in the partner's CRM system. This
	// value is essential to track and reconcile because it's included in the outbound
	// payload sent back to the partner.
	PartnerOpportunityIdentifier *string

	// Identifies the type of support the partner needs from Amazon Web Services.
	//
	// Valid values:
	//
	//   - Cosell—Architectural Validation: Confirmation from Amazon Web Services that
	//   the partner's proposed solution architecture is aligned with Amazon Web Services
	//   best practices and poses minimal architectural risks.
	//
	//   - Cosell—Business Presentation: Request Amazon Web Services seller's
	//   participation in a joint customer presentation.
	//
	//   - Cosell—Competitive Information: Access to Amazon Web Services competitive
	//   resources and support for the partner's proposed solution.
	//
	//   - Cosell—Pricing Assistance: Connect with an Amazon Web Services seller for
	//   support situations where a partner may be receiving an upfront discount on a
	//   service (for example: EDP deals).
	//
	//   - Cosell—Technical Consultation: Connect with an Amazon Web Services
	//   Solutions Architect to address the partner's questions about the proposed
	//   solution.
	//
	//   - Cosell—Total Cost of Ownership Evaluation: Assistance with quoting
	//   different cost savings of proposed solutions on Amazon Web Services versus
	//   on-premises or a traditional hosting environment.
	//
	//   - Cosell—Deal Support: Request Amazon Web Services seller's support to
	//   progress the opportunity (for example: joint customer call, strategic
	//   positioning).
	//
	//   - Cosell—Support for Public Tender/RFx: Opportunity related to the public
	//   sector where the partner needs Amazon Web Services RFx support.
	PrimaryNeedsFromAws []types.PrimaryNeedFromAws

	// An object that contains project details summary for the Opportunity .
	Project *types.Project

	// Specifies details of a customer's procurement terms. Required only for partners
	// in eligible programs.
	SoftwareRevenue *types.SoftwareRevenue

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetOpportunityMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson10_serializeOpGetOpportunity{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson10_deserializeOpGetOpportunity{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "GetOpportunity"); err != nil {
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
	if err = addOpGetOpportunityValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetOpportunity(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opGetOpportunity(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "GetOpportunity",
	}
}
