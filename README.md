# Start hosting server content
* Ensure `go` is available in your `$PATH`.
* Ensure `npm` is available in your `$PATH`.
* With the current directory inside the `scripts` folder, execute the `startProd.sh` script.
* with the current directory inside the `scripts` folder, execute the `initAdmin.sh` script. You will be prompted for a username/password for the websites admin privileges.

# Backend
* REST API writting in go for managing admins and hosted music.

# Frontend
* Single page react app written in Typescript, served with an express.js server. Has a dependency on the above REST API for displaying music.
