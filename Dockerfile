# Dockerfile
FROM alpine:latest
COPY ghReview /usr/bin/ghReview
CMD ["/usr/bin/ghReview"]
