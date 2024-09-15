import * as React from "react";
import Box from "@mui/material/Box";
// import Button from "@mui/material/Button";
import Modal from "@mui/material/Modal";
import { Form } from "./form";
import styled from "styled-components";
import { PALETTE } from "../const";

const style = {
  position: "absolute" as "absolute",
  top: "50%",
  left: "50%",
  transform: "translate(-50%, -50%)",
};

const ButtonDiv = styled.div`
  display: inline-block;
  button{
    background-color: ${PALETTE.SECONDARY};
    padding: 2px 5px;
    border-radius: 5px;
  }
    :hover {
    background-color: ${PALETTE.THEIRD};
  }
`;

export const Modals = ({ tag }: { tag: string }) => {
  const [open, setOpen] = React.useState(false);
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
          {/* <Typography id="modal-modal-title" variant="h6" component="h2">
            Text in a modal
          </Typography>
          <Typography id="modal-modal-description" sx={{ mt: 2 }}>
            Duis mollis, est non commodo luctus, nisi erat porttitor ligula.
          </Typography> */}
          <Form></Form>
        </Box>
      </Modal>
    </div>
  );
};
