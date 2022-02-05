FROM ruby:3

WORKDIR /home/migrate

ADD ./_migrate/Gemfile ./_migrate/Gemfile.lock /home/migrate
RUN bundle install

ENTRYPOINT ["rake"]
CMD ["db:migrate"]
