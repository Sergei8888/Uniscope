FROM node:18-alpine as uniscope-base

ENV PUPPETEER_SKIP_DOWNLOAD=true

WORKDIR /usr/src/uniscope

RUN apk add --no-cache libc6-compat
RUN npm install pnpm@8.6.6 -g
RUN npm install turbo@1.9.3 -g

COPY . .

RUN pnpm install -r --frozen-lockfile && \
	rm -rf ~/.pnpm-store;

# Build shared package
WORKDIR /usr/src/uniscope/packages/shared
RUN pnpm run build
