import React from "react";

import "./Video.css";
import { IDanmuItem, IDanmuBroadCastItem } from "./types/type";
import { sendDanmuToServer } from "./network/sendDanmu";
import { connectToWebsocket } from "./network/websocket";
import {
  FormControl,
  InputLabel,
  OutlinedInput,
  InputAdornment,
  IconButton,
  FormControlLabel,
  Switch,
} from "@material-ui/core";

import TelegramIcon from "@material-ui/icons/Telegram";
import ColorizeIcon from "@material-ui/icons/Colorize";

const VIDRO_URL =
  "https://s1.pstatp.com/cdn/expire-1-M/byted-player-videos/1.0.0/xgplayer-demo.mp4";

class Video extends React.Component {
  state = {
    danmuContent: "",
    danmuColor: "#ff0000",

    uid: "1",
    rid: "1",

    danmuOpen: false,
  };

  danmuContainerRef: HTMLDivElement | null = null;
  colorPickerButtonRef: HTMLInputElement | null = null;

  addDanmu = (danmu: IDanmuBroadCastItem) => {
    requestAnimationFrame(() => {
      const danmuNode = document.createElement("div");
      danmuNode.classList.add("danmu-item");
      danmuNode.onanimationend = () => {
        console.log("removed ", danmu);
        requestAnimationFrame(() => {
          this.danmuContainerRef?.removeChild(danmuNode);
        });
      };
      danmuNode.style.transform = `translateY(${~~(Math.random() * 550)}px)`;
      danmuNode.style.color = danmu.color;
      //   danmuNode.style.willChange = "transform";
      danmuNode.innerText = danmu.content;
      this.danmuContainerRef?.appendChild(danmuNode);

      console.log("added ", danmu);
    });
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
    const { uid, rid } = this.state;

    const danmu: IDanmuItem = {
      id: new Date().getTime().toString(),
      content: this.state.danmuContent,
      color: this.state.danmuColor,
    };

    sendDanmuToServer({
      ...danmu,
      uid: +uid,
      rid: +rid,
    });
    console.log("send danmu", danmu);
  };

  sendDanmuDirectly = () => {
    const danmu: IDanmuItem = {
      id: new Date().getTime().toString(),
      content: this.state.danmuContent,
      color: this.state.danmuColor,
    };
    this.addDanmu(danmu);
  };

  socket: WebSocket | null = null;

  openWs = () => {
    const { uid, rid } = this.state;
    this.socket = connectToWebsocket(
      +uid,
      +rid,
      (data) => {
        this.addDanmu(data);
      },
      () => {
        this.socket = null;
        this.setState({
          danmuOpen: false,
        });
      }
    );
    console.log("new ws", this.socket);
  };

  closeWs = () => {
    this.socket?.close();
    this.socket = null;
    console.log("close ws");
  };

  selectColor = () => {
    this.colorPickerButtonRef?.click();
  };

  componentWillUnmount() {
    this.closeWs();
  }

  componentDidMount() {
    // this.openWs();
  }

  handleDanmuToggle = () => {
    const { danmuOpen } = this.state;
    if (danmuOpen) {
      this.closeWs();
    } else {
      this.openWs();
    }

    this.setState({
      danmuOpen: !danmuOpen,
    });
  };

  render() {
    const { danmuContent, danmuColor, uid, rid, danmuOpen } = this.state;

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

        <div className="send-danmu-container">
          <FormControl className="form-control" variant="outlined">
            <InputLabel>弹幕内容</InputLabel>
            <OutlinedInput
              type="text"
              value={danmuContent}
              onChange={this.handleChange("danmuContent")}
              style={{ color: danmuColor }}
              endAdornment={
                <InputAdornment position="end">
                  <IconButton onClick={this.selectColor} edge="end">
                    <input
                      type="color"
                      style={{
                        position: "absolute",
                        opacity: 0,
                        right: "36px",
                      }}
                      onChange={this.handleChange("danmuColor")}
                      value={danmuColor}
                      ref={(elem) => (this.colorPickerButtonRef = elem)}
                    />
                    <ColorizeIcon />
                  </IconButton>

                  <IconButton onClick={this.sendDanmu} edge="end">
                    <TelegramIcon />
                  </IconButton>
                </InputAdornment>
              }
              labelWidth={80}
            />
          </FormControl>
        </div>
        <div className="danmu-switcher-container">
          <FormControl
            variant="outlined"
            style={{ opacity: danmuOpen ? 0 : 1 }}
          >
            <InputLabel>uid</InputLabel>
            <OutlinedInput
              type="text"
              value={uid}
              onChange={this.handleChange("uid")}
              labelWidth={60}
            />
          </FormControl>

          <FormControl
            variant="outlined"
            style={{ opacity: danmuOpen ? 0 : 1 }}
          >
            <InputLabel>rid</InputLabel>
            <OutlinedInput
              type="text"
              value={rid}
              onChange={this.handleChange("rid")}
              labelWidth={60}
            />
          </FormControl>

          <FormControlLabel
            control={
              <Switch
                checked={danmuOpen}
                onChange={this.handleDanmuToggle}
                color="primary"
              />
            }
            label="开启弹幕"
          />
        </div>
        <div>
          <button onClick={this.sendDanmuDirectly}>[测试] 发送本地弹幕</button>
        </div>
      </div>
    );
  }
}

export default Video;
