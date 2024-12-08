function BlogPosts({ html, state }) {
  const { store } = state;
  const { posts = [], offset, limit } = store;

  const cards = posts
    .slice(offset, offset + limit)
    .map((o, i) => `<blog-card key="${i + offset}"></blog-card>`)
    .join("");

  return html` <section class="mi-auto my-8 ">${cards}</section> `;
}
