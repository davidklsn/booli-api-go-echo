// Custom styling
import "../css/main.css";
import "../css/json_view.css"
import "../css/json_editor.css"


import "htmx.org";

import { checkForJson, checkForJsonEditor } from "./json_handler";

window.onload = function () {
  // JSON handling on pages
  checkForJson(document);
  checkForJsonEditor(document);
};
