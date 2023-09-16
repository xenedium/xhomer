FROM node:18-bookworm as frontend-builder

WORKDIR /app
COPY frontend/package.json frontend/yarn.lock ./
RUN yarn install --frozen-lockfile
COPY frontend/ ./
RUN yarn build

FROM golang:1.21.1-bullseye as backend-builder

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -v -o /app/bin/ ./...

FROM golang:1.21.1-bullseye as runtime

WORKDIR /app
COPY --from=backend-builder /app/bin/xhomer .
COPY --from=frontend-builder /app/dist ./dist

EXPOSE 8080
CMD ["./xhomer"]

