const express = require("express");
const path = require('path')

const app = express(); // create express app
const port = 8080;
app.use(express.static("build"));

app.use('*', (req, res) => {
    res.sendFile(path.join(__dirname, '/build/index.html'));
});

// start express server on port 5000
app.listen(port, () => {
    console.log(`server started on port: ${port}`);
});