import { StrictMode } from "react";
import { createRoot } from "react-dom/client";
import { Dashboard } from "./pages/dashboard";
import "./index.css";
import { History } from "./pages/history";

createRoot(document.getElementById("root")!).render(
  <StrictMode>
    <History />
  </StrictMode>
);
