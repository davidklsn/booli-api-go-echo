// Custom styling
import "../css/main.css";
import "../css/json_view.css"
import "../css/json_editor.css"


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

  var jsonInput = document.getElementById("jsonInput");
  if (jsonInput) {
    const rawData = jsonInput.getAttribute("json");
    const elem = document.getElementById("jsonInput");


    if (elem && rawData) {
      const jsonData: object = JSON.parse(rawData);
      
      // @ts-ignore
      const editor = new JSONEditor(elem, {})
      editor.set(jsonData)
      editor.expandAll()
    }
  }
};
