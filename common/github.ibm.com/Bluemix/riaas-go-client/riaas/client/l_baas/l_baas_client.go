// Code generated by go-swagger; DO NOT EDIT.

package l_baas

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new l baas API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for l baas API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
DeleteLoadBalancersID deletes a load balancer

Deletes a load balancer. This operation cannot be reversed.
*/
func (a *Client) DeleteLoadBalancersID(params *DeleteLoadBalancersIDParams, authInfo runtime.ClientAuthInfoWriter) (*DeleteLoadBalancersIDNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteLoadBalancersIDParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "DeleteLoadBalancersID",
		Method:             "DELETE",
		PathPattern:        "/load_balancers/{id}",
		ProducesMediaTypes: []string{""},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &DeleteLoadBalancersIDReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*DeleteLoadBalancersIDNoContent), nil

}

/*
DeleteLoadBalancersIDListenersListenerID deletes a listener

Deletes a listener. This operation cannot be reversed.
*/
func (a *Client) DeleteLoadBalancersIDListenersListenerID(params *DeleteLoadBalancersIDListenersListenerIDParams, authInfo runtime.ClientAuthInfoWriter) (*DeleteLoadBalancersIDListenersListenerIDNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteLoadBalancersIDListenersListenerIDParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "DeleteLoadBalancersIDListenersListenerID",
		Method:             "DELETE",
		PathPattern:        "/load_balancers/{id}/listeners/{listener_id}",
		ProducesMediaTypes: []string{""},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &DeleteLoadBalancersIDListenersListenerIDReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*DeleteLoadBalancersIDListenersListenerIDNoContent), nil

}

/*
DeleteLoadBalancersIDListenersListenerIDPoliciesPolicyID deletes a policy of the load balancer listener

Deletes a policy of the load balancer listener. This operation cannot be reversed.
*/
func (a *Client) DeleteLoadBalancersIDListenersListenerIDPoliciesPolicyID(params *DeleteLoadBalancersIDListenersListenerIDPoliciesPolicyIDParams, authInfo runtime.ClientAuthInfoWriter) (*DeleteLoadBalancersIDListenersListenerIDPoliciesPolicyIDNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteLoadBalancersIDListenersListenerIDPoliciesPolicyIDParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "DeleteLoadBalancersIDListenersListenerIDPoliciesPolicyID",
		Method:             "DELETE",
		PathPattern:        "/load_balancers/{id}/listeners/{listener_id}/policies/{policy_id}",
		ProducesMediaTypes: []string{""},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &DeleteLoadBalancersIDListenersListenerIDPoliciesPolicyIDReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*DeleteLoadBalancersIDListenersListenerIDPoliciesPolicyIDNoContent), nil

}

/*
DeleteLoadBalancersIDListenersListenerIDPoliciesPolicyIDRulesRuleID deletes a rule from the load balancer listener policy

Deletes a rule from the load balancer listener policy. This operation cannot be reversed.
*/
func (a *Client) DeleteLoadBalancersIDListenersListenerIDPoliciesPolicyIDRulesRuleID(params *DeleteLoadBalancersIDListenersListenerIDPoliciesPolicyIDRulesRuleIDParams, authInfo runtime.ClientAuthInfoWriter) (*DeleteLoadBalancersIDListenersListenerIDPoliciesPolicyIDRulesRuleIDNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteLoadBalancersIDListenersListenerIDPoliciesPolicyIDRulesRuleIDParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "DeleteLoadBalancersIDListenersListenerIDPoliciesPolicyIDRulesRuleID",
		Method:             "DELETE",
		PathPattern:        "/load_balancers/{id}/listeners/{listener_id}/policies/{policy_id}/rules/{rule_id}",
		ProducesMediaTypes: []string{""},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &DeleteLoadBalancersIDListenersListenerIDPoliciesPolicyIDRulesRuleIDReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*DeleteLoadBalancersIDListenersListenerIDPoliciesPolicyIDRulesRuleIDNoContent), nil

}

/*
DeleteLoadBalancersIDPoolsPoolID deletes a pool

Deletes a pool. This operation cannot be reversed.
*/
func (a *Client) DeleteLoadBalancersIDPoolsPoolID(params *DeleteLoadBalancersIDPoolsPoolIDParams, authInfo runtime.ClientAuthInfoWriter) (*DeleteLoadBalancersIDPoolsPoolIDNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteLoadBalancersIDPoolsPoolIDParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "DeleteLoadBalancersIDPoolsPoolID",
		Method:             "DELETE",
		PathPattern:        "/load_balancers/{id}/pools/{pool_id}",
		ProducesMediaTypes: []string{""},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &DeleteLoadBalancersIDPoolsPoolIDReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*DeleteLoadBalancersIDPoolsPoolIDNoContent), nil

}

