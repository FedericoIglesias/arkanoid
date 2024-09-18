import Box from "@mui/material/Box";
import Modal from "@mui/material/Modal";
import { Form } from "./form";
import styled from "styled-components";
import { PALETTE } from "../const";
import { Children, useState } from "react";

const style = {
  position: "absolute" as "absolute",
  top: "50%",
  left: "50%",
  transform: "translate(-50%, -50%)",
};

const ButtonDiv = styled.div`
  display: inline-block;
  button {
    background-color: ${PALETTE.SECONDARY};
    padding: 2px 5px;
    border-radius: 5px;
  }
  :hover {
    background-color: ${PALETTE.THEIRD};
  }
`;

export const Modals = ({
  tag,
  children,
}: {
  tag: string;
  children: JSX.Element;
}) => {
  const [open, setOpen] = useState(false);
  const handleOpen = () => setOpen(true);
  const handleClose = () => setOpen(false);

  return (
    <div>
      <ButtonDiv>
        <button onClick={handleOpen}>{tag}</button>
      </ButtonDiv>
      <Modal
        open={open}
        onClose={handleClose}
        aria-labelledby="modal-modal-title"
        aria-describedby="modal-modal-description"
      >
        <Box sx={style}>
          {children}
        </Box>
      </Modal>
    </div>
  );
};
