# Guidelines: Getting started with Community Engagement for Open Source projects

> **Required within the VMware Tanzu GitHub Organization**

The below document complements VMware's Open Source Guidelines and provides a checklist of what is needed by each open source project to be set up for healthy community engagement.

## Share Knowledge

1. The project README should be simple, focused, and mostly bullet item listed with steps to get up and running quickly. The README or landing page should include a project statement (what is does, why it’s important, and how it works), a list of features, screenshots (if a graphical UI exists), contribution guidelines, code of conduct, and links to relevant outside sources. ([EXAMPLE](https://gist.github.com/jonasrosland/8bf2e270887aa8514a19fd55335e9915)).

1. Add Code of Conduct that you receive from OSPO.

1. Have a clear “Getting started” guide, for both [users](https://velero.io/docs/v1.5.0/basic-install/) and [contributors](https://velero.io/docs/v1.5.0/start-contributing/).

1. Design and architecture documents should be made available to help users and contributors understand the underpinnings of the project ([EXAMPLE](https://github.com/vmware-tanzu/antrea/blob/main/docs/design/architecture.md)).

1. Make sure the main documentation page for the project is:

   1. Up to date.

   1. Works properly in all browsers and mobile devices.

   1. Is easy to contribute to - using tools such as Hugo to publish documentation and similar content has worked really well.

1. Shared communication should be easy to find, for example [blogs](https://octant.dev/blog/), [recorded meetings](https://www.youtube.com/playlist?list=PL7bmigfV0EqQRysvqvqOtRNk4L5S7uqwM), [meeting notes](https://hackmd.io/Jq6F5zqZR7S80CeDWUklkA?view), and [presentations](https://velero.io/resources/).

1. Have an initial blog post that states the vision, why the project was created versus contributing to an existing project, and gives a commitment to the community ([EXAMPLE](https://tanzu.vmware.com/content/blog/seeing-is-believing-octant-reveals-the-objects-running-in-kubernetes-clusters)).

## 2. Collaborate

1. Development solely done in public and upstream.

1. Use public tools for issue tracking and planning such as GitHub Projects and Zenhub Boards.

1. Ensure that each open source project has a Slack presence - make sure you are where your contributors are.

1. In the repositories, there needs to be an MAINTAINERS file that lists all maintainers of the project, and potential stakeholders ([EXAMPLE](https://github.com/goharbor/community/blob/master/MAINTAINERS.md)).

1. Document who owns key roles in each project, and include that in MAINTAINERS ([EXAMPLE 1](https://github.com/vmware-tanzu/velero/blob/main/MAINTAINERS.md) and [EXAMPLE 2](https://github.com/goharbor/community/blob/master/MAINTAINERS.md)), or have them listed on the project landing page ([EXAMPLE](https://carvel.dev/)).

1. Create sample issue templates to ensure predictable triaging of bugs, features and support issues  ([EXAMPLE 1](https://github.com/vmware-tanzu/sonobuoy/blob/master/.github/ISSUE_TEMPLATE/bug_report.md) and [EXAMPLE 2](https://github.com/vmware-tanzu/velero/blob/master/.github/ISSUE_TEMPLATE/feature-enhancement-request.md)).

1. Be a good open source citizen - “wear your agenda on your sleeve”. If you are clear about that then folks will be able to predict what you are looking for, and it is much easier to get to a "win-win" solution.

1. As the project grows, include a GOVERNANCE directive with leadership defined, community engagement, maintainers, expected maturity or growth patterns, and commitment or non-commitment to open governance, and architecture decisions.

1. Have a clearly defined path for a contributor to become a reviewer, and ultimately a maintainer ([EXAMPLE](https://github.com/vmware-tanzu/antrea/blob/main/GOVERNANCE.md#maintainers)).

1. Have a common style guide to use for documentation ([EXAMPLE](https://github.com/vmware-tanzu/velero/blob/main/site/content/docs/v1.5/style-guide.md)).

1. Have comprehensive end-to-end and unit testing using public CI/CD systems such as CircleCI or GitHub Actions.

1. Use inclusive terminology throughout your project and docs. For example, the default branch in Github should be `main` instead of `master`. Refer to the inclusive terminology list for more examples.

## Build and Increase our credibility

1. Deliver on features that are set in the roadmap - stay focused.

1. Contribute to related open source projects, tying yourself and your project to another part of the open source ecosystem - don’t work in a bubble.

1. Be public and be seen, by writing blog posts, speaking at events, hosting open community calls, engaging with the community on Slack, being a part of podcasts, or hosting IRL Community Meetings.

## Engage with users and contributors

1. When getting users who want to publicly state that they are using your project, utilize an ADOPTERS file ([EXAMPLE](https://github.com/vmware-tanzu/velero/blob/main/ADOPTERS.md)).

1. Be engaged and open in your project Slack channel, keep project discussions as much as possible in there and not in the internal VMware Slack.

1. Don’t immediately jump on fixing the small and easy things - create and use the “good first issue” and “help wanted” tags to give new contributors an easy way into the project.

1. Hold regular community meetings through Zoom, explaining finished work, what’s currently on deck, and roadmap items that should be discussed with the community - we’ve had success with running these weekly instead of monthly.

1. Be cognizant of time zones, not everyone is online when you are - make sure the work you do supports an asynchronous community.

1. Establish a Twitter profile ([EXAMPLE](https://twitter.com/projectvelero)) or Twitter outlet for each open source project - this could be a separate account, or amplified messaging through @vmwopensource and others.

1. Submit CFPs and be active at relevant conferences, events, and meetups - there is more we can do than just be speakers, for example lead live community meetings, workshops, and code jams at these events.

## Evolve perception of VMware

1. Ensure a complementary presence in existing communities such as Kubernetes Special Interest Groups (SIGs) and Working Groups (WG).

1. Ensure that blogs, CFPs and tweets are community first, not VMware first.
