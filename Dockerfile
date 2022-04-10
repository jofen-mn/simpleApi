FROM alpine
RUN mkdir -p /opt/jinfa
ADD  jinfa  /opt/jinfa
RUN chmod +x /opt/jinfa/simpleApi

ENTRYPOINT ["/opt/jinfa/simpleApi"]
CMD ["-c=/opt/jinfa/config.yaml"]
EXPOSE 8081
EXPOSE 8082