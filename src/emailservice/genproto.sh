#!/bin/bash -eu
#


# [START gke_emailservice_genproto]

python -m grpc_tools.protoc -I../../pb --python_out=. --grpc_python_out=. ../../pb/demo.proto

# [END gke_emailservice_genproto]