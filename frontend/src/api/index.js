/* eslint-disable no-restricted-globals */
var socket = new WebSocket("ws://localhost:8080/ws");

let connect = (cb) => {
    console.log("Attempting to connect");

    socket.onopen = () => {
        console.log("successfully connected");
    };

    socket.onmessage = (msg) => {
        console.log(msg);
        cb(msg);
    };

    socket.onclose = (event) => {
        console.log("Closed Connection", event);
    };

    socket.onerror = (error) => {
        console.error("Socket Error: ", error);
    }
};

let sendMessage = (msg) => {
    console.log("sending: ", msg);
    socket.send(msg);
}

export {connect, sendMessage};