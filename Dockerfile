FROM scratch
COPY twilioagg-linux-amd64 /bin/twilioagg-linux-amd64
ENTRYPOINT ["/bin/twilioagg-linux-amd64"]
