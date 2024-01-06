FROM node:alpine AS builder

WORKDIR /app
COPY package*.json .

RUN yarn install
COPY . .
RUN yarn run build

FROM nginx:alpine as runner
COPY --from=builder /app/build /usr/share/nginx/html
EXPOSE 80
CMD ["nginx", "-g", "daemon off;"]