/*
DeleteLoadBalancersIDPoolsPoolIDMembersMemberID deletes a member from the pool

Deletes a member from the pool. This operation cannot be reversed.
*/
func (a *Client) DeleteLoadBalancersIDPoolsPoolIDMembersMemberID(params *DeleteLoadBalancersIDPoolsPoolIDMembersMemberIDParams, authInfo runtime.ClientAuthInfoWriter) (*DeleteLoadBalancersIDPoolsPoolIDMembersMemberIDNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteLoadBalancersIDPoolsPoolIDMembersMemberIDParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "DeleteLoadBalancersIDPoolsPoolIDMembersMemberID",
		Method:             "DELETE",
		PathPattern:        "/load_balancers/{id}/pools/{pool_id}/members/{member_id}",
		ProducesMediaTypes: []string{""},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &DeleteLoadBalancersIDPoolsPoolIDMembersMemberIDReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*DeleteLoadBalancersIDPoolsPoolIDMembersMemberIDNoContent), nil

}

/*
GetLoadBalancers retrieves all load balancers

Retrieves a paginated list of all load balancers belonging to this account.
*/
func (a *Client) GetLoadBalancers(params *GetLoadBalancersParams, authInfo runtime.ClientAuthInfoWriter) (*GetLoadBalancersOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetLoadBalancersParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetLoadBalancers",
		Method:             "GET",
		PathPattern:        "/load_balancers/",
		ProducesMediaTypes: []string{""},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetLoadBalancersReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetLoadBalancersOK), nil

}

/*
GetLoadBalancersID retrieves a load balancer

Retrieves a single load balancer specified by the identifier in the URL path.
*/
func (a *Client) GetLoadBalancersID(params *GetLoadBalancersIDParams, authInfo runtime.ClientAuthInfoWriter) (*GetLoadBalancersIDOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetLoadBalancersIDParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetLoadBalancersID",
		Method:             "GET",
		PathPattern:        "/load_balancers/{id}",
		ProducesMediaTypes: []string{""},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetLoadBalancersIDReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetLoadBalancersIDOK), nil

}

/*
GetLoadBalancersIDListeners retrieves all listeners of the load balancer

Retrieves a list of all listeners belonging to the load balancer.
*/
func (a *Client) GetLoadBalancersIDListeners(params *GetLoadBalancersIDListenersParams, authInfo runtime.ClientAuthInfoWriter) (*GetLoadBalancersIDListenersOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetLoadBalancersIDListenersParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetLoadBalancersIDListeners",
		Method:             "GET",
		PathPattern:        "/load_balancers/{id}/listeners",
		ProducesMediaTypes: []string{""},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetLoadBalancersIDListenersReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetLoadBalancersIDListenersOK), nil

}

/*
GetLoadBalancersIDListenersListenerID retrieves a listener

Retrieves a single listener specified by the identifier in the URL path.
*/
func (a *Client) GetLoadBalancersIDListenersListenerID(params *GetLoadBalancersIDListenersListenerIDParams, authInfo runtime.ClientAuthInfoWriter) (*GetLoadBalancersIDListenersListenerIDOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetLoadBalancersIDListenersListenerIDParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetLoadBalancersIDListenersListenerID",
		Method:             "GET",
		PathPattern:        "/load_balancers/{id}/listeners/{listener_id}",
		ProducesMediaTypes: []string{""},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetLoadBalancersIDListenersListenerIDReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetLoadBalancersIDListenersListenerIDOK), nil

}

/*
GetLoadBalancersIDListenersListenerIDPolicies retrieves all policies of the listener

Retrieves a list of all policies belonging to the load balancer listener.
*/
func (a *Client) GetLoadBalancersIDListenersListenerIDPolicies(params *GetLoadBalancersIDListenersListenerIDPoliciesParams, authInfo runtime.ClientAuthInfoWriter) (*GetLoadBalancersIDListenersListenerIDPoliciesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetLoadBalancersIDListenersListenerIDPoliciesParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetLoadBalancersIDListenersListenerIDPolicies",
		Method:             "GET",
		PathPattern:        "/load_balancers/{id}/listeners/{listener_id}/policies",
		ProducesMediaTypes: []string{""},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetLoadBalancersIDListenersListenerIDPoliciesReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetLoadBalancersIDListenersListenerIDPoliciesOK), nil

}

