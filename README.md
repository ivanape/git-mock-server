# git-mock-server

Git mock server based in https://github.com/sosedoff/gitkit.

## Create sample git repositories

```
mkdir my-bare-repos
cd my-bare-repos
git init --bare example01.git
```

## How to run

`docker-compose up --build`

## How to clone a repository

`git clone http://localhost:5000/example01.git`