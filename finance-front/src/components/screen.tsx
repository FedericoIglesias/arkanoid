import styled from "styled-components";
import { PALETTE2 } from "../const";
import { Header } from "./header";
import { Panel } from "./panel";

const Body = styled.body`
  display: grid;
  grid-template-areas:
    "header header header header header"
    "aside aside main main main main"
    "footer footer footer footer footer";
`;

export const Screen = () => {
  return (
    <>
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
    </>
  );
};
