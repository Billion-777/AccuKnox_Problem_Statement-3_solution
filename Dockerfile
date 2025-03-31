FROM golang:alpine
WORKDIR /app
COPY . .
RUN go build -o datetime .
EXPOSE 8083
CMD ["./datetime"]
