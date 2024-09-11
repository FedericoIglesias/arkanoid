const main = () => {
  fetch("http://localhost:3000/transaction/2", {
    'content-type': 'application/json; charset=utf-8',
    // method: "GET",
    // mode: "no-cors",
    // cache: "no-cache",
    // credentials: "same-origin",
  }).then(res => {
    console.log(res.text())
  }).then(data => {
    console.log(data);
  })
  // const data = await res.text()
  // const json = JSON.parse(data)
  // console.log(json);

}


main()

export { onBeforeRender }

import fetch from 'node-fetch'

async function onBeforeRender(pageContext) {
  // `.page.server.js` files always run in Node.js; we could use SQL/ORM queries here.
  const response = await fetch('https://movies.example.org/api')
  let movies = await response.json()

  // `movies` will be serialized and passed to the browser; we select only the data we
  // need in order to minimize what is sent to the browser.
  movies = movies.map(({ title, release_date }) => ({ title, release_date }))

  // We make `movies` available as `pageContext.pageProps.movies`
  const pageProps = { movies }
  return {
    pageContext: {
      pageProps
    }
  }
}