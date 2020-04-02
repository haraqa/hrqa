Hrqa
===
[![Go Report Card](https://goreportcard.com/badge/github.com/haraqa/hrqa)](https://goreportcard.com/report/haraqa/hrqa)
[![License](https://img.shields.io/github/license/haraqa/hrqa.svg)](https://github.com/haraqa/hrqa/blob/master/LICENSE)
[![build](https://github.com/haraqa/hrqa/workflows/build/badge.svg)](https://github.com/haraqa/hrqa/blob/master/.github/workflows/go.yml)
[![Docker Build](https://img.shields.io/docker/cloud/build/haraqa/hrqa.svg)](https://hub.docker.com/r/haraqa/hrqa/)
[![Release](https://img.shields.io/github/release/haraqa/hrqa.svg)](https://github.com/haraqa/hrqa/releases)

<h2 align="center">Command Line Haraqa Client</h2>

<div align="center">
  <a href="https://github.com/haraqa/haraqa">
    <img src="https://raw.githubusercontent.com/haraqa/haraqa/media/mascot.png"/>
  </a>
</div>

**hrqa** is a command line tool for interacting with a [haraqa](https://github.com/haraqa/haraqa) broker.

## Install

#### From source:
```
go get github.com/haraqa/hrqa
```

#### From docker
```
docker run haraqa/hrqa --help
```

## Run

```
# create a new topic
hrqa topic create -t my_topic

# list all topics
hrqa topic list

# send a message to a topic
hrqa produce -t my_topic -m "hello there"

# pipe messages to a topic
echo "hello world" | hrqa produce -t my_topic

# get the offsets of a topic
hrqa topic offsets -t my_topic

# consume the latest message from a topic
hrqa consume -t my_topic

# consume from a topic at a specific offset
hrqa consume -t my_topic -o 0

# remove all messages from a topic
hrqa topic delete -t my_topic
```

## Contributing

We want this project to be the best it can be and all feedback, feature requests or pull requests are welcome.

## License

MIT © 2019 [haraqa](https://github.com/haraqa/) and [contributors](https://github.com/haraqa/haraqa/graphs/contributors). See `LICENSE` for more information.