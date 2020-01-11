FROM ruby:2.7

# 参考: https://blog.justwe.site/post/docker-arg-env/
ARG VERSION=0.4.0
ENV VERSION=${VERSION}

RUN \
   gem install redis-dump -v ${VERSION}