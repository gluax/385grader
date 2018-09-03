<<<<<<< HEAD
FROM node:alpine
=======
FROM node:slim
>>>>>>> a2e735d6e1b103c9ce97d2595aa32951cd806a9d

RUN apt-get update && apt-get install -y \
    build-essential \
    software-properties-common \
    bc \
    expect \
    && mkdir grader

WORKDIR grader

COPY . .

<<<<<<< HEAD
RUN npm i

ENTRYPOINT ["npm", "grade"]
=======
RUN npm i && npm start
>>>>>>> a2e735d6e1b103c9ce97d2595aa32951cd806a9d
