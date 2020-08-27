#!/bin/bash
VERSION=v0.0.1
docker build -t=push.soh.re/ghreview:$VERSION .
docker tag push.soh.re/ghreview:$VERSION push.soh.re/ghreview:latest
docker push push.soh.re/ghreview:$VERSION
docker push push.soh.re/ghreview:latest
