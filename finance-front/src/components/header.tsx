import { ChangeEvent, useContext } from "react";
import styled from "styled-components";
import { ValueContext } from "../Context/valuesContext";
import { Context } from "../vite-env";

export const Header = () => {
  const Header = styled.section`
    padding: 2px;
    display: flex;
    justify-content: space-between;
    align-items: center;
    div {
      display: flex;
    }
    select {
      margin-left: 2px;
      font-size: 10px;
      padding-bottom: 1px;
      border: 1px solid #aaa;
      border-radius: 10px;
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
      <select onChange={(e) => handlerChange(e)} value={values.language}>
        <option value={0}>En</option>
        <option value={1}>Es</option>
      </select>
      <p>Log Out</p>
    </Header>
  );
};
