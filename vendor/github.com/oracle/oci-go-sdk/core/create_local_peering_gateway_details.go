// Copyright (c) 2016, 2018, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// Core Services API
//
// APIs for Networking Service, Compute Service, and Block Volume Service.
//

package core

import (
	"github.com/oracle/oci-go-sdk/common"
)

// CreateLocalPeeringGatewayDetails The representation of CreateLocalPeeringGatewayDetails
type CreateLocalPeeringGatewayDetails struct {

	// The OCID of the compartment containing the local peering gateway (LPG).
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The OCID of the VCN the LPG belongs to.
	VcnId *string `mandatory:"true" json:"vcnId"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// For more information, see Resource Tags (https://docs.us-phoenix-1.oraclecloud.com/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Operations": {"CostCenter": "42"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// A user-friendly name. Does not have to be unique, and it's changeable. Avoid
	// entering confidential information.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// Free-form tags for this resource. Each tag is a simple key-value pair with no
	// predefined name, type, or namespace. For more information, see
	// Resource Tags (https://docs.us-phoenix-1.oraclecloud.com/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// The OCID of the route table the LPG will use.
	// If you don't specify a route table here, the LPG is created without an associated route
	// table. The Networking service does NOT automatically associate the attached VCN's default route table
	// with the LPG.
	RouteTableId *string `mandatory:"false" json:"routeTableId"`
}

func (m CreateLocalPeeringGatewayDetails) String() string {
	return common.PointerString(m)
}
