import React, { useState } from "react";
import "./App.css";

interface IDanmuItem {
  id: string;
  content: string;
  color: string;
}

class App extends React.Component {
  state = {
    danmuContent: '',
    danmuColor: 'red'
  }
  danmuContainerRef: HTMLDivElement | null = null;

  addDanmu = (danmu: IDanmuItem) => {
    const danmuNode = document.createElement("div")
    danmuNode.classList.add("danmu-item");
    danmuNode.onanimationend = () => {
      console.log("removed ", danmu);
      this.danmuContainerRef?.removeChild(danmuNode);
    };
    danmuNode.style.transform = `translateY(${~~(Math.random() * 550)}px)`;
    danmuNode.style.color = danmu.color;
    danmuNode.innerText = danmu.content
    this.danmuContainerRef?.appendChild(danmuNode);

    console.log("added ", danmu)
  }

  handleContentChange = (event: React.ChangeEvent<HTMLInputElement>) => {
    const newVal = event.target.value
    this.setState({
      danmuContent: newVal
    })
  }

  handleColorChange = (event: React.ChangeEvent<HTMLInputElement>) => {
    const newVal = event.target.value
    this.setState({
      danmuColor: newVal
    })
  }

  sendDanmu = () => {
    const danmu: IDanmuItem = {
      id: new Date().getTime().toString(),
      content: this.state.danmuContent,
      color: this.state.danmuColor,
    };
    this.addDanmu(danmu);
    console.log("send danmu", danmu);
  };

  render() {
    const { danmuContent, danmuColor } = this.state;

    return (
      <div className="App">
        <div className="video">
          <div
            className="danmu-container"
            id="danmu-container"
            ref={(ref) => {
              this.danmuContainerRef = ref;
            }}
          ></div>
        </div>
        
        <input type="text" onChange={this.handleContentChange} value={danmuContent}/>
        <input type="text" onChange={this.handleColorChange} value={danmuColor}/>
        <button onClick={this.sendDanmu}>send danmu</button>
      </div>
    );
  }
}

export default App;