/*
GetLoadBalancersIDListenersListenerIDPoliciesPolicyID retrieves a policy of the load balancer listener

Retrieve a single policy specified by the identifier in the URL path.
*/
func (a *Client) GetLoadBalancersIDListenersListenerIDPoliciesPolicyID(params *GetLoadBalancersIDListenersListenerIDPoliciesPolicyIDParams, authInfo runtime.ClientAuthInfoWriter) (*GetLoadBalancersIDListenersListenerIDPoliciesPolicyIDOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetLoadBalancersIDListenersListenerIDPoliciesPolicyIDParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetLoadBalancersIDListenersListenerIDPoliciesPolicyID",
		Method:             "GET",
		PathPattern:        "/load_balancers/{id}/listeners/{listener_id}/policies/{policy_id}",
		ProducesMediaTypes: []string{""},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetLoadBalancersIDListenersListenerIDPoliciesPolicyIDReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetLoadBalancersIDListenersListenerIDPoliciesPolicyIDOK), nil

}

/*
GetLoadBalancersIDListenersListenerIDPoliciesPolicyIDRules lists all rules of the load balancer listener policy

Retrieves a list of all rules belonging to the load balancer listener policy.
*/
func (a *Client) GetLoadBalancersIDListenersListenerIDPoliciesPolicyIDRules(params *GetLoadBalancersIDListenersListenerIDPoliciesPolicyIDRulesParams, authInfo runtime.ClientAuthInfoWriter) (*GetLoadBalancersIDListenersListenerIDPoliciesPolicyIDRulesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetLoadBalancersIDListenersListenerIDPoliciesPolicyIDRulesParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetLoadBalancersIDListenersListenerIDPoliciesPolicyIDRules",
		Method:             "GET",
		PathPattern:        "/load_balancers/{id}/listeners/{listener_id}/policies/{policy_id}/rules",
		ProducesMediaTypes: []string{""},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetLoadBalancersIDListenersListenerIDPoliciesPolicyIDRulesReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetLoadBalancersIDListenersListenerIDPoliciesPolicyIDRulesOK), nil

}

/*
GetLoadBalancersIDListenersListenerIDPoliciesPolicyIDRulesRuleID retrieves a rule of the load balancer listener policy

Retrieves a single rule specified by the identifier in the URL path.
*/
func (a *Client) GetLoadBalancersIDListenersListenerIDPoliciesPolicyIDRulesRuleID(params *GetLoadBalancersIDListenersListenerIDPoliciesPolicyIDRulesRuleIDParams, authInfo runtime.ClientAuthInfoWriter) (*GetLoadBalancersIDListenersListenerIDPoliciesPolicyIDRulesRuleIDOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetLoadBalancersIDListenersListenerIDPoliciesPolicyIDRulesRuleIDParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetLoadBalancersIDListenersListenerIDPoliciesPolicyIDRulesRuleID",
		Method:             "GET",
		PathPattern:        "/load_balancers/{id}/listeners/{listener_id}/policies/{policy_id}/rules/{rule_id}",
		ProducesMediaTypes: []string{""},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetLoadBalancersIDListenersListenerIDPoliciesPolicyIDRulesRuleIDReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetLoadBalancersIDListenersListenerIDPoliciesPolicyIDRulesRuleIDOK), nil

}

/*
GetLoadBalancersIDPools retrieves all pools of the load balancer

Retrieves a list of all pools belonging to the load balancer.
*/
func (a *Client) GetLoadBalancersIDPools(params *GetLoadBalancersIDPoolsParams, authInfo runtime.ClientAuthInfoWriter) (*GetLoadBalancersIDPoolsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetLoadBalancersIDPoolsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetLoadBalancersIDPools",
		Method:             "GET",
		PathPattern:        "/load_balancers/{id}/pools",
		ProducesMediaTypes: []string{""},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetLoadBalancersIDPoolsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetLoadBalancersIDPoolsOK), nil

}

/*
GetLoadBalancersIDPoolsPoolID retrieves a pool

Retrieves a single pool specified by the identifier in the URL path.
*/
func (a *Client) GetLoadBalancersIDPoolsPoolID(params *GetLoadBalancersIDPoolsPoolIDParams, authInfo runtime.ClientAuthInfoWriter) (*GetLoadBalancersIDPoolsPoolIDOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetLoadBalancersIDPoolsPoolIDParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetLoadBalancersIDPoolsPoolID",
		Method:             "GET",
		PathPattern:        "/load_balancers/{id}/pools/{pool_id}",
		ProducesMediaTypes: []string{""},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetLoadBalancersIDPoolsPoolIDReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetLoadBalancersIDPoolsPoolIDOK), nil

}

