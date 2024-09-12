const transaction = {
  userId: 2,
  amount: 1,
  description: "proof from App.js",
  categoryID: 1,
  flowID: 2
};

fetch("http://localhost:3000/transaction", {
  method: "POST",
  mode: 'cors',
  cache: 'no-cache',
  credentials: 'omit',
  headers: {
    "content-type": "application/json",
  },
  body: JSON.stringify(transaction),
})
  .then((response) => {
    console.log(response);
  })
  .then((json) => {
    console.log(json);
  })
  .catch((err) => {
    console.log("Error: " + err);
  });