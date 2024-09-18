import styled from "styled-components";
import { WORDS } from "../const";
import { Context } from "../vite-env";
import { FormEvent, useContext, useState } from "react";
import { ValueContext } from "../Context/valuesContext";

const Form = styled.div`
  background-color: #ddd;
  padding: 10px;
  display: flex;
  flex-direction: column;
  justify-content: center;
  label,
  input {
    margin-bottom: 5px;
  }
`;
export const FormCategory = () => {
  const { values }: { values: Context } = useContext<any>(ValueContext);

  const [category, setCategory] = useState<string>("");

  const handlerSubmit = async (e: FormEvent<HTMLButtonElement>) => {
    e.preventDefault();
    if (category == "") {
      alert("Complet the field");
      return;
    }
    console.log(category);
    fetch("http://localhost:3000/category", {
      method: "POST",
      mode: "cors",
      cache: "no-cache",
      credentials: "same-origin",
      headers: {
        "content-type": "application/json",
      },
      body: JSON.stringify({ nameCategory: category }),
    })
      .then((response) => {
        console.log(response.statusText);
        response.status != 201
          ? alert("Unsucces to process your request")
          : null;
      })
      .catch((e) => {
        alert(e);
      });
  };

  return (
    <Form>
      <label>{WORDS.CATEGORY[values.language]}:</label>
      <input type="text" onChange={(e) => setCategory(e.target.value)} />
      <button onClick={(e) => handlerSubmit(e)}>
        {WORDS.SEND[values.language]}
      </button>
    </Form>
  );
};