/*
GetLoadBalancersIDPoolsPoolIDMembers retrieves all members of the pool

Retrieves a paginated list of all members belonging to the pool.
*/
func (a *Client) GetLoadBalancersIDPoolsPoolIDMembers(params *GetLoadBalancersIDPoolsPoolIDMembersParams, authInfo runtime.ClientAuthInfoWriter) (*GetLoadBalancersIDPoolsPoolIDMembersOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetLoadBalancersIDPoolsPoolIDMembersParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetLoadBalancersIDPoolsPoolIDMembers",
		Method:             "GET",
		PathPattern:        "/load_balancers/{id}/pools/{pool_id}/members",
		ProducesMediaTypes: []string{""},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetLoadBalancersIDPoolsPoolIDMembersReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetLoadBalancersIDPoolsPoolIDMembersOK), nil

}

/*
GetLoadBalancersIDPoolsPoolIDMembersMemberID retrieves a member

Retrieves a single member specified by the identifier in the URL path.
*/
func (a *Client) GetLoadBalancersIDPoolsPoolIDMembersMemberID(params *GetLoadBalancersIDPoolsPoolIDMembersMemberIDParams, authInfo runtime.ClientAuthInfoWriter) (*GetLoadBalancersIDPoolsPoolIDMembersMemberIDOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetLoadBalancersIDPoolsPoolIDMembersMemberIDParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetLoadBalancersIDPoolsPoolIDMembersMemberID",
		Method:             "GET",
		PathPattern:        "/load_balancers/{id}/pools/{pool_id}/members/{member_id}",
		ProducesMediaTypes: []string{""},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetLoadBalancersIDPoolsPoolIDMembersMemberIDReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetLoadBalancersIDPoolsPoolIDMembersMemberIDOK), nil

}

/*
GetLoadBalancersIDStatistics retrieves statistics of a load balancer

Retrieves statistics of a load balancer specified by the identifier in the URL path.
*/
func (a *Client) GetLoadBalancersIDStatistics(params *GetLoadBalancersIDStatisticsParams, authInfo runtime.ClientAuthInfoWriter) (*GetLoadBalancersIDStatisticsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetLoadBalancersIDStatisticsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetLoadBalancersIDStatistics",
		Method:             "GET",
		PathPattern:        "/load_balancers/{id}/statistics",
		ProducesMediaTypes: []string{""},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetLoadBalancersIDStatisticsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetLoadBalancersIDStatisticsOK), nil

}

/*
PatchLoadBalancersID updates a load balancer

Updates a load balancer.
*/
func (a *Client) PatchLoadBalancersID(params *PatchLoadBalancersIDParams, authInfo runtime.ClientAuthInfoWriter) (*PatchLoadBalancersIDOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPatchLoadBalancersIDParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "PatchLoadBalancersID",
		Method:             "PATCH",
		PathPattern:        "/load_balancers/{id}",
		ProducesMediaTypes: []string{""},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &PatchLoadBalancersIDReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*PatchLoadBalancersIDOK), nil

}

/*
PatchLoadBalancersIDListenersListenerID updates a listener

Updates a listener from a listener template.
*/
func (a *Client) PatchLoadBalancersIDListenersListenerID(params *PatchLoadBalancersIDListenersListenerIDParams, authInfo runtime.ClientAuthInfoWriter) (*PatchLoadBalancersIDListenersListenerIDOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPatchLoadBalancersIDListenersListenerIDParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "PatchLoadBalancersIDListenersListenerID",
		Method:             "PATCH",
		PathPattern:        "/load_balancers/{id}/listeners/{listener_id}",
		ProducesMediaTypes: []string{""},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &PatchLoadBalancersIDListenersListenerIDReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*PatchLoadBalancersIDListenersListenerIDOK), nil

}

