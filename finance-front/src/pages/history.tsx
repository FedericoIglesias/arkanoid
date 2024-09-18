import { useContext } from "react";
import { Screen } from "../components/screen";
import { tags, WORDS } from "../const";
import { ValueContext } from "../Context/valuesContext";
import { Context } from "../vite-env";
import { Table } from "../components/table";
import { Form } from "../components/form";
import { Modals } from "../components/modals";
import { FormCategory } from "../components/formCategory";

export const History = () => {
  const { values }: { values: Context } = useContext<any>(ValueContext);

  const listTags: Array<string> = [
    tags.AMOUNT[values.language],
    tags.CATEGORY[values.language],
    tags.DESCRIPTION[values.language],
    tags.FLOW[values.language],
    tags.DATE[values.language],
    tags.MONEY_TYPE[values.language]
  ];

  return (
    <Screen>
      <>
        <div style={{ display: "flex", justifyContent: "space-between" }}>
          <Modals tag={WORDS.ADD_RECORD[values.language]}>
            <Form />
          </Modals>
          <Modals tag={WORDS.ADD_CATEGORY[values.language]}>
            <FormCategory />
          </Modals>
        </div>
        <Table listHead={listTags} />
      </>
    </Screen>
  );
};
