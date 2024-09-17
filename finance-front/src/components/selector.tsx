import { Pagination } from "@mui/material";

export function Selector({
  setRowsPerPage,
  page,
  rowsPerPage,
  setNumberPage,
}: {
  setRowsPerPage: any;
  page: number;
  rowsPerPage: number;
  setNumberPage: any;
}) {
  return (
    <section
      style={{
        width: "85%",
        margin: "0 auto",
        display: "flex",
        justifyContent: "space-between",
        alignItems: "center",
        padding: "1px 10px",
        fontSize: "10px",
      }}
    >
      <div>
        <label>Rows per page:</label>
        <select
          style={{ fontSize: "10px", backgroundColor: "#CCC" }}
          value={rowsPerPage}
          onChange={(e) => setRowsPerPage(e.target.value)}
        >
          <option value={10}>10</option>
          <option value={25}>25</option>
          <option value={50}>50</option>
          <option value={100}>100</option>
        </select>
      </div>
      <Pagination
        count={page}
        size="small"
        onChange={(event, nPage) => setNumberPage(nPage)}
      />
      {/* <Pagination color='primary' style={{ margin: ' 10px auto', backgroundColor:'white' }} shape='rounded' count={page} variant="outlined" onChange={(event, nPage) => setNp(nPage)} /> */}
    </section>
  );
}
