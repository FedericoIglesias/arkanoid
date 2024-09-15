import { Screen } from "../components/screen";
import { Table } from "../components/table";
import { lenguage, tags } from "../const";

const listTags: Array<string> = [
  tags.AMOUNT[lenguage],
  tags.CATEGORY[lenguage],
  tags.DESCRIPTION[lenguage],
  tags.FLOW[lenguage],
];

export const History = () => {
  return (
    <Screen>
      <>
        <h2>History</h2>
        <Table listHead={listTags} />
      </>
    </Screen>
  );
};
