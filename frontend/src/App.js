// App.js
import React, { Component } from "react";
import "./App.css";
import { connect, sendMessage } from "./api/index";
import Header from "./components/Header/Header";
import ChatHistory from "./components/ChatHistory/ChatHistory";
import ChatInput from "./components/ChatInput/ChatInput";

class App extends Component {
  constructor(props) {
    super(props);
    connect();
    this.state = {
      chatHistory: []
    };
  }

  componentDidMount() {
    connect((msg) => {
      console.log("New Message");
      this.setState(prevState => ({
        chatHistory: [...this.state.chatHistory, msg]
      }))
      console.log(this.state);
    }); 
  }

  send() {
    if (KeyboardEvent.keyCode  === 13) {
      sendMessage(currentTarget.value);
    }
    console.log("hello");
    sendMessage("hello");
  }

  render() {
    return (
      <div className="App">
        <link rel="icon" href="Chattr_icon_v1.png"/> 
        <Header />
        <ChatHistory chatHistory={this.state.chatHistory} />
        <ChatInput send={this.send} />
      </div>
    );
  }
}

export default App;