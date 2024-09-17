import styled from "styled-components";

const Aside = styled.aside``;

export const Panel = () => {
  return (
    <Aside>
      <p>Write entry</p>
      <ul>
        <li>
          <a href="/dashboard">Dashboard</a>
        </li>
        <li>
          <a href="/history">History</a>
        </li>
      </ul>
    </Aside>
  );
};
