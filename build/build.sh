#!/bin/bash
base=$1
ref=$2

rm -rf .dynamic-env && mkdir -p .dynamic-env && cd .dynamic-env

kustomize create --resources ../deployment/env/"$base"
kustomize edit set image vmerino/go-app:"$ref"
kustomize build . | kubectl apply -f -