/*
PatchLoadBalancersIDListenersListenerIDPoliciesPolicyID updates a policy of the load balancer listener

Updates a policy from a policy template.
*/
func (a *Client) PatchLoadBalancersIDListenersListenerIDPoliciesPolicyID(params *PatchLoadBalancersIDListenersListenerIDPoliciesPolicyIDParams, authInfo runtime.ClientAuthInfoWriter) (*PatchLoadBalancersIDListenersListenerIDPoliciesPolicyIDOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPatchLoadBalancersIDListenersListenerIDPoliciesPolicyIDParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "PatchLoadBalancersIDListenersListenerIDPoliciesPolicyID",
		Method:             "PATCH",
		PathPattern:        "/load_balancers/{id}/listeners/{listener_id}/policies/{policy_id}",
		ProducesMediaTypes: []string{""},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &PatchLoadBalancersIDListenersListenerIDPoliciesPolicyIDReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*PatchLoadBalancersIDListenersListenerIDPoliciesPolicyIDOK), nil

}

/*
PatchLoadBalancersIDListenersListenerIDPoliciesPolicyIDRulesRuleID updates a rule of the load balancer listener policy

Updates a rule of the load balancer listener policy.
*/
func (a *Client) PatchLoadBalancersIDListenersListenerIDPoliciesPolicyIDRulesRuleID(params *PatchLoadBalancersIDListenersListenerIDPoliciesPolicyIDRulesRuleIDParams, authInfo runtime.ClientAuthInfoWriter) (*PatchLoadBalancersIDListenersListenerIDPoliciesPolicyIDRulesRuleIDOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPatchLoadBalancersIDListenersListenerIDPoliciesPolicyIDRulesRuleIDParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "PatchLoadBalancersIDListenersListenerIDPoliciesPolicyIDRulesRuleID",
		Method:             "PATCH",
		PathPattern:        "/load_balancers/{id}/listeners/{listener_id}/policies/{policy_id}/rules/{rule_id}",
		ProducesMediaTypes: []string{""},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &PatchLoadBalancersIDListenersListenerIDPoliciesPolicyIDRulesRuleIDReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*PatchLoadBalancersIDListenersListenerIDPoliciesPolicyIDRulesRuleIDOK), nil

}

/*
PatchLoadBalancersIDPoolsPoolID updates a pool

Updates a pool from a pool template.
*/
func (a *Client) PatchLoadBalancersIDPoolsPoolID(params *PatchLoadBalancersIDPoolsPoolIDParams, authInfo runtime.ClientAuthInfoWriter) (*PatchLoadBalancersIDPoolsPoolIDOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPatchLoadBalancersIDPoolsPoolIDParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "PatchLoadBalancersIDPoolsPoolID",
		Method:             "PATCH",
		PathPattern:        "/load_balancers/{id}/pools/{pool_id}",
		ProducesMediaTypes: []string{""},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &PatchLoadBalancersIDPoolsPoolIDReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*PatchLoadBalancersIDPoolsPoolIDOK), nil

}

/*
PatchLoadBalancersIDPoolsPoolIDMembersMemberID updates a member

Updates an existing member from a member template.
*/
func (a *Client) PatchLoadBalancersIDPoolsPoolIDMembersMemberID(params *PatchLoadBalancersIDPoolsPoolIDMembersMemberIDParams, authInfo runtime.ClientAuthInfoWriter) (*PatchLoadBalancersIDPoolsPoolIDMembersMemberIDOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPatchLoadBalancersIDPoolsPoolIDMembersMemberIDParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "PatchLoadBalancersIDPoolsPoolIDMembersMemberID",
		Method:             "PATCH",
		PathPattern:        "/load_balancers/{id}/pools/{pool_id}/members/{member_id}",
		ProducesMediaTypes: []string{""},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &PatchLoadBalancersIDPoolsPoolIDMembersMemberIDReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*PatchLoadBalancersIDPoolsPoolIDMembersMemberIDOK), nil

}

/*
PostLoadBalancers creates and provisions a load balancer

Creates and provisions a new load balancer.
*/
func (a *Client) PostLoadBalancers(params *PostLoadBalancersParams, authInfo runtime.ClientAuthInfoWriter) (*PostLoadBalancersCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostLoadBalancersParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "PostLoadBalancers",
		Method:             "POST",
		PathPattern:        "/load_balancers/",
		ProducesMediaTypes: []string{""},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &PostLoadBalancersReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*PostLoadBalancersCreated), nil

}

/*
PostLoadBalancersIDListeners creates a listener

Creates a new listener to the load balancer.
*/
func (a *Client) PostLoadBalancersIDListeners(params *PostLoadBalancersIDListenersParams, authInfo runtime.ClientAuthInfoWriter) (*PostLoadBalancersIDListenersCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostLoadBalancersIDListenersParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "PostLoadBalancersIDListeners",
		Method:             "POST",
		PathPattern:        "/load_balancers/{id}/listeners",
		ProducesMediaTypes: []string{""},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &PostLoadBalancersIDListenersReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*PostLoadBalancersIDListenersCreated), nil

}

