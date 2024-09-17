import { useEffect, useState } from "react";
import styled from "styled-components";
import { Selector } from "./selector";

type getRaw = {
  id: number;
  amount: number;
  description: string;
  categoryId: number;
  userId: number;
  flowId: number;
  Date: string;
};

export const Table = ({ listHead }: { listHead: string[] }) => {
  const [data, setData] = useState<any>([]);

  useEffect(() => {
    fetch("http://localhost:3000/transaction/2", {
      mode: "cors",
    })
      .then((response) => response.json())
      .then((json) => setData(json.data))
      .catch((error) => {
        console.error("Error al hacer la petición:", error);
      });
  }, []);

  const [rowsPerPage, setRowsPerPage] = useState(10);
  const [numberPage, setNumberPage] = useState(1);
  const page = Math.ceil(data.length / rowsPerPage);
  const dataRows = data.slice(
    (numberPage - 1) * rowsPerPage,
    numberPage * rowsPerPage
  );
  const Table = styled.section`
    font-size: 10px;
    max-width: 85%;
    max-height: 80vh;
    margin: 0 auto;
    overflow-y: auto;
    div {
      background-color: pink;
      margin-bottom: 2px;
      display: flex;
      border-radius: 10px;
      p {
        width: calc(100% / ${listHead.length});
        padding: 3px;
      }
    }
    div:nth-child(1) {
      top: 0;
      position: sticky;
    }
  `;

  return (
    <>
      <Selector
        setRowsPerPage={setRowsPerPage}
        page={page}
        rowsPerPage={rowsPerPage}
        setNumberPage={setNumberPage}
      />
      <Table>
        <div>
          {listHead.map((tag: string) => {
            return <p key={tag}>{tag}</p>;
          })}
        </div>
        {dataRows.map((row: getRaw) => {
          return (
            <div key={JSON.stringify(row.id)}>
              <p>{row.amount}</p>
              <p>{row.categoryId}</p>
              <p>{row.description.slice(0, 10)}</p>
              <p>{row.flowId}</p>
              <p>{new Date(row.Date).toLocaleDateString("en-EN")}</p>
            </div>
          );
        })}
      </Table>
    </>
  );
};
