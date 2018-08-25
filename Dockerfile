FROM node:slim

RUN apt-get update && apt-get install -y \
    build-essential \
    software-properties-common \
    bc \
    expect \
    && mkdir grader

WORKDIR grader

COPY . .

RUN npm i && npm start
