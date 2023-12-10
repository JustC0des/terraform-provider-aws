// Code generated by internal/generate/servicepackages/main.go; DO NOT EDIT.

package ssoadmin

import (
	"context"

	"github.com/hashicorp/terraform-provider-aws/internal/conns"
	"github.com/hashicorp/terraform-provider-aws/internal/types"
	"github.com/hashicorp/terraform-provider-aws/names"
)

type servicePackage struct{}

func (p *servicePackage) FrameworkDataSources(ctx context.Context) []*types.ServicePackageFrameworkDataSource {
	return []*types.ServicePackageFrameworkDataSource{
		{
			Factory: newDataSourceApplication,
			Name:    "Application",
		},
		{
			Factory: newDataSourceApplicationAssignments,
			Name:    "Application Assignments",
		},
		{
			Factory: newDataSourceApplicationProviders,
			Name:    "Application Providers",
		},
		{
			Factory: newDataSourcePrincipalApplicationAssignments,
			Name:    "Principal Application Assignments",
		},
	}
}

func (p *servicePackage) FrameworkResources(ctx context.Context) []*types.ServicePackageFrameworkResource {
	return []*types.ServicePackageFrameworkResource{
		{
			Factory: newResourceApplication,
			Name:    "Application",
			Tags:    &types.ServicePackageResourceTags{},
		},
		{
			Factory: newResourceApplicationAssignment,
			Name:    "Application Assignment",
		},
		{
			Factory: newResourceApplicationAssignmentConfiguration,
			Name:    "Application Assignment Configuration",
		},
	}
}

func (p *servicePackage) SDKDataSources(ctx context.Context) []*types.ServicePackageSDKDataSource {
	return []*types.ServicePackageSDKDataSource{
		{
			Factory:  DataSourceInstances,
			TypeName: "aws_ssoadmin_instances",
		},
		{
			Factory:  DataSourcePermissionSet,
			TypeName: "aws_ssoadmin_permission_set",
		},
	}
}

func (p *servicePackage) SDKResources(ctx context.Context) []*types.ServicePackageSDKResource {
	return []*types.ServicePackageSDKResource{
		{
			Factory:  ResourceAccountAssignment,
			TypeName: "aws_ssoadmin_account_assignment",
		},
		{
			Factory:  ResourceCustomerManagedPolicyAttachment,
			TypeName: "aws_ssoadmin_customer_managed_policy_attachment",
		},
		{
			Factory:  ResourceAccessControlAttributes,
			TypeName: "aws_ssoadmin_instance_access_control_attributes",
		},
		{
			Factory:  ResourceManagedPolicyAttachment,
			TypeName: "aws_ssoadmin_managed_policy_attachment",
		},
		{
			Factory:  ResourcePermissionSet,
			TypeName: "aws_ssoadmin_permission_set",
			Name:     "Permission Set",
			Tags:     &types.ServicePackageResourceTags{},
		},
		{
			Factory:  ResourcePermissionSetInlinePolicy,
			TypeName: "aws_ssoadmin_permission_set_inline_policy",
		},
		{
			Factory:  ResourcePermissionsBoundaryAttachment,
			TypeName: "aws_ssoadmin_permissions_boundary_attachment",
		},
		{
			Factory:  ResourceTrustedTokenIssuer,
			TypeName: "aws_ssoadmin_trusted_token_issuer",
			Tags:     &types.ServicePackageResourceTags{},
		},
	}
}

func (p *servicePackage) ServicePackageName() string {
	return names.SSOAdmin
}

func ServicePackage(ctx context.Context) conns.ServicePackage {
	return &servicePackage{}
}