/*
PostLoadBalancersIDListenersListenerIDPolicies creates a policy for the load balancer listener

Creates a new policy to the load balancer listener.
*/
func (a *Client) PostLoadBalancersIDListenersListenerIDPolicies(params *PostLoadBalancersIDListenersListenerIDPoliciesParams, authInfo runtime.ClientAuthInfoWriter) (*PostLoadBalancersIDListenersListenerIDPoliciesCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostLoadBalancersIDListenersListenerIDPoliciesParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "PostLoadBalancersIDListenersListenerIDPolicies",
		Method:             "POST",
		PathPattern:        "/load_balancers/{id}/listeners/{listener_id}/policies",
		ProducesMediaTypes: []string{""},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &PostLoadBalancersIDListenersListenerIDPoliciesReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*PostLoadBalancersIDListenersListenerIDPoliciesCreated), nil

}

/*
PostLoadBalancersIDListenersListenerIDPoliciesPolicyIDRules creates a rule for the load balancer listener policy

Creates a new rule for the load balancer listener policy.
*/
func (a *Client) PostLoadBalancersIDListenersListenerIDPoliciesPolicyIDRules(params *PostLoadBalancersIDListenersListenerIDPoliciesPolicyIDRulesParams, authInfo runtime.ClientAuthInfoWriter) (*PostLoadBalancersIDListenersListenerIDPoliciesPolicyIDRulesCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostLoadBalancersIDListenersListenerIDPoliciesPolicyIDRulesParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "PostLoadBalancersIDListenersListenerIDPoliciesPolicyIDRules",
		Method:             "POST",
		PathPattern:        "/load_balancers/{id}/listeners/{listener_id}/policies/{policy_id}/rules",
		ProducesMediaTypes: []string{""},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &PostLoadBalancersIDListenersListenerIDPoliciesPolicyIDRulesReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*PostLoadBalancersIDListenersListenerIDPoliciesPolicyIDRulesCreated), nil

}

/*
PostLoadBalancersIDPools creates a pool

This request creates a new pool from a pool template.
*/
func (a *Client) PostLoadBalancersIDPools(params *PostLoadBalancersIDPoolsParams, authInfo runtime.ClientAuthInfoWriter) (*PostLoadBalancersIDPoolsCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostLoadBalancersIDPoolsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "PostLoadBalancersIDPools",
		Method:             "POST",
		PathPattern:        "/load_balancers/{id}/pools",
		ProducesMediaTypes: []string{""},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &PostLoadBalancersIDPoolsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*PostLoadBalancersIDPoolsCreated), nil

}

/*
PostLoadBalancersIDPoolsPoolIDMembers creates a member

Creates a new member and adds the member to the pool.
*/
func (a *Client) PostLoadBalancersIDPoolsPoolIDMembers(params *PostLoadBalancersIDPoolsPoolIDMembersParams, authInfo runtime.ClientAuthInfoWriter) (*PostLoadBalancersIDPoolsPoolIDMembersCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostLoadBalancersIDPoolsPoolIDMembersParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "PostLoadBalancersIDPoolsPoolIDMembers",
		Method:             "POST",
		PathPattern:        "/load_balancers/{id}/pools/{pool_id}/members",
		ProducesMediaTypes: []string{""},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &PostLoadBalancersIDPoolsPoolIDMembersReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*PostLoadBalancersIDPoolsPoolIDMembersCreated), nil

}

/*
PutLoadBalancersIDPoolsPoolIDMembers updates members of the pool

Updates members of the load balancer pool from a member template collection.
*/
func (a *Client) PutLoadBalancersIDPoolsPoolIDMembers(params *PutLoadBalancersIDPoolsPoolIDMembersParams, authInfo runtime.ClientAuthInfoWriter) (*PutLoadBalancersIDPoolsPoolIDMembersAccepted, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPutLoadBalancersIDPoolsPoolIDMembersParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "PutLoadBalancersIDPoolsPoolIDMembers",
		Method:             "PUT",
		PathPattern:        "/load_balancers/{id}/pools/{pool_id}/members",
		ProducesMediaTypes: []string{""},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &PutLoadBalancersIDPoolsPoolIDMembersReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*PutLoadBalancersIDPoolsPoolIDMembersAccepted), nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}