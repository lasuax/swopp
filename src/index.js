import React from "react";
import App from "./App";
import Kafalik from "./components/Kafalik";
import { BrowserRouter } from "react-router-dom";
import ReactDOM from "react-dom";

ReactDOM.render(
  <BrowserRouter>
    <Kafalik />
    <App />
  </BrowserRouter>,
  document.getElementById("root")
);
