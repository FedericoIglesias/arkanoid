import { useEffect, useState } from "react";
import { Form } from "./form";

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
  const listBody: any[] = [1, 2, 3, 4, 5];

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

  return (
    <>
      <Form listBody={listBody} />
      <table>
        <thead>
          <tr key={1}>
            {listHead.map((tag: string) => {
              return <th>{tag}</th>;
            })}
          </tr>
        </thead>
        <tbody>
          {data.map((raw: getRaw) => {
            return (
              <tr key={JSON.stringify(raw.id)}>
                <td>{raw.amount}</td>
                <td>{raw.categoryId}</td>
                <td>{raw.description.slice(0,10)}</td>
                <td>{raw.flowId}</td>
              </tr>
            );
          })}
        </tbody>
      </table>
    </>
  );
};
