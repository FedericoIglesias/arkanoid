import { ChangeEvent } from "react";
import { lenguage, PALETTE, tags } from "../const";
import { red } from "@mui/material/colors";
import styled from "styled-components";

type Transaction = {
  userID: number;
  amount: number;
  description: string;
  categoryID: number;
  flowID: number;
};


const FormStyle = styled.form`
  background-color: ${PALETTE.SECONDARY};
  display: flex;
  flex-direction: column;
  padding: 20px;
  label {
    margin-bottom: 3px;
  }
  input,
  select {
    margin-bottom: 5px;
  }
  input[type="submit" i] {
    margin: 0 auto;
    padding: 4px 7px;
    border-radius: 10px;
  }
  input[type="submit" i]:hover {
    background-color: #6c6c07;
  }
`;

export const Form = () => {
  const transaction: Transaction = {
    userID: 2,
    amount: 0,
    description: "",
    categoryID: 1,
    flowID: 2,
  };

  const handleSubmit = async (event: { preventDefault: () => void }) => {
    event.preventDefault();
    fetch("http://localhost:3000/transaction", {
      method: "POST",
      mode: "cors",
      cache: "no-cache",
      credentials: "same-origin",
      headers: {
        "content-type": "application/json",
      },
      body: JSON.stringify(transaction),
    })
      .then((response) => {
        if (response.status !== 200){
          alert("Error process your request")
        }
      })
      .catch((err) => {
        console.log("Error: " + err);
      });
  };

  const handleChange = (
    e: ChangeEvent<HTMLInputElement> | ChangeEvent<HTMLSelectElement>
  ) => {
    switch (e.target.name) {
      case tags.AMOUNT[0]:
        transaction.amount = Number(e.target.value);
        break;
      case tags.CATEGORY[0]:
        transaction.categoryID = Number(e.target.value);
        break;
      case tags.DESCRIPTION[0]:
        transaction.description = e.target.value;
        break;
      case tags.FLOW[0]:
        transaction.flowID = e.target.value == tags.INCOME[lenguage] ? 1 : 2;
        break;
    }
  };

  return (
    <FormStyle
      onSubmit={(e: { preventDefault: () => void }) => handleSubmit(e)}
    >
      <label htmlFor="price">{tags.AMOUNT[lenguage]}</label>
      <input
        type="number"
        name={tags.AMOUNT[0]}
        id="transaction"
        onChange={(e) => handleChange(e)}
      />
      <label htmlFor="description">{tags.DESCRIPTION[lenguage]}</label>
      <input
        type="text"
        name={tags.DESCRIPTION[0]}
        id="transaction"
        onChange={(e) => handleChange(e)}
      />
      <label htmlFor="flow">{tags.FLOW[lenguage]}</label>
      <select
        name={tags.FLOW[0]}
        id="transaction"
        onChange={(e) => handleChange(e)}
      >
        <option value={tags.INCOME[lenguage]}>{tags.INCOME[lenguage]}</option>
        <option value={tags.OUTFLOW[lenguage]}>{tags.OUTFLOW[lenguage]}</option>
      </select>
      <label htmlFor="category">{tags.CATEGORY[lenguage]}</label>
      <select
        name={tags.CATEGORY[0]}
        id="transaction"
        onChange={(e) => handleChange(e)}
      >
        <option>1</option>
        <option>2</option>
        <option>3</option>
        <option>4</option>
      </select>
      <input type="submit" value={tags.SUBMIT[lenguage]} />
    </FormStyle>
  );
};
