import { Modals } from "../components/modals";
import { Screen } from "../components/screen";
import { lenguage, WORDS } from "../const";

export const Dashboard = () => {
  return (
    <main className="bodyScreen">
      <Screen></Screen>
      <section style={{ gridArea: "main" }}>
        <h2>DASHBOARD</h2>
        <Modals tag={WORDS.ADD_RECORD[lenguage]} />
      </section>
    </main>
  );
};
