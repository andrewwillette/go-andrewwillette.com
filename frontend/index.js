const express = require("express");
const path = require('path')

const app = express(); // create express app
const port = 80;
app.use(express.static("build"));

app.use('*', (req, res) => {
    res.sendFile(path.join(__dirname, '/build/index.html'));
});

// start express server on port 5000
//
var server = require('http').createServer(app);
server.listen(8080, "0.0.0.0");
//app.listen(port, () => {
//    console.log(`server started on port: ${port}`);
//});
