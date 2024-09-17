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
            height: '6vh'
          }}
        >
          <Header></Header>
        </section>
        <section
          style={{ gridArea: "aside", backgroundColor: `${PALETTE2.PRIMARY}`,position: 'static',top:0 }}
        >
          <Panel></Panel>
        </section>
        <section style={{ gridArea: "main", backgroundColor: `#DDD`, position: 'relative',top:0  }}>
          {children}
        </section>
      </main>
    </>
  );
};
