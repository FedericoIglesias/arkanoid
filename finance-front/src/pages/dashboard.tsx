import { useContext } from "react";
import { Modals } from "../components/modals";
import { Screen } from "../components/screen";
import { WORDS } from "../const";
import { ValueContext } from "../Context/valuesContext";
import { Context } from "../vite-env";
import { Form } from "../components/form";

export const Dashboard = () => {
  const { values }: { values: Context } = useContext<any>(ValueContext);

  return (
    <Screen>
      <>
        <h2>DASHBOARD</h2>
        <Modals tag={WORDS.ADD_RECORD[values.language]}>
          <Form></Form>
        </Modals>
      </>
    </Screen>
  );
};
