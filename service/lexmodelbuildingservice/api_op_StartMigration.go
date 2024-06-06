// Code generated by smithy-go-codegen DO NOT EDIT.

package lexmodelbuildingservice

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/lexmodelbuildingservice/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Starts migrating a bot from Amazon Lex V1 to Amazon Lex V2. Migrate your bot
// when you want to take advantage of the new features of Amazon Lex V2.
//
// For more information, see [Migrating a bot] in the Amazon Lex developer guide.
//
// [Migrating a bot]: https://docs.aws.amazon.com/lex/latest/dg/migrate.html
func (c *Client) StartMigration(ctx context.Context, params *StartMigrationInput, optFns ...func(*Options)) (*StartMigrationOutput, error) {
	if params == nil {
		params = &StartMigrationInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "StartMigration", params, optFns, c.addOperationStartMigrationMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*StartMigrationOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type StartMigrationInput struct {

	// The strategy used to conduct the migration.
	//
	//   - CREATE_NEW - Creates a new Amazon Lex V2 bot and migrates the Amazon Lex V1
	//   bot to the new bot.
	//
	//   - UPDATE_EXISTING - Overwrites the existing Amazon Lex V2 bot metadata and the
	//   locale being migrated. It doesn't change any other locales in the Amazon Lex V2
	//   bot. If the locale doesn't exist, a new locale is created in the Amazon Lex V2
	//   bot.
	//
	// This member is required.
	MigrationStrategy types.MigrationStrategy

	// The name of the Amazon Lex V1 bot that you are migrating to Amazon Lex V2.
	//
	// This member is required.
	V1BotName *string

	// The version of the bot to migrate to Amazon Lex V2. You can migrate the $LATEST
	// version as well as any numbered version.
	//
	// This member is required.
	V1BotVersion *string

	// The name of the Amazon Lex V2 bot that you are migrating the Amazon Lex V1 bot
	// to.
	//
	//   - If the Amazon Lex V2 bot doesn't exist, you must use the CREATE_NEW
	//   migration strategy.
	//
	//   - If the Amazon Lex V2 bot exists, you must use the UPDATE_EXISTING migration
	//   strategy to change the contents of the Amazon Lex V2 bot.
	//
	// This member is required.
	V2BotName *string

	// The IAM role that Amazon Lex uses to run the Amazon Lex V2 bot.
	//
	// This member is required.
	V2BotRole *string

	noSmithyDocumentSerde
}

type StartMigrationOutput struct {

	// The unique identifier that Amazon Lex assigned to the migration.
	MigrationId *string

	// The strategy used to conduct the migration.
	MigrationStrategy types.MigrationStrategy

	// The date and time that the migration started.
	MigrationTimestamp *time.Time

	// The locale used for the Amazon Lex V1 bot.
	V1BotLocale types.Locale

	// The name of the Amazon Lex V1 bot that you are migrating to Amazon Lex V2.
	V1BotName *string

	// The version of the bot to migrate to Amazon Lex V2.
	V1BotVersion *string

	// The unique identifier for the Amazon Lex V2 bot.
	V2BotId *string

	// The IAM role that Amazon Lex uses to run the Amazon Lex V2 bot.
	V2BotRole *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationStartMigrationMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpStartMigration{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpStartMigration{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "StartMigration"); err != nil {
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
	if err = addOpStartMigrationValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opStartMigration(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opStartMigration(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "StartMigration",
	}
}
