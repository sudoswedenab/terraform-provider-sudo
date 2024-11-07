package provider

import (
	"context"
	"crypto/sha256"
	"crypto/x509"
	"encoding/base64"
	"encoding/json"
	"encoding/pem"
	"errors"

	josev4 "github.com/go-jose/go-jose/v4"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

var _ datasource.DataSource = &JWKSDataSource{}

func NewJWKSDataSource() datasource.DataSource {
	return &JWKSDataSource{}
}

type JWKSDataSource struct{}

type JWKSDataSourceModel struct {
	RSAPrivateKeyPem types.String `tfsdk:"rsa_private_key_pem"`
	JSON             types.String `tfsdk:"json"`
}

func (s *JWKSDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_jwks"
}

func (s *JWKSDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data JWKSDataSourceModel

	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	rsaPrivateKeyPEM := data.RSAPrivateKeyPem.ValueString()

	block, _ := pem.Decode([]byte(rsaPrivateKeyPEM))
	if block == nil {
		panic(errors.New("empty block"))
	}

	rsaPrivateKey, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		panic(err)
	}

	publicKeyDER, err := x509.MarshalPKIXPublicKey(&rsaPrivateKey.PublicKey)
	if err != nil {
		panic(err)
	}

	hash := sha256.New()
	hash.Write(publicKeyDER)
	sha256sum := hash.Sum(nil)

	keyID := base64.RawURLEncoding.EncodeToString(sha256sum)

	jwks := josev4.JSONWebKeySet{
		Keys: []josev4.JSONWebKey{
			{
				Key:       &rsaPrivateKey.PublicKey,
				KeyID:     keyID,
				Algorithm: string(josev4.RS256),
				Use:       "sig",
			},
		},
	}

	b, err := json.Marshal(jwks)
	if err != nil {
		panic(err)
	}

	data.JSON = basetypes.NewStringValue(string(b))

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (s *JWKSDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"rsa_private_key_pem": schema.StringAttribute{
				Required:  true,
				Sensitive: true,
			},
			"json": schema.StringAttribute{
				Computed: true,
			},
		},
	}
}
