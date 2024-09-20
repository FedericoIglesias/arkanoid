import { Dashboard } from "./pages/dashboard";
import "./index.css";
import { History } from "./pages/history";
import { ValuesContext } from "./Context/valuesContext";
import {
  createBrowserRouter,
  createRoutesFromElements,
  Route,
  RouterProvider,
} from "react-router-dom";
import { Login } from "./pages/login";

export const App = () => {
  const router = createBrowserRouter(
    createRoutesFromElements(
      <>
        <Route path="/dashboard" element={<Dashboard />} />
        <Route path="/history" element={<History />} />
        <Route path="/login" element={<Login />} />
        <Route path="/" element={<Login />} />
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
