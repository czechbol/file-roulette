FROM alpine:3.17

# Copy librespeedtest binary
COPY file-roulette /bin/file-roulette

ENTRYPOINT ["/bin/librespeedtest"]