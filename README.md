[![Codacy Badge](https://api.codacy.com/project/badge/Grade/b07a461e6a9a48fc84226baefff06423)](https://www.codacy.com/app/OpenDevSecOps/go-post2slack?utm_source=github.com&amp;utm_medium=referral&amp;utm_content=opendevsecops/go-post2slack&amp;utm_campaign=Badge_Grade)
[![Follow on Twitter](https://img.shields.io/twitter/follow/opendevsecops.svg?logo=twitter)](https://twitter.com/opendevsecops)

# go-post2slack

Simple utility to post messages to slack.

## Rationale

Although this task can be done with cURL, this command is meant to save you some troubles with the API and future-proof it.

## Getting Started

To install post2slack simply do:

```sh
$ go get github.com/opendevsecops/go-post2slack
```

Once the command is installed in your home go/bin folder execute it like this:

```sh
$ ~/go/bin/go-post2slack --help
```

To send message to slack simply do:

```sh
$ ~/go/bin/go-post2slack --url "https://hooks.slack.com/services/T00000000/B00000000/XXXXXXXXXXXXXXXXXXXXXXXX" --text "Your text message!"
```
