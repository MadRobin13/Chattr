// App.js
import React, { Component } from "react";
import "./App.css";
import { connect, sendMessage } from "./api/index";
import Header from "./components/Header/Header";
import ChatHistory from "./components/Header/ChatHistory/ChatHistory";

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
    console.log("hello");
    sendMessage("hello");
  }

  render() {
    return (
      <div className="App">
        <Header />
        <ChatHistory chatHistory={this.state.chatHistory} />
        <button onClick={this.send}>Hit</button>
      </div>
    );
  }
}

export default App;