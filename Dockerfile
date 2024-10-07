# backend/Dockerfile

# 1단계: 빌드
FROM golang:1.22-alpine AS builder

# 환경 변수 설정
ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

WORKDIR /build

COPY go.mod go.sum main.go ./
COPY . .

RUN ls -al

RUN go mod download

RUN go mod tidy

RUN go build -o main .

WORKDIR /dist

RUN cp /build/main .

# 2단계: 실행
FROM scratch

# 빌드된 바이너리 복사
COPY --from=builder /dist/main .

# 포트 노출
EXPOSE 8080

# 애플리케이션 실행
ENTRYPOINT ["/main"]
