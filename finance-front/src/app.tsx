import { createContext, useState } from "react";
import { Dashboard } from "./pages/dashboard";
import "./index.css";
import { History } from "./pages/history";
import { Context } from "./vite-env";
import { ValuesContext } from "./Context/valuesContext";
import {
  createBrowserRouter,
  createRoutesFromElements,
  Route,
  RouterProvider,
} from "react-router-dom";

export const App = () => {
  const router = createBrowserRouter(
    createRoutesFromElements(
      <>
        <Route path="/dashboard" element={<Dashboard />} />
        <Route path="/history" element={<History />} />
      </>
    )
  );
  return (
    <>
      <ValuesContext>
        <RouterProvider router={router} />
      </ValuesContext>
    </>
  );
};

