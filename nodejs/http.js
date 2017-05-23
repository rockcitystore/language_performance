/**
 * Created by root on 2017/5/23.
 */
"use strict";

const http = require('http');

const server = http.createServer((req, res) => {
    res.end('This is sparta!!!');
});
server.on('clientError', (err, socket) => {
    socket.end('HTTP/1.1 400 Bad Request\r\n\r\n');
});
server.listen(8000);