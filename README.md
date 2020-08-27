# ghReview
[![Release](https://img.shields.io/github/release/Jmainguy/ghReview.svg?style=flat-square)](https://github.com/Jmainguy/ghReview/releases/latest)

Application to check Github for Pull Requests, where the user is assigned as a reviewer, and ping them on Slack about it.

Will only ping once while app is running, and only if the review is in waiting, not if comments or previous review has been made.

## Deployments
Edit the values in deployments/kubernetes/deployment.yaml to match your env

## Releases
Grab Release from [The Releases Page](https://github.com/Jmainguy/ghReview/releases)

## Build
```/bin/bash
export GO111MODULE=on
go build
```
