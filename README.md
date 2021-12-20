# andrewwillette.com
My personal website which hosts my CV and music recordings.

## Build and host the website
* Ensure `go` is available in your `$PATH`.
* Ensure `npm` is available in your `$PATH`.
* With the current directory inside the `scripts` folder, execute the `startProd.sh` script.
* With the current directory inside the `scripts` folder, execute the `initAdmin.sh` script. You will be prompted for a username/password for the websites admin privileges.

## Backend directory
* REST API written in Go for managing admins and music recordings.

## Frontend directory
* Single page react app written in Typescript, served with an express.js server. Has a dependency on the above REST API for displaying music.
