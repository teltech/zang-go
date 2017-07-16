// Copyright 2017 The Zang Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package zang

import (
	"fmt"

	"github.com/zang-cloud/zang-go/helpers"
)

// ListCredentialLists -
func (c *Client) ListCredentialLists(params map[string]string) (resp *Response, err error) {
	if cerr := c.Config.Validate(); cerr != nil {
		return nil, cerr
	}

	resp, err = c.Request.Get(
		fmt.Sprintf("%s/SIP/CredentialLists", c.Config.AccountSid),
		params,
	)
	return
}

// GetCredentialList -
func (c *Client) GetCredentialList(sid string, params map[string]string) (resp *Response, err error) {
	if cerr := c.Config.Validate(); cerr != nil {
		return nil, cerr
	}

	if verr := helpers.ValidateSid(sid); err != nil {
		return nil, verr
	}

	resp, err = c.Request.Get(
		fmt.Sprintf("%s/SIP/CredentialLists/%s", c.Config.AccountSid, sid),
		params,
	)
	return
}

// CreateCredentialList -
func (c *Client) CreateCredentialList(params map[string]string) (resp *Response, err error) {
	if cerr := c.Config.Validate(); cerr != nil {
		return nil, cerr
	}

	resp, err = c.Request.Create(
		fmt.Sprintf("%s/SIP/CredentialLists", c.Config.AccountSid),
		params,
	)
	return
}

// UpdateCredentialList -
func (c *Client) UpdateCredentialList(sid string, params map[string]string) (resp *Response, err error) {
	if cerr := c.Config.Validate(); cerr != nil {
		return nil, cerr
	}

	if verr := helpers.ValidateSid(sid); err != nil {
		return nil, verr
	}

	resp, err = c.Request.Update(
		fmt.Sprintf("%s/SIP/CredentialLists/%s", c.Config.AccountSid, sid),
		params,
	)
	return
}

// DeleteCredentialList -
func (c *Client) DeleteCredentialList(sid string, params map[string]string) (resp *Response, err error) {
	if cerr := c.Config.Validate(); cerr != nil {
		return nil, cerr
	}

	if verr := helpers.ValidateSid(sid); err != nil {
		return nil, verr
	}

	resp, err = c.Request.Delete(
		fmt.Sprintf("%s/SIP/CredentialLists/%s", c.Config.AccountSid, sid),
		params,
	)
	return
}

// ListSipCredential -
func (c *Client) ListSipCredential(credentialListSid string, params map[string]string) (resp *Response, err error) {
	if cerr := c.Config.Validate(); cerr != nil {
		return nil, cerr
	}

	resp, err = c.Request.Get(
		fmt.Sprintf("%s/SIP/CredentialLists/%s/Credentials", c.Config.AccountSid, credentialListSid),
		params,
	)
	return
}

// GetSipCredential -
func (c *Client) GetSipCredential(credentialListSid string, sid string, params map[string]string) (resp *Response, err error) {
	if cerr := c.Config.Validate(); cerr != nil {
		return nil, cerr
	}

	if verr := helpers.ValidateSid(sid); err != nil {
		return nil, verr
	}

	resp, err = c.Request.Get(
		fmt.Sprintf("%s/SIP/CredentialLists/%s/Credentials/%s", c.Config.AccountSid, credentialListSid, sid),
		params,
	)
	return
}

// CreateSipCredential -
func (c *Client) CreateSipCredential(credentialListSid string, params map[string]string) (resp *Response, err error) {
	if cerr := c.Config.Validate(); cerr != nil {
		return nil, cerr
	}

	resp, err = c.Request.Create(
		fmt.Sprintf("%s/SIP/CredentialLists/%s/Credentials", c.Config.AccountSid, credentialListSid),
		params,
	)
	return
}

// UpdateSipCredential -
func (c *Client) UpdateSipCredential(credentialListSid string, sid string, params map[string]string) (resp *Response, err error) {
	if cerr := c.Config.Validate(); cerr != nil {
		return nil, cerr
	}

	if verr := helpers.ValidateSid(sid); err != nil {
		return nil, verr
	}

	resp, err = c.Request.Update(
		fmt.Sprintf("%s/SIP/CredentialLists/%s/Credentials/%s", c.Config.AccountSid, credentialListSid, sid),
		params,
	)
	return
}

// DeleteSipCredential -
func (c *Client) DeleteSipCredential(credentialListSid string, sid string, params map[string]string) (resp *Response, err error) {
	if cerr := c.Config.Validate(); cerr != nil {
		return nil, cerr
	}

	if verr := helpers.ValidateSid(sid); err != nil {
		return nil, verr
	}

	resp, err = c.Request.Delete(
		fmt.Sprintf("%s/SIP/CredentialLists/%s/Credentials/%s", c.Config.AccountSid, credentialListSid, sid),
		params,
	)
	return
}