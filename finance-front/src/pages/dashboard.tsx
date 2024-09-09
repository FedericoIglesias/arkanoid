import { useState } from "react";
import { Form } from "../components/form";
import { Table } from "../components/table";
import { lenguage, tags } from "../const";

// const transaction = {
//   Amount: 0,
//   Description: "",
//   CategoryID: 1,
//   FlowID: tags.INCOME[0].toLocaleLowerCase(),
// };

export const Dashboard = () => {
  
  const listTags: Array<string> = [
    tags.AMOUNT[lenguage],
    tags.CATEGORY[lenguage],
    tags.DESCRIPTION[lenguage],
    tags.FLOW[lenguage],
  ];

  return (
    <>
      <h1>DASHBOARD</h1>
      <Table listHead={listTags}/>
    </>
  );
};
