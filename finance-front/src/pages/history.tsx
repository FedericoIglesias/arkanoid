import { useContext } from "react";
import { Screen } from "../components/screen";
import { lenguage, tags } from "../const";
import { ValueContext } from "../Context/valuesContext";
import { Context } from "../vite-env";
import { Table } from "../components/table";
// import { TableGeneral as Table } from "../components/table";

// const listTags: Array<string> = [
//   tags.AMOUNT[lenguage],
//   tags.CATEGORY[lenguage],
//   tags.DESCRIPTION[lenguage],
//   tags.FLOW[lenguage],
// ];

export const History = () => {
  const { values }: { values: Context } = useContext<any>(ValueContext);

  const listTags: Array<string> = [
    tags.AMOUNT[values.language],
    tags.CATEGORY[values.language],
    tags.DESCRIPTION[values.language],
    tags.FLOW[values.language],
    "Date",
  ];

  return (
    <Screen>
      <>
        <Table listHead={listTags} />
      </>
    </Screen>
  );
};
