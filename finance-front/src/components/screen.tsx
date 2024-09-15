import { PALETTE2 } from "../const";
import { Header } from "./header";
import { Panel } from "./panel";


export const Screen = ({ children }: { children: JSX.Element }) => {
  return (
    <>
      <main className="bodyScreen">
        <section
          style={{
            gridArea: "header",
            backgroundColor: `${PALETTE2.SECONDARY}`,
          }}
        >
          <Header></Header>
        </section>
        <section
          style={{ gridArea: "aside", backgroundColor: `${PALETTE2.PRIMARY}` }}
        >
          <Panel></Panel>
        </section>
        <section style={{ gridArea: "main", backgroundColor: `#222` }}>
          {children}
        </section>
      </main>
    </>
  );
};
