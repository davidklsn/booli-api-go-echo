// Custom styling
import "../css/main.css";
import "../css/json_view.css"


import "htmx.org";
import { prettyPrintJson, FormatOptions } from "pretty-print-json";

window.onload = function () {
  // Your code goes here
  var jsonView = document.getElementById("jsonView");
  if (jsonView) {
    const rawData = jsonView.getAttribute("json");
    const elem = document.getElementById("jsonView");

    if (elem && rawData) {
      const jsonData: object = JSON.parse(rawData);
      const options: FormatOptions = { linkUrls: true };
      elem.innerHTML = prettyPrintJson.toHtml(jsonData, options);
    }
  }
};
