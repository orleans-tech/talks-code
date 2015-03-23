var express = require('express');
var path = require('path');
var app = express();
var http = require('http').Server(app);
var io = require('socket.io')(http);

app.use(express.static(path.join(__dirname)));

app.get('/', function(req, res){
    res.sendFile(__dirname + '/index.html');
});

var numberOfConnections = 0;
var messageId = 0;
var emitNumberOfConnections = function() {
    io.emit('number_connections', numberOfConnections);
};

io.on('connection', function(socket) {
    numberOfConnections++;
    emitNumberOfConnections();

    socket.on('chat message', function(msg) {
        messageId++;
        console.log('message: ' + msg);
        io.emit('chat message', {
            message: msg,
            id: messageId
        });
    });

    socket.on('disconnect', function() {
        numberOfConnections--;
        emitNumberOfConnections();
    });
});

http.listen(8888, function(){
    console.log('listening on *:3000');
});