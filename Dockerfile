FROM alpine

WORKDIR /app
COPY ./main .

EXPOSE 8081
ENTRYPOINT ["./main"]