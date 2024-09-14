import styled from "styled-components";



export const Header = () => {
  const Header = styled.section`
  padding: 5px;
  max-height: 5vh;
  display: flex;
`;

  return (
    <Header>
      <h1>Finance App</h1>
      <p>Iglesias Federico</p>
      <p>Log Out</p>
      <select>
        <option value={0}>English</option>
        <option value={1}>Español</option>
      </select>
    </Header>
  );
};
