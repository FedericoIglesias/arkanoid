import { ChangeEvent } from "react";

const tags = {
  AMOUNT: ["Amount", "Monto"],
  DESCRIPTION: ["Description", "Descripción"],
  CATEGORY: ["Category", "Categoria"],
  FLOW: ["Flow", "Flujo"],
  INCOME: ["Income", "Ingreso"],
  OUTFLOW: ["Outflow", "Gasto"],
  SUBMIT: ["Submit", "Enviar"],
};

export const Form = () => {
  const lenguage = 0;

  const transaction = {
    Amount: 0,
    Description: "",
    CategoryID: 1,
    FlowID: tags.INCOME[0].toLocaleLowerCase(),
  };

  const handleSubmit = (event: { preventDefault: () => void }) => {
    event.preventDefault();
    console.log(transaction);
  };

  const handleChange = (
    e: ChangeEvent<HTMLInputElement> | ChangeEvent<HTMLSelectElement>
  ) => {
    console.log(e.target.value);
    console.log(e.target.name);
    switch (e.target.name) {
      case tags.AMOUNT[0]:
        transaction.Amount = Number(e.target.value);
        break;
      case tags.CATEGORY[0]:
        transaction.CategoryID = Number(e.target.value);
        break;
      case tags.DESCRIPTION[0]:
        transaction.Description = e.target.value;
        break;
      case tags.FLOW[0]:
        transaction.FlowID = e.target.value.toLocaleLowerCase();
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
