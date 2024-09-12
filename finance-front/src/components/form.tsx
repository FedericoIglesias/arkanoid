import { ChangeEvent } from "react";
import { lenguage, tags } from "../const";

type Transaction = {
  userID: number;
  amount: number;
  description: string;
  categoryID: number;
  flowID: number;
};

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
        console.log(response);
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
    <form onSubmit={(e) => handleSubmit(e)}>
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
        <option value="income">{tags.INCOME[lenguage]}</option>
        <option value="outflow">{tags.OUTFLOW[lenguage]}</option>
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
    </form>
  );
};
