// Copyright 2021 VMware Shoutouter contributors. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0

package main

import (
	"flag"
	"fmt"
	"time"

	log "github.com/sirupsen/logrus"

	"github.com/google/go-github/v32/github"
)

type args struct {
	debug                    bool
	config, token, team, org string
	days                     int
}

var a *args

func main() {

	c, err := makeConfig()
	if err != nil {
		log.Fatal(err)
	}

	if err := c.Validate(); err != nil {
		log.Fatal(err)
	}

	p := NewProject(c)

	err = p.GetDevs()
	if err != nil {
		log.Fatal(err)
	}

	for _, d := range p.Devs {
		if d != nil {
			log.Debug("Login: %s\n", d.GetLogin())
		}
	}

	err = p.GetReposFromTeam()
	if err != nil {
		log.Fatal(err)
	}

	log.Debug("---- REPOS ----")
	for _, r := range p.Repos {
		log.Debug("%s\n", *r.Name)
	}

	log.Debug("---- PRs ----")
	err = p.GetPullRequests()
	if err != nil {
		log.Fatal(err)
	}
	log.Debug("PRs Found: %d\n", len(p.PullRequests))

	i := 0
	emptyTime := time.Time{}
	for _, pr := range p.PullRequests {

		if pr.GetMergedAt() != emptyTime {
			// Skip PRs older than specified dates
			// this works for now, but should probably be put in a function or method.
			if a.days > -1 {
				age := time.Since(pr.GetMergedAt())
				if int(age.Hours()/24) > a.days {
					continue
				}
			}

			if fromMember(pr, p) {
				continue
			}
			i++
			printShoutout(pr)
		}
	}

	fmt.Printf("PRs merged from non-core memembers: %d\n", i)
}

func init() {
	a = &args{}
	flag.StringVar(&a.config, "config", "", "--config file.yaml")
	flag.StringVar(&a.token, "token", "", "--token ${GITHUB_API_TOKEN}")
	flag.StringVar(&a.org, "org", "", "--org ${GITHUB_ORG_NAME}")
	flag.StringVar(&a.team, "team", "", "--team ${GITHUB_TEAM_NAME}")
	flag.BoolVar(&a.debug, "debug", false, "--debug for detailed logging")
	flag.IntVar(&a.days, "days", -1, "--days for a window to check PRs (-1 for all (default), 30 for a month)")

	flag.Parse()

	if a.debug {
		log.SetLevel(log.DebugLevel)
	}
}

func makeConfig() (*ProjectConfig, error) {
	var err error
	c := &ProjectConfig{}

	if a != nil && a.config != "" {
		c, err = NewConfigFromFile(a.config)
	}
	if err != nil {
		return nil, err
	}

	if a != nil && a.org != "" {
		c.OrgName = a.org
	}

	if a.team != "" {
		c.TeamNames = append(c.TeamNames, a.team)
	}

	if a.token != "" {
		c.Token = a.token
	}

	return c, nil
}

// Formats the shoutouts in a consistent format.
func printShoutout(pr *github.PullRequest) {
	fmt.Printf("- [@%s](%s): [%s](%s)\n", *pr.User.Login, *pr.User.HTMLURL, *pr.Title, *pr.HTMLURL)
}

// fromMember determines if a PR was written by a core team member
func fromMember(pr *github.PullRequest, p *Project) bool {
	for _, m := range p.Devs {
		if pr.GetUser().GetLogin() == m.GetLogin() {
			return true
		}
	}
	return false
}
