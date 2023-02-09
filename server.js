const express = require('express');
const bodyParser = require('body-parser');
const app = express();
const http = require('http').Server(app);
const io = require('socket.io')(http);
var WebSocket = require('faye-websocket');
var client = new WebSocket.Client('ws://localhost:8080/ws');
client.on('open', function(message) {
  console.log('Connection established!');
});

client.on('close', function(message) {
  console.log('Connection closed!', message.code, message.reason);
  
  client = null;
});

app.use(bodyParser.urlencoded( {extended: false}))
app.use(bodyParser.json());

function add(a,b){
    return a+b;
}
app.post('/server/sendA', (request, response)=>{
    if(request && request.body && request.body['a']){
        console.log('Sending received data to go-websocket: ',JSON.stringify(request.body));
        client.send(JSON.stringify(request.body));
        response.send("Changed the value for A")
    }
    else{
        response.send("Value for A not given")
    }
})

app.post('/server/sendB', (request, response)=>{
    if(request && request.body && request.body['b']){
        console.log('Sending received data to go-websocket: ',JSON.stringify(request.body));
        client.send(JSON.stringify(request.body));
        response.send("Changed the value for B")
    }
    else{
        response.send("Value for B not given")
    }
})

app.get('/server/sum', (request, response)=>{
    client.send(JSON.stringify({"sum":0}));
    client.on('message', (event)=>{
        console.log(event['data']);
        let resp = Buffer.from(event['data']).toString();
        sum  = JSON.parse(resp).sum;
        response.send(sum.toString())
    });
})

app.listen(8000, '127.0.0.1', ()=>{
    console.log("Server is running...")
});