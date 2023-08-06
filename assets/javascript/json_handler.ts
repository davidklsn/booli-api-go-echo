import { prettyPrintJson, FormatOptions } from "pretty-print-json";

export const checkForJson = function (document: Document) {
  console.log("dallas");
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
// Your code goes here

export const checkForJsonEditor = function (document: Document) {
  var jsonInput = document.getElementById("jsonInput");
  if (jsonInput) {
    const rawData = jsonInput.getAttribute("json");
    const elem = document.getElementById("jsonInput");

    if (elem && rawData) {
      const jsonData: object = JSON.parse(rawData);

      // @ts-ignore
      const editor = new JSONEditor(elem, {});
      editor.set(jsonData);
      editor.expandAll();
    }
  }
};

// @ts-ignore
const postEditedJson = function (event: any) {
  console.log("dallassallad", event);
};
