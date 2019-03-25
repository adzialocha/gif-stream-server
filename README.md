# gif-stream-server

A server which awaits webcam image uploads from [gif-stream](https://github.com/adzialocha/gif-stream) clients to generate .gifs and upload them to an AWS S3 bucket. Can be used together with [HOFFNUNG 3000](https://github.com/adzialocha/hoffnung3000).

## Requirements

* Go environment
* AWS S3 instance

## Setup via Heroku

1. Make sure to set the `GO_INSTALL_PACKAGE_SPEC` variable to `./cmd/...` to make sure Heroku builds both separate binaries.

    ```
    heroku config:set GO_INSTALL_PACKAGE_SPEC=./cmd/...
    ```

2. Configure the following environmental variables:

    ```
    AWS_REGION=eu-central-1
    AWS_ACCESS_KEY_ID=
    AWS_SECRET_ACCESS_KEY=
    AWS_BUCKET_NAME=
    ```

3. Install the `Heroku Scheduler` Add On and put up a job which executes `worker` every 10 minutes.

4. Activate both Dynos `web` and `worker` (put them to "ON" under "Configure Dynos").
