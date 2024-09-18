import { ChangeEvent, useContext, useEffect, useState } from "react";
import { PALETTE, tags } from "../const";
import styled from "styled-components";
import { Error } from "../error";
import { Category, Context, MoneyType } from "../vite-env";
import { ValueContext } from "../Context/valuesContext";

type Transaction = {
  userID: number;
  amount: number;
  description: string;
  categoryID: number;
  flowID: number;
  moneyTypeID: number;
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

const transaction: Transaction = {
  userID: 2,
  amount: 0,
  description: "",
  categoryID: 1,
  flowID: 2,
  moneyTypeID: 1,
};

export const Form = () => {
  const { values }: { values: Context } = useContext<any>(ValueContext);

  const [listCategory, setListCategory] = useState<[Category] | null>(null);
  const [listMoneyType, setListMoneyType] = useState<[MoneyType] | null>(null);

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
        if (response.status !== 200) {
          alert(Error.RESPONSE[values.language]);
        } else {
          alert("Succes to proccess your request");
        }
      })
      .catch((err) => {
        console.log("Error: " + err);
      });
  };

  useEffect(() => {
    fetch("http://localhost:3000/categories", {
      mode: "cors",
    })
      .then((response) => response.json())
      .then((json) => setListCategory(json.data))
      .catch((error) => {
        console.error("Error al hacer la petición:", error);
      });
    fetch("http://localhost:3000/moneyType", {
      mode: "cors",
    })
      .then((response) => response.json())
      .then((json) => setListMoneyType(json.data))
      .catch((error) => {
        console.error("Error al hacer la petición:", error);
      });
  }, []);

  const handleChange = (
    e:
      | ChangeEvent<HTMLInputElement>
      | ChangeEvent<HTMLSelectElement>
      | ChangeEvent<HTMLTextAreaElement>
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
        transaction.flowID =
          e.target.value == tags.INCOME[values.language] ? 1 : 2;
        break;
      case tags.MONEY_TYPE[0]:
        transaction.description = e.target.value;
        break;
    }
  };

  return (
    <FormStyle
      onSubmit={(e: { preventDefault: () => void }) => handleSubmit(e)}
    >
      <label htmlFor="price">{tags.AMOUNT[values.language]}</label>
      <input
        type="number"
        name={tags.AMOUNT[0]}
        id="transaction"
        onChange={(e) => handleChange(e)}
      />
      <label htmlFor="description">{tags.DESCRIPTION[values.language]}</label>
      <textarea
        rows={4}
        cols={5}
        maxLength={50}
        name={tags.DESCRIPTION[0]}
        id="transaction"
        onChange={(e) => handleChange(e)}
        style={{ resize: "none" }}
      />
      <label htmlFor="flow">{tags.FLOW[values.language]}</label>
      <select
        name={tags.FLOW[0]}
        id="transaction"
        onChange={(e) => handleChange(e)}
      >
        <option value={tags.INCOME[values.language]}>
          {tags.INCOME[values.language]}
        </option>
        <option value={tags.OUTFLOW[values.language]}>
          {tags.OUTFLOW[values.language]}
        </option>
      </select>
      <label htmlFor="category">{tags.CATEGORY[values.language]}</label>
      <select
        name={tags.CATEGORY[0]}
        id="transaction"
        onChange={(e) => handleChange(e)}
      >
        {listCategory?.map((element: Category) => {
          return (
            <option value={element.ID}>
              {element.nameCategory.toUpperCase()}
            </option>
          );
        })}
      </select>
      <label htmlFor="moneyType">{tags.MONEY_TYPE[values.language]}</label>
      <select
        name={tags.MONEY_TYPE[0]}
        id="transaction"
        onChange={(e) => handleChange(e)}
      >
        {listMoneyType?.map((element: MoneyType) => {
          return (
            <option value={element.ID}>
              {element.moneyType.toUpperCase()}
            </option>
          );
        })}
      </select>
      <input type="submit" value={tags.SUBMIT[values.language]} />
    </FormStyle>
  );
};
