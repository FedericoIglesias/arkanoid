import { ChangeEvent, useContext } from "react";
import styled from "styled-components";
import { ValueContext } from "../Context/valuesContext";
import { Context } from "../vite-env";

export const Header = () => {
  const Header = styled.section`
    padding: 5px;
    /* max-height: 5vh; */
    display: flex;
    justify-content: space-between;
    div{
      display: flex;
    }
  `;

  const { values, setValues } = useContext<any>(ValueContext);

  const handlerChange = (
    e: ChangeEvent<HTMLSelectElement> | ChangeEvent<HTMLInputElement>
  ) => {
    setValues((prevState: Context) => ({
      ...prevState,
      language: Number(e.target.value),
    }));
  };

  return (
    <Header>
      <h1>Finance App</h1>
      <p>Iglesias Federico</p>
      <div>
        <p>Log Out</p>
        <select onChange={(e) => handlerChange(e)} value={values.language}>
          <option value={0}>English</option>
          <option value={1}>Español</option>
        </select>
      </div>
    </Header>
  );
};
