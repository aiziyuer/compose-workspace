FROM ruby:2.7

# 参考: https://blog.justwe.site/post/docker-arg-env/
ARG version=0.4.0
ENV VERSION=${version}

RUN \
  gem install redis-dump -v ${VERSION}