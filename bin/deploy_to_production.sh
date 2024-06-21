#!/bin/bash

gcloud builds submit . --config=cloudbuild.yaml --project jiikko
