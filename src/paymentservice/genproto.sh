#!/bin/bash -eu
#


# [START gke_paymentservice_genproto]

# protos are loaded dynamically for node, simply copies over the proto.
mkdir -p proto
cp -r ../../pb/* ./proto

# [END gke_paymentservice_genproto]