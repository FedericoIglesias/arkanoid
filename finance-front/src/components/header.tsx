import { ChangeEvent, useContext, useEffect } from "react";
import styled from "styled-components";
import { ValueContext } from "../Context/valuesContext";
import { Context } from "../vite-env";

export const Header = () => {
  const Header = styled.section`
    padding: 5px;
    max-height: 5vh;
    display: flex;
  `;

  const { values, setValues } = useContext<any>(ValueContext);

  const handlerChange = (
    e: ChangeEvent<HTMLSelectElement> | ChangeEvent<HTMLInputElement>
  ) => {
    // e.preventDefault();
    // values.language = Number(e.target.value)
    // console.log(values);
    setValues((prevState: Context) => ({
      ...prevState,
      language: Number(e.target.value),
    }));
  };

  return (
    <Header>
      <h1>Finance App</h1>
      <p>Iglesias Federico</p>
      <p>Log Out</p>
      <select onChange={(e) => handlerChange(e)} value={values.language}>
        <option value={0}>English</option>
        <option value={1}>Español</option>
      </select>
    </Header>
  );
};
