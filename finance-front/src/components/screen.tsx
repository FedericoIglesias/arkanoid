import { Header } from "./header";
import { Panel } from "./panel";

export const Screen = ({ children }: { children: JSX.Element }) => {
  return (
    <>
      <main className="bodyScreen">
        <section
          style={{
            gridArea: "header",
            backgroundColor: `#FFF`,
            fontSize: "12px",
          }}
        >
          <Header></Header>
        </section>
        <section
          style={{
            gridArea: "aside",
            backgroundColor: `#FFF`,
            position: "static",
            top: 0,
          }}
        >
          <Panel></Panel>
        </section>
        <section
          style={{
            gridArea: "main",
            backgroundColor: `#EEE`,
            position: "relative",
            top: 0,
            boxShadow: "inset 2px 2px 2px #AAA",
            padding: '10px 10px '
          }}
        >
          {children}
        </section>
      </main>
    </>
  );
};
