var socket = new WebSocket("ws://localhost:8080/ws");

let connect = () => {
    socket.connect();
    console.log("Attempting to connect");

    socket.onopen = () => {
        console.log("successfully connected");
    };

    socket.onmessage = (msg) => {
        console.log(msg);
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