import React, { Component } from "react";
import "./ChatHistory.scss";
import Message from "../Message/Message";
// import ChatInput from "../ChatInput/ChatInput";

class ChatHistory extends Component {
  render() {
    const messages = this.props.chatHistory.map(msg => <Message message={msg.data} />);

    return (
      <div className="ChatHistory">
        <h2>Chat</h2>
        {messages}
        {/* <ChatInput onSend={this.props.onSend} /> */}
      </div>
    );
  }
}

export default ChatHistory;