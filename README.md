# spawn

Spawn is a project scaffolding tool.  It is designed to take an opinion on how to build web and mobile application front and backends.  It relies on a number of open source tools and frameworks to help bootstrap you applications. 

## Supported Platforms
1. Heroku
2. Gitlab
3. Gitlab CI

Spawn relies on tokens from the various supported platforms to create resources on your behalf.  The following documentation can help to get setup on the supported platforms
1. [GitLab](#https://docs.gitlab.com/ee/user/profile/personal_access_tokens.html)
2. [Heroku](https://help.heroku.com/PBGP6IDE/how-should-i-generate-an-api-key-that-allows-me-to-use-the-heroku-platform-api)


## Supported Languages
1. React
2. NodeJs

The supported languages are based off of a set of curated template applications.  They have been designed and built for the purpose of the [Elko](elko.dev).  However, Spawn allows you to bring your own template as well!  Simply override the Elko template and provide your own git repository and let Spawn do the rest!

### Usage
``` bash
./spawn application
```

### Installation
