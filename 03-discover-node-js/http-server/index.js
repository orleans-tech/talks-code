'use strict';

var express     = require('express'),
    path        = require('path'),
    app         = express(),
    http        = require('http').Server(app),
    io          = require('socket.io')(http);

var numberOfConnections = 0;
var messageId = 0;
var port = process.env.PORT || 5050;

app.use(express.static(path.join(__dirname, 'assets')));

app.get('/', function(req, res){
    res.sendFile(__dirname + '/index.html');
});

var emitNumberOfConnections = function() {
    io.emit('number_connections', numberOfConnections);
};

io.on('connection', function(socket) {
    numberOfConnections++;
    emitNumberOfConnections();

    socket.on('chat message', function(msg) {
        messageId++;
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

http.listen(port, function(){
    console.log('listening on localhost:' + port);
});