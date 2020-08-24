import React from "react";
import "./App.css";
import { IDanmuItem, IDanmuBroadCastItem } from "./types/type";
import { sendDanmuToServer } from "./network/sendDanmu";
import { connectToWebsocket } from "./network/websocket";

const VIDRO_URL =
  "https://s1.pstatp.com/cdn/expire-1-M/byted-player-videos/1.0.0/xgplayer-demo.mp4";

class App extends React.Component {
  state = {
    danmuContent: "",
    danmuColor: "#ff0000",

    uid: "1",
    rid: "1",
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

  handleChange = (key: string) => (
    event: React.ChangeEvent<HTMLInputElement>
  ) => {
    const newVal = event.target.value;
    this.setState({
      [key]: newVal,
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

  openWs = () => {
    const { uid, rid } = this.state;
    this.socket = connectToWebsocket(+uid, +rid, (data) => {
      this.addDanmu(data)
    });
    console.log("new ws", this.socket);
  };

  componentWillUnmount() {
    this.socket?.close();
    console.log("close ws");
  }

  componentDidMount() {
    // this.openWs();
  }

  render() {
    const { danmuContent, danmuColor, uid, rid } = this.state;

    return (
      <div className="App">
        <div className="video">
          <video controls className="video-component" src={VIDRO_URL}></video>
          <div
            className="danmu-container"
            id="danmu-container"
            ref={(ref) => {
              this.danmuContainerRef = ref;
            }}
          ></div>
        </div>

        <div>
          <input
            type="text"
            onChange={this.handleChange("danmuContent")}
            value={danmuContent}
          />
          <input
            type="color"
            onChange={this.handleChange("danmuColor")}
            value={danmuColor}
          />
          <button onClick={this.sendDanmu}>send danmu</button>
        </div>
        <div>
          <input type="text" onChange={this.handleChange("uid")} value={uid} />
          <input type="text" onChange={this.handleChange("rid")} value={rid} />
          <button onClick={this.openWs}>open ws</button>
        </div>
      </div>
    );
  }
}

export default App;
