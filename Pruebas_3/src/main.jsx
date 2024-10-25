import { createRoot } from "react-dom/client";
import { TextProvider } from "./context/TextContext";
import App from "./App";
import "./main.css";

createRoot(document.getElementById("root")).render(
  <TextProvider>
    <App />
  </TextProvider>
);
