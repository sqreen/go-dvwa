FROM golang AS build
WORKDIR /app
COPY . .
RUN make dvwa

FROM debian:stable-slim
RUN apt-get update && apt-get install -y ca-certificates
RUN mkdir -p /app
COPY --from=build /app/dvwa /app/
COPY --from=build /app/template /app/template
CMD /app/dvwa