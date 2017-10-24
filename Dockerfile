FROM scratch
COPY helloip /helloip
ENTRYPOINT ["/helloip"]
EXPOSE 8080
