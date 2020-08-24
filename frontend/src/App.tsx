import React from "react";
import "./App.css";
import { IDanmuItem, IDanmuBroadCastItem } from "./types/type";
import { sendDanmuToServer } from "./network/sendDanmu";
import { connectToWebsocket } from "./network/websocket";

const VIDRO_URL = 'https://s1.pstatp.com/cdn/expire-1-M/byted-player-videos/1.0.0/xgplayer-demo.mp4'

class App extends React.Component {
  state = {
    danmuContent: "",
    danmuColor: "#ff0000",
  };
  danmuContainerRef: HTMLDivElement | null = null;

  addDanmu = (danmu: IDanmuBroadCastItem) => {
    const danmuNode = document.createElement("div");
    danmuNode.classList.add("danmu-item");
    danmuNode.onanimationend = () => {
      console.log("removed ", danmu);
      this.danmuContainerRef?.removeChild(danmuNode);
    };
    danmuNode.style.transform = `translateY(${~~(Math.random() * 550)}px)`;
    danmuNode.style.color = danmu.color;
    danmuNode.innerText = danmu.content;
    this.danmuContainerRef?.appendChild(danmuNode);

    console.log("added ", danmu);
  };

  handleContentChange = (event: React.ChangeEvent<HTMLInputElement>) => {
    const newVal = event.target.value;
    this.setState({
      danmuContent: newVal,
    });
  };

  handleColorChange = (event: React.ChangeEvent<HTMLInputElement>) => {
    const newVal = event.target.value;
    this.setState({
      danmuColor: newVal,
    });
  };

  sendDanmu = () => {
    const danmu: IDanmuItem = {
      id: new Date().getTime().toString(),
      content: this.state.danmuContent,
      color: this.state.danmuColor,
    };

    sendDanmuToServer({
      ...danmu,
      uid: 1,
      rid: 1,
    });
    console.log("send danmu", danmu);
  };

  socket: WebSocket | null = null;

  componentWillUnmount() {
    this.socket?.close();
    console.log('close ws')
  }

  componentDidMount() {
    this.socket = connectToWebsocket(1, 1, (data) => {
      console.log(data);
    });
    console.log('new ws', this.socket)
  }

  render() {
    const { danmuContent, danmuColor } = this.state;

    return (
      <div className="App">
        <div className="video">
          <video
            controls
            className="video-component"
            src={VIDRO_URL}
          ></video>
          <div
            className="danmu-container"
            id="danmu-container"
            ref={(ref) => {
              this.danmuContainerRef = ref;
            }}
          ></div>
        </div>

        <input
          type="text"
          onChange={this.handleContentChange}
          value={danmuContent}
        />
        <input
          type="color"
          onChange={this.handleColorChange}
          value={danmuColor}
        />
        <button onClick={this.sendDanmu}>send danmu</button>
      </div>
    );
  }
}

export default App;
