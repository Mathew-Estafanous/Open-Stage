FROM node:16 AS build
COPY package.json .
COPY public ./public
COPY src ./src
RUN npm install
ARG REACT_APP_API_URL
RUN npm run build

FROM node:16-alpine
COPY --from=build ./build .
RUN npm install -g serve
CMD serve -s .