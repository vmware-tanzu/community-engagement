// Copyright 2021 VMware Shoutouter contributors. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0

package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidateOrg(t *testing.T) {
	// Missing org name should return an error
	c := ProjectConfig{
		Token:     "asdf",
		TeamNames: []string{"team"},
	}
	assert.Error(t, c.Validate())

	// Adding an org name will make the error go away
	c.OrgName = "org"
	assert.NoError(t, c.Validate())
}

func TestValidateTeam(t *testing.T) {
	// Missing team names should return an error
	c := ProjectConfig{
		Token:   "asdf",
		OrgName: "org",
	}
	assert.Error(t, c.Validate())

	// Adding a team name will make the error go away
	c.TeamNames = []string{"team"}
	assert.NoError(t, c.Validate())
}

func TestValidateToken(t *testing.T) {
	// Missing token should return an error
	c := ProjectConfig{
		OrgName:   "org",
		TeamNames: []string{"team"},
	}
	assert.Error(t, c.Validate())

	// Adding a token will make the error go away
	c.Token = "asdf"
	assert.NoError(t, c.Validate())
}

func TestNewConfigFromFileSucceeds(t *testing.T) {
	c, err := NewConfigFromFile("testdata/good_example.yaml")
	assert.Nil(t, err)
	assert.NoError(t, c.Validate())
}

func TestNewConfigFromFileMissingTeams(t *testing.T) {
	c, err := NewConfigFromFile("testdata/missing_teams.yaml")
	assert.NoError(t, err)
	assert.Equal(t, NoTeamError, c.Validate())
}
