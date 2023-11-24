FROM golang:1.18-alpine AS buildStage
WORKDIR /app
COPY . .
RUN go mod tidy
RUN go build -o ./output ./main.go

FROM alpine
WORKDIR /app
COPY --from=buildStage /app/output ./output
COPY ./.env ./.env
COPY ./template_forget_password.html ./template_forget_password.html
COPY ./template_change_password.html ./template_change_password.html
EXPOSE 8080
CMD [ "./output" ]