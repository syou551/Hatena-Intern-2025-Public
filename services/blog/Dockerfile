# syntax=docker/dockerfile:1

FROM node:22-slim AS node-builder
RUN corepack enable

WORKDIR /services/blog
RUN --mount=type=bind,source=package.json,target=package.json \
  --mount=type=bind,source=yarn.lock,target=yarn.lock \
  --mount=type=bind,source=.yarnrc.yml,target=.yarnrc.yml \
  --mount=type=bind,source=.yarn,target=.yarn,readwrite \
  --mount=type=cache,id=yarn,target=/root/.yarn/berry/cache \
  yarn install --immutable
COPY package.json yarn.lock .yarnrc.yml tsconfig.json vite.config.ts ./
COPY .yarn/ ./.yarn/
COPY src/ ./src/
COPY templates/ ./templates/
COPY images/ ./images/
COPY assets/ ./assets/
# publicにファイルを設置した場合コメントを消す
# COPY public/ ./public/
RUN yarn build


FROM golang:1.24-bookworm AS go-builder

WORKDIR /services/blog
RUN --mount=type=bind,source=go.mod,target=go.mod \
    --mount=type=bind,source=go.sum,target=go.sum \
    go mod download
COPY . .
RUN --mount=type=cache,target=/root/.cache/go-build make build


FROM golang:1.24-bookworm

WORKDIR /services/blog
COPY --from=node-builder /services/blog/dist ./dist
COPY --from=node-builder /services/blog/images ./dist/images
COPY --from=node-builder /services/blog/assets ./dist/assets
COPY --from=go-builder /services/blog/bin/server ./bin/server

USER 1000

ENTRYPOINT ["./bin/server"]
