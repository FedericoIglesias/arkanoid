import { createContext, useState } from "react";
import { Context } from "../vite-env";

export const ValueContext = createContext({});

export const ValuesContext = ({ children }: { children: JSX.Element }) => {
  const [values, setValues] = useState<Context>({
    language: 0,
    userId: 2,
  });

  const transportData = { values, setValues };

  return (
    <ValueContext.Provider value={transportData}>
      {children}
    </ValueContext.Provider>
  );
};
