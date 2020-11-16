FROM alpine

COPY ./gorecipes /gorecipes

CMD ["./gorecipes"]

