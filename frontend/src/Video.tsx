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

let send_timestamp: { [key: string]: number | undefined } = {};

class Video extends React.Component {
  state = {
    danmuContent: "",
    danmuColor: "#ff0000",

    uid: "1",
    rid: "1",

    // onlines: 0,
    danmuOpen: false,
    // onlineOpen: false,
  };

  danmuContainerRef: HTMLDivElement | null = null;
  colorPickerButtonRef: HTMLInputElement | null = null;

  addDanmu = (danmu: IDanmuBroadCastItem) => {
    const start_timestamp = send_timestamp[danmu.content];
    const now_timestamp = Date.now();
    if (start_timestamp) {
      console.log(
        `[PERFORMANCE] send and receive danmu ${
        now_timestamp - start_timestamp
        }ms`,
        danmu
      );
      delete send_timestamp[danmu.content];
    }

    requestAnimationFrame(() => {
      if (Date.now() - now_timestamp >= 2 * 1000) return; 
      // optimize: if a danmu not in the screen more than 10 seconds, should not be ignored 

      const danmuNode = document.createElement("div");
      danmuNode.classList.add("danmu-item");

      const containerComputedStyle = getComputedStyle(this.danmuContainerRef!);
      const containerWidth = containerComputedStyle.width;
      const containerHeight = +(containerComputedStyle.height.replace('px', ''));

      danmuNode.style.top = `${~~(Math.random() * (containerHeight - 24))}px`;
      danmuNode.style.color = danmu.color;
      danmuNode.innerText = danmu.content;
      this.danmuContainerRef?.appendChild(danmuNode);

      const danmuWidth = getComputedStyle(danmuNode).width;
      const animation = danmuNode.animate(
        [
          {
            transform: `translateX(0px)`,
            offset: 0,
          },
          {
            transform: `translateX(calc(-${danmuWidth} - ${containerWidth}))`,
            offset: 1,
          },
        ],
        {
          duration: 12000,
          easing: 'linear'
        }
      );

      animation.onfinish = () => {
        requestAnimationFrame(() => {
          this.danmuContainerRef?.removeChild(danmuNode);
        });
      };
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

    send_timestamp[danmu.content] = Date.now();
    sendDanmuToServer({
      ...danmu,
      uid: +uid,
      rid: +rid,
      time: Date.now(),
    });
  };

  testLocalPressure = () => {
    setInterval(() => this.sendDanmuDirectly(), 200);
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
    this.onlineHandler && clearInterval(this.onlineHandler);
    send_timestamp = {};
  }

  componentDidMount() {
    // this.openWs();
  }

  onlineHandler: NodeJS.Timeout | undefined = undefined;
  // handleOnlineToggle = () => {
  //   const { onlineOpen, rid } = this.state;
  //   if (onlineOpen) {
  //     this.onlineHandler && clearInterval(this.onlineHandler);
  //     this.onlineHandler = undefined;
  //   } else {
  //     this.onlineHandler = setInterval(async () => {
  //       const res = await getOnlineUser(rid);
  //       const res_json = await res.json();
  //       console.log('onlines: ', res_json.message)
  //       this.setState({
  //         onlines: res_json.message,
  //       });
  //     }, 5000);
  //   }

  //   this.setState({
  //     onlineOpen: !onlineOpen,
  //   });
  // };

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
    const {
      danmuContent,
      danmuColor,
      uid,
      rid,
      danmuOpen,
      // onlineOpen,
      // onlines,
    } = this.state;

    return (
      <div className="App">
        {/* {onlineOpen && <div className="online-container">{onlines} 人在线</div>} */}
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

          {/* <FormControlLabel
            control={
              <Switch
                checked={onlineOpen}
                onChange={this.handleOnlineToggle}
                color="primary"
              />
            }
            label="在线人数"
          /> */}
        </div>
        <div>
          <button id="test-local-danmu-send" onClick={this.sendDanmuDirectly}>
            [测试] 发送本地弹幕
          </button>
          <button id="test-local-pressure" onClick={this.testLocalPressure}>
            [测试] 前端压力性能测试
          </button>
        </div>
      </div>
    );
  }
}

export default Video;
