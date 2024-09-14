import { useEffect, useState } from "react";
import styled from "styled-components";
import { PALETTE } from "../const";

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

  PALETTE;

  const TableSty = styled.table`
    color: #333;
    margin: 0 auto;
    border-spacing: 0;
    -webkit-border-vertical-spacing: 0.2cap;
    thead, tbody{
      background-color: ${PALETTE.SECONDARY};
    }
    th,
    td {
      padding: 5px 15px;
    }
  `;

  return (
    <>
      {/* <Form /> */}
      <TableSty>
          <thead>
            {listHead.map((tag: string) => {
              return <th key={tag}>{tag}</th>;
            })}
          </thead>
        <tbody>
          {data.map((raw: getRaw) => {
            return (
              <tr key={JSON.stringify(raw.id)}>
                <td>{raw.amount}</td>
                <td>{raw.categoryId}</td>
                <td>{raw.description.slice(0, 6)}</td>
                <td>{raw.flowId}</td>
              </tr>
            );
          })}
        </tbody>
      </TableSty>
    </>
  );
};
