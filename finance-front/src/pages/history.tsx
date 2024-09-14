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
    <main className="bodyScreen">
      <Screen></Screen>
      <section style={{ gridArea: "main" }}>
        <h2>History</h2>
        <Table listHead={listTags} />
      </section>
    </main>
  );
};
