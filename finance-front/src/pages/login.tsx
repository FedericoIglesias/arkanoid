import styled from "styled-components";
import { Context, ReqLogin, ResLogin } from "../vite-env";
import { useContext, useEffect } from "react";
import { ValueContext } from "../Context/valuesContext";
import { Email } from "@mui/icons-material";

const LoginSTY = styled.section`
  background-color: #222;
  height: 100vh;
  display: flex;
  justify-content: center;
  align-items: center;
  div {
    background-color: white;
    display: flex;
    flex-direction: column;
    padding: 10px 20px;
    box-shadow: 2px 2px 5px 2px #956403;
    label,
    input {
      margin-bottom: 5px;
    }
    input {
      border: 1px solid grey;
      padding: 1px 2px;
    }
    button {
      padding: 4px 7px;
      display: inline-block;
      margin: 0 auto;
    }
    button:hover {
      background-color: #956403;
    }
  }
`;

export const Login = () => {
  const login: ReqLogin = { email: "", password: "" };

  const { setValues }: { values: Context; setValues: any } =
    useContext<any>(ValueContext);

  const handlerSubmit = (e: any) => {
    e.preventDefault();

    fetch("http://localhost:3000/login", {
      method: "POST",
      mode: "cors",
      body: JSON.stringify(login),
    })
      .then((response) => response.text())
      .then((json) => {
        const data: ResLogin = JSON.parse(json);
        setValues((prevState: Context) => ({
          ...prevState,
          ID: data.ID,
          email: data.email,
          jwt: data.jwt,
        }));
      })
      .catch((e)=>{
        alert(e)
      });
  };

  return (
    <>
      <LoginSTY>
        <div>
          <label>Email</label>
          <input
            type="text"
            onChange={(e) => {
              login.email = e.target.value.trim();
            }}
          />
          <label htmlFor="">Password</label>
          <input
            type="password"
            onChange={(e) => {
              login.password = e.target.value;
            }}
          />
          <button onClick={(e) => handlerSubmit(e)}>Login</button>
        </div>
      </LoginSTY>
    </>
  );
};
