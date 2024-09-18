import { useContext, useEffect, useState } from "react";
import styled from "styled-components";
import { Selector } from "./selector";
import { Category, Context } from "../vite-env";
import { ValueContext } from "../Context/valuesContext";
import { tags } from "../const";

type getRaw = {
  id: number;
  amount: number;
  description: string;
  categoryId: number;
  userId: number;
  flowId: number;
  Date: string;
};

const getDateFormat = (values: Context) => {
  if (values.language == 0) {
    return "en-EN";
  }
  return "es-ES";
};

export const Table = ({ listHead }: { listHead: string[] }) => {
  const [data, setData] = useState<any>([]);
  const [listCategory, setListCategory] = useState<[Category] | null>(null);

  const { values }: { values: Context } = useContext<any>(ValueContext);

  useEffect(() => {
    fetch("http://localhost:3000/transaction/2", {
      mode: "cors",
    })
      .then((response) => response.json())
      .then((json) => {
        const listData = json.data.sort(function (a: getRaw, b: getRaw) {
          if (a.Date < b.Date) {
            return 1;
          }
          if (a.Date > b.Date) {
            return -1;
          }
          // a must be equal to b
          return 0;
        });
        setData(listData);
      })
      .catch((error) => {
        console.error("Error al hacer la petición:", error);
      });
  }, []);

  useEffect(() => {
    fetch("http://localhost:3000/categories", {
      mode: "cors",
    })
      .then((response) => response.json())
      .then((json) => setListCategory(json.data))
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
  const Thead = styled.div`
    font-size: 10px;
    max-width: 85%;
    max-height: 80vh;
    margin: 0 auto;
    background-color: #fff;
    margin-bottom: 2px;
    display: flex;
    border-radius: 10px;
    p {
      width: calc(100% / ${listHead.length});
      padding: 3px 3px;
    }
  `;
  const Table = styled.section`
    font-size: 10px;
    max-width: 85%;
    max-height: 75vh;
    margin: 0 auto;
    overflow-y: auto;
    div {
      background-color: #fff;
      margin-bottom: 2px;
      display: flex;
      border-radius: 10px;
      align-items: center;
      position: relative;
      padding: 3px 0;
      p {
        width: calc(100% / ${listHead.length});
        padding: 2px 0px 2px 6px;
      }
    }
  `;

  const putCategory = (id: number) => {
    let name = "";
    if (listCategory == null) {
      return "Wait request";
    }
    listCategory.map((element: Category) => {
      if (element.ID == id) {
        return (name = element.nameCategory);
      }
    });
    return name.toUpperCase();
  };

  return (
    <>
      <Selector
        setRowsPerPage={setRowsPerPage}
        page={page}
        rowsPerPage={rowsPerPage}
        setNumberPage={setNumberPage}
      />
      <Thead>
        {listHead.map((tag: string) => {
          return <p key={tag}>{tag}</p>;
        })}
      </Thead>
      <Table>
        {dataRows.map((row: getRaw) => {
          return (
            <div key={JSON.stringify(row.id)}>
              <p>{row.amount}</p>
              <p>{putCategory(row.categoryId)}</p>
              <p>{row.description.slice(0, 10)}</p>
              <p>
                {row.flowId == 2
                  ? tags.INCOME[values.language]
                  : tags.OUTFLOW[values.language]}
              </p>
              <p>
                {new Date(row.Date).toLocaleDateString(getDateFormat(values))}
              </p>
            </div>
          );
        })}
      </Table>
    </>
  );
};
