# go-rest-client-examples

## Quick Start Golang Module

```
go mod init github.com/earthquakesan/go-rest-client-examples
go mod tidy

# library for REST API requests
go get github.com/imroc/req/v3

# library for testing assertions
go get github.com/stretchr/testify

# library for logging
go get github.com/rs/zerolog/log

# library for yaml demarshalling (reading config)
go get gopkg.in/yaml.v2
```

## Setup Virtual Env for Python

```
sudo apt-get install python3-pip
pip install virtualenvwrapper

cat <<EOF >> ~/.bashrc
export WORKON_HOME=\$HOME/.virtualenvs
export PROJECT_HOME=\$HOME/Devel
export VIRTUALENVWRAPPER_PYTHON=python3
source \$HOME/.local/bin/virtualenvwrapper.sh

export PATH=$PATH:$HOME/.local/bin
EOF

source ~/.bashrc
```

## Start REST Server

```
mkvirtualenv go-rest-client-examples
pip install -r rest-server/requirements.txt
make start-webserver
```

## Run Tests

```
make test
```