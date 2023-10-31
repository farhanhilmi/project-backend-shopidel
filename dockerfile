FROM alpine
COPY /output /output
COPY .env .env
EXPOSE 8080
CMD [ "/output" ]