import { useEffect, useState } from "react";
import { Form } from "./form";

export const Table = ({ listHead }: { listHead: string[] }) => {
  const listBody: any[] = [1, 2, 3, 4, 5];

  const [data, setData] = useState(null);

  useEffect(() => {
    fetch("http://localhost:3000/transaction/2", {
      mode: 'cors'
    })
      .then((response) => response.json())
      .then((json) => console.log(json))
      .catch((error) => {
        console.error("Error al hacer la petición:", error);
      });
  }, []);

  return (
    <>
      <Form listBody={listBody} />
      <table>
        <thead>
          <tr id="1">
            {listHead.map((tag: string) => {
              return <th>{tag}</th>;
            })}
          </tr>
        </thead>
        <tbody>
          {listBody.map((data) => {
            return (
              <tr id={data}>
                <td>{data}</td>
                <td>{data}</td>
                <td>{data}</td>
                <td>{data}</td>
              </tr>
            );
          })}
        </tbody>
      </table>
    </>
  );
};
