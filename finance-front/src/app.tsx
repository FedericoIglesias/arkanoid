import { createContext, useState } from "react";
import { Dashboard } from "./pages/dashboard";
import "./index.css";
import { History } from "./pages/history";
import { Context } from "./vite-env";
import { ValuesContext } from "./Context/valuesContext";

export const App = () => {
  return (
    <>
      <ValuesContext>
        <Dashboard />
      </ValuesContext>
    </>
  );
};
// {/* <History /> */}
