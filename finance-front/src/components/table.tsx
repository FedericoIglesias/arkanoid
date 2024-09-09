// import { lenguage, tags } from "../const";
import { useEffect } from "react";
import { Form } from "./form";

export const Table = ({ listHead }: { listHead: string[] }) => {
  const listBody: any[] = [1, 2, 3, 4, 5];

  return (
    <>
      <Form listBody={listBody} />
      <table>
        <thead>
          <tr>
            {listHead.map((tag: string) => {
              return <th>{tag}</th>;
            })}
            {/* <th>{tags.AMOUNT[lenguage]}</th>
          <th>{tags.CATEGORY[lenguage]}</th>
          <th>{tags.DESCRIPTION[lenguage]}</th>
          <th>{tags.FLOW[lenguage]}</th> */}
          </tr>
        </thead>
        <tbody>
          {listBody.map((data) => {
            return (
              <tr>
                <td>1</td>
                <td>2</td>
                <td>3</td>
                <td>4</td>
              </tr>
            );
          })}
        </tbody>
      </table>
    </>
  );
};
