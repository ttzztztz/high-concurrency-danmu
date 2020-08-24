import React from "react";
import "./App.css";
import Video from "./Video";

import AppBar from "@material-ui/core/AppBar";
import Toolbar from "@material-ui/core/Toolbar";
import Typography from "@material-ui/core/Typography";
import { ThemeProvider } from "@material-ui/styles";
import CssBaseline from "@material-ui/core/CssBaseline";
import createMuiTheme from "@material-ui/core/styles/createMuiTheme";
import blue from "@material-ui/core/colors/blue";
import purple from "@material-ui/core/colors/purple";

const theme = createMuiTheme({
  palette: {
    primary: blue,
    secondary: purple,
  },
});

class App extends React.Component {
  render() {
    return (
      <ThemeProvider theme={theme}>
        <CssBaseline />
        <AppBar position="fixed">
          <Toolbar variant="dense">
            <Typography variant="h6" color="inherit">
              高性能弹幕系统
            </Typography>
          </Toolbar>
        </AppBar>
        <div className="content-container">
          <Video />
        </div>
      </ThemeProvider>
    );
  }
}

export default App;
