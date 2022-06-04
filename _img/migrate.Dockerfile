FROM ruby:3

WORKDIR /home/migrate

ADD ./_migrate/Gemfile ./_migrate/Gemfile.lock /home/migrate
RUN bundle install

RUN useradd migrate
USER migrate

ENTRYPOINT ["rake"]
CMD ["db:migrate"]
