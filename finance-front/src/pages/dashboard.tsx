import { useContext, useEffect } from "react";
import { Modals } from "../components/modals";
import { Screen } from "../components/screen";
import { lenguage, WORDS } from "../const";
import { ValueContext } from "../Context/valuesContext";
import { Context } from "../vite-env";

// console.log(values);

export const Dashboard = () => {
  const { values } = useContext<any>(ValueContext);
  console.log(values);
useEffect(()=>{},[values])

  return (
    <Screen>
      <>
        <h2>DASHBOARD</h2>
        <Modals tag={WORDS.ADD_RECORD[values.language]} />
      </>
    </Screen>
  );
};